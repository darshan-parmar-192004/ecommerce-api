<script setup>
import { RouterLink } from 'vue-router'
import { useCartStore } from '@/stores/cart'

const cartStore = useCartStore()
</script>

<template>
  <Transition name="drawer">
    <div v-if="cartStore.isDrawerOpen" class="fixed inset-0 z-50 overflow-hidden">
      <div class="absolute inset-0 bg-black/50" @click="cartStore.toggleDrawer()" />
      <div class="absolute right-0 top-0 h-full w-full max-w-md bg-white shadow-xl flex flex-col">
        <div class="flex items-center justify-between p-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">Shopping Cart</h2>
          <button @click="cartStore.toggleDrawer()" class="p-2 text-gray-500 hover:text-gray-700">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <div class="flex-1 overflow-y-auto p-4">
          <TransitionGroup v-if="cartStore.items.length > 0" name="cart-item" tag="div" class="space-y-4">
            <div
              v-for="item in cartStore.items"
              :key="item.id"
              class="flex gap-4 p-4 bg-gray-50 rounded-lg"
            >
              <img :src="item.image || '/placeholder.jpg'" :alt="item.name" class="w-20 h-20 object-cover rounded-md" />
              <div class="flex-1">
                <h3 class="text-sm font-medium text-gray-900">{{ item.name }}</h3>
                <p class="text-sm text-gray-600 mt-1">₹{{ (item.price * item.quantity).toFixed(2) }}</p>
                <div class="flex items-center gap-2 mt-2">
                  <button
                    @click="cartStore.updateQuantity(item.id, item.quantity - 1)"
                    class="w-6 h-6 flex items-center justify-center border border-gray-300 rounded hover:bg-gray-100"
                  >
                    -
                  </button>
                  <span class="text-sm w-8 text-center">{{ item.quantity }}</span>
                  <button
                    @click="cartStore.updateQuantity(item.id, item.quantity + 1)"
                    class="w-6 h-6 flex items-center justify-center border border-gray-300 rounded hover:bg-gray-100"
                  >
                    +
                  </button>
                  <button
                    @click="cartStore.removeFromCart(item.id)"
                    class="ml-auto text-sm text-red-600 hover:text-red-700"
                  >
                    Remove
                  </button>
                </div>
              </div>
            </div>
          </TransitionGroup>
          <div v-else class="text-center py-12 text-gray-500">
            Your cart is empty
          </div>
        </div>

        <div v-if="cartStore.items.length > 0" class="border-t border-gray-200 p-4 space-y-4">
          <div class="flex justify-between text-base font-semibold text-gray-900">
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
