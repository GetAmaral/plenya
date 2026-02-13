"use client";

import { useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AlertCircle, RefreshCw } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import { labResultBatchApi } from "@/lib/api/lab-result-batch-api";
import { toast } from "sonner";

interface UnclassifiedResultsAlertProps {
  batches: Array<{
    id: string;
    laboratoryName: string;
    collectionDate: string;
    labResults: Array<{
      id: string;
      testName: string;
      resultNumeric?: number | null;
      labTestDefinitionId?: string | null;
      level?: number | null;
    }>;
  }>;
}

export function UnclassifiedResultsAlert({
  batches,
}: UnclassifiedResultsAlertProps) {
  const queryClient = useQueryClient();
  const [classifyingBatchId, setClassifyingBatchId] = useState<string | null>(
    null
  );

  const unclassifiedStats = batches.reduce(
    (acc, batch) => {
      const unclassifiedInBatch = batch.labResults.filter(
        (result) =>
          result.resultNumeric != null &&
          result.labTestDefinitionId != null &&
          result.level == null
      );

      if (unclassifiedInBatch.length > 0) {
        acc.batchesWithUnclassified.push({
          batchId: batch.id,
          laboratoryName: batch.laboratoryName,
          collectionDate: batch.collectionDate,
          unclassifiedCount: unclassifiedInBatch.length,
        });
        acc.totalUnclassified += unclassifiedInBatch.length;
      }

      return acc;
    },
    {
      batchesWithUnclassified: [] as Array<{
        batchId: string;
        laboratoryName: string;
        collectionDate: string;
        unclassifiedCount: number;
      }>,
      totalUnclassified: 0,
    }
  );

  const classifyMutation = useMutation({
    mutationFn: async (batchId: string) => {
      return labResultBatchApi.classify(batchId);
    },
    onSuccess: (_, batchId) => {
      toast.success("Lote classificado com sucesso");
      queryClient.invalidateQueries({ queryKey: ["lab-result-batches"] });
      setClassifyingBatchId(null);
    },
    onError: (error: any) => {
      toast.error("Erro ao classificar lote", {
        description: error.message || "Tente novamente",
      });
      setClassifyingBatchId(null);
    },
  });

  const handleClassify = (batchId: string) => {
    setClassifyingBatchId(batchId);
    classifyMutation.mutate(batchId);
  };

  if (unclassifiedStats.totalUnclassified === 0) {
    return null;
  }

  return (
    <Alert variant="default" className="border-orange-200 bg-orange-50">
      <AlertCircle className="h-4 w-4 text-orange-600" />
      <AlertTitle className="text-orange-900">
        Resultados não classificados
      </AlertTitle>
      <AlertDescription className="text-orange-800">
        <p className="mb-3">
          Há <strong>{unclassifiedStats.totalUnclassified} resultados</strong>{" "}
          em <strong>{unclassifiedStats.batchesWithUnclassified.length} lotes</strong> que ainda não foram
          classificados em níveis de risco.
        </p>

        <div className="space-y-2">
          {unclassifiedStats.batchesWithUnclassified.map((batch) => (
            <div
              key={batch.batchId}
              className="flex items-center justify-between rounded-md border border-orange-200 bg-white p-2"
            >
              <div className="flex-1">
                <p className="text-sm font-medium text-gray-900">
                  {batch.laboratoryName}
                </p>
                <p className="text-xs text-gray-500">
                  {new Date(batch.collectionDate).toLocaleDateString("pt-BR")} •{" "}
                  {batch.unclassifiedCount} resultado(s) não classificado(s)
                </p>
              </div>

              <Button
                size="sm"
                variant="outline"
                onClick={() => handleClassify(batch.batchId)}
                disabled={classifyingBatchId === batch.batchId}
                className="ml-2"
              >
                {classifyingBatchId === batch.batchId ? (
                  <>
                    <RefreshCw className="mr-1 h-3 w-3 animate-spin" />
                    Classificando...
                  </>
                ) : (
                  <>
                    <RefreshCw className="mr-1 h-3 w-3" />
                    Classificar
                  </>
                )}
              </Button>
            </div>
          ))}
        </div>

        <p className="mt-3 text-xs text-orange-700">
          A classificação automática usa os Score Items configurados para
          determinar o nível de risco baseado nos valores dos exames.
        </p>
      </AlertDescription>
    </Alert>
  );
}
