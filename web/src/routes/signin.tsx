import { SignIn } from "@clerk/clerk-react";

export default function SignInPage() {
  return <SignIn path="/signin" forceRedirectUrl="/settlements" />;
}
