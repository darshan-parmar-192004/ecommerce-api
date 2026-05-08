<script setup>
import { onMounted, ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useProducts } from '@/composables/useProducts'
import ProductCard from '@/components/product/ProductCard.vue'
import SkeletonCard from '@/components/ui/SkeletonCard.vue'
import { ArrowRight, Check, Package, CreditCard } from 'lucide-vue-next'

const router = useRouter()
const productsStore = useProducts()
  const featuredProducts = computed(() => Array.isArray(productsStore.products) ? productsStore.products.slice(0, 8) : [])
const loading = ref(true)

  onMounted(async () => {
  try {
    await productsStore.fetchProducts()
    // Poll for products to ensure they are loaded
    let attempts = 0
    while (productsStore.products.length === 0 && attempts < 10) {
      await new Promise(resolve => setTimeout(resolve, 100))
      attempts++
    }
  } finally {
    loading.value = false
  }
})
</script>
<template>
  <div class="home-page">
    <!-- Hero Section -->
    <section class="relative bg-gray-900 dark:bg-brand-900 text-white overflow-hidden">
      <div class="absolute inset-0">
        <img
          src="/placeholder.svg"
          alt="Hero background"
          class="w-full h-full object-cover opacity-30"
        />
        <div class="absolute inset-0 bg-gradient-to-r from-gray-900 via-gray-900/95 to-transparent dark:from-brand-900 dark:via-brand-900/95" />
      </div>

      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24 md:py-32">
        <div class="max-w-2xl">
          <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold tracking-tight">
            Discover Premium
            <span class="text-gray-300 dark:text-gray-400">Products</span>
          </h1>
          <p class="mt-6 text-lg md:text-xl text-gray-300 dark:text-gray-400 leading-relaxed">
            Experience the future of online shopping with our curated collection of premium products designed for modern living.
          </p>
            <div class="mt-10 flex flex-col sm:flex-row gap-4">
              <button
                @click="router.push('/products')"
                class="inline-flex items-center justify-center px-8 py-3 text-base font-medium bg-white text-gray-900 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-200 transition-all duration-200 hover:shadow-xl hover:-translate-y-0.5"
              >
                Shop Now
                <ArrowRight class="ml-2 w-5 h-5" />
              </button>
              <button
                @click="router.push('/auth/register')"
                class="inline-flex items-center justify-center px-8 py-3 text-base font-medium bg-transparent text-white border-2 border-white rounded-lg hover:bg-white hover:text-gray-900 transition-all duration-200"
              >
                Learn More
              </button>
            </div>
        </div>
      </div>
    </section>

    <!-- Featured Products -->
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="text-center mb-12">
        <h2 class="text-3xl font-bold text-gray-900 dark:text-gray-100">Featured Products</h2>
        <p class="mt-4 text-gray-600 dark:text-gray-400 max-w-2xl mx-auto">
          Discover our handpicked selection of premium products, crafted for exceptional quality and design.
        </p>
      </div>

      <div v-if="loading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <SkeletonCard v-for="n in 8" :key="n" />
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div
            v-for="(product, index) in featuredProducts"
            :key="product.productId"
            :style="{ animationDelay: `${index * 100}ms` }"
            class="animate-fade-in-up"
          >
            <div class="card cursor-pointer" @click="router.push({ name: 'ProductDetail', params: { id: product.productId } })">
              <div class="aspect-square overflow-hidden bg-gray-100 dark:bg-brand-700">
                <img
                  :src="product.images?.[0] || '/placeholder.svg'"
                  :alt="product.name"
                  class="w-full h-full object-cover transition-transform duration-700 hover:scale-110"
                />
              </div>
              <div class="p-5">
                <h3 class="text-sm font-medium text-gray-900 dark:text-gray-100 line-clamp-2">{{ product.name }}</h3>
                <div class="flex items-center justify-between mt-3">
                  <span class="text-xl font-bold text-gray-900 dark:text-gray-100">₹{{ product.price?.toFixed(2) }}</span>
                </div>
              </div>
            </div>
          </div>
      </div>

       <div class="text-center mt-12">
            <button
              @click="router.push('/products')"
              class="inline-flex items-center gap-2 px-6 py-3 text-sm font-medium text-gray-900 dark:text-gray-100 border border-gray-300 dark:border-brand-600 rounded-lg hover:bg-gray-50 dark:hover:bg-brand-700 transition-all duration-200"
            >
              View All Products
              <ArrowRight class="w-4 h-4" />
            </button>
       </div>
    </section>

    <!-- Features Section -->
    <section class="bg-gray-50 dark:bg-brand-900 py-16">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="text-center p-6">
            <div class="w-16 h-16 bg-gray-900 dark:bg-brand-700 rounded-full flex items-center justify-center mx-auto mb-4">
              <Check class="w-8 h-8 text-white" />
            </div>
            <h3 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-2">Premium Quality</h3>
            <p class="text-gray-600 dark:text-gray-400">Curated products that meet our highest standards for quality and durability.</p>
          </div>

          <div class="text-center p-6">
            <div class="w-16 h-16 bg-gray-900 dark:bg-brand-700 rounded-full flex items-center justify-center mx-auto mb-4">
              <Package class="w-8 h-8 text-white" />
            </div>
            <h3 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-2">Free Shipping</h3>
            <p class="text-gray-600 dark:text-gray-400">Enjoy free shipping on all orders, no minimum purchase required.</p>
          </div>

          <div class="text-center p-6">
            <div class="w-16 h-16 bg-gray-900 dark:bg-brand-700 rounded-full flex items-center justify-center mx-auto mb-4">
              <CreditCard class="w-8 h-8 text-white" />
            </div>
            <h3 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-2">Secure Payment</h3>
            <p class="text-gray-600 dark:text-gray-400">Your transactions are protected with bank-level security encryption.</p>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="bg-gray-900 dark:bg-brand-900 text-white py-16">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl font-bold mb-4">Ready to Elevate Your Experience?</h2>
        <p class="text-gray-300 dark:text-gray-400 mb-8 max-w-2xl mx-auto">
          Join thousands of satisfied customers who have transformed their lives with our premium products.
        </p>
        <button
          @click="router.push('/auth/register')"
          class="inline-flex items-center px-8 py-3 text-base font-medium bg-white text-gray-900 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-200 transition-all duration-200 hover:shadow-xl hover:-translate-y-0.5"
        >
          Get Started Today
        </button>
      </div>
    </section>
  </div>
</template>

 <style scoped>
 .home-page {
   min-height: 100vh;
 }
 </style>
