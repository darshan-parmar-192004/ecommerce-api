<script setup>
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
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900">Shopping Cart</h1>
        <p class="text-gray-600 mt-1">Manage your cart items</p>
      </div>

      <div v-if="cartStore.isEmpty" class="bg-white rounded-xl shadow-sm p-12 text-center">
        <svg class="mx-auto h-24 w-24 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
        </svg>
        <h2 class="mt-4 text-xl font-semibold text-gray-900">Your cart is empty</h2>
        <p class="mt-2 text-gray-500">Add some products to get started!</p>
        <NuxtLink to="/products" class="inline-block mt-6 px-6 py-3 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-colors font-medium">
          Browse Products
        </NuxtLink>
      </div>

      <div v-else class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div class="lg:col-span-2 space-y-4">
          <div 
            v-for="item in cartStore.items" 
            :key="item.product.product_id"
            class="bg-white rounded-xl shadow-sm p-6 flex gap-6 hover:shadow-md transition-shadow"
          >
            <div class="w-28 h-28 bg-gradient-to-br from-gray-50 to-gray-100 rounded-xl flex-shrink-0 flex items-center justify-center">
              <svg class="w-14 h-14 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
            
            <div class="flex-1 min-w-0">
              <NuxtLink 
                :to="`/products/${item.product.product_id}`"
                class="font-semibold text-gray-900 hover:text-indigo-600 line-clamp-1 text-lg"
              >
                {{ item.product.name }}
              </NuxtLink>
              <p class="text-indigo-600 font-bold text-lg mt-1">
                ₹ {{ Number(item.product.price).toFixed(2) }}
              </p>
              
              <div class="flex items-center gap-4 mt-4">
                <div class="flex items-center border border-gray-200 rounded-lg bg-gray-50">
                  <button 
                    @click="cartStore.updateQuantity(item.product.product_id, item.quantity - 1)"
                    class="px-4 py-2 text-gray-600 hover:bg-gray-100 hover:text-indigo-600 transition-colors rounded-l-lg"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4" />
                    </svg>
                  </button>
                  <span class="px-4 py-2 font-medium min-w-[3rem] text-center">{{ item.quantity }}</span>
                  <button 
                    @click="cartStore.updateQuantity(item.product.product_id, item.quantity + 1)"
                    class="px-4 py-2 text-gray-600 hover:bg-gray-100 hover:text-indigo-600 transition-colors rounded-r-lg"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                    </svg>
                  </button>
                </div>
                
                <button 
                  @click="cartStore.removeItem(item.product.product_id)"
                  class="text-red-500 hover:text-red-700 text-sm flex items-center gap-1 px-3 py-2 hover:bg-red-50 rounded-lg transition-colors"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                  Remove
                </button>
              </div>
            </div>

            <div class="text-right">
              <p class="text-xl font-bold text-gray-900">
                ₨{{ (Number(item.product.price) * item.quantity).toFixed(2) }}
              </p>
            </div>
          </div>
        </div>

        <div class="lg:col-span-1">
          <div class="bg-white rounded-xl shadow-sm p-6 sticky top-24">
            <h2 class="text-xl font-bold text-gray-900 mb-6">Order Summary</h2>
            
            <div class="space-y-4 text-base">
              <div class="flex justify-between text-gray-600">
                <span>Subtotal ({{ cartStore.totalItems }} items)</span>
                <span class="font-medium">₹ {{ cartStore.subtotal.toFixed(2) }}</span>
              </div>
              <div class="flex justify-between text-gray-600">
                <span>Tax (8%)</span>
                <span class="font-medium">₹ {{ cartStore.tax.toFixed(2) }}</span>
              </div>
              <div class="border-t pt-4 flex justify-between">
                <span class="font-bold text-lg">Total</span>
                <span class="font-bold text-2xl text-indigo-600">₹ {{ cartStore.total.toFixed(2) }}</span>
              </div>
            </div>

            <button 
              @click="proceedToCheckout"
              class="w-full mt-6 py-4 bg-indigo-600 text-white rounded-xl hover:bg-indigo-700 transition-all flex items-center justify-center gap-2 font-semibold text-lg hover:scale-[1.02] active:scale-[0.98]"
            >
              Proceed to Checkout
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
              </svg>
            </button>
            
            <NuxtLink 
              to="/products" 
              class="block text-center text-indigo-600 hover:text-indigo-700 mt-4 font-medium"
            >
              Continue Shopping
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>