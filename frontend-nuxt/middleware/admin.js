export default defineNuxtRouteMiddleware(async (to) => {
  let token = null
  if (import.meta.client) {
    token = localStorage.getItem('auth_token')
  }

  if (!token) {
    return navigateTo(`/auth/login?redirect=${encodeURIComponent(to.fullPath)}`)
  }

  try {
    const { auth } = useApi()
    const data = await auth.me()
    const role = data?.data?.role || data?.role
    if (role !== 'admin') {
      throw createError({
        statusCode: 403,
        statusMessage: 'Forbidden - Admin access required'
      })
    }
  } catch {
    return navigateTo(`/auth/login?redirect=${encodeURIComponent(to.fullPath)}`)
  }
})
