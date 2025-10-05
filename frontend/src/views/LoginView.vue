<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { RouterLink } from 'vue-router';

const authStore = useAuthStore();
const username = ref('');
const password = ref('');

const handleSubmit = () => {
  if (username.value && password.value) {
    authStore.login({ username: username.value, password: password.value });
  }
};
</script>

<template>
  <div class="auth-container">
    <div class="auth-card">
      <h2 class="auth-title">Welcome Back</h2>
      <p class="auth-subtitle">Log in to track your transactions.</p>
      <form @submit.prevent="handleSubmit" class="auth-form">
        <div v-if="authStore.error" class="error-message">{{ authStore.error }}</div>
        <div class="form-group">
          <label for="username">Username</label>
          <input id="username" type="text" v-model="username" required />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input id="password" type="password" v-model="password" required />
        </div>
        <button type="submit" class="submit-button" :disabled="authStore.isLoading">
          {{ authStore.isLoading ? 'Logging in...' : 'Log In' }}
        </button>
      </form>
      <p class="switch-form">
        Don't have an account? <RouterLink to="/register">Register here</RouterLink>
      </p>
    </div>
  </div>
</template>

<style scoped>
/* Add styles for auth views */
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: calc(100vh - 200px);
}
.auth-card {
  width: 100%;
  max-width: 400px;
  padding: 2.5rem;
  background-color: var(--bg-card);
  border-radius: var(--radius);
  box-shadow: var(--shadow-md);
  border: 1px solid var(--border-color);
}
.auth-title {
  font-size: 1.75rem;
  font-weight: 700;
  text-align: center;
  margin: 0 0 0.5rem;
}
.auth-subtitle {
  text-align: center;
  color: var(--text-secondary);
  margin: 0 0 2rem;
}
.auth-form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}
.form-group label {
  display: block;
  font-weight: 500;
  margin-bottom: 0.5rem;
}
input {
  width: 100%;
  padding: 0.75rem;
  font-size: 1rem;
  border-radius: 6px;
  border: 1px solid var(--border-color);
  background-color: var(--bg-main);
  box-sizing: border-box;
}
.submit-button {
  padding: 0.75rem;
  font-size: 1rem;
  font-weight: 600;
  color: #fff;
  background: var(--primary-gradient);
  border: none;
  border-radius: 6px;
  cursor: pointer;
  margin-top: 0.5rem;
}
.submit-button:disabled {
  opacity: 0.6;
  cursor: wait;
}
.error-message {
  color: var(--color-debit);
  background-color: #fef2f2;
  padding: 0.75rem;
  border-radius: 6px;
  text-align: center;
  font-weight: 500;
}
.switch-form {
  margin-top: 1.5rem;
  text-align: center;
  font-size: 0.9rem;
  color: var(--text-secondary);
}
.switch-form a {
  color: var(--primary-color-start);
  font-weight: 600;
  text-decoration: none;
}
</style>
