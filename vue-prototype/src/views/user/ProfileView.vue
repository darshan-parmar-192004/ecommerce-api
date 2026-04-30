<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useUserStore } from '@/stores/user'
import { useToastStore } from '@/stores/toast'

const authStore = useAuthStore()
const userStore = useUserStore()
const toastStore = useToastStore()

const form = ref({
  name: '',
  email: '',
  phone: '',
  country: '',
  address: ''
})

const successMessage = ref('')
const isEditing = ref(false)

onMounted(async () => {
  await userStore.fetchProfile()
  if (userStore.profile) {
    form.value = {
      name: userStore.profile.name || '',
      email: userStore.profile.email || '',
      phone: userStore.profile.phone || '',
      country: userStore.profile.country || '',
      address: userStore.profile.address || ''
    }
  } else if (authStore.user) {
    form.value = {
      name: authStore.user.name || '',
      email: authStore.user.email || '',
      phone: '',
      country: '',
      address: ''
    }
  }
})

const updateProfile = async () => {
  try {
    await userStore.updateProfile(form.value)
    successMessage.value = 'Profile updated successfully!'
    toastStore.success('Profile updated successfully!')
    isEditing.value = false
    setTimeout(() => { successMessage.value = '' }, 3000)
  } catch (err) {
    console.error('Update failed', err)
    toastStore.error('Failed to update profile.')
  }
}
</script>

<template>
  <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Profile Settings</h1>

    <div useMotion="{ initial: { opacity: 0, y: 20 }, enter: { opacity: 1, y: 0 } }" class="bg-white rounded-xl border border-gray-200 p-6 space-y-6">
      <div v-if="successMessage" class="p-3 bg-green-50 text-green-700 rounded-lg text-sm">
        {{ successMessage }}
      </div>

      <div class="flex items-center gap-4 mb-6">
        <div class="w-16 h-16 bg-gray-900 text-white rounded-full flex items-center justify-center text-2xl font-bold">
          {{ form.name ? form.name.charAt(0).toUpperCase() : '' }}
        </div>
        <div>
          <h2 class="text-xl font-semibold text-gray-900">{{ form.name }}</h2>
          <p class="text-gray-600">{{ form.email }}</p>
        </div>
      </div>

      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Full Name</label>
          <input v-model="form.name" :disabled="!isEditing" class="input" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Email</label>
          <input v-model="form.email" :disabled="!isEditing" type="email" class="input" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Phone</label>
          <input v-model="form.phone" :disabled="!isEditing" type="tel" class="input" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Country</label>
          <input v-model="form.country" :disabled="!isEditing" type="text" class="input" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Address</label>
          <textarea v-model="form.address" :disabled="!isEditing" rows="3" class="input" />
        </div>
      </div>

      <div class="flex gap-4">
        <button
          v-if="!isEditing"
          @click="isEditing = true"
          class="btn-secondary"
        >
          Edit Profile
        </button>
        <template v-else>
          <button @click="isEditing = false" class="btn-secondary">Cancel</button>
          <button @click="updateProfile" :disabled="userStore.loading" class="btn-primary">
            {{ userStore.loading ? 'Saving...' : 'Save Changes' }}
          </button>
        </template>
      </div>
    </div>
  </div>
</template>
