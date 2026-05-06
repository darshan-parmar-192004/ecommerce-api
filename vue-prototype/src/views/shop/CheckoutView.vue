<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'
import { useUserStore } from '@/stores/user'
import { useToastStore } from '@/stores/toast'
import { useFormValidation } from '@/composables/useFormValidation'

const router = useRouter()
const cartStore = useCartStore()
const userStore = useUserStore()
const toastStore = useToastStore()

const step = ref(1)
const loading = ref(false)
const orderPlaced = ref(false)
const orderId = ref('N/A')

const shippingSchema = {
  address: {
    required: true,
    requiredMessage: 'Address is required',
    minLength: 5,
    minLengthMessage: 'Address must be at least 5 characters'
  },
  city: {
    required: true,
    requiredMessage: 'City is required'
  },
  state: {
    required: true,
    requiredMessage: 'State is required'
  },
  zip: {
    required: true,
    requiredMessage: 'ZIP code is required',
    pattern: /^\d{5}(-\d{4})?$/,
    patternMessage: 'Invalid ZIP code format'
  },
  country: {
    required: true,
    requiredMessage: 'Country is required'
  }
}

const {
  form: shippingInfo,
  errors: shippingErrors,
  validate: validateShipping,
  handleBlur: handleShippingBlur
} = useFormValidation(shippingSchema, {
  address: '',
  city: '',
  state: '',
  zip: '',
  country: ''
})

onMounted(async () => {
  await userStore.fetchProfile()
  if (userStore.profile) {
    shippingInfo.address = userStore.profile.address || ''
    shippingInfo.city = userStore.profile.city || ''
    shippingInfo.state = userStore.profile.state || ''
    shippingInfo.zip = userStore.profile.zip || ''
    shippingInfo.country = userStore.profile.country || ''
  }
})

const validateShippingStep = () => {
  return validateShipping()
}

const goToPayment = () => {
  if (validateShippingStep()) {
    step.value = 2
  }
}

