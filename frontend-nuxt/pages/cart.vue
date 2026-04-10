<script setup>
import { useCartStore } from '~/stores/cart'

const cartStore = useCartStore()
const router = useRouter()

useSeoMeta({
  title: 'Shopping Cart - E-Commerce Store'
})

const proceedToCheckout = () => {
  router.push('/checkout')
}
</script>

<template>
  <div class="min-h-screen bg-surface py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-on_surface font-display">Shopping Cart</h1>
        <p class="text-on_surface_variant mt-1">Manage your cart items</p>
      </div>

      <div v-if="cartStore.isEmpty" class="bg-surface-container-lowest rounded-xl shadow-ambient p-12 text-center">
        <svg class="mx-auto h-24 w-24 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <h2 class="mt-4 text-xl font-semibold text-on_surface">Your cart is empty</h2>
        <p class="mt-2 text-on_surface_variant">Add some products to get started!</p>
        <NuxtLink to="/products" class="inline-block mt-6 px-6 py-3 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90 transition-opacity font-medium">
          Browse Products
        </NuxtLink>
      </div>

      <div v-else class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div class="lg:col-span-2 space-y-4">
          <TransitionGroup name="cart-item">
            <div
              v-for="item in cartStore.items"
              :key="item.product.product_id"
              class="bg-surface-container-lowest rounded-xl shadow-ambient p-6 flex gap-6 hover:shadow-lg transition-all duration-300 group"
            >
              <div class="w-28 h-28 bg-gradient-to-br from-surface-container to-surface-container-high rounded-xl flex-shrink-0 flex items-center justify-center shadow-inner">
                <svg class="w-14 h-14 text-outline group-hover:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
              </div>

              <div class="flex-1 min-w-0">
                <NuxtLink
                  :to="`/products/${item.product.product_id}`"
                  class="font-semibold text-on_surface hover:text-primary line-clamp-1 text-lg transition-colors duration-200"
                >
                  {{ item.product.name }}
                </NuxtLink>
                <p class="text-primary font-bold text-lg mt-1">
                  ₹ {{ Number(item.product.price).toFixed(2) }}
                </p>

                <div class="flex items-center gap-4 mt-4">
                  <div class="flex items-center border border-outline-variant/30 rounded-lg bg-surface-container overflow-hidden shadow-inner">
                    <button
                      @click="cartStore.updateQuantity(item.product.product_id, item.quantity - 1)"
                      class="px-4 py-2 text-on_surface_variant hover:bg-surface-container-high hover:text-primary transition-all duration-200 hover:scale-110 active:scale-95"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
                      </svg>
                    </button>
                    <span class="px-4 py-2 font-medium min-w-[3rem] text-center bg-surface-container-low">{{ item.quantity }}</span>
                    <button
                      @click="cartStore.updateQuantity(item.product.product_id, item.quantity + 1)"
                      class="px-4 py-2 text-on_surface_variant hover:bg-surface-container-high hover:text-primary transition-all duration-200 hover:scale-110 active:scale-95"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                      </svg>
                    </button>
                  </div>

                  <button
                    @click="cartStore.removeItem(item.product.product_id)"
                    class="text-error hover:text-error/80 text-sm flex items-center gap-1 px-3 py-2 hover:bg-error-container/30 rounded-lg transition-all duration-200 hover:scale-105"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                    Remove
                  </button>

                  <button
                    @click="cartStore.saveForLater(item.product.product_id)"
                    class="text-on_surface_variant hover:text-primary text-sm flex items-center gap-1 px-3 py-2 hover:bg-surface-container rounded-lg transition-all duration-200 hover:scale-105"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                    </svg>
                    Save for Later
                  </button>
                </div>
              </div>

              <div class="text-right">
                <p class="text-xl font-bold text-on_surface">
                  ₹ {{ (Number(item.product.price) * item.quantity).toFixed(2) }}
                </p>
              </div>
            </div>
          </TransitionGroup>

          <!-- Saved for Later Section -->
          <div v-if="cartStore.savedCount > 0" class="mt-8">
            <h2 class="text-xl font-bold text-on_surface font-display mb-4">Saved for Later ({{ cartStore.savedCount }})</h2>
            <div class="space-y-4">
              <TransitionGroup name="cart-item">
                <div
                  v-for="item in cartStore.savedForLater"
                  :key="item.product.product_id"
                  class="bg-surface-container-lowest rounded-xl shadow-ambient p-6 flex gap-6 hover:shadow-lg transition-all duration-300 group"
                >
                  <div class="w-24 h-24 bg-gradient-to-br from-surface-container to-surface-container-high rounded-xl flex-shrink-0 flex items-center justify-center shadow-inner">
                    <svg class="w-12 h-12 text-outline group-hover:scale-110 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                    </svg>
                  </div>

                  <div class="flex-1 min-w-0">
                    <NuxtLink
                      :to="`/products/${item.product.product_id}`"
                      class="font-semibold text-on_surface hover:text-primary line-clamp-1 transition-colors duration-200"
                    >
                      {{ item.product.name }}
                    </NuxtLink>
                    <p class="text-primary font-bold mt-1">
                      ₹ {{ Number(item.product.price).toFixed(2) }}
                    </p>

                    <div class="flex items-center gap-3 mt-3">
                      <button
                        @click="cartStore.moveToCart(item.product.product_id)"
                        class="text-primary hover:text-primary/80 text-sm font-medium flex items-center gap-1 px-3 py-2 hover:bg-primary-container/20 rounded-lg transition-all duration-200"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
                        </svg>
                        Move to Cart
                      </button>

                      <button
                        @click="cartStore.removeFromSaved(item.product.product_id)"
                        class="text-error hover:text-error/80 text-sm flex items-center gap-1 px-3 py-2 hover:bg-error-container/30 rounded-lg transition-all duration-200"
                      >
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                        Remove
                      </button>
                    </div>
                  </div>

                  <div class="text-right">
                    <p class="font-bold text-on_surface">
                      ₹ {{ (Number(item.product.price) * item.quantity).toFixed(2) }}
                    </p>
                  </div>
                </div>
              </TransitionGroup>
            </div>
          </div>
        </div>

        <div class="lg:col-span-1">
          <div class="bg-surface-container-lowest/80 backdrop-blur-md rounded-xl shadow-ambient p-6 sticky top-24 border border-white/10">
            <h2 class="text-xl font-bold text-on_surface font-display mb-6">Order Summary</h2>

            <div class="space-y-4 text-base">
              <div class="flex justify-between text-on_surface_variant">
                <span>Subtotal ({{ cartStore.totalItems }} items)</span>
                <span class="font-medium text-on_surface">₹ {{ cartStore.subtotal.toFixed(2) }}</span>
              </div>
              <div class="flex justify-between text-on_surface_variant">
                <span>Tax (8%)</span>
                <span class="font-medium text-on_surface">₹ {{ cartStore.tax.toFixed(2) }}</span>
              </div>
              <div class="border-t border-outline-variant/20 pt-4 flex justify-between">
                <span class="font-bold text-lg text-on_surface">Total</span>
                <span class="font-bold text-2xl text-primary">₹ {{ cartStore.total.toFixed(2) }}</span>
              </div>
            </div>

            <button
              @click="proceedToCheckout"
              class="w-full mt-6 py-4 bg-gradient-to-r from-primary to-primary-container text-white rounded-xl hover:opacity-90 transition-all flex items-center justify-center gap-2 font-semibold text-lg shadow-lg shadow-primary/25 hover:scale-[1.02] active:scale-[0.98]"
            >
              Proceed to Checkout
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
              </svg>
            </button>

            <NuxtLink
              to="/products"
              class="block text-center text-primary hover:text-primary/80 mt-4 font-medium"
            >
              Continue Shopping
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.cart-item-enter-active {
  animation: slideIn 0.4s ease-out;
}

.cart-item-leave-active {
  animation: slideOut 0.3s ease-in;
}

.cart-item-move {
  transition: transform 0.4s ease;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateX(-30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes slideOut {
  from {
    opacity: 1;
    transform: translateX(0);
  }
  to {
    opacity: 0;
    transform: translateX(30px);
  }
}
</style>
