export default defineEventHandler(async (event) => {
  const token = getCookie(event, 'auth_token')
  const id = getRouterParam(event, 'id')
  
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
  
  // Fetch product from backend
  const res = await fetch(`${config.public.apiBase}/products/${id}`)
  
  if (!res.ok) {
    throw createError({
      statusCode: res.status,
      statusMessage: 'Product not found'
    })
  }
  
  const data = await res.json()
  
  return {
    data: data.data || data
  }
})