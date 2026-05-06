/**
 * Centralized error handling utility
 * Provides consistent error parsing and user-friendly messages
 */

/**
 * Parse API error response and return user-friendly message
 * @param {Error} error - Axios error object
 * @returns {string} User-friendly error message
 */
export function parseApiError(error) {
  // Network error (no response)
  if (!error.response) {
    return 'Network error. Please check your internet connection.'
  }

  const { status, data } = error.response

  // Handle specific status codes
  switch (status) {
    case 400:
      return data?.message || 'Invalid request. Please check your input.'
    case 401:
      return 'Your session has expired. Please login again.'
    case 403:
      return 'You do not have permission to perform this action.'
    case 404:
      return 'The requested resource was not found.'
    case 409:
      return data?.message || 'This resource already exists.'
    case 422:
      // Validation errors
      if (data?.errors) {
        const firstError = Array.isArray(data.errors) ? data.errors[0] : data.errors
        return firstError?.message || firstError || 'Validation failed.'
      }
      return data?.message || 'Please check your input.'
    case 429:
      return 'Too many requests. Please try again later.'
    case 500:
      return 'Server error. Please try again later.'
    case 502:
    case 503:
    case 504:
      return 'Service temporarily unavailable. Please try again later.'
    default:
      return data?.message || 'An unexpected error occurred.'
  }
}

/**
 * Handle error with toast notification
 * @param {Error} error - Error object
 * @param {Object} toastStore - Pinia toast store
 * @param {string} fallbackMessage - Fallback message if parsing fails
 */
export function handleErrorWithToast(error, toastStore, fallbackMessage = 'An error occurred') {
  const message = parseApiError(error)
  toastStore.error(message)

  // Log detailed error for debugging (only in development)
  if (import.meta.env.DEV) {
    console.error('API Error:', {
      message: error.message,
      status: error.response?.status,
      data: error.response?.data,
      config: error.config
    })
  }
}

/**
 * Create error object with type information
 * @param {string} type - Error type (network, validation, auth, etc.)
 * @param {string} message - Error message
 * @param {*} details - Additional error details
 */
export function createAppError(type, message, details = null) {
  const error = new Error(message)
  error.type = type
  error.details = details
  return error
}

export default {
  parseApiError,
  handleErrorWithToast,
  createAppError
}
