import api from './api'

const userService = {
  getOrders: (customerId, params) => api.get(`/customers/${customerId}/orders`, { params }),
  getOrderById: (id) => api.get(`/orders/${id}`),
  getOrderDetail: (id) => api.get(`/orders/${id}/detail`),
  createOrder: (data) => api.post('/orders', data),
  cancelOrder: (id) => api.put(`/orders/${id}/cancel`),
  getProfile: () => api.get('/customers/me'),
  updateProfile: (data) => api.put('/customers/me', data)
}

export default userService
