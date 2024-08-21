import {Survivor} from '@types';
import {DefaultSurvivor} from '@/lib/services/survivor';
import {ReactNode, useState} from 'react';
import {PopulationContextType, PopulationContext} from './populationContext';

export default function PopulationContextProvider({
  children,
}: {
  children: ReactNode;
}) {
  const [currentSurvivor, setCurrentSurvivor] =
    useState<Survivor>(DefaultSurvivor);
  const [dialogOpen, setDialogOpen] = useState(false);
  const [edit, setEdit] = useState(true);

  const context: PopulationContextType = {
    currentSurvivor,
    setCurrentSurvivor,
    dialogOpen,
    setDialogOpen,
    edit,
    setEdit,
  };
  return (
    <PopulationContext.Provider value={context}>
      {children}
    </PopulationContext.Provider>
  );
}
