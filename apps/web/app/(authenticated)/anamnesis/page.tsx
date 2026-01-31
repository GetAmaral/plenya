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
  ClipboardList,
  Search,
  Plus,
  ChevronLeft,
  ChevronRight,
  Heart,
  Activity,
  XCircle,
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

interface Anamnesis {
  id: string;
  patientId: string;
  doctorId: string;
  chiefComplaint: string;
  historyOfPresentIllness: string;
  pastMedicalHistory: string;
  familyHistory: string;
  allergies: string;
  medications: string;
  socialHistory: string;
  reviewOfSystems: string;
  createdAt: string;
  updatedAt: string;
}

interface AnamnesisResponse {
  data: Anamnesis[];
  total: number;
  page: number;
  limit: number;
}

export default function AnamnesisPage() {
  useRequireAuth();
  useRequireSelectedPatient();

  const [sorting, setSorting] = useState<SortingState>([]);
  const [globalFilter, setGlobalFilter] = useState("");
  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 10,
  });

  const { data, isLoading, error } = useQuery({
    queryKey: ["anamnesis", pagination.pageIndex, pagination.pageSize],
    queryFn: async () => {
      const offset = pagination.pageIndex * pagination.pageSize;
      const result = await apiClient.get<Anamnesis[] | AnamnesisResponse>(
        `/api/v1/anamnesis?limit=${pagination.pageSize}&offset=${offset}`
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

  const columns: ColumnDef<Anamnesis>[] = [
    {
      accessorKey: "chiefComplaint",
      header: "Queixa Principal",
      cell: ({ row }) => (
        <div className="max-w-md truncate font-medium">
          {row.getValue("chiefComplaint")}
        </div>
      ),
    },
    {
      accessorKey: "allergies",
      header: "Alergias",
      cell: ({ row }) => {
        const allergies = row.getValue("allergies") as string;
        return allergies ? (
          <Badge variant="outline" className="gap-1">
            <Heart className="h-3 w-3" />
            {allergies.split(",").length} alergia(s)
          </Badge>
        ) : (
          <span className="text-sm text-muted-foreground">Nenhuma</span>
        );
      },
    },
    {
      accessorKey: "medications",
      header: "Medicamentos",
      cell: ({ row }) => {
        const medications = row.getValue("medications") as string;
        return medications ? (
          <div className="max-w-xs truncate text-sm">
            {medications}
          </div>
        ) : (
          <span className="text-sm text-muted-foreground">Nenhum</span>
        );
      },
    },
    {
      accessorKey: "createdAt",
      header: "Data",
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
      accessorKey: "updatedAt",
      header: "Atualizado",
      cell: ({ row }) => {
        const date = new Date(row.getValue("updatedAt"));
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
        const anamnesis = row.original;

        return (
          <div className="flex gap-2">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => console.log("Ver", anamnesis.id)}
            >
              Ver
            </Button>
            <Button
              variant="ghost"
              size="sm"
              onClick={() => console.log("Editar", anamnesis.id)}
            >
              Editar
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

  const stats = [
    {
      label: "Total de Anamneses",
      value: data?.total || 0,
      icon: ClipboardList,
    },
    {
      label: "Com Alergias",
      value:
        data?.data?.filter((a) => a.allergies && a.allergies.trim() !== "")
          .length || 0,
      icon: Heart,
    },
    {
      label: "Em Medicação",
      value:
        data?.data?.filter((a) => a.medications && a.medications.trim() !== "")
          .length || 0,
      icon: Activity,
    },
  ];

  return (
    <div className="container mx-auto py-8 space-y-8">
      {/* Selected Patient Header */}
      <SelectedPatientHeader />

      {/* Header */}
      <PageHeader
        breadcrumbs={[{ label: 'Anamnese' }]}
        title="Anamnese"
        description={`${data?.total || 0} registros de anamnese`}
        actions={[
          {
            label: 'Novo',
            icon: <Plus className="h-4 w-4" />,
            onClick: () => console.log('Nova anamnese'),
            variant: 'default',
          },
        ]}
      />

        {/* Stats */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
          className="mb-6 grid gap-4 md:grid-cols-3"
        >
          {stats.map((stat) => {
            const Icon = stat.icon;
            return (
              <Card key={stat.label} className="border-0 shadow-sm">
                <CardContent className="p-4">
                  <div className="flex items-center justify-between">
                    <div>
                      <p className="text-sm text-muted-foreground">
                        {stat.label}
                      </p>
                      <p className="text-2xl font-bold">{stat.value}</p>
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
                <span>Lista de Anamneses</span>
                <div className="relative w-64">
                  <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                  <Input
                    placeholder="Buscar anamneses..."
                    value={globalFilter ?? ""}
                    onChange={(e) => setGlobalFilter(e.target.value)}
                    className="pl-9"
                  />
                </div>
              </CardTitle>
            </CardHeader>
            <CardContent>
              {error ? (
                <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-8 text-center">
                  <XCircle className="mx-auto h-12 w-12 text-destructive mb-4" />
                  <h3 className="text-lg font-semibold mb-2">Erro ao carregar anamneses</h3>
                  <p className="text-sm text-muted-foreground mb-4">
                    {error instanceof Error ? error.message : "Não foi possível conectar ao servidor"}
                  </p>
                  <Button
                    variant="outline"
                    onClick={() => window.location.reload()}
                  >
                    Tentar novamente
                  </Button>
                </div>
              ) : (
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
                            Nenhuma anamnese encontrada.
                          </td>
                        </tr>
                      )}
                    </tbody>
                  </table>
                </div>
              )}

              {/* Pagination */}
              <div className="mt-4 flex items-center justify-between">
                <div className="text-sm text-muted-foreground">
                  Mostrando{" "}
                  {pagination.pageIndex * pagination.pageSize + 1} até{" "}
                  {Math.min(
                    (pagination.pageIndex + 1) * pagination.pageSize,
                    data?.total || 0
                  )}{" "}
                  de {data?.total || 0} anamneses
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
  );
}
