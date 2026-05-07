import axios from 'axios'
import { getToken, clearAuth } from '@/lib/auth'
import router from '@/router'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: 10000,
  withCredentials: true
})

api.interceptors.request.use(
  config => {
    const token = getToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => Promise.reject(error)
)

api.interceptors.response.use(
  response => response,
  async error => {
    if (error.response?.status === 401) {
      clearAuth()
      router.push({ name: 'Login' })
    }
    return Promise.reject(error)
  }
)

export default api
