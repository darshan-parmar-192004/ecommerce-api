<script setup>
const { toasts, removeToast } = useToast()
</script>

<template>
  <Teleport to="body">
    <div class="fixed bottom-6 right-6 z-50 flex flex-col gap-3">
      <TransitionGroup name="toast">
        <div
          v-for="toast in toasts"
          :key="toast.id"
          :class="[
            'px-5 py-4 rounded-xl shadow-ambient flex items-center gap-3 min-w-[320px] max-w-md',
            'border backdrop-blur-sm bg-surface-container-lowest',
            toast.type === 'success' ? 'border-green-500/20 text-on_surface' :
            toast.type === 'error' ? 'border-error/20 text-on_surface bg-error-container/10' :
            'border-primary/20 text-on_surface'
          ]"
        >
          <div :class="[
            'w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0',
            toast.type === 'success' ? 'bg-green-500/10' :
            toast.type === 'error' ? 'bg-error/10' :
            'bg-primary/10'
          ]">
            <svg v-if="toast.type === 'success'" class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
            <svg v-else-if="toast.type === 'error'" class="w-5 h-5 text-error" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
            <svg v-else class="w-5 h-5 text-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>

          <p class="text-sm font-medium flex-1 text-on_surface">{{ toast.message }}</p>

          <button
            @click="removeToast(toast.id)"
            class="p-1.5 rounded-lg hover:bg-surface-container transition-colors"
          >
            <svg class="w-4 h-4 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<style scoped>
.toast-enter-active {
  transition: all 0.3s ease-out;
}

.toast-leave-active {
  transition: all 0.2s ease-in;
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

.toast-move {
  transition: transform 0.3s ease;
}
</style>
