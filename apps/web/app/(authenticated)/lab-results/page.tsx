"use client";

import { useState, useMemo } from "react";
import { useQuery } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { motion } from "framer-motion";
import { formatDate } from "@/lib/format-date";
import {
  TestTube,
  Search,
  Plus,
  Edit,
  Eye,
  FileText,
} from "lucide-react";
import { toast } from "sonner";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
} from "@/components/ui/context-menu";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { labResultBatchApi, openLabBatchPDF } from "@/lib/api/lab-result-batch-api";
import { labResultViewApi } from "@/lib/api/lab-result-view-api";
import { useLabResultViewPreference } from "@/hooks/useLabResultViewPreference";
import { UnclassifiedResultsAlert } from "@/components/lab-results/UnclassifiedResultsAlert";

// Tipos para a estrutura pivot
interface PivotCell {
  value: string;
  batchId: string;
  resultId: string;
  unit?: string;
  level?: number; // Nível para coloração
  createdAt: string; // Para ordenação quando múltiplos lotes na mesma data
}

// Cores por nível (mesmas do poster de scores, com transparência ~50%)
const getLevelColor = (level?: number | null): string => {
  if (level == null) return "bg-transparent";

  switch (level) {
    case 0:
      return "bg-red-100/50";
    case 1:
      return "bg-orange-100/50";
    case 2:
      return "bg-yellow-100/50";
    case 3:
      return "bg-blue-100/50";
    case 4:
      return "bg-green-100/50";
    case 5:
      return "bg-emerald-100/50";
    default:
      return "bg-transparent";
  }
};

interface PivotRow {
  testType: string;
  testName: string;
  dates: Map<string, PivotCell>; // key = collectionDate (ISO string)
}

