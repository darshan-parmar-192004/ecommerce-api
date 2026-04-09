import { useAuthStore } from '~/stores/auth'

export const useApi = () => {
  const config = useRuntimeConfig()
  const apiBase = config.public.apiBase
  const authStore = useAuthStore()

  const getToken = () => {
    const tokenCookie = useCookie('auth_token')
    return tokenCookie.value || authStore.token
  }

  const fetchJson = async (url, options = {}) => {
    const headers = {
      'Content-Type': 'application/json',
      ...(options.headers || {})
    }

    const token = getToken()
    if (token) {
      headers['Authorization'] = `Bearer ${token}`
    }

    try {
      const response = await fetch(`${apiBase}${url}`, {
        ...options,
        headers
      })

      const data = await response.json().catch(() => ({}))

      if (!response.ok) {
        let message = 'Request failed'
        
        if (data.error && typeof data.error === 'object') {
          message = data.error.message || message
        } else if (data.message) {
          message = data.message
        } else if (data.error) {
          message = String(data.error)
        }

        if (response.status === 401) {
          authStore.clearAuth()
        }

        throw { message, status: response.status }
      }

      return data
    } catch (err) {
      if (err.message && err.status !== undefined) {
        throw err
      }
      throw { message: 'Network error. Please check your connection.', status: 0 }
    }
  }

  const products = {
    list: (params = {}) => {
      const query = new URLSearchParams(
        Object.entries(params).filter(([_, v]) => v !== undefined && v !== '')
      ).toString()
      return fetchJson(`/products${query ? `?${query}` : ''}`)
    },
    get: (id) => fetchJson(`/products/${id}`)
  }

  const categories = {
    list: () => fetchJson('/categories'),
    get: (id) => fetchJson(`/categories/${id}`),
    products: (id) => fetchJson(`/categories/${id}/products`)
  }

  const auth = {
    register: async (data) => {
      const response = await fetch(`${apiBase}/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
      })
      const result = await response.json()
      
      if (!response.ok) {
        throw { message: result.message || result.error || 'Registration failed', status: response.status }
      }
      
      if (result.token) {
        const tokenCookie = useCookie('auth_token', { maxAge: 60 * 60 * 24 })
        tokenCookie.value = result.token
        await authStore.setAuth({
          token: result.token,
          customer: result.data?.customer || result.customer
        })
      }
      return result
    },
    login: async (data) => {
      const response = await fetch(`${apiBase}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
      })
      const result = await response.json()
      
      if (!response.ok) {
        throw { message: result.message || result.error || 'Login failed', status: response.status }
      }
      
      if (result.token) {
        const tokenCookie = useCookie('auth_token', { maxAge: 60 * 60 * 24 })
        tokenCookie.value = result.token
        await authStore.setAuth({
          token: result.token,
          customer: result.customer
        })
      }
      return result
    },
    logout: async () => {
      const token = getToken()
      try {
        if (token) {
          await fetch(`${apiBase}/auth/logout`, { 
            method: 'POST',
            headers: { 'Authorization': `Bearer ${token}` }
          })
        }
      } finally {
        const tokenCookie = useCookie('auth_token')
        tokenCookie.value = null
        authStore.clearAuth()
      }
    },
    me: () => fetchJson('/auth/me')
  }

  const orders = {
    list: () => fetchJson('/orders'),
    get: (id) => fetchJson(`/orders/${id}`),
    create: (data) => fetchJson('/orders', { method: 'POST', body: JSON.stringify(data) })
  }

  return { fetchJson, products, categories, auth, orders }
}
