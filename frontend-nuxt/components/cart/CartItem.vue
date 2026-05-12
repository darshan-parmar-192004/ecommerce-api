<script setup>
import { useCartStore } from '~/stores/cart'

const props = defineProps({
  item: { type: Object, required: true }
})

const cartStore = useCartStore()
</script>

<template>
  <div class="flex items-center gap-4 p-4 bg-surface-container-low rounded-lg">
    <div class="w-16 h-16 bg-surface-container rounded-lg flex items-center justify-center flex-shrink-0">
      <i class="pi pi-box text-2xl text-outline" />
    </div>
    <div class="flex-1 min-w-0">
      <p class="font-medium text-on_surface truncate">{{ item.product.name }}</p>
      <p class="text-sm text-primary font-medium">₹{{ Number(item.product.price).toFixed(2) }}</p>
    </div>
    <div class="flex items-center gap-2">
      <Button
        @click="cartStore.updateQuantity(item.product.product_id, item.quantity - 1)"
        icon="pi pi-minus"
        text
        rounded
        size="small"
      />
      <span class="w-8 text-center font-medium">{{ item.quantity }}</span>
      <Button
        @click="cartStore.updateQuantity(item.product.product_id, item.quantity + 1)"
        icon="pi pi-plus"
        text
        rounded
        size="small"
      />
      <Button
        @click="cartStore.removeItem(item.product.product_id)"
        icon="pi pi-trash"
        severity="danger"
        text
        rounded
        size="small"
      />
    </div>
  </div>
</template>
