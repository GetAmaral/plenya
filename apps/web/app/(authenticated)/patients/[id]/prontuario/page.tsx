'use client';

/**
 * /patients/[id]/prontuario — prontuário agregado.
 *
 * Lista cronológica (desc) de TODOS os eventos do paciente: anamneses, exames,
 * scores, avaliações, prescrições, consultas. Filtros multi-select por tipo
 * + por especialidade do autor. Paginação infinite scroll.
 *
 * Não específico de Continuum — vale pra qualquer paciente.
 */
import { useMemo, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import {
  Loader2,
  FileText,
  Microscope,
  Activity,
  Target,
  User as UserIcon,
  Pill,
  Calendar,
  Stethoscope,
} from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Checkbox } from '@/components/ui/checkbox';
import { PageHeader } from '@/components/layout/page-header';
import {
  useMedicalRecord,
  MR_TYPE_LABELS,
  MR_TYPE_COLORS,
  type MedicalRecordType,
  type MedicalRecordEntry,
} from '@/lib/api/medical-record-api';
import { cn } from '@/lib/utils';

const TYPE_ICON: Record<MedicalRecordType, any> = {
  anamnesis: Stethoscope,
  lab_batch: Microscope,
  score_snapshot: Target,
  physical_assessment: Activity,
  fitness_test: Target,
  postural_assessment: UserIcon,
  prescription: Pill,
  appointment: Calendar,
};

// Detecta especialidade a partir do array JSON de roles serializado em texto.
const ROLE_FILTERS = [
  { key: 'doctor', label: 'Médico' },
  { key: 'nutritionist', label: 'Nutricionista' },
  { key: 'psychologist', label: 'Psicólogo' },
  { key: 'physicalEducator', label: 'Educador Físico' },
];

export default function MedicalRecordPage() {
  const params = useParams<{ id: string }>();
  const router = useRouter();
  useRequireAuth();

  const [selectedTypes, setSelectedTypes] = useState<MedicalRecordType[]>([]);
  const [roleFilter, setRoleFilter] = useState<string>('all');

  const query = useMedicalRecord(params.id, {
    types: selectedTypes.length > 0 ? selectedTypes : undefined,
  });
  const allEntries = useMemo(
    () => (query.data?.pages ?? []).flat(),
    [query.data],
  );

  // Filtro client-side por especialidade do autor (roles é CSV/JSON).
  const filtered = useMemo(() => {
    if (roleFilter === 'all') return allEntries;
    return allEntries.filter((e) =>
      (e.authorRoles ?? '').toLowerCase().includes(roleFilter.toLowerCase()),
    );
  }, [allEntries, roleFilter]);

  // Agrupa por data (dia) pra timeline visual.
  const grouped = useMemo(() => {
    const map = new Map<string, MedicalRecordEntry[]>();
    for (const e of filtered) {
      const key = e.eventDate.slice(0, 10);
      const arr = map.get(key) ?? [];
      arr.push(e);
      map.set(key, arr);
    }
    return Array.from(map.entries());
  }, [filtered]);

  const toggleType = (t: MedicalRecordType) => {
    setSelectedTypes((prev) =>
      prev.includes(t) ? prev.filter((x) => x !== t) : [...prev, t],
    );
  };

  return (
    <div className="container mx-auto space-y-4 py-6">
      <PageHeader
        breadcrumbs={[
          { label: 'Pacientes', href: '/patients' },
          { label: 'Paciente', href: `/patients/${params.id}` },
          { label: 'Prontuário' },
        ]}
        title="Prontuário"
        description="Linha do tempo agregada — anamneses, exames, escores, avaliações, prescrições e consultas."
      />

      {/* Filtros */}
      <div className="space-y-3 rounded-md border bg-muted/30 p-3">
        <div>
          <p className="mb-1 text-xs font-medium uppercase text-muted-foreground">Tipo</p>
          <div className="flex flex-wrap gap-2">
            {(Object.keys(MR_TYPE_LABELS) as MedicalRecordType[]).map((t) => {
              const checked = selectedTypes.includes(t);
              return (
                <button
                  key={t}
                  type="button"
                  onClick={() => toggleType(t)}
                  className={cn(
                    'flex items-center gap-1 rounded-md border px-2 py-1 text-xs transition',
                    checked
                      ? 'border-primary bg-primary/10 text-primary'
                      : 'border-border bg-background hover:bg-accent',
                  )}
                >
                  <Checkbox checked={checked} className="h-3 w-3 pointer-events-none" />
                  {MR_TYPE_LABELS[t]}
                </button>
              );
            })}
          </div>
        </div>
        <div>
          <p className="mb-1 text-xs font-medium uppercase text-muted-foreground">
            Profissional
          </p>
          <div className="flex flex-wrap gap-2">
            <button
              type="button"
              onClick={() => setRoleFilter('all')}
              className={cn(
                'rounded-md border px-2 py-1 text-xs',
                roleFilter === 'all'
                  ? 'border-primary bg-primary/10 text-primary'
                  : 'border-border bg-background hover:bg-accent',
              )}
            >
              Todos
            </button>
            {ROLE_FILTERS.map((r) => (
              <button
                key={r.key}
                type="button"
                onClick={() => setRoleFilter(r.key)}
                className={cn(
                  'rounded-md border px-2 py-1 text-xs',
                  roleFilter === r.key
                    ? 'border-primary bg-primary/10 text-primary'
                    : 'border-border bg-background hover:bg-accent',
                )}
              >
                {r.label}
              </button>
            ))}
          </div>
        </div>
      </div>

      {/* Timeline */}
      {query.isLoading ? (
        <Card className="flex h-40 items-center justify-center">
          <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />
        </Card>
      ) : grouped.length === 0 ? (
        <Card>
          <CardContent className="py-10 text-center text-sm text-muted-foreground">
            <FileText className="mx-auto mb-2 h-8 w-8" />
            Nenhum registro nessa visão.
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-4">
          {grouped.map(([day, entries]) => (
            <div key={day}>
              <p className="mb-2 text-xs font-medium uppercase text-muted-foreground">
                {format(new Date(day), "EEEE, dd 'de' MMM yyyy", { locale: ptBR })}
              </p>
              <div className="space-y-2">
                {entries.map((e) => {
                  const Icon = TYPE_ICON[e.type] ?? FileText;
                  return (
                    <Card
                      key={e.id}
                      className="cursor-pointer transition hover:border-primary"
                      onClick={() => navigateTo(router, e)}
                    >
                      <CardContent className="flex items-start gap-3 p-3">
                        <div className="mt-0.5 rounded-md border bg-muted/40 p-2">
                          <Icon className="h-4 w-4 text-muted-foreground" />
                        </div>
                        <div className="min-w-0 flex-1">
                          <div className="flex flex-wrap items-center gap-2">
                            <Badge variant="outline" className={cn('text-xs', MR_TYPE_COLORS[e.type])}>
                              {MR_TYPE_LABELS[e.type]}
                            </Badge>
                            {e.status && (
                              <span className="text-xs text-muted-foreground">{e.status}</span>
                            )}
                          </div>
                          <p className="mt-1 text-sm font-medium">{e.title}</p>
                          {e.subtitle && (
                            <p className="text-xs text-muted-foreground">{e.subtitle}</p>
                          )}
                          {e.authorName && (
                            <p className="mt-1 text-xs text-muted-foreground">
                              por {e.authorName}
                            </p>
                          )}
                        </div>
                        <p className="shrink-0 text-xs text-muted-foreground">
                          {format(new Date(e.eventDate), 'HH:mm')}
                        </p>
                      </CardContent>
                    </Card>
                  );
                })}
              </div>
            </div>
          ))}

          {query.hasNextPage && (
            <div className="flex justify-center pt-2">
              <Button
                variant="outline"
                onClick={() => query.fetchNextPage()}
                disabled={query.isFetchingNextPage}
              >
                {query.isFetchingNextPage && (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                )}
                Carregar mais
              </Button>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

// Roteamento contextual por tipo de registro.
function navigateTo(router: ReturnType<typeof useRouter>, e: MedicalRecordEntry) {
  switch (e.type) {
    case 'appointment':
      router.push(`/appointments/${e.id}`);
      break;
    case 'anamnesis':
      router.push(`/anamnesis/${e.id}`);
      break;
    case 'lab_batch':
      router.push(`/lab-results/${e.id}`);
      break;
    case 'score_snapshot':
      router.push(`/health-scores/${e.id}`);
      break;
    case 'prescription':
      router.push(`/prescriptions/${e.id}`);
      break;
    case 'physical_assessment':
      router.push(`/training/physical-assessments/${e.id}`);
      break;
    case 'fitness_test':
      router.push(`/training/fitness-tests/${e.id}`);
      break;
    case 'postural_assessment':
      router.push(`/training/posture/${e.id}`);
      break;
  }
}
