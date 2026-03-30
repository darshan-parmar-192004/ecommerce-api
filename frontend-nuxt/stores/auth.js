import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: null,
    user: null,
    loading: false
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    isAdmin: (state) => state.user?.role === 'admin'
  },

  actions: {
    setAuth(data) {
      this.token = data.token
      this.user = data.customer
      this.persistAuth()
    },

    setToken(token) {
      this.token = token
      this.persistAuth()
    },

    setUser(user) {
      this.user = user
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
      }
    },

    clearAuthStorage() {
      if (import.meta.client) {
        document.cookie = 'auth_token=; path=/; max-age=0'
        localStorage.removeItem('auth_user')
      }
    }
  }
})
