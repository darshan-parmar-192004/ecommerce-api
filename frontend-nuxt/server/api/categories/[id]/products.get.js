export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id')
  const query = getQuery(event)
  
  const params = new URLSearchParams()
  Object.entries(query).forEach(([key, value]) => {
    if (value !== undefined && value !== '') {
      params.append(key, String(value))
    }
  })

  const url = `/categories/${id}/products${params.toString() ? `?${params.toString()}` : ''}`
  
  const config = useRuntimeConfig()
  const response = await fetch(`${config.public.apiBase}${url}`, {
    headers: {
      'Content-Type': 'application/json'
    }
  })

  const data = await response.json()

  if (!response.ok) {
    throw createError({
      statusCode: response.status,
      message: data.error || data.message || 'Failed to fetch category products'
    })
  }

  return data
})