const placeOrder = async () => {
  if (!validateShippingStep()) return

  loading.value = true
  try {
    const orderItems = cartStore.items.map(item => ({
      product_id: item.id,
      quantity: item.quantity,
      unit_price: item.price
    }))

    if (!userStore.profile?.customer_id) {
      toastStore.error('User profile not loaded. Please try again.')
      return
    }

    const orderData = {
      order: {
        customer_id: userStore.profile.customer_id,
        total_amount: cartStore.cartTotal,
        status: 'pending',
        shipping_address: `${shippingInfo.address}, ${shippingInfo.city}, ${shippingInfo.state} ${shippingInfo.zip}, ${shippingInfo.country}`
      },
      items: orderItems
    }

    const result = await userStore.createOrder(orderData)
    orderId.value = result.order_id || 'N/A'
    orderPlaced.value = true
    cartStore.clearCart()
    toastStore.success('Order placed successfully!')
  } catch (err) {
    toastStore.error('Failed to place order. Please try again.')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-8">Checkout</h1>

    <div v-if="orderPlaced" class="text-center py-16">
      <div class="w-16 h-16 bg-green-100 dark:bg-green-900/30 rounded-full flex items-center justify-center mx-auto mb-4">
        <svg class="w-8 h-8 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
        </svg>
      </div>
      <h2 class="text-2xl font-bold text-gray-900 dark:text-gray-100 mb-2">Order Placed!</h2>
      <p class="text-gray-600 dark:text-gray-400 mb-2">Thank you for your purchase.</p>
      <p class="text-sm text-gray-500 dark:text-gray-400 mb-6">Order ID: {{ orderId }}</p>
      <RouterLink to="/user/orders" class="btn-primary">
        View Order
      </RouterLink>
    </div>

    <div v-else class="space-y-8">
      <!-- Progress Steps -->
      <div class="flex items-center gap-4 mb-8">
        <div :class="step >= 1 ? 'bg-brand-900 text-white' : 'bg-gray-200 dark:bg-brand-700'" class="w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium transition-colors duration-200">
          1
        </div>
        <div class="flex-1 h-1 bg-gray-200 dark:bg-brand-700">
          <div :class="step >= 2 ? 'bg-brand-900' : 'bg-gray-200 dark:bg-brand-700'" class="h-full transition-all duration-300" :style="{ width: step >= 2 ? '100%' : '0%' }" />
        </div>
        <div :class="step >= 2 ? 'bg-brand-900 text-white' : 'bg-gray-200 dark:bg-brand-700'" class="w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium transition-colors duration-200">
          2
        </div>
        <div class="flex-1 h-1 bg-gray-200 dark:bg-brand-700">
          <div :class="step >= 3 ? 'bg-brand-900' : 'bg-gray-200 dark:bg-brand-700'" class="h-full transition-all duration-300" :style="{ width: step >= 3 ? '100%' : '0%' }" />
        </div>
        <div :class="step >= 3 ? 'bg-brand-900 text-white' : 'bg-gray-200 dark:bg-brand-700'" class="w-8 h-8 rounded-full flex items-center justify-center text-sm font-medium transition-colors duration-200">
          3
        </div>
      </div>

      <!-- Step 1: Shipping -->
      <div v-if="step === 1" class="space-y-4">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-gray-100">Shipping Information</h2>
        <div>
          <input
            v-model="shippingInfo.address"
            @blur="handleShippingBlur('address')"
            placeholder="Address"
            class="input"
            :class="{ 'border-red-300 focus:ring-red-500': shippingErrors.address }"
          />
          <p v-if="shippingErrors.address" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ shippingErrors.address }}</p>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div>
            <input
              v-model="shippingInfo.city"
              @blur="handleShippingBlur('city')"
              placeholder="City"
              class="input"
              :class="{ 'border-red-300 focus:ring-red-500': shippingErrors.city }"
            />
            <p v-if="shippingErrors.city" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ shippingErrors.city }}</p>
          </div>
          <div>
            <input
              v-model="shippingInfo.state"
              @blur="handleShippingBlur('state')"
              placeholder="State"
              class="input"
              :class="{ 'border-red-300 focus:ring-red-500': shippingErrors.state }"
            />
            <p v-if="shippingErrors.state" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ shippingErrors.state }}</p>
          </div>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div>
            <input
              v-model="shippingInfo.zip"
              @blur="handleShippingBlur('zip')"
              placeholder="ZIP Code"
              class="input"
              :class="{ 'border-red-300 focus:ring-red-500': shippingErrors.zip }"
            />
            <p v-if="shippingErrors.zip" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ shippingErrors.zip }}</p>
          </div>
          <div>
            <input
              v-model="shippingInfo.country"
              @blur="handleShippingBlur('country')"
              placeholder="Country"
              class="input"
              :class="{ 'border-red-300 focus:ring-red-500': shippingErrors.country }"
            />
            <p v-if="shippingErrors.country" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ shippingErrors.country }}</p>
          </div>
        </div>
        <button @click="goToPayment" class="btn-primary w-full mt-4">
          Continue to Review
        </button>
      </div>

      <!-- Step 2: Review Order -->
      <div v-else-if="step === 2" class="space-y-6">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-gray-100">Review Your Order</h2>
        <div class="bg-gray-50 dark:bg-brand-800 p-4 rounded-lg space-y-3">
          <div v-for="item in cartStore.items" :key="item.id" class="flex justify-between text-sm text-gray-700 dark:text-gray-300">
            <span>{{ item.name }} x{{ item.quantity }}</span>
            <span>₹{{ (item.price * item.quantity).toFixed(2) }}</span>
          </div>
          <hr class="border-gray-200 dark:border-brand-700" />
          <div class="flex justify-between font-semibold text-gray-900 dark:text-gray-100">
            <span>Total</span>
            <span>₹{{ cartStore.cartTotal.toFixed(2) }}</span>
          </div>
        </div>

        <div class="bg-gray-50 dark:bg-brand-800 p-4 rounded-lg">
          <h3 class="text-sm font-semibold text-gray-900 dark:text-gray-100 mb-2">Shipping Address</h3>
          <p class="text-sm text-gray-600 dark:text-gray-400">
            {{ shippingInfo.address }}, {{ shippingInfo.city }}, {{ shippingInfo.state }} {{ shippingInfo.zip }}, {{ shippingInfo.country }}
          </p>
        </div>

        <div class="flex gap-4">
          <button @click="step = 1" class="btn-secondary flex-1">Back</button>
          <button @click="placeOrder" :disabled="loading || cartStore.items.length === 0" class="btn-primary flex-1">
            <span v-if="loading">Placing Order...</span>
            <span v-else>Place Order</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
