import {useAuth} from '@/auth/hooks';
import Spinner from '@/components/ui/spinner';
import {Link, Navigate} from 'react-router-dom';

export default function Home() {
  const {isLoading, isAuthenticated} = useAuth();

  if (isLoading) {
    return <Spinner />;
  }

  if (isAuthenticated) {
    return <Navigate to="/settlements" replace={true} />;
  }

  return (
    <>
      <h1 className="mb-4 text-5xl font-extrabold leading-none tracking-tight">
        Datamonster
      </h1>
      <Link to="/signin">Sign In</Link>
    </>
  );
}
