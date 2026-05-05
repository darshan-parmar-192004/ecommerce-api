<script setup>
import { useToastStore } from '@/stores/toast'
import { onMounted, onUnmounted } from 'vue'

const toastStore = useToastStore()
let autoDismissInterval = null

onMounted(() => {
  autoDismissInterval = setInterval(() => {
    const now = Date.now()
    toastStore.toasts.forEach(toast => {
      if (now - toast.timestamp > 5000) {
        toastStore.removeToast(toast.id)
      }
    })
  }, 1000)
})

onUnmounted(() => {
  if (autoDismissInterval) {
    clearInterval(autoDismissInterval)
  }
})
</script>

<template>
  <div class="fixed bottom-4 right-4 z-50 space-y-2 max-w-sm" role="alert" aria-live="polite">
    <TransitionGroup name="toast">
      <div
        v-for="toast in toastStore.toasts"
        :key="toast.id"
        :class="{
          'bg-green-500': toast.type === 'success',
          'bg-red-500': toast.type === 'error',
          'bg-blue-500': toast.type === 'info',
          'bg-gray-800 dark:bg-brand-700': toast.type === 'default'
        }"
        class="px-4 py-3 rounded-lg shadow-lg text-white flex items-center justify-between gap-3"
      >
        <span class="text-sm">{{ toast.message }}</span>
        <button @click="toastStore.removeToast(toast.id)" class="text-white/80 hover:text-white transition-colors">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped>
.toast-enter-active {
  animation: slide-up 0.3s ease-out;
}
.toast-leave-active {
  animation: slide-down 0.3s ease-in;
}
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
@keyframes slide-down {
  from {
    opacity: 1;
    transform: translateY(0);
  }
  to {
    opacity: 0;
    transform: translateY(20px);
  }
}
</style>