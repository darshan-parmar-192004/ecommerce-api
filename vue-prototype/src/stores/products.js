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
      const response = await productService.getProducts(params)
      const responseData = response.data || response
      products.value = responseData.data || []
      pagination.value.total = responseData.pagination?.total_items || 0
      pagination.value.total_pages = responseData.pagination?.total_pages || 1
    } catch (err) {
      error.value = err.response?.data?.message || err.message || 'Failed to fetch products'
    } finally {
      loading.value = false
    }
  }

  const fetchProductById = async (id) => {
    loading.value = true
    error.value = null
    try {
      const response = await productService.getProductById(id)
      const responseData = response.data || response
      currentProduct.value = responseData.data || responseData
    } catch (err) {
      error.value = err.response?.data?.message || err.message || 'Failed to fetch product'
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
