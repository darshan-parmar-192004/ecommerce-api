<script setup>
defineProps({
  count: { type: Number, default: 1 }
})
</script>

<template>
  <div 
    v-for="n in count" 
    :key="n" 
    class="group relative bg-gradient-to-br from-white/90 to-white/70 backdrop-blur-xl rounded-2xl border border-white/40 shadow-lg shadow-gray-200/50 overflow-hidden transition-all duration-500 hover:shadow-xl hover:shadow-gray-900/10"
  >
    <!-- Image Skeleton with Premium Shimmer -->
    <div class="relative w-full h-56 overflow-hidden bg-gradient-to-br from-gray-50/50 via-gray-50 to-gray-50/50">
      <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/60 to-transparent animate-shimmer"></div>
    </div>

    <!-- Content Skeleton -->
    <div class="p-6 space-y-4">
      <!-- Category Badge Skeleton -->
      <div class="h-3 w-16 rounded-full bg-gradient-to-r from-gray-200/80 via-gray-100/80 to-gray-200/80 animate-shimmer opacity-70"></div>
      
      <!-- Title Skeleton -->
      <div class="h-6 w-3/4 rounded-lg bg-gradient-to-r from-gray-200/80 via-gray-100/80 to-gray-200/80 animate-shimmer"></div>
      
      <!-- Description Skeleton -->
      <div class="h-4 rounded bg-gradient-to-r from-gray-200/80 via-gray-100/80 to-gray-200/80 animate-shimmer"></div>
      <div class="h-4 w-5/6 rounded bg-gradient-to-r from-gray-200/80 via-gray-100/80 to-gray-200/80 animate-shimmer"></div>
      
      <!-- Divider -->
      <div class="h-px w-full bg-gradient-to-r from-transparent via-gray-300/50 to-transparent"></div>
      
      <!-- Price & Button Row -->
      <div class="flex items-center justify-between pt-2">
        <div class="h-7 w-24 rounded-lg bg-gradient-to-r from-gray-100/80 via-gray-200/80 to-gray-100/80 animate-shimmer"></div>
        <div class="h-9 w-20 rounded-xl bg-gradient-to-r from-gray-200/80 via-gray-100/80 to-gray-200/80 animate-shimmer"></div>
      </div>
    </div>

    <!-- Premium hover glow effect -->
    <div class="absolute inset-0 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity duration-500 pointer-events-none">
      <div class="absolute inset-0 bg-gradient-to-br from-gray-900/5 via-transparent to-gray-700/5"></div>
    </div>
  </div>
</template>