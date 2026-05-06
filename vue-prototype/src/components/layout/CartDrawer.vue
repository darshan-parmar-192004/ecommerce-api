<script setup>
import { RouterLink } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import CartItem from '@/components/cart/CartItem.vue'

const cartStore = useCartStore()
</script>

<template>
  <Transition name="drawer">
    <div v-if="cartStore.isDrawerOpen" class="fixed inset-0 z-50 overflow-hidden">
      <div class="absolute inset-0 bg-black/50" @click="cartStore.toggleDrawer()" />
      <div class="absolute right-0 top-0 h-full w-full sm:w-full sm:max-w-md md:max-w-lg bg-white dark:bg-brand-800 shadow-xl flex flex-col" role="dialog" aria-modal="true" aria-label="Shopping Cart">
        <div class="flex items-center justify-between p-4 border-b border-gray-200 dark:border-brand-700">
          <h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100">Shopping Cart</h2>
          <button @click="cartStore.toggleDrawer()" class="p-2 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200" aria-label="Close cart">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <div class="flex-1 overflow-y-auto p-4">
          <TransitionGroup v-if="cartStore.items.length > 0" name="cart-item" tag="div" class="space-y-4">
            <CartItem
              v-for="item in cartStore.items"
              :key="item.id"
              :item="item"
            />
          </TransitionGroup>
          <div v-else class="text-center py-12 text-gray-500 dark:text-gray-400">
            Your cart is empty
          </div>
        </div>

        <div v-if="cartStore.items.length > 0" class="border-t border-gray-200 dark:border-brand-700 p-4 space-y-4">
          <div class="flex justify-between text-base font-semibold text-gray-900 dark:text-gray-100">
            <span>Total</span>
            <span>₹{{ cartStore.cartTotal.toFixed(2) }}</span>
          </div>
          <RouterLink
            to="/checkout"
            @click="cartStore.toggleDrawer()"
            class="btn-primary w-full text-center"
          >
            Proceed to Checkout
          </RouterLink>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.drawer-enter-active,
.drawer-leave-active {
  transition: all 0.3s ease;
}
.drawer-enter-from,
.drawer-leave-to {
  opacity: 0;
}
.drawer-enter-from .absolute,
.drawer-leave-to .absolute {
  transform: translateX(100%);
}
.drawer-enter-active .absolute,
.drawer-leave-active .absolute {
  transition: transform 0.3s ease;
}

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
</style>
