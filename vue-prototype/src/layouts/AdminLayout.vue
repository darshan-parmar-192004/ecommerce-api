<template>
<div class="min-h-screen bg-gray-50 dark:bg-brand-900 flex">
  <AdminSidebar />
  <main class="flex-1 p-8 overflow-auto dark:bg-brand-900">
    <router-view v-slot="{ Component }">
      <Transition name="page" mode="out-in">
        <component :is="Component" />
      </Transition>
    </router-view>
  </main>
</div>
</template>

<script setup>
import AdminSidebar from '@/components/layout/AdminSidebar.vue'
</script>

<style scoped>
.page-enter-active,
.page-leave-active {
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
