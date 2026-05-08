export default defineNuxtPlugin(() => {
  if (process.client) {
    const isDark = useState('theme-dark', () => {
      const stored = localStorage.getItem('theme')
      if (stored) {
        return stored === 'dark'
      }
      return window.matchMedia('(prefers-color-scheme: dark)').matches
    })

    const applyTheme = () => {
      if (isDark.value) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
      if (process.client) {
        localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
      }
    }

    onMounted(() => {
      applyTheme()
    })

    watch(isDark, () => {
      applyTheme()
    })
  }
})