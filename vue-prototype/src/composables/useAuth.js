import { ref, computed } from 'vue'
import authService from '@/lib/authService'
import { useCart } from '@/composables/useCart'
import { STORAGE_KEYS } from '@/constants'

// Module-level singleton state (shared across all calls to useAuth)
const token = ref(localStorage.getItem(STORAGE_KEYS.TOKEN) || null)
const user = ref(null)
const loading = ref(false)
const error = ref(null)

// Initialize user from localStorage
const initUser = () => {
  try {
    const userStr = localStorage.getItem(STORAGE_KEYS.USER)
    if (userStr) {
      user.value = JSON.parse(userStr)
    }
  } catch (err) {
    console.warn('Failed to parse user from localStorage:', err)
    user.value = null
    localStorage.removeItem(STORAGE_KEYS.USER)
  }
}
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
    const { data: registerData } = await authService.register(userData)
    setToken(registerData.token)
    setUser(registerData.data)
    return registerData
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
    const { clearCart } = useCart()
    clearCart()
  }
}

/**
 * Auth composable - provides authentication state and methods.
 * Uses module-level singleton state so all callers share the same reactive state.
 */
export function useAuth() {
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
}
