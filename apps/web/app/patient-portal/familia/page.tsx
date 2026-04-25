"use client";

/**
 * /familia — paciente convida familiares pra ver dados (escopo granular).
 */
import { useState } from "react";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { toast } from "sonner";
import {
  Loader2,
  UserPlus,
  Users,
  Trash2,
  CheckCircle2,
  Clock,
} from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import {
  useFamilyGrants,
  useCreateFamilyInvite,
  useRevokeFamilyGrant,
  useUpdateFamilyScope,
  type FamilyAccessScope,
  type FamilyGrantView,
} from "@/lib/api/patient-portal-api";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Badge } from "@/components/ui/badge";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";

const SCOPE_LABELS: Record<keyof FamilyAccessScope, string> = {
  appointments: "Consultas",
  continuum: "Continuum (programa)",
  exams: "Exames e prescrições",
  scores: "Escores",
  documents: "Documentos",
  boxes: "Boxes",
  messages: "Mensagens (somente leitura)",
};

const DEFAULT_SCOPE: FamilyAccessScope = {
  appointments: true,
  continuum: true,
  exams: false,
  scores: false,
  documents: false,
  boxes: false,
  messages: false,
};

export default function MyFamilyPage() {
  const { ready } = useRequirePatientAuth();
  const { data, isLoading } = useFamilyGrants();
  const create = useCreateFamilyInvite();
  const revoke = useRevokeFamilyGrant();

  const [open, setOpen] = useState(false);
  const [email, setEmail] = useState("");
  const [label, setLabel] = useState("");
  const [scope, setScope] = useState<FamilyAccessScope>(DEFAULT_SCOPE);

  const handleCreate = async () => {
    if (!email || !label) {
      toast.error("Email e rótulo obrigatórios");
      return;
    }
    try {
      await create.mutateAsync({ granteeEmail: email, granteeLabel: label, scope });
      toast.success("Convite enviado por email");
      setOpen(false);
      setEmail("");
      setLabel("");
      setScope(DEFAULT_SCOPE);
    } catch (e: any) {
      toast.error(e?.message ?? "Falha ao convidar");
    }
  };

  const handleRevoke = async (id: string) => {
    if (!confirm("Revogar acesso? A pessoa não verá mais seus dados.")) return;
    try {
      await revoke.mutateAsync(id);
      toast.success("Acesso revogado");
    } catch (e: any) {
      toast.error(e?.message ?? "Falha");
    }
  };

  if (!ready) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <header className="flex items-start justify-between">
        <div className="space-y-2">
          <p className="text-sm uppercase tracking-wider text-muted-foreground">Compartilhamento</p>
          <h1 className="text-3xl font-light">Família e cuidadores</h1>
          <p className="max-w-xl text-muted-foreground">
            Convide pessoas de confiança pra acompanhar parte da sua saúde. Você escolhe o que cada uma vê.
          </p>
        </div>
        <Dialog open={open} onOpenChange={setOpen}>
          <DialogTrigger asChild>
            <Button>
              <UserPlus className="mr-1 h-4 w-4" /> Convidar
            </Button>
          </DialogTrigger>
          <DialogContent className="max-w-md">
            <DialogHeader>
              <DialogTitle>Convidar familiar</DialogTitle>
              <DialogDescription>
                Enviaremos um link por email. A pessoa precisa ter (ou criar) uma conta no portal.
              </DialogDescription>
            </DialogHeader>
            <div className="space-y-3">
              <div className="space-y-2">
                <Label htmlFor="email">Email</Label>
                <Input id="email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
              </div>
              <div className="space-y-2">
                <Label htmlFor="label">Como chamar essa pessoa?</Label>
                <Input
                  id="label"
                  placeholder="Ex: Esposa Mariana, Filho João"
                  value={label}
                  onChange={(e) => setLabel(e.target.value)}
                />
              </div>
              <div className="space-y-2">
                <Label>O que ela pode ver?</Label>
                <div className="grid grid-cols-1 gap-2">
                  {(Object.keys(SCOPE_LABELS) as (keyof FamilyAccessScope)[]).map((k) => (
                    <Label key={k} className="flex cursor-pointer items-center gap-2 text-sm font-normal">
                      <Checkbox
                        checked={scope[k]}
                        onCheckedChange={(v) => setScope({ ...scope, [k]: !!v })}
                      />
                      {SCOPE_LABELS[k]}
                    </Label>
                  ))}
                </div>
              </div>
            </div>
            <DialogFooter>
              <Button variant="outline" onClick={() => setOpen(false)}>
                Cancelar
              </Button>
              <Button onClick={handleCreate} disabled={create.isPending}>
                {create.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : "Enviar convite"}
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </header>

      {isLoading ? (
        <Card className="flex h-32 items-center justify-center">
          <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
        </Card>
      ) : !data?.length ? (
        <Card>
          <CardContent className="py-10 text-center text-sm text-muted-foreground">
            <Users className="mx-auto mb-3 h-8 w-8" />
            Nenhum compartilhamento ativo.
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-3">
          {data.map((g) => (
            <GrantCard key={g.id} g={g} onRevoke={() => handleRevoke(g.id)} />
          ))}
        </div>
      )}
    </div>
  );
}

