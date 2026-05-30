"use client";

import { useRef, useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
  DialogDescription,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { createPatient, type Patient } from "@/lib/api/patients";
import { useFormNavigation } from "@/lib/use-form-navigation";

interface QuickAddPatientDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  // Chamado após criar; recebe o paciente novo. Use para selecionar/encadear.
  onCreated?: (patient: Patient) => void;
}

// Cadastro rápido (recepção): nome obrigatório, telefone recomendado.
// Demais dados completados depois na ficha.
export function QuickAddPatientDialog({
  open,
  onOpenChange,
  onCreated,
}: QuickAddPatientDialogProps) {
  const formRef = useRef<HTMLFormElement>(null);
  useFormNavigation({ formRef });
  const queryClient = useQueryClient();

  const [name, setName] = useState("");
  const [phone, setPhone] = useState("");

  const createMutation = useMutation({
    mutationFn: () => createPatient({ name: name.trim(), phone: phone.trim() || undefined }),
    onSuccess: (patient) => {
      queryClient.invalidateQueries({ queryKey: ["patients"] });
      toast.success("Paciente cadastrado");
      setName("");
      setPhone("");
      onOpenChange(false);
      onCreated?.(patient);
    },
    onError: (e: unknown) => {
      toast.error(e instanceof Error ? e.message : "Falha ao cadastrar");
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (name.trim().length < 3) {
      toast.error("Informe o nome completo (mínimo 3 letras)");
      return;
    }
    createMutation.mutate();
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Cadastro rápido</DialogTitle>
          <DialogDescription>
            Registre nome e telefone agora. O restante da ficha pode ser
            completado depois.
          </DialogDescription>
        </DialogHeader>
        <form ref={formRef} onSubmit={handleSubmit} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="qa-name">Nome completo</Label>
            <Input
              id="qa-name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="Maria da Silva"
              autoFocus
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="qa-phone">Telefone</Label>
            <Input
              id="qa-phone"
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
              placeholder="(43) 99999-9999"
            />
          </div>
          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={() => onOpenChange(false)}
            >
              Cancelar
            </Button>
            <Button type="submit" disabled={createMutation.isPending}>
              {createMutation.isPending ? "Salvando..." : "Cadastrar"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}

export default QuickAddPatientDialog;
