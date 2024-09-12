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
import * as VisuallyHidden from '@radix-ui/react-visually-hidden';
import {Checkbox} from '@/components/ui/checkbox';
import {RadioGroup, RadioGroupItem} from '@/components/ui/radio-group';
import {Separator} from '@/components/ui/separator';
import {Input} from '@/components/ui/input';
import {cn} from '@/lib/utils';
import {z} from 'zod';
import {DefaultSurvivor, GET_SURVIVORS} from '@/lib/services/survivor';
import {Control, FieldPath, useForm} from 'react-hook-form';

import {zodResolver} from '@hookform/resolvers/zod';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import {useContext, useEffect, useState} from 'react';
import {PopulationContext} from '@/components/context/populationContext';
import {Toggle} from '@/components/ui/toggle';
import {PencilSimple, Plus} from '@phosphor-icons/react';
import {useMutation} from '@apollo/client';
import {
  CreateSurvivor,
  CreateSurvivorRequest,
  UpdateSurvivor,
} from '@/lib/services/survivor';
import {SurvivorGender} from '@/__generated__/graphql';

const formSchema = z.object({
  name: z
    .string()
    .min(1, {message: 'Name cannot be empty'})
    .max(50, {message: 'Too long.'}),
  gender: z.enum(['M', 'F']),
  survival: z.coerce.number().min(0, {message: 'Cannot be negative'}),
  systemicPressure: z.coerce.number(),
  movement: z.coerce.number(),
  accuracy: z.coerce.number(),
  strength: z.coerce.number(),
  evasion: z.coerce.number(),
  luck: z.coerce.number(),
  speed: z.coerce.number(),
  lumi: z.coerce.number().min(0, {message: 'Cannot be negative.'}),
  insanity: z.coerce.number().min(0, {message: 'Cannot be negative.'}),
  torment: z.coerce.number(),
});

type SurvivorFormFields = z.infer<typeof formSchema>;

type SurvivorDialogProps = {
  settlementId: string;
};

