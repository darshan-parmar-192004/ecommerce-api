<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useCart } from '@/composables/useCart'
import CartItem from '@/components/cart/CartItem.vue'
import { ShoppingBag, ArrowLeft } from 'lucide-vue-next'

const router = useRouter()
const { items, cartCount, cartTotal, clearCart } = useCart()

const hasItems = computed(() => items.value.length > 0)
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-8">Shopping Cart</h1>

    <div v-if="!hasItems" class="text-center py-16">
      <ShoppingBag class="w-16 h-16 text-gray-400 dark:text-gray-500 mx-auto mb-4" />
      <p class="text-gray-500 dark:text-gray-400 mb-4">Your cart is empty</p>
      <RouterLink to="/" class="btn-primary">
        Continue Shopping
      </RouterLink>
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <div class="lg:col-span-2">
        <div class="flex items-center justify-between mb-4">
          <span class="text-sm text-gray-600 dark:text-gray-400">{{ cartCount }} items</span>
          <button
            @click="clearCart()"
            class="text-sm text-red-600 hover:text-red-700 dark:text-red-400 dark:hover:text-red-300 transition-colors"
          >
            Clear Cart
          </button>
        </div>

        <TransitionGroup name="cart-item" tag="div" class="space-y-4">
          <CartItem
            v-for="item in items"
            :key="item.id"
            :item="item"
          />
        </TransitionGroup>

        <RouterLink to="/" class="inline-flex items-center gap-2 text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-100 mt-6">
          <ArrowLeft class="w-4 h-4" />
          Continue Shopping
        </RouterLink>
      </div>

      <div class="bg-white dark:bg-brand-800 p-6 rounded-xl border border-gray-200 dark:border-brand-700 h-fit">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-4">Order Summary</h2>
        <div class="space-y-3 mb-6">
          <div class="flex justify-between text-gray-600 dark:text-gray-400">
            <span>Subtotal ({{ cartCount }} items)</span>
            <span>₹{{ cartTotal.toFixed(2) }}</span>
          </div>
          <div class="flex justify-between text-gray-600 dark:text-gray-400">
            <span>Shipping</span>
            <span class="text-green-600 dark:text-green-400">Free</span>
          </div>
          <hr class="border-gray-200 dark:border-brand-700" />
          <div class="flex justify-between text-lg font-semibold text-gray-900 dark:text-gray-100">
            <span>Total</span>
            <span>₹{{ cartTotal.toFixed(2) }}</span>
          </div>
        </div>
        <button @click="router.push('/checkout')" class="btn-primary w-full">
          Proceed to Checkout
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.cart-item-enter-active,
.cart-item-leave-active {
  transition: all 0.3s ease;
}
.cart-item-enter-from {
  opacity: 0;
  transform: translateX(20px);
}
.cart-item-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}
.cart-item-move {
  transition: transform 0.3s ease;
}
</style>
