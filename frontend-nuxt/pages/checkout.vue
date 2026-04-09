<script setup>
import { useCartStore } from '~/stores/cart'
import { useAuthStore } from '~/stores/auth'

definePageMeta({
  middleware: ['auth']
})

const cartStore = useCartStore()
const authStore = useAuthStore()
const router = useRouter()
const { orders: ordersApi } = useApi()

useSeoMeta({
  title: 'Checkout - E-Commerce Store'
})

const loading = ref(false)
const error = ref('')
const success = ref(false)

const form = reactive({
  shipping_address: ''
})

const handleSubmit = async () => {
  if (cartStore.isEmpty) {
    error.value = 'Your cart is empty'
    return
  }

  if (!form.shipping_address.trim()) {
    error.value = 'Please enter a shipping address'
    return
  }

  loading.value = true
  error.value = ''

  try {
    const orderData = {
      order: {
        shipping_address: form.shipping_address,
        total_amount: cartStore.total
      },
      items: cartStore.items.map(item => ({
        product_id: item.product.product_id,
        quantity: item.quantity,
        unit_price: item.product.price
      }))
    }

    const result = await ordersApi.create(orderData)

    cartStore.clearCart()
    success.value = true
    
    setTimeout(() => {
      router.push(`/orders/${result.order_id || result.data?.order_id}`)
    }, 1500)
  } catch (err) {
    error.value = err.message || 'Failed to create order. Please try again.'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-surface">
    <h1 class="text-3xl font-bold text-on_surface font-display mb-8">Checkout</h1>

    <div v-if="cartStore.isEmpty" class="text-center py-16 bg-surface-container-lowest rounded-xl shadow-ambient">
      <svg class="mx-auto h-24 w-24 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
      </svg>
      <h2 class="mt-4 text-xl font-medium text-on_surface">Your cart is empty</h2>
      <p class="mt-2 text-on_surface_variant">Add some products before checking out.</p>
      <NuxtLink to="/products" class="inline-block mt-6 px-6 py-3 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90">
        Browse Products
      </NuxtLink>
    </div>

    <div v-else-if="success" class="text-center py-16 bg-surface-container-lowest rounded-xl shadow-ambient">
      <svg class="mx-auto h-24 w-24 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
      </svg>
      <h2 class="mt-4 text-2xl font-medium text-on_surface">Order Placed Successfully!</h2>
      <p class="mt-2 text-on_surface_variant">Redirecting to your order...</p>
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div>
        <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
          <h2 class="text-lg font-semibold text-on_surface font-display mb-4">Shipping Information</h2>
          
          <div v-if="error" class="bg-error-container text-error p-3 rounded-lg text-sm mb-4 flex items-start gap-2">
            <svg class="w-5 h-5 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            {{ error }}
          </div>

          <form @submit.prevent="handleSubmit" class="space-y-4">
            <div>
              <label for="address" class="block text-sm font-medium text-on_surface_variant mb-2">
                Shipping Address <span class="text-error">*</span>
              </label>
              <textarea
                id="address"
                v-model="form.shipping_address"
                rows="4"
                required
                class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                placeholder="Enter your full shipping address including street, city, state, and zip code"
              />
            </div>

            <div class="pt-4">
              <button 
                type="submit"
                :disabled="loading"
                class="w-full py-4 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90 disabled:opacity-50 flex items-center justify-center gap-2 font-semibold shadow-lg shadow-primary/25"
              >
                <svg v-if="loading" class="animate-spin h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                {{ loading ? 'Processing Order...' : 'Place Order' }}
              </button>
            </div>
          </form>
        </div>
      </div>

      <div>
        <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
          <h2 class="text-lg font-semibold text-on_surface font-display mb-4">Order Summary</h2>
          
          <div class="space-y-4 max-h-80 overflow-y-auto">
            <div 
              v-for="item in cartStore.items" 
              :key="item.product.product_id"
              class="flex items-center gap-4 py-3 border-b border-outline-variant/20 last:border-0"
            >
              <div class="w-14 h-14 bg-surface-container rounded-lg flex-shrink-0 flex items-center justify-center">
                <svg class="w-8 h-8 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
              </div>
              
              <div class="flex-1 min-w-0">
                <p class="font-medium text-on_surface truncate">{{ item.product.name }}</p>
                <p class="text-sm text-on_surface_variant">Qty: {{ item.quantity }} × ₹ {{ Number(item.product.price).toFixed(2) }}</p>
              </div>

              <p class="font-semibold text-on_surface">
                ₹ {{ (Number(item.product.price) * item.quantity).toFixed(2) }}
              </p>
            </div>
          </div>

          <div class="mt-6 pt-4 border-t border-outline-variant/20 space-y-2">
            <div class="flex justify-between text-sm text-on_surface_variant">
              <span>Subtotal</span>
              <span class="font-medium text-on_surface">₹ {{ cartStore.subtotal.toFixed(2) }}</span>
            </div>
            <div class="flex justify-between text-sm text-on_surface_variant">
              <span>Tax (8%)</span>
              <span class="font-medium text-on_surface">₹ {{ cartStore.tax.toFixed(2) }}</span>
            </div>
            <div class="flex justify-between text-base pt-2 border-t border-outline-variant/20">
              <span class="font-semibold text-on_surface">Total</span>
              <span class="font-bold text-xl text-primary">₹ {{ cartStore.total.toFixed(2) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
