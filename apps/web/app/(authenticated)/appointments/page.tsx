"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { motion } from "framer-motion";
import {
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
  type ColumnDef,
  type SortingState,
} from "@tanstack/react-table";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  Calendar,
  Search,
  Plus,
  ChevronLeft,
  ChevronRight,
  Clock,
  CheckCircle2,
  XCircle,
  CalendarCheck,
} from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { apiClient } from "@/lib/api-client";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";

interface Appointment {
  id: string;
  patientId: string;
  doctorId: string;
  scheduledAt: string;
  durationMinutes: number;
  type: string;
  status: "scheduled" | "confirmed" | "completed" | "cancelled";
  reason: string;
  createdAt: string;
}

interface AppointmentsResponse {
  data: Appointment[];
  total: number;
  page: number;
  limit: number;
}

const statusConfig = {
  scheduled: { label: "Agendado", variant: "outline" as const, icon: Clock },
  confirmed: {
    label: "Confirmado",
    variant: "default" as const,
    icon: CalendarCheck,
  },
  completed: { label: "Concluído", variant: "stable" as const, icon: CheckCircle2 },
  cancelled: {
    label: "Cancelado",
    variant: "destructive" as const,
    icon: XCircle,
  },
};

export default function AppointmentsPage() {
  useRequireAuth();
  useRequireSelectedPatient();

  const [sorting, setSorting] = useState<SortingState>([]);
  const [globalFilter, setGlobalFilter] = useState("");
  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 10,
  });

  const { data, isLoading } = useQuery({
    queryKey: ["appointments", pagination.pageIndex, pagination.pageSize],
    queryFn: async () => {
      const offset = pagination.pageIndex * pagination.pageSize;
      const result = await apiClient.get<Appointment[] | AppointmentsResponse>(
        `/api/v1/appointments?limit=${pagination.pageSize}&offset=${offset}`
      );

      // Se a API retornar um array direto, converter para o formato esperado
      if (Array.isArray(result)) {
        return {
          data: result,
          total: result.length,
          page: pagination.pageIndex + 1,
          limit: pagination.pageSize,
        };
      }

      return result;
    },
  });

  const columns: ColumnDef<Appointment>[] = [
    {
      accessorKey: "scheduledAt",
      header: "Data/Hora",
      cell: ({ row }) => {
        const date = new Date(row.getValue("scheduledAt"));
        return (
          <div>
            <div className="font-medium">
              {format(date, "dd/MM/yyyy", { locale: ptBR })}
            </div>
            <div className="text-sm text-muted-foreground">
              {format(date, "HH:mm", { locale: ptBR })}
            </div>
          </div>
        );
      },
    },
    {
      accessorKey: "type",
      header: "Tipo",
      cell: ({ row }) => (
        <Badge variant="outline">{row.getValue("type")}</Badge>
      ),
    },
    {
      accessorKey: "reason",
      header: "Motivo",
      cell: ({ row }) => (
        <div className="max-w-md truncate text-sm">
          {row.getValue("reason")}
        </div>
      ),
    },
    {
      accessorKey: "durationMinutes",
      header: "Duração",
      cell: ({ row }) => (
        <div className="text-sm">{row.getValue("durationMinutes")} min</div>
      ),
    },
    {
      accessorKey: "status",
      header: "Status",
      cell: ({ row }) => {
        const status = row.getValue("status") as keyof typeof statusConfig;
        const config = statusConfig[status];
        const Icon = config.icon;

        return (
          <Badge variant={config.variant} className="gap-1">
            <Icon className="h-3 w-3" />
            {config.label}
          </Badge>
        );
      },
    },
    {
      id: "actions",
      cell: ({ row }) => {
        const appointment = row.original;

        return (
          <div className="flex gap-2">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => console.log("Ver", appointment.id)}
            >
              Ver
            </Button>
            {appointment.status === "scheduled" && (
              <Button
                variant="ghost"
                size="sm"
                onClick={() => console.log("Confirmar", appointment.id)}
              >
                Confirmar
              </Button>
            )}
          </div>
        );
      },
    },
  ];

  const table = useReactTable({
    data: data?.data || [],
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onSortingChange: setSorting,
    onGlobalFilterChange: setGlobalFilter,
    onPaginationChange: setPagination,
    state: {
      sorting,
      globalFilter,
      pagination,
    },
  });

  return (
    <div className="min-h-screen p-6">
      <div className="mx-auto max-w-7xl space-y-8">
        {/* Header */}
        <PageHeader
          breadcrumbs={[{ label: 'Consultas' }]}
          title="Consultas"
          description={`${data?.total || 0} consultas agendadas`}
          actions={[
            {
              label: 'Novo',
              icon: <Plus className="h-4 w-4" />,
              onClick: () => console.log('Nova consulta'),
              variant: 'default',
            },
          ]}
        />

        {/* Selected Patient Header */}
        <SelectedPatientHeader />

        {/* Stats */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
          className="mb-6 grid gap-4 md:grid-cols-4"
        >
          {Object.entries(statusConfig).map(([key, config]) => {
            const Icon = config.icon;
            return (
              <Card key={key} className="border-0 shadow-sm">
                <CardContent className="p-4">
                  <div className="flex items-center justify-between">
                    <div>
                      <p className="text-sm text-muted-foreground">
                        {config.label}
                      </p>
                      <p className="text-2xl font-bold">
                        {data?.data?.filter((a) => a.status === key).length || 0}
                      </p>
                    </div>
                    <Icon className="h-8 w-8 text-muted-foreground" />
                  </div>
                </CardContent>
              </Card>
            );
          })}
        </motion.div>

        {/* Table */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.2 }}
        >
          <Card className="border-0 shadow-lg">
            <CardHeader>
              <CardTitle className="flex items-center justify-between">
                <span>Lista de Consultas</span>
                <div className="relative w-64">
                  <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                  <Input
                    placeholder="Buscar consultas..."
                    value={globalFilter ?? ""}
                    onChange={(e) => setGlobalFilter(e.target.value)}
                    className="pl-9"
                  />
                </div>
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="rounded-md border">
                <table className="w-full">
                  <thead className="bg-muted/50">
                    {table.getHeaderGroups().map((headerGroup) => (
                      <tr key={headerGroup.id}>
                        {headerGroup.headers.map((header) => (
                          <th
                            key={header.id}
                            className="px-4 py-3 text-left text-sm font-semibold"
                          >
                            {header.isPlaceholder
                              ? null
                              : flexRender(
                                  header.column.columnDef.header,
                                  header.getContext()
                                )}
                          </th>
                        ))}
                      </tr>
                    ))}
                  </thead>
                  <tbody className="divide-y">
                    {isLoading ? (
                      Array.from({ length: 5 }).map((_, i) => (
                        <tr key={i}>
                          {Array.from({ length: 6 }).map((_, j) => (
                            <td key={j} className="px-4 py-3">
                              <Skeleton className="h-6 w-full" />
                            </td>
                          ))}
                        </tr>
                      ))
                    ) : table.getRowModel().rows?.length ? (
                      table.getRowModel().rows.map((row) => (
                        <tr
                          key={row.id}
                          className="transition-colors hover:bg-muted/50"
                        >
                          {row.getVisibleCells().map((cell) => (
                            <td key={cell.id} className="px-4 py-3">
                              {flexRender(
                                cell.column.columnDef.cell,
                                cell.getContext()
                              )}
                            </td>
                          ))}
                        </tr>
                      ))
                    ) : (
                      <tr>
                        <td
                          colSpan={columns.length}
                          className="px-4 py-8 text-center text-muted-foreground"
                        >
                          Nenhuma consulta encontrada.
                        </td>
                      </tr>
                    )}
                  </tbody>
                </table>
              </div>

              {/* Pagination */}
              <div className="mt-4 flex items-center justify-between">
                <div className="text-sm text-muted-foreground">
                  Mostrando{" "}
                  {pagination.pageIndex * pagination.pageSize + 1} até{" "}
                  {Math.min(
                    (pagination.pageIndex + 1) * pagination.pageSize,
                    data?.total || 0
                  )}{" "}
                  de {data?.total || 0} consultas
                </div>
                <div className="flex gap-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => table.previousPage()}
                    disabled={!table.getCanPreviousPage()}
                  >
                    <ChevronLeft className="h-4 w-4" />
                  </Button>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => table.nextPage()}
                    disabled={!table.getCanNextPage()}
                  >
                    <ChevronRight className="h-4 w-4" />
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
        </motion.div>
      </div>
    </div>
  );
}
