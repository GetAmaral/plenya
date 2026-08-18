"use client";

import { useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  AlertCircle,
  ChevronDown,
  ChevronRight,
  FileText,
  RefreshCw,
  Sparkles,
} from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import {
  labResultBatchApi,
  openLabBatchPDF,
} from "@/lib/api/lab-result-batch-api";
import { ProcessingStatus } from "@/components/lab-results/ProcessingStatus";
import { toast } from "sonner";

interface AlertLabResult {
  id: string;
  testName: string;
  labTestDefinition?: { name: string } | null;
  resultNumeric?: number | null;
  resultText?: string | null;
  unit?: string | null;
  labTestDefinitionId?: string | null;
  level?: number | null;
  classifyReason?: string | null;
  matchReason?: string | null;
}

interface UnclassifiedResultsAlertProps {
  batches: Array<{
    id: string;
    laboratoryName: string;
    collectionDate: string;
    hasPdf?: boolean;
    labResults: AlertLabResult[];
  }>;
}

/**
 * Nem todo "sem nível" é pendência. Dois casos são decisão do sistema, não trabalho parado:
 * o exame não se aplica ao paciente (sexo/idade/menopausa) e o exame não entra no escore
 * (sem ScoreItem configurado). Contá-los fazia o aviso nunca zerar, com um botão Classificar
 * que não tinha o que resolver. Eles continuam visíveis na lista expandida, com o motivo.
 */
function isInformationalReason(result: AlertLabResult) {
  const reason = (result.classifyReason ?? "").toLowerCase();
  return (
    reason.startsWith("não se aplica") || reason.includes("não entra no escore")
  );
}

function isUnclassified(result: AlertLabResult) {
  return (
    result.resultNumeric != null &&
    result.labTestDefinitionId != null &&
    result.level == null &&
    !isInformationalReason(result)
  );
}

/** Sem nível por decisão do sistema — listado como contexto, fora da contagem de pendências. */
function isInformational(result: AlertLabResult) {
  return (
    result.labTestDefinitionId != null &&
    result.level == null &&
    isInformationalReason(result)
  );
}

/** Exame que a IA leu mas não achou no catálogo — fica fora do escore. */
function isUnmatched(result: AlertLabResult) {
  return result.labTestDefinitionId == null;
}

function resultName(result: AlertLabResult) {
  return result.labTestDefinition?.name || result.testName || "(sem nome)";
}

function resultValue(result: AlertLabResult) {
  const value =
    result.resultNumeric != null
      ? String(result.resultNumeric)
      : result.resultText ?? "";
  return [value, result.unit].filter(Boolean).join(" ");
}

