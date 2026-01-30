"use client";

import { useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { LabResultEntryForm } from "./LabResultEntryForm";
import { labResultsApi, labTestDefinitionsApi } from "@/lib/api/lab-test-api";
import type { LabResultValuesBatchInput } from "@/lib/validations/lab-result-value";
import { valueToApiPayload } from "@/lib/validations/lab-result-value";

interface LabResultEntryDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  currentUserId: string;
}

export function LabResultEntryDialog({
  open,
  onOpenChange,
  currentUserId,
}: LabResultEntryDialogProps) {
  const queryClient = useQueryClient();
  const [isSubmitting, setIsSubmitting] = useState(false);

  const createMutation = useMutation({
    mutationFn: async (data: {
      labResult: any;
      values: any[];
    }) => {
      return labResultsApi.createWithValues(data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["lab-results"] });
      toast.success("Resultado salvo", {
        description: "O resultado do exame foi salvo com sucesso.",
      });
      onOpenChange(false);
    },
    onError: (error: any) => {
      toast.error("Erro ao salvar", {
        description:
          error.message || "Não foi possível salvar o resultado do exame.",
      });
    },
  });

  const handleSubmit = async (formData: LabResultValuesBatchInput) => {
    setIsSubmitting(true);
    try {
      // Get the selected test to get its name
      const test = await labTestDefinitionsApi.getById(formData.selectedTestId);

      // Create LabResult container
      // Note: patientId is omitted - backend will auto-fill from selectedPatient
      const labResult = {
        doctorId: formData.doctorId,
        testType: test.name,
        testDate: formData.testDate.toISOString(),
        notes: formData.notes,
        status: "completed" as const,
      };

      // Convert form values to API payload
      const values = formData.values.map((value: any, index: number) => {
        // Get test definition from the value
        const testDef = test.subTests?.find(
          (t) => t.id === value.labTestDefinitionId
        ) || test;

        return valueToApiPayload(
          value.labTestDefinitionId,
          value,
          testDef.resultType,
          testDef.unit
        );
      });

      await createMutation.mutateAsync({ labResult, values });
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleCancel = () => {
    onOpenChange(false);
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-4xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Novo Resultado de Exame</DialogTitle>
          <DialogDescription>
            Registre os resultados de um exame laboratorial realizado.
          </DialogDescription>
        </DialogHeader>

        <LabResultEntryForm
          currentUserId={currentUserId}
          onSubmit={handleSubmit}
          onCancel={handleCancel}
          isSubmitting={isSubmitting}
        />
      </DialogContent>
    </Dialog>
  );
}
