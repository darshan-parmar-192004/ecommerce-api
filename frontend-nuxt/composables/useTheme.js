export const useTheme = () => {
  const isDark = useState('theme-dark')

  const toggleTheme = () => {
    isDark.value = !isDark.value
  }

  return {
    isDark,
    toggleTheme
  }
}