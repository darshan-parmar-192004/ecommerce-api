<script setup>
import { useRouter } from 'vue-router'
import { useMotion } from '@vueuse/motion'

const router = useRouter()
</script>

<template>
  <div class="min-h-[60vh] flex items-center justify-center px-4">
    <div useMotion="{ initial: { opacity: 0, scale: 0.9 }, enter: { opacity: 1, scale: 1 } }" class="text-center">
      <div class="w-20 h-20 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-6">
        <svg class="w-10 h-10 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
      </div>
      <h2 class="text-2xl font-bold text-gray-900">Unauthorized</h2>
      <p class="text-gray-600 mt-2 mb-8">You don't have permission to access this page.</p>
      <button @click="router.push('/')" class="btn-primary">
        Go Home
      </button>
    </div>
  </div>
</template>
