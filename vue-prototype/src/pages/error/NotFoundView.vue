<script setup>
import { useRouter } from 'vue-router'
import { useMotion } from '@vueuse/motion'

const router = useRouter()
</script>

<template>
  <div class="min-h-[60vh] flex items-center justify-center px-4">
    <div useMotion="{ initial: { opacity: 0, scale: 0.9 }, enter: { opacity: 1, scale: 1 } }" class="text-center">
      <h1 class="text-9xl font-bold text-gray-200">404</h1>
      <h2 class="text-2xl font-bold text-gray-900 mt-4">Page Not Found</h2>
      <p class="text-gray-600 mt-2 mb-8">The page you're looking for doesn't exist.</p>
      <button @click="router.push('/')" class="btn-primary">
        Go Home
      </button>
    </div>
  </div>
</template>
