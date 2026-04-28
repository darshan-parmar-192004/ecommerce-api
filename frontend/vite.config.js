import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
  server: {
    port: 3000,
    proxy: {
      "/auth": {
        target: "http://127.0.0.1:8080",
        changeOrigin: true,
        configure: (proxy, options) => {
          proxy.on('proxyRes', (proxyRes, req, res) => {
            const setCookie = proxyRes.headers['set-cookie'];
            if (setCookie) {
              res.setHeader('set-cookie', setCookie);
            }
          });
        },
      },
      "/products": {
        target: "http://127.0.0.1:8080",
        changeOrigin: true,
      },
      "/categories": {
        target: "http://127.0.0.1:8080",
        changeOrigin: true,
      },
      "/orders": {
        target: "http://127.0.0.1:8080",
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: "dist",
    sourcemap: true,
  },
});
