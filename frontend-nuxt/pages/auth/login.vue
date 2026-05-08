<script setup>
definePageMeta({
  layout: 'auth'
})

const route = useRoute()
const { success: showSuccess } = useToast()

const form = reactive({
  email: '',
  password: ''
})

const showPassword = ref(false)
const error = ref('')
const loading = ref(false)

const { auth } = useApi()

const handleSubmit = async () => {
  if (!form.email || !form.password) {
    error.value = 'Please fill in all fields'
    return
  }

  error.value = ''
  loading.value = true

  try {
    const result = await auth.login(form)
    console.log('Login result:', result)

    showSuccess(`Welcome back, ${result.customer?.name || 'User'}!`)

    // Use navigateTo for reliable redirect after auth state is set
    const redirect = route.query.redirect || '/'
    await navigateTo(redirect)
  } catch (err) {
    error.value = err.message || 'Login failed'
  } finally {
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
      <form @submit.prevent="handleSubmit" class="space-y-5">
        <!-- Error Alert -->
        <div v-if="error" class="p-4 rounded-xl text-sm flex items-center gap-3 animate-shake bg-error-container text-error">
          <svg class="w-5 h-5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 11-18 0 9 9 0 0118 0zm-9 3.75h.008v.008H12v-.008z" />
          </svg>
          {{ error }}
        </div>

          <!-- Email Field -->
          <div>
            <label for="email" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant">
              Email Address
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75" />
                </svg>
              </div>
              <input 
                id="email"
                v-model="form.email"
                type="email"
                required
                class="w-full pl-12 pr-4 py-3.5 rounded-md transition-all duration-300 focus:outline-none bg-surface-container-low text-on_surface placeholder:text-outline/60 focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary border border-transparent"
                placeholder="you@example.com"
              />
            </div>
          </div>

          <!-- Password Field -->
          <div>
            <div class="flex items-center justify-between mb-3">
              <label for="password" class="block text-xs font-semibold uppercase tracking-wider text-on_surface_variant">
                Password
              </label>
              <NuxtLink to="#" class="text-xs font-medium transition-colors hover:text-primary text-primary">
                Forgot Password
              </NuxtLink>
            </div>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                <svg class="w-5 h-5 text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z" />
                </svg>
              </div>
              <input 
                id="password"
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                required
                class="w-full pl-12 pr-12 py-3.5 rounded-md transition-all duration-300 focus:outline-none bg-surface-container-low text-on_surface placeholder:text-outline/60 focus:bg-surface-container-lowest focus:ring-2 focus:ring-primary/20 focus:border-primary border border-transparent"
                placeholder="Enter your password"
              />
              <button 
                type="button"
                @click="showPassword = !showPassword"
                class="absolute inset-y-0 right-0 pr-4 flex items-center"
              >
                <svg v-if="!showPassword" class="w-5 h-5 transition-colors text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
                  <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                <svg v-else class="w-5 h-5 transition-colors text-outline" fill="none" stroke="currentColor" viewBox="0 0 24 24" stroke-width="1.5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M3.98 8.223A10.477 10.477 0 001.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M6.228 6.228A10.45 10.45 0 0112 4.5c4.756 0 8.773 3.162 10.065 7.498a10.52 10.52 0 01-4.293 5.774M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 10-4.243-4.243m4.242 4.242L9.88 9.88" />
                </svg>
              </button>
            </div>
          </div>

          <!-- Submit Button -->
          <button 
            type="submit"
            :disabled="loading"
            class="w-full py-3.5 px-4 text-white font-semibold rounded-md transition-all duration-300 flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed hover:shadow-glow hover:scale-[1.01] active:scale-[0.99] bg-gradient-to-r from-primary to-primary-container shadow-lg shadow-primary/25"
          >
            <svg v-if="loading" class="animate-spin h-5 w-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ loading ? 'Signing in...' : 'Sign In' }}
          </button>

          <!-- Register Link -->
          <p class="text-center pt-2 text-outline">
            <span class="text-sm">Don't have an account?</span>
            <NuxtLink to="/auth/register" class="text-sm font-semibold ml-1 transition-colors hover:text-primary text-primary">
              Create Account
            </NuxtLink>
          </p>
        </form>
    </div>

    <!-- Footer -->
    <div class="text-center mt-8">
      <p class="text-xs uppercase tracking-widest text-outline">
        Curated by Editorial Team
      </p>
    </div>
  </div>
</template>
