<script setup>
const name = ref('')
const email = ref('')
const password = ref('')

definePageMeta({
  layout: 'auth'
})

const { success: showSuccess } = useAppToast()

const showPassword = ref(false)
const error = ref('')
const loading = ref(false)
const success = ref(false)

const { auth } = useApi()

const passwordStrength = computed(() => {
  const pwd = password.value
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
  if (!email.value) return null
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email.value)
})

const onSubmit = async () => {
  if (!name.value || !email.value || !password.value) {
    error.value = 'All fields are required'
    return
  }
  if (password.value.length < 6) {
    error.value = 'Password must be at least 6 characters'
    return
  }
  try {
    loading.value = true
    await auth.register({ name: name.value, email: email.value, password: password.value })

    success.value = true
    showSuccess('Account created successfully!')
  } catch (err) {
    error.value = err.message || 'Registration failed'
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
    <div v-if="!success" class="text-center mb-10">
      <NuxtLink to="/" class="inline-block">
        <span class="text-3xl font-bold font-display tracking-tight text-primary">
          The Curator
        </span>
      </NuxtLink>
      <h2 class="mt-6 text-2xl font-bold text-on_surface font-display">
        Create your account
      </h2>
      <p class="mt-2 text-on_surface_variant">
        Join us and start shopping
      </p>
    </div>

    <Transition name="fade" mode="out-in">
      <RedirectScreen
        v-if="success"
        message="Registration Successful!"
        subtitle="Your account has been created."
        redirect-to="/"
      />

      <div v-else class="rounded-2xl p-8 relative overflow-hidden bg-surface-container-lowest shadow-ambient">
        <form @submit.prevent="onSubmit" class="flex flex-col gap-6">
          <Message v-if="error" severity="error" :closable="false">
            <i class="pi pi-exclamation-triangle mr-2" />
            {{ error }}
          </Message>

          <div>
            <label for="name" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant">
              Full Name
            </label>
            <IconField>
              <InputIcon class="pi pi-user" />
              <InputText
                id="name"
                :model-value="name"
                @update:model-value="name = $event"
                type="text"
                required
                class="w-full !pl-10"
                placeholder="John Doe"
              />
            </IconField>
          </div>

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
                :class="isValidEmail === false ? '!border-error !bg-error-container' : ''"
                placeholder="you@example.com"
              />
            </IconField>
            <div class="min-h-[1.25rem] mt-1">
              <p v-if="isValidEmail === false" class="text-xs text-error">
                Please enter a valid email address
              </p>
            </div>
          </div>

          <div>
            <label for="password" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant">
              Password
            </label>
            <IconField>
              <InputIcon class="pi pi-lock" />
              <InputText
                id="password"
                :model-value="password"
                @update:model-value="password = $event"
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
            <div v-if="password" class="mt-2">
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

          <p class="text-center text-on_surface_variant pt-2">
            Already have an account?
            <NuxtLink to="/auth/login" class="text-primary hover:text-primary/80 font-semibold">
              Sign in
            </NuxtLink>
          </p>
        </form>
      </div>
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
