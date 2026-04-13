import { defineStore } from 'pinia'
import { shallowRef } from 'vue'

const TAX_RATE = 0.08

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: shallowRef([]),
    savedForLater: shallowRef([]),
    isOpen: false,
    _loaded: false
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
    isEmpty: (state) => state.items.length === 0,
    savedCount: (state) => state.savedForLater.length
  },

  actions: {
    addItem(product, quantity = 1) {
      const existing = this.items.find(item => item.product.product_id === product.product_id)
      if (existing) {
        existing.quantity += quantity
      } else {
        this.items = [...this.items, { product, quantity }]
      }
      this.persistCart()
    },

    removeItem(productId) {
      const index = this.items.findIndex(item => item.product.product_id === productId)
      if (index !== -1) {
        this.items = this.items.filter((_, i) => i !== index)
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
          this.items = [...this.items]
          this.persistCart()
        }
      }
    },

    clearCart() {
      this.items = []
      this.persistCart()
    },

    saveForLater(productId) {
      const index = this.items.findIndex(item => item.product.product_id === productId)
      if (index !== -1) {
        const [item] = this.items.splice(index, 1)
        this.savedForLater = [...this.savedForLater, item]
        this.persistCart()
      }
    },

    moveToCart(productId) {
      const index = this.savedForLater.findIndex(item => item.product.product_id === productId)
      if (index !== -1) {
        const [item] = this.savedForLater.splice(index, 1)
        this.items = [...this.items, item]
        this.persistCart()
      }
    },

    removeFromSaved(productId) {
      this.savedForLater = this.savedForLater.filter(item => item.product.product_id !== productId)
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
      if (import.meta.client) {
        localStorage.setItem('cart', JSON.stringify(this.items))
        localStorage.setItem('savedForLater', JSON.stringify(this.savedForLater))
        useCookie('cart').value = JSON.stringify(this.items)
      }
    },

    loadCart() {
      if (import.meta.client && !this._loaded) {
        let stored = localStorage.getItem('cart')
        if (!stored) {
          const cartCookie = useCookie('cart')
          stored = cartCookie.value
        }
        if (stored) {
          try {
            this.items = shallowRef(JSON.parse(stored))
          } catch {
            this.items = shallowRef([])
          }
        }

        const savedStored = localStorage.getItem('savedForLater')
        if (savedStored) {
          try {
            this.savedForLater = shallowRef(JSON.parse(savedStored))
          } catch {
            this.savedForLater = shallowRef([])
          }
        }
        this._loaded = true
      }
    },

    reloadCart() {
      if (import.meta.client) {
        this._loaded = false
        this.loadCart()
      }
    },

    init() {
      if (import.meta.client && !this._loaded) {
        this.loadCart()
      }
    }
  }
})
