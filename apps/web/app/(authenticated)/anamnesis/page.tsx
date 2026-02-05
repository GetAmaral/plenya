"use client";

import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { motion } from "framer-motion";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  ClipboardList,
  Search,
  Plus,
  ChevronLeft,
  ChevronRight,
  XCircle,
  Eye,
  EyeOff,
  Lock,
  Calendar,
  FileText,
  User,
} from "lucide-react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";
import { SelectedPatientHeader } from "@/components/patients/SelectedPatientHeader";
import { PageHeader } from "@/components/layout/page-header";
import { getAllAnamnesis, type Anamnesis } from "@/lib/api/anamnesis";
import { CreateAnamnesisForm, EditAnamnesisForm } from "@/components/anamnesis/AnamnesisForm";

type AnamnesisVisibility = "all" | "medicalOnly" | "psychOnly" | "authorOnly";

export default function AnamnesisPage() {
  useRequireAuth();
  const { selectedPatient } = useRequireSelectedPatient();

  const [searchQuery, setSearchQuery] = useState("");
  const [showCreateForm, setShowCreateForm] = useState(false);
  const [editingAnamnesis, setEditingAnamnesis] = useState<Anamnesis | null>(null);
  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 10,
  });

  const { data, isLoading, error } = useQuery({
    queryKey: ["anamnesis", selectedPatient?.id, pagination.pageIndex, pagination.pageSize],
    queryFn: () => getAllAnamnesis({
      limit: pagination.pageSize,
      offset: pagination.pageIndex * pagination.pageSize,
    }),
    enabled: !!selectedPatient,
  });

  const getVisibilityLabel = (visibility: AnamnesisVisibility) => {
    switch (visibility) {
      case "all":
        return "Todos";
      case "medicalOnly":
        return "Apenas Médicos";
      case "psychOnly":
        return "Apenas Psicólogos";
      case "authorOnly":
        return "Apenas Autor";
      default:
        return visibility;
    }
  };

  const isContentRestricted = (anamnesis: Anamnesis) => {
    return anamnesis.summary === "Conteúdo restrito" || anamnesis.content === "Conteúdo restrito";
  };

  const getVisibilityIcon = (visibility: AnamnesisVisibility) => {
    switch (visibility) {
      case "all":
        return Eye;
      case "medicalOnly":
      case "psychOnly":
        return Lock;
      default:
        return EyeOff;
    }
  };

  const filteredData = data?.data.filter((anamnesis) => {
    if (!searchQuery) return true;
    const query = searchQuery.toLowerCase();
    return (
      anamnesis.summary?.toLowerCase().includes(query) ||
      anamnesis.content?.toLowerCase().includes(query) ||
      anamnesis.notes?.toLowerCase().includes(query) ||
      anamnesis.anamnesisTemplate?.name.toLowerCase().includes(query)
    );
  }) || [];

  return (
    <div className="container mx-auto py-8 space-y-8">
      {/* Selected Patient Header */}
      <SelectedPatientHeader />

      {/* Header */}
      <PageHeader
        breadcrumbs={[{ label: 'Anamneses' }]}
        title="Anamneses"
        description={`${data?.total || 0} ${data?.total === 1 ? 'anamnese registrada' : 'anamneses registradas'}`}
        actions={[
          {
            label: 'Nova Anamnese',
            icon: <Plus className="h-4 w-4" />,
            onClick: () => setShowCreateForm(!showCreateForm),
            variant: 'default',
          },
        ]}
      />

      {/* Create Form */}
      {showCreateForm && (
        <CreateAnamnesisForm onSuccess={() => setShowCreateForm(false)} />
      )}

      {/* Edit Form */}
      {editingAnamnesis && (
        <EditAnamnesisForm
          anamnesis={editingAnamnesis}
          onSuccess={() => setEditingAnamnesis(null)}
          onCancel={() => setEditingAnamnesis(null)}
        />
      )}

      {/* Search */}
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.1 }}
      >
        <div className="relative">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Buscar por resumo, conteúdo, observações ou template..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="pl-9"
          />
        </div>
      </motion.div>

      {/* Error State */}
      {error && (
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
        >
          <Card className="border-destructive/50 bg-destructive/10">
            <CardContent className="p-8 text-center">
              <XCircle className="mx-auto h-12 w-12 text-destructive mb-4" />
              <h3 className="text-lg font-semibold mb-2">Erro ao carregar anamneses</h3>
              <p className="text-sm text-muted-foreground mb-4">
                {error instanceof Error ? error.message : "Não foi possível conectar ao servidor"}
              </p>
              <Button variant="outline" onClick={() => window.location.reload()}>
                Tentar novamente
              </Button>
            </CardContent>
          </Card>
        </motion.div>
      )}

      {/* Loading State */}
      {isLoading && (
        <div className="space-y-4">
          {Array.from({ length: 3 }).map((_, i) => (
            <Card key={i}>
              <CardHeader>
                <Skeleton className="h-6 w-3/4" />
                <Skeleton className="h-4 w-1/2" />
              </CardHeader>
              <CardContent>
                <Skeleton className="h-20 w-full" />
              </CardContent>
            </Card>
          ))}
        </div>
      )}

      {/* Anamnesis Cards */}
      {!isLoading && !error && (
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.2 }}
          className="space-y-4"
        >
          {filteredData.length === 0 ? (
            <Card>
              <CardContent className="p-12 text-center">
                <ClipboardList className="mx-auto h-12 w-12 text-muted-foreground mb-4" />
                <h3 className="text-lg font-semibold mb-2">
                  {searchQuery ? "Nenhuma anamnese encontrada" : "Nenhuma anamnese registrada"}
                </h3>
                <p className="text-sm text-muted-foreground mb-4">
                  {searchQuery
                    ? "Tente ajustar os termos da busca"
                    : "Comece criando sua primeira anamnese para este paciente"}
                </p>
                {!searchQuery && (
                  <Button onClick={() => setShowCreateForm(true)}>
                    <Plus className="mr-2 h-4 w-4" />
                    Nova Anamnese
                  </Button>
                )}
              </CardContent>
            </Card>
          ) : (
            filteredData.map((anamnesis) => {
              const VisibilityIcon = getVisibilityIcon(anamnesis.visibility);
              const consultationDate = new Date(anamnesis.consultationDate);

              return (
                <Card key={anamnesis.id} className="hover:shadow-md transition-shadow">
                  <CardHeader>
                    <div className="flex items-start justify-between gap-4">
                      <div className="flex-1 space-y-2">
                        <div
                          className="html-content font-semibold"
                          dangerouslySetInnerHTML={{
                            __html: anamnesis.summaryHtml || anamnesis.summary || "Sem resumo"
                          }}
                        />
                        <CardDescription className="flex items-center gap-4 flex-wrap text-sm">
                          <span className="flex items-center gap-1">
                            <Calendar className="h-3 w-3" />
                            {format(consultationDate, "dd/MM/yyyy 'às' HH:mm", { locale: ptBR })}
                          </span>
                          {anamnesis.anamnesisTemplate && (
                            <span className="flex items-center gap-1">
                              <FileText className="h-3 w-3" />
                              {anamnesis.anamnesisTemplate.name}
                            </span>
                          )}
                          {anamnesis.author && (
                            <span className="flex items-center gap-1">
                              <User className="h-3 w-3" />
                              {anamnesis.author.name}
                            </span>
                          )}
                        </CardDescription>
                      </div>
                      <div className="flex items-center gap-2">
                        {anamnesis.visibility !== "all" && (
                          <Badge variant="outline" className="gap-1">
                            <VisibilityIcon className="h-3 w-3" />
                            {getVisibilityLabel(anamnesis.visibility)}
                          </Badge>
                        )}
                        {!isContentRestricted(anamnesis) && (
                          <Button
                            variant="outline"
                            size="sm"
                            onClick={() => setEditingAnamnesis(anamnesis)}
                          >
                            Editar
                          </Button>
                        )}
                      </div>
                    </div>
                  </CardHeader>
                  {(anamnesis.contentHtml || anamnesis.content || anamnesis.notes || (anamnesis.items && anamnesis.items.length > 0)) && (
                    <CardContent className="space-y-3">
                      {isContentRestricted(anamnesis) ? (
                        <div className="flex items-center gap-2 p-3 rounded-md bg-muted/50 border border-muted">
                          <Lock className="h-4 w-4 text-muted-foreground" />
                          <p className="text-sm text-muted-foreground italic">
                            Conteúdo restrito
                          </p>
                        </div>
                      ) : (
                        <>
                          {(anamnesis.contentHtml || anamnesis.content) && (
                            <div>
                              <h4 className="text-sm font-medium mb-1">Conteúdo:</h4>
                              <div
                                className="text-sm text-muted-foreground line-clamp-3 html-content"
                                style={{ whiteSpace: 'pre-wrap' }}
                                dangerouslySetInnerHTML={{
                                  __html: anamnesis.contentHtml || anamnesis.content || ''
                                }}
                              />
                            </div>
                          )}
                          {anamnesis.notes && (
                            <div>
                              <h4 className="text-sm font-medium mb-1">Observações:</h4>
                              <p className="text-sm text-muted-foreground line-clamp-2 whitespace-pre-wrap">
                                {anamnesis.notes}
                              </p>
                            </div>
                          )}
                        </>
                      )}
                      {anamnesis.items && anamnesis.items.length > 0 && (() => {
                        // Organize items by group and subgroup
                        const itemsByGroup = new Map<string, {
                          group: any,
                          subgroups: Map<string, {
                            subgroup: any,
                            items: any[]
                          }>
                        }>();

                        anamnesis.items
                          .filter(item => item.scoreItem?.subgroup?.group)
                          .forEach(item => {
                            const group = item.scoreItem!.subgroup!.group!;
                            const subgroup = item.scoreItem!.subgroup!;

                            if (!itemsByGroup.has(group.id)) {
                              itemsByGroup.set(group.id, {
                                group,
                                subgroups: new Map()
                              });
                            }

                            const groupData = itemsByGroup.get(group.id)!;

                            if (!groupData.subgroups.has(subgroup.id)) {
                              groupData.subgroups.set(subgroup.id, {
                                subgroup,
                                items: []
                              });
                            }

                            groupData.subgroups.get(subgroup.id)!.items.push(item);
                          });

                        const LEVEL_STYLES = {
                          0: { bg: 'bg-red-100', text: 'text-red-900', border: 'border-red-500' },
                          1: { bg: 'bg-orange-100', text: 'text-orange-900', border: 'border-orange-500' },
                          2: { bg: 'bg-yellow-100', text: 'text-yellow-900', border: 'border-yellow-500' },
                          3: { bg: 'bg-blue-100', text: 'text-blue-900', border: 'border-blue-500' },
                          4: { bg: 'bg-green-100', text: 'text-green-900', border: 'border-green-500' },
                          5: { bg: 'bg-emerald-100', text: 'text-emerald-900', border: 'border-emerald-500' },
                          6: { bg: 'bg-gray-100', text: 'text-gray-900', border: 'border-gray-500' },
                        };

                        return (
                          <div>
                            <h4 className="text-sm font-medium mb-3">Items Preenchidos:</h4>
                            <div className="space-y-3">
                              {Array.from(itemsByGroup.values()).map(({ group, subgroups }) => (
                                <Card key={group.id} className="overflow-hidden border-gray-200">
                                  <CardHeader className="bg-gray-50 pb-2">
                                    <CardTitle className="text-sm font-semibold text-gray-900">
                                      {group.name}
                                    </CardTitle>
                                  </CardHeader>
                                  <CardContent className="pt-3 space-y-3">
                                    {Array.from(subgroups.values()).map(({ subgroup, items }) => (
                                      <div key={subgroup.id} className="space-y-2">
                                        <h5 className="text-xs font-medium text-muted-foreground">
                                          {subgroup.name}
                                        </h5>
                                        <div className="space-y-2">
                                          {items.map((item: any) => {
                                            const scoreItem = item.scoreItem;
                                            const selectedLevel = scoreItem.levels?.find((l: any) => l.level === item.numericValue);
                                            const levelStyle = selectedLevel ? LEVEL_STYLES[selectedLevel.level as keyof typeof LEVEL_STYLES] || LEVEL_STYLES[6] : null;

                                            return (
                                              <div key={item.id} className="text-xs border rounded-md p-2 bg-card">
                                                <div className="flex items-start gap-2">
                                                  <div className="flex-1">
                                                    <span className="font-semibold">{scoreItem.name}</span>
                                                    {item.textValue && (
                                                      <span className="font-normal text-muted-foreground">
                                                        {': '}{item.textValue}
                                                      </span>
                                                    )}
                                                  </div>
                                                  {selectedLevel && levelStyle && (
                                                    <span className={`text-xs px-2 py-0.5 rounded-full border-2 font-bold whitespace-nowrap ${levelStyle.bg} ${levelStyle.text} ${levelStyle.border}`}>
                                                      N{selectedLevel.level}: {selectedLevel.name}
                                                    </span>
                                                  )}
                                                </div>
                                              </div>
                                            );
                                          })}
                                        </div>
                                      </div>
                                    ))}
                                  </CardContent>
                                </Card>
                              ))}
                            </div>
                          </div>
                        );
                      })()}
                    </CardContent>
                  )}
                </Card>
              );
            })
          )}
        </motion.div>
      )}

      {/* Pagination */}
      {!isLoading && !error && filteredData.length > 0 && (
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ delay: 0.3 }}
          className="flex items-center justify-between pt-4"
        >
          <div className="text-sm text-muted-foreground">
            Mostrando {filteredData.length} de {data?.total || 0} anamneses
          </div>
          <div className="flex gap-2">
            <Button
              variant="outline"
              size="sm"
              onClick={() => setPagination((p) => ({ ...p, pageIndex: p.pageIndex - 1 }))}
              disabled={pagination.pageIndex === 0}
            >
              <ChevronLeft className="h-4 w-4" />
              Anterior
            </Button>
            <Button
              variant="outline"
              size="sm"
              onClick={() => setPagination((p) => ({ ...p, pageIndex: p.pageIndex + 1 }))}
              disabled={
                !data?.total ||
                (pagination.pageIndex + 1) * pagination.pageSize >= data.total
              }
            >
              Próxima
              <ChevronRight className="h-4 w-4" />
            </Button>
          </div>
        </motion.div>
      )}
    </div>
  );
}
