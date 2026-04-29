import api from './api'

const productService = {
  getProducts: (params) => api.get('/products', { params }),
  getProductById: (id) => api.get(`/products/${id}`),
  createProduct: (data) => api.post('/products', data),
  updateProduct: (id, data) => api.put(`/products/${id}`, data),
  deleteProduct: (id) => api.delete(`/products/${id}`),
  getInventory: (params) => api.get('/inventory', { params }),
  updateInventory: (id, data) => api.put(`/inventory/${id}`, data)
}

export default productService
