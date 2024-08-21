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
import {Settlement} from '@/lib/types/settlements';
import {useState} from 'react';
import {Plus} from 'lucide-react';
import {useMutation} from '@apollo/client';
import {gql} from '@/__generated__';
import {useUser} from '@clerk/clerk-react';

const CREATE_SETTLEMENT = gql(/* GraphQL */ `
  mutation CreateSettlement($input: CreateSettlementInput!) {
    createSettlement(input: $input) {
      id
      name
      owner
    }
  }
`);

const schema = {
  settlementName: z
    .string()
    .min(1, 'Settlement name is too short')
    .max(25, 'Settlement name is too long'),
};
export const AddSettlementSchema = z.object(schema);

export type AddSettlementFields = z.infer<typeof AddSettlementSchema>;

export interface CreateSettlementProps {
  update: (s: Settlement) => void;
}
export function CreateSettlementDialog({refresh}: {refresh: () => void}) {
  const [open, setOpen] = useState(false);
  const {user} = useUser();
  const form = useForm<AddSettlementFields>({
    resolver: zodResolver(AddSettlementSchema),
    defaultValues: {
      settlementName: '',
    },
  });
  const [createSettlement, {loading}] = useMutation(CREATE_SETTLEMENT);
  const submitForm = async (data: AddSettlementFields) => {
    const {settlementName} = AddSettlementSchema.parse({
      settlementName: data.settlementName,
    });

    if (!user) {
      throw new Error('user is required for settlement creation');
    }
    createSettlement({
      variables: {
        input: {
          name: settlementName,
          owner: user.id,
        },
      },
    });
    refresh();
    setOpen(false);
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
              <Button disabled={loading} type="submit">
                Add
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
