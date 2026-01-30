"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
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
  TestTube,
  Search,
  Plus,
  ChevronLeft,
  ChevronRight,
  Download,
  TrendingUp,
  TrendingDown,
  Minus,
} from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { apiClient } from "@/lib/api-client";
import { useRequireAuth, useAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import Link from "next/link";
import { Settings } from "lucide-react";
import { LabResultEntryDialog } from "@/components/lab-tests/LabResultEntryDialog";
import { PageHeader } from "@/components/layout/page-header";

interface LabResult {
  id: string;
  patientId: string;
  doctorId: string;
  testName: string;
  testType: string;
  orderedAt: string;
  performedAt?: string;
  status: "pending" | "in_progress" | "completed" | "cancelled";
  result?: string;
  unit?: string;
  referenceRange?: string;
  interpretation?: string;
  createdAt: string;
  updatedAt: string;
}

interface LabResultsResponse {
  data: LabResult[];
  total: number;
  page: number;
  limit: number;
}

const statusConfig = {
  pending: { label: "Pendente", variant: "outline" as const, icon: Minus },
  in_progress: {
    label: "Em Andamento",
    variant: "secondary" as const,
    icon: TrendingUp,
  },
  completed: {
    label: "Concluído",
    variant: "default" as const,
    icon: TrendingUp,
  },
  cancelled: {
    label: "Cancelado",
    variant: "destructive" as const,
    icon: Minus,
  },
};

export default function LabResultsPage() {
  useRequireAuth();
  useRequireSelectedPatient();
  const { user } = useAuth();
  const router = useRouter();

  const [sorting, setSorting] = useState<SortingState>([]);
  const [globalFilter, setGlobalFilter] = useState("");
  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 10,
  });
  const [entryDialogOpen, setEntryDialogOpen] = useState(false);

  const { data, isLoading } = useQuery({
    queryKey: ["lab-results", pagination.pageIndex, pagination.pageSize],
    queryFn: async () => {
      const offset = pagination.pageIndex * pagination.pageSize;
      const result = await apiClient.get<LabResult[] | LabResultsResponse>(
        `/api/v1/lab-results?limit=${pagination.pageSize}&offset=${offset}`
      );

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

  const columns: ColumnDef<LabResult>[] = [
    {
      accessorKey: "testName",
      header: "Nome do Exame",
      cell: ({ row }) => (
        <div className="font-medium">{row.getValue("testName")}</div>
      ),
    },
    {
      accessorKey: "testType",
      header: "Tipo",
      cell: ({ row }) => (
        <div className="text-sm text-muted-foreground">
          {row.getValue("testType")}
        </div>
      ),
    },
    {
      accessorKey: "orderedAt",
      header: "Data da Solicitação",
      cell: ({ row }) => {
        const dateStr = row.getValue("orderedAt") as string;
        if (!dateStr) return <div className="text-sm text-muted-foreground">-</div>;
        try {
          const date = new Date(dateStr);
          return (
            <div className="text-sm text-muted-foreground">
              {format(date, "dd/MM/yyyy", { locale: ptBR })}
            </div>
          );
        } catch {
          return <div className="text-sm text-muted-foreground">-</div>;
        }
      },
    },
    {
      accessorKey: "result",
      header: "Resultado",
      cell: ({ row }) => {
        const result = row.getValue("result") as string | undefined;
        const unit = row.original.unit;
        return (
          <div className="max-w-xs truncate text-sm">
            {result ? `${result}${unit ? ` ${unit}` : ""}` : "-"}
          </div>
        );
      },
    },
    {
      accessorKey: "referenceRange",
      header: "Valores de Referência",
      cell: ({ row }) => {
        const range = row.getValue("referenceRange") as string | undefined;
        return range ? (
          <Badge variant="outline">{range}</Badge>
        ) : (
          <span className="text-muted-foreground">-</span>
        );
      },
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
      accessorKey: "createdAt",
      header: "Registrado em",
      cell: ({ row }) => {
        const date = new Date(row.getValue("createdAt"));
        return (
          <div className="text-sm text-muted-foreground">
            {format(date, "dd/MM/yyyy", { locale: ptBR })}
          </div>
        );
      },
    },
    {
      id: "actions",
      cell: ({ row }) => {
        const labResult = row.original;

        return (
          <div className="flex gap-2">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => console.log("Ver", labResult.id)}
            >
              Ver
            </Button>
            <Button
              variant="ghost"
              size="sm"
              onClick={() => console.log("Download", labResult.id)}
            >
              <Download className="h-4 w-4" />
            </Button>
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
          breadcrumbs={[{ label: 'Exames' }]}
          title="Exames Laboratoriais"
          description={`${data?.total || 0} resultados de exames`}
          actions={[
            ...(user?.role === "admin" ? [{
              label: 'Definições',
              icon: <Settings className="h-4 w-4" />,
              onClick: () => router.push('/lab-results/definitions'),
              variant: 'outline' as const,
            }] : []),
            {
              label: 'Novo',
              icon: <Plus className="h-4 w-4" />,
              onClick: () => setEntryDialogOpen(true),
              variant: 'default' as const,
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
          className="mb-6 grid gap-4 md:grid-cols-3"
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
                        {data?.data?.filter((r) => r.status === key).length || 0}
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
                <span>Lista de Resultados</span>
                <div className="relative w-64">
                  <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                  <Input
                    placeholder="Buscar resultados..."
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
                          {Array.from({ length: 7 }).map((_, j) => (
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
                          Nenhum resultado encontrado.
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
                  de {data?.total || 0} resultados
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

        {/* Lab Result Entry Dialog */}
        <LabResultEntryDialog
          open={entryDialogOpen}
          onOpenChange={setEntryDialogOpen}
          currentUserId={user?.id || ""}
        />
      </div>
    </div>
  );
}
