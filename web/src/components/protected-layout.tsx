/* eslint-disable react-hooks/exhaustive-deps */
import {useAuth} from '@clerk/clerk-react';
import Spinner from '@/components/ui/spinner';
import {Outlet, useNavigate} from 'react-router-dom';
import PopulationContextProvider from '@/components/context/populationContextProvider';
import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  createHttpLink,
} from '@apollo/client';
import {setContext} from '@apollo/client/link/context';

const httpLink = createHttpLink({
  uri: '/graphql',
  credentials: 'include',
});

export default function ProtectedLayout() {
  const {userId, isLoaded} = useAuth();
  const navigate = useNavigate();
  const {getToken} = useAuth();
  const authLink = setContext(async (_, {headers}) => {
    const token = await getToken();
    return {headers: {...headers, Authorization: `Bearer ${token}`}};
  });
  const apolloClient = new ApolloClient({
    cache: new InMemoryCache(),
    link: authLink.concat(httpLink),
  });

  if (isLoaded && !userId) {
    navigate('/sign-in');
  }

  if (!isLoaded) return <Spinner />;

  return (
    <ApolloProvider client={apolloClient}>
      <div className="flex h-screen flex-col items-center justify-center">
        <PopulationContextProvider>
          <Outlet />
        </PopulationContextProvider>
      </div>
    </ApolloProvider>
  );
}
