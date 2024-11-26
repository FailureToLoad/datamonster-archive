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
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from '@/components/ui/form';
import {Input} from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {SurvivorStatus} from '@/lib/types/survivor';
import {zodResolver} from '@hookform/resolvers/zod';
import {useContext} from 'react';
import {useForm} from 'react-hook-form';
import {z} from 'zod';

const FormSchema = z.object({
  status: z.nativeEnum(SurvivorStatus),
  year: z.coerce.number().min(0).max(30),
});

interface StatusDialogProps {
  open: boolean;
  setOpen: (isOpen: boolean) => void;
}

export default function StatusDialog({open, setOpen}: StatusDialogProps) {
  const {currentSurvivor} = useContext(PopulationContext);
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  function onSubmit(data: z.infer<typeof FormSchema>) {
    // create request
    // call update
    // refetch survivors / invalidate cache
    console.log(data);
    setOpen(false);
  }
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="mx-auto min-w-fit px-6 flex-grow">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
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
                  <FormField
                    control={form.control}
                    name="status"
                    render={({field}) => (
                      <FormItem>
                        <Select
                          onValueChange={field.onChange}
                          defaultValue={field.value}
                        >
                          <FormControl>
                            <SelectTrigger className="border-black border-b">
                              <SelectValue placeholder="Status" />
                            </SelectTrigger>
                          </FormControl>
                          <SelectContent>
                            <SelectGroup>
                              <SelectItem value={SurvivorStatus.Alive}>
                                Alive
                              </SelectItem>
                              <SelectItem value={SurvivorStatus.CeasedToExist}>
                                Ceased to Exist
                              </SelectItem>
                              <SelectItem value={SurvivorStatus.Dead}>
                                Dead
                              </SelectItem>
                              <SelectItem value={SurvivorStatus.Retired}>
                                Retired
                              </SelectItem>
                              <SelectItem value={SurvivorStatus.SkipHunt}>
                                Skip Hunt
                              </SelectItem>
                            </SelectGroup>
                          </SelectContent>
                        </Select>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
              </div>
              <div className="text-start">Year</div>
              <div className="col-span-2 flex flex-col">
                <div className="w-1/2 place-self-start">
                  <FormField
                    control={form.control}
                    name="year"
                    render={({field}) => (
                      <FormItem>
                        <FormControl>
                          <Input
                            type="number"
                            value={field.value}
                            onChange={field.onChange}
                            className="text-start border-black border-b focus:outline-none focus:border-b-2 focus:border-b-accent"
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
              </div>
            </div>
            <DialogFooter>
              <Button variant="link" type="submit">
                Save
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
