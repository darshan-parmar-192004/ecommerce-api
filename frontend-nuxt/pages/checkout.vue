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

const currentStep = ref(1)
const loading = ref(false)
const error = ref('')
const success = ref(false)
const paymentError = ref('')

const shippingForm = reactive({
  fullName: '',
  address: '',
  city: '',
  state: '',
  zipCode: '',
  phone: ''
})

const paymentForm = reactive({
  cardNumber: '',
  expiry: '',
  cvv: '',
  nameOnCard: ''
})

const shippingErrors = reactive({})
const paymentErrors = reactive({})

const validateShipping = () => {
  shippingErrors.fullName = !shippingForm.fullName.trim() ? 'Full name is required' : ''
  shippingErrors.address = !shippingForm.address.trim() ? 'Address is required' : ''
  shippingErrors.city = !shippingForm.city.trim() ? 'City is required' : ''
  shippingErrors.state = !shippingForm.state.trim() ? 'State is required' : ''
  shippingErrors.zipCode = !shippingForm.zipCode.trim() ? 'ZIP code is required' : ''
  shippingErrors.phone = !shippingForm.phone.trim() ? 'Phone number is required' : ''
  
  return !Object.values(shippingErrors).some(e => e)
}

const validatePayment = () => {
  paymentErrors.nameOnCard = !paymentForm.nameOnCard.trim() ? 'Name on card is required' : ''
  paymentErrors.cardNumber = !paymentForm.cardNumber.trim() ? 'Card number is required' : ''
  paymentErrors.expiry = !paymentForm.expiry.trim() ? 'Expiry date is required' : ''
  paymentErrors.cvv = !paymentForm.cvv.trim() ? 'CVV is required' : ''
  
  return !Object.values(paymentErrors).some(e => e)
}

const goToShipping = () => {
  currentStep.value = 1
}

const goToPayment = () => {
  if (validateShipping()) {
    currentStep.value = 2
  }
}

const processPayment = async () => {
  if (!validatePayment()) return
  
  loading.value = true
  paymentError.value = ''
  
  await new Promise(resolve => setTimeout(resolve, 1500))
  
  const lastFour = paymentForm.cardNumber.slice(-4)
  if (lastFour === '0000') {
    paymentError.value = 'Payment declined. Please try a different card.'
    loading.value = false
    return
  }
  
  await handleSubmit()
}

