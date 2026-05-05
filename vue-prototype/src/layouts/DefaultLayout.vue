<script setup>
import Navbar from '@/components/layout/Navbar.vue'
import Footer from '@/components/layout/Footer.vue'
import ToastContainer from '@/components/common/ToastContainer.vue'
import CartDrawer from '@/components/layout/CartDrawer.vue'
</script>

<template>
  <div class="default-layout dark:bg-brand-900">
    <Navbar />

    <main class="main-content dark:bg-brand-900">
      <router-view v-slot="{ Component }">
        <Transition name="page" mode="out-in">
          <component :is="Component" class="page-content" />
        </Transition>
      </router-view>
    </main>

    <Footer />
    <ToastContainer />
    <CartDrawer />
  </div>
</template>
<style scoped>
.default-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
.main-content {
  flex: 1;
  padding: 1rem;
  max-width: 1400px;
  width: 100%;
  margin: 0 auto;
}
.page-content {
  width: 100%;
}
/* Page transition animations */
.page-enter-active,
.page-leave-active {
  transition: opacity 0.25s ease;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
}
</style>
