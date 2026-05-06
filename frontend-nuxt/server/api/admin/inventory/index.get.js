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
  
  // Fetch products with inventory info from backend
  const res = await fetch(`${config.public.apiBase}/products`)
  const data = await res.json()
  
  // Format for inventory view
  const products = (data.data || data).map(p => ({
    id: p.id || p.product_id,
    name: p.name,
    sku: p.sku || p.sku,
    price: p.price,
    stock_quantity: p.stock_quantity || p.inventory_count || 0,
    low_stock_threshold: p.low_stock_threshold || 10,
    status: (p.stock_quantity || 0) > 0 ? 'in_stock' : 'out_of_stock'
  }))
  
  return {
    data: products
  }
})