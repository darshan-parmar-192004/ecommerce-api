<script setup>
import { ReloadIcon } from '@radix-icons/vue';
import { cn } from "@/lib/utils";

const props = defineProps({
  class: {
    type: [Boolean, null, String, Object, Array],
    required: false,
    skipCheck: true,
  },
});
</script>

<template>
  <ReloadIcon role="status" aria-label="Loading" :class="cn('size-4 animate-spin', props.class)" />
</template>
