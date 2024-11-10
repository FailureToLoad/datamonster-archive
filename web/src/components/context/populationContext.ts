import {Survivor, DefaultSurvivor} from '@/lib/types/survivor';
import {createContext} from 'react';

export interface PopulationContextType {
  currentSurvivor: Survivor;
  setCurrentSurvivor: (survivor: Survivor) => void;
  dialogOpen: boolean;
  setDialogOpen: (open: boolean) => void;
  edit: boolean;
  setEdit: (isReadOnly: boolean) => void;
}
export const PopulationContext = createContext<PopulationContextType>({
  currentSurvivor: DefaultSurvivor,
  setCurrentSurvivor: function (): void {
    throw new Error('Function not implemented.');
  },
  dialogOpen: false,
  setDialogOpen: function (): void {
    throw new Error('Function not implemented.');
  },
  edit: true,
  setEdit: function (): void {
    throw new Error('Function not implemented.');
  },
});
