export default defineNuxtPlugin(async () => {
  const authStore = useAuthStore()

  if (import.meta.client) {
    authStore.loadAuth()
    await authStore.verifyAuth()
  }
})
