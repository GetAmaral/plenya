"use client";

/**
 * /preparacao — preparação pré-consulta do paciente.
 *
 * Depois de agendar e pagar, o paciente preenche um formulário curto (subset curado
 * do Escore) e envia exames anteriores, para o Dr. chegar preparado na consulta.
 * Os itens do formulário serão curados depois; até lá a config vem vazia.
 */
import { useEffect, useMemo, useRef, useState } from "react";
import { Loader2, Upload, FileText, CheckCircle2 } from "lucide-react";

import { formatDate } from "@/lib/format-date";
import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import {
  useMyAppointments,
  usePrepConfig,
  usePrepPrefill,
  useMyPrep,
  useSubmitPrep,
  useUploadExam,
} from "@/lib/api/patient-portal-api";
import { PrepForm, type PrepAnswers } from "@/components/patient-portal/prep-form";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";

export default function PreparacaoPage() {
  const { ready } = useRequirePatientAuth();
  const { data: appointments } = useMyAppointments("upcoming");
  const nextAppt = appointments?.[0];
  const appointmentId = nextAppt?.id;

  const { data: config, isLoading: loadingConfig } = usePrepConfig(appointmentId);
  const { data: existing } = useMyPrep(appointmentId);
  const { data: prefill } = usePrepPrefill(appointmentId);
  const submit = useSubmitPrep();
  const uploadExam = useUploadExam();

  const [chiefComplaint, setChiefComplaint] = useState("");
  const [answers, setAnswers] = useState<PrepAnswers>({});
  const [savedMsg, setSavedMsg] = useState<string | null>(null);
  const [uploaded, setUploaded] = useState<string[]>([]);
  const [prefilledFromTriage, setPrefilledFromTriage] = useState(false);
  const fileRef = useRef<HTMLInputElement>(null);
  const hydrated = useRef(false);

  // Hidrata uma vez: prefill da Triagem como base, rascunho salvo sobrescreve.
  useEffect(() => {
    if (hydrated.current) return;
    if (existing === undefined) return; // ainda carregando o rascunho
    if (appointmentId && prefill === undefined) return; // espera o prefill quando há consulta
    hydrated.current = true;

    const map: PrepAnswers = {};
    let usedPrefill = false;
    for (const r of prefill?.responses ?? []) {
      map[r.scoreItemId] = r;
      usedPrefill = true;
    }
    if (existing) {
      setChiefComplaint(existing.chiefComplaint ?? "");
      const saved = existing.responses ?? [];
      for (const r of saved) map[r.scoreItemId] = r;
      if (saved.length > 0) usedPrefill = false; // já tem rascunho próprio
    }
    setAnswers(map);
    setPrefilledFromTriage(usedPrefill && Object.keys(map).length > 0);
  }, [existing, prefill, appointmentId]);

  const isSubmitted = existing?.status === "submitted";
  const hasItems = useMemo(
    () => (config?.groups ?? []).some((g) => g.subgroups.some((sg) => sg.items.length > 0)),
    [config],
  );

  const save = (finalize: boolean) => {
    setSavedMsg(null);
    submit.mutate(
      {
        appointmentId,
        chiefComplaint: chiefComplaint || undefined,
        responses: Object.values(answers),
        submit: finalize,
      },
      {
        onSuccess: () =>
          setSavedMsg(finalize ? "Preparação enviada. Obrigado!" : "Rascunho salvo."),
        onError: (e: any) => setSavedMsg(e?.message ?? "Erro ao salvar."),
      },
    );
  };

  const onPickFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    const fd = new FormData();
    fd.append("file", file);
    fd.append("title", file.name);
    uploadExam.mutate(fd, {
      onSuccess: () => setUploaded((u) => [...u, file.name]),
    });
    if (fileRef.current) fileRef.current.value = "";
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
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">Preparação</p>
        <h1 className="text-2xl font-bold">Antes da sua consulta</h1>
        {nextAppt ? (
          <p className="text-sm text-muted-foreground">
            Consulta em{" "}
            {formatDate(nextAppt.scheduledAt, "d 'de' MMMM 'às' HH:mm")}.
            Quanto mais o Dr. souber antes, melhor a gente aproveita o tempo de vocês.
          </p>
        ) : (
          <p className="text-sm text-muted-foreground">
            Quando você tiver uma consulta agendada, a preparação aparece aqui.
          </p>
        )}
      </header>

      {isSubmitted && (
        <div className="flex items-center gap-2 rounded-lg border border-emerald-300 bg-emerald-50 px-4 py-3 text-sm text-emerald-800">
          <CheckCircle2 className="h-4 w-4" />
          Preparação enviada. Você ainda pode atualizar e enviar de novo se precisar.
        </div>
      )}

      {/* Queixa / objetivo principal */}
      <Card>
        <CardContent className="space-y-3 pt-6">
          <label className="text-sm font-medium">
            O que mais te incomoda hoje, ou o que você quer resolver?
          </label>
          <textarea
            value={chiefComplaint}
            onChange={(e) => setChiefComplaint(e.target.value)}
            rows={3}
            placeholder="Conte com suas palavras."
            className="w-full rounded-lg border border-border bg-background px-3 py-2 text-sm"
          />
        </CardContent>
      </Card>

      {/* Upload de exames anteriores */}
      <Card>
        <CardContent className="space-y-3 pt-6">
          <div className="space-y-1">
            <p className="text-sm font-medium">Exames anteriores</p>
            <p className="text-xs text-muted-foreground">
              Envie exames recentes (PDF ou imagem) para o Dr. analisar antes. Tudo confidencial.
            </p>
          </div>
          <input
            ref={fileRef}
            type="file"
            accept="application/pdf,image/jpeg,image/png"
            onChange={onPickFile}
            className="hidden"
          />
          <Button
            type="button"
            variant="outline"
            onClick={() => fileRef.current?.click()}
            disabled={uploadExam.isPending}
          >
            {uploadExam.isPending ? (
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            ) : (
              <Upload className="mr-2 h-4 w-4" />
            )}
            Enviar exame
          </Button>
          {uploaded.length > 0 && (
            <ul className="space-y-1">
              {uploaded.map((name, i) => (
                <li key={i} className="flex items-center gap-2 text-sm text-muted-foreground">
                  <FileText className="h-4 w-4" />
                  {name}
                </li>
              ))}
            </ul>
          )}
        </CardContent>
      </Card>

      {/* Formulário curado */}
      <Card>
        <CardContent className="pt-6">
          {loadingConfig ? (
            <div className="flex items-center justify-center py-8">
              <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
            </div>
          ) : hasItems && config ? (
            <>
              {prefilledFromTriage && (
                <p className="mb-4 rounded-md border border-amber-200 bg-amber-50 px-3 py-2 text-xs text-amber-800">
                  Pré-preenchemos algumas respostas a partir da sua triagem. Confira e atualize o que tiver mudado.
                </p>
              )}
              <PrepForm config={config} answers={answers} onChange={setAnswers} />
            </>
          ) : (
            <p className="text-sm text-muted-foreground">
              Nada a preencher por enquanto. Você já pode enviar seus exames acima.
            </p>
          )}
        </CardContent>
      </Card>

      {savedMsg && <p className="text-sm text-muted-foreground">{savedMsg}</p>}

      <div className="flex flex-wrap gap-3">
        <Button variant="outline" onClick={() => save(false)} disabled={submit.isPending}>
          Salvar rascunho
        </Button>
        <Button onClick={() => save(true)} disabled={submit.isPending}>
          {submit.isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
          Enviar para o Dr.
        </Button>
      </div>
    </div>
  );
}
