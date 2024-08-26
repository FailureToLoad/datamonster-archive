import {defineConfig, loadEnv} from 'vite';
import path from 'path';
import react from '@vitejs/plugin-react-swc';

export default defineConfig(({mode}) => {
  const env = loadEnv(mode, process.cwd());
  const PORT = `${env.VITE_PORT ?? '8090'}`;
  return {
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
      port: Number(PORT),
    },
    preview: {
      host: '0.0.0.0',
      strictPort: true,
      port: 8090,
    },
  };
});
