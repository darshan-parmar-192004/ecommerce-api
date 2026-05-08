<script setup>
import { useCartStore } from '~/stores/cart'
import { useAuthStore } from '~/stores/auth'
import { X, ShoppingBag, Minus, Plus, Trash2 } from 'lucide-vue-next'

const cartStore = useCartStore()
const { isAuthenticated } = useAuthStore()
const router = useRouter()
const { error: showError } = useToast()

const checkout = async () => {
  if (!isAuthenticated.value) {
    router.push('/auth/login?redirect=/checkout')
    return
  }
  router.push('/checkout')
}

const formatPrice = (price) => {
  return `₹${Number(price).toFixed(2)}`
}
</script>

<template>
  <Teleport to="body">
    <Transition name="slide">
      <div
        v-if="cartStore.isOpen"
        class="fixed inset-0 z-50 overflow-hidden"
        @keyup.escape="cartStore.toggleCart"
      >
        <div
          class="absolute inset-0 bg-on_surface/30 backdrop-blur-md"
          @click="cartStore.toggleCart"
        />

        <div class="absolute inset-y-0 right-0 max-w-md w-full bg-surface-container-lowest/95 backdrop-blur-xl shadow-ambient flex flex-col border-l border-white/10">
          <!-- Header -->
          <div class="flex items-center justify-between p-6 border-b border-outline-variant/20">
            <div>
              <h2 class="text-xl font-bold text-on_surface font-display">Shopping Cart</h2>
              <p class="text-xs text-outline mt-1">
                {{ cartStore.totalItems }} item{{ cartStore.totalItems !== 1 ? 's' : '' }}
              </p>
            </div>
            <button
              @click="cartStore.toggleCart"
              class="p-2 rounded-full hover:bg-surface-container transition-all duration-200 hover:scale-110 active:scale-95"
            >
              <X class="w-6 h-6 text-outline" />
            </button>
          </div>

          <!-- Cart Items -->
          <div class="flex-1 overflow-y-auto p-6">
            <template v-if="cartStore.isEmpty">
              <div class="h-full flex flex-col items-center justify-center text-center">
                <div class="w-24 h-24 rounded-full bg-surface-container flex items-center justify-center mb-6 animate-bounce-slow">
                  <ShoppingBag class="w-12 h-12 text-outline/40" />
                </div>
                <h3 class="text-lg font-semibold text-on_surface font-display">Your cart is empty</h3>
                <p class="text-on_surface_variant mt-2 font-body">Add products to get started</p>
                <button
                  @click="cartStore.toggleCart"
                  class="mt-6 px-6 py-3 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg font-semibold transition-all hover:shadow-lg hover:scale-[1.02] active:scale-[0.98] shadow-primary/25"
                >
                  Continue Shopping
                </button>
              </div>
            </template>
            <template v-else>
              <div class="space-y-4">
                <TransitionGroup name="cart-item">
                  <div
                    v-for="item in cartStore.items"
                    :key="item.product.product_id"
                    class="flex gap-4 p-4 rounded-xl bg-surface-container-low/80 backdrop-blur-sm transition-all duration-300 hover:bg-surface-container hover:shadow-lg group"
                  >
                    <div class="w-24 h-24 bg-surface-container rounded-lg flex-shrink-0 overflow-hidden flex items-center justify-center shadow-inner">
                      <ShoppingBag class="w-12 h-12 text-outline/30 group-hover:scale-110 transition-transform duration-300" />
                    </div>

                    <div class="flex-1 min-w-0">
                      <h3 class="font-semibold text-on_surface font-display truncate group-hover:text-primary transition-colors duration-200">
                        {{ item.product.name }}
                      </h3>
                      <p class="text-primary font-bold mt-1 font-display">
                        {{ formatPrice(item.product.price) }}
                      </p>

<div class="flex items-center gap-2 mt-3">
                          <button
                            @click="cartStore.updateQuantity(item.product.product_id, item.quantity - 1)"
                            class="w-8 h-8 rounded-lg bg-surface-container hover:bg-surface-container-high flex items-center justify-center transition-all duration-200 hover:scale-110 active:scale-95"
                          >
                            <Minus class="w-4 h-4 text-on_surface" />
                          </button>
                          <span class="w-10 text-center font-medium text-on_surface bg-surface-container-low rounded">{{ item.quantity }}</span>
                          <button
                            @click="cartStore.updateQuantity(item.product.product_id, item.quantity + 1)"
                            class="w-8 h-8 rounded-lg bg-surface-container hover:bg-surface-container-high flex items-center justify-center transition-all duration-200 hover:scale-110 active:scale-95"
                          >
                            <Plus class="w-4 h-4 text-on_surface" />
                          </button>
                          <button
                            @click="cartStore.removeItem(item.product.product_id)"
                            class="ml-auto text-error hover:text-error/80 transition-all duration-200 hover:scale-105 text-sm font-medium flex items-center gap-1"
                          >
                            <Trash2 class="w-4 h-4" />
                            Remove
                          </button>
                        </div>
                    </div>
                  </div>
                </TransitionGroup>
              </div>
            </template>
          </div>

          <!-- Footer -->
          <div v-if="!cartStore.isEmpty" class="border-t border-outline-variant/20 p-6 space-y-4 bg-surface-container-low/80 backdrop-blur-md">
            <div class="space-y-2">
              <div class="flex justify-between text-sm text-on_surface_variant">
                <span>Subtotal</span>
                <span>{{ formatPrice(cartStore.subtotal) }}</span>
              </div>
              <div class="flex justify-between text-sm text-on_surface_variant">
                <span>Tax (8%)</span>
                <span>{{ formatPrice(cartStore.tax) }}</span>
              </div>
              <div class="flex justify-between text-lg font-bold text-on_surface pt-2 border-t border-outline-variant/20">
                <span>Total</span>
                <span class="text-primary">{{ formatPrice(cartStore.total) }}</span>
              </div>
            </div>
            <button
              @click="checkout"
              class="w-full py-4 px-6 bg-gradient-to-r from-primary to-primary-container text-white font-semibold rounded-xl hover:shadow-xl hover:shadow-primary/30 hover:scale-[1.02] active:scale-[0.98] transition-all shadow-lg shadow-primary/25"
            >
              Checkout
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: opacity 0.3s ease;
}

.slide-enter-active > div:last-child,
.slide-leave-active > div:last-child {
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-enter-from,
.slide-leave-to {
  opacity: 0;
}

.slide-enter-from > div:last-child,
.slide-leave-to > div:last-child {
  transform: translateX(100%);
}

.cart-item-enter-active {
  animation: slideInCart 0.4s ease-out;
}

.cart-item-leave-active {
  animation: slideOutCart 0.3s ease-in;
}

.cart-item-move {
  transition: transform 0.4s ease;
}

@keyframes slideInCart {
  from {
    opacity: 0;
    transform: translateX(30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes slideOutCart {
  from {
    opacity: 1;
    transform: translateX(0);
  }
  to {
    opacity: 0;
    transform: translateX(-30px);
  }
}

@keyframes bounce-slow {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

.animate-bounce-slow {
  animation: bounce-slow 2s ease-in-out infinite;
}
</style>
