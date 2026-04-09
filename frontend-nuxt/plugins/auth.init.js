export default defineNuxtPlugin(() => {
  const authStore = useAuthStore()
  authStore.loadAuthFromCookie()
})
