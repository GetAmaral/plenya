"use client";

import { useState, useMemo, useCallback } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
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
import {
  TestTube2,
  Search,
  Plus,
  ChevronLeft,
  ChevronRight,
  Edit,
  Trash2,
  Network,
  ChevronDown,
  ChevronRight as ChevronRightIcon,
} from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import {
  labTestDefinitionsApi,
  labTestHelpers,
  type LabTestDefinition,
  type LabTestCategory,
} from "@/lib/api/lab-test-api";
import { LabTestDefinitionDialog } from "@/components/lab-tests/LabTestDefinitionDialog";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { toast } from "sonner";
import { PageHeader } from "@/components/layout/page-header";

export default function LabTestDefinitionsPage() {
  useRequireAuth();
  const queryClient = useQueryClient();

  const [sorting, setSorting] = useState<SortingState>([]);
  const [globalFilter, setGlobalFilter] = useState("");
  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 20,
  });
  const [deleteDialogOpen, setDeleteDialogOpen] = useState(false);
  const [testToDelete, setTestToDelete] = useState<string | null>(null);
  const [expandedRows, setExpandedRows] = useState<Set<string>>(new Set());
  const [createDialogOpen, setCreateDialogOpen] = useState(false);
  const [editDialogOpen, setEditDialogOpen] = useState(false);
  const [testToEdit, setTestToEdit] = useState<LabTestDefinition | null>(null);

  const { data: labTests, isLoading } = useQuery({
    queryKey: ["lab-test-definitions"],
    queryFn: labTestDefinitionsApi.getAll,
  });

  const deleteMutation = useMutation({
    mutationFn: labTestDefinitionsApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["lab-test-definitions"] });
      toast.success("Definição removida", {
        description: "A definição de exame foi removida com sucesso.",
      });
      setDeleteDialogOpen(false);
      setTestToDelete(null);
    },
    onError: () => {
      toast.error("Erro ao remover", {
        description: "Não foi possível remover a definição de exame.",
      });
    },
  });

  const toggleRow = useCallback((id: string) => {
    setExpandedRows((prev) => {
      const next = new Set(prev);
      if (next.has(id)) {
        next.delete(id);
      } else {
        next.add(id);
      }
      return next;
    });
  }, []);

  const handleDelete = useCallback((id: string) => {
    setTestToDelete(id);
    setDeleteDialogOpen(true);
  }, []);

  const confirmDelete = useCallback(() => {
    if (testToDelete) {
      deleteMutation.mutate(testToDelete);
    }
  }, [testToDelete, deleteMutation]);

  const handleEdit = useCallback((test: LabTestDefinition) => {
    setTestToEdit(test);
    setEditDialogOpen(true);
  }, []);

  const handleCreate = useCallback(() => {
    setCreateDialogOpen(true);
  }, []);

  // Category stats
  const categoryStats = useMemo(() => {
    return labTests?.reduce(
      (acc, test) => {
        acc[test.category] = (acc[test.category] || 0) + 1;
        return acc;
      },
      {} as Record<LabTestCategory, number>
    ) || {};
  }, [labTests]);

  const columns = useMemo<ColumnDef<LabTestDefinition>[]>(() => [
    {
      accessorKey: "name",
      header: "Nome do Exame",
      cell: ({ row }) => {
        const test = row.original;
        const hasSubTests = test.subTests && test.subTests.length > 0;
        const isExpanded = expandedRows.has(test.id);

        return (
          <div className="flex items-center gap-2">
            {hasSubTests && (
              <button
                onClick={() => toggleRow(test.id)}
                className="text-muted-foreground hover:text-foreground"
              >
                {isExpanded ? (
                  <ChevronDown className="h-4 w-4" />
                ) : (
                  <ChevronRightIcon className="h-4 w-4" />
                )}
              </button>
            )}
            <div>
              <div className="font-medium">{test.name}</div>
              {test.shortName && (
                <div className="text-xs text-muted-foreground">
                  {test.shortName}
                </div>
              )}
            </div>
          </div>
        );
      },
    },
    {
      accessorKey: "code",
      header: "Código",
      cell: ({ row }) => (
        <div className="font-mono text-xs">{row.getValue("code")}</div>
      ),
    },
    {
      accessorKey: "category",
      header: "Categoria",
      cell: ({ row }) => {
        const category = row.getValue("category") as LabTestCategory;
        return (
          <Badge variant="outline">
            {labTestHelpers.getCategoryLabel(category)}
          </Badge>
        );
      },
    },
    {
      accessorKey: "isRequestable",
      header: "Solicitável",
      cell: ({ row }) => {
        const isRequestable = row.getValue("isRequestable") as boolean;
        return (
          <Badge variant={isRequestable ? "default" : "secondary"}>
            {isRequestable ? "Sim" : "Não"}
          </Badge>
        );
      },
    },
    {
      accessorKey: "resultType",
      header: "Tipo",
      cell: ({ row }) => (
        <div className="text-sm text-muted-foreground">
          {labTestHelpers.getResultTypeLabel(
            row.getValue("resultType") as "numeric" | "text" | "boolean" | "categorical"
          )}
        </div>
      ),
    },
    {
      accessorKey: "unit",
      header: "Unidade",
      cell: ({ row }) => {
        const unit = row.getValue("unit") as string | undefined;
        return (
          <div className="text-sm text-muted-foreground">{unit || "—"}</div>
        );
      },
    },
    {
      id: "subTests",
      header: "Sub-exames",
      cell: ({ row }) => {
        const subTests = row.original.subTests;
        if (!subTests || subTests.length === 0) return <div>—</div>;
        return (
          <Badge variant="outline" className="gap-1">
            <Network className="h-3 w-3" />
            {subTests.length}
          </Badge>
        );
      },
    },
    {
      id: "actions",
      cell: ({ row }) => {
        const test = row.original;

        return (
          <div className="flex gap-2">
            <Button
              variant="ghost"
              size="sm"
              onClick={() => handleEdit(test)}
            >
              <Edit className="h-4 w-4" />
            </Button>
            <Button
              variant="ghost"
              size="sm"
              onClick={() => handleDelete(test.id)}
              className="text-destructive hover:text-destructive"
            >
              <Trash2 className="h-4 w-4" />
            </Button>
          </div>
        );
      },
    },
  ], [expandedRows, toggleRow, handleEdit, handleDelete]);

  const table = useReactTable({
    data: labTests || [],
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

  if (isLoading) {
    return (
      <div className="min-h-screen p-6">
        <div className="mx-auto max-w-7xl">
          <div className="space-y-4">
            <Skeleton className="h-12 w-full" />
            <Skeleton className="h-32 w-full" />
            <Skeleton className="h-96 w-full" />
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen p-6">
      <div className="mx-auto max-w-7xl">
        {/* Header */}
        <PageHeader
          title="Definições de Exames"
          description="Gerencie catálogo de exames laboratoriais e parâmetros"
          actions={
            <>
              <Button className="gap-2" onClick={handleCreate}>
                <Plus className="h-4 w-4" />
                Nova Definição
              </Button>
            </>
          }
        />

        {/* Stats */}
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.1 }}
          className="mb-6 grid gap-4 md:grid-cols-4"
        >
          <Card className="border-0 shadow-sm">
            <CardContent className="p-4">
              <div>
                <p className="text-sm text-muted-foreground">
                  Total de Exames
                </p>
                <p className="text-2xl font-bold">{labTests?.length || 0}</p>
              </div>
            </CardContent>
          </Card>

          <Card className="border-0 shadow-sm">
            <CardContent className="p-4">
              <div>
                <p className="text-sm text-muted-foreground">Solicitáveis</p>
                <p className="text-2xl font-bold">
                  {labTests?.filter((t) => t.isRequestable).length || 0}
                </p>
              </div>
            </CardContent>
          </Card>

          <Card className="border-0 shadow-sm">
            <CardContent className="p-4">
              <div>
                <p className="text-sm text-muted-foreground">Parâmetros</p>
                <p className="text-2xl font-bold">
                  {labTests?.filter((t) => !t.isRequestable).length || 0}
                </p>
              </div>
            </CardContent>
          </Card>

          <Card className="border-0 shadow-sm">
            <CardContent className="p-4">
              <div>
                <p className="text-sm text-muted-foreground">Categorias</p>
                <p className="text-2xl font-bold">
                  {Object.keys(categoryStats).length}
                </p>
              </div>
            </CardContent>
          </Card>
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
                <span>Catálogo de Exames</span>
                <div className="relative w-64">
                  <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                  <Input
                    placeholder="Buscar exames..."
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
                      Array.from({ length: 10 }).map((_, i) => (
                        <tr key={i}>
                          {Array.from({ length: 8 }).map((_, j) => (
                            <td key={j} className="px-4 py-3">
                              <Skeleton className="h-6 w-full" />
                            </td>
                          ))}
                        </tr>
                      ))
                    ) : table.getRowModel().rows?.length ? (
                      <>
                        {table.getRowModel().rows.map((row) => (
                          <>
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
                            {expandedRows.has(row.original.id) &&
                              row.original.subTests &&
                              row.original.subTests.length > 0 && (
                                <tr>
                                  <td
                                    colSpan={columns.length}
                                    className="bg-muted/30 px-4 py-2"
                                  >
                                    <div className="ml-8 space-y-1">
                                      <p className="text-xs font-semibold text-muted-foreground mb-2">
                                        Parâmetros do exame:
                                      </p>
                                      {row.original.subTests.map((subTest) => (
                                        <div
                                          key={subTest.id}
                                          className="flex items-center gap-4 text-sm"
                                        >
                                          <span className="font-medium">
                                            {subTest.name}
                                          </span>
                                          <Badge
                                            variant="outline"
                                            className="text-xs"
                                          >
                                            {subTest.code}
                                          </Badge>
                                          {subTest.unit && (
                                            <span className="text-muted-foreground">
                                              {subTest.unit}
                                            </span>
                                          )}
                                        </div>
                                      ))}
                                    </div>
                                  </td>
                                </tr>
                              )}
                          </>
                        ))}
                      </>
                    ) : (
                      <tr>
                        <td
                          colSpan={columns.length}
                          className="px-4 py-8 text-center text-muted-foreground"
                        >
                          Nenhuma definição encontrada.
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
                    labTests?.length || 0
                  )}{" "}
                  de {labTests?.length || 0} definições
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

        {/* Delete Confirmation Dialog */}
        <AlertDialog open={deleteDialogOpen} onOpenChange={setDeleteDialogOpen}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Confirmar exclusão</AlertDialogTitle>
              <AlertDialogDescription>
                Tem certeza que deseja remover esta definição de exame? Esta
                ação não pode ser desfeita.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>Cancelar</AlertDialogCancel>
              <AlertDialogAction
                onClick={confirmDelete}
                className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
              >
                Remover
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>

        {/* Create Dialog */}
        <LabTestDefinitionDialog
          mode="create"
          open={createDialogOpen}
          onOpenChange={setCreateDialogOpen}
          availableTests={labTests || []}
        />

        {/* Edit Dialog */}
        {testToEdit && (
          <LabTestDefinitionDialog
            mode="edit"
            open={editDialogOpen}
            onOpenChange={setEditDialogOpen}
            initialData={testToEdit}
            availableTests={labTests || []}
          />
        )}
      </div>
    </div>
  );
}
