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

  modules: ["@pinia/nuxt", "@nuxtjs/tailwindcss"],

  pinia: {
    storesDirs: ["./stores/**"],
  },

  imports: {
    dirs: ["./composables"],
  },

  css: ["~/assets/css/main.css"],

  tailwindcss: {
    cssPath: "~/assets/css/tailwind.css",
    configPath: "tailwind.config.js",
  },

  app: {
    head: {
      title: "E-Commerce Store",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        { name: "description", content: "Modern e-commerce store" },
      ],
      link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
    },
    pageTransition: { name: "page", mode: "out-in" },
  },

  nitro: {
    serveStatic: {
      strict: false,
    },
  },

  runtimeConfig: {
    public: {
      apiBase: "http://127.0.0.1:8080",
    },
  },

  compatibilityDate: "2024-04-03",
});
