export default defineEventHandler(async (event) => {
  const query = getQuery(event)
  const cookie = getCookie(event, 'auth_token')
  
  const params = new URLSearchParams()
  Object.entries(query).forEach(([key, value]) => {
    if (value !== undefined && value !== '') {
      params.append(key, String(value))
    }
  })

  const url = `/products${params.toString() ? `?${params.toString()}` : ''}`
  
  const headers = {
    'Content-Type': 'application/json'
  }
  
  if (cookie) {
    headers['Authorization'] = `Bearer ${cookie}`
  }

  const config = useRuntimeConfig()
  const response = await fetch(`${config.public.apiBase}${url}`, {
    headers
  })

  const data = await response.json()

  if (!response.ok) {
    throw createError({
      statusCode: response.status,
      message: data.error || data.message || 'Failed to fetch products'
    })
  }

  return data
})
