<script setup>
import { useRouter } from 'vue-router'
import { useMotion } from '@vueuse/motion'
import { AlertTriangle } from 'lucide-vue-next'

const router = useRouter()
</script>

<template>
  <div class="min-h-[60vh] flex items-center justify-center px-4">
    <div useMotion="{ initial: { opacity: 0, scale: 0.9 }, enter: { opacity: 1, scale: 1 } }" class="text-center">
      <div class="w-20 h-20 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-6">
        <AlertTriangle class="w-10 h-10 text-red-600" />
      </div>
      <h2 class="text-2xl font-bold text-gray-900">Unauthorized</h2>
      <p class="text-gray-600 mt-2 mb-8">You don't have permission to access this page.</p>
      <button @click="router.push('/')" class="btn-primary">
        Go Home
      </button>
    </div>
  </div>
</template>
