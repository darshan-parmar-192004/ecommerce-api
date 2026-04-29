import { defineStore } from 'pinia'
import { ref } from 'vue'
import userService from '@/services/userService'

export const useUserStore = defineStore('user', () => {
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
      const { data } = await userService.getOrders({ ...pagination.value, ...params })
      orders.value = data.orders
      pagination.value.total = data.total
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
      const { data } = await userService.getOrderById(id)
      currentOrder.value = data
    } catch (err) {
      error.value = err.response?.data?.message || 'Failed to fetch order'
    } finally {
      loading.value = false
    }
  }

  const createOrder = async (orderData) => {
    loading.value = true
    error.value = null
    try {
      const { data } = await userService.createOrder(orderData)
      return data
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
      const { data } = await userService.cancelOrder(id)
      return data
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
      const { data } = await userService.getProfile()
      profile.value = data
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
      const { data: updated } = await userService.updateProfile(data)
      profile.value = updated
      return updated
    } catch (err) {
      error.value = err.response?.data?.message || 'Failed to update profile'
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    orders,
    currentOrder,
    profile,
    loading,
    error,
    pagination,
    fetchOrders,
    fetchOrderById,
    createOrder,
    cancelOrder,
    fetchProfile,
    updateProfile
  }
})
