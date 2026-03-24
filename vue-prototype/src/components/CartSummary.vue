<script setup>
import { inject, computed } from 'vue'

const cart = inject('cart')

const itemCount = computed(() => cart.items.length)

const formattedTotal = computed(() => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD'
  }).format(cart.total)
})

const handleRemoveItem = (productId) => {
  cart.removeItem(productId)
}
</script>

<template>
  <div class="bg-white rounded-xl shadow-lg border border-gray-200 overflow-hidden">
    <div class="bg-gradient-to-r from-gray-900 to-gray-800 px-5 py-4">
      <div class="flex items-center gap-3">
        <div class="p-2 bg-white/10 rounded-lg">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
          </svg>
        </div>
        <div>
          <h3 class="text-white font-semibold">Your Cart</h3>
          <p class="text-gray-400 text-xs">{{ itemCount }} {{ itemCount === 1 ? 'item' : 'items' }}</p>
        </div>
      </div>
    </div>
    
    <div class="p-5">
      <div v-if="cart.items.length === 0" class="text-center py-8">
        <div class="w-16 h-16 mx-auto mb-3 bg-gray-100 rounded-full flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
          </svg>
        </div>
        <p class="text-gray-500 font-medium">Your cart is empty</p>
        <p class="text-gray-400 text-sm mt-1">Add some products to get started</p>
      </div>
      
      <template v-else>
        <div class="max-h-[280px] overflow-y-auto space-y-2 -mx-2 px-2">
          <div 
            v-for="item in cart.items" 
            :key="item.cartId" 
            class="flex items-center justify-between py-2.5 px-3 bg-gray-50 rounded-lg hover:bg-gray-100 transition-colors group"
          >
            <div class="flex-1 min-w-0 mr-3">
              <p class="text-sm font-medium text-gray-900 truncate">{{ item.name }}</p>
              <p class="text-xs text-gray-500">${{ item.price.toFixed(2) }}</p>
            </div>
            <button 
              class="p-1.5 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors opacity-0 group-hover:opacity-100"
              @click="handleRemoveItem(item.id)"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </div>
        
        <div class="mt-4 pt-4 border-t border-gray-200">
          <div class="flex items-center justify-between mb-3">
            <span class="text-gray-500 font-medium">Subtotal</span>
            <span class="text-gray-900 font-semibold">{{ formattedTotal }}</span>
          </div>
          <div class="flex items-center justify-between mb-4">
            <span class="text-gray-500">Shipping</span>
            <span class="text-green-600 font-medium">Free</span>
          </div>
          <div class="flex items-center justify-between pt-3 border-t-2 border-gray-100">
            <span class="text-gray-900 font-bold">Total</span>
            <span class="text-2xl font-extrabold text-gray-900">{{ formattedTotal }}</span>
          </div>
        </div>

        <button class="w-full mt-4 py-3 bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 text-white font-semibold rounded-lg transition-all duration-200 shadow-lg hover:shadow-xl">
          Checkout Now
        </button>
      </template>
    </div>
  </div>
</template>
