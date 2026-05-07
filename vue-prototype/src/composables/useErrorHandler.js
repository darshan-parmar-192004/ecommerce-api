import { useToast } from 'vue-toastification'

export function useErrorHandler() {
  const toast = useToast()

  function parseApiError(error) {
    if (!error.response) return 'Network error. Please check your internet connection.'
    const { status, data } = error.response
    switch (status) {
      case 400: return data?.message || 'Invalid request. Please check your input.'
      case 401: return 'Your session has expired. Please login again.'
      case 403: return 'You do not have permission to perform this action.'
      case 404: return 'The requested resource was not found.'
      case 409: return data?.message || 'This resource already exists.'
      case 422:
        if (data?.errors) {
          const firstError = Array.isArray(data.errors) ? data.errors[0] : data.errors
          return firstError?.message || firstError || 'Validation failed.'
        }
        return data?.message || 'Please check your input.'
      case 429: return 'Too many requests. Please try again later.'
      case 500: return 'Server error. Please try again later.'
      case 502: case 503: case 504: return 'Service temporarily unavailable. Please try again later.'
      default: return data?.message || 'An unexpected error occurred.'
    }
  }

  function showError(error, fallbackMessage = 'An error occurred') {
    const message = parseApiError(error)
    toast.error(message)
    if (import.meta.env.DEV) {
      console.error('API Error:', { message: error.message, status: error.response?.status, data: error.response?.data, config: error.config })
    }
  }

  function showSuccess(message) { toast.success(message) }
  function showInfo(message) { toast.info(message) }
  function showWarning(message) { toast.warning(message) }

  return { parseApiError, showError, showSuccess, showInfo, showWarning }
}
