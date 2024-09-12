import {gql} from '@/__generated__/';
import {PopulationContext} from '@/components/context/populationContext';
import {Button} from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {DefaultSurvivor, GET_SURVIVORS} from '@/lib/services/survivor';
import {useMutation} from '@apollo/client';
import {useContext} from 'react';

const DELETE_SURVIVOR = gql(`
  mutation DeleteSurvivor($id: ID!) {
    deleteSurvivor(id: $id)
  }
`);

interface DeleteDialogProps {
  open: boolean;
  setOpen: (isOpen: boolean) => void;
}

export default function DeleteDialog({open, setOpen}: DeleteDialogProps) {
  const {currentSurvivor, setCurrentSurvivor} =
    useContext(PopulationContext);
  const [deleteSurvivor, {error}] = useMutation(DELETE_SURVIVOR, {
    refetchQueries: [GET_SURVIVORS],
    awaitRefetchQueries: true,
    variables: {id: currentSurvivor.id},
  });
  async function handleDelete() {
    await deleteSurvivor();
    if (error) {
      console.log(error);
    }
    setCurrentSurvivor(DefaultSurvivor);
    setOpen(false);
  }
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="mx-auto w-3/5 px-6 flex-grow">
        <DialogHeader>
          <DialogTitle>Delete Survivor </DialogTitle>
          <DialogDescription>
            This is a destructive action with no means of recovery. Please
            confirm you wish to permanently delete this survivor.
          </DialogDescription>
        </DialogHeader>
        <div className="w-full flex text-center justify-center font-bold">
          {currentSurvivor.name}
        </div>
        <Button variant="destructive" onClick={() => handleDelete()}>
          Delete Forever
        </Button>
      </DialogContent>
    </Dialog>
  );
}
