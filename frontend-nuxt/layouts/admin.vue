<template>
  <div class="min-h-screen bg-gray-50 dark:bg-brand-900 flex">
    <AdminSidebar />
    <main class="flex-1 p-8 overflow-auto dark:bg-brand-900">
      <NuxtPage />
    </main>
    <Toast position="bottom-right" />
  </div>
</template>
