<script setup>
import { RouterView } from 'vue-router'
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <div class="min-h-screen bg-gray-50 dark:bg-brand-900 flex items-center justify-center px-4">
    <div class="w-full max-w-md">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100">E-Commerce</h1>
        <p class="text-gray-600 dark:text-gray-400 mt-2">Modern Shopping Experience</p>
      </div>
      <RouterView v-slot="{ Component }">
        <Transition name="fade" mode="out-in">
          <component v-if="Component" :is="Component" :key="route.path" />
        </Transition>
      </RouterView>
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
