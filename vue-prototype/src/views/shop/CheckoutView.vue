<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import { useUserStore } from '@/stores/user'
import { useToastStore } from '@/stores/toast'

const router = useRouter()
const cartStore = useCartStore()
const userStore = useUserStore()
const toastStore = useToastStore()

const step = ref(1)
const loading = ref(false)
const orderPlaced = ref(false)
const orderId = ref('N/A')

const shippingInfo = ref({
  address: '',
  city: '',
  state: '',
  zip: '',
  country: ''
})

const paymentInfo = ref({
  cardNumber: '',
  expiry: '',
  cvv: ''
})

onMounted(async () => {
  await userStore.fetchProfile()
  if (userStore.profile) {
    shippingInfo.value = {
      address: userStore.profile.address || '',
      city: '',
      state: '',
      zip: '',
      country: ''
    }
  }
})

const placeOrder = async () => {
  loading.value = true
  try {
    // Transform cart items to order items format matching backend
    const orderItems = cartStore.items.map(item => ({
      product_id: item.id,
      quantity: item.quantity,
      unit_price: item.price
    }))
    
    const orderData = {
      order: {
        customer_id: 'CUST-001', // Should come from auth
        total_amount: cartStore.cartTotal,
        status: 'pending',
        shipping_address: `${shippingInfo.value.address}, ${shippingInfo.value.city}, ${shippingInfo.value.state} ${shippingInfo.value.zip}`
      },
      items: orderItems
    }
    const result = await userStore.createOrder(orderData)
    orderId.value = result.order_id || 'N/A'
    orderPlaced.value = true
    cartStore.clearCart()
  } catch (err) {
    console.error('Order failed', err)
    toastStore.error('Failed to place order. Please try again.')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Checkout</h1>

    <div v-if="orderPlaced" class="text-center py-16">
      <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
        </svg>
      </div>
      <h2 class="text-2xl font-bold text-gray-900 mb-2">Order Placed!</h2>
      <p class="text-gray-600 mb-2">Thank you for your purchase.</p>
      <p class="text-sm text-gray-500 mb-6">Order ID: {{ orderId }}</p>
      <RouterLink to="/user/orders" class="btn-primary">
        View Order
      </RouterLink>
    </div>

    <div v-else class="space-y-8">
      <!-- Progress Steps -->
      <div class="flex items-center gap-4 mb-8">
        <div :class="step >= 1 ? 'bg-gray-900 text-white' : 'bg-gray-200'" class="w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium">
          1
        </div>
        <div class="flex-1 h-1 bg-gray-200">
          <div :class="step >= 2 ? 'bg-gray-900' : 'bg-gray-200'" class="h-full transition-all duration-300" :style="{ width: step >= 2 ? '100%' : '0%' }" />
        </div>
        <div :class="step >= 2 ? 'bg-gray-900 text-white' : 'bg-gray-200'" class="w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium">
          2
        </div>
        <div class="flex-1 h-1 bg-gray-200">
          <div :class="step >= 3 ? 'bg-gray-900' : 'bg-gray-200'" class="h-full transition-all duration-300" :style="{ width: step >= 3 ? '100%' : '0%' }" />
        </div>
        <div :class="step >= 3 ? 'bg-gray-900 text-white' : 'bg-gray-200'" class="w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium">
          3
        </div>
      </div>

      <!-- Step 1: Shipping -->
      <div v-if="step === 1" class="space-y-4">
        <h2 class="text-xl font-semibold text-gray-900">Shipping Information</h2>
        <input v-model="shippingInfo.address" placeholder="Address" class="input" />
        <div class="grid grid-cols-2 gap-4">
          <input v-model="shippingInfo.city" placeholder="City" class="input" />
          <input v-model="shippingInfo.state" placeholder="State" class="input" />
        </div>
        <div class="grid grid-cols-2 gap-4">
          <input v-model="shippingInfo.zip" placeholder="ZIP Code" class="input" />
          <input v-model="shippingInfo.country" placeholder="Country" class="input" />
        </div>
        <button @click="step = 2" class="btn-primary w-full mt-4">Continue to Payment</button>
      </div>

      <!-- Step 2: Payment -->
      <div v-else-if="step === 2" class="space-y-4">
        <h2 class="text-xl font-semibold text-gray-900">Payment Information</h2>
        <input v-model="paymentInfo.cardNumber" placeholder="Card Number" class="input" />
        <div class="grid grid-cols-2 gap-4">
          <input v-model="paymentInfo.expiry" placeholder="MM/YY" class="input" />
          <input v-model="paymentInfo.cvv" placeholder="CVV" class="input" />
        </div>
        <div class="flex gap-4 mt-4">
          <button @click="step = 1" class="btn-secondary flex-1">Back</button>
          <button @click="step = 3" class="btn-primary flex-1">Review Order</button>
        </div>
      </div>

      <!-- Step 3: Review -->
      <div v-else class="space-y-6">
        <h2 class="text-xl font-semibold text-gray-900">Review Your Order</h2>
        <div class="bg-gray-50 p-4 rounded-lg space-y-3">
          <div v-for="item in cartStore.items" :key="item.id" class="flex justify-between text-sm">
            <span>{{ item.name }} x{{ item.quantity }}</span>
            <span>₹{{ (item.price * item.quantity).toFixed(2) }}</span>
          </div>
          <hr class="border-gray-200" />
          <div class="flex justify-between font-semibold">
            <span>Total</span>
            <span>₹{{ cartStore.cartTotal.toFixed(2) }}</span>
          </div>
        </div>
        <div class="flex gap-4">
          <button @click="step = 2" class="btn-secondary flex-1">Back</button>
          <button @click="placeOrder" :disabled="loading" class="btn-primary flex-1">
            {{ loading ? 'Placing Order...' : 'Place Order' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
