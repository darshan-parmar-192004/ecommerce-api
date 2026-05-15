import { errorMessages } from '../constants/errorMessages'

export const useErrorHandler = () => {
  const toast = useAppToast()

  const parseApiError = (error) => {
    if (!error) return errorMessages.unknown

    if (error.response?.data?.message) {
      return error.response.data.message
    }

    if (error.response?.status === 400) {
      return errorMessages.badRequest
    }
    if (error.response?.status === 401) {
      return errorMessages.unauthorized
    }
    if (error.response?.status === 403) {
      return errorMessages.forbidden
    }
    if (error.response?.status === 404) {
      return errorMessages.notFound
    }
    if (error.response?.status === 409) {
      return errorMessages.conflict
    }
    if (error.response?.status === 422) {
      return errorMessages.validationError
    }
    if (error.response?.status === 429) {
      return errorMessages.tooManyRequests
    }
    if (error.response?.status >= 500) {
      return errorMessages.serverError
    }

    if (error.message) {
      return error.message
    }

    return errorMessages.genericError
  }

  const showError = (error, fallback = errorMessages.defaultFallback) => {
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
