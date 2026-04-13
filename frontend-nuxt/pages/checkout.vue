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
const processingPayment = ref(false)
const error = ref('')
const success = ref(false)
const paymentError = ref('')

const selectedPaymentMethod = ref('razorpay')
const paymentMethods = [
  { id: 'razorpay', name: 'Pay with Razorpay', icon: 'credit-card', description: 'Instant payment' },
  { id: 'upi', name: 'UPI', icon: 'smartphone', description: 'Pay via UPI app' },
  { id: 'wallet', name: 'Wallet', icon: 'wallet', description: 'Paytm, PhonePe, etc.' },
  { id: 'cod', name: 'Cash on Delivery', icon: 'cash', description: 'Pay when you receive' }
]

const promoCode = ref('')
const promoApplied = ref(false)
const promoDiscount = ref(0)
const promoError = ref('')

const applyPromoCode = () => {
  promoError.value = ''
  const code = promoCode.value.trim().toUpperCase()
  
  const promoCodes = {
    'SAVE10': 10,
    'FLAT20': 20,
    'FIRST50': 50
  }
  
  if (promoCodes[code]) {
    promoDiscount.value = promoCodes[code]
    promoApplied.value = true
  } else {
    promoError.value = 'Invalid promo code'
  }
}

const removePromo = () => {
  promoCode.value = ''
  promoDiscount.value = 0
  promoApplied.value = false
}

const discountAmount = computed(() => {
  return (cartStore.subtotal * promoDiscount.value) / 100
})

const discountedTotal = computed(() => {
  return cartStore.total - discountAmount.value
})

const shippingForm = reactive({
  fullName: '',
  address: '',
  city: '',
  state: '',
  zipCode: '',
  phone: ''
})

const shippingErrors = reactive({})

const validateShipping = () => {
  shippingErrors.fullName = !shippingForm.fullName.trim() ? 'Full name is required' : ''
  shippingErrors.address = !shippingForm.address.trim() ? 'Address is required' : ''
  shippingErrors.city = !shippingForm.city.trim() ? 'City is required' : ''
  shippingErrors.state = !shippingForm.state.trim() ? 'State is required' : ''
  shippingErrors.zipCode = !shippingForm.zipCode.trim() ? 'ZIP code is required' : ''
  shippingErrors.phone = !shippingForm.phone.trim() ? 'Phone number is required' : ''
  
  return !Object.values(shippingErrors).some(e => e)
}

const goToShipping = () => {
  currentStep.value = 1
}

const goToPayment = () => {
  if (validateShipping()) {
    currentStep.value = 2
  }
}

