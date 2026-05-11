<script setup>
defineProps({
  variant: {
    type: String,
    default: 'full' // 'full' | 'hero' | 'minimal'
  }
})
</script>

<template>
  <div class="fixed inset-0 -z-10 overflow-hidden pointer-events-none">
    <!-- Base gradient background -->
    <div class="absolute inset-0 bg-gradient-to-br from-gray-50 via-white to-gray-100" />
    
    <!-- Animated grid pattern -->
    <div 
      :class="[
        'absolute inset-0 opacity-[0.03]',
        variant === 'full' ? 'bg-[linear-gradient(#1e293b_1px,transparent_1px),linear-gradient(90deg,#1e293b_1px,transparent_1px)] bg-[size:50px_50px] [animation:grid_20s_linear_infinite]' : '',
        variant === 'hero' ? 'bg-[linear-gradient(#1e293b_1px,transparent_1px),linear-gradient(90deg,#1e293b_1px,transparent_1px)] bg-[size:30px_30px]' : '',
        variant === 'minimal' ? 'bg-[linear-gradient(#1e293b_1px,transparent_1px),linear-gradient(90deg,#1e293b_1px,transparent_1px)] bg-[size:100px_100px]' : ''
      ]" 
    />
    
    <!-- Gradient orbs for depth -->
    <div class="absolute top-0 left-1/4 w-96 h-96 bg-primary-500/5 rounded-full blur-3xl animate-float" />
    <div class="absolute bottom-0 right-1/4 w-96 h-96 bg-accent-500/5 rounded-full blur-3xl animate-float" style="animation-delay: -3s" />
    
    <!-- Top highlight -->
    <div class="absolute top-0 left-0 right-0 h-px bg-gradient-to-r from-transparent via-primary-200/30 to-transparent" />
  </div>
</template>

<style scoped>
@keyframes grid {
  0% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(50px);
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0) scale(1);
    opacity: 0.5;
  }
  50% {
    transform: translateY(-20px) scale(1.05);
    opacity: 0.7;
  }
}
</style>
