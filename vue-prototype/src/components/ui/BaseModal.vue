<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  show: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: ''
  },
  size: {
    type: String,
    default: 'md',
    validator: (val) => ['sm', 'md', 'lg', 'xl'].includes(val)
  },
  persistent: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['close', 'confirm'])

const modalRef = ref(null)

const sizes = {
  sm: 'max-w-sm',
  md: 'max-w-md',
  lg: 'max-w-lg',
  xl: 'max-w-xl'
}

const handleBackdropClick = (e) => {
  if (!props.persistent && e.target === e.currentTarget) {
    closeModal()
  }
}

const closeModal = () => {
  if (!props.persistent) {
    emit('close')
  }
}

const handleKeydown = (e) => {
  if (e.key === 'Escape' && !props.persistent) {
    closeModal()
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
  if (props.show) {
    document.body.style.overflow = 'hidden'
  }
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
  document.body.style.overflow = ''
})

const handleConfirm = () => {
  emit('confirm')
}
</script>

<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="show"
        class="fixed inset-0 z-50 overflow-y-auto"
        @click="handleBackdropClick"
      >
        <div class="flex min-h-screen items-center justify-center p-4">
          <div class="fixed inset-0 bg-black/50 dark:bg-black/70" />

          <div
            ref="modalRef"
            :class="[sizes[size], 'relative w-full bg-white dark:bg-brand-800 rounded-xl shadow-2xl']"
            role="dialog"
            aria-modal="true"
            :aria-label="title || 'Dialog'"
          >
            <div class="flex items-center justify-between p-6 pb-0">
              <h3 class="text-lg font-semibold text-gray-900 dark:text-gray-100">
                {{ title }}
              </h3>
              <button
                v-if="!persistent"
                @click="closeModal"
                class="text-gray-400 hover:text-gray-600 dark:text-gray-500 dark:hover:text-gray-300 transition-colors"
                aria-label="Close dialog"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>

            <div class="p-6">
              <slot />
            </div>

            <div
              v-if="$slots.footer"
              class="flex items-center justify-end gap-3 p-6 pt-0"
            >
              <slot name="footer">
                <button
                  @click="closeModal"
                  class="btn btn-secondary"
                >
                  Cancel
                </button>
                <button
                  @click="handleConfirm"
                  class="btn btn-primary"
                >
                  Confirm
                </button>
              </slot>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>
