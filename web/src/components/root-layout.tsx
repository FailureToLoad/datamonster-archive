import {Outlet} from 'react-router-dom';
import {ClerkLoaded, ClerkLoading} from '@clerk/clerk-react';
import Spinner from '@/components/ui/spinner';
import AuthProvider from './auth-provider';

const PUBLISHABLE_KEY = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY;

if (!PUBLISHABLE_KEY) {
  throw new Error('Missing Publishable Key');
}

export default function RootLayout() {
  return (
    <AuthProvider>
        <div className="default flex h-screen flex-col items-center justify-center bg-background">
          <ClerkLoading>
            <Spinner />
          </ClerkLoading>
          <ClerkLoaded>
            <Outlet />
          </ClerkLoaded>
        </div>
    </AuthProvider>
  );
}
