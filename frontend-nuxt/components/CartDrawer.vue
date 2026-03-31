<script setup lang="ts">
const cartStore = useCartStore()
const { isAuthenticated } = useAuth()
const router = useRouter()

const checkout = async () => {
  if (!isAuthenticated.value) {
    router.push('/auth/login')
    return
  }
}
</script>

<template>
  <Teleport to="body">
    <Transition name="slide">
      <div
        v-if="cartStore.isOpen"
        class="fixed inset-0 z-50 overflow-hidden"
      >
        <div
          class="absolute inset-0 bg-black/30 backdrop-blur-sm"
          @click="cartStore.toggleCart"
        />

        <div class="absolute inset-y-0 right-0 max-w-md w-full bg-white shadow-xl flex flex-col">
          <div class="flex items-center justify-between p-4 border-b">
            <h2 class="text-lg font-semibold">Shopping Cart</h2>
            <button
              @click="cartStore.toggleCart"
              class="p-2 text-gray-400 hover:text-gray-600"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>

          <div class="flex-1 overflow-y-auto p-4">
            <template v-if="cartStore.isEmpty">
              <p class="text-center text-gray-500 py-8">
                Your cart is empty
              </p>
            </template>
            <template v-else>
              <div class="space-y-4">
                <div
                  v-for="item in cartStore.items"
                  :key="item.product.product_id"
                  class="flex gap-4 p-3 bg-gray-50 rounded-lg"
                >
                  <div class="w-20 h-20 bg-gray-200 rounded-lg flex-shrink-0" />

                  <div class="flex-1 min-w-0">
                    <h3 class="font-medium text-gray-900 truncate">
                      {{ item.product.name }}
                    </h3>
                    <p class="text-primary-600 font-semibold">
                      ₹ {{ item.product.price.toFixed(2) }}
                    </p>

                    <div class="flex items-center gap-2 mt-2">
                      <button
                        @click="cartStore.updateQuantity(item.product.product_id, item.quantity - 1)"
                        class="w-8 h-8 rounded bg-gray-200 hover:bg-gray-300 flex items-center justify-center transition-transform active:scale-90"
                      >
                        -
                      </button>
                      <span class="w-8 text-center font-medium">{{ item.quantity }}</span>
                      <button
                        @click="cartStore.updateQuantity(item.product.product_id, item.quantity + 1)"
                        class="w-8 h-8 rounded bg-gray-200 hover:bg-gray-300 flex items-center justify-center transition-transform active:scale-90"
                      >
                        +
                      </button>
                      <button
                        @click="cartStore.removeItem(item.product.product_id)"
                        class="ml-auto text-red-500 hover:text-red-700"
                      >
                        Remove
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </template>
          </div>

          <div v-if="!cartStore.isEmpty" class="border-t p-4 space-y-4">
            <div class="flex justify-between text-lg font-semibold">
              <span>Total</span>
              <span>₹ {{ cartStore.total.toFixed(2) }}</span>
            </div>
            <button
              @click="checkout"
              class="w-full py-3 px-4 bg-gradient-to-r from-primary-600 to-accent-600 text-white font-semibold rounded-xl hover:from-primary-700 hover:to-accent-700 transition-all shadow-lg shadow-primary-500/25"
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
  transition: transform 0.3s ease;
}

.slide-enter-from,
.slide-leave-to {
  opacity: 0;
}

.slide-enter-from > div:last-child,
.slide-leave-to > div:last-child {
  transform: translateX(100%);
}
</style>
