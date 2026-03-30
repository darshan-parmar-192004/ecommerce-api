export default defineEventHandler(async (event) => {
  const cookie = getCookie(event, 'auth_token')
  
  const headers = {
    'Content-Type': 'application/json'
  }
  
  if (cookie) {
    headers['Authorization'] = `Bearer ${cookie}`
  }
  
  const config = useRuntimeConfig()
  const response = await fetch(`${config.public.apiBase}/auth/me`, {
    headers
  })

  const data = await response.json()

  if (!response.ok) {
    throw createError({
      statusCode: response.status,
      message: data.error || data.message || 'Failed to get user'
    })
  }

  return data
})