import { defineStore } from 'pinia'
import { ref } from 'vue'
import productService from '@/services/productService'

export const useProductsStore = defineStore('products', () => {
  const products = ref([])
  const currentProduct = ref(null)
  const filters = ref({
    category: '',
    minPrice: null,
    maxPrice: null,
    search: '',
    sort: 'newest'
  })
  const pagination = ref({ page: 1, limit: 12, total: 0 })
  const loading = ref(false)
  const error = ref(null)

  const fetchProducts = async () => {
    loading.value = true
    error.value = null
    try {
      const params = { ...filters.value, ...pagination.value }
      const { data } = await productService.getProducts(params)
      products.value = data.products
      pagination.value.total = data.total
    } catch (err) {
      error.value = err.response?.data?.message || 'Failed to fetch products'
    } finally {
      loading.value = false
    }
  }

  const fetchProductById = async (id) => {
    loading.value = true
    error.value = null
    try {
      const { data } = await productService.getProductById(id)
      currentProduct.value = data
    } catch (err) {
      error.value = err.response?.data?.message || 'Failed to fetch product'
    } finally {
      loading.value = false
    }
  }

  const setFilters = (newFilters) => {
    filters.value = { ...filters.value, ...newFilters }
    pagination.value.page = 1
    fetchProducts()
  }

  const resetFilters = () => {
    filters.value = { category: '', minPrice: null, maxPrice: null, search: '', sort: 'newest' }
    pagination.value.page = 1
    fetchProducts()
  }

  return {
    products,
    currentProduct,
    filters,
    pagination,
    loading,
    error,
    fetchProducts,
    fetchProductById,
    setFilters,
    resetFilters
  }
})
