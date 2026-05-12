<script setup>
import AnimatedGrid from '~/components/common/AnimatedGrid.vue'
</script>

<template>
  <div class="min-h-screen w-full flex items-center justify-center p-4 relative overflow-hidden bg-surface">
    <AnimatedGrid variant="full" />
    <div class="relative z-10 w-full max-w-md">
      <slot />
    </div>
  </div>
</template>
