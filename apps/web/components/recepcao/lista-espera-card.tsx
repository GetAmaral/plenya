"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import {
  Clock,
  Plus,
  Loader2,
  CalendarPlus,
  X,
  Search,
  Check,
} from "lucide-react";
import { toast } from "sonner";
import { useRouter } from "next/navigation";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { apiClient } from "@/lib/api-client";
import type { Patient } from "@/lib/api/patient-api";
import { useDebouncedValue } from "@/lib/use-debounced-value";
import {
  APPOINTMENT_TYPE_LABELS,
  type AppointmentType,
} from "@/lib/api/calendar-api";
import {
  useWaitlist,
  useCreateWaitlistEntry,
  useUpdateWaitlistEntry,
  type WaitlistEntry,
} from "@/lib/api/waitlist";
import { EmptyState } from "./empty-state";
import { ErrorState } from "./error-state";

const TYPE_OPTIONS = Object.entries(APPOINTMENT_TYPE_LABELS) as [
  AppointmentType,
  string,
][];

/**
 * Linha de uma pessoa aguardando encaixe. Monta as proprias mutations com o id:
 * converter (vira consulta -> abre Nova consulta com o paciente selecionado) e
 * cancelar (marca status=cancelled). Ambas invalidam waitlistKeys.all, entao a
 * lista reage sozinha.
 */
function WaitlistRow({ entry }: { entry: WaitlistEntry }) {
  const router = useRouter();
  const update = useUpdateWaitlistEntry(entry.id);
  const patientName = entry.patient?.name ?? "Paciente";

  async function handleCancel() {
    try {
      await update.mutateAsync({ status: "cancelled" });
      toast.success("Removido da lista de espera", {
        description: patientName,
      });
    } catch (err) {
      toast.error("Nao foi possivel remover", {
        description: err instanceof Error ? err.message : "Tente novamente.",
      });
    }
  }

  function handleConvert() {
    // Encaminha para Nova consulta ja com o paciente em contexto. O agendamento
    // efetivo e a baixa da espera (status=scheduled) sao feitos no fluxo de
    // agenda; aqui so levamos a secretaria pra la sem perder o contexto.
    const params = new URLSearchParams({ patientId: entry.patientId });
    if (entry.preferredType) params.set("type", entry.preferredType);
    router.push(`/appointments/new?${params.toString()}`);
  }

  return (
    <div className="flex items-center justify-between gap-3 p-3 rounded-lg border border-border hover:bg-accent/50 transition-colors">
      <div className="min-w-0">
        <p className="font-medium truncate">{patientName}</p>
        <p className="text-sm text-muted-foreground truncate">
          {entry.preferredType
            ? APPOINTMENT_TYPE_LABELS[entry.preferredType]
            : "Sem preferencia de tipo"}
        </p>
        {entry.notes && (
          <p className="text-xs text-muted-foreground truncate">
            {entry.notes}
          </p>
        )}
      </div>
      <div className="flex items-center gap-2 shrink-0">
        <Button
          size="sm"
          variant="outline"
          onClick={handleConvert}
          title="Converter em consulta"
        >
          <CalendarPlus className="h-4 w-4" />
          <span className="hidden sm:inline">Agendar</span>
        </Button>
        <Button
          size="sm"
          variant="ghost"
          onClick={handleCancel}
          disabled={update.isPending}
          title="Cancelar"
        >
          {update.isPending ? (
            <Loader2 className="h-4 w-4 animate-spin" />
          ) : (
            <X className="h-4 w-4" />
          )}
        </Button>
      </div>
    </div>
  );
}

/**
 * Dialog de "Adicionar a lista de espera". Reusa o mesmo padrao de busca
 * server-side de pacientes da PatientQuickSearch (GET /api/v1/patients?search),
 * sem inventar endpoint. Seleciona o paciente, o tipo preferido (opcional) e
 * notas (opcional) e chama useCreateWaitlistEntry.
 */
