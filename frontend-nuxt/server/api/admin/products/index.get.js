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
  console.log('Admin API role check:', role)
  const allowed = ['admin','Administrator','ADMIN']
  if (!allowed.includes(role)) {
    throw createError({
      statusCode: 403,
      statusMessage: 'Forbidden - Admin access required'
    })
  }
  
  // Fetch products from backend
  const res = await fetch(`${config.public.apiBase}/products`)
  const data = await res.json()
  
  return {
    data: data.data || data,
    pagination: data.pagination
  }
})