function GrantCard({ g, onRevoke }: { g: FamilyGrantView; onRevoke: () => void }) {
  const update = useUpdateFamilyScope();
  const [scope, setScope] = useState<FamilyAccessScope>(g.scope);
  const [editing, setEditing] = useState(false);

  const handleSave = async () => {
    try {
      await update.mutateAsync({ id: g.id, scope });
      toast.success("Permissões atualizadas");
      setEditing(false);
    } catch (e: any) {
      toast.error(e?.message ?? "Falha");
    }
  };

  const allowed = (Object.keys(SCOPE_LABELS) as (keyof FamilyAccessScope)[]).filter((k) => g.scope[k]);

  return (
    <Card>
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center justify-between text-base">
          <span>{g.granteeLabel}</span>
          {g.status === "active" ? (
            <Badge variant="default" className="gap-1">
              <CheckCircle2 className="h-3 w-3" /> Ativo
            </Badge>
          ) : (
            <Badge variant="outline" className="gap-1">
              <Clock className="h-3 w-3" /> Convite enviado
            </Badge>
          )}
        </CardTitle>
        <CardDescription>{g.granteeEmail}</CardDescription>
      </CardHeader>
      <CardContent className="space-y-3">
        {editing ? (
          <div className="space-y-2">
            {(Object.keys(SCOPE_LABELS) as (keyof FamilyAccessScope)[]).map((k) => (
              <Label key={k} className="flex cursor-pointer items-center gap-2 text-sm font-normal">
                <Checkbox
                  checked={scope[k]}
                  onCheckedChange={(v) => setScope({ ...scope, [k]: !!v })}
                />
                {SCOPE_LABELS[k]}
              </Label>
            ))}
            <div className="flex gap-2 pt-1">
              <Button size="sm" onClick={handleSave} disabled={update.isPending}>
                {update.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : "Salvar"}
              </Button>
              <Button size="sm" variant="outline" onClick={() => { setScope(g.scope); setEditing(false); }}>
                Cancelar
              </Button>
            </div>
          </div>
        ) : (
          <>
            <div className="flex flex-wrap gap-1.5">
              {allowed.length === 0 ? (
                <Badge variant="outline" className="text-muted-foreground">Sem permissões</Badge>
              ) : (
                allowed.map((k) => (
                  <Badge key={k} variant="secondary">
                    {SCOPE_LABELS[k]}
                  </Badge>
                ))
              )}
            </div>
            <p className="text-xs text-muted-foreground">
              {g.acceptedAt
                ? `Aceito em ${format(new Date(g.acceptedAt), "dd/MM/yyyy", { locale: ptBR })}`
                : `Convite válido até ${format(new Date(g.expiresAt), "dd/MM/yyyy", { locale: ptBR })}`}
            </p>
            <div className="flex gap-2">
              <Button size="sm" variant="outline" onClick={() => setEditing(true)}>
                Editar permissões
              </Button>
              <Button size="sm" variant="outline" onClick={onRevoke}>
                <Trash2 className="mr-1 h-3.5 w-3.5" /> Revogar
              </Button>
            </div>
          </>
        )}
      </CardContent>
    </Card>
  );
}
