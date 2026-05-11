export default defineNuxtPlugin(() => {
  const isDark = useState('theme-dark', () => {
    if (process.client) {
      const stored = localStorage.getItem('theme')
      if (stored) {
        return stored === 'dark'
      }
      return window.matchMedia('(prefers-color-scheme: dark)').matches
    }
    return false
  })

  const applyTheme = () => {
    if (process.client) {
      if (isDark.value) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
      localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
    }
  }

  // Initialize theme on client side
  if (process.client) {
    applyTheme()
  }

  // Watch for theme changes
  watch(isDark, () => {
    applyTheme()
  })
})