import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: null,
    user: null,
    loading: false
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    isAdmin: (state) => {
      console.log('isAdmin check:', state.user?.role)
      return ['admin', 'Administrator', 'ADMIN'].includes(state.user?.role)
    },
    userName: (state) => state.user?.name || 'User',
    userEmail: (state) => state.user?.email || ''
  },

  actions: {
    async setAuth(data) {
      this.token = data.token
      const userFromResponse = data.data || data.customer || {}
      this.user = {
        customer_id: userFromResponse.customer_id || null,
        email:        userFromResponse.email        || null,
        name:         userFromResponse.name         || null,
        role:         userFromResponse.role || 'user'
      }
      console.log('setAuth - user with role:', this.user)
      this.persistAuth()

      if (import.meta.client) {
        const { useCartStore } = await import('./cart')
        const cartStore = useCartStore()
        cartStore.reloadCart()
      }
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
            const parsed = JSON.parse(userStr)
            this.user = {
              customer_id: parsed.customer_id || null,
              email:        parsed.email        || null,
              name:         parsed.name         || null,
              role:         parsed.role || 'user'
            }
            if (parsed.role) {
              console.log('Loaded role from localStorage:', parsed.role)
            }
          } catch (e) {
            console.error('Failed to parse stored user:', e)
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

    loadAuthFromCookie() {
      const tokenCookie = useCookie('auth_token')
      if (tokenCookie.value) {
        this.token = tokenCookie.value
      }
      
      const userStr = useCookie('auth_user')
      if (userStr.value) {
        try {
          this.user = typeof userStr.value === 'string' ? JSON.parse(userStr.value) : userStr.value
        } catch {
          this.user = null
        }
      }
    },

    clearAuthStorage() {
      if (import.meta.client) {
        document.cookie = 'auth_token=; path=/; max-age=0; samesite=lax'
        localStorage.removeItem('auth_token')
        localStorage.removeItem('auth_user')
      }
      
      const tokenCookie = useCookie('auth_token')
      tokenCookie.value = null
      
      const userCookie = useCookie('auth_user')
      userCookie.value = null
    },

    async verifyAuth() {
      const token = this.token || localStorage.getItem('auth_token')
      if (!token) return false

      try {
        const config = useRuntimeConfig()
        const apiBase = config.public.apiBase

        const response = await fetch(`${apiBase}/auth/me`, {
          headers: {
            'Authorization': `Bearer ${token}`
          }
        })

        if (response.ok) {
          const data = await response.json()
          this.token = token
          this.user = {
            customer_id: data.data?.customer_id || data.customer_id,
            email: data.data?.email || data.email,
            name: data.data?.name || data.name,
            role: data.data?.role || data.role
          }
          this.persistAuth()
          return true
        } else {
          this.clearAuth()
          return false
        }
      } catch (error) {
        console.error('Auth verification failed:', error)
        this.clearAuth()
        return false
      }
    },

    handleSessionExpiration() {
      this.clearAuth()
      navigateTo('/auth/login?message=session_expired')
    }
  }
})
