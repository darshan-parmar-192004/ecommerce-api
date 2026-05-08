<script setup>
const props = defineProps({
  count: {
    type: Number,
    default: 1
  }
})
</script>

<template>
  <div v-for="i in count" :key="i" class="overflow-hidden rounded-lg bg-surface-container-lowest">
    <div class="aspect-[4/3] bg-surface-container animate-pulse" />
    <div class="p-5 space-y-3">
      <div class="h-4 bg-surface-container rounded w-3/4 animate-pulse" />
      <div class="h-3 bg-surface-container rounded w-1/2 animate-pulse" />
      <div class="flex justify-between pt-2">
        <div class="h-5 bg-surface-container rounded w-20 animate-pulse" />
        <div class="h-8 bg-surface-container rounded w-16 animate-pulse" />
      </div>
    </div>
  </div>
</template>