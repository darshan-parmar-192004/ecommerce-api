<script setup>
import { useCartStore } from '~/stores/cart'

const props = defineProps({
  product: {
    type: Object,
    required: true
  }
})

const cartStore = useCartStore()
const isAdding = ref(false)
const cardRef = ref(null)
const tiltStyle = ref({})

const addToCart = async () => {
  isAdding.value = true
  cartStore.addItem(props.product)
  cartStore.openCart()
  setTimeout(() => {
    isAdding.value = false
  }, 500)
}

const handleMouseMove = (e) => {
  if (!cardRef.value) return
  
  const rect = cardRef.value.getBoundingClientRect()
  const x = e.clientX - rect.left
  const y = e.clientY - rect.top
  
  const centerX = rect.width / 2
  const centerY = rect.height / 2
  
  const rotateX = (y - centerY) / 20
  const rotateY = (centerX - x) / 20
  
  tiltStyle.value = {
    transform: `perspective(1000px) rotateX(${rotateX}deg) rotateY(${rotateY}deg) scale3d(1.02, 1.02, 1.02)`,
    transition: 'transform 0.1s ease-out'
  }
}

const resetTilt = () => {
  tiltStyle.value = {
    transform: 'perspective(1000px) rotateX(0) rotateY(0) scale3d(1, 1, 1)',
    transition: 'transform 0.3s ease-out'
  }
}
</script>

<template>
  <div 
    ref="cardRef"
    @mousemove="handleMouseMove"
    @mouseleave="resetTilt"
    class="group overflow-hidden rounded-lg bg-surface-container-lowest shadow-ambient transition-all duration-300 hover:shadow-xl"
    :style="tiltStyle"
  >
    <NuxtLink
      :to="`/products/${product.product_id}`"
      class="block aspect-[4/3] bg-surface-container relative overflow-hidden"
    >
      <div class="absolute inset-0 flex items-center justify-center transition-transform duration-700 ease-out group-hover:scale-110">
        <svg class="w-20 h-20 text-outline/30" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1">
          <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
        </svg>
      </div>
      
      <!-- Quick Add Button Overlay -->
      <Transition name="fade">
        <div class="absolute inset-0 bg-on_surface/40 opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex items-center justify-center">
          <button
            @click.prevent="addToCart"
            :disabled="isAdding || product.stock === 0"
            class="transform scale-90 group-hover:scale-100 transition-transform duration-300 bg-white text-on_surface px-6 py-3 rounded-lg font-semibold flex items-center gap-2 shadow-lg hover:bg-primary hover:text-white"
          >
            <svg v-if="isAdding" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
            </svg>
            {{ isAdding ? 'Adding...' : 'Quick Add' }}
          </button>
        </div>
      </Transition>

      <Transition name="fade">
        <div v-if="product.stock === 0" class="absolute top-3 right-3 bg-error text-on_error text-[10px] font-bold uppercase tracking-wider px-3 py-1.5 rounded">
          Out of Stock
        </div>
      </Transition>

      <!-- Category Tag -->
      <div v-if="product.category_name" class="absolute top-3 left-3">
        <span class="text-[10px] font-bold uppercase tracking-wider px-2 py-1 rounded bg-surface/80 backdrop-blur-sm text-on_surface">
          {{ product.category_name }}
        </span>
      </div>
    </NuxtLink>

    <div class="p-5">
      <NuxtLink :to="`/products/${product.product_id}`" class="block">
        <h3 class="font-display font-semibold text-on_surface transition-colors duration-300 group-hover:text-primary line-clamp-2 min-h-[2.5rem] text-base">
          {{ product.name }}
        </h3>
        <p v-if="product.description" class="mt-1.5 text-sm text-outline line-clamp-2 font-body">
          {{ product.description }}
        </p>
      </NuxtLink>

      <div class="mt-5 flex items-center justify-between">
        <div>
          <span class="text-xl font-bold text-primary font-display">
            ₹{{ Number(product.price).toFixed(2) }}
          </span>
          <p v-if="product.original_price && product.original_price > product.price" class="text-xs text-outline line-through mt-0.5">
            ₹{{ Number(product.original_price).toFixed(2) }}
          </p>
        </div>

        <button
          @click.prevent="addToCart"
          :disabled="isAdding || product.stock === 0"
          class="flex items-center gap-2 bg-gradient-to-r from-primary to-primary-container text-white px-4 py-2.5 rounded-md disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-300 hover:shadow-lg hover:scale-[1.03] active:scale-[0.97] shadow-lg shadow-primary/25"
        >
          <svg v-if="isAdding" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
          </svg>
          <span class="font-medium text-xs uppercase tracking-wider">Add</span>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
