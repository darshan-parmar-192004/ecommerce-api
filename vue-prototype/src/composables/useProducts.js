import { ref, reactive } from 'vue'
import productService from '@/lib/productService'
import categoryService from '@/lib/categoryService'

const mapProduct = (product) => ({
  productId: product.product_id,
  categoryId: product.category_id,
  name: product.name,
  description: product.description,
  price: product.price,
  stockQuantity: product.stock_quantity,
  images: product.images,
  rating: product.rating,
  categoryName: product.category_name || 'Uncategorized'
})

const mapCategory = (category) => ({
  categoryId: category.category_id,
  name: category.name,
  description: category.description,
  productCount: category.product_count
})

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
const pagination = ref({ page: 1, limit: 12, total: 0, totalPages: 1 })
const loading = ref(false)
const error = ref(null)

const fetchProducts = async () => {
  loading.value = true
  error.value = null
  try {
    const params = { ...filters.value, page: pagination.value.page, limit: pagination.value.limit }
    console.log('Fetching products with', params);
    const response = await productService.getProducts(params);
    console.log('Response', response);
    const responseData = response.data || response
    const rawProducts = (responseData.data || []).map(mapProduct)

    // Enrich products with category names from local categories
    products.value = rawProducts.map(product => {
      const category = categories.value ? categories.value.find(c => c.categoryId === product.categoryId) : null
      return {
        ...product,
        categoryName: category?.name || product.categoryName || 'Uncategorized'
      }
    })

    pagination.value.total = responseData.pagination?.total_items || 0
    pagination.value.totalPages = responseData.pagination?.total_pages || 1
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
    categories.value = (responseData.data || []).map(mapCategory)
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
    const rawProduct = responseData.data || responseData
    currentProduct.value = mapProduct(rawProduct)
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
  if (pagination.value.page < pagination.value.totalPages) {
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
  if (page >= 1 && page <= pagination.value.totalPages) {
    pagination.value.page = page
    fetchProducts()
  }
}

export function useProducts() {
  return reactive({
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
  })
}
