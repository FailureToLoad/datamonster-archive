import {defineConfig} from 'vitest/config';
import path from 'path';
import react from '@vitejs/plugin-react-swc';

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@types': path.resolve(__dirname, './src/__generated__/graphql'),
    },
  },
  server: {
    host: '0.0.0.0',
    strictPort: true,
    port: Number(8090),
  },
});
