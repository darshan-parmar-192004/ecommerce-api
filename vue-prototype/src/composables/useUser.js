import { ref } from 'vue'
import userService from '@/lib/userService'
import { useAuth } from '@/composables/useAuth'

// Module-level singleton state (shared across all calls to useUser)
const orders = ref([])
const currentOrder = ref(null)
const profile = ref(null)
const loading = ref(false)
const error = ref(null)
const pagination = ref({ page: 1, limit: 10, total: 0 })

const fetchOrders = async (params = {}) => {
  loading.value = true
  error.value = null
  try {
    const { user } = useAuth()
    const customerId = user.value?.customerId
    if (!customerId) {
      throw new Error('User not authenticated')
    }
    const response = await userService.getOrders(customerId, { ...pagination.value, ...params })
    const responseData = response.data || response
    orders.value = responseData.data || []
    pagination.value.total = responseData.pagination?.total_items || 0
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to fetch orders'
  } finally {
    loading.value = false
  }
}

const fetchOrderById = async (id) => {
  loading.value = true
  error.value = null
  try {
    const response = await userService.getOrderById(id)
    const responseData = response.data || response
    currentOrder.value = responseData.data || responseData
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to fetch order'
  } finally {
    loading.value = false
  }
}

const fetchOrderDetail = async (id) => {
  loading.value = true
  error.value = null
  try {
    const response = await userService.getOrderDetail(id)
    const responseData = response.data || response
    const detailData = responseData.data || responseData
    currentOrder.value = {
      ...detailData.order,
      items: detailData.items
    }
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to fetch order detail'
  } finally {
    loading.value = false
  }
}

const createOrder = async (orderData) => {
  loading.value = true
  error.value = null
  try {
    const response = await userService.createOrder(orderData)
    const responseData = response.data || response
    return responseData.data || responseData
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to create order'
    throw err
  } finally {
    loading.value = false
  }
}

const cancelOrder = async (id) => {
  loading.value = true
  error.value = null
  try {
    const response = await userService.cancelOrder(id)
    const responseData = response.data || response
    return responseData.data || responseData
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to cancel order'
    throw err
  } finally {
    loading.value = false
  }
}

const fetchProfile = async () => {
  loading.value = true
  try {
    const response = await userService.getProfile()
    const responseData = response.data || response
    profile.value = responseData.data || responseData
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to fetch profile'
  } finally {
    loading.value = false
  }
}

const updateProfile = async (data) => {
  loading.value = true
  error.value = null
  try {
    const response = await userService.updateProfile({
      name: data.name,
      country: data.country,
      phone: data.phone
    })
    const responseData = response.data || response
    profile.value = responseData.data || responseData
    return profile.value
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to update profile'
    throw err
  } finally {
    loading.value = false
  }
}

/**
 * User composable - provides user/order state and methods.
 * Uses module-level singleton state so all callers share the same reactive state.
 */
export function useUser() {
  return {
    orders,
    currentOrder,
    profile,
    loading,
    error,
    pagination,
    fetchOrders,
    fetchOrderById,
    fetchOrderDetail,
    createOrder,
    cancelOrder,
    fetchProfile,
    updateProfile
  }
}
