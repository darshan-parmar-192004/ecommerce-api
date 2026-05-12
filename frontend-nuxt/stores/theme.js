import { defineStore } from 'pinia'

export const useThemeStore = defineStore('theme', {
  state: () => ({
    themeMode: 'light', // 'light' or 'dark'
    initialized: false
  }),

  getters: {
    isDark: (state) => state.themeMode === 'dark',
    currentMode: (state) => state.themeMode
  },

  actions: {
    initializeTheme() {
      if (import.meta.client && !this.initialized) {
        const stored = localStorage.getItem('theme')
        if (stored) {
          this.themeMode = stored
        } else {
          const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
          this.themeMode = prefersDark ? 'dark' : 'light'
        }
        this.applyTheme()
        this.initialized = true
      }
    },

    toggleTheme() {
      this.themeMode = this.themeMode === 'dark' ? 'light' : 'dark'
      this.applyTheme()
    },

    setThemeMode(mode) {
      this.themeMode = mode
      this.applyTheme()
    },

    applyTheme() {
      if (import.meta.client) {
        const html = document.documentElement
        if (this.themeMode === 'dark') {
          html.classList.add('dark')
          html.classList.remove('light')
        } else {
          html.classList.add('light')
          html.classList.remove('dark')
        }
        localStorage.setItem('theme', this.themeMode)
      }
    }
  }
})