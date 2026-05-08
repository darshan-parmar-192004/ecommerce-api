import api from './api'

const categoryService = {
  getCategories: () => api.get('/categories'),
  getCategoryHierarchy: () => api.get('/categories/hierarchy'),
  createCategory: (data) => api.post('/categories', data),
  updateCategory: (id, data) => api.put(`/categories/${id}`, data),
  deleteCategory: (id) => api.delete(`/categories/${id}`)
}

export default categoryService
