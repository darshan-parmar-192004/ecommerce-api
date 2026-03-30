export default defineEventHandler(async (event) => {
  const cookie = getCookie(event, 'auth_token')
  const body = await readBody(event)
  
  if (!cookie) {
    throw createError({
      statusCode: 401,
      message: 'Authentication required'
    })
  }

  if (!body.items || body.items.length === 0) {
    throw createError({
      statusCode: 400,
      message: 'Order must contain at least one item'
    })
  }

  const config = useRuntimeConfig()
  const response = await fetch(`${config.public.apiBase}/orders`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${cookie}`,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(body)
  })

  const data = await response.json()

  if (!response.ok) {
    throw createError({
      statusCode: response.status,
      message: data.error || data.message || 'Failed to create order'
    })
  }

  return data
})
