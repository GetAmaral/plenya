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
import { LabTestDefinitionForm } from "./LabTestDefinitionForm";
import {
  labTestDefinitionsApi,
  type LabTestDefinition,
} from "@/lib/api/lab-test-api";
import type { LabTestDefinitionFormValues } from "@/lib/validations/lab-test-definition";

interface LabTestDefinitionDialogProps {
  mode: "create" | "edit";
  open: boolean;
  onOpenChange: (open: boolean) => void;
  initialData?: LabTestDefinition;
  availableTests: LabTestDefinition[];
}

export function LabTestDefinitionDialog({
  mode,
  open,
  onOpenChange,
  initialData,
  availableTests,
}: LabTestDefinitionDialogProps) {
  const queryClient = useQueryClient();
  const [isSubmitting, setIsSubmitting] = useState(false);

  const createMutation = useMutation({
    mutationFn: labTestDefinitionsApi.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["lab-test-definitions"] });
      toast.success("Definição criada", {
        description: "A definição de exame foi criada com sucesso.",
      });
      onOpenChange(false);
    },
    onError: (error: any) => {
      toast.error("Erro ao criar", {
        description: error.message || "Não foi possível criar a definição.",
      });
    },
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<LabTestDefinitionFormValues> }) =>
      labTestDefinitionsApi.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["lab-test-definitions"] });
      toast.success("Definição atualizada", {
        description: "A definição de exame foi atualizada com sucesso.",
      });
      onOpenChange(false);
    },
    onError: (error: any) => {
      toast.error("Erro ao atualizar", {
        description: error.message || "Não foi possível atualizar a definição.",
      });
    },
  });

  const handleSubmit = async (values: LabTestDefinitionFormValues) => {
    setIsSubmitting(true);
    try {
      if (mode === "create") {
        await createMutation.mutateAsync(values as any);
      } else if (initialData) {
        await updateMutation.mutateAsync({
          id: initialData.id,
          data: values,
        });
      }
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleCancel = () => {
    onOpenChange(false);
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-3xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>
            {mode === "create"
              ? "Nova Definição de Exame"
              : "Editar Definição de Exame"}
          </DialogTitle>
          <DialogDescription>
            {mode === "create"
              ? "Crie uma nova definição de exame laboratorial ou parâmetro."
              : "Atualize as informações da definição de exame."}
          </DialogDescription>
        </DialogHeader>

        <LabTestDefinitionForm
          mode={mode}
          initialData={initialData}
          availableTests={availableTests}
          onSubmit={handleSubmit}
          onCancel={handleCancel}
          isSubmitting={isSubmitting}
        />
      </DialogContent>
    </Dialog>
  );
}