function AddWaitlistDialog({
  open,
  onOpenChange,
}: {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}) {
  const create = useCreateWaitlistEntry();

  const [selected, setSelected] = useState<Patient | null>(null);
  const [term, setTerm] = useState("");
  const [preferredType, setPreferredType] = useState<AppointmentType | "">("");
  const [notes, setNotes] = useState("");

  const debounced = useDebouncedValue(term.trim(), 300);
  const searchEnabled = !selected && debounced.length >= 2;

  const { data: results = [], isFetching } = useQuery({
    queryKey: ["recepcao", "waitlist-patient-search", debounced],
    queryFn: () =>
      apiClient.get<Patient[]>(
        `/api/v1/patients?search=${encodeURIComponent(debounced)}&limit=8`,
      ),
    enabled: searchEnabled,
    staleTime: 10_000,
  });

  function reset() {
    setSelected(null);
    setTerm("");
    setPreferredType("");
    setNotes("");
  }

  function handleOpenChange(next: boolean) {
    if (!next) reset();
    onOpenChange(next);
  }

  async function handleSubmit() {
    if (!selected) {
      toast.error("Selecione um paciente");
      return;
    }
    try {
      await create.mutateAsync({
        patientId: selected.id,
        preferredType: preferredType || undefined,
        notes: notes.trim() || undefined,
      });
      toast.success("Adicionado a lista de espera", {
        description: selected.name,
      });
      reset();
      onOpenChange(false);
    } catch (err) {
      toast.error("Nao foi possivel adicionar", {
        description: err instanceof Error ? err.message : "Tente novamente.",
      });
    }
  }

  return (
    <Dialog open={open} onOpenChange={handleOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Adicionar a lista de espera</DialogTitle>
          <DialogDescription>
            Registre quem aguarda encaixe. Quando abrir um horario, converta em
            consulta direto daqui.
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          {/* Paciente */}
          <div className="space-y-2">
            <Label>Paciente</Label>
            {selected ? (
              <div className="flex items-center justify-between gap-3 rounded-lg border border-border p-3">
                <div className="min-w-0">
                  <p className="font-medium truncate">{selected.name}</p>
                  {(selected.phone || selected.email) && (
                    <p className="text-xs text-muted-foreground truncate">
                      {[selected.phone, selected.email]
                        .filter(Boolean)
                        .join(" · ")}
                    </p>
                  )}
                </div>
                <Button
                  size="sm"
                  variant="ghost"
                  onClick={() => setSelected(null)}
                >
                  Trocar
                </Button>
              </div>
            ) : (
              <div className="relative">
                <div className="relative">
                  <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
                  <Input
                    value={term}
                    onChange={(e) => setTerm(e.target.value)}
                    placeholder="Buscar paciente por nome, telefone ou CPF"
                    className="pl-9"
                    aria-label="Buscar paciente"
                    autoFocus
                  />
                  {isFetching && searchEnabled && (
                    <Loader2 className="absolute right-3 top-1/2 -translate-y-1/2 h-4 w-4 animate-spin text-muted-foreground" />
                  )}
                </div>
                {searchEnabled && (
                  <Card className="absolute z-20 mt-2 w-full border shadow-lg">
                    <CardContent className="p-2">
                      {results.length === 0 && !isFetching ? (
                        <p className="px-3 py-4 text-sm text-muted-foreground">
                          Nenhum paciente encontrado. Cadastre o paciente antes
                          de adicionar a lista.
                        </p>
                      ) : (
                        <ul className="max-h-60 overflow-auto">
                          {results.map((p) => (
                            <li key={p.id}>
                              <button
                                type="button"
                                onClick={() => {
                                  setSelected(p);
                                  setTerm("");
                                }}
                                className="flex w-full items-center justify-between gap-2 rounded-md px-3 py-2 text-left hover:bg-accent transition-colors"
                              >
                                <div className="min-w-0">
                                  <span className="block font-medium truncate">
                                    {p.name}
                                  </span>
                                  {(p.phone || p.email) && (
                                    <span className="block text-xs text-muted-foreground truncate">
                                      {[p.phone, p.email]
                                        .filter(Boolean)
                                        .join(" · ")}
                                    </span>
                                  )}
                                </div>
                                <Check className="h-4 w-4 shrink-0 opacity-0" />
                              </button>
                            </li>
                          ))}
                        </ul>
                      )}
                    </CardContent>
                  </Card>
                )}
              </div>
            )}
          </div>

          {/* Tipo preferido */}
          <div className="space-y-2">
            <Label htmlFor="waitlist-type">Tipo preferido</Label>
            <Select
              value={preferredType}
              onValueChange={(v) => setPreferredType(v as AppointmentType)}
            >
              <SelectTrigger id="waitlist-type">
                <SelectValue placeholder="Opcional" />
              </SelectTrigger>
              <SelectContent>
                {TYPE_OPTIONS.map(([value, label]) => (
                  <SelectItem key={value} value={value}>
                    {label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          {/* Notas */}
          <div className="space-y-2">
            <Label htmlFor="waitlist-notes">Notas</Label>
            <Textarea
              id="waitlist-notes"
              value={notes}
              onChange={(e) => setNotes(e.target.value)}
              placeholder="Preferencia de horario, urgencia, observacoes da secretaria"
              rows={3}
            />
          </div>
        </div>

        <DialogFooter>
          <Button
            type="button"
            variant="outline"
            onClick={() => handleOpenChange(false)}
          >
            Cancelar
          </Button>
          <Button
            type="button"
            onClick={handleSubmit}
            disabled={!selected || create.isPending}
          >
            {create.isPending ? (
              <Loader2 className="h-4 w-4 animate-spin" />
            ) : (
              <Plus className="h-4 w-4" />
            )}
            Adicionar
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

/**
 * "Lista de espera": pessoas aguardando encaixe (status=waiting). Lista
 * paciente, tipo preferido e notas, com acoes de agendar (converte em Nova
 * consulta) e cancelar por linha, mais o botao de adicionar via dialog.
 */
export function ListaEsperaCard() {
  const [addOpen, setAddOpen] = useState(false);
  const { data: entries = [], isLoading, isError, refetch } = useWaitlist("waiting");

  return (
    <Card className="border-0 shadow-md">
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span className="flex items-center gap-2">
            <Clock className="h-5 w-5 text-sky-600" />
            Lista de espera
          </span>
          {!isLoading && entries.length > 0 && (
            <Badge variant="observation">{entries.length}</Badge>
          )}
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {isError ? (
          <ErrorState
            title="Nao foi possivel carregar a lista de espera"
            onRetry={() => refetch()}
          />
        ) : isLoading ? (
          <div className="space-y-3">
            {Array.from({ length: 3 }).map((_, i) => (
              <Skeleton key={i} className="h-16 w-full" />
            ))}
          </div>
        ) : entries.length === 0 ? (
          <EmptyState
            icon={Clock}
            title="Ninguem aguardando encaixe"
            hint="Adicione um paciente a lista para chamar assim que abrir um horario."
          />
        ) : (
          <div className="space-y-3">
            {entries.map((entry) => (
              <WaitlistRow key={entry.id} entry={entry} />
            ))}
          </div>
        )}

        <Button
          variant="outline"
          className="w-full"
          onClick={() => setAddOpen(true)}
        >
          <Plus className="h-4 w-4" />
          Adicionar a lista
        </Button>
      </CardContent>

      <AddWaitlistDialog open={addOpen} onOpenChange={setAddOpen} />
    </Card>
  );
}
