<script setup>
useHead({
  title: 'Unauthorized - The Curator'
})
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-surface">
    <div class="text-center p-8">
      <div class="mb-8">
        <i class="pi pi-exclamation-triangle text-6xl text-error" />
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
