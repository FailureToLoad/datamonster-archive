// src/lib/auth/GoogleAuthContext.tsx
import {createContext, useEffect, useState} from 'react';
import {AuthUser} from './types';

interface AuthContextType {
  user: AuthUser | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  getToken: () => Promise<string | null>;
  signIn: () => void;
  signOut: () => void;
}

export const AuthContext = createContext<AuthContextType | null>(null);

const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID;

export function GoogleAuthProvider({children}: {children: React.ReactNode}) {
  const [user, setUser] = useState<AuthUser | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Initialize Google OAuth
    window.google?.accounts.id.initialize({
      client_id: GOOGLE_CLIENT_ID,
      callback: handleCredentialResponse,
    });

    // Check existing token
    const token = localStorage.getItem('auth_token');
    if (token) {
      validateToken(token);
    } else {
      setIsLoading(false);
    }
  }, []);

  const handleCredentialResponse = async (response: {credential: string}) => {
    try {
      const result = await fetch('/auth/callback', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({credential: response.credential}),
      });

      if (!result.ok) {
        throw new Error('Failed to authenticate');
      }

      const data = await result.json();
      localStorage.setItem('auth_token', data.token);

      setUser({
        id: data.userId,
        email: data.email,
        name: data.name,
        picture: data.picture,
      });
    } catch (error) {
      console.error('Authentication error:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const validateToken = async (token: string) => {
    try {
      const result = await fetch('/auth/validate', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!result.ok) {
        localStorage.removeItem('auth_token');
        setUser(null);
        return;
      }

      const data = await result.json();
      setUser({
        id: data.userId,
        email: data.email,
        name: data.name,
        picture: data.picture,
      });
    } catch (error) {
      console.error('Token validation error:', error);
      localStorage.removeItem('auth_token');
      setUser(null);
    } finally {
      setIsLoading(false);
    }
  };

  const getToken = async () => {
    return localStorage.getItem('auth_token');
  };

  const signIn = () => {
    window.google?.accounts.id.prompt();
  };

  const signOut = () => {
    localStorage.removeItem('auth_token');
    setUser(null);
    if (user?.email) {
      window.google?.accounts.id.revoke(user.email, () => {
        console.log('User signed out');
      });
    }
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        isLoading,
        isAuthenticated: !!user,
        getToken,
        signIn,
        signOut,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
