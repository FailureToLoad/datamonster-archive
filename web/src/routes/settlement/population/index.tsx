import Spinner from '@/components/ui/spinner';
import {useAuth} from '@clerk/clerk-react';

import {useParams} from 'react-router-dom';
import {SurvivorTable} from './survivorTable';
import NewSurvivorDialog from './survivorDialog';
import {Survivor} from '@types';
import {useQuery} from '@apollo/client';
import {gql} from '@/__generated__';

const GET_SURVIVORS = gql(/* GraphQL */ `
  query GetSurvivors($settlementId: ID!) {
    survivors(filter: {settlementID: $settlementId}) {
      id
      accuracy
      born
      courage
      evasion
      gender
      huntxp
      insanity
      luck
      lumi
      movement
      name
      speed
      strength
      survival
      systemicpressure
      torment
      understanding
    }
  }
`);

export const PopulationQueryKey = 'population';
export default function PopulationTab() {
  const {settlementId} = useParams();
  if (!settlementId) {
    throw Error('settlement id is required');
  }
  const {loading, error, refetch, data} = useQuery(GET_SURVIVORS, {
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
        <NewSurvivorDialog settlementId={settlementId} afterSubmit={refetch} />
      </div>
      <SurvivorTable data={population} />
    </div>
  );
}
