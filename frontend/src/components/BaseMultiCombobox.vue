<script setup lang="ts">
import { computed, onMounted, onBeforeUnmount, ref, watch, nextTick } from 'vue';

type Option = { label: string; value: string };

const props = defineProps<{
  modelValue: string[];
  options: Option[];
  placeholder?: string;
  searchPlaceholder?: string;
  disabled?: boolean;
  maxMenuHeight?: number;
}>();

const emit = defineEmits<{ (e: 'update:modelValue', v: string[]): void }>();

const open = ref(false);
const query = ref('');
const rootEl = ref<HTMLElement | null>(null);
const inputEl = ref<HTMLInputElement | null>(null);
const listEl = ref<HTMLUListElement | null>(null);
const highlight = ref(-1);
const menuStyle = ref<Record<string, string>>({});

const normalized = computed(() => props.options.map(o => ({ ...o, key: o.value })));
const filtered = computed(() => {
  const q = query.value.trim().toLowerCase();
  if (!q) return normalized.value;
  return normalized.value.filter(o => o.label.toLowerCase().includes(q));
});
const selectedSet = computed(() => new Set(props.modelValue));

function toggleOpen(next?: boolean) {
  if (props.disabled) return;
  open.value = typeof next === 'boolean' ? next : !open.value;
  if (open.value) {
    nextTick(() => {
      reposition();
      inputEl.value?.focus();
      addGlobalListeners();
    });
  } else {
    removeGlobalListeners();
  }
}

function reposition() {
  const el = rootEl.value;
  if (!el) return;
  const r = el.getBoundingClientRect();
  const gutter = 8;
  const desiredMax = Math.max(160, Math.min(props.maxMenuHeight ?? 360, 360));
  const downSpace = window.innerHeight - r.bottom - gutter;
  const upSpace = r.top - gutter;
  const openUp = downSpace < 180 && upSpace > downSpace;
  const maxH = openUp ? Math.min(desiredMax, upSpace) : Math.min(desiredMax, downSpace);
  const top = openUp ? Math.max(8, r.top - maxH) : Math.min(window.innerHeight - maxH - 8, r.bottom + 4);

  menuStyle.value = {
    position: 'fixed',
    zIndex: '2000',
    left: `${r.left}px`,
    top: `${top}px`,
    width: `${r.width}px`,
    maxHeight: `${maxH}px`,
  };
}

function onWinEvent() { if (open.value) reposition(); }
function addGlobalListeners() {
  window.addEventListener('scroll', onWinEvent, true);
  window.addEventListener('resize', onWinEvent);
  window.addEventListener('orientationchange', onWinEvent);
}
function removeGlobalListeners() {
  window.removeEventListener('scroll', onWinEvent, true);
  window.removeEventListener('resize', onWinEvent);
  window.removeEventListener('orientationchange', onWinEvent);
}

function select(value: string) {
  if (selectedSet.value.has(value)) {
    emit('update:modelValue', props.modelValue.filter(v => v !== value));
  } else {
    emit('update:modelValue', [...props.modelValue, value]);
  }
}

function removeChip(value: string) {
  emit('update:modelValue', props.modelValue.filter(v => v !== value));
}

function clearAll() {
  emit('update:modelValue', []);
}

function onKeydown(e: KeyboardEvent) {
  if (!open.value && (e.key === 'ArrowDown' || e.key === 'Enter')) {
    e.preventDefault();
    return toggleOpen(true);
  }
  if (!open.value) return;

  if (e.key === 'ArrowDown') {
    e.preventDefault();
    highlight.value = (highlight.value + 1) % filtered.value.length;
    ensureVisible();
  } else if (e.key === 'ArrowUp') {
    e.preventDefault();
    highlight.value = (highlight.value - 1 + filtered.value.length) % filtered.value.length;
    ensureVisible();
  } else if (e.key === 'Enter') {
    e.preventDefault();
    const opt = filtered.value[highlight.value];
    if (opt) select(opt.value);
  } else if (e.key === 'Escape') {
    e.preventDefault();
    toggleOpen(false);
  }
}

