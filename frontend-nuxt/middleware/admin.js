export default defineNuxtRouteMiddleware((to) => {
  // Run only on client side to avoid SSR blocking
  if (process.client) {
    const authStore = useAuthStore()
    if (!authStore.isAuthenticated || !authStore.token) {
      return navigateTo(`/auth/login?redirect=${encodeURIComponent(to.fullPath)}`)
    }
    if (!authStore.isAdmin) {
      throw createError({
        statusCode: 403,
        message: 'Access denied. Admin privileges required.'
      })
    }
  }
})