const handleSubmit = async () => {
  if (cartStore.isEmpty) {
    error.value = 'Your cart is empty'
    return
  }

  loading.value = true
  error.value = ''

  try {
    const shippingAddress = `${shippingForm.fullName}, ${shippingForm.address}, ${shippingForm.city}, ${shippingForm.state} ${shippingForm.zipCode}, ${shippingForm.phone}`
    
    const orderData = {
      order: {
        customer_id: authStore.user?.customer_id,
        shipping_address: shippingAddress,
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

const formatCardNumber = (e) => {
  let value = e.target.value.replace(/\D/g, '')
  value = value.substring(0, 16)
  paymentForm.cardNumber = value.replace(/(\d{4})(?=\d)/g, '$1 ')
}

const formatExpiry = (e) => {
  let value = e.target.value.replace(/\D/g, '')
  value = value.substring(0, 4)
  if (value.length >= 2) {
    value = value.substring(0, 2) + '/' + value.substring(2)
  }
  paymentForm.expiry = value
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

    <div v-else>
      <div class="flex items-center justify-center mb-8">
        <div class="flex items-center">
          <div :class="['w-10 h-10 rounded-full flex items-center justify-center font-semibold transition-all', currentStep >= 1 ? 'bg-gradient-to-r from-primary to-primary-container text-white' : 'bg-surface-container text-on_surface_variant']">
            1
          </div>
          <div :class="['w-20 h-1 mx-2 transition-all', currentStep >= 2 ? 'bg-gradient-to-r from-primary to-primary-container' : 'bg-surface-container']" />
          <div :class="['w-10 h-10 rounded-full flex items-center justify-center font-semibold transition-all', currentStep >= 2 ? 'bg-gradient-to-r from-primary to-primary-container text-white' : 'bg-surface-container text-on_surface_variant']">
            2
          </div>
          <div :class="['w-20 h-1 mx-2 transition-all', currentStep >= 3 ? 'bg-gradient-to-r from-primary to-primary-container' : 'bg-surface-container']" />
          <div :class="['w-10 h-10 rounded-full flex items-center justify-center font-semibold transition-all', currentStep >= 3 ? 'bg-gradient-to-r from-primary to-primary-container text-white' : 'bg-surface-container text-on_surface_variant']">
            3
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <div>
          <div v-if="currentStep === 1" class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
            <h2 class="text-lg font-semibold text-on_surface font-display mb-4">Shipping Information</h2>
            
            <div v-if="error" class="bg-error-container text-error p-3 rounded-lg text-sm mb-4 flex items-start gap-2">
              <svg class="w-5 h-5 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              {{ error }}
            </div>

            <form @submit.prevent="goToPayment" class="space-y-4">
              <div>
                <label for="fullName" class="block text-sm font-medium text-on_surface_variant mb-2">
                  Full Name <span class="text-error">*</span>
                </label>
                <input
                  id="fullName"
                  v-model="shippingForm.fullName"
                  type="text"
                  required
                  class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                  placeholder="John Doe"
                />
                <p v-if="shippingErrors.fullName" class="text-error text-xs mt-1">{{ shippingErrors.fullName }}</p>
              </div>

              <div>
                <label for="address" class="block text-sm font-medium text-on_surface_variant mb-2">
                  Address <span class="text-error">*</span>
                </label>
                <textarea
                  id="address"
                  v-model="shippingForm.address"
                  rows="2"
                  required
                  class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                  placeholder="123 Main Street"
                />
                <p v-if="shippingErrors.address" class="text-error text-xs mt-1">{{ shippingErrors.address }}</p>
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label for="city" class="block text-sm font-medium text-on_surface_variant mb-2">
                    City <span class="text-error">*</span>
                  </label>
                  <input
                    id="city"
                    v-model="shippingForm.city"
                    type="text"
                    required
                    class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                    placeholder="New York"
                  />
                  <p v-if="shippingErrors.city" class="text-error text-xs mt-1">{{ shippingErrors.city }}</p>
                </div>

                <div>
                  <label for="state" class="block text-sm font-medium text-on_surface_variant mb-2">
                    State <span class="text-error">*</span>
                  </label>
                  <input
                    id="state"
                    v-model="shippingForm.state"
                    type="text"
                    required
                    class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                    placeholder="NY"
                  />
                  <p v-if="shippingErrors.state" class="text-error text-xs mt-1">{{ shippingErrors.state }}</p>
                </div>
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label for="zipCode" class="block text-sm font-medium text-on_surface_variant mb-2">
                    ZIP Code <span class="text-error">*</span>
                  </label>
                  <input
                    id="zipCode"
                    v-model="shippingForm.zipCode"
                    type="text"
                    required
                    class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                    placeholder="10001"
                  />
                  <p v-if="shippingErrors.zipCode" class="text-error text-xs mt-1">{{ shippingErrors.zipCode }}</p>
                </div>

                <div>
                  <label for="phone" class="block text-sm font-medium text-on_surface_variant mb-2">
                    Phone <span class="text-error">*</span>
                  </label>
                  <input
                    id="phone"
                    v-model="shippingForm.phone"
                    type="tel"
                    required
                    class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                    placeholder="+1 234 567 8900"
                  />
                  <p v-if="shippingErrors.phone" class="text-error text-xs mt-1">{{ shippingErrors.phone }}</p>
                </div>
              </div>

              <div class="pt-4">
                <button 
                  type="submit"
                  class="w-full py-4 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90 flex items-center justify-center gap-2 font-semibold shadow-lg shadow-primary/25"
                >
                  Continue to Payment
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
                  </svg>
                </button>
              </div>
            </form>
          </div>

          <div v-if="currentStep === 2" class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
            <h2 class="text-lg font-semibold text-on_surface font-display mb-4">Payment Information</h2>
            
            <div v-if="paymentError" class="bg-error-container text-error p-3 rounded-lg text-sm mb-4 flex items-start gap-2">
              <svg class="w-5 h-5 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              {{ paymentError }}
            </div>

            <form @submit.prevent="processPayment" class="space-y-4">
              <div>
                <label for="nameOnCard" class="block text-sm font-medium text-on_surface_variant mb-2">
                  Name on Card <span class="text-error">*</span>
                </label>
                <input
                  id="nameOnCard"
                  v-model="paymentForm.nameOnCard"
                  type="text"
                  required
                  class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                  placeholder="John Doe"
                />
                <p v-if="paymentErrors.nameOnCard" class="text-error text-xs mt-1">{{ paymentErrors.nameOnCard }}</p>
              </div>

              <div>
                <label for="cardNumber" class="block text-sm font-medium text-on_surface_variant mb-2">
                  Card Number <span class="text-error">*</span>
                </label>
                <input
                  id="cardNumber"
                  :value="paymentForm.cardNumber"
                  @input="formatCardNumber"
                  type="text"
                  required
                  class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                  placeholder="1234 5678 9012 3456"
                />
                <p v-if="paymentErrors.cardNumber" class="text-error text-xs mt-1">{{ paymentErrors.cardNumber }}</p>
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label for="expiry" class="block text-sm font-medium text-on_surface_variant mb-2">
                    Expiry Date <span class="text-error">*</span>
                  </label>
                  <input
                    id="expiry"
                    :value="paymentForm.expiry"
                    @input="formatExpiry"
                    type="text"
                    required
                    class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                    placeholder="MM/YY"
                  />
                  <p v-if="paymentErrors.expiry" class="text-error text-xs mt-1">{{ paymentErrors.expiry }}</p>
                </div>

                <div>
                  <label for="cvv" class="block text-sm font-medium text-on_surface_variant mb-2">
                    CVV <span class="text-error">*</span>
                  </label>
                  <input
                    id="cvv"
                    v-model="paymentForm.cvv"
                    type="text"
                    maxlength="4"
                    required
                    class="w-full px-4 py-3 bg-surface-container border border-transparent rounded-lg focus:outline-none focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary text-on_surface placeholder:text-outline"
                    placeholder="123"
                  />
                  <p v-if="paymentErrors.cvv" class="text-error text-xs mt-1">{{ paymentErrors.cvv }}</p>
                </div>
              </div>

              <div class="pt-4 flex gap-3">
                <button 
                  type="button"
                  @click="goToShipping"
                  class="flex-1 py-4 bg-surface-container border border-outline-variant text-on_surface rounded-lg hover:bg-surface-container-high flex items-center justify-center gap-2 font-semibold"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16l-4-4m0 0l4-4m-4 4h18" />
                  </svg>
                  Back
                </button>
                <button 
                  type="submit"
                  :disabled="loading"
                  class="flex-1 py-4 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90 disabled:opacity-50 flex items-center justify-center gap-2 font-semibold shadow-lg shadow-primary/25"
                >
                  <svg v-if="loading" class="animate-spin h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  {{ loading ? 'Processing...' : 'Place Order' }}
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
  </div>
</template>
