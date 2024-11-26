// src/lib/auth/utils.ts
import {AuthUser, CredentialResponse} from './types';

export const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID;

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
    });

    if (!result.ok) {
      throw new Error('Failed to authenticate');
    }

    const data = await result.json();

    // Store the token
    localStorage.setItem('auth_token', data.token);

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

export async function validateAuthToken(
  token: string
): Promise<AuthUser | null> {
  try {
    const result = await fetch('/api/auth/validate', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!result.ok) {
      localStorage.removeItem('auth_token');
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
    localStorage.removeItem('auth_token');
    return null;
  }
}
