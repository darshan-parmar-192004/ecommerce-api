import { ref, computed } from 'vue'
import { STORAGE_KEYS } from '@/constants'

// Module-level singleton state (shared across all calls to useCart)
const loadCart = () => {
  try {
    const cartStr = localStorage.getItem(STORAGE_KEYS.CART)
    return cartStr ? JSON.parse(cartStr) : []
  } catch (error) {
    console.error('Failed to load cart from localStorage:', error)
    return []
  }
}

const items = ref(loadCart())
const isDrawerOpen = ref(false)

const cartCount = computed(() => items.value.reduce((sum, item) => sum + item.quantity, 0))
const cartTotal = computed(() => items.value.reduce((sum, item) => sum + (item.price * item.quantity), 0))

const saveCart = () => {
  try {
    localStorage.setItem(STORAGE_KEYS.CART, JSON.stringify(items.value))
  } catch (error) {
    console.error('Failed to save cart to localStorage:', error)
  }
}

const addToCart = (product, quantity = 1) => {
  const existing = items.value.find(item => item.id === product.productId)
  if (existing) {
    existing.quantity += quantity
  } else {
    items.value.push({
      id: product.productId,
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

const toggleDrawer = () => {
  isDrawerOpen.value = !isDrawerOpen.value
}

const clearCart = () => {
  items.value = []
  saveCart()
}

/**
 * Cart composable - provides cart state and methods.
 * Uses module-level singleton state so all callers share the same reactive state.
 */
export function useCart() {
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
}
