import { useUser, SignInButton } from "@clerk/clerk-react";
import Spinner from "@/components/ui/spinner";
import { Navigate } from "react-router-dom";

export default function Home() {
  const { isLoaded, user } = useUser();
  if (!isLoaded) {
    return <Spinner />;
  }
  if (user) {
    return <Navigate to="/settlements" replace={true} />;
  }
  return (
    <>
      <h1 className="mb-4 text-5xl font-extrabold leading-none tracking-tight ">
        Datamonster
      </h1>
      <SignInButton />
    </>
  );
}
