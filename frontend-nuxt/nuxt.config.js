export default defineNuxtConfig({
  ssr: true,
  devtools: { enabled: true },

  experimental: {
    payloadExtraction: false,
    inlineSSGStyles: false,
    renderJsonPayloads: false
  },

  nitro: {
    openAPI: false,
    preCompression: false
  },

  routeRules: {
    '/api/**': { cors: true },
    '/**': { ssr: true }
  },

  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || process.env.API_BASE_URL || 'http://localhost:8080'
    }
  },

  modules: [
    '@pinia/nuxt',
    '@nuxtjs/tailwindcss'
  ],

  app: {
    head: {
      title: 'E-Commerce Store',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Modern e-commerce store built with Nuxt 3' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    },
    errorPage: '/error'
  },

  css: ['~/assets/css/main.css'],

  tailwindcss: {
    cssPath: '~/assets/css/tailwind.css',
    configPath: 'tailwind.config.js'
  },

  pinia: {
    storesDirs: ['./stores/**']
  },

  compatibilityDate: '2026-03-27'
})