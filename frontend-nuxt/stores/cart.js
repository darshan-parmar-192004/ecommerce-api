import { defineStore } from 'pinia'
import { useAuthStore } from './auth'

const TAX_RATE = 0.08

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: [],
    isOpen: false
  }),

  getters: {
    totalItems: (state) => state.items.reduce((sum, item) => sum + item.quantity, 0),
    subtotal: (state) => state.items.reduce((sum, item) => sum + item.product.price * item.quantity, 0),
    tax() {
      return this.subtotal * TAX_RATE
    },
    total() {
      return this.subtotal + this.tax
    },
    isEmpty: (state) => state.items.length === 0
  },

  actions: {
    addItem(product, quantity = 1) {
      const authStore = useAuthStore()
      if (!authStore.isAuthenticated) {
        return
      }
      
      const existing = this.items.find(item => item.product.product_id === product.product_id)
      if (existing) {
        existing.quantity += quantity
      } else {
        this.items.push({ product, quantity })
      }
      this.persistCart()
    },

    removeItem(productId) {
      const index = this.items.findIndex(item => item.product.product_id === productId)
      if (index !== -1) {
        this.items.splice(index, 1)
        this.persistCart()
      }
    },

    updateQuantity(productId, quantity) {
      const item = this.items.find(item => item.product.product_id === productId)
      if (item) {
        if (quantity <= 0) {
          this.removeItem(productId)
        } else {
          item.quantity = quantity
          this.persistCart()
        }
      }
    },

    clearCart() {
      this.items = []
      this.persistCart()
    },

    toggleCart() {
      this.isOpen = !this.isOpen
    },

    openCart() {
      this.isOpen = true
    },

    closeCart() {
      this.isOpen = false
    },

    persistCart() {
      const authStore = useAuthStore()
      if (!authStore.isAuthenticated) {
        return
      }
      
      if (import.meta.client) {
        localStorage.setItem('cart', JSON.stringify(this.items))
      }
    },

    loadCart() {
      const authStore = useAuthStore()
      if (!authStore.isAuthenticated) {
        this.items = []
        return
      }
      
      if (import.meta.client) {
        const stored = localStorage.getItem('cart')
        if (stored) {
          try {
            this.items = JSON.parse(stored)
          } catch {
            this.items = []
          }
        }
      }
    }
  }
})
