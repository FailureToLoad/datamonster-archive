import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {createBrowserRouter, RouterProvider} from 'react-router-dom';
import RootLayout from '@/components/root-layout';

import IndexPage from '@/routes';
import SignInPage from '@/routes/signin';
import ProtectedLayout from '@/components/protected-layout';
import SettlementsPage from '@/routes/settlements/index.tsx';

import StorageTab from '@/routes/settlement/settlementStorage.tsx';
import TimelineTab from '@/routes/settlement/timeline.tsx';
import PopulationTab from '@/routes/settlement/population';
import SettlementPage from '@/routes/settlement/index.tsx';
import ErrorPage from './components/error';
import {QueryClient, QueryClientProvider} from '@tanstack/react-query';

const queryClient = new QueryClient();
const router = createBrowserRouter([
  {
    element: <RootLayout />,
    errorElement: <ErrorPage />,
    children: [
      {path: '/', element: <IndexPage />},
      {path: '/signin/*', element: <SignInPage />},
      {
        element: <ProtectedLayout />,
        children: [
          {
            path: '/settlements',
            element: <SettlementsPage />,
          },
          {
            path: '/settlements/:settlementId',
            element: <SettlementPage />,
            children: [
              {path: 'population', element: <PopulationTab />},
              {path: 'storage', element: <StorageTab />},
              {path: 'timeline', element: <TimelineTab />},
            ],
          },
        ],
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>
);
