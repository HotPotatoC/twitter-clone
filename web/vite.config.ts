import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/api': {
        target:
          process.env.APP_ENV === 'production'
            ? process.env.API_URL
            : 'http://localhost:5000',
        changeOrigin: true,
        secure: process.env.APP_ENV === 'production',
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
