import { STORAGE_KEYS } from '@/constants'

/**
 * Non-composable auth utility for use outside Vue component context
 * (e.g., router guards, axios interceptors).
 *
 * Directly reads/writes localStorage since composables require setup context.
 */

export function getToken() {
  return localStorage.getItem(STORAGE_KEYS.TOKEN) || null
}

export function getUser() {
  try {
    const userStr = localStorage.getItem(STORAGE_KEYS.USER)
    return userStr ? JSON.parse(userStr) : null
  } catch {
    return null
  }
}

export function isAuthenticated() {
  return !!getToken()
}

export function isAdmin() {
  const user = getUser()
  return user?.role === 'admin'
}

export function clearAuth() {
  localStorage.removeItem(STORAGE_KEYS.TOKEN)
  localStorage.removeItem(STORAGE_KEYS.USER)
}

export function setToken(token) {
  localStorage.setItem(STORAGE_KEYS.TOKEN, token)
}
