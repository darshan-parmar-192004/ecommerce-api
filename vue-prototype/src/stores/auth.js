import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import authService from '@/services/authService'
import { useCartStore } from './cart'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || null)
  const user = ref(null)
  const loading = ref(false)
  const error = ref(null)
  
  // Initialize user from localStorage safely
  const initUser = () => {
    try {
      const userStr = localStorage.getItem('user')
      if (userStr) {
        user.value = JSON.parse(userStr)
      }
    } catch (error) {
      console.warn('Failed to parse user from localStorage:', error)
      user.value = null
      localStorage.removeItem('user')
    }
  }
  
  // Run initialization
  initUser()

  const isAuthenticated = computed(() => !!token.value)
  const isAdmin = computed(() => user.value?.role === 'admin')

  const setToken = (newToken) => {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  const setUser = (userData) => {
    user.value = userData
    localStorage.setItem('user', JSON.stringify(userData))
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
      console.error('Failed to fetch user profile', err)
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
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      const cartStore = useCartStore()
      cartStore.clearCart()
    }
  }

  if (isAuthenticated.value && !user.value) {
    fetchUser()
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
