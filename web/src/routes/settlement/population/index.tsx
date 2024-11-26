import Spinner from '@/components/ui/spinner';
import {useAuth} from '@clerk/clerk-react';
import {SurvivorTable} from './survivorTable';
import NewSurvivorDialog from './survivorDialog';
import {Survivor} from '@/lib/types/survivor';
import {useParams} from 'react-router-dom';
import {FetchSurvivors} from '@/lib/services/survivor';
import {useQuery, useQueryClient} from '@tanstack/react-query';
import {PopulationQueryKey} from '@/components/context/populationContextProvider';

export default function PopulationTab() {
  const {settlementId} = useParams();
  const queryClient = useQueryClient();
  const {getToken, isLoaded} = useAuth();
  const afterSubmit = () => {
    queryClient.invalidateQueries({queryKey: [PopulationQueryKey]});
  };

  if (!settlementId) {
    throw Error('settlement id is required');
  }
  const getPopulation = async () => {
    try {
      const token = await getToken();
      if (!token) {
        throw Error('must be logged in');
      }
      const response = await FetchSurvivors(settlementId, token);
      if (!response) return null;
      return response;
    } catch (e) {
      console.log(e);
      return null;
    }
  };
  const {isPending, isError, data, error} = useQuery({
    queryKey: [PopulationQueryKey],
    queryFn: getPopulation,
  });

  if (isPending || !isLoaded) {
    return <Spinner />;
  }

  if (isError) {
    throw new Error(error.message);
  }

  const population = data as Array<Survivor>;
  return (
    <div id="population" className="max-w-fit py-4">
      <div className="flex flex-row-reverse items-center py-4">
        <NewSurvivorDialog
          settlementId={settlementId}
          getToken={getToken}
          afterSubmit={afterSubmit}
        />
      </div>
      <SurvivorTable data={population} />
    </div>
  );
}
