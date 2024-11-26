// src/components/root-layout.tsx
import {Outlet} from 'react-router-dom';
import Spinner from '@/components/ui/spinner';
import {useAuth} from '@/auth/hooks';

const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID;
if (!GOOGLE_CLIENT_ID) {
  throw new Error('Missing Google Client ID');
}

export default function RootLayout() {
  const {isLoading} = useAuth();

  return (
    <div className="default flex h-screen flex-col items-center justify-center bg-background">
      {isLoading ? <Spinner /> : <Outlet />}
    </div>
  );
}
