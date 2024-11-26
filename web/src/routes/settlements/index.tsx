// src/routes/settlements/index.tsx
import {SettlementCard} from '@/routes/settlements/card';
import {Settlement} from '@/lib/types/settlements';
import Spinner from '@/components/ui/spinner';
import {CreateSettlementDialog} from './creationDialog';
import {useAuth} from '@/auth/hooks';
import {useQuery} from '@tanstack/react-query';
import {GetSettlements} from '@/lib/services/settlement';

export const SettlementsQueryKey = 'settlements';
export default function SettlementsPage() {
  const {getToken} = useAuth();

  const getSettlements = async () => {
    try {
      const token = await getToken();
      if (!token) {
        return null;
      }
      const settlements = await GetSettlements(token);
      return settlements;
    } catch (e) {
      console.error('Failed to fetch settlements:', e);
      return null;
    }
  };

  const {isPending, isError, data, error} = useQuery({
    queryKey: [SettlementsQueryKey],
    queryFn: getSettlements,
  });

  if (isPending) {
    return <Spinner />;
  }

  if (isError) {
    throw new Error(error.message);
  }

  const settlements = data as Array<Settlement>;
  return (
    <main className="flex w-screen h-screen flex-col items-center justify-center overflow-hidden">
      <ul className="w-1/4 space-y-4 ">
        {settlements &&
          settlements.map((settlement) => (
            <li key={settlement.id}>
              <SettlementCard settlement={settlement} />
            </li>
          ))}
        <li key={-1}>
          <CreateSettlementDialog />
        </li>
      </ul>
    </main>
  );
}
