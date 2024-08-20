import {SettlementCard} from '@/routes/settlements/card';
import {SettlementId} from '@/lib/types/settlements';
import Spinner from '@/components/ui/spinner';
import {CreateSettlementDialog} from './creationDialog';
import {useQuery} from '@apollo/client';
import {gql} from '@/__generated__';

const GET_SETTLEMENTS = gql(/* GraphQL */ `
  query GetSettlements {
    settlements {
      id
      name
    }
  }
`);

export default function SettlementsPage() {
  const {loading, error, refetch, data} = useQuery(GET_SETTLEMENTS);

  if (loading) {
    return <Spinner />;
  }

  if (error) {
    throw new Error(error.message);
  }

  const settlements = data?.settlements as Array<SettlementId>;
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
          <CreateSettlementDialog refresh={refetch} />
        </li>
      </ul>
    </main>
  );
}
