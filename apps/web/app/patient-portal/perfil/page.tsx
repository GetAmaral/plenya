"use client";

/**
 * /perfil — dados pessoais editáveis + segurança + comunicação + LGPD.
 */
import { useState, useEffect } from "react";
import { toast } from "sonner";
import {
  Loader2,
  Lock,
  Download,
  Trash2,
  Save,
  User as UserIcon,
} from "lucide-react";

import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import {
  usePatientMe,
  useUpdatePatientProfile,
  useSetPatientPassword,
  useRequestAccountDelete,
  patientProfileApi,
  type UpdatableProfile,
} from "@/lib/api/patient-portal-api";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { Textarea } from "@/components/ui/textarea";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { useAuthStore } from "@/lib/auth-store";

export default function MyProfilePage() {
  const { ready } = useRequirePatientAuth();
  const { user } = useAuthStore();
  const { data: me, isLoading } = usePatientMe();
  const update = useUpdatePatientProfile();
  const setPassword = useSetPatientPassword();
  const requestDelete = useRequestAccountDelete();

  const [profile, setProfile] = useState<UpdatableProfile>({});
  const [pwd, setPwd] = useState("");
  const [pwd2, setPwd2] = useState("");
  const [deleteOpen, setDeleteOpen] = useState(false);
  const [deleteReason, setDeleteReason] = useState("");

  useEffect(() => {
    if (me?.patient) {
      setProfile({
        phone: me.patient.phone ?? "",
        email: (me.patient as any).email ?? "",
        address: me.patient.address ?? "",
        municipality: me.patient.municipality ?? "",
        state: me.patient.state ?? "",
        emergencyPhone: (me.patient as any).emergencyPhone ?? "",
      });
    }
  }, [me]);

  if (!ready || isLoading) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  const handleSaveProfile = async () => {
    try {
      await update.mutateAsync(profile);
      toast.success("Dados atualizados");
    } catch (e: any) {
      toast.error(e?.message ?? "Falha ao salvar");
    }
  };

  const handleSavePassword = async () => {
    if (pwd.length < 8) {
      toast.error("Senha precisa de ao menos 8 caracteres");
      return;
    }
    if (pwd !== pwd2) {
      toast.error("Senhas não conferem");
      return;
    }
    try {
      await setPassword.mutateAsync(pwd);
      toast.success("Senha atualizada");
      setPwd("");
      setPwd2("");
    } catch (e: any) {
      toast.error(e?.message ?? "Falha ao salvar senha");
    }
  };

  const handleExport = () => {
    // Browser baixa diretamente — backend retorna JSON com Content-Disposition
    const url = patientProfileApi.exportURL();
    const token = useAuthStore.getState().accessToken;
    if (!token) {
      toast.error("Sessão expirou");
      return;
    }
    fetch(url, { headers: { Authorization: `Bearer ${token}` } })
      .then(async (r) => {
        if (!r.ok) throw new Error("Falha ao exportar");
        const blob = await r.blob();
        const link = document.createElement("a");
        link.href = URL.createObjectURL(blob);
        link.download = "plenya-meus-dados.json";
        link.click();
      })
      .catch((e) => toast.error(e?.message ?? "Falha"));
  };

  const handleRequestDelete = async () => {
    try {
      await requestDelete.mutateAsync(deleteReason);
      toast.success("Solicitação enviada. Nossa equipe entra em contato.");
      setDeleteOpen(false);
      setDeleteReason("");
    } catch (e: any) {
      toast.error(e?.message ?? "Falha ao solicitar");
    }
  };

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">Conta</p>
        <h1 className="text-3xl font-light">{user?.name ?? "Meu perfil"}</h1>
        <p className="text-sm text-muted-foreground">{user?.email}</p>
      </header>

      <Tabs defaultValue="data">
        <TabsList>
          <TabsTrigger value="data">Dados</TabsTrigger>
          <TabsTrigger value="security">Segurança</TabsTrigger>
          <TabsTrigger value="lgpd">LGPD</TabsTrigger>
        </TabsList>

        <TabsContent value="data" className="mt-4">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-base">
                <UserIcon className="h-4 w-4" /> Dados de contato
              </CardTitle>
              <CardDescription>
                Para corrigir nome, CPF ou data de nascimento, fale com a equipe.
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid gap-3 md:grid-cols-2">
                <Field
                  label="Telefone"
                  value={profile.phone}
                  onChange={(v) => setProfile({ ...profile, phone: v })}
                />
                <Field
                  label="Email pessoal"
                  type="email"
                  value={profile.email}
                  onChange={(v) => setProfile({ ...profile, email: v })}
                />
                <Field
                  label="Endereço"
                  value={profile.address}
                  onChange={(v) => setProfile({ ...profile, address: v })}
                />
                <Field
                  label="Cidade"
                  value={profile.municipality}
                  onChange={(v) => setProfile({ ...profile, municipality: v })}
                />
                <Field
                  label="UF"
                  value={profile.state}
                  onChange={(v) => setProfile({ ...profile, state: v.toUpperCase().slice(0, 2) })}
                />
                <Field
                  label="Telefone de emergência"
                  value={profile.emergencyPhone}
                  onChange={(v) => setProfile({ ...profile, emergencyPhone: v })}
                />
              </div>
              <Button onClick={handleSaveProfile} disabled={update.isPending}>
                {update.isPending ? <Loader2 className="mr-1 h-4 w-4 animate-spin" /> : <Save className="mr-1 h-4 w-4" />}
                Salvar alterações
              </Button>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="security" className="mt-4 space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-base">
                <Lock className="h-4 w-4" /> Definir / trocar senha
              </CardTitle>
              <CardDescription>Mínimo 8 caracteres. Pode entrar sempre por link mágico se preferir não usar senha.</CardDescription>
            </CardHeader>
            <CardContent className="space-y-3">
              <div className="grid gap-3 md:grid-cols-2">
                <div className="space-y-2">
                  <Label htmlFor="pwd">Nova senha</Label>
                  <Input id="pwd" type="password" value={pwd} onChange={(e) => setPwd(e.target.value)} autoComplete="new-password" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="pwd2">Confirme</Label>
                  <Input id="pwd2" type="password" value={pwd2} onChange={(e) => setPwd2(e.target.value)} autoComplete="new-password" />
                </div>
              </div>
              <Button onClick={handleSavePassword} disabled={setPassword.isPending || !pwd}>
                {setPassword.isPending ? <Loader2 className="h-4 w-4 animate-spin" /> : "Atualizar senha"}
              </Button>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="lgpd" className="mt-4 space-y-4">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-base">
                <Download className="h-4 w-4" /> Exportar meus dados
              </CardTitle>
              <CardDescription>
                LGPD art. 18 V — portabilidade. Baixa um arquivo JSON com tudo que mantemos sobre você.
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Button variant="outline" onClick={handleExport}>
                <Download className="mr-1 h-4 w-4" /> Baixar dados
              </Button>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-base text-destructive">
                <Trash2 className="h-4 w-4" /> Solicitar exclusão
              </CardTitle>
              <CardDescription>
                LGPD art. 18 VI. Nossa equipe analisa o pedido (alguns dados clínicos têm retenção legal de 20 anos).
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Button variant="outline" onClick={() => setDeleteOpen(true)}>
                <Trash2 className="mr-1 h-4 w-4" /> Solicitar exclusão da conta
              </Button>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>

      <AlertDialog open={deleteOpen} onOpenChange={setDeleteOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Solicitar exclusão de dados?</AlertDialogTitle>
            <AlertDialogDescription>
              A equipe vai entrar em contato para validar a solicitação. Alguns dados clínicos podem precisar ser mantidos por exigência legal.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <Textarea
            placeholder="Conte o motivo (opcional)"
            value={deleteReason}
            onChange={(e) => setDeleteReason(e.target.value)}
            rows={3}
          />
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={handleRequestDelete}>Enviar solicitação</AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  );
}

function Field({
  label,
  value,
  onChange,
  type = "text",
}: {
  label: string;
  value?: string;
  onChange: (v: string) => void;
  type?: string;
}) {
  return (
    <div className="space-y-2">
      <Label>{label}</Label>
      <Input
        type={type}
        value={value ?? ""}
        onChange={(e) => onChange(e.target.value)}
      />
    </div>
  );
}
