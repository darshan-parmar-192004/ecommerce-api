/**
 * Application constants
 * Centralized location for all hardcoded values
 * Change once here, applies everywhere
 */

// Storage Keys
export const STORAGE_KEYS = {
  TOKEN: 'token',
  USER: 'user',
  CART: 'cart'
}

// Pagination
export const PAGINATION = {
  DEFAULT_LIMIT: 12,
  PRODUCT_MANAGE_LIMIT: 50,
  INVENTORY_LIMIT: 200,
  VISIBLE_PAGES: 7
}

// Toast
export const TOAST = {
  DEFAULT_DURATION: 4000,
  MAX_TOASTS: 5
}

// Fallback Values
export const FALLBACKS = {
  CUSTOMER_ID: 'CUST-001',
  IMAGE_PLACEHOLDER: '/placeholder.jpg',
  IMAGE_PLACEHOLDER_SVG: '/placeholder.svg'
}

// API
export const API = {
  BASE_URL: import.meta.env.VITE_API_BASE_URL || '/api',
  TIMEOUT: 10000
}

// Animation
export const ANIMATION = {
  STAGGER_DELAY: 100, // ms per item
  DRAWER_DURATION: 300,
  CART_ITEM_DURATION: 300
}

// Route Names
export const ROUTES = {
  HOME: 'Home',
  LOGIN: 'Login',
  REGISTER: 'Register',
  UNAUTHORIZED: 'Unauthorized',
  NOT_FOUND: 'NotFound',
  ADMIN_DASHBOARD: 'AdminDashboard'
}

// User Roles
export const ROLES = {
  CUSTOMER: 'customer',
  ADMIN: 'admin'
}

// Password Requirements
export const PASSWORD_POLICY = {
  MIN_LENGTH: 8,
  REQUIRE_UPPERCASE: true,
  REQUIRE_LOWERCASE: true,
  REQUIRE_DIGIT: true,
  REQUIRE_SPECIAL: true
}
