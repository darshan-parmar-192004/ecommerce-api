import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    profile: null,
    orders: [],
    loading: false,
    ordersLoading: false
  }),

  getters: {
    hasProfile: (state) => !!state.profile,
    orderCount: (state) => state.orders.length
  },

  actions: {
    async fetchProfile() {
      this.loading = true
      try {
        const { auth } = useApi()
        const data = await auth.me()
        this.profile = data.data || data.customer || data
        return this.profile
      } catch (error) {
        console.error('Failed to fetch profile:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchOrders() {
      this.ordersLoading = true
      try {
        const { orders } = useApi()
        const data = await orders.list()
        this.orders = data.data || data || []
        return this.orders
      } catch (error) {
        console.error('Failed to fetch orders:', error)
        throw error
      } finally {
        this.ordersLoading = false
      }
    },

    async fetchOrder(id) {
      try {
        const { orders } = useApi()
        const data = await orders.get(id)
        return data.data || data
      } catch (error) {
        console.error('Failed to fetch order:', error)
        throw error
      }
    },

    async updateProfile(profileData) {
      this.loading = true
      try {
        const { fetchJson } = useApi()
        const data = await fetchJson('/customers/me', {
          method: 'PUT',
          body: JSON.stringify(profileData)
        })
        this.profile = { ...this.profile, ...data }
        return this.profile
      } catch (error) {
        console.error('Failed to update profile:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    clearUserData() {
      this.profile = null
      this.orders = []
    }
  }
})