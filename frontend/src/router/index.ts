import { createRouter, createWebHistory } from 'vue-router';
import TransactionsView from '@/views/TransactionsView.vue';
import { useAuthStore } from '@/stores/auth';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'transactions',
      component: TransactionsView,
      meta: { requiresAuth: true },
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/views/DashboardView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
      meta: { requiresGuest: true },
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('@/views/RegisterView.vue'),
      meta: { requiresGuest: true },
    },
  ],
});
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore();

  if (to.meta.requiresAuth && !authStore.isLoggedIn) {
    // Redirect to login if route requires auth and user is not logged in
    next({ name: 'login' });
  } else if (to.meta.requiresGuest && authStore.isLoggedIn) {
    // Redirect to home if route is for guests and user is logged in
    next({ name: 'transactions' });
  } else {
    // Otherwise, allow navigation
    next();
  }
});

export default router;
