<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useToastStore } from '@/stores/toast'
import { useFormValidation } from '@/composables/useFormValidation'

const router = useRouter()
const authStore = useAuthStore()
const toastStore = useToastStore()

const registerSchema = {
  name: {
    required: true,
    requiredMessage: 'Full name is required',
    minLength: 2,
    minLengthMessage: 'Name must be at least 2 characters'
  },
  email: {
    required: true,
    requiredMessage: 'Email is required',
    email: true,
    emailMessage: 'Please enter a valid email address'
  },
  password: {
    required: true,
    requiredMessage: 'Password is required',
    minLength: 6,
    minLengthMessage: 'Password must be at least 6 characters'
  },
  confirmPassword: {
    required: true,
    requiredMessage: 'Please confirm your password',
    custom: (value, form) => {
      if (value !== form.password) {
        return 'Passwords do not match'
      }
      return null
    }
  }
}

const {
  form,
  errors,
  validate,
  handleBlur
} = useFormValidation(registerSchema, {
  name: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const showPassword = ref(false)

const handleRegister = async () => {
  if (!validate()) return

  try {
    await authStore.register({
      name: form.name,
      email: form.email,
      password: form.password
    })
    toastStore.success(`Welcome, ${authStore.user?.name || 'User'}!`)
    if (authStore.isAdmin) {
      router.push({ name: 'AdminDashboard' })
    } else {
      router.push('/')
    }
  } catch (err) {
    toastStore.error(err.response?.data?.message || 'Registration failed. Please try again.')
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <RouterLink to="/" class="inline-flex items-center text-sm text-gray-600 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-100 mb-4">
        <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
        </svg>
        Back to home
      </RouterLink>
    </div>
    <div>
      <h2 class="text-2xl font-bold text-gray-900 dark:text-gray-100">Create account</h2>
      <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">Get started with your free account</p>
    </div>

    <form @submit.prevent="handleRegister" class="space-y-4">
      <div>
        <label for="name" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Full Name</label>
        <input
          id="name"
          v-model="form.name"
          type="text"
          required
          class="input"
          :class="{ 'border-red-300 focus:ring-red-500': errors.name }"
          placeholder="John Doe"
          @blur="handleBlur('name')"
        />
        <p v-if="errors.name" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ errors.name }}</p>
      </div>

      <div>
        <label for="email" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Email</label>
        <input
          id="email"
          v-model="form.email"
          type="email"
          required
          class="input"
          :class="{ 'border-red-300 focus:ring-red-500': errors.email }"
          placeholder="you@example.com"
          @blur="handleBlur('email')"
        />
        <p v-if="errors.email" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ errors.email }}</p>
      </div>

      <div>
        <label for="password" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Password</label>
        <div class="relative">
          <input
            id="password"
            v-model="form.password"
            :type="showPassword ? 'text' : 'password'"
            required
            class="input pr-10"
            :class="{ 'border-red-300 focus:ring-red-500': errors.password }"
            placeholder="••••••••"
            @blur="handleBlur('password')"
          />
          <button
            type="button"
            @click="showPassword = !showPassword"
            class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
            :aria-label="showPassword ? 'Hide password' : 'Show password'"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path v-if="!showPassword" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
            </svg>
          </button>
        </div>
        <p v-if="errors.password" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ errors.password }}</p>
      </div>

      <div>
        <label for="confirmPassword" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Confirm Password</label>
        <input
          id="confirmPassword"
          v-model="form.confirmPassword"
          type="password"
          required
          class="input"
          :class="{ 'border-red-300 focus:ring-red-500': errors.confirmPassword }"
          placeholder="••••••••"
          @blur="handleBlur('confirmPassword')"
        />
        <p v-if="errors.confirmPassword" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ errors.confirmPassword }}</p>
      </div>

      <div v-if="authStore.error" class="text-sm text-red-600 dark:text-red-400 bg-red-50 dark:bg-red-900/20 p-3 rounded-lg">
        {{ authStore.error }}
      </div>

      <button
        type="submit"
        :disabled="authStore.loading"
        class="btn-primary w-full"
      >
        <span v-if="authStore.loading">Creating account...</span>
        <span v-else>Create account</span>
      </button>
    </form>

    <p class="text-sm text-center text-gray-600 dark:text-gray-400">
      Already have an account?
      <RouterLink to="/auth/login" class="font-medium text-gray-900 dark:text-gray-100 hover:underline">
        Sign in
      </RouterLink>
    </p>
  </div>
</template>
