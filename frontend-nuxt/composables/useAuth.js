export const useAuth = () => {
  const authStore = useAuthStore()

  const token = computed(() => authStore.token)
  const user = computed(() => authStore.user)
  const isAuthenticated = computed(() => authStore.isAuthenticated)
  const isAdmin = computed(() => authStore.isAdmin)
  const loading = computed(() => authStore.loading)

  const setAuth = (data) => {
    authStore.setAuth(data)
  }

  const setToken = (newToken) => {
    authStore.setToken(newToken)
  }

  const setUser = (newUser) => {
    authStore.setUser(newUser)
  }

  const clearToken = () => {
    authStore.clearAuth()
  }

  const loadAuth = () => {
    authStore.loadAuth()
  }

  const getTokenFromCookie = (event) => {
    return getCookie(event, 'auth_token')
  }

  return {
    token,
    user,
    isAuthenticated,
    isAdmin,
    loading,
    setAuth,
    setToken,
    setUser,
    clearToken,
    loadAuth,
    getTokenFromCookie
  }
}
