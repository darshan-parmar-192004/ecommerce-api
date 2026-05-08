<script setup>
import { computed } from 'vue'
import { ChevronsLeft, ChevronLeft, ChevronRight, ChevronsRight } from 'lucide-vue-next'

const props = defineProps({
  currentPage: { type: Number, default: 1 },
  totalPages: { type: Number, default: 1 },
  maxVisible: { type: Number, default: 7 }
})

const emit = defineEmits(['page-change'])

const pages = computed(() => {
  const half = Math.floor(props.maxVisible / 2)
  let start = Math.max(1, props.currentPage - half)
  let end = Math.min(props.totalPages, start + props.maxVisible - 1)
  
  if (end - start + 1 < props.maxVisible) {
    start = Math.max(1, end - props.maxVisible + 1)
  }
  
  const result = []
  for (let i = start; i <= end; i++) {
    result.push(i)
  }
  return result
})

const showLeftEllipsis = computed(() => pages.value[0] > 1)
const showRightEllipsis = computed(() => pages.value[pages.value.length - 1] < props.totalPages)

const goToPage = (page) => {
  if (page >= 1 && page <= props.totalPages && page !== props.currentPage) {
    emit('page-change', page)
  }
}
</script>

<template>
  <nav class="flex items-center justify-center gap-2 py-8" aria-label="Product pagination">
    <!-- First Page -->
    <button
      @click="goToPage(1)"
      :disabled="currentPage === 1"
      class="relative px-4 py-2.5 text-sm font-semibold rounded-xl backdrop-blur-xl transition-all duration-300 focus:ring-4 focus:ring-gray-900/30 dark:focus:ring-brand-700/30 focus:outline-none disabled:opacity-40 disabled:cursor-not-allowed group overflow-hidden"
      :class="currentPage === 1
        ? 'bg-gray-200/50 text-gray-400 border border-gray-300/40'
        : 'bg-white/50 dark:bg-brand-800/50 text-gray-700 dark:text-gray-300 border border-white/30 dark:border-brand-600/30 hover:bg-white/70 dark:hover:bg-brand-800/70 hover:border-gray-300/50 dark:hover:border-brand-600/50 hover:text-gray-900 dark:hover:text-gray-100 hover:shadow-lg hover:shadow-gray-900/20 dark:hover:shadow-brand-900/20 active:scale-95'"
      aria-label="First page"
    >
      <span class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></span>
      <ChevronsLeft class="w-4 h-4 relative z-10" aria-hidden="true" />
    </button>

    <!-- Previous -->
    <button
      @click="goToPage(currentPage - 1)"
      :disabled="currentPage === 1"
      class="relative flex items-center gap-1.5 px-5 py-2.5 text-sm font-semibold rounded-xl backdrop-blur-xl transition-all duration-300 focus:ring-4 focus:ring-gray-900/30 dark:focus:ring-brand-700/30 focus:outline-none disabled:opacity-40 disabled:cursor-not-allowed group overflow-hidden"
      :class="currentPage === 1
        ? 'bg-gray-200/50 text-gray-400 border border-gray-300/40'
        : 'bg-white/50 dark:bg-brand-800/50 text-gray-700 dark:text-gray-300 border border-white/30 dark:border-brand-600/30 hover:bg-white/70 dark:hover:bg-brand-800/70 hover:border-gray-300/50 dark:hover:border-brand-600/50 hover:text-gray-900 dark:hover:text-gray-100 hover:shadow-lg hover:shadow-gray-900/20 dark:hover:shadow-brand-900/20 active:scale-95'"
      aria-label="Previous page"
    >
      <span class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></span>
      <ChevronLeft class="w-4 h-4 relative z-10" aria-hidden="true" />
      <span class="hidden sm:inline relative z-10">Prev</span>
    </button>

    <!-- Page Numbers -->
    <div class="flex items-center gap-1.5">
      <span v-if="showLeftEllipsis" class="px-3 py-2 text-sm font-medium text-gray-500/80 dark:text-gray-400/80">...</span>

      <button
        v-for="page in pages"
        :key="page"
        @click="goToPage(page)"
        class="relative min-w-[44px] px-3.5 py-2.5 text-sm font-bold rounded-xl backdrop-blur-xl transition-all duration-300 focus:ring-4 focus:ring-gray-900/30 dark:focus:ring-brand-700/30 focus:outline-none overflow-hidden group border-2"
        :class="page === currentPage
          ? 'bg-gradient-to-br from-gray-900 to-gray-800 dark:from-brand-700 dark:to-brand-800 text-white border-gray-700 dark:border-brand-600 shadow-lg shadow-gray-900/40 dark:shadow-brand-900/40 scale-105'
          : 'bg-white/60 dark:bg-brand-800/60 text-gray-700 dark:text-gray-300 border-white/40 dark:border-brand-600/40 hover:bg-white/80 dark:hover:bg-brand-800/80 hover:border-gray-300/60 dark:hover:border-brand-600/60 hover:text-gray-900 dark:hover:text-gray-100 hover:shadow-md hover:shadow-gray-900/20 dark:hover:shadow-brand-900/20 active:scale-95'"
        :aria-current="page === currentPage ? 'page' : undefined"
      >
        <span class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></span>
        <span class="relative z-10">{{ page }}</span>
      </button>

      <span v-if="showRightEllipsis" class="px-3 py-2 text-sm font-medium text-gray-500/80 dark:text-gray-400/80">...</span>
    </div>

    <!-- Next -->
    <button
      @click="goToPage(currentPage + 1)"
      :disabled="currentPage === totalPages"
      class="relative flex items-center gap-1.5 px-5 py-2.5 text-sm font-semibold rounded-xl backdrop-blur-xl transition-all duration-300 focus:ring-4 focus:ring-gray-900/30 dark:focus:ring-brand-700/30 focus:outline-none disabled:opacity-40 disabled:cursor-not-allowed group overflow-hidden"
      :class="currentPage === totalPages
        ? 'bg-gray-200/50 text-gray-400 border border-gray-300/40'
        : 'bg-white/50 dark:bg-brand-800/50 text-gray-700 dark:text-gray-300 border border-white/30 dark:border-brand-600/30 hover:bg-white/70 dark:hover:bg-brand-800/70 hover:border-gray-300/50 dark:hover:border-brand-600/50 hover:text-gray-900 dark:hover:text-gray-100 hover:shadow-lg hover:shadow-gray-900/20 dark:hover:shadow-brand-900/20 active:scale-95'"
      aria-label="Next page"
    >
      <span class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></span>
      <span class="hidden sm:inline relative z-10">Next</span>
      <ChevronRight class="w-4 h-4 relative z-10" aria-hidden="true" />
    </button>

    <!-- Last Page -->
    <button
      @click="goToPage(totalPages)"
      :disabled="currentPage === totalPages"
      class="relative px-4 py-2.5 text-sm font-semibold rounded-xl backdrop-blur-xl transition-all duration-300 focus:ring-4 focus:ring-gray-900/30 dark:focus:ring-brand-700/30 focus:outline-none disabled:opacity-40 disabled:cursor-not-allowed group overflow-hidden"
      :class="currentPage === totalPages
        ? 'bg-gray-200/50 text-gray-400 border border-gray-300/40'
        : 'bg-white/50 dark:bg-brand-800/50 text-gray-700 dark:text-gray-300 border border-white/30 dark:border-brand-600/30 hover:bg-white/70 dark:hover:bg-brand-800/70 hover:border-gray-300/50 dark:hover:border-brand-600/50 hover:text-gray-900 dark:hover:text-gray-100 hover:shadow-lg hover:shadow-gray-900/20 dark:hover:shadow-brand-900/20 active:scale-95'"
      aria-label="Last page"
    >
      <span class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></span>
      <ChevronsRight class="w-4 h-4 relative z-10" aria-hidden="true" />
    </button>
  </nav>
</template>