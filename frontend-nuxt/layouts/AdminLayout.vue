/* AdminLayout.vue for Frontend-Nuxt */

<template>
  <div class="min-h-screen bg-gray-50 dark:bg-brand-900 flex">
    <AdminSidebar />
    <main class="flex-1 p-8 overflow-auto dark:bg-brand-900">
      <NuxtPage />
    </main>
    <ToastContainer />
  </div>
</template>

<script setup>
  // Similar logic to vue-prototype's AdminLayout
</script>

<style scoped>
  .page-enter-active, .page-leave-active {
    transition: opacity 0.3s ease, transform 0.3s ease;
  }
  .page-enter-from {
    opacity: 0;
    transform: translateX(10px);
  }
  .page-leave-to {
    opacity: 0;
    transform: translateX(-10px);
  }
</style>