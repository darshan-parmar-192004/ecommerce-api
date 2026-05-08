<script setup>
import { ref, onMounted } from 'vue'
import { useAuth } from '@/composables/useAuth'
import { useUser } from '@/composables/useUser'
import { useErrorHandler } from '@/composables/useErrorHandler'

const { user: authUser } = useAuth()
const { profile, loading: userLoading, fetchProfile, updateProfile: saveProfile } = useUser()
const { showError, showSuccess } = useErrorHandler()

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
  await fetchProfile()
  if (profile.value) {
    form.value = {
      name: profile.value.name || '',
      email: profile.value.email || '',
      phone: profile.value.phone || '',
      country: profile.value.country || '',
      address: profile.value.address || ''
    }
  } else if (authUser.value) {
    form.value = {
      name: authUser.value.name || '',
      email: authUser.value.email || '',
      phone: '',
      country: '',
      address: ''
    }
  }
})

const updateProfile = async () => {
  try {
    await saveProfile(form.value)
    successMessage.value = 'Profile updated successfully!'
    showSuccess('Profile updated successfully!')
    isEditing.value = false
    setTimeout(() => { successMessage.value = '' }, 3000)
  } catch (err) {
    showError(err, 'Failed to update profile.')
  }
}
</script>

<template>
  <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-8">Profile Settings</h1>

    <div useMotion="{ initial: { opacity: 0, y: 20 }, enter: { opacity: 1, y: 0 } }" class="bg-white dark:bg-brand-800 rounded-xl border border-gray-200 dark:border-brand-700 p-6 space-y-6">
      <div v-if="successMessage" class="p-3 bg-green-50 dark:bg-green-900/20 text-green-700 dark:text-green-400 rounded-lg text-sm">
        {{ successMessage }}
      </div>

      <div class="flex items-center gap-4 mb-6">
        <div class="w-16 h-16 bg-gray-900 dark:bg-brand-700 text-white rounded-full flex items-center justify-center text-2xl font-bold">
          {{ form.name ? form.name.charAt(0).toUpperCase() : '' }}
        </div>
        <div>
          <h2 class="text-xl font-semibold text-gray-900 dark:text-gray-100">{{ form.name }}</h2>
          <p class="text-gray-600 dark:text-gray-400">{{ form.email }}</p>
        </div>
      </div>

      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Full Name</label>
          <input v-model="form.name" :disabled="!isEditing" class="input dark:bg-brand-700 dark:border-brand-600 dark:text-gray-100" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Email</label>
          <input v-model="form.email" :disabled="!isEditing" type="email" class="input dark:bg-brand-700 dark:border-brand-600 dark:text-gray-100" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Phone</label>
          <input v-model="form.phone" :disabled="!isEditing" type="tel" class="input dark:bg-brand-700 dark:border-brand-600 dark:text-gray-100" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Country</label>
          <input v-model="form.country" :disabled="!isEditing" type="text" class="input dark:bg-brand-700 dark:border-brand-600 dark:text-gray-100" />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Address</label>
          <textarea v-model="form.address" :disabled="!isEditing" rows="3" class="input dark:bg-brand-700 dark:border-brand-600 dark:text-gray-100" />
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
          <button @click="updateProfile" :disabled="userLoading" class="btn-primary">
            {{ userLoading ? 'Saving...' : 'Save Changes' }}
          </button>
        </template>
      </div>
    </div>
  </div>
</template>
