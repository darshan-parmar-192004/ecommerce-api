export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id')
  const cookie = getCookie(event, 'auth_token')
  
  if (!cookie) {
    throw createError({
      statusCode: 401,
      message: 'Authentication required'
    })
  }

  const config = useRuntimeConfig()
  const response = await fetch(`${config.public.apiBase}/orders/${id}`, {
    headers: {
      'Authorization': `Bearer ${cookie}`,
      'Content-Type': 'application/json'
    }
  })

  const data = await response.json()

  if (!response.ok) {
    throw createError({
      statusCode: response.status,
      message: data.error || data.message || 'Failed to fetch order'
    })
  }

  return data
})
