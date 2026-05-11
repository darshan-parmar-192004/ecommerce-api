<script setup>
import { useForm } from 'vue-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import * as z from 'zod'

const route = useRoute()
const { success: showSuccess } = useAppToast()

const loginSchema = z.object({
  email: z.string().email('Please enter a valid email address'),
  password: z.string().min(1, 'Password is required')
})

const form = useForm({
  resolver: zodResolver(loginSchema),
  defaultValues: {
    email: '',
    password: ''
  }
})

const { handleSubmit, reset, setError, setErrors } = form

const { auth } = useApi()

const onSubmit = async (data) => {
  try {
    const result = await auth.login(data)
    console.log('Login result:', result)

    showSuccess(`Welcome back, ${result.customer?.name || 'User'}!`)

    const redirect = route.query.redirect || '/'
    await navigateTo(redirect)
  } catch (err) {
    setError('root', { message: err.message || 'Login failed' })
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
      <form @submit="handleSubmit(onSubmit)" class="space-y-5">
        <div>
          <label for="email" class="block text-xs font-semibold uppercase tracking-wider mb-3 text-on_surface_variant">
            Email Address
          </label>
          <IconField>
            <InputIcon class="pi pi-envelope" />
            <InputText
              id="email"
              v-model="form.email"
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
              v-model="form.watch('password')"
              type="password"
              required
              class="w-full !pl-10"
              placeholder="Enter your password"
            />
          </IconField>
        </div>

        <Button
          type="submit"
          label="Sign In"
          class="w-full !py-3"
        />

        <p class="text-center pt-2 text-outline">
          <span class="text-sm">Don't have an account?</span>
          <NuxtLink to="/auth/register" class="text-sm font-semibold ml-1 transition-colors hover:text-primary text-primary">
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
