<script setup>
import { useCartStore } from '~/stores/cart'

const route = useRoute()
const cartStore = useCartStore()

const { data: product, pending, error } = await useFetch(`/api/products/${route.params.id}`)

useSeoMeta({
  title: () => product.value ? `${product.value.name} - E-Commerce Store` : 'Product',
  description: () => product.value?.description || 'Product details'
})

const quantity = ref(1)
const added = ref(false)
const isZoomed = ref(false)
const mousePosition = ref({ x: 0, y: 0 })

const reviews = ref([
  { id: 1, name: 'John D.', rating: 5, comment: 'Excellent product! Very satisfied with the quality.', date: '2024-01-15' },
  { id: 2, name: 'Sarah M.', rating: 4, comment: 'Good value for money. Fast delivery.', date: '2024-01-10' },
  { id: 3, name: 'Alex K.', rating: 5, comment: 'Highly recommend! Exceeded expectations.', date: '2024-01-05' }
])

const incrementQuantity = () => {
  quantity.value++
}

const decrementQuantity = () => {
  if (quantity.value > 1) {
    quantity.value--
  }
}

const addToCart = () => {
  if (product.value) {
    for (let i = 0; i < quantity.value; i++) {
      cartStore.addItem(product.value)
    }
    added.value = true
    cartStore.openCart()
    setTimeout(() => {
      added.value = false
    }, 2000)
  }
}

const handleMouseMove = (e) => {
  if (!isZoomed.value) return
  const rect = e.target.getBoundingClientRect()
  mousePosition.value = {
    x: ((e.clientX - rect.left) / rect.width) * 100,
    y: ((e.clientY - rect.top) / rect.height) * 100
  }
}

const averageRating = computed(() => {
  if (!reviews.value.length) return 0
  return reviews.value.reduce((sum, r) => sum + r.rating, 0) / reviews.value.length
})
</script>

<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 bg-surface">
    <NuxtLink to="/products" class="inline-flex items-center text-sm text-on_surface_variant hover:text-primary mb-6 transition-colors">
      <i class="pi pi-arrow-left mr-1" />
      Back to Products
    </NuxtLink>

    <div v-if="pending" class="animate-pulse">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="bg-surface-container aspect-square rounded-lg" />
        <div>
          <div class="h-8 bg-surface-container rounded w-3/4 mb-4" />
          <div class="h-6 bg-surface-container rounded w-1/4 mb-6" />
          <div class="h-4 bg-surface-container rounded w-full mb-2" />
          <div class="h-4 bg-surface-container rounded w-2/3" />
        </div>
      </div>
    </div>

    <div v-else-if="error" class="bg-surface-container-lowest rounded-xl shadow-ambient p-8 text-center">
      <i class="pi pi-exclamation-triangle text-5xl text-error mb-4" />
      <h1 class="mt-4 text-2xl font-bold text-on_surface mb-2">Product Not Found</h1>
      <p class="text-on_surface_variant mb-6">The product you're looking for doesn't exist.</p>
      <NuxtLink to="/products">
        <Button label="Browse Products" icon="pi pi-shopping-bag" />
      </NuxtLink>
    </div>

    <div v-else-if="product" class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div class="relative group">
        <div 
          class="bg-surface-container aspect-square rounded-lg flex items-center justify-center overflow-hidden cursor-zoom-in"
          @mouseenter="isZoomed = true"
          @mouseleave="isZoomed = false"
          @mousemove="handleMouseMove"
        >
          <div 
            class="w-full h-full flex items-center justify-center transition-transform duration-200"
            :style="{
              transform: isZoomed ? 'scale(1.5)' : 'scale(1)',
              transformOrigin: `${mousePosition.x}% ${mousePosition.y}%`
            }"
          >
            <i class="pi pi-box text-8xl text-outline" />
          </div>
          
          <div v-if="isZoomed" class="absolute inset-0 pointer-events-none border-2 border-primary/50 rounded-lg"></div>
        </div>
        
        <Button
          icon="pi pi-heart"
          severity="danger"
          text
          rounded
          class="absolute top-4 right-4"
        />
        
        <Tag v-if="product.stock === 0" value="Out of Stock" severity="danger" class="absolute top-4 left-4" />
      </div>

      <div>
        <h1 class="text-3xl font-bold text-on_surface font-display mb-2">
          {{ product.name }}
        </h1>
        
        <div class="flex items-center gap-4 mb-4">
          <div class="flex items-center gap-1">
            <Rating :modelValue="averageRating" readonly :cancel="false" />
            <span class="ml-2 text-sm text-on_surface_variant">{{ averageRating.toFixed(1) }} ({{ reviews.length }} reviews)</span>
          </div>
        </div>

        <p class="text-3xl font-semibold text-primary mb-4">
          ₹ {{ Number(product.price).toFixed(2) }}
        </p>

        <p v-if="product.description" class="text-on_surface_variant mb-6">
          {{ product.description }}
        </p>

        <div class="mb-6">
          <p class="text-sm text-on_surface_variant">
            <span v-if="product.stock > 10" class="text-green-600 flex items-center gap-1">
              <i class="pi pi-check-circle" />
              In Stock ({{ product.stock }} available)
            </span>
            <span v-else-if="product.stock > 0" class="text-yellow-600 flex items-center gap-1">
              <i class="pi pi-exclamation-circle" />
              Only {{ product.stock }} left
            </span>
          </p>
        </div>

        <div class="mb-6">
          <label class="block text-sm font-semibold text-on_surface_variant mb-3">Quantity</label>
          <InputNumber
            v-model="quantity"
            :min="1"
            :max="product.stock"
            showButtons
            class="w-auto"
          />
        </div>

        <Button
          @click="addToCart"
          :label="added ? 'Added to Cart!' : 'Add to Cart'"
          :icon="added ? 'pi pi-check' : 'pi pi-shopping-cart'"
          :severity="added ? 'success' : 'primary'"
          class="!py-3 !px-8"
        />

        <div class="border-t border-outline-variant/20 pt-8 mt-8">
          <h3 class="text-xl font-bold text-on_surface font-display mb-6">Customer Reviews</h3>
          
          <div class="space-y-4">
            <div 
              v-for="review in reviews" 
              :key="review.id"
              class="bg-surface-container-low p-4 rounded-lg"
            >
              <div class="flex items-center justify-between mb-2">
                <div class="flex items-center gap-2">
                  <Avatar :label="review.name.charAt(0)" shape="circle" class="bg-gradient-to-r from-primary to-primary-container text-white" />
                  <span class="font-semibold text-on_surface">{{ review.name }}</span>
                </div>
                <Rating :modelValue="review.rating" readonly :cancel="false" />
              </div>
              <p class="text-on_surface_variant">{{ review.comment }}</p>
              <p class="text-xs text-outline mt-2">{{ review.date }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
