import Spinner from '@/components/ui/spinner';
import {useAuth} from '@clerk/clerk-react';
import {SurvivorTable} from './survivorTable';
import NewSurvivorDialog from './survivorDialog';
import {Survivor} from '@types';
import {useParams} from 'react-router-dom';
import {useQuery} from '@apollo/client';
import {GET_SURVIVORS} from '@/lib/services/survivor';

export default function PopulationTab() {
  const {settlementId} = useParams();
  if (!settlementId) {
    throw Error('settlement id is required');
  }
  const {loading, error, data} = useQuery(GET_SURVIVORS, {
    variables: {settlementId},
  });
  const {isLoaded} = useAuth();

  if (loading || !isLoaded) {
    return <Spinner />;
  }

  if (error) {
    throw new Error(error.message);
  }

  const population = data?.survivors as Array<Survivor>;
  return (
    <div id="population" className="max-w-fit py-4">
      <div className="flex flex-row-reverse items-center py-4">
        <NewSurvivorDialog settlementId={settlementId} />
      </div>
      <SurvivorTable data={population} />
    </div>
  );
}
