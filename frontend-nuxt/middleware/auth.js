import { nextTick } from 'vue'

export default defineNuxtRouteMiddleware(async (to) => {
  const authStore = useAuthStore()
  const token = useCookie('auth_token')

  if (!token.value && !authStore.token) {
    await nextTick()
    if (token.value || authStore.token) return
    return navigateTo(`/auth/login?redirect=${encodeURIComponent(to.fullPath)}`)
  }
})
