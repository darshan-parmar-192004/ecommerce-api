import api from './api'

const productService = {
  getProducts: (params) => {
    // Transform camelCase to snake_case for backend API
    const transformedParams = { ...params }
    if (transformedParams.minPrice !== undefined && transformedParams.minPrice !== null) {
      transformedParams.min_price = transformedParams.minPrice
      delete transformedParams.minPrice
    }
    if (transformedParams.maxPrice !== undefined && transformedParams.maxPrice !== null) {
      transformedParams.max_price = transformedParams.maxPrice
      delete transformedParams.maxPrice
    }
    return api.get('/products', { params: transformedParams })
  },
  getProductById: (id) => api.get(`/products/${id}`),
  createProduct: (data) => api.post('/products', data),
  updateProduct: (id, data) => api.put(`/products/${id}`, data),
  deleteProduct: (id) => api.delete(`/products/${id}`),
  getInventory: (params) => api.get('/inventory', { params }),
  updateInventory: (id, data) => api.put(`/inventory/${id}`, data),
  getTopSellers: () => api.get('/inventory/top-sellers'),
  getCustomerLifetimeValue: () => api.get('/inventory/customer-lifetime-value'),
  getStockLevels: () => api.get('/inventory/stock')
}

export default productService
