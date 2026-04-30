<script setup>
import { useToastStore } from '@/stores/toast'

const toastStore = useToastStore()
</script>

<template>
  <div class="fixed bottom-4 right-4 z-50 space-y-2 max-w-sm">
    <div
      v-for="toast in toastStore.toasts"
      :key="toast.id"
      :class="{
        'bg-green-500': toast.type === 'success',
        'bg-red-500': toast.type === 'error',
        'bg-blue-500': toast.type === 'info',
        'bg-gray-800': toast.type === 'default'
      }"
      class="px-4 py-3 rounded-lg shadow-lg text-white flex items-center justify-between gap-3 animate-slide-up"
    >
      <span class="text-sm">{{ toast.message }}</span>
      <button @click="toastStore.removeToast(toast.id)" class="text-white/80 hover:text-white">
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
  </div>
</template>

<style scoped>
@keyframes slide-up {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.animate-slide-up {
  animation: slide-up 0.3s ease-out;
}
</style>