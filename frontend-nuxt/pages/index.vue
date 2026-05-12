<script setup>
definePageMeta({
})

import ProductCard from '~/components/products/ProductCard.vue'

const { products: productsApi } = useApi()

const { data, pending, error, refresh } = await useAsyncData(
  'home-products',
  () => productsApi.list({ page: 1, limit: 8 })
)

useSeoMeta({
  title: 'Home - E-Commerce Store',
  description: 'Shop the best products at our store'
})

const products = computed(() => data.value?.data || [])

const parallaxOffset = ref(0)

const handleScroll = () => {
  if (typeof window !== 'undefined') {
    parallaxOffset.value = window.scrollY * 0.15
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll, { passive: true })
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})

const floatingProducts = computed(() => products.value.slice(0, 3))
</script>

<template>
  <div>
    <section class="relative overflow-hidden bg-surface-container-low min-h-[85vh] flex items-center">
      <div 
        class="absolute inset-0 overflow-hidden pointer-events-none"
        :style="{ transform: `translateY(${parallaxOffset}px)` }"
      >
        <div 
          class="absolute -top-20 -right-20 w-[600px] h-[600px] rounded-full blur-3xl opacity-20" 
          style="background: radial-gradient(circle, #3e51fb, transparent); animation: float 8s ease-in-out infinite;"
        ></div>
        <div 
          class="absolute -bottom-40 -left-40 w-[500px] h-[500px] rounded-full blur-3xl opacity-15" 
          style="background: radial-gradient(circle, #1c31e3, transparent); animation: float 10s ease-in-out infinite reverse;"
        ></div>
        <div 
          class="absolute top-1/2 left-1/2 -translate-x-1/2 w-[800px] h-[800px] rounded-full blur-3xl opacity-10" 
          style="background: radial-gradient(circle, #6366f1, transparent); animation: pulse-glow 6s ease-in-out infinite;"
        ></div>
      </div>

      <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <div 
          v-for="(product, index) in floatingProducts" 
          :key="product.product_id"
          class="absolute hidden lg:block w-40 h-48 bg-surface-container-lowest rounded-xl shadow-2xl border border-outline-variant/20 overflow-hidden animate-float-delayed"
          :style="{
            top: `${15 + index * 25}%`,
            right: `${12 + index * 6}%`,
            animationDelay: `${index * 2}s`,
            transform: `translateY(${parallaxOffset * 0.5}px)`,
          }"
        >
          <div class="aspect-[4/3] bg-surface-container flex items-center justify-center">
            <i class="pi pi-box text-4xl text-outline/30" />
          </div>
          <div class="p-2 text-center">
            <p class="text-xs font-medium text-on_surface truncate">{{ product.name }}</p>
            <p class="text-xs text-primary font-bold">₹{{ Number(product.price).toFixed(0) }}</p>
          </div>
        </div>
      </div>

      <div class="relative w-full max-w-7xl mx-auto px-6 lg:px-8 py-28 md:py-40">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
          <div class="max-w-xl">
            <p class="text-xs font-semibold uppercase tracking-[0.2em] text-outline mb-6 animate-fade-in-up">
              Curated Collection
            </p>
            <h1 class="text-5xl md:text-6xl lg:text-7xl font-bold leading-[1.05] tracking-tight text-on_surface mb-8 animate-fade-in-up font-display" style="animation-delay: 100ms; letter-spacing: -0.02em;">
              Discover Products<br />
              <span class="text-primary">That Matter</span>
            </h1>
            <p class="text-lg md:text-xl text-on_surface_variant mb-10 max-w-xl animate-fade-in-up font-body leading-relaxed" style="animation-delay: 200ms;">
              Shop the latest trends with unbeatable prices and fast delivery
            </p>
            <div class="flex flex-wrap gap-4 animate-fade-in-up" style="animation-delay: 300ms;">
              <NuxtLink 
                to="/products" 
                class="group relative inline-flex items-center gap-2 bg-gradient-to-r from-primary to-primary-container text-white px-8 py-4 rounded-md font-semibold transition-all duration-400 hover:shadow-xl hover:shadow-primary/30 hover:scale-105 active:scale-[0.98]"
              >
                <span class="relative z-10">Shop Now</span>
                <i class="pi pi-arrow-right relative z-10 transition-transform group-hover:translate-x-1" />
                <span class="absolute inset-0 rounded-md bg-white/20 scale-0 group-hover:scale-100 transition-transform duration-300"></span>
              </NuxtLink>
              <NuxtLink 
                to="/cart" 
                class="inline-flex items-center gap-2 px-8 py-4 rounded-md font-semibold transition-all duration-300 hover:bg-surface-container-high bg-surface-container-highest text-on_surface hover:scale-105 active:scale-[0.98]"
              >
                View Cart
              </NuxtLink>
            </div>
          </div>
          <div class="hidden lg:block"></div>
        </div>
      </div>
    </section>

    <section class="max-w-7xl mx-auto px-6 lg:px-8 py-20">
      <div class="flex items-end justify-between mb-12">
        <div>
          <h2 class="text-3xl font-bold text-on_surface font-display tracking-tight">Featured Products</h2>
          <p class="text-on_surface_variant mt-2 font-body">Handpicked for you</p>
        </div>
        <NuxtLink to="/products" class="text-primary hover:text-primary-container font-semibold flex items-center gap-1 transition-colors duration-300 text-sm uppercase tracking-wider">
          View All
          <i class="pi pi-arrow-right text-xs" />
        </NuxtLink>
      </div>

      <div v-if="pending" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        <div v-for="i in 4" :key="i" class="rounded-md overflow-hidden" style="border: 1px solid rgba(197, 197, 217, 0.2);">
          <div class="aspect-[4/3] bg-surface-container-low relative overflow-hidden">
            <div class="absolute inset-0 animate-shimmer" style="background: linear-gradient(90deg, transparent, rgba(255,255,255,0.4), transparent); background-size: 200% 100%;"></div>
          </div>
          <div class="p-5 bg-surface-container-lowest">
            <div class="h-5 bg-surface-container rounded w-3/4 mb-3 animate-shimmer" style="background: linear-gradient(90deg, #eeeef0, #e8e8ea, #eeeef0); background-size: 200% 100%; animation: shimmer 1.5s infinite;"></div>
            <div class="h-4 bg-surface-container rounded w-1/2 mb-5 animate-shimmer" style="background: linear-gradient(90deg, #eeeef0, #e8e8ea, #eeeef0); background-size: 200% 100%; animation: shimmer 1.5s infinite;"></div>
            <div class="flex justify-between items-center">
              <div class="h-6 bg-surface-container rounded w-20 animate-shimmer" style="background: linear-gradient(90deg, #eeeef0, #e8e8ea, #eeeef0); background-size: 200% 100%; animation: shimmer 1.5s infinite;"></div>
              <div class="h-10 bg-surface-container rounded-md w-20 animate-shimmer" style="background: linear-gradient(90deg, #eeeef0, #e8e8ea, #eeeef0); background-size: 200% 100%; animation: shimmer 1.5s infinite;"></div>
            </div>
          </div>
        </div>
      </div>

      <div v-else-if="error" class="bg-surface-container-lowest rounded-md p-12 text-center shadow-ambient" style="border: 1px solid rgba(197, 197, 217, 0.2);">
        <i class="pi pi-exclamation-triangle text-5xl text-error mb-4" />
        <h3 class="mt-4 text-lg font-semibold text-on_surface font-display">Failed to load products</h3>
        <p class="mt-2 text-on_surface_variant font-body">Something went wrong. Please try again.</p>
        <Button @click="refresh" label="Try Again" class="mt-6" />
      </div>

      <div v-else-if="products.length === 0" class="bg-surface-container-lowest rounded-md p-16 text-center shadow-ambient" style="border: 1px solid rgba(197, 197, 217, 0.2);">
        <i class="pi pi-box text-5xl text-outline/30 mb-4" />
        <h2 class="mt-4 text-xl font-semibold text-on_surface font-display">No products available</h2>
        <p class="mt-2 text-on_surface_variant font-body">Check back soon for new products.</p>
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        <ProductCard 
          v-for="(product, index) in products" 
          :key="product.product_id" 
          :product="product"
          class="animate-stagger-fade"
          :style="{ animationDelay: `${index * 80}ms` }"
        />
      </div>
    </section>

    <section class="bg-surface-container-low py-20">
      <div class="max-w-7xl mx-auto px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-12">
          <div class="text-center group">
            <div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container">
              <i class="pi pi-truck text-white text-2xl" />
            </div>
            <h3 class="text-lg font-semibold text-on_surface mb-2 font-display">Free Shipping</h3>
            <p class="text-on_surface_variant font-body">On orders over ₹500</p>
          </div>
          <div class="text-center group">
            <div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container">
              <i class="pi pi-shield text-white text-2xl" />
            </div>
            <h3 class="text-lg font-semibold text-on_surface mb-2 font-display">Secure Payment</h3>
            <p class="text-on_surface_variant font-body">100% secure checkout</p>
          </div>
          <div class="text-center group">
            <div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container">
              <i class="pi pi-refresh text-white text-2xl" />
            </div>
            <h3 class="text-lg font-semibold text-on_surface mb-2 font-display">Easy Returns</h3>
            <p class="text-on_surface_variant font-body">30-day return policy</p>
          </div>
        </div>
      </div>
    </section>

    <section class="py-20 bg-surface">
      <div class="max-w-7xl mx-auto px-6 lg:px-8">
        <div class="bg-gradient-to-r from-primary to-primary-container rounded-2xl p-12 text-center relative overflow-hidden">
          <div class="absolute inset-0 bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNjAiIGhlaWdodD0iNjAiIHZpZXdCb3g9IjAgMCA2MCA2MCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48ZyBmaWxsPSJub25lIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiPjxwYXRoIGQ9Ik0zNiAxOGMtOS45NDEgMC0xOCA4LjA1OS0xOCAxOHM4LjA1OSAxOCAxOCAxOCAxOC04LjA1OSAxOC0xOC04LjA1OS0xOC0xOC0xOHptMCAzMmMtMy4zMTIgMC02LTIuNjg4LTYtNnMyLjY4OC02IDYtNiA2IDIuNjg4IDYgNi0yLjY4OCA2LTYgNnoiIGZpbGw9IiNmZmZmZmYiIGZpbGwtb3BhY2l0eT0iLjEiLz48L2c+PC9zdmc+')] opacity-20"></div>
          <div class="relative z-10">
            <h2 class="text-3xl font-bold text-white font-display mb-4">Stay Updated</h2>
            <p class="text-white/80 mb-8 max-w-xl mx-auto font-body">Subscribe to our newsletter for exclusive deals, new arrivals, and insider-only discounts.</p>
            <div class="flex flex-col sm:flex-row gap-4 max-w-md mx-auto">
              <InputText 
                type="email" 
                placeholder="Enter your email"
                class="flex-1 px-6 py-4 rounded-lg bg-white/20 backdrop-blur-sm text-white placeholder:text-white/60 border border-white/30 focus:outline-none focus:bg-white/30 focus:border-white"
              />
              <Button 
                type="submit"
                label="Subscribe"
                class="px-8 py-4 !bg-white !text-primary font-semibold rounded-lg transition-all duration-300 hover:bg-white/90 hover:scale-[1.02] active:scale-[0.98] white"
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.animate-fade-in-up {
  animation: fadeInUp 0.6s ease-out forwards;
  opacity: 0;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0) scale(1);
  }
  50% {
    transform: translateY(-25px) scale(1.02);
  }
}

@keyframes float-delayed {
  0%, 100% {
    transform: translateY(0) rotate(-5deg);
  }
  50% {
    transform: translateY(-20px) rotate(3deg);
  }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

.animate-float-delayed {
  animation: float-delayed 8s ease-in-out infinite;
}

@keyframes pulse-glow {
  0%, 100% {
    opacity: 0.1;
    transform: translate(-50%, -50%) scale(1);
  }
  50% {
    opacity: 0.2;
    transform: translate(-50%, -50%) scale(1.1);
  }
}

@keyframes shimmer {
  0% { background-position: -200% 0; }
  100% { background-position: 200% 0; }
}
</style>
