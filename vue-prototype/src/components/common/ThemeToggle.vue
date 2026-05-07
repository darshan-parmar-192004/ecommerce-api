<script setup>
import { Sun, Moon } from 'lucide-vue-next'
import { useTheme } from '@/composables/useTheme'
const { isDark, toggleTheme } = useTheme()
</script>
<template>
  <button @click="toggleTheme" class="p-2 text-gray-700 hover:text-gray-900 dark:text-gray-300 dark:hover:text-gray-100 transition-colors" :aria-label="isDark ? 'Switch to light mode' : 'Switch to dark mode'">
    <Moon v-if="!isDark" class="w-5 h-5" />
    <Sun v-else class="w-5 h-5" />
  </button>
</template>
