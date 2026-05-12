export default defineNuxtRouteMiddleware((to) => {
  const authStore = useAuthStore()

  if (authStore.isAuthenticated && authStore.token) {
    return navigateTo('/')
  }
})