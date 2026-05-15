<script setup>
import { useUserStore } from '~/stores/user'
import { useAuthStore } from '~/stores/auth'
import { successMessages } from '~/constants/errorMessages'
const userStore = useUserStore()
const authStore = useAuthStore()

const form = reactive({ name: '', email: '', phone: '' })
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
    useAppToast().success(successMessages.profileUpdated)
    setTimeout(() => success.value = false, 3000)
  } catch (error) {
    console.error('Failed to update profile:', error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="max-w-2xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <h1 class="text-3xl font-bold text-on_surface mb-8">Profile Settings</h1>
    
    <div class="bg-surface-container-lowest rounded-lg p-6 shadow-ambient">
      <Message v-if="success" severity="success" :closable="false" class="mb-4">
        {{ successMessages.profileUpdated }}
      </Message>
      <form @submit.prevent="updateProfile" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-on_surface mb-2">Name</label>
          <InputText v-model="form.name" type="text" class="w-full" />
        </div>
        <div>
          <label class="block text-sm font-medium text-on_surface mb-2">Email</label>
          <InputText v-model="form.email" type="email" disabled class="w-full !bg-surface-container !cursor-not-allowed" />
        </div>
        <div>
          <label class="block text-sm font-medium text-on_surface mb-2">Phone</label>
          <InputText v-model="form.phone" type="tel" class="w-full" />
        </div>
        <Button type="submit" :loading="loading" :label="loading ? 'Saving...' : 'Save Changes'" />
      </form>
    </div>
  </div>
</template>
