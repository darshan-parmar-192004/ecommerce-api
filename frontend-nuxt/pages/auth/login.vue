<script setup>
import { routes } from '~/utils/constants'

const email = ref('')
const password = ref('')

definePageMeta({
  layout: 'auth'
})

const route = useRoute()
const { success: showSuccess, error: showError } = useAppToast()

const loading = ref(false)
const formError = ref('')
const redirectUrl = ref(route.query.redirect || '/')

const { auth } = useApi()

const onSubmit = async () => {
  if (!email.value || !password.value) {
    formError.value = 'All fields are required'
    return
  }
  loading.value = true
  formError.value = ''
  try {
    const result = await auth.login({ email: email.value, password: password.value })
    showSuccess(`Welcome back, ${result.customer?.name || 'User'}!`)
    await navigateTo(redirectUrl.value)
  } catch (err) {
    const message = err.message || 'Login failed'
    formError.value = message
    showError(message)
    loading.value = false
  }
}

useSeoMeta({
  title: 'Login - E-Commerce Store'
})
</script>

<template>
  <div class="w-full max-w-md">
    <div class="text-center mb-10">
      <NuxtLink to="/" class="inline-block">
        <span class="text-3xl font-bold font-display tracking-tight text-primary">
          The Curator
        </span>
      </NuxtLink>
      <h2 class="mt-8 text-xl font-semibold text-on_surface font-display">
        Welcome back to your collection.
      </h2>
    </div>

    <div class="rounded-2xl p-8 relative overflow-hidden bg-surface-container-lowest shadow-ambient">
      <form @submit.prevent="onSubmit" class="space-y-5">
        <Message v-if="formError" severity="error" :closable="false" class="!mb-0">
          <i class="pi pi-exclamation-triangle mr-2" />
          {{ formError }}
        </Message>

        <div>
          <label for="email" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant">
            Email Address
          </label>
          <IconField>
            <InputIcon class="pi pi-envelope" />
            <InputText
              id="email"
              :model-value="email"
              @update:model-value="email = $event"
              type="email"
              required
              class="w-full !pl-10"
              placeholder="you@example.com"
            />
          </IconField>
        </div>

        <div>
          <div class="flex items-center justify-between mb-3">
            <label for="password" class="block text-xs font-semibold uppercase tracking-wider text-on_surface_variant">
              Password
            </label>
            <NuxtLink to="#" class="text-xs font-medium transition-colors hover:text-primary text-primary">
              Forgot Password
            </NuxtLink>
          </div>
          <IconField>
            <InputIcon class="pi pi-lock" />
            <InputText
              id="password"
              :model-value="password"
              @update:model-value="password = $event"
              type="password"
              required
              class="w-full !pl-10"
              placeholder="Enter your password"
            />
          </IconField>
        </div>

        <Button
          type="submit"
          :loading="loading"
          :label="loading ? 'Signing in...' : 'Sign In'"
          class="w-full !py-3"
        />

        <p class="text-center pt-2 text-outline">
          <span class="text-sm">Don't have an account?</span>
          <NuxtLink :to="routes.register" class="text-sm font-semibold ml-1 transition-colors hover:text-primary text-primary" aria-label="Create new account">
            Create Account
          </NuxtLink>
        </p>
      </form>
    </div>

    <div class="text-center mt-8">
      <p class="text-xs uppercase tracking-widest text-outline">
        Curated by Editorial Team
      </p>
    </div>
  </div>
</template>
