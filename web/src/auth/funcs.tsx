import {AuthUser, CredentialResponse} from './types';

export const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID;

// In-memory token storage
let authToken: string | null = null;

export async function handleGoogleResponse(
  response: CredentialResponse
): Promise<AuthUser | null> {
  const {credential} = response;
  try {
    const result = await fetch('/api/auth/google', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({credential}),
      credentials: 'include',
    });

    if (!result.ok) {
      throw new Error('Failed to authenticate');
    }

    const data = await result.json();
    authToken = data.token;

    return {
      id: data.userId,
      email: data.email,
      name: data.name,
      picture: data.picture,
    };
  } catch (error) {
    console.error('Authentication error:', error);
    return null;
  }
}

export async function validateAuthToken(): Promise<AuthUser | null> {
  try {
    if (!authToken) {
      return null;
    }

    const result = await fetch('/api/auth/validate', {
      headers: {
        Authorization: `Bearer ${authToken}`,
      },
      credentials: 'include',
    });

    if (!result.ok) {
      authToken = null;
      return null;
    }

    const data = await result.json();
    return {
      id: data.userId,
      email: data.email,
      name: data.name,
      picture: data.picture,
    };
  } catch (error) {
    console.error('Token validation error:', error);
    authToken = null;
    return null;
  }
}

export function getAuthToken(): string | null {
  return authToken;
}

export function clearAuth(): void {
  authToken = null;
}
