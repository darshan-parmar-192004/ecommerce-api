<script setup>
import { ShoppingBag, ShoppingCart, ClipboardList, Inbox } from 'lucide-vue-next'

defineProps({
  title: {
    type: String,
    default: 'No items found'
  },
  description: {
    type: String,
    default: ''
  },
  icon: {
    type: String,
    default: 'shopping-bag'
  },
  actionText: {
    type: String,
    default: ''
  },
  actionLink: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['action'])
</script>

<template>
  <div class="flex flex-col items-center justify-center py-12 px-4">
    <div class="w-24 h-24 mb-6 text-gray-300 dark:text-gray-600">
      <ShoppingBag
        v-if="icon === 'shopping-bag'"
        class="w-full h-full"
      />
      <ShoppingCart
        v-else-if="icon === 'shopping-cart'"
        class="w-full h-full"
      />
      <ClipboardList
        v-else-if="icon === 'orders'"
        class="w-full h-full"
      />
      <Inbox
        v-else
        class="w-full h-full"
      />
    </div>

    <h3 class="text-lg font-medium text-gray-900 dark:text-gray-100 mb-2">
      {{ title }}
    </h3>

    <p
      v-if="description"
      class="text-sm text-gray-500 dark:text-gray-400 text-center mb-6 max-w-sm"
    >
      {{ description }}
    </p>

    <RouterLink
      v-if="actionText && actionLink"
      :to="actionLink"
      class="btn btn-primary"
    >
      {{ actionText }}
    </RouterLink>

    <button
      v-else-if="actionText"
      @click="emit('action')"
      class="btn btn-primary"
    >
      {{ actionText }}
    </button>
  </div>
</template>
