import { useToast as usePrimeToast } from 'primevue/usetoast'

export const useAppToast = () => {
  const toast = usePrimeToast()

  const success = (message) => {
    toast.add({ severity: 'success', summary: 'Success', detail: message, life: 3000 })
  }

  const error = (message) => {
    toast.add({ severity: 'error', summary: 'Error', detail: message, life: 5000 })
  }

  const info = (message) => {
    toast.add({ severity: 'info', summary: 'Info', detail: message, life: 3000 })
  }

  const warning = (message) => {
    toast.add({ severity: 'warn', summary: 'Warning', detail: message, life: 4000 })
  }

  return {
    success,
    error,
    info,
    warning
  }
}
