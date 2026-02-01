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

type AnamnesisVisibility = "all" | "medicalOnly" | "psychOnly";

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
      default:
        return visibility;
    }
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
                        <div className="flex items-center gap-2 flex-wrap">
                          <div
                            className="html-content font-semibold"
                            dangerouslySetInnerHTML={{
                              __html: anamnesis.summaryHtml || anamnesis.summary || "Sem resumo"
                            }}
                          />
                          <Badge variant="outline" className="gap-1">
                            <VisibilityIcon className="h-3 w-3" />
                            {getVisibilityLabel(anamnesis.visibility)}
                          </Badge>
                        </div>
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
                              {anamnesis.author.email}
                            </span>
                          )}
                        </CardDescription>
                      </div>
                      <div className="flex gap-2">
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => setEditingAnamnesis(anamnesis)}
                        >
                          Editar
                        </Button>
                      </div>
                    </div>
                  </CardHeader>
                  {(anamnesis.contentHtml || anamnesis.content || anamnesis.notes) && (
                    <CardContent className="space-y-3">
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
