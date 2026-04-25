"use client";

/**
 * /  — Home da área do paciente. Placeholder Fase 1.
 * Fase 2 vai trazer dashboard completo (próxima consulta, status Continuum, escore, etc).
 */
import { useRequirePatientAuth } from "@/lib/use-patient-auth";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Loader2 } from "lucide-react";

export default function PatientHomePage() {
  const { ready, user } = useRequirePatientAuth();

  if (!ready) {
    return (
      <div className="flex h-[60vh] items-center justify-center">
        <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
      </div>
    );
  }

  const firstName = user?.name?.split(" ")[0] ?? "paciente";

  return (
    <div className="mx-auto max-w-5xl space-y-8">
      <header className="space-y-2">
        <p className="text-sm uppercase tracking-wider text-muted-foreground">
          Início
        </p>
        <h1 className="text-3xl font-light md:text-4xl">
          Olá, {firstName}.
        </h1>
        <p className="text-muted-foreground">
          Aqui você acompanha seu programa, consultas, exames e fala com a equipe.
        </p>
      </header>

      <div className="grid gap-4 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Meu Continuum</CardTitle>
          </CardHeader>
          <CardContent className="text-sm text-muted-foreground">
            Em breve — o quadro do seu programa entra aqui na Fase 2.
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle className="text-base">Próxima consulta</CardTitle>
          </CardHeader>
          <CardContent className="text-sm text-muted-foreground">
            Em breve — sua próxima consulta aparece aqui.
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
