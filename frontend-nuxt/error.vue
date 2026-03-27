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
        <svg class="mx-auto h-24 w-24 text-indigo-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
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
      <button 
        @click="handleError"
        class="btn-primary inline-flex items-center gap-2"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
        </svg>
        Go Home
      </button>
    </div>
  </div>
</template>