export function UnclassifiedResultsAlert({
  batches,
}: UnclassifiedResultsAlertProps) {
  const queryClient = useQueryClient();
  const [busyBatchId, setBusyBatchId] = useState<string | null>(null);
  const [expandedBatchId, setExpandedBatchId] = useState<string | null>(null);
  const [reinterpretJob, setReinterpretJob] = useState<{
    batchId: string;
    jobId: string;
  } | null>(null);

  const pendingBatches = batches
    .map((batch) => ({
      batch,
      unclassified: batch.labResults.filter(isUnclassified),
      unmatched: batch.labResults.filter(isUnmatched),
      informational: batch.labResults.filter(isInformational),
    }))
    .filter((entry) => entry.unclassified.length > 0);

  const totalUnclassified = pendingBatches.reduce(
    (sum, entry) => sum + entry.unclassified.length,
    0
  );

  const classifyMutation = useMutation({
    mutationFn: (batchId: string) => labResultBatchApi.classify(batchId),
    onSuccess: () => {
      toast.success("Lote reclassificado");
      queryClient.invalidateQueries({ queryKey: ["lab-result-batches"] });
      setBusyBatchId(null);
    },
    onError: (error: any) => {
      toast.error("Erro ao classificar lote", {
        description: error?.message || "Tente novamente",
      });
      setBusyBatchId(null);
    },
  });

  const reinterpretMutation = useMutation({
    mutationFn: (batchId: string) => labResultBatchApi.reinterpret(batchId),
    onSuccess: (data, batchId) => {
      toast.success("A IA está relendo o laudo", {
        description: `${data.removedResults} resultado(s) da leitura anterior foram substituídos.`,
      });
      setReinterpretJob({ batchId, jobId: data.jobId });
      setBusyBatchId(null);
    },
    onError: (error: any) => {
      toast.error("Não deu para reinterpretar", {
        description: error?.message || "Tente novamente",
      });
      setBusyBatchId(null);
    },
  });

  const handleClassify = (batchId: string) => {
    setBusyBatchId(batchId);
    classifyMutation.mutate(batchId);
  };

  const handleReinterpret = (batchId: string) => {
    setBusyBatchId(batchId);
    reinterpretMutation.mutate(batchId);
  };

  if (totalUnclassified === 0 && !reinterpretJob) {
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
          Há <strong>{totalUnclassified} resultados</strong> em{" "}
          <strong>{pendingBatches.length} lote(s)</strong> que ainda não foram
          classificados em níveis de risco. Abra o lote para ver quais são e o
          motivo de cada um.
        </p>

        {/* Progresso da releitura fica fora da lista: quando a IA termina, o lote pode
            sair da lista de pendentes e o acompanhamento sumiria no meio do caminho. */}
        {reinterpretJob && (
          <div className="mb-3">
            <ProcessingStatus
              jobId={reinterpretJob.jobId}
              onCompleted={() => {
                queryClient.invalidateQueries({
                  queryKey: ["lab-result-batches"],
                });
                toast.success("Laudo reinterpretado");
                setReinterpretJob(null);
              }}
              onFailed={(error) => {
                toast.error("A releitura falhou", { description: error });
                setReinterpretJob(null);
              }}
            />
          </div>
        )}

        <div className="space-y-2">
          {pendingBatches.map(({ batch, unclassified, unmatched, informational }) => {
            const expanded = expandedBatchId === batch.id;
            const busy = busyBatchId === batch.id;

            return (
              <div
                key={batch.id}
                className="rounded-md border border-orange-200 bg-white"
              >
                <div className="flex flex-wrap items-center justify-between gap-2 p-2">
                  <button
                    type="button"
                    onClick={() =>
                      setExpandedBatchId(expanded ? null : batch.id)
                    }
                    className="flex flex-1 items-center gap-2 text-left"
                    aria-expanded={expanded}
                  >
                    {expanded ? (
                      <ChevronDown className="h-4 w-4 shrink-0 text-muted-foreground" />
                    ) : (
                      <ChevronRight className="h-4 w-4 shrink-0 text-muted-foreground" />
                    )}
                    <span className="flex-1">
                      <span className="block text-sm font-medium text-foreground">
                        {batch.laboratoryName}
                      </span>
                      <span className="block text-xs text-muted-foreground">
                        {new Date(batch.collectionDate).toLocaleDateString(
                          "pt-BR"
                        )}{" "}
                        • {unclassified.length} sem nível
                        {unmatched.length > 0 &&
                          ` • ${unmatched.length} fora do catálogo`}
                        {informational.length > 0 &&
                          ` • ${informational.length} fora do escore`}
                      </span>
                    </span>
                  </button>

                  <div className="flex items-center gap-2">
                    <Button
                      size="sm"
                      variant="outline"
                      onClick={() => handleClassify(batch.id)}
                      disabled={busy}
                    >
                      <RefreshCw
                        className={`mr-1 h-3 w-3 ${
                          busy && classifyMutation.isPending
                            ? "animate-spin"
                            : ""
                        }`}
                      />
                      Classificar
                    </Button>

                    {batch.hasPdf !== false && (
                      <>
                        <Button
                          size="sm"
                          variant="outline"
                          onClick={() => handleReinterpret(batch.id)}
                          disabled={busy}
                          title="Apaga os resultados vindos do PDF e manda a IA reler o laudo original"
                        >
                          <Sparkles className="mr-1 h-3 w-3" />
                          Reinterpretar com IA
                        </Button>

                        <Button
                          size="sm"
                          variant="ghost"
                          onClick={() => {
                            openLabBatchPDF(batch.id).catch(() =>
                              toast.error("Não foi possível abrir o PDF original")
                            );
                          }}
                          title="Abre o laudo original em PDF para conferir o que o laboratório reportou"
                        >
                          <FileText className="mr-1 h-3 w-3" />
                          Ver PDF
                        </Button>
                      </>
                    )}
                  </div>
                </div>

                {expanded && (
                  <div className="overflow-x-auto border-t border-orange-100 p-2">
                    <table className="w-full min-w-[420px] text-left text-xs">
                      <thead className="text-muted-foreground">
                        <tr>
                          <th className="py-1 pr-2 font-medium">Exame</th>
                          <th className="py-1 pr-2 font-medium">Resultado</th>
                          <th className="py-1 font-medium">Motivo</th>
                        </tr>
                      </thead>
                      <tbody>
                        {unclassified.map((result) => (
                          <tr
                            key={result.id}
                            className="border-t border-orange-50"
                          >
                            <td className="py-1 pr-2 font-medium text-foreground">
                              {resultName(result)}
                            </td>
                            <td className="py-1 pr-2 tabular-nums">
                              {resultValue(result)}
                            </td>
                            <td className="py-1 text-muted-foreground">
                              {result.classifyReason ||
                                "Sem faixa configurada para este exame"}
                            </td>
                          </tr>
                        ))}
                        {unmatched.map((result) => (
                          <tr
                            key={result.id}
                            className="border-t border-orange-50"
                          >
                            <td className="py-1 pr-2 font-medium text-foreground">
                              {resultName(result)}
                            </td>
                            <td className="py-1 pr-2 tabular-nums">
                              {resultValue(result)}
                            </td>
                            <td className="py-1 text-muted-foreground">
                              {result.matchReason ||
                                "Não encontrado no catálogo de exames"}
                            </td>
                          </tr>
                        ))}
                        {informational.map((result) => (
                          <tr
                            key={result.id}
                            className="border-t border-orange-50 text-muted-foreground"
                          >
                            <td className="py-1 pr-2 font-medium">
                              {resultName(result)}
                            </td>
                            <td className="py-1 pr-2 tabular-nums">
                              {resultValue(result)}
                            </td>
                            <td className="py-1">{result.classifyReason}</td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                )}
              </div>
            );
          })}
        </div>

        <p className="mt-3 text-xs text-orange-700">
          &quot;Classificar&quot; reaplica os Score Items configurados sobre os
          valores já lidos. &quot;Reinterpretar com IA&quot; volta ao PDF
          original e refaz a leitura — use quando o problema for o que foi
          extraído, não a faixa de risco.
        </p>
      </AlertDescription>
    </Alert>
  );
}
