import { defineStore } from 'pinia'
import { PAGINATION } from '~/utils/constants'

export const useProductsStore = defineStore('products', {
  state: () => ({
    products: [],
    categories: [],
    currentProduct: null,
    loading: false,
    categoriesLoading: false,
    pagination: {
      page: PAGINATION.DEFAULT_PAGE,
      limit: PAGINATION.DEFAULT_LIMIT,
      total: 0,
      pages: 0
    },
    filters: {
      categoryId: null,
      search: '',
      sortBy: 'created_at',
      sortOrder: 'desc'
    }
  }),

  getters: {
    productList: (state) => state.products,
    categoryList: (state) => state.categories,
    currentPage: (state) => state.pagination.page,
    totalPages: (state) => state.pagination.pages,
    hasNextPage: (state) => state.pagination.page < state.pagination.pages,
    hasPrevPage: (state) => state.pagination.page > 1
  },

  actions: {
    async fetchProducts(params = {}) {
      this.loading = true
      try {
        const { products } = useApi()
        const response = await products.list({
          page: params.page || this.pagination.page,
          limit: params.limit || this.pagination.limit,
          category_id: params.categoryId || this.filters.categoryId,
          search: params.search || this.filters.search,
          sort_by: params.sortBy || this.filters.sortBy,
          sort_order: params.sortOrder || this.filters.sortOrder
        })

        this.products = response.data || response || []
        if (response.pagination) {
          this.pagination = { ...this.pagination, ...response.pagination }
        }
        return this.products
      } catch (error) {
        console.error('Failed to fetch products:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchProduct(id) {
      this.loading = true
      try {
        const { products } = useApi()
        const response = await products.get(id)
        this.currentProduct = response.data || response
        return this.currentProduct
      } catch (error) {
        console.error('Failed to fetch product:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchCategories() {
      this.categoriesLoading = true
      try {
        const { categories } = useApi()
        const response = await categories.list()
        this.categories = response.data || response || []
        return this.categories
      } catch (error) {
        console.error('Failed to fetch categories:', error)
        throw error
      } finally {
        this.categoriesLoading = false
      }
    },

    setFilters(filters) {
      this.filters = { ...this.filters, ...filters }
      this.pagination.page = PAGINATION.DEFAULT_PAGE
    },

    setPage(page) {
      this.pagination.page = page
    },

    clearCurrentProduct() {
      this.currentProduct = null
    }
  }
})