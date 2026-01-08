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
import { useRequireAuth } from "@/lib/use-auth";

interface LabResult {
  id: string;
  patientId: string;
  doctorId: string;
  testType: string;
  testDate: string;
  results: string;
  normalRange: string;
  status: "pending" | "completed" | "reviewed";
  notes?: string;
  createdAt: string;
}

interface LabResultsResponse {
  data: LabResult[];
  total: number;
  page: number;
  limit: number;
}

const statusConfig = {
  pending: { label: "Pendente", variant: "outline" as const, icon: Minus },
  completed: {
    label: "Concluído",
    variant: "default" as const,
    icon: TrendingUp,
  },
  reviewed: {
    label: "Revisado",
    variant: "stable" as const,
    icon: TrendingDown,
  },
};

export default function LabResultsPage() {
  useRequireAuth();

  const [sorting, setSorting] = useState<SortingState>([]);
  const [globalFilter, setGlobalFilter] = useState("");
  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 10,
  });

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
      accessorKey: "testType",
      header: "Tipo de Exame",
      cell: ({ row }) => (
        <div className="font-medium">{row.getValue("testType")}</div>
      ),
    },
    {
      accessorKey: "testDate",
      header: "Data do Exame",
      cell: ({ row }) => {
        const date = new Date(row.getValue("testDate"));
        return (
          <div className="text-sm text-muted-foreground">
            {format(date, "dd/MM/yyyy", { locale: ptBR })}
          </div>
        );
      },
    },
    {
      accessorKey: "results",
      header: "Resultado",
      cell: ({ row }) => (
        <div className="max-w-xs truncate text-sm">
          {row.getValue("results")}
        </div>
      ),
    },
    {
      accessorKey: "normalRange",
      header: "Valores de Referência",
      cell: ({ row }) => (
        <Badge variant="outline">{row.getValue("normalRange")}</Badge>
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
      <div className="mx-auto max-w-7xl">
        {/* Header */}
        <motion.div
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          className="mb-8 flex items-center justify-between"
        >
          <div>
            <h1 className="text-3xl font-bold tracking-tight flex items-center gap-3">
              <TestTube className="h-8 w-8 text-purple-600" />
              Exames Laboratoriais
            </h1>
            <p className="mt-2 text-muted-foreground">
              Gerencie resultados de exames laboratoriais
            </p>
          </div>
          <Button className="gap-2">
            <Plus className="h-4 w-4" />
            Novo Resultado
          </Button>
        </motion.div>

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
      </div>
    </div>
  );
}
