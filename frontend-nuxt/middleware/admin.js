export default defineNuxtRouteMiddleware((to) => {
  const authStore = useAuthStore()
  if (import.meta.server) {
    authStore.loadAuthFromCookie()
  }
  if (!authStore.isAuthenticated || !authStore.token) {
    return navigateTo(`/auth/login?redirect=${encodeURIComponent(to.fullPath)}`)
  }
  if (!authStore.isAdmin) {
    throw createError({
      statusCode: 403,
      message: 'Access denied. Admin privileges required.'
    })
  }
})