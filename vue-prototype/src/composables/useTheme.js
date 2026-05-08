import { ref, watch } from 'vue'
const isDark = ref(false)
let initialized = false
export function useTheme() {
  const initTheme = () => {
    if (initialized) return
    const savedTheme = localStorage.getItem('theme')
    if (savedTheme === 'dark' || (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
      isDark.value = true
      document.documentElement.classList.add('dark')
    }
    initialized = true
  }
  const toggleTheme = () => {
    isDark.value = !isDark.value
    if (isDark.value) {
      document.documentElement.classList.add('dark')
      localStorage.setItem('theme', 'dark')
    } else {
      document.documentElement.classList.remove('dark')
      localStorage.setItem('theme', 'light')
    }
  }
  watch(isDark, (value) => {
    if (value) document.documentElement.classList.add('dark')
    else document.documentElement.classList.remove('dark')
  })
  return { isDark, initTheme, toggleTheme }
}
