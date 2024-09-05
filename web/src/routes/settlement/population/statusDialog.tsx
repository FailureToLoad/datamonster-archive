import {PopulationContext} from '@/components/context/populationContext';
import {Button} from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {Input} from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {useContext} from 'react';

interface StatusDialogProps {
  open: boolean;
  setOpen: (isOpen: boolean) => void;
}

export default function StatusDialog({open, setOpen}: StatusDialogProps) {
  const {currentSurvivor} = useContext(PopulationContext);
  function setStatus() {
    console.log(currentSurvivor.id);
    setOpen(false);
  }
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="mx-auto min-w-fit px-6 flex-grow">
        <DialogHeader>
          <DialogTitle>Set Status </DialogTitle>
          <DialogDescription>Choose a status.</DialogDescription>
        </DialogHeader>
        <div className="grid grid-cols-3 gap-4 w-full">
          <div className="text-start">Name</div>
          <div className="col-span-2 flex flex-col items-start">
            <div className="w-1/2 text-start border-black border-b">
              {currentSurvivor.name}
            </div>
          </div>
          <div className="text-start">Status</div>
          <div className="col-span-2 flex flex-col">
            <div className="w-1/2 place-self-start">
              <Select>
                <SelectTrigger className="border-black border-b">
                  <SelectValue placeholder="Status" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectItem value="cease-exist">Cease to Exist</SelectItem>
                    <SelectItem value="dead">Dead</SelectItem>
                    <SelectItem value="retired">Retired</SelectItem>
                    <SelectItem value="skip">Skip Hunt</SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
            </div>
          </div>
          <div className="text-start">Year</div>
          <div className="col-span-2 flex flex-col">
            <div className="w-1/2 place-self-start">
              <Input
                type="number"
                className="text-start border-black border-b focus:outline-none focus:border-b-2 focus:border-b-accent"
              />
            </div>
          </div>
        </div>
        <DialogFooter>
          <Button variant="link" onClick={() => setStatus()}>
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
