<script setup>
const props = defineProps({
  product: {
    type: Object,
    required: true
  }
})

const cartStore = useCartStore()
const isAdding = ref(false)

const addToCart = async () => {
  isAdding.value = true
  cartStore.addItem(props.product)
  cartStore.openCart()
  setTimeout(() => {
    isAdding.value = false
  }, 500)
}
</script>

<template>
  <div class="bg-white rounded-xl shadow-sm hover:shadow-xl transition-all duration-300 overflow-hidden group -translate-y-0 hover:-translate-y-1">
    <NuxtLink
      :to="`/products/${product.product_id}`"
      class="block aspect-square bg-gradient-to-br from-gray-50 to-gray-100 relative overflow-hidden"
    >
      <div class="absolute inset-0 flex items-center justify-center group-hover:scale-110 transition-transform duration-500">
        <svg class="w-20 h-20 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
        </svg>
      </div>
      <Transition name="fade">
        <div v-if="product.stock === 0" class="absolute top-3 right-3 bg-red-500 text-white text-xs font-medium px-2 py-1 rounded-full">
          Out of Stock
        </div>
      </Transition>
    </NuxtLink>

    <div class="p-4 relative">
      <NuxtLink
        :to="`/products/${product.product_id}`"
        class="block"
      >
        <h3 class="font-semibold text-gray-900 group-hover:text-primary-600 transition-colors line-clamp-2 min-h-[2.5rem]">
          {{ product.name }}
        </h3>
        <p v-if="product.description" class="mt-1 text-sm text-gray-500 line-clamp-2">
          {{ product.description }}
        </p>
      </NuxtLink>

      <div class="mt-4 flex items-center justify-between">
        <div>
          <span class="text-xl font-bold text-primary-600">
            ₹{{ Number(product.price).toFixed(2) }}
          </span>
        </div>

        <button
          @click.prevent="addToCart"
          :disabled="isAdding || product.stock === 0"
          class="flex items-center gap-2 bg-gradient-to-r from-primary-600 to-accent-600 text-white px-4 py-2 rounded-lg hover:from-primary-700 hover:to-accent-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all hover:scale-105 active:scale-95 shadow-md shadow-primary-500/25"
        >
          <svg v-if="isAdding" class="w-4 h-4 animate-spin" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
          </svg>
          <span class="font-medium text-sm">Add</span>
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
