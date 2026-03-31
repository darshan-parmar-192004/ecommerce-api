<script setup>
definePageMeta({
  layout: 'auth'
})

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const { success: showSuccess } = useToast()

const form = reactive({
  email: '',
  password: '',
  name: ''
})

const showPassword = ref(false)
const error = ref('')
const loading = ref(false)
const success = ref(false)
const countdown = ref(0)

const emailTouched = ref(false)

const passwordStrength = computed(() => {
  const pwd = form.password
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
  if (!emailTouched.value) return null
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(form.email)
})

const handleSubmit = async () => {
  if (!form.email || !form.password || !form.name) {
    error.value = 'Please fill in all required fields'
    return
  }

  if (form.password.length < 6) {
    error.value = 'Password must be at least 6 characters'
    return
  }

  error.value = ''
  loading.value = true

  try {
    const response = await fetch('/api/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form)
    })
    
    const data = await response.json()
    
    if (!response.ok) {
      let errMsg = 'Registration failed'
      if (data.error) {
        if (typeof data.error === 'object') {
          errMsg = data.error.message || JSON.stringify(data.error)
        } else {
          errMsg = data.error
        }
      } else if (data.message) {
        errMsg = data.message
      }
      throw new Error(errMsg)
    }
    
    success.value = true
    showSuccess('Account created successfully!')
    
    countdown.value = 3
    const interval = setInterval(() => {
      countdown.value--
      if (countdown.value <= 0) {
        clearInterval(interval)
        router.push('/auth/login')
      }
    }, 1000)
  } catch (err) {
    error.value = err.message || 'Something went wrong'
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
        <span class="text-4xl font-bold bg-gradient-to-r from-primary-600 to-accent-600 bg-clip-text text-transparent animate-pulse-slow">
          Store
        </span>
      </NuxtLink>
      <h2 class="mt-6 text-2xl font-bold text-gray-900">
        {{ success ? 'Account Created!' : 'Create your account' }}
      </h2>
      <p class="mt-2 text-gray-600">
        {{ success ? 'Redirecting to login...' : 'Join us and start shopping' }}
      </p>
    </div>

    <div class="bg-white/80 backdrop-blur-sm rounded-2xl shadow-xl p-8 border border-gray-100 relative overflow-hidden">
      <!-- Success State -->
      <Transition name="fade" mode="out-in">
        <div v-if="success" class="text-center py-8">
          <div class="w-20 h-20 mx-auto mb-4 bg-green-100 rounded-full flex items-center justify-center animate-bounce-subtle">
            <svg class="w-10 h-10 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <p class="text-lg font-medium text-gray-900 mb-2">Registration Successful!</p>
          <p class="text-gray-500 mb-4">Your account has been created.</p>
          <p class="text-sm text-gray-400">Redirecting to login in {{ countdown }} second{{ countdown !== 1 ? 's' : '' }}...</p>
        </div>

        <form v-else @submit.prevent="handleSubmit" class="space-y-5">
          <!-- Error Alert -->
          <div v-if="error" class="bg-red-50 border border-red-200 text-red-600 p-4 rounded-xl text-sm flex items-center gap-3 animate-shake">
            <svg class="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            {{ error }}
          </div>

          <!-- Name Field -->
          <div>
            <label for="name" class="block text-sm font-semibold text-gray-700 mb-2">
              Full Name
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>
              <input 
                id="name"
                v-model="form.name"
                type="text"
                required
                class="w-full pl-12 pr-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-primary-500 focus:ring-primary-200 focus:border-primary-500 focus:scale-[1.01] transition-all duration-200 bg-white"
                placeholder="John Doe"
              />
            </div>
          </div>

          <!-- Email Field -->
          <div>
            <label for="email" class="block text-sm font-semibold text-gray-700 mb-2">
              Email Address
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" />
                </svg>
              </div>
              <input 
                id="email"
                v-model="form.email"
                type="email"
                required
                @blur="emailTouched = true"
                :class="[
                  'w-full pl-12 pr-4 py-3 border rounded-xl focus:ring-2 focus:ring-primary-500 focus:ring-primary-200 focus:border-primary-500 focus:scale-[1.01] transition-all duration-200 bg-white',
                  isValidEmail === false ? 'border-red-500 bg-red-50' : 'border-gray-300'
                ]"
                placeholder="you@example.com"
              />
            </div>
            <p v-if="isValidEmail === false" class="mt-1 text-xs text-red-500">Please enter a valid email address</p>
          </div>

          <!-- Password Field -->
          <div>
            <label for="password" class="block text-sm font-semibold text-gray-700 mb-2">
              Password
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
              <input 
                id="password"
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                required
                minlength="6"
                class="w-full pl-12 pr-12 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-primary-500 focus:ring-primary-200 focus:border-primary-500 focus:scale-[1.01] transition-all duration-200 bg-white"
                placeholder="Min 6 characters"
              />
              <button 
                type="button"
                @click="showPassword = !showPassword"
                class="absolute inset-y-0 right-0 pr-4 flex items-center"
              >
                <svg v-if="!showPassword" class="w-5 h-5 text-gray-400 hover:text-gray-600 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                <svg v-else class="w-5 h-5 text-gray-400 hover:text-gray-600 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
                </svg>
              </button>
            </div>
            <!-- Password Strength -->
            <div v-if="form.password" class="mt-2">
              <div class="flex gap-1 mb-1">
                <div :class="['h-1 flex-1 rounded-full transition-all duration-300', passwordStrength.level >= 1 ? passwordStrength.color : 'bg-gray-200']" />
                <div :class="['h-1 flex-1 rounded-full transition-all duration-300', passwordStrength.level >= 2 ? passwordStrength.color : 'bg-gray-200']" />
                <div :class="['h-1 flex-1 rounded-full transition-all duration-300', passwordStrength.level >= 3 ? passwordStrength.color : 'bg-gray-200']" />
              </div>
              <p :class="['text-xs', passwordStrength.level === 1 ? 'text-red-500' : passwordStrength.level === 2 ? 'text-yellow-500' : 'text-green-500']">{{ passwordStrength.text }}</p>
            </div>
          </div>

          <!-- Submit Button -->
          <button 
            type="submit"
            :disabled="loading"
            class="w-full py-3 px-4 bg-gradient-to-r from-primary-600 to-accent-600 text-white font-semibold rounded-xl hover:from-primary-700 hover:to-accent-700 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-all flex items-center justify-center gap-2 active:scale-[0.98] hover:scale-[1.01] shadow-lg shadow-primary-500/25"
          >
            <svg v-if="loading" class="animate-spin h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ loading ? 'Creating account...' : 'Create Account' }}
          </button>

          <!-- Login Link -->
          <p class="text-center text-gray-600">
            Already have an account?
            <NuxtLink to="/auth/login" class="text-primary-600 hover:text-primary-700 font-semibold">
              Sign in
            </NuxtLink>
          </p>
        </form>
      </Transition>
    </div>
  </div>
</template>

<style scoped>
.animate-pulse-slow {
  animation: pulse 3s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.animate-bounce-subtle {
  animation: bounceSubtle 0.5s ease-out;
}

@keyframes bounceSubtle {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
