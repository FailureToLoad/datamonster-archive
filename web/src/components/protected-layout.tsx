/* eslint-disable react-hooks/exhaustive-deps */
import {useAuth} from '@/auth/hooks';
import Spinner from '@/components/ui/spinner';
import {Outlet, useNavigate} from 'react-router-dom';
import PopulationContextProvider from '@/components/context/populationContextProvider';
import {useEffect} from 'react';

export default function ProtectedLayout() {
  const {user, isLoading} = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (!isLoading && !user) {
      navigate('/signin');
    }
  }, [isLoading, user]);

  if (isLoading) return <Spinner />;

  return (
    <div className="flex h-screen flex-col items-center justify-center">
      <PopulationContextProvider>
        <Outlet />
      </PopulationContextProvider>
    </div>
  );
}