export default function LabResultsPage() {
  useRequireAuth();
  const { selectedPatient } = useRequireSelectedPatient();
  const router = useRouter();
  const { currentOrderBy, setOrderBy } = useLabResultViewPreference();

  const [searchTerm, setSearchTerm] = useState("");

  // Buscar todos os batches com resultados (agora retorna DetailResponse)
  const { data: batches, isLoading } = useQuery({
    queryKey: ["lab-result-batches"],
    queryFn: async () => {
      const result = await labResultBatchApi.list({
        limit: 1000,
        offset: 0,
      });
      return result;
    },
    enabled: !!selectedPatient,
  });

  // Buscar views disponíveis
  const { data: views = [] } = useQuery({
    queryKey: ["lab-result-views"],
    queryFn: () => labResultViewApi.list(false, false),
  });

  // Buscar view selecionada (se houver)
  const { data: selectedView } = useQuery({
    queryKey: ["lab-result-view", currentOrderBy.viewId],
    queryFn: () => labResultViewApi.getById(currentOrderBy.viewId!),
    enabled: currentOrderBy.type === "view" && !!currentOrderBy.viewId,
  });

  // Processar dados para estrutura pivot
  const { pivotData, uniqueDates } = useMemo(() => {
    if (!batches) return { pivotData: [], uniqueDates: [] };

    const rowsMap = new Map<string, PivotRow>();
    const datesSet = new Set<string>();

    // Processar cada batch e seus resultados
    batches.forEach((batch) => {
      if (!batch.labResults || batch.labResults.length === 0) return;

      batch.labResults.forEach((result) => {
        // Se tem link com definição, usar dados da definição
        // Senão, usar dados manuais do resultado
        let key: string;
        let displayName: string;
        let displayType: string;

        if (result.labTestDefinitionId && result.labTestDefinition) {
          // IMPORTANTE: Agrupar pelo ID da definição quando há link
          key = `def::${result.labTestDefinitionId}`;
          displayName = result.labTestDefinition.name;
          displayType = result.labTestDefinition.category || "other";
        } else {
          // Sem link: agrupar pelo nome manual
          key = `manual::${result.testType}::${result.testName}`;
          displayName = result.testName;
          displayType = result.testType;
        }

        if (!rowsMap.has(key)) {
          rowsMap.set(key, {
            testType: displayType,
            testName: displayName,
            dates: new Map(),
          });
        }

        const row = rowsMap.get(key)!;
        // Coluna pela data DO RESULTADO (laudos trazem coleta por exame); fallback p/ a do lote.
        const dateKey = result.collectionDate ?? batch.collectionDate;
        datesSet.add(dateKey);

        // Valor do resultado (numérico ou texto)
        const value =
          result.resultNumeric !== undefined && result.resultNumeric !== null
            ? `${result.resultNumeric}`
            : result.resultText || "-";

        const newCell: PivotCell = {
          value,
          batchId: batch.id,
          resultId: result.id,
          unit: result.unit,
          level: result.level,
          createdAt: batch.createdAt,
        };

        // Se já existe resultado nesta data, manter o mais recente (por createdAt)
        const existingCell = row.dates.get(dateKey);
        if (!existingCell || new Date(newCell.createdAt) > new Date(existingCell.createdAt)) {
          row.dates.set(dateKey, newCell);
        }
      });
    });

    // Ordenar datas (mais recente à esquerda)
    const sortedDates = Array.from(datesSet).sort(
      (a, b) => new Date(b).getTime() - new Date(a).getTime()
    );

    const rows = Array.from(rowsMap.values());

    return { pivotData: rows, uniqueDates: sortedDates };
  }, [batches]);

  // Filtrar por busca
  const filteredData = useMemo(() => {
    if (!searchTerm) return pivotData;
    const term = searchTerm.toLowerCase();
    return pivotData.filter(
      (row) =>
        row.testName.toLowerCase().includes(term) ||
        row.testType.toLowerCase().includes(term)
    );
  }, [pivotData, searchTerm]);

  // Ordenar dados conforme preferência
  const sortedData = useMemo(() => {
    let sorted = [...filteredData];

    if (currentOrderBy.type === "alpha-asc") {
      sorted.sort((a, b) => a.testName.localeCompare(b.testName));
    } else if (currentOrderBy.type === "alpha-desc") {
      sorted.sort((a, b) => b.testName.localeCompare(a.testName));
    } else if (currentOrderBy.type === "view" && selectedView?.items) {
      // Separar: exames no view vs fora do view
      const inView: PivotRow[] = [];
      const notInView: PivotRow[] = [];

      sorted.forEach((row) => {
        const matchingItem = selectedView.items!.find(
          (item) => item.labTestDefinition?.name === row.testName
        );
        if (matchingItem) {
          inView.push(row);
        } else {
          notInView.push(row);
        }
      });

      // Ordenar inView pelo order
      inView.sort((a, b) => {
        const aOrder =
          selectedView.items!.find((i) => i.labTestDefinition?.name === a.testName)
            ?.order || 0;
        const bOrder =
          selectedView.items!.find((i) => i.labTestDefinition?.name === b.testName)
            ?.order || 0;
        return aOrder - bOrder;
      });

      // Ordenar notInView alfabeticamente
      notInView.sort((a, b) => a.testName.localeCompare(b.testName));

      sorted = [...inView, ...notInView];
    }

    return sorted;
  }, [filteredData, currentOrderBy, selectedView]);

  // Handler para editar lote
  const handleEditBatch = (batchId: string) => {
    router.push(`/lab-results/${batchId}/edit`);
  };

  // Abre a tela de detalhe do lote (resultados, status/motivos, Ver PDF)
  const handleViewBatch = (batchId: string) => {
    router.push(`/lab-results/${batchId}`);
  };

  // Abre o PDF original do laudo numa nova aba
  const handleViewPDF = (batchId: string) => {
    openLabBatchPDF(batchId).catch(() =>
      toast.error("Não foi possível abrir o PDF original deste laudo"),
    );
  };


  return (
    <div className="container mx-auto py-8 space-y-8">
      {/* Selected Patient Header */}
      <SelectedPatientHeader />

      {/* Header */}
      <PageHeader
        breadcrumbs={[{ label: "Exames" }]}
        title="Resultados de Exames"
        description={`${pivotData.length} tipos de exames • ${uniqueDates.length} datas`}
        actions={[
          {
            label: "Novo Lote",
            icon: <Plus className="h-4 w-4" />,
            onClick: () => router.push("/lab-results/new"),
            variant: "default" as const,
          },
        ]}
      />

      {/* Alert de resultados não classificados */}
      {batches && <UnclassifiedResultsAlert batches={batches} />}

      {/* Tabela Pivot */}
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.1 }}
      >
        <Card className="border-0 shadow-lg">
          <CardHeader>
            {/* No celular vira coluna: em linha, o select de 200px + a busca de 256px não
                cabem nos 393px do iPhone e a tela ganhava rolagem horizontal (a busca
                ficava fora do enquadramento). */}
            <CardTitle className="flex flex-col items-stretch gap-3 sm:flex-row sm:items-center sm:justify-between sm:gap-4">
              <span>Histórico de Exames</span>
              <div className="flex flex-col gap-2 sm:flex-row sm:items-center">
                <Select
                  value={
                    currentOrderBy.type === "view"
                      ? `view:${currentOrderBy.viewId}`
                      : currentOrderBy.type
                  }
                  onValueChange={(value) => {
                    if (value.startsWith("view:")) {
                      setOrderBy({ type: "view", viewId: value.split(":")[1] });
                    } else {
                      setOrderBy({ type: value as any });
                    }
                  }}
                >
                  <SelectTrigger className="w-full sm:w-[200px]">
                    <SelectValue placeholder="Ordenação" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="alpha-asc">A → Z</SelectItem>
                    <SelectItem value="alpha-desc">Z → A</SelectItem>
                    {views.map((view) => (
                      <SelectItem key={view.id} value={`view:${view.id}`}>
                        {view.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                <div className="relative w-full sm:w-64">
                  <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
                  <Input
                    placeholder="Buscar exame..."
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    className="pl-9"
                  />
                </div>
              </div>
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="rounded-md border overflow-x-auto">
              <table className="w-full text-sm">
                <thead className="bg-muted/50 sticky top-0">
                  <tr>
                    <th className="px-4 py-3 text-left font-semibold w-48 sticky left-0 bg-muted/50 z-10">
                      Exame
                    </th>
                    {uniqueDates.map((date) => (
                      <th
                        key={date}
                        className="px-3 py-3 text-center font-semibold min-w-[120px]"
                      >
                        <div className="text-xs">
                          {formatDate(date, "dd/MM/yyyy")}
                        </div>
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody className="divide-y">
                  {isLoading ? (
                    Array.from({ length: 5 }).map((_, i) => (
                      <tr key={i}>
                        <td className="px-4 py-3">
                          <Skeleton className="h-6 w-full" />
                        </td>
                        {uniqueDates.map((_, j) => (
                          <td key={j} className="px-3 py-3">
                            <Skeleton className="h-6 w-full" />
                          </td>
                        ))}
                      </tr>
                    ))
                  ) : sortedData.length > 0 ? (
                    sortedData.map((row, rowIndex) => (
                      <tr
                        key={`${row.testType}-${row.testName}`}
                        className="hover:bg-muted/30"
                      >
                        <td className="px-4 py-3 font-medium sticky left-0 bg-background">
                          <div className="font-semibold">{row.testName}</div>
                        </td>
                        {uniqueDates.map((date) => {
                          const cell = row.dates.get(date);
                          return (
                            <td key={date} className="px-3 py-3">
                              {cell ? (
                                <ContextMenu>
                                  <ContextMenuTrigger asChild>
                                    <div
                                      onClick={() => handleViewBatch(cell.batchId)}
                                      title="Clique para ver o lote · botão direito para mais ações"
                                      className={`cursor-pointer p-2 rounded hover:opacity-80 text-center transition-opacity ${getLevelColor(
                                        cell.level
                                      )}`}
                                    >
                                      <span className="font-medium">
                                        {cell.value}
                                        {cell.unit && (
                                          <span className="text-xs text-muted-foreground ml-1">
                                            {cell.unit}
                                          </span>
                                        )}
                                      </span>
                                    </div>
                                  </ContextMenuTrigger>
                                  <ContextMenuContent>
                                    <ContextMenuItem
                                      onClick={() => handleViewBatch(cell.batchId)}
                                    >
                                      <Eye className="h-4 w-4 mr-2" />
                                      Ver detalhes do lote
                                    </ContextMenuItem>
                                    <ContextMenuItem
                                      onClick={() => handleViewPDF(cell.batchId)}
                                    >
                                      <FileText className="h-4 w-4 mr-2" />
                                      Ver PDF do laudo
                                    </ContextMenuItem>
                                    <ContextMenuItem
                                      onClick={() =>
                                        handleEditBatch(cell.batchId)
                                      }
                                    >
                                      <Edit className="h-4 w-4 mr-2" />
                                      Editar Lote
                                    </ContextMenuItem>
                                  </ContextMenuContent>
                                </ContextMenu>
                              ) : (
                                <div className="text-center text-muted-foreground">
                                  -
                                </div>
                              )}
                            </td>
                          );
                        })}
                      </tr>
                    ))
                  ) : (
                    <tr>
                      <td
                        colSpan={uniqueDates.length + 1}
                        className="px-4 py-8 text-center text-muted-foreground"
                      >
                        Nenhum resultado encontrado.
                      </td>
                    </tr>
                  )}
                </tbody>
              </table>
            </div>
          </CardContent>
        </Card>
      </motion.div>
    </div>
  );
}
