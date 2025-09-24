import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
  server: {
    port: 5173,
    proxy: {
      // Proxy API calls to PocketBase during development
      '/api': {
        target: 'http://localhost:8090',
        changeOrigin: true,
      },
      '/_': {
        target: 'http://localhost:8090',
        changeOrigin: true,
      }
    }
  },

  build: {
    emptyOutDir: true,
    sourcemap: true,
  },
});