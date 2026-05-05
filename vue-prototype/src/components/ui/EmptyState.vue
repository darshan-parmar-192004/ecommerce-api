<script setup>
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
      <svg
        v-if="icon === 'shopping-bag'"
        class="w-full h-full"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="1.5"
          d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"
        />
      </svg>
      <svg
        v-else-if="icon === 'shopping-cart'"
        class="w-full h-full"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="1.5"
          d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 100 4 2 2 0 000-4z"
        />
      </svg>
      <svg
        v-else-if="icon === 'orders'"
        class="w-full h-full"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="1.5"
          d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"
        />
      </svg>
      <svg
        v-else
        class="w-full h-full"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="1.5"
          d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4"
        />
      </svg>
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
