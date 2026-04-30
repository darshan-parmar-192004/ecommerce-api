import { defineStore } from 'pinia'
import { ref } from 'vue'
import productService from '@/services/productService'
import categoryService from '@/services/categoryService'

export const useProductsStore = defineStore('products', () => {
  const products = ref([])
  const currentProduct = ref(null)
  const categories = ref([])
  const filters = ref({
    category: '',
    minPrice: null,
    maxPrice: null,
    search: '',
    sort: 'newest'
  })
  const pagination = ref({ page: 1, limit: 12, total: 0, total_pages: 1 })
  const loading = ref(false)
  const error = ref(null)

  const fetchProducts = async () => {
    loading.value = true
    error.value = null
    try {
      const params = { ...filters.value, page: pagination.value.page, limit: pagination.value.limit }
      const response = await productService.getProducts(params)
      const responseData = response.data || response
      const rawProducts = responseData.data || []
      
      // Enrich products with category names
      products.value = rawProducts.map(product => {
        const category = categories.value.find(c => c.category_id === product.category_id)
        return {
          ...product,
          category_name: category?.name || 'Uncategorized'
        }
      })
      
      pagination.value.total = responseData.pagination?.total_items || 0
      pagination.value.total_pages = responseData.pagination?.total_pages || 1
    } catch (err) {
      error.value = err.response?.data?.message || err.message || 'Failed to fetch products'
    } finally {
      loading.value = false
    }
  }

  const fetchCategories = async () => {
    try {
      const response = await categoryService.getCategories()
      const responseData = response.data || response
      categories.value = responseData.data || []
    } catch (err) {
      console.error('Failed to fetch categories', err)
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

  const nextPage = () => {
    if (pagination.value.page < pagination.value.total_pages) {
      pagination.value.page += 1
      fetchProducts()
    }
  }

  const prevPage = () => {
    if (pagination.value.page > 1) {
      pagination.value.page -= 1
      fetchProducts()
    }
  }

  const setPage = (page) => {
    if (page >= 1 && page <= pagination.value.total_pages) {
      pagination.value.page = page
      fetchProducts()
    }
  }

  return {
    products,
    currentProduct,
    categories,
    filters,
    pagination,
    loading,
    error,
    fetchProducts,
    fetchProductById,
    setFilters,
    resetFilters,
    nextPage,
    prevPage,
    setPage,
    fetchCategories
  }
})