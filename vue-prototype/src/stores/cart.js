import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useCartStore = defineStore('cart', () => {
  const items = ref(JSON.parse(localStorage.getItem('cart') || '[]'))
  const isDrawerOpen = ref(false)

  const cartCount = computed(() => items.value.reduce((sum, item) => sum + item.quantity, 0))
  const cartTotal = computed(() => items.value.reduce((sum, item) => sum + (item.price * item.quantity), 0))

  const saveCart = () => localStorage.setItem('cart', JSON.stringify(items.value))

  const addToCart = (product, quantity = 1) => {
    const existing = items.value.find(item => item.id === product.product_id)
    if (existing) {
      existing.quantity += quantity
    } else {
      items.value.push({
        id: product.product_id,
        name: product.name,
        price: product.price,
        image: product.images?.[0] || '',
        quantity
      })
    }
    saveCart()
  }

  const removeFromCart = (productId) => {
    items.value = items.value.filter(item => item.id !== productId)
    saveCart()
  }

  const updateQuantity = (productId, quantity) => {
    const item = items.value.find(item => item.id === productId)
    if (item) {
      item.quantity = Math.max(1, quantity)
      saveCart()
    }
  }

  const toggleDrawer = () => isDrawerOpen.value = !isDrawerOpen.value
  const clearCart = () => { items.value = []; saveCart() }

  return {
    items,
    isDrawerOpen,
    cartCount,
    cartTotal,
    addToCart,
    removeFromCart,
    updateQuantity,
    toggleDrawer,
    clearCart
  }
})
