<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import { useErrorHandler } from '@/composables/useErrorHandler'
import { useFormValidation } from '@/composables/useFormValidation'
import { ArrowLeft, Eye, EyeOff } from 'lucide-vue-next'

const router = useRouter()
const { user, loading, error, isAdmin, register } = useAuth()
const { showError, showSuccess } = useErrorHandler()

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
    await register({
      name: form.name,
      email: form.email,
      password: form.password
    })
    showSuccess(`Welcome, ${user?.name || 'User'}!`)
    if (isAdmin.value) {
      router.push({ name: 'AdminDashboard' })
    } else {
      router.push('/')
    }
  } catch (err) {
    showError(err, 'Registration failed. Please try again.')
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
            <Eye v-if="!showPassword" class="w-5 h-5" />
              <EyeOff v-else class="w-5 h-5" />
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

      <div v-if="error" class="text-sm text-red-600 dark:text-red-400 bg-red-50 dark:bg-red-900/20 p-3 rounded-lg">
        {{ error }}
      </div>

      <button
        type="submit"
        :disabled="loading"
        class="btn-primary w-full"
      >
        <span v-if="loading">Creating account...</span>
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
