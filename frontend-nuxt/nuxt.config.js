import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  ssr: true,

  devtools: { enabled: false },

  experimental: {
    ownMiddlewareResolution: true,
    payloadExtraction: true,
    inlineSSRStyles: false,
    appManifest: false,
  },

  nitro: {
    serveStatic: {
      strict: false,
    },
    compressPublicAssets: true,
  },

  modules: ["@pinia/nuxt", "@primevue/nuxt-module"],

  primevue: {
    options: {
      theme: {
        preset: "Aura",
        options: {
          darkModeSelector: ".dark",
        },
      },
    },
    autoImport: true,
  },

  pinia: {
    storesDirs: ["./stores/**"],
  },

  imports: {
    dirs: ["./composables"],
  },

  css: ["~/assets/css/main.css"],

  vite: {
    plugins: [tailwindcss()],
  },

  app: {
    head: {
      title: "E-Commerce Store",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        { name: "description", content: "Modern e-commerce store" },
      ],
      link: [
        { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" },
        { rel: "shortcut icon", type: "image/svg+xml", href: "/favicon.svg" },
      ],
    },
    pageTransition: { name: "page", mode: "out-in" },
  },

  runtimeConfig: {
    public: {
      apiBase: "http://127.0.0.1:8080",
    },
  },

  compatibilityDate: "2024-04-03",
});
