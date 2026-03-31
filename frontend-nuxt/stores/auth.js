import { defineStore } from 'pinia'
import { useCartStore } from './cart'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: null,
    user: null,
    loading: false
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    isAdmin: (state) => state.user?.role === 'admin',
    userName: (state) => state.user?.name || 'User',
    userEmail: (state) => state.user?.email || ''
  },

  actions: {
    setAuth(data) {
      this.token = data.token
      this.user = data.customer
      this.persistAuth()
      
      const cartStore = useCartStore()
      cartStore.loadCart()
    },

    setToken(token) {
      this.token = token
      this.persistAuth()
    },

    setUser(user) {
      this.user = user
      this.persistAuth()
    },

    clearAuth() {
      this.token = null
      this.user = null
      this.clearAuthStorage()
    },

    persistAuth() {
      if (import.meta.client) {
        if (this.token) {
          document.cookie = `auth_token=${this.token}; path=/; max-age=86400; samesite=lax`
          localStorage.setItem('auth_token', this.token)
        }
        if (this.user) {
          localStorage.setItem('auth_user', JSON.stringify(this.user))
        }
      }
    },

    loadAuth() {
      if (import.meta.client) {
        const userStr = localStorage.getItem('auth_user')
        if (userStr) {
          try {
            this.user = JSON.parse(userStr)
          } catch {
            this.user = null
          }
        }
        
        const tokenFromStorage = localStorage.getItem('auth_token')
        if (tokenFromStorage) {
          this.token = tokenFromStorage
        }
        
        const tokenMatch = document.cookie.match(/auth_token=([^;]+)/)
        if (tokenMatch && !this.token) {
          this.token = tokenMatch[1]
        }
      }
    },

    clearAuthStorage() {
      if (import.meta.client) {
        document.cookie = 'auth_token=; path=/; max-age=0; samesite=lax'
        localStorage.removeItem('auth_token')
        localStorage.removeItem('auth_user')
      }
    },

    async verifyAuth() {
      if (!this.token) return false
      
      try {
        const response = await fetch('/api/auth/me', {
          headers: {
            'Authorization': `Bearer ${this.token}`
          }
        })
        
        if (response.ok) {
          const data = await response.json()
          this.user = data.data || data
          return true
        } else {
          this.clearAuth()
          return false
        }
      } catch (error) {
        console.error('Auth verification failed:', error)
        return false
      }
    }
  }
})
