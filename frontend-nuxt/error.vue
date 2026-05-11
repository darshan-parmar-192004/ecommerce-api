<script setup>
const error = useError()

const handleError = () => {
  clearError({ redirect: '/' })
}

useSeoMeta({
  title: () => `Error - ${error.value?.statusCode || '500'}`
})
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <div class="text-center max-w-md mx-auto px-4">
      <div class="mb-6">
        <i class="pi pi-exclamation-triangle text-6xl text-primary" />
      </div>
      <h1 class="text-6xl font-bold text-gray-900 mb-4">
        {{ error?.statusCode || '500' }}
      </h1>
      <p class="text-xl text-gray-600 mb-2">
        {{ error?.statusCode === 404 ? 'Page not found' : 'Something went wrong' }}
      </p>
      <p class="text-gray-500 mb-8">
        {{ error?.message || 'An unexpected error occurred' }}
      </p>
      <Button
        @click="handleError"
        icon="pi pi-home"
        label="Go Home"
      />
    </div>
  </div>
</template>
