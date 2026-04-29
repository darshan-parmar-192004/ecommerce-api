<script setup>
import Navbar from '@/components/layout/Navbar.vue'
import Footer from '@/components/layout/Footer.vue'
// Transition hook handlers (JavaScript animation hooks as you requested)
const onBeforeEnter = (el) => {
  el.style.opacity = 0
}
const onEnter = (el, done) => {
  try {
    el.style.transition = 'opacity 0.25s ease'
    // Force reflow to trigger transition
    void el.offsetWidth
    el.style.opacity = 1
    // Listen for transition end, with fallback timeout
    const onTransitionEnd = () => {
      el.removeEventListener('transitionend', onTransitionEnd)
      done()
    }
    el.addEventListener('transitionend', onTransitionEnd)

    // Fallback: ensure done() is called even if transition event fails
    setTimeout(done, 300)
  } catch (err) {
    // Never leave transition hanging
    done()
  }
}
const onAfterEnter = (el) => {
  el.style.transition = ''
}
const onBeforeLeave = (el) => {
  el.style.opacity = 1
}
const onLeave = (el, done) => {
  try {
    el.style.transition = 'opacity 0.25s ease'
    void el.offsetWidth
    el.style.opacity = 0
    const onTransitionEnd = () => {
      el.removeEventListener('transitionend', onTransitionEnd)
      done()
    }
    el.addEventListener('transitionend', onTransitionEnd)

    setTimeout(done, 300)
  } catch (err) {
    done()
  }
}
const onAfterLeave = (el) => {
  el.style.transition = ''
}
</script>
<template>
  <div class="default-layout">
    <Navbar />

    <main class="main-content">
      <!-- Scoped router-view slot: only renders when Component is resolved -->
      <router-view v-slot="{ Component }">
        <Transition mode="out-in" @before-enter="onBeforeEnter" @enter="onEnter" @after-enter="onAfterEnter" @before-leave="onBeforeLeave" @leave="onLeave" @after-leave="onAfterLeave">
          <!-- Prevent null components from breaking the transition -->
          <component v-if="Component" :is="Component" class="page-content" />
          <!-- Optional loading fallback -->
          <div v-else class="loading-fallback">Loading page...</div>
        </Transition>
      </router-view>
    </main>
    <Footer />
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
.loading-fallback {
  text-align: center;
  padding: 4rem 2rem;
  color: #666;
}
</style>
