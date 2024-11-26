// src/lib/auth/types.ts
export interface AuthUser {
  id: string;
  email: string;
  name: string | null;
  picture: string | null;
}

export interface AuthContextType {
  user: AuthUser | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  getToken: () => Promise<string | null>;
  signIn: () => void;
  signOut: () => void;
}

export interface CredentialResponse {
  credential: string;
  select_by: string;
}

export interface GoogleIdentity {
  accounts: {
    id: {
      initialize: (config: {
        client_id: string;
        callback: (response: CredentialResponse) => void;
        auto_select?: boolean;
      }) => void;
      prompt: (
        callback?: (notification: {
          isNotDisplayed: () => boolean;
          isSkippedMoment: () => boolean;
        }) => void
      ) => void;
      revoke: (email: string, callback: () => void) => void;
    };
  };
}

declare global {
  interface Window {
    google?: GoogleIdentity;
  }
}
