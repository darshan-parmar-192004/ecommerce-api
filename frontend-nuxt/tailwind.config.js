export default {
  content: [
    './components/**/*.{js,vue}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './composables/**/*.{js,ts}',
    './plugins/**/*.{js,ts}',
    './app.vue'
  ],
  theme: {
    extend: {
      colors: {
        background: '#f9f9fc',
        surface: '#f9f9fc',
        'surface-container-lowest': '#ffffff',
        'surface-container-low': '#f3f3f6',
        'surface-container': '#eeeef0',
        'surface-container-high': '#e8e8ea',
        'surface-container-highest': '#e2e2e5',
        on_surface: '#1a1c1e',
        on_surface_variant: '#444656',
        primary: '#1c31e3',
        'primary-container': '#3e51fb',
        'primary-fixed': '#dfe0ff',
        'primary-fixed-dim': '#bdc2ff',
        on_primary: '#ffffff',
        secondary: '#5f5e5e',
        'secondary-container': '#e5e2e1',
        on_secondary: '#ffffff',
        tertiary: '#903000',
        'tertiary-container': '#b84000',
        on_tertiary: '#ffffff',
        outline: '#757688',
        'outline-variant': '#c5c5d9',
        error: '#ba1a1a',
        'error-container': '#ffdad6',
        on_error: '#ffffff',
      },
      fontFamily: {
        display: ['Manrope', 'sans-serif'],
        body: ['Inter', 'sans-serif'],
      },
      borderRadius: {
        'sm': '0.375rem',
        'md': '0.75rem',
        'lg': '1rem',
        'xl': '1.25rem',
      },
      boxShadow: {
        'ambient': '0px 12px 32px -4px rgba(26, 28, 30, 0.06)',
        'glow': '0 4px 16px rgba(28, 49, 227, 0.25)',
      },
      animation: {
        'grid': 'grid 20s linear infinite',
        'float': 'float 6s ease-in-out infinite',
        'pulse-slow': 'pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite',
        'bounce-subtle': 'bounceSubtle 0.5s ease-out',
        'shimmer': 'shimmer 1.5s infinite',
        'stagger-fade': 'staggerFade 0.6s ease-out forwards',
      },
      keyframes: {
        grid: {
          '0%': { transform: 'translateY(0)' },
          '100%': { transform: 'translateY(50px)' },
        },
        float: {
          '0%, 100%': { transform: 'translateY(0)' },
          '50%': { transform: 'translateY(-20px)' },
        },
        bounceSubtle: {
          '0%, 100%': { transform: 'scale(1)' },
          '50%': { transform: 'scale(1.1)' },
        },
        shimmer: {
          '0%': { backgroundPosition: '-200% 0' },
          '100%': { backgroundPosition: '200% 0' },
        },
        staggerFade: {
          '0%': { opacity: '0', transform: 'translateY(24px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
      },
      transitionDuration: {
        '400': '400ms',
        '700': '700ms',
      },
    }
  },
  plugins: []
}
