<script setup>
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import { useErrorHandler } from '@/composables/useErrorHandler'
import { useFormValidation } from '@/composables/useFormValidation'
import { ArrowLeft, Eye, EyeOff } from 'lucide-vue-next'

const router = useRouter()
const route = useRoute()
const { user, loading, error, isAdmin, login } = useAuth()
const { showError, showSuccess } = useErrorHandler()

const loginSchema = {
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
  }
}

const {
  form,
  errors,
  validate,
  handleBlur
} = useFormValidation(loginSchema, {
  email: '',
  password: ''
})

const showPassword = ref(false)

const handleLogin = async () => {
  if (!validate()) return

  try {
    await login(form)
    showSuccess(`Welcome back, ${user?.name || 'User'}!`)
    if (isAdmin.value) {
      router.push({ name: 'AdminDashboard' })
    } else {
      const redirect = route.query.redirect || '/'
      router.push(redirect)
    }
  } catch (err) {
    showError(err, 'Login failed. Please try again.')
  }
}
</script>

<template>
  <div class="space-y-6">
    <div>
      <RouterLink to="/" class="inline-flex items-center text-sm text-gray-600 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-100 mb-4">
        <ArrowLeft class="w-4 h-4 mr-1" />
        Back to home
      </RouterLink>
    </div>
    <div>
      <h2 class="text-2xl font-bold text-gray-900 dark:text-gray-100">Welcome back</h2>
      <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">Sign in to your account to continue</p>
    </div>

    <form @submit.prevent="handleLogin" class="space-y-4">
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
            <Eye v-if="!showPassword" class="w-5 h-5" />
              <EyeOff v-else class="w-5 h-5" />
          </button>
        </div>
        <p v-if="errors.password" class="mt-1 text-xs text-red-600 dark:text-red-400">{{ errors.password }}</p>
      </div>

      <div v-if="error" class="text-sm text-red-600 dark:text-red-400 bg-red-50 dark:bg-red-900/20 p-3 rounded-lg">
        {{ error }}
      </div>

      <button
        type="submit"
        :disabled="loading"
        class="btn-primary w-full"
      >
        <span v-if="loading">Signing in...</span>
        <span v-else>Sign in</span>
      </button>
    </form>

    <p class="text-sm text-center text-gray-600 dark:text-gray-400">
      Don't have an account?
      <RouterLink to="/auth/register" class="font-medium text-gray-900 dark:text-gray-100 hover:underline">
        Sign up
      </RouterLink>
    </p>
  </div>
</template>
