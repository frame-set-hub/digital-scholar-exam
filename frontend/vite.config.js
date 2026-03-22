import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'
import { fileURLToPath, URL } from 'node:url'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  const proxyTarget = env.API_PROXY_TARGET || 'http://localhost:8080'
  const rawPort = env.DEV_SERVER_PORT
  const parsedPort = rawPort ? Number(rawPort) : 5173
  const devPort =
    Number.isFinite(parsedPort) && parsedPort > 0 && parsedPort < 65536
      ? parsedPort
      : 5173

  return {
    plugins: [vue(), tailwindcss()],
    server: {
      port: devPort,
      proxy: {
        '/api': {
          target: proxyTarget,
          changeOrigin: true,
        },
      },
    },
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
  }
})
