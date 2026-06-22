"use client";

/**
 * PortalAccessCard — exibido no perfil do Patient (lado staff).
 * Mostra status do convite + botão "Convidar para o portal".
 *
 * Aparece pra admin/manager/secretary; outros roles podem ver mas não acionam.
 */
import { useState } from "react";
import { formatDate } from "@/lib/format-date";
import { toast } from "sonner";
import { Loader2, Send, Copy, MailCheck, MailWarning, Mail, MessageSquare } from "lucide-react";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Checkbox } from "@/components/ui/checkbox";
import { Label } from "@/components/ui/label";
import { useAuthStore, isGranted } from "@/lib/auth-store";
import {
  usePortalInviteStatus,
  useCreatePortalInvite,
} from "@/lib/api/patient-portal-api";

const ROLES_THAT_CAN_INVITE = ["admin", "manager", "secretary"] as const;

export function PortalAccessCard({ patientId }: { patientId: string }) {
  const { user } = useAuthStore();
  const canInvite = ROLES_THAT_CAN_INVITE.some((r) => isGranted(user, r));

  const { data: status, isLoading } = usePortalInviteStatus(patientId);
  const create = useCreatePortalInvite(patientId);
  const [sendEmail, setSendEmail] = useState(true);
  const [sendWA, setSendWA] = useState(false);
  const [lastResult, setLastResult] = useState<{ link: string } | null>(null);

  const handleInvite = async () => {
    try {
      const res = await create.mutateAsync({ sendEmail, sendWA });
      setLastResult({ link: res.link });
      const channels = [];
      if (res.sentEmail) channels.push("email");
      if (res.sentWA) channels.push("WhatsApp");
      toast.success(
        channels.length
          ? `Convite enviado por ${channels.join(" e ")}`
          : "Convite criado (envio falhou — copie o link manualmente)",
      );
    } catch (err: any) {
      toast.error(err?.response?.data?.error ?? "Falha ao gerar convite");
    }
  };

  const copy = (text: string) => {
    navigator.clipboard.writeText(text);
    toast.success("Link copiado");
  };

  return (
    <Card className="border-0 shadow-lg">
      <CardHeader>
        <CardTitle className="flex items-center justify-between text-base">
          Acesso à área do paciente
          {!isLoading && status?.status === "accepted" && (
            <Badge variant="default" className="gap-1">
              <MailCheck className="h-3 w-3" /> Ativo
            </Badge>
          )}
          {!isLoading && status?.status === "pending" && (
            <Badge variant="outline" className="gap-1">
              <Mail className="h-3 w-3" /> Convite pendente
            </Badge>
          )}
          {!isLoading && status?.status === "expired" && (
            <Badge variant="destructive" className="gap-1">
              <MailWarning className="h-3 w-3" /> Convite expirado
            </Badge>
          )}
          {!isLoading && status?.status === "none" && (
            <Badge variant="secondary">Sem convite</Badge>
          )}
        </CardTitle>
        <CardDescription>
          Portal em minha.plenyasaude.com.br — paciente vê Continuum, consultas, exames e fala com a equipe.
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        {status?.status === "accepted" && status.acceptedAt && (
          <p className="text-sm text-muted-foreground">
            Acesso ativado em{" "}
            {formatDate(status.acceptedAt, "dd/MM/yyyy 'às' HH:mm")}.
          </p>
        )}
        {status?.status === "pending" && status.expiresAt && (
          <p className="text-sm text-muted-foreground">
            Convite enviado, expira em{" "}
            {formatDate(status.expiresAt, "dd/MM/yyyy")}. Reenvie se necessário.
          </p>
        )}

        {canInvite && (
          <>
            <div className="flex flex-wrap items-center gap-4 text-sm">
              <Label className="flex cursor-pointer items-center gap-2">
                <Checkbox
                  checked={sendEmail}
                  onCheckedChange={(v) => setSendEmail(!!v)}
                />
                <Mail className="h-3.5 w-3.5" /> Email
              </Label>
              <Label className="flex cursor-pointer items-center gap-2">
                <Checkbox
                  checked={sendWA}
                  onCheckedChange={(v) => setSendWA(!!v)}
                />
                <MessageSquare className="h-3.5 w-3.5" /> WhatsApp
              </Label>
            </div>
            <div className="flex gap-2">
              <Button
                onClick={handleInvite}
                disabled={create.isPending || (!sendEmail && !sendWA)}
                size="sm"
              >
                {create.isPending ? (
                  <Loader2 className="mr-1 h-4 w-4 animate-spin" />
                ) : (
                  <Send className="mr-1 h-4 w-4" />
                )}
                {status?.status === "accepted" ? "Reenviar acesso" : "Convidar para o portal"}
              </Button>
            </div>
          </>
        )}

        {lastResult?.link && (
          <div className="rounded-md border bg-muted/40 p-3">
            <p className="mb-1 text-xs font-medium text-muted-foreground">
              Link do convite (válido por 14 dias)
            </p>
            <div className="flex items-center gap-2">
              <code className="flex-1 truncate text-xs">{lastResult.link}</code>
              <Button
                size="sm"
                variant="outline"
                onClick={() => copy(lastResult.link)}
              >
                <Copy className="h-3.5 w-3.5" />
              </Button>
            </div>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
