import tailwindcss from "@tailwindcss/vite";
import Aura from "@primeuix/themes/aura";

export default defineNuxtConfig({
  ssr: true,

  devtools: { enabled: false },

  experimental: {
    ownMiddlewareResolution: true,
    payloadExtraction: false,
    inlineSSRStyles: true,
    appManifest: true,
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
        preset: Aura,
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

  css: ["primeicons/primeicons.css", "~/assets/css/main.css"],

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
        { rel: "preconnect", href: "https://fonts.googleapis.com" },
        { rel: "preconnect", href: "https://fonts.gstatic.com", crossorigin: "" },
        { rel: "stylesheet", href: "https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Manrope:wght@500;600;700;800&display=swap" },
      ],
    },
  },

  runtimeConfig: {
    public: {
      apiBase: "http://127.0.0.1:8080",
    },
  },

  compatibilityDate: "2024-04-03",
});
