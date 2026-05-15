export const useApi = () => {
  const config = useRuntimeConfig()
  const apiBase = config.public?.apiBase || 'http://localhost:8080'
  const token = useCookie('auth_token')

  const headers = () => {
    const h = { 'Content-Type': 'application/json' }
    if (token.value) {
      h['Authorization'] = `Bearer ${token.value}`
    }
    return h
  }

  const apiFetch = (url, options = {}) => {
    return $fetch(`${apiBase}${url}`, {
      ...options,
      headers: {
        ...headers(),
        ...options.headers
      }
    })
  }

  const auth = {
    me: () => apiFetch('/auth/me')
  }

  const products = {
    list: () => apiFetch('/products'),
    get: (id) => apiFetch(`/products/${id}`),
    create: (data) => apiFetch('/products', { method: 'POST', body: data }),
    update: (id, data) => apiFetch(`/products/${id}`, { method: 'PUT', body: data }),
    delete: (id) => apiFetch(`/products/${id}`, { method: 'DELETE' })
  }

  const orders = {
    list: () => apiFetch('/orders'),
    get: (id) => apiFetch(`/orders/${id}`),
    updateStatus: (id, status) => apiFetch(`/orders/${id}/status`, { method: 'PUT', body: { status } }),
    cancel: (id) => apiFetch(`/orders/${id}/cancel`, { method: 'PUT' })
  }

  const admin = {
    products,
    orders,
    categories: {
      list: () => apiFetch('/categories')
    }
  }

  const fetchJson = (url, options = {}) => {
    return apiFetch(url, options)
  }

  return { auth, products, orders, admin, fetchJson }
}
