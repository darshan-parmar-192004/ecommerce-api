export default defineNuxtPlugin(async () => {
  const authStore = useAuthStore()
  
  // On server, just initialize from cookies
  if (import.meta.server) {
    authStore.loadAuthFromCookie()
  }
})