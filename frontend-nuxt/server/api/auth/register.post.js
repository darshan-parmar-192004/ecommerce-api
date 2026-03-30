export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  
  const config = useRuntimeConfig()
  const response = await fetch(`${config.public.apiBase}/auth/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(body)
  })

  const data = await response.json()

  if (!response.ok) {
    const errorMessage = data.error?.message || data.message || data.error || 'Failed to register'
    throw createError({
      statusCode: response.status,
      message: errorMessage
    })
  }

  if (response.ok && data.token) {
    setCookie(event, 'auth_token', data.token, {
      httpOnly: false,
      maxAge: 60 * 60 * 24,
      path: '/'
    })
  }

  return data
})