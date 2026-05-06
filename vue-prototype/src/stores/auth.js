import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import authService from '@/services/authService'
import { useCartStore } from './cart'
import { STORAGE_KEYS } from '@/constants'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem(STORAGE_KEYS.TOKEN) || null)
  const user = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Initialize user from localStorage safely (call explicitly when needed)
  const initUser = () => {
    try {
      const userStr = localStorage.getItem(STORAGE_KEYS.USER)
      if (userStr) {
        user.value = JSON.parse(userStr)
      }
    } catch (error) {
      console.warn('Failed to parse user from localStorage:', error)
      user.value = null
      localStorage.removeItem(STORAGE_KEYS.USER)
    }
  }

  // Initialize user on store creation (safe to call)
  initUser()

  const isAuthenticated = computed(() => !!token.value)
  const isAdmin = computed(() => user.value?.role === 'admin')

  const setToken = (newToken) => {
    token.value = newToken
    localStorage.setItem(STORAGE_KEYS.TOKEN, newToken)
  }

  const setUser = (userData) => {
    user.value = userData
    try {
      localStorage.setItem(STORAGE_KEYS.USER, JSON.stringify(userData))
    } catch (err) {
      console.error('Failed to save user to localStorage:', err)
    }
  }

  const login = async (credentials) => {
    loading.value = true
    error.value = null
    try {
      const { data } = await authService.login(credentials)
      setToken(data.token)
      setUser(data.customer)
      return data
    } catch (err) {
      error.value = err.response?.data?.message || 'Login failed'
      throw err
    } finally {
      loading.value = false
    }
  }

  const register = async (userData) => {
    loading.value = true
    error.value = null
    try {
      const { data } = await authService.register(userData)
      setToken(data.token)
      setUser(data.customer)
      return data
    } catch (err) {
      error.value = err.response?.data?.message || 'Registration failed'
      throw err
    } finally {
      loading.value = false
    }
  }

  const fetchUser = async () => {
    try {
      const { data } = await authService.getProfile()
      setUser(data)
    } catch (err) {
      error.value = 'Failed to fetch user profile'
    }
  }

  const logout = async () => {
    try {
      await authService.logout()
    } catch (err) {
      console.error('Logout failed', err)
    } finally {
      token.value = null
      user.value = null
      localStorage.removeItem(STORAGE_KEYS.TOKEN)
      localStorage.removeItem(STORAGE_KEYS.USER)
      const cartStore = useCartStore()
      cartStore.clearCart()
    }
  }

  return {
    token,
    user,
    loading,
    error,
    isAuthenticated,
    isAdmin,
    login,
    register,
    logout,
    fetchUser,
    setToken
  }
})

