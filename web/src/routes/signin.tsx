import {SignIn} from '@clerk/clerk-react';

export default function SignInPage() {
  return (
    <SignIn
      path="/signin"
      forceRedirectUrl="/settlements"
      appearance={{
        elements: {
          footer: {
            display: 'none',
          },
        },
      }}
    />
  );
}