const mockRazorpayPayment = async () => {
  processingPayment.value = true
  paymentError.value = ''
  error.value = ''
  
  try {
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    const mockSuccess = true
    
    if (!mockSuccess) {
      paymentError.value = 'Payment failed. Please try again.'
      processingPayment.value = false
      return
    }
    
    await handleSubmit()
  } catch (err) {
    console.error('Payment error:', err)
    paymentError.value = err.message || 'Payment failed. Please try again.'
  } finally {
    processingPayment.value = false
  }
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

    console.log('Creating order with data:', JSON.stringify(orderData, null, 2))
    
    const result = await ordersApi.create(orderData)
    
    console.log('Order created successfully:', result)

    cartStore.clearCart()
    success.value = true
    
    setTimeout(() => {
      router.push(`/orders/${result.order_id || result.data?.order_id}`)
    }, 1500)
  } catch (err) {
    console.error('Order creation error:', err)
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
            
            <div v-if="error" class="bg-error-container text-error p-3 rounded-lg text-sm mb-4 flex items-start gap-2">
              <svg class="w-5 h-5 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              {{ error }}
            </div>

            <div v-if="paymentError" class="bg-error-container text-error p-3 rounded-lg text-sm mb-4 flex items-start gap-2">
              <svg class="w-5 h-5 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              {{ paymentError }}
            </div>

            <!-- Promo Code Section -->
            <div class="mb-6 p-4 bg-surface-container rounded-lg">
              <label class="block text-sm font-medium text-on_surface_variant mb-2">Promo Code</label>
              <div v-if="promoApplied" class="flex items-center justify-between">
                <div class="flex items-center gap-2">
                  <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <span class="text-green-600 font-medium">{{ promoCode.value }} applied (-{{ promoDiscount.value }}%)</span>
                </div>
                <button @click="removePromo" class="text-sm text-error hover:text-error/80">Remove</button>
              </div>
              <div v-else class="flex gap-2">
                <input 
                  v-model="promoCode"
                  type="text"
                  placeholder="Enter promo code"
                  class="flex-1 px-4 py-2 bg-surface-container-low border border-transparent rounded-lg focus:outline-none focus:ring-2 focus:ring-primary/20 text-on_surface placeholder:text-outline"
                  @keyup.enter="applyPromoCode"
                />
                <button 
                  @click="applyPromoCode"
                  class="px-4 py-2 bg-primary text-white rounded-lg hover:bg-primary/90 transition-colors"
                >
                  Apply
                </button>
              </div>
              <p v-if="promoError" class="text-error text-xs mt-2">{{ promoError }}</p>
              <p class="text-xs text-outline mt-2">Try: SAVE10, FLAT20, FIRST50</p>
            </div>

            <!-- Payment Method Selection -->
            <div class="mb-6">
              <label class="block text-sm font-medium text-on_surface_variant mb-3">Payment Method</label>
              <div class="grid grid-cols-2 gap-3">
                <button
                  v-for="method in paymentMethods"
                  :key="method.id"
                  @click="selectedPaymentMethod = method.id"
                  :class="[
                    'p-4 rounded-lg border-2 text-left transition-all duration-200',
                    selectedPaymentMethod === method.id 
                      ? 'border-primary bg-primary/10 shadow-lg shadow-primary/20' 
                      : 'border-outline-variant/30 hover:border-primary/50 bg-surface-container'
                  ]"
                >
                  <div class="flex items-center gap-3">
                    <div :class="[
                      'w-10 h-10 rounded-full flex items-center justify-center',
                      selectedPaymentMethod === method.id ? 'bg-primary text-white' : 'bg-surface-container-high text-on_surface_variant'
                    ]">
                      <svg v-if="method.icon === 'credit-card'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h4m1 4l-4-4m0 0l4-4m-4 4h10" />
                      </svg>
                      <svg v-else-if="method.icon === 'smartphone'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z" />
                      </svg>
                      <svg v-else-if="method.icon === 'wallet'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                      </svg>
                      <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z" />
                      </svg>
                    </div>
                    <div>
                      <p class="font-semibold text-on_surface text-sm">{{ method.name }}</p>
                      <p class="text-xs text-on_surface_variant">{{ method.description }}</p>
                    </div>
                  </div>
                </button>
              </div>
            </div>

            <div class="text-center py-4">
              <div class="bg-surface-container rounded-lg p-4 mb-6 text-left">
                <div class="flex justify-between items-center py-2 border-b border-outline-variant/20">
                  <span class="text-on_surface_variant">Order Total</span>
                  <span class="text-xl font-bold text-primary">₹ {{ discountedTotal.toFixed(2) }}</span>
                </div>
                <div v-if="promoApplied" class="flex justify-between items-center py-2 border-b border-outline-variant/20 text-green-600">
                  <span>Discount ({{ promoDiscount.value }}%)</span>
                  <span>-₹ {{ discountAmount.toFixed(2) }}</span>
                </div>
                <div class="flex justify-between items-center py-2 text-sm">
                  <span class="text-on_surface_variant">Payment Method</span>
                  <span class="text-on_surface capitalize">{{ paymentMethods.find(m => m.id === selectedPaymentMethod)?.name }}</span>
                </div>
              </div>

              <button 
                @click="mockRazorpayPayment"
                :disabled="processingPayment"
                class="w-full py-4 bg-gradient-to-r from-primary to-primary-container text-white rounded-lg hover:opacity-90 disabled:opacity-50 flex items-center justify-center gap-2 font-semibold shadow-lg shadow-primary/25 hover:scale-[1.02] active:scale-[0.98] transition-all duration-200"
              >
                <svg v-if="processingPayment" class="animate-spin h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M3 10h18M7 15h4m1 4l-4-4m0 0l4-4m-4 4h10" />
                </svg>
                {{ processingPayment ? 'Processing Payment...' : `Pay ₹${discountedTotal.toFixed(2)}` }}
              </button>

              <div class="flex gap-3 mt-4">
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
              </div>

              <p class="text-xs text-on_surface_variant mt-6">
                <svg class="w-4 h-4 inline mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
                This is a demo payment. No real money will be charged.
              </p>
            </div>
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
              <div v-if="promoApplied" class="flex justify-between text-sm text-green-600">
                <span>Discount ({{ promoDiscount.value }}%)</span>
                <span>-₹ {{ discountAmount.toFixed(2) }}</span>
              </div>
              <div class="flex justify-between text-sm text-on_surface_variant">
                <span>Tax (8%)</span>
                <span class="font-medium text-on_surface">₹ {{ cartStore.tax.toFixed(2) }}</span>
              </div>
              <div class="flex justify-between text-base pt-2 border-t border-outline-variant/20">
                <span class="font-semibold text-on_surface">Total</span>
                <span class="font-bold text-xl text-primary">₹ {{ discountedTotal.toFixed(2) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
