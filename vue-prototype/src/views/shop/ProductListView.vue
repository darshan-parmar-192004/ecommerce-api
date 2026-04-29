<script setup>
import { ref, onMounted } from 'vue'
// Import your existing components (e.g., ProductGrid) here
const products = ref([])
const isLoading = ref(true)
onMounted(async () => {
  try {
    // Your existing data fetching logic here
    // Example: products.value = await fetchProducts()
  } finally {
    isLoading.value = false
  }
})
</script>
<template>
  <!-- SINGLE ROOT WRAPPER (REQUIRED FOR TRANSITIONS) -->
  <div class="product-list-page">
    <h1>Products</h1>

    <div v-if="isLoading">Loading products...</div>
    <div v-else-if="products.length === 0">No products found.</div>
    <div v-else>
      <!-- Your existing product list content here -->
      <!-- <ProductGrid :products="products" /> -->
    </div>
  </div>
</template>
<style scoped>
.product-list-page {
  padding: 1rem;
  max-width: 1200px;
  margin: 0 auto;
}
h1 {
  margin-bottom: 2rem;
}
</style>
