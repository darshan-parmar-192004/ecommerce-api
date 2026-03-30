export default defineEventHandler(async (event) => {
  const cookie = getCookie(event, 'auth_token')
  
  const headers = {
    'Content-Type': 'application/json'
  }
  
  if (cookie) {
    headers['Authorization'] = `Bearer ${cookie}`
  }
  
  const config = useRuntimeConfig()
  const response = await fetch(`${config.public.apiBase}/auth/logout`, {
    method: 'POST',
    headers
  })

  deleteCookie(event, 'auth_token')

  const data = await response.json()

  if (!response.ok) {
    const errorMessage = data.error?.message || data.message || data.error || 'Failed to logout'
    throw createError({
      statusCode: response.status,
      message: errorMessage
    })
  }

  return data
})