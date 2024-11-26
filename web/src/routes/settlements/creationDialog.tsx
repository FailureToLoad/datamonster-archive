// src/routes/settlements/creationDialog.tsx
import {Button} from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import {Input} from '@/components/ui/input';
import {
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  Form,
} from '@/components/ui/form';
import * as z from 'zod';
import {useForm} from 'react-hook-form';
import {zodResolver} from '@hookform/resolvers/zod';
import {useState} from 'react';
import {Plus} from 'lucide-react';
import {useQueryClient} from '@tanstack/react-query';
import {CreateSettlement} from '@/lib/services/settlement';
import {SettlementsQueryKey} from '.';
import {useAuth} from '@/auth/hooks';

const schema = {
  settlementName: z
    .string()
    .min(1, 'Settlement name is too short')
    .max(25, 'Settlement name is too long'),
};
export const AddSettlementSchema = z.object(schema);

export type AddSettlementFields = z.infer<typeof AddSettlementSchema>;

export function CreateSettlementDialog() {
  const {getToken} = useAuth();
  const queryClient = useQueryClient();
  const [open, setOpen] = useState(false);
  const form = useForm<AddSettlementFields>({
    resolver: zodResolver(AddSettlementSchema),
    defaultValues: {
      settlementName: '',
    },
  });

  const submitForm = async (data: AddSettlementFields) => {
    try {
      const token = await getToken();
      if (!token) {
        throw new Error('Not authenticated');
      }

      const {settlementName} = AddSettlementSchema.parse({
        settlementName: data.settlementName,
      });

      await CreateSettlement(settlementName, token);
      queryClient.invalidateQueries({queryKey: [SettlementsQueryKey]});
      setOpen(false);
      form.reset();
    } catch (error) {
      console.error('Failed to create settlement:', error);
      // You might want to add error handling UI here
    }
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button
          variant="outline"
          className="w-full"
          aria-label="Create Settlement"
        >
          <Plus className="h-6 w-6" />
        </Button>
      </DialogTrigger>
      <DialogContent
        className="sm:max-w-[425px]"
        data-testid="settlement-modal"
      >
        <Form {...form}>
          <form className="space-y-8" onSubmit={form.handleSubmit(submitForm)}>
            <DialogHeader>
              <DialogTitle>Add Settlement</DialogTitle>
              <DialogDescription>Enter settlement details.</DialogDescription>
            </DialogHeader>

            <FormField
              control={form.control}
              name="settlementName"
              render={({field}) => (
                <FormItem>
                  <FormLabel>Settlement Name</FormLabel>
                  <FormControl>
                    <Input
                      type="text"
                      className="flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-sm shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
                      {...field}
                    />
                  </FormControl>
                </FormItem>
              )}
            />
            <DialogFooter>
              <Button type="submit">Add</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
