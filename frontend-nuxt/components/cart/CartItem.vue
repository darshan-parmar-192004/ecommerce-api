<script setup>
import { Plus, Minus, Trash2, ShoppingBag } from 'lucide-vue-next'
import { useCartStore } from '~/stores/cart'

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})

const cartStore = useCartStore()

const formatPrice = (price) => {
  return `₹${Number(price).toFixed(2)}`
}
</script>

<template>
  <div class="flex gap-4 p-4 rounded-xl bg-surface-container-low/80 backdrop-blur-sm transition-all duration-300 hover:bg-surface-container hover:shadow-lg group">
    <div class="w-24 h-24 bg-surface-container rounded-lg flex-shrink-0 overflow-hidden flex items-center justify-center shadow-inner">
      <ShoppingBag class="w-12 h-12 text-outline/30 group-hover:scale-110 transition-transform duration-300" />
    </div>

    <div class="flex-1 min-w-0">
      <h3 class="font-semibold text-on_surface font-display truncate group-hover:text-primary transition-colors duration-200">
        {{ item.product.name }}
      </h3>
      <p class="text-primary font-bold mt-1 font-display">
        {{ formatPrice(item.product.price) }}
      </p>

      <div class="flex items-center gap-2 mt-3">
        <button
          @click="cartStore.updateQuantity(item.product.product_id, item.quantity - 1)"
          class="w-8 h-8 rounded-lg bg-surface-container hover:bg-surface-container-high flex items-center justify-center transition-all duration-200 hover:scale-110 active:scale-95"
        >
          <Minus class="w-4 h-4 text-on_surface" />
        </button>
        <span class="w-10 text-center font-medium text-on_surface bg-surface-container-low rounded">{{ item.quantity }}</span>
        <button
          @click="cartStore.updateQuantity(item.product.product_id, item.quantity + 1)"
          class="w-8 h-8 rounded-lg bg-surface-container hover:bg-surface-container-high flex items-center justify-center transition-all duration-200 hover:scale-110 active:scale-95"
        >
          <Plus class="w-4 h-4 text-on_surface" />
        </button>
        <button
          @click="cartStore.removeItem(item.product.product_id)"
          class="ml-auto text-error hover:text-error/80 transition-all duration-200 hover:scale-105 text-sm font-medium flex items-center gap-1"
        >
          <Trash2 class="w-4 h-4" />
          Remove
        </button>
      </div>
    </div>
  </div>
</template>