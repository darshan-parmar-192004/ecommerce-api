export default defineNuxtRouteMiddleware((to) => {
  const token = useCookie('auth_token')
  
  if (!token.value) {
    return navigateTo(`/auth/login?redirect=${encodeURIComponent(to.fullPath)}`)
  }
})