function ensureVisible() {
  const list = listEl.value;
  if (!list) return;
  const item = list.children[highlight.value] as HTMLElement | undefined;
  if (!item) return;
  const top = item.offsetTop;
  const bottom = top + item.offsetHeight;
  if (top < list.scrollTop) list.scrollTop = top;
  else if (bottom > list.scrollTop + list.clientHeight) list.scrollTop = bottom - list.clientHeight;
}

function onClickAway(e: MouseEvent) {
  const root = rootEl.value;
  if (root && !root.contains(e.target as Node)) toggleOpen(false);
}

onMounted(() => document.addEventListener('mousedown', onClickAway));
onBeforeUnmount(() => {
  document.removeEventListener('mousedown', onClickAway);
  removeGlobalListeners();
});

watch(query, () => { highlight.value = 0; });
</script>

<template>
  <div class="mmc" :class="{ disabled }" ref="rootEl">
    <div class="mmc-input" @click="toggleOpen(true)">
      <div class="chips" v-if="modelValue.length">
        <span class="chip" v-for="v in modelValue" :key="v">
          {{ options.find(o => o.value === v)?.label ?? v }}
          <button type="button" class="x" @click.stop="removeChip(v)" aria-label="Remove">×</button>
        </span>
      </div>
      <input
        ref="inputEl"
        :placeholder="modelValue.length ? '' : (placeholder ?? 'Select...')"
        :disabled="disabled"
        v-model="query"
        @keydown="onKeydown"
      />
      <button type="button" class="caret" :disabled="disabled" @click.stop="toggleOpen()">
        ▼
      </button>
      <button v-if="modelValue.length" type="button" class="clear" @click.stop="clearAll" aria-label="Clear">✕</button>
    </div>

    <teleport to="body">
      <div class="mmc-menu" v-show="open" :style="menuStyle">
        <input
          class="search"
          :placeholder="searchPlaceholder ?? 'Search...'"
          v-model="query"
          @keydown="onKeydown"
        />
        <ul ref="listEl" class="list" role="listbox">
          <li
            v-for="(o, i) in filtered"
            :key="o.key"
            class="opt"
            :class="{ active: i === highlight, selected: selectedSet.has(o.value) }"
            role="option"
            :aria-selected="selectedSet.has(o.value)"
            @mouseenter="highlight = i"
            @mousedown.prevent="select(o.value)"
          >
            <input type="checkbox" :checked="selectedSet.has(o.value)" @change.prevent />
            <span>{{ o.label }}</span>
          </li>
          <li v-if="!filtered.length" class="empty">No matches</li>
        </ul>
      </div>
    </teleport>
  </div>
</template>

<style scoped>
.mmc { position: relative; font: inherit; }
.mmc.disabled { opacity: 0.6; pointer-events: none; }

.mmc-input {
  display: flex; align-items: center; gap: .5rem;
  border: 1px solid var(--border-color); border-radius: 6px;
  background: var(--bg-main); padding: .5rem .75rem;
}
.mmc-input input {
  border: none; outline: none; background: transparent; flex: 1; font: inherit;
}
.caret, .clear { background: transparent; border: none; cursor: pointer; color: var(--text-secondary); }

.chips { display: flex; gap: .375rem; flex-wrap: wrap; }
.chip {
  background: var(--bg-card); border: 1px solid var(--border-color);
  border-radius: 999px; padding: .125rem .5rem; font-size: .85rem;
  display: inline-flex; align-items: center; gap: .25rem;
}
.chip .x { background: transparent; border: none; cursor: pointer; color: var(--text-secondary); }

.mmc-menu {
  /* positioned via inline styles (fixed) for no clipping */
  border: 1px solid var(--border-color); border-radius: 8px; background: var(--bg-card);
  box-shadow: var(--shadow-md); overflow: auto;
}
.mmc-menu .search {
  width: 100%; padding: .5rem .75rem; border: none; border-bottom: 1px solid var(--border-color);
  outline: none; font: inherit; background: var(--bg-main);
}
.list { list-style: none; margin: 0; padding: .25rem; overflow: auto; }
.opt {
  display: flex; align-items: center; gap: .5rem;
  padding: .375rem .5rem; border-radius: 6px; cursor: pointer;
}
.opt.active { background: var(--bg-main); }
.opt.selected { font-weight: 600; }
.empty { padding: .75rem; color: var(--text-secondary); text-align: center; }
</style>
