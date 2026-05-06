export default defineEventHandler(async (event) => {
  const token = getCookie(event, 'auth_token')
  
  if (!token) {
    throw createError({
      statusCode: 401,
      statusMessage: 'Unauthorized'
    })
  }
  
  const config = useRuntimeConfig()
  
  // Verify user is admin
  const verifyRes = await fetch(`${config.public.apiBase}/auth/me`, {
    headers: { 'Authorization': `Bearer ${token}` }
  })
  
  if (!verifyRes.ok) {
    throw createError({
      statusCode: 401,
      statusMessage: 'Unauthorized'
    })
  }
  
  const userData = await verifyRes.json()
  const role = userData.data?.role || userData.role
  
  if (role !== 'admin') {
    throw createError({
      statusCode: 403,
      statusMessage: 'Forbidden - Admin access required'
    })
  }
  
  // Fetch orders from backend
  const res = await fetch(`${config.public.apiBase}/orders`)
  const data = await res.json()
  
  return {
    data: data.data || data
  }
})