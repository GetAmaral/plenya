'use client';

/**
 * /continuum — panorama da equipe (Fase 7 ainda não implementada).
 *
 * Fase 1 entrega só placeholder com atalhos pros editores de templates,
 * que é onde admin/manager configura programa Semestral, Anual e boxes.
 * Quando Fase 2 (inscrição) e Fase 7 (dashboard) entrarem, esta página vira
 * o hub operacional com 3 visões: por paciente / por semana / por alertas.
 */
import Link from 'next/link';
import { useRequireAuth } from '@/lib/use-auth';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { PageHeader } from '@/components/layout/page-header';
import { LayoutTemplate, Package, Workflow } from 'lucide-react';
import { useContinuumTemplates, useContinuumBoxTemplates } from '@/lib/api/continuum-api';

export default function ContinuumPage() {
  useRequireAuth();
  const { data: templates = [] } = useContinuumTemplates();
  const { data: boxes = [] } = useContinuumBoxTemplates();

  return (
    <div className="container mx-auto space-y-6 py-6">
      <PageHeader
        breadcrumbs={[{ label: 'Continuum' }]}
        title="Continuum Plenya"
        description="Programa de acompanhamento longitudinal multidisciplinar — Médico, Nutricionista, Psicólogo e Educador Físico organizados pelo Método AGIR e mensurados pelo Escore Plenya."
      />

      <div className="grid gap-4 md:grid-cols-3">
        <Card>
          <CardHeader>
            <div className="flex items-center gap-2">
              <Workflow className="h-5 w-5 text-primary" />
              <CardTitle className="text-base">Inscrições ativas</CardTitle>
            </div>
            <CardDescription>Pacientes em programa</CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-semibold">0</p>
            <p className="mt-1 text-xs text-muted-foreground">
              Disponível na Fase 2 (inscrição de paciente)
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <div className="flex items-center gap-2">
              <LayoutTemplate className="h-5 w-5 text-primary" />
              <CardTitle className="text-base">Templates de programa</CardTitle>
            </div>
            <CardDescription>Semestral, Anual e variações</CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-semibold">{templates.length}</p>
            <Link href="/continuum/templates" className="mt-2 inline-block">
              <Button size="sm" variant="outline">
                Gerenciar templates
              </Button>
            </Link>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <div className="flex items-center gap-2">
              <Package className="h-5 w-5 text-primary" />
              <CardTitle className="text-base">Templates de Box</CardTitle>
            </div>
            <CardDescription>Boas-vindas, mensal, reavaliação</CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-3xl font-semibold">{boxes.length}</p>
            <Link href="/continuum/box-templates" className="mt-2 inline-block">
              <Button size="sm" variant="outline">
                Gerenciar boxes
              </Button>
            </Link>
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Próximas fases</CardTitle>
          <CardDescription>
            O Continuum está sendo construído em fases incrementais. Esta tela vai virar o hub
            operacional da equipe quando as próximas fases entrarem.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-2 text-sm">
          <p>• <b>Fase 2:</b> Inscrever paciente em um template (gera timeline pré-agendada).</p>
          <p>• <b>Fase 3:</b> Ancorar consultas reais aos marcos do programa.</p>
          <p>• <b>Fase 4:</b> Plano integrado markdown colaborativo.</p>
          <p>• <b>Fase 5:</b> Logística de Box (planned → shipped → delivered).</p>
          <p>• <b>Fase 6:</b> Prontuário agregado (tab paralela ao paciente).</p>
          <p>• <b>Fase 7:</b> Panorama equipe — por paciente, por semana, por alertas.</p>
        </CardContent>
      </Card>
    </div>
  );
}
