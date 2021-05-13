import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

const apiURL: string = process.env.API_URL
  ? process.env.API_URL
  : 'http://localhost:5000'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/api': {
        target: apiURL,
        changeOrigin: true,
        secure: process.env.APP_ENV === 'production',
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
