<script setup>
import { useCartStore } from '@/stores/cart'

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})

const cartStore = useCartStore()

const decreaseQuantity = () => {
  cartStore.updateQuantity(props.item.id, props.item.quantity - 1)
}

const increaseQuantity = () => {
  cartStore.updateQuantity(props.item.id, props.item.quantity + 1)
}

const removeItem = () => {
  cartStore.removeFromCart(props.item.id)
}
</script>

<template>
  <div class="flex gap-4 p-4 bg-gray-50 dark:bg-brand-700/50 rounded-lg">
    <img :src="item.image || '/placeholder.jpg'" :alt="item.name" class="w-20 h-20 object-cover rounded-md" />
    <div class="flex-1">
      <h3 class="text-sm font-medium text-gray-900 dark:text-gray-100">{{ item.name }}</h3>
      <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">₹{{ (item.price * item.quantity).toFixed(2) }}</p>
      <div class="flex items-center gap-2 mt-2">
        <button
          @click="decreaseQuantity()"
          class="w-6 h-6 flex items-center justify-center border border-gray-300 dark:border-brand-600 rounded hover:bg-gray-100 dark:hover:bg-brand-700"
          :aria-label="`Decrease quantity of ${item.name}`"
        >
          -
        </button>
        <span class="text-sm w-8 text-center text-gray-900 dark:text-gray-100">{{ item.quantity }}</span>
        <button
          @click="increaseQuantity()"
          class="w-6 h-6 flex items-center justify-center border border-gray-300 dark:border-brand-600 rounded hover:bg-gray-100 dark:hover:bg-brand-700"
          :aria-label="`Increase quantity of ${item.name}`"
        >
          +
        </button>
        <button
          @click="removeItem()"
          class="ml-auto text-sm text-red-600 hover:text-red-700 dark:text-red-400 dark:hover:text-red-300"
          :aria-label="`Remove ${item.name} from cart`"
        >
          Remove
        </button>
      </div>
    </div>
  </div>
</template>
