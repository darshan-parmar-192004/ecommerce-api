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
        <i class="pi pi-shopping-cart text-6xl text-outline mb-4" />
        <h2 class="mt-4 text-xl font-semibold text-on_surface">Your cart is empty</h2>
        <p class="mt-2 text-on_surface_variant">Add some products to get started!</p>
        <NuxtLink to="/products">
          <Button label="Browse Products" icon="pi pi-shopping-bag" class="mt-6" />
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
                <i class="pi pi-box text-4xl text-outline group-hover:scale-110 transition-transform duration-300" />
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
                    <Button
                      @click="cartStore.updateQuantity(item.product.product_id, item.quantity - 1)"
                      icon="pi pi-minus"
                      text
                      class="!px-4 !py-2"
                      size="small"
                    />
                    <span class="px-4 py-2 font-medium min-w-[3rem] text-center bg-surface-container-low">{{ item.quantity }}</span>
                    <Button
                      @click="cartStore.updateQuantity(item.product.product_id, item.quantity + 1)"
                      icon="pi pi-plus"
                      text
                      class="!px-4 !py-2"
                      size="small"
                    />
                  </div>

                  <Button
                    @click="cartStore.removeItem(item.product.product_id)"
                    severity="danger"
                    text
                    size="small"
                  >
                    <i class="pi pi-trash mr-1" />
                    Remove
                  </Button>

                  <Button
                    @click="cartStore.saveForLater(item.product.product_id)"
                    severity="secondary"
                    text
                    size="small"
                  >
                    <i class="pi pi-bookmark mr-1" />
                    Save for Later
                  </Button>
                </div>
              </div>

              <div class="text-right">
                <p class="text-xl font-bold text-on_surface">
                  ₹ {{ (Number(item.product.price) * item.quantity).toFixed(2) }}
                </p>
              </div>
            </div>
          </TransitionGroup>

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
                    <i class="pi pi-box text-3xl text-outline group-hover:scale-110 transition-transform duration-300" />
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
                      <Button
                        @click="cartStore.moveToCart(item.product.product_id)"
                        severity="secondary"
                        size="small"
                      >
                        <i class="pi pi-shopping-cart mr-1" />
                        Move to Cart
                      </Button>
                      <Button
                        @click="cartStore.removeFromSaved(item.product.product_id)"
                        severity="danger"
                        text
                        size="small"
                      >
                        <i class="pi pi-trash mr-1" />
                        Remove
                      </Button>
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

            <Button
              @click="proceedToCheckout"
              label="Proceed to Checkout"
              icon="pi pi-arrow-right"
              iconPos="right"
              class="w-full !mt-6 !py-4"
            />

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
