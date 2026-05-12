<script setup>
import { useForm, useController } from '@vuehookform/core'
import * as z from 'zod'

definePageMeta({
  layout: 'auth'
})

const { success: showSuccess } = useAppToast()

const registerSchema = z.object({
  name: z.string().min(2, 'Full name is required'),
  email: z.string().email('Please enter a valid email address'),
  password: z.string().min(6, 'Password must be at least 6 characters')
})

const form = useForm({
  schema: registerSchema,
  defaultValues: {
    name: '',
    email: '',
    password: ''
  }
})

const nameControl = useController({ name: 'name', control: form.control })
const emailControl = useController({ name: 'email', control: form.control })
const passwordControl = useController({ name: 'password', control: form.control })

const { handleSubmit, reset, setError, watch } = form

const showPassword = ref(false)
const error = ref('')
const loading = ref(false)
const success = ref(false)
const countdown = ref(0)

const { auth } = useApi()

const passwordStrength = computed(() => {
  const pwd = watch('password')
  if (!pwd) return { level: 0, text: '', color: '' }
  
  let score = 0
  if (pwd.length >= 6) score++
  if (pwd.length >= 8) score++
  if (/[a-z]/.test(pwd) && /[A-Z]/.test(pwd)) score++
  if (/\d/.test(pwd)) score++
  if (/[^a-zA-Z0-9]/.test(pwd)) score++
  
  if (score <= 2) return { level: 1, text: 'Weak', color: 'bg-red-500' }
  if (score <= 3) return { level: 2, text: 'Medium', color: 'bg-yellow-500' }
  return { level: 3, text: 'Strong', color: 'bg-green-500' }
})

const isValidEmail = computed(() => {
  const email = watch('email')
  if (!email) return null
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
})

const onSubmit = async (data) => {
  try {
    loading.value = true
    await auth.register(data)
    
    success.value = true
    showSuccess('Account created successfully!')
    
    countdown.value = 3
    const interval = setInterval(() => {
      countdown.value--
      if (countdown.value <= 0) {
        clearInterval(interval)
        navigateTo('/auth/login')
      }
    }, 1000)
  } catch (err) {
    setError('root', { message: err.message || 'Registration failed' })
  } finally {
    loading.value = false
  }
}

useSeoMeta({
  title: 'Register - E-Commerce Store'
})
</script>

<template>
  <div class="w-full max-w-md">
    <div class="text-center mb-8">
      <NuxtLink to="/" class="inline-block">
        <span class="text-4xl font-bold text-primary font-display">
          The Curator
        </span>
      </NuxtLink>
      <h2 class="mt-6 text-2xl font-bold text-on_surface font-display">
        {{ success ? 'Account Created!' : 'Create your account' }}
      </h2>
      <p class="mt-2 text-on_surface_variant">
        {{ success ? 'Redirecting to login...' : 'Join us and start shopping' }}
      </p>
    </div>

    <Transition name="fade" mode="out-in">
      <div v-if="success" class="text-center py-8">
        <div class="w-20 h-20 mx-auto mb-4 bg-primary-fixed rounded-full flex items-center justify-center">
          <i class="pi pi-check text-3xl text-primary" />
        </div>
        <p class="text-lg font-medium text-on_surface mb-2">Registration Successful!</p>
        <p class="text-on_surface_variant mb-4">Your account has been created.</p>
        <p class="text-sm text-outline">Redirecting to login in {{ countdown }} second{{ countdown !== 1 ? 's' : '' }}...</p>
      </div>

      <form v-else @submit="handleSubmit(onSubmit)" class="space-y-5">
        <Message v-if="error" severity="error" :closable="false">
          <i class="pi pi-exclamation-triangle mr-2" />
          {{ error }}
        </Message>

        <div>
          <label for="name" class="block text-sm font-semibold text-on_surface_variant mb-2">
            Full Name
          </label>
          <IconField>
            <InputIcon class="pi pi-user" />
            <InputText
              id="name"
              :value="nameControl.field.value"
              @update:model-value="nameControl.field.onChange"
              @blur="nameControl.field.onBlur"
              type="text"
              required
              class="w-full !pl-10"
              placeholder="John Doe"
            />
          </IconField>
        </div>

        <div>
          <label for="email" class="block text-sm font-semibold text-on_surface_variant mb-2">
            Email Address
          </label>
          <IconField>
            <InputIcon class="pi pi-envelope" />
            <InputText
              id="email"
              :value="emailControl.field.value"
              @update:model-value="emailControl.field.onChange"
              @blur="emailControl.field.onBlur"
              type="email"
              required
              :class="isValidEmail === false ? '!border-error !bg-error-container' : ''"
              placeholder="you@example.com"
            />
          </IconField>
          <Message v-if="isValidEmail === false" severity="error" :closable="false" class="!mt-1 !text-xs">
            Please enter a valid email address
          </Message>
        </div>

        <div>
          <label for="password" class="block text-sm font-semibold text-on_surface_variant mb-2">
            Password
          </label>
          <IconField>
            <InputIcon class="pi pi-lock" />
            <InputText
              id="password"
              :value="passwordControl.field.value"
              @update:model-value="passwordControl.field.onChange"
              @blur="passwordControl.field.onBlur"
              :type="showPassword ? 'text' : 'password'"
              required
              minlength="6"
              class="w-full !pl-10"
              placeholder="Min 6 characters"
            />
            <Button
              type="button"
              @click="showPassword = !showPassword"
              :icon="showPassword ? 'pi pi-eye-slash' : 'pi pi-eye'"
              text
              class="absolute right-2 top-1/2 -translate-y-1/2"
            />
          </IconField>
          <div v-if="form.watch('password')" class="mt-2">
            <div class="flex gap-1 mb-1">
              <div :class="['h-1 flex-1 rounded-full transition-all duration-300', passwordStrength.level >= 1 ? passwordStrength.color : 'bg-surface-container-high']" />
              <div :class="['h-1 flex-1 rounded-full transition-all duration-300', passwordStrength.level >= 2 ? passwordStrength.color : 'bg-surface-container-high']" />
              <div :class="['h-1 flex-1 rounded-full transition-all duration-300', passwordStrength.level >= 3 ? passwordStrength.color : 'bg-surface-container-high']" />
            </div>
            <p :class="['text-xs', passwordStrength.level === 1 ? 'text-error' : passwordStrength.level === 2 ? 'text-yellow-600' : 'text-green-600']">{{ passwordStrength.text }}</p>
          </div>
        </div>

        <Button
          type="submit"
          :loading="loading"
          :label="loading ? 'Creating account...' : 'Create Account'"
          class="w-full !py-3"
        />

        <p class="text-center text-on_surface_variant">
          Already have an account?
          <NuxtLink to="/auth/login" class="text-primary hover:text-primary/80 font-semibold">
            Sign in
          </NuxtLink>
        </p>
      </form>
    </Transition>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>