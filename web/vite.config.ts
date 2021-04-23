import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    https: true,
    host: 'dev.localhost.com',
    proxy: {
      '/api': {
        target:
          process.env.APP_ENV === 'production'
            ? process.env.API_URL
            : 'https://dev.localhost.com:5000',
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
