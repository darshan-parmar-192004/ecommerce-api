<template>
  <div class="min-h-screen bg-surface flex">
    <AdminSidebar />
    <main class="flex-1 p-8 overflow-auto">
      <NuxtPage />
    </main>
    <Toast position="bottom-right" />
  </div>
</template>
