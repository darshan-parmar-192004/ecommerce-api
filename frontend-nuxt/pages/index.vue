<script setup>
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
</script>

<template>
  <div>
    <!-- Hero Section -->
    <section class="relative overflow-hidden bg-surface-container-low">
      <div class="absolute inset-0 overflow-hidden">
        <div class="absolute -top-40 -right-40 w-[500px] h-[500px] rounded-full blur-3xl animate-float opacity-20" style="background: radial-gradient(circle, #3e51fb, transparent);"></div>
        <div class="absolute -bottom-40 -left-40 w-[400px] h-[400px] rounded-full blur-3xl animate-float opacity-15" style="background: radial-gradient(circle, #1c31e3, transparent); animation-delay: -3s;"></div>
      </div>

      <div class="relative max-w-7xl mx-auto px-6 lg:px-8 py-28 md:py-40">
        <div class="max-w-3xl">
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
              class="inline-flex items-center gap-2 bg-gradient-to-r from-primary to-primary-container text-white px-8 py-4 rounded-md font-semibold transition-all duration-400 hover:shadow-lg hover:scale-[1.02] active:scale-[0.98] shadow-lg shadow-primary/25"
            >
              Shop Now
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3" />
              </svg>
            </NuxtLink>
            <NuxtLink 
              to="/cart" 
              class="inline-flex items-center gap-2 px-8 py-4 rounded-md font-semibold transition-all duration-300 hover:bg-surface-container-high bg-surface-container-highest text-on_surface"
            >
              View Cart
            </NuxtLink>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Products -->
    <section class="max-w-7xl mx-auto px-6 lg:px-8 py-20">
      <div class="flex items-end justify-between mb-12">
        <div>
          <h2 class="text-3xl font-bold text-on_surface font-display tracking-tight">Featured Products</h2>
          <p class="text-on_surface_variant mt-2 font-body">Handpicked for you</p>
        </div>
        <NuxtLink to="/products" class="text-primary hover:text-primary-container font-semibold flex items-center gap-1 transition-colors duration-300 text-sm uppercase tracking-wider">
          View All
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3" />
          </svg>
        </NuxtLink>
      </div>

      <!-- Skeleton Shimmer Loading -->
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

      <!-- Error State -->
      <div v-else-if="error" class="bg-surface-container-lowest rounded-md p-12 text-center shadow-ambient" style="border: 1px solid rgba(197, 197, 217, 0.2);">
        <svg class="mx-auto h-16 w-16 text-error" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v3.75m9-.75a9 9 0 11-18 0 9 9 0 0118 0zm-9 3.75h.008v.008H12v-.008z" />
        </svg>
        <h3 class="mt-4 text-lg font-semibold text-on_surface font-display">Failed to load products</h3>
        <p class="mt-2 text-on_surface_variant font-body">Something went wrong. Please try again.</p>
        <button @click="refresh" class="mt-6 px-6 py-3 bg-gradient-to-r from-primary to-primary-container text-white rounded-md font-medium transition-all duration-300 hover:shadow-lg shadow-lg shadow-primary/25">
          Try Again
        </button>
      </div>

      <!-- Empty State -->
      <div v-else-if="products.length === 0" class="bg-surface-container-lowest rounded-md p-16 text-center shadow-ambient" style="border: 1px solid rgba(197, 197, 217, 0.2);">
        <svg class="mx-auto h-24 w-24 text-outline/30" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1">
          <path stroke-linecap="round" stroke-linejoin="round" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
        </svg>
        <h2 class="mt-4 text-xl font-semibold text-on_surface font-display">No products available</h2>
        <p class="mt-2 text-on_surface_variant font-body">Check back soon for new products.</p>
      </div>

      <!-- Product Grid -->
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

    <!-- Features Section -->
    <section class="bg-surface-container-low py-20">
      <div class="max-w-7xl mx-auto px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-12">
          <div class="text-center group">
            <div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M8.25 18.75a1.5 1.5 0 01-3 0m3 0a1.5 1.5 0 00-3 0m3 0h6m-9 0H3.375a1.125 1.125 0 01-1.125-1.125V14.25m17.25 4.5a1.5 1.5 0 01-3 0m3 0a1.5 1.5 0 00-3 0m3 0h1.125c.621 0 1.129-.504 1.09-1.124a17.902 17.902 0 00-3.213-9.193 2.056 2.056 0 00-1.58-.86H14.25M16.5 18.75h-2.25m0-11.177v-.958c0-.568-.422-1.048-.987-1.106a48.554 48.554 0 00-10.026 0 1.106 1.106 0 00-.987 1.106v7.635m12-6.677v6.677m0 4.5v-4.5m0 0h-12" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-on_surface mb-2 font-display">Free Shipping</h3>
            <p class="text-on_surface_variant font-body">On orders over ₹500</p>
          </div>
          <div class="text-center group">
            <div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75m-3-7.036A11.959 11.959 0 013.598 6 11.99 11.99 0 003 9.749c0 5.592 3.824 10.29 9 11.623 5.176-1.332 9-6.03 9-11.622 0-1.31-.21-2.571-.598-3.751h-1.564A11.959 11.959 0 0112 2.714z" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-on_surface mb-2 font-display">Secure Payment</h3>
            <p class="text-on_surface_variant font-body">100% secure checkout</p>
          </div>
          <div class="text-center group">
            <div class="w-16 h-16 rounded-full flex items-center justify-center mx-auto mb-5 transition-transform duration-400 group-hover:scale-110 bg-gradient-to-r from-primary to-primary-container">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0013.803-3.7M4.031 9.865a8.25 8.25 0 0113.803-3.7l3.181 3.182m0-4.991v4.99" />
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-on_surface mb-2 font-display">Easy Returns</h3>
            <p class="text-on_surface_variant font-body">30-day return policy</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Newsletter Section -->
    <section class="py-20 bg-surface">
      <div class="max-w-7xl mx-auto px-6 lg:px-8">
        <div class="bg-gradient-to-r from-primary to-primary-container rounded-2xl p-12 text-center relative overflow-hidden">
          <div class="absolute inset-0 bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNjAiIGhlaWdodD0iNjAiIHZpZXdCb3g9IjAgMCA2MCA2MCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48ZyBmaWxsPSJub25lIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiPjxwYXRoIGQ9Ik0zNiAxOGMtOS45NDEgMC0xOCA4LjA1OS0xOCAxOHM4LjA1OSAxOCAxOCAxOCAxOC04LjA1OSAxOC0xOC04LjA1OS0xOC0xOC0xOHptMCAzMmMtMy4zMTIgMC02LTIuNjg4LTYtNnMyLjY4OC02IDYtNiA2IDIuNjg4IDYgNi0yLjY4OCA2LTYgNnoiIGZpbGw9IiNmZmZmZmYiIGZpbGwtb3BhY2l0eT0iLjEiLz48L2c+PC9zdmc+')] opacity-20"></div>
          <div class="relative z-10">
            <h2 class="text-3xl font-bold text-white font-display mb-4">Stay Updated</h2>
            <p class="text-white/80 mb-8 max-w-xl mx-auto font-body">Subscribe to our newsletter for exclusive deals, new arrivals, and insider-only discounts.</p>
            <form class="flex flex-col sm:flex-row gap-4 max-w-md mx-auto" @submit.prevent>
              <input 
                type="email" 
                placeholder="Enter your email"
                class="flex-1 px-6 py-4 rounded-lg bg-white/20 backdrop-blur-sm text-white placeholder:text-white/60 border border-white/30 focus:outline-none focus:bg-white/30 focus:border-white"
              />
              <button 
                type="submit"
                class="px-8 py-4 bg-white text-primary font-semibold rounded-lg transition-all duration-300 hover:bg-white/90 hover:scale-[1.02] active:scale-[0.98]"
              >
                Subscribe
              </button>
            </form>
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
    transform: translateY(-20px) scale(1.05);
  }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

@keyframes shimmer {
  0% { background-position: -200% 0; }
  100% { background-position: 200% 0; }
}
</style>
