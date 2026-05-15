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
const { success: showSuccess, error: showError } = useAppToast()

useSeoMeta({
  title: 'Checkout - E-Commerce Store'
})

const currentStep = ref(1)
const loading = ref(false)
const processingPayment = ref(false)
const error = ref('')
const success = ref(false)

const selectedPaymentMethod = ref('razorpay')
const paymentMethods = [
  { id: 'razorpay', name: 'Pay with Razorpay', icon: 'credit-card', description: 'Instant payment' },
  { id: 'upi', name: 'UPI', icon: 'mobile', description: 'Pay via UPI app' },
  { id: 'wallet', name: 'Wallet', icon: 'wallet', description: 'Paytm, PhonePe, etc.' },
  { id: 'cod', name: 'Cash on Delivery', icon: 'money', description: 'Pay when you receive' }
]

const promoCode = ref('')
const promoApplied = ref(false)
const promoDiscount = ref(0)
const promoError = ref('')

const applyPromoCode = () => {
  promoError.value = ''
  const code = promoCode.value.trim().toUpperCase()
  const promoCodes = { 'SAVE10': 10, 'FLAT20': 20, 'FIRST50': 50 }
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

const discountAmount = computed(() => (cartStore.subtotal * promoDiscount.value) / 100)
const discountedTotal = computed(() => cartStore.total - discountAmount.value)

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

const goToShipping = () => { currentStep.value = 1 }
const goToPayment = () => { if (validateShipping()) currentStep.value = 2 }

const mockRazorpayPayment = async () => {
  processingPayment.value = true
  try {
    await handleSubmit()
  } catch (err) {
    showError('Payment failed. Please try again.')
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
    const shippingAddress = shippingForm.fullName + ', ' + shippingForm.address + ', ' + shippingForm.city + ', ' + shippingForm.state + ' ' + shippingForm.zipCode + ', ' + shippingForm.phone
    const orderData = {
      order: { customer_id: authStore.user?.customer_id, shipping_address: shippingAddress, total_amount: discountedTotal.value },
      items: cartStore.items.map(item => ({ product_id: item.product.product_id, quantity: item.quantity, unit_price: item.product.price }))
    }
    const result = await ordersApi.create(orderData)
    cartStore.clearCart()
    success.value = true
    showSuccess('Order placed successfully!')
    const orderId = result.order_id || result.data?.order_id
    setTimeout(() => { router.push('/orders/' + orderId) }, 1500)
  } catch (err) {
    error.value = err.message || 'Failed to create order. Please try again.'
    showError(error.value)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-surface">
    <h1 class="text-3xl font-bold text-on_surface font-display mb-8">Checkout</h1>

    <div v-if="cartStore.isEmpty" class="text-center py-16 bg-surface-container-lowest rounded-xl shadow-ambient">
      <i class="pi pi-shopping-cart text-6xl text-outline mb-4" />
      <h2 class="mt-4 text-xl font-medium text-on_surface">Your cart is empty</h2>
      <p class="mt-2 text-on_surface_variant">Add some products before checking out.</p>
      <NuxtLink to="/products">
        <Button label="Browse Products" class="mt-6" icon="pi pi-shopping-bag" />
      </NuxtLink>
    </div>

    <div v-else-if="success" class="text-center py-16 bg-surface-container-lowest rounded-xl shadow-ambient">
      <i class="pi pi-check-circle text-6xl text-green-600 mb-4" />
      <h2 class="mt-4 text-2xl font-medium text-on_surface">Order Placed Successfully!</h2>
      <p class="mt-2 text-on_surface_variant">Redirecting to your order...</p>
    </div>

    <div v-else>
      <StepList :model="[{label: 'Shipping'}, {label: 'Payment'}, {label: 'Confirm'}]" :activeStep="currentStep - 1" class="mb-8" />

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <div>
          <div v-if="currentStep === 1" class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
            <h2 class="text-lg font-semibold text-on_surface font-display mb-4">Shipping Information</h2>
            <Message v-if="error" severity="error" :closable="false" class="mb-4" />
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-on_surface_variant mb-2">Full Name <span class="text-error">*</span></label>
                <InputText v-model="shippingForm.fullName" class="w-full" placeholder="John Doe" />
                <small v-if="shippingErrors.fullName" class="text-error">{{ shippingErrors.fullName }}</small>
              </div>
              <div>
                <label class="block text-sm font-medium text-on_surface_variant mb-2">Address <span class="text-error">*</span></label>
                <Textarea v-model="shippingForm.address" rows="2" class="w-full" placeholder="123 Main Street" />
                <small v-if="shippingErrors.address" class="text-error">{{ shippingErrors.address }}</small>
              </div>
              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-on_surface_variant mb-2">City <span class="text-error">*</span></label>
                  <InputText v-model="shippingForm.city" class="w-full" placeholder="New York" />
                  <small v-if="shippingErrors.city" class="text-error">{{ shippingErrors.city }}</small>
                </div>
                <div>
                  <label class="block text-sm font-medium text-on_surface_variant mb-2">State <span class="text-error">*</span></label>
                  <InputText v-model="shippingForm.state" class="w-full" placeholder="NY" />
                  <small v-if="shippingErrors.state" class="text-error">{{ shippingErrors.state }}</small>
                </div>
              </div>
              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-on_surface_variant mb-2">ZIP Code <span class="text-error">*</span></label>
                  <InputText v-model="shippingForm.zipCode" class="w-full" placeholder="10001" />
                  <small v-if="shippingErrors.zipCode" class="text-error">{{ shippingErrors.zipCode }}</small>
                </div>
                <div>
                  <label class="block text-sm font-medium text-on_surface_variant mb-2">Phone <span class="text-error">*</span></label>
                  <InputText v-model="shippingForm.phone" class="w-full" placeholder="+1 234 567 8900" />
                  <small v-if="shippingErrors.phone" class="text-error">{{ shippingErrors.phone }}</small>
                </div>
              </div>
              <Button @click="goToPayment" label="Continue to Payment" icon="pi pi-arrow-right" iconPos="right" class="w-full !mt-4 !py-3" />
            </div>
          </div>

          <div v-if="currentStep === 2" class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
            <h2 class="text-lg font-semibold text-on_surface font-display mb-4">Payment Information</h2>
            <Message v-if="error" severity="error" :closable="false" class="mb-4" />
            <div class="mb-6 p-4 bg-surface-container rounded-lg">
              <label class="block text-sm font-medium text-on_surface_variant mb-2">Promo Code</label>
              <div v-if="promoApplied" class="flex items-center justify-between">
                <div class="flex items-center gap-2">
                  <i class="pi pi-check-circle text-green-600" />
                  <span class="text-green-600 font-medium">{{ promoCode }} applied (-{{ promoDiscount }}%)</span>
                </div>
                <Button @click="removePromo" label="Remove" text severity="danger" size="small" />
              </div>
              <div v-else class="flex gap-2">
                <InputText v-model="promoCode" placeholder="Enter promo code" class="flex-1" @keyup.enter="applyPromoCode" />
                <Button @click="applyPromoCode" label="Apply" />
              </div>
              <small v-if="promoError" class="text-error">{{ promoError }}</small>
              <p class="text-xs text-outline mt-2">Try: SAVE10, FLAT20, FIRST50</p>
            </div>
            <div class="mb-6">
              <label class="block text-sm font-medium text-on_surface_variant mb-3">Payment Method</label>
              <div class="space-y-3">
                <div v-for="method in paymentMethods" :key="method.id" class="flex items-center gap-3 p-3 rounded-lg border-2 cursor-pointer transition-all"
                  :class="selectedPaymentMethod === method.id ? 'border-primary bg-primary/10' : 'border-outline-variant/30 hover:border-primary/50 bg-surface-container'"
                  @click="selectedPaymentMethod = method.id"
                >
                  <i :class="'pi pi-' + method.icon" class="text-xl" />
                  <div>
                    <p class="font-semibold text-on_surface text-sm">{{ method.name }}</p>
                    <p class="text-xs text-on_surface_variant">{{ method.description }}</p>
                  </div>
                </div>
              </div>
            </div>
            <div class="bg-surface-container rounded-lg p-4 mb-6 text-left">
              <div class="flex justify-between items-center py-2 border-b border-outline-variant/20">
                <span class="text-on_surface_variant">Order Total</span>
                <span class="text-xl font-bold text-primary">Rs {{ discountedTotal.toFixed(2) }}</span>
              </div>
              <div v-if="promoApplied" class="flex justify-between items-center py-2 border-b border-outline-variant/20 text-green-600">
                <span>Discount ({{ promoDiscount }}%)</span>
                <span>-Rs {{ discountAmount.toFixed(2) }}</span>
              </div>
            </div>
            <Button @click="mockRazorpayPayment" :loading="processingPayment" :label="processingPayment ? 'Processing Payment...' : 'Pay Rs ' + discountedTotal.toFixed(2)" icon="pi pi-credit-card" class="w-full !py-4" />
            <Button @click="goToShipping" label="Back" icon="pi pi-arrow-left" text class="w-full mt-3" />
            <p class="text-xs text-on_surface_variant mt-4 text-center"><i class="pi pi-lock mr-1" />This is a demo payment. No real money will be charged.</p>
          </div>
        </div>

        <div>
          <div class="bg-surface-container-lowest rounded-xl shadow-ambient p-6">
            <h2 class="text-lg font-semibold text-on_surface font-display mb-4">Order Summary</h2>
            <div class="space-y-4 max-h-80 overflow-y-auto">
              <div v-for="item in cartStore.items" :key="item.product.product_id" class="flex items-center gap-4 py-3 border-b border-outline-variant/20 last:border-0">
                <div class="w-14 h-14 bg-surface-container rounded-lg flex-shrink-0 flex items-center justify-center">
                  <i class="pi pi-box text-2xl text-outline" />
                </div>
                <div class="flex-1 min-w-0">
                  <p class="font-medium text-on_surface truncate">{{ item.product.name }}</p>
                  <p class="text-sm text-on_surface_variant">Qty: {{ item.quantity }} x Rs {{ Number(item.product.price).toFixed(2) }}</p>
                </div>
                <p class="font-semibold text-on_surface">Rs {{ (Number(item.product.price) * item.quantity).toFixed(2) }}</p>
              </div>
            </div>
            <div class="mt-6 pt-4 border-t border-outline-variant/20 space-y-2">
              <div class="flex justify-between text-sm text-on_surface_variant">
                <span>Subtotal</span>
                <span class="font-medium text-on_surface">Rs {{ cartStore.subtotal.toFixed(2) }}</span>
              </div>
              <div v-if="promoApplied" class="flex justify-between text-sm text-green-600">
                <span>Discount ({{ promoDiscount }}%)</span>
                <span>-Rs {{ discountAmount.toFixed(2) }}</span>
              </div>
              <div class="flex justify-between text-sm text-on_surface_variant">
                <span>Tax (8%)</span>
                <span class="font-medium text-on_surface">Rs {{ cartStore.tax.toFixed(2) }}</span>
              </div>
              <div class="flex justify-between text-base pt-2 border-t border-outline-variant/20">
                <span class="font-semibold text-on_surface">Total</span>
                <span class="font-bold text-xl text-primary">Rs {{ discountedTotal.toFixed(2) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
