<script setup>
useHead({
  title: 'Unauthorized - The Curator'
})
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-surface">
    <div class="text-center p-8">
      <div class="mb-8">
        <svg class="w-24 h-24 mx-auto text-error" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
      </div>
      <h1 class="text-4xl font-bold text-on_surface mb-4">Access Denied</h1>
      <p class="text-on_surface_variant mb-8 max-w-md mx-auto">
        You don't have permission to access this page. Please contact an administrator if you believe this is an error.
      </p>
      <NuxtLink 
        to="/" 
        class="inline-flex items-center px-6 py-3 rounded-md font-medium text-white bg-gradient-to-r from-primary to-primary-container shadow-glow hover:shadow-xl transition-all duration-300"
      >
        Go to Homepage
      </NuxtLink>
    </div>
  </div>
</template>