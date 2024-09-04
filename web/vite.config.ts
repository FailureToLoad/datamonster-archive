import {defineConfig} from 'vite';
import path from 'path';
import react from '@vitejs/plugin-react-swc';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@types': path.resolve(__dirname, './src/__generated__/graphql'),
    },
  },
  server: {
    host: true,
    strictPort: true,
    port: 8090,
  },
});
