import { defineConfig } from "vite";

export default defineConfig({
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
    },
  },
  build: {
    outDir: "dist",
    sourcemap: true,
  },
});
