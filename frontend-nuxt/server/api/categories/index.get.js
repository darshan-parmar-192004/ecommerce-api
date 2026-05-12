export default defineEventHandler(async (event) => {
  const query = getQuery(event)
  
  const params = new URLSearchParams()
  Object.entries(query).forEach(([key, value]) => {
    if (value !== undefined && value !== '') {
      params.append(key, String(value))
    }
  })

  const url = `/categories${params.toString() ? `?${params.toString()}` : ''}`
  
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
      message: data.error || data.message || 'Failed to fetch categories'
    })
  }

  return data
})