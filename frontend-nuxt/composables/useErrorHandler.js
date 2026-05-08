import { toast } from 'vue-sonner'

export const useErrorHandler = () => {
  const parseApiError = (error) => {
    if (!error) return 'An unknown error occurred'

    if (error.response?.data?.message) {
      return error.response.data.message
    }

    if (error.response?.status === 400) {
      return 'Bad request. Please check your input.'
    }
    if (error.response?.status === 401) {
      return 'Unauthorized. Please log in again.'
    }
    if (error.response?.status === 403) {
      return 'Access denied. You do not have permission.'
    }
    if (error.response?.status === 404) {
      return 'Resource not found.'
    }
    if (error.response?.status === 409) {
      return 'Conflict. The resource already exists.'
    }
    if (error.response?.status === 422) {
      return 'Validation error. Please check your input.'
    }
    if (error.response?.status === 429) {
      return 'Too many requests. Please try again later.'
    }
    if (error.response?.status >= 500) {
      return 'Server error. Please try again later.'
    }

    if (error.message) {
      return error.message
    }

    return 'An error occurred. Please try again.'
  }

  const showError = (error, fallback = 'An error occurred') => {
    const message = parseApiError(error) || fallback
    toast.error(message)
  }

  const showSuccess = (message) => {
    toast.success(message)
  }

  const showInfo = (message) => {
    toast.info(message)
  }

  const showWarning = (message) => {
    toast.warning(message)
  }

  return {
    parseApiError,
    showError,
    showSuccess,
    showInfo,
    showWarning
  }
}