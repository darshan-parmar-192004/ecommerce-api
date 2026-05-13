<script setup>
const props = defineProps({
  message: { type: String, required: true },
  subtitle: { type: String, default: '' },
  redirectTo: { type: String, required: true }
})

const countdown = ref(4)
let interval

onMounted(() => {
  interval = setInterval(() => {
    countdown.value--
    if (countdown.value <= 0) {
      clearInterval(interval)
      navigateTo(props.redirectTo)
    }
  }, 1000)
})

onUnmounted(() => {
  if (interval) clearInterval(interval)
})
</script>

<template>
  <div class="rounded-2xl p-12 bg-surface-container-lowest shadow-ambient text-center">
    <NuxtLink to="/" class="inline-block mb-6">
      <span class="text-3xl font-bold font-display tracking-tight text-primary">
        The Curator
      </span>
    </NuxtLink>

    <div class="flex justify-center mb-6">
      <i class="pi pi-spinner text-3xl text-primary animate-spin" />
    </div>

    <p class="text-lg font-medium text-on_surface mb-1">{{ message }}</p>
    <p v-if="subtitle" class="text-sm text-on_surface_variant mb-4">{{ subtitle }}</p>

    <p class="text-xs text-outline">
      Redirecting in {{ countdown }} second{{ countdown !== 1 ? 's' : '' }}...
    </p>
  </div>
</template>