export default function NewSurvivorDialog({settlementId}: SurvivorDialogProps) {
  const [loading, setLoading] = useState(false);
  const {
    currentSurvivor,
    setCurrentSurvivor,
    dialogOpen,
    setDialogOpen,
    edit,
    setEdit,
  } = useContext(PopulationContext);

  const form = useForm<SurvivorFormFields>({
    resolver: zodResolver(formSchema),
    criteriaMode: 'all',
    defaultValues: {
      name: currentSurvivor.name,
      gender: currentSurvivor.gender,
      survival: currentSurvivor.survival,
      systemicPressure: currentSurvivor.systemicpressure,
      movement: currentSurvivor.movement,
      accuracy: currentSurvivor.accuracy,
      strength: currentSurvivor.strength,
      evasion: currentSurvivor.evasion,
      luck: currentSurvivor.luck,
      speed: currentSurvivor.speed,
      lumi: currentSurvivor.lumi,
      insanity: currentSurvivor.insanity,
      torment: currentSurvivor.torment,
    },
  });
  const [createSurvivor] = useMutation(CreateSurvivor);
  const [updateSurvivor] = useMutation(UpdateSurvivor);
  async function onSubmit(values: SurvivorFormFields) {
    try {
      setLoading(true);
      if (!settlementId || settlementId === '') {
        throw Error('settlementId is required');
      }

      formSchema.parse(values);

      const survivor: CreateSurvivorRequest = {
        name: values.name,
        born: 1,
        gender: values.gender as SurvivorGender,
        huntxp: 0,
        survival: values.survival,
        movement: values.movement,
        accuracy: values.accuracy,
        strength: values.strength,
        evasion: values.evasion,
        luck: values.luck,
        speed: values.speed,
        insanity: values.insanity,
        systemicpressure: values.systemicPressure,
        torment: values.torment,
        lumi: values.lumi,
        courage: 0,
        understanding: 0,
        settlementID: settlementId,
      };

      if (currentSurvivor.id.length > 0) {
        updateSurvivor({
          variables: {
            id: currentSurvivor.id,
            input: {
              ...survivor,
            },
          },
          refetchQueries: [GET_SURVIVORS],
        });
      } else {
        createSurvivor({
          variables: {
            input: {
              ...survivor,
            },
          },
          refetchQueries: [GET_SURVIVORS],
        });
      }

      form.reset({});
      setLoading(false);
      setDialogOpen(false);
    } catch (error) {
      console.log(error);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    form.setValue('name', currentSurvivor.name || 'Meat');
    form.setValue('gender', currentSurvivor.gender || 'M');
    form.setValue('survival', currentSurvivor.survival || 0);
    form.setValue('systemicPressure', currentSurvivor.systemicpressure || 0);
    form.setValue('movement', currentSurvivor.movement || 0);
    form.setValue('accuracy', currentSurvivor.accuracy || 0);
    form.setValue('strength', currentSurvivor.strength || 0);
    form.setValue('evasion', currentSurvivor.evasion || 0);
    form.setValue('luck', currentSurvivor.luck || 0);
    form.setValue('speed', currentSurvivor.speed || 0);
    form.setValue('lumi', currentSurvivor.lumi || 0);
    form.setValue('insanity', currentSurvivor.insanity || 0);
    form.setValue('torment', currentSurvivor.torment || 0);
  }, [form, currentSurvivor]);

  const newSurvivorState = () => {
    setEdit(true);
    setCurrentSurvivor(DefaultSurvivor);
  };
  return (
    <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
        <Button variant="outline" onClick={newSurvivorState}>
          <Plus className="size-4" />
        </Button>
      </DialogTrigger>
      <DialogContent className="mx-auto w-3/5 px-6 flex-grow">
        <VisuallyHidden.Root>
          <DialogHeader>
            <DialogTitle>Create Survivor</DialogTitle>
            <DialogDescription>
              Fill in all required fields to create a new survivor.
            </DialogDescription>
          </DialogHeader>
        </VisuallyHidden.Root>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <section className="grid grid-cols-2 items-center justify-center gap-4">
              <div className="mb-4 flex flex-row items-center justify-between col-span-2 h-full border-b-2 border-black">
                <div className="flex flex-row gap-2 w-full">
                  <EditButton value={edit} onChange={setEdit} />
                  <p className="text-2xl font-serif font-light tracking-wide">
                    Name
                  </p>
                  <FormField
                    control={form.control}
                    name="name"
                    render={({field}) => (
                      <FormItem>
                        <FormMessage />
                        <FormControl>
                          <Input
                            type="text"
                            id="name-input"
                            className="w-full text-lg"
                            value={field.value}
                            onChange={field.onChange}
                            disabled={!edit}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
                <FormField
                  control={form.control}
                  name="gender"
                  render={({field}) => (
                    <FormItem>
                      <FormControl>
                        <RadioGroup
                          className="flex flex-row"
                          onValueChange={field.onChange}
                          defaultValue={field.value}
                          disabled={!edit}
                        >
                          <FormItem className="flex items-center space-x-2">
                            <FormControl>
                              <RadioGroupItem value="M" id="gender-m" />
                            </FormControl>
                            <FormLabel>M</FormLabel>
                          </FormItem>
                          <FormItem className="flex items-center space-x-2">
                            <FormControl>
                              <RadioGroupItem value="F" id="r2" />
                            </FormControl>
                            <FormLabel>F</FormLabel>
                          </FormItem>
                        </RadioGroup>
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <div className="flex flex-row items-center col-span-2 border border-black h-full gap-2">
                <FormField
                  control={form.control}
                  name="survival"
                  render={({field}) => (
                    <FormItem>
                      <FormControl>
                        <div className="order-1 ml-6 my-4 border border-black size-20 place-content-around">
                          <Input
                            id="survival-input"
                            type="number"
                            className="size-full text-center text-2xl"
                            value={field.value}
                            onChange={field.onChange}
                            disabled={!edit}
                          />
                        </div>
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <div className="order-2 my-4 flex flex-col flex-1 h-20 items-start justify-between">
                  <p className="text-2xl font-serif font-light tracking-wide">
                    Survival
                  </p>
                  <div className="flex w-full ">
                    <div className="flex items-start space-x-2">
                      <Checkbox id="cannot-spend" />
                      <label
                        htmlFor="cannot-spend"
                        className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                      >
                        Cannot spend survival
                      </label>
                    </div>
                  </div>
                </div>
                <div className="order-3 my-4 items-start justify-start content-start space-y-1">
                  <CheckboxItem id="dodge" label="Dodge" />
                  <CheckboxItem id="encourage" label="Encourage" />
                  <CheckboxItem id="surge" label="Surge" />
                  <CheckboxItem id="dash" label="Dash" />
                  <CheckboxItem id="fistpump" label="Fist Pump" />
                </div>
                <div className="order-last border-l border-black justify-self-end flex flex-col">
                  <Stat
                    id="systemicPressure"
                    label={`Systemic Pressure`}
                    control={form.control}
                    className="mx-4 size-min"
                    disabled={!edit}
                  />
                </div>
              </div>
              <div className="flex flex-row items-center justify-between col-span-2 border border-black h-32 gap-2">
                <Stat
                  id="movement"
                  label="Movement"
                  control={form.control}
                  className="ml-4 mr-2 size-full"
                  disabled={!edit}
                />
                <Separator className="bg-black" orientation="vertical" />
                <Stat
                  id="accuracy"
                  label="Accuracy"
                  control={form.control}
                  className="mx-2 size-full"
                  disabled={!edit}
                />
                <Separator className="bg-black" orientation="vertical" />
                <Stat
                  id="strength"
                  label="Strength"
                  control={form.control}
                  className="mx-2 size-full"
                  disabled={!edit}
                />
                <Separator className="bg-black" orientation="vertical" />
                <Stat
                  id="evasion"
                  label="Evasion"
                  control={form.control}
                  className="mx-2 size-full"
                  disabled={!edit}
                />
                <Separator className="bg-black" orientation="vertical" />
                <Stat
                  id="luck"
                  label="Luck"
                  control={form.control}
                  className="mx-2 size-full"
                  disabled={!edit}
                />
                <Separator className="bg-black" orientation="vertical" />
                <Stat
                  id="speed"
                  label="Speed"
                  control={form.control}
                  className="mx-2 size-full"
                  disabled={!edit}
                />
                <Separator className="bg-black" orientation="vertical" />
                <Stat
                  id="lumi"
                  label="Lumi"
                  control={form.control}
                  className="ml-2 mr-4 size-full"
                  disabled={!edit}
                />
              </div>
              <div className="flex flex-row items-center col-span-2 border border-black h-full gap-2">
                <Stat
                  id="insanity"
                  label="Insanity"
                  control={form.control}
                  className="order-first ml-4 mr-2"
                  disabled={!edit}
                />
                <Separator className="bg-black" orientation="vertical" />
                <div className="my-4 flex flex-col flex-1 h-20 items-start justify-between">
                  <div className="w-full flex flex-row justify-between">
                    <p className="text-2xl font-serif font-light tracking-wide">
                      Brain
                    </p>
                    <Checkbox id="brainbox" className="size-6" />
                  </div>

                  <div className="flex w-full ">
                    <p className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                      If your insanity is 3+, you are <b>insane</b>
                    </p>
                  </div>
                </div>
                <Separator className="bg-black" orientation="vertical" />
                <Stat
                  id="torment"
                  label="Torment"
                  control={form.control}
                  className="order-last ml-2 mr-4"
                  disabled={!edit}
                />
              </div>
            </section>
            <DialogFooter className="pt-4 h-9">
              {edit && (
                <Button type="submit" disabled={loading}>
                  Submit
                </Button>
              )}
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}

function EditButton({
  value,
  onChange,
}: {
  value: boolean;
  onChange: (val: boolean) => void;
}) {
  return (
    <Toggle aria-label="Toggle Edit" pressed={value} onPressedChange={onChange}>
      <PencilSimple className="h-4 w-4" />
    </Toggle>
  );
}

function CheckboxItem({id, label}: {id: string; label?: string}) {
  return (
    <div className="flex items-start space-x-1 rounded-none">
      <Checkbox id={id} />
      {label && (
        <label
          htmlFor={id}
          className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
        >
          {label}
        </label>
      )}
    </div>
  );
}

function Stat({
  id,
  label,
  className = '',
  control,
  disabled,
}: {
  id: FieldPath<SurvivorFormFields>;
  label: string;
  className?: string;
  control: Control<SurvivorFormFields>;
  disabled: boolean;
}) {
  return (
    <div
      id={id}
      className={cn(
        'flex flex-col items-center justify-between w-fit',
        className
      )}
    >
      <div className="mt-4 flex border border-black h-16 w-14 ">
        <FormField
          control={control}
          name={id}
          render={({field}) => (
            <FormItem>
              <FormControl>
                <Input
                  id={`${id}-input`}
                  type="number"
                  className="size-full text-center text-lg"
                  disabled={disabled}
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
      </div>
      <p className="flex my-4 text-xs text-wrap whitespace-break-spaces text-center">
        {label}
      </p>
    </div>
  );
}
