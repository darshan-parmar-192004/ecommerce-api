<script setup>
import { useUserStore } from '~/stores/user'
import { useAuthStore } from '~/stores/auth'

const userStore = useUserStore()
const authStore = useAuthStore()

const form = reactive({
  name: '',
  email: '',
  phone: ''
})

const loading = ref(false)
const success = ref(false)

onMounted(async () => {
  await userStore.fetchProfile()
  if (userStore.profile) {
    form.name = userStore.profile.name || ''
    form.email = userStore.profile.email || ''
    form.phone = userStore.profile.phone || ''
  }
})

const updateProfile = async () => {
  loading.value = true
  success.value = false
  try {
    await userStore.updateProfile(form)
    success.value = true
    setTimeout(() => success.value = false, 3000)
  } catch (error) {
    console.error('Failed to update profile:', error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-2xl mx-auto">
    <h1 class="text-3xl font-bold text-on_surface mb-8">Profile Settings</h1>
    
    <div class="bg-surface-container-lowest rounded-lg p-6">
      <form @submit.prevent="updateProfile" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-on_surface mb-2">Name</label>
          <input
            v-model="form.name"
            type="text"
            class="w-full px-4 py-2 rounded-md border border-outline-variant focus:border-primary focus:outline-none bg-surface-container-lowest text-on_surface"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-on_surface mb-2">Email</label>
          <input
            v-model="form.email"
            type="email"
            disabled
            class="w-full px-4 py-2 rounded-md border border-outline-variant bg-surface-container text-on_surface_variant cursor-not-allowed"
          />
        </div>
        
        <div>
          <label class="block text-sm font-medium text-on_surface mb-2">Phone</label>
          <input
            v-model="form.phone"
            type="tel"
            class="w-full px-4 py-2 rounded-md border border-outline-variant focus:border-primary focus:outline-none bg-surface-container-lowest text-on_surface"
          />
        </div>
        
        <div v-if="success" class="p-4 rounded-md bg-success-container text-success">
          Profile updated successfully!
        </div>
        
        <button
          type="submit"
          :disabled="loading"
          class="px-6 py-2 rounded-md font-medium text-white bg-gradient-to-r from-primary to-primary-container shadow-glow hover:shadow-xl transition-all disabled:opacity-50"
        >
          {{ loading ? 'Saving...' : 'Save Changes' }}
        </button>
      </form>
    </div>
  </div>
</template>