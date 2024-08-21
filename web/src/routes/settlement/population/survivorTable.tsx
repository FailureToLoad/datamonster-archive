import {
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  useReactTable,
  getSortedRowModel,
  getPaginationRowModel,
  ColumnDef,
} from '@tanstack/react-table';

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import {Button} from '@/components/ui/button';
import {useContext, useState} from 'react';
import {Survivor} from '@types';
import {PopulationContext} from '@/components/context/populationContext';

interface DataTableProps<Survivor> {
  data: Survivor[];
}

export function SurvivorTable<TData extends Survivor>({
  data,
}: DataTableProps<TData>) {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({
    born: false,
    status: false,
    survival: false,
    insanity: false,
    systemicPressure: false,
    torment: false,
    lumi: false,
    courage: false,
    understanding: false,
  });
  const {setDialogOpen, setEdit, setCurrentSurvivor} =
    useContext(PopulationContext);
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    onSortingChange: setSorting,
    getSortedRowModel: getSortedRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    state: {
      sorting,
      columnVisibility,
    },
  });
  const viewSurvivor = (survivor: Survivor) => {
    setCurrentSurvivor(survivor);
    setEdit(false);
    setDialogOpen(true);
  };

  return (
    <div>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>

          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && 'selected'}
                  onClick={() => viewSurvivor(row.original)}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id} className="text-center">
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <Button
          variant="outline"
          size="sm"
          onClick={() => table.previousPage()}
          disabled={!table.getCanPreviousPage()}
        >
          Previous
        </Button>
        <Button
          variant="outline"
          size="sm"
          onClick={() => table.nextPage()}
          disabled={!table.getCanNextPage()}
        >
          Next
        </Button>
      </div>
    </div>
  );
}

enum Keys {
  born = 'born',
  gender = 'gender',
  status = 'status',
  name = 'name',
  xp = 'huntXp',
  survival = 'survival',
  movement = 'movement',
  accuracy = 'accuracy',
  strength = 'strength',
  evasion = 'evasion',
  luck = 'luck',
  speed = 'speed',
  insanity = 'insanity',
  sp = 'systemicPressure',
  torment = 'torment',
  lumi = 'lumi',
  courage = 'courage',
  understanding = 'understanding',
}

const columns: ColumnDef<Survivor>[] = [
  {
    id: Keys.name,
    accessorKey: Keys.name,
    header: 'Name',
  },
  {
    id: Keys.born,
    accessorKey: Keys.born,
    header: 'Born',
  },
  {
    id: Keys.gender,
    accessorKey: Keys.gender,
    header: 'Gender',
  },
  {
    id: Keys.status,
    accessorKey: Keys.status,
    header: 'Status',
  },
  {
    id: Keys.xp,
    accessorKey: Keys.xp,
    header: 'XP',
  },
  {
    id: Keys.survival,
    accessorKey: Keys.survival,
    header: 'Survival',
  },
  {
    id: Keys.movement,
    accessorKey: Keys.movement,
    header: 'Movement',
  },
  {
    id: Keys.accuracy,
    accessorKey: Keys.accuracy,
    header: 'Accuracy',
  },
  {
    id: Keys.strength,
    accessorKey: Keys.strength,
    header: 'Strength',
  },
  {
    id: Keys.evasion,
    accessorKey: Keys.evasion,
    header: 'Evasion',
  },
  {
    id: Keys.luck,
    accessorKey: Keys.luck,
    header: 'Luck',
  },
  {
    id: Keys.speed,
    accessorKey: Keys.speed,
    header: 'Speed',
  },
  {
    id: Keys.insanity,
    accessorKey: Keys.insanity,
    header: 'Insanity',
  },
  {
    id: Keys.sp,
    accessorKey: Keys.sp,
    header: 'Systemic Pressure',
  },
  {
    id: Keys.torment,
    accessorKey: Keys.torment,
    header: 'Torment',
  },
  {
    id: Keys.lumi,
    accessorKey: Keys.lumi,
    header: 'Lumi',
  },
  {
    id: Keys.courage,
    accessorKey: Keys.courage,
    header: 'Courage',
  },
  {
    id: Keys.understanding,
    accessorKey: Keys.understanding,
    header: 'Understanding',
  },
];
