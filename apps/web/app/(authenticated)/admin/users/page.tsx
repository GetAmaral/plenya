"use client";

import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { motion } from "framer-motion";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  Users,
  Search,
  Plus,
  ChevronLeft,
  ChevronRight,
  XCircle,
  Shield,
  Stethoscope,
  User as UserIcon,
  Heart,
  Edit,
  Trash2,
} from "lucide-react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { useRequireAuth } from "@/lib/use-auth";
import { PageHeader } from "@/components/layout/page-header";
import { getAllUsers, deleteUser, type User } from "@/lib/api/users";
import { CreateUserForm, EditUserForm } from "@/components/admin/UserForm";
import { toast } from "sonner";
import { ROLES, ROLE_BADGE_COLORS, getRoleColor, getRoleLabel } from "@/lib/roles";
import { cn } from "@/lib/utils";
import type { UserRole } from "@/lib/auth-store";
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
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export default function UsersPage() {
  const { user: currentUser } = useRequireAuth();

  const [searchQuery, setSearchQuery] = useState("");
  const [roleFilter, setRoleFilter] = useState<string>("all");
  const [showCreateForm, setShowCreateForm] = useState(false);
  const [editingUser, setEditingUser] = useState<User | null>(null);
  const [deletingUser, setDeletingUser] = useState<User | null>(null);
  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 20,
  });

  const queryClient = useQueryClient();

  const { data, isLoading, error } = useQuery({
    queryKey: ["users", roleFilter, pagination.pageIndex, pagination.pageSize],
    queryFn: () => getAllUsers({
      role: roleFilter === "all" ? undefined : roleFilter,
      limit: pagination.pageSize,
      offset: pagination.pageIndex * pagination.pageSize,
    }),
  });

  const deleteMutation = useMutation({
    mutationFn: deleteUser,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["users"] });
      toast.success("Usuário deletado com sucesso!");
      setDeletingUser(null);
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || "Erro ao deletar usuário");
    },
  });


  const filteredData = data?.data.filter((user) => {
    if (!searchQuery) return true;
    const query = searchQuery.toLowerCase();
    return (
      user.name.toLowerCase().includes(query) ||
      user.email.toLowerCase().includes(query)
    );
  }) || [];

  const totalPages = data ? Math.ceil(data.total / pagination.pageSize) : 0;

  // Don't render form modals if we're still showing it
  if (showCreateForm) {
    return (
      <div className="container mx-auto py-8 space-y-8">
        <PageHeader
          breadcrumbs={[
            { label: 'Admin', href: '/admin' },
            { label: 'Usuários' }
          ]}
        />
        <CreateUserForm
          onSuccess={() => {
            setShowCreateForm(false);
            queryClient.invalidateQueries({ queryKey: ["users"] });
          }}
          onCancel={() => setShowCreateForm(false)}
        />
      </div>
    );
  }

  if (editingUser) {
    return (
      <div className="container mx-auto py-8 space-y-8">
        <PageHeader
          breadcrumbs={[
            { label: 'Admin', href: '/admin' },
            { label: 'Usuários', href: '/admin/users' },
            { label: 'Editar' }
          ]}
        />
        <EditUserForm
          user={editingUser}
          onSuccess={() => {
            setEditingUser(null);
            queryClient.invalidateQueries({ queryKey: ["users"] });
          }}
          onCancel={() => setEditingUser(null)}
        />
      </div>
    );
  }

  return (
    <div className="container mx-auto py-8 space-y-8">
      {/* Header */}
      <PageHeader
        breadcrumbs={[{ label: 'Admin' }, { label: 'Usuários' }]}
        actions={
          <Button onClick={() => setShowCreateForm(true)}>
            <Plus className="mr-2 h-4 w-4" />
            Novo Usuário
          </Button>
        }
      />

      {/* Filters */}
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ delay: 0.1 }}
        className="flex flex-col sm:flex-row gap-4"
      >
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
          <Input
            placeholder="Buscar por nome ou email..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="pl-9"
          />
        </div>
        <Select value={roleFilter} onValueChange={setRoleFilter}>
          <SelectTrigger className="w-full sm:w-[200px]">
            <SelectValue placeholder="Filtrar por role" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">Todos</SelectItem>
            <SelectItem value="admin">Administradores</SelectItem>
            <SelectItem value="doctor">Médicos</SelectItem>
            <SelectItem value="nurse">Enfermeiros</SelectItem>
            <SelectItem value="patient">Pacientes</SelectItem>
          </SelectContent>
        </Select>
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
              <h3 className="text-lg font-semibold mb-2">Erro ao carregar usuários</h3>
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
          {Array.from({ length: 5 }).map((_, i) => (
            <Card key={i}>
              <CardHeader>
                <Skeleton className="h-6 w-3/4" />
                <Skeleton className="h-4 w-1/2" />
              </CardHeader>
            </Card>
          ))}
        </div>
      )}

      {/* Users List */}
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
                <Users className="mx-auto h-12 w-12 text-muted-foreground mb-4" />
                <h3 className="text-lg font-semibold mb-2">
                  {searchQuery || (roleFilter && roleFilter !== "all") ? "Nenhum usuário encontrado" : "Nenhum usuário cadastrado"}
                </h3>
                <p className="text-sm text-muted-foreground mb-4">
                  {searchQuery || (roleFilter && roleFilter !== "all")
                    ? "Tente ajustar os filtros de busca"
                    : "Comece criando o primeiro usuário"}
                </p>
                {!searchQuery && (!roleFilter || roleFilter === "all") && (
                  <Button onClick={() => setShowCreateForm(true)}>
                    <Plus className="mr-2 h-4 w-4" />
                    Novo Usuário
                  </Button>
                )}
              </CardContent>
            </Card>
          ) : (
            <>
              <div className="space-y-2">
                {filteredData.map((user) => {
                  const isCurrentUser = currentUser?.email === user.email;

                  return (
                    <Card key={user.id} className="hover:shadow-md transition-shadow">
                      <CardHeader>
                        <div className="flex items-start justify-between gap-4">
                          <div className="flex-1 space-y-2">
                            <div className="flex items-center gap-3 flex-wrap">
                              <div className="font-semibold text-lg">{user.name}</div>
                              {(user.roles || []).map((role) => {
                                const roleData = ROLES[role as UserRole];
                                if (!roleData) return null;
                                const RoleIcon = roleData.icon;
                                const colorScheme = getRoleColor(role as UserRole);
                                const badgeClasses = ROLE_BADGE_COLORS[colorScheme as keyof typeof ROLE_BADGE_COLORS];

                                return (
                                  <Badge
                                    key={role}
                                    className={cn("gap-1 border", badgeClasses.active, badgeClasses.border)}
                                    variant="outline"
                                  >
                                    <RoleIcon className="h-3 w-3" />
                                    {getRoleLabel(role as UserRole)}
                                  </Badge>
                                );
                              })}
                              {isCurrentUser && (
                                <Badge variant="outline">Você</Badge>
                              )}
                            </div>
                            <CardDescription className="flex items-center gap-4 flex-wrap text-sm">
                              <span className="flex items-center gap-1">
                                <UserIcon className="h-3 w-3" />
                                {user.email}
                              </span>
                              <span className="text-muted-foreground">
                                Criado em {format(new Date(user.createdAt), "dd/MM/yyyy", { locale: ptBR })}
                              </span>
                            </CardDescription>
                          </div>
                          <div className="flex gap-2">
                            <Button
                              variant="outline"
                              size="sm"
                              onClick={() => setEditingUser(user)}
                            >
                              <Edit className="h-4 w-4 mr-1" />
                              Editar
                            </Button>
                            {!isCurrentUser && (
                              <Button
                                variant="outline"
                                size="sm"
                                onClick={() => setDeletingUser(user)}
                              >
                                <Trash2 className="h-4 w-4 text-destructive" />
                              </Button>
                            )}
                          </div>
                        </div>
                      </CardHeader>
                    </Card>
                  );
                })}
              </div>

              {/* Pagination */}
              {totalPages > 1 && (
                <div className="flex items-center justify-between">
                  <p className="text-sm text-muted-foreground">
                    Mostrando {pagination.pageIndex * pagination.pageSize + 1} até{" "}
                    {Math.min((pagination.pageIndex + 1) * pagination.pageSize, data?.total || 0)} de {data?.total || 0} usuários
                  </p>
                  <div className="flex gap-2">
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={() => setPagination((p) => ({ ...p, pageIndex: p.pageIndex - 1 }))}
                      disabled={pagination.pageIndex === 0}
                    >
                      <ChevronLeft className="h-4 w-4 mr-1" />
                      Anterior
                    </Button>
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={() => setPagination((p) => ({ ...p, pageIndex: p.pageIndex + 1 }))}
                      disabled={pagination.pageIndex >= totalPages - 1}
                    >
                      Próxima
                      <ChevronRight className="h-4 w-4 ml-1" />
                    </Button>
                  </div>
                </div>
              )}
            </>
          )}
        </motion.div>
      )}

      {/* Delete Confirmation Dialog */}
      <AlertDialog open={!!deletingUser} onOpenChange={() => setDeletingUser(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja deletar o usuário <strong>{deletingUser?.name}</strong>?
              Esta ação não pode ser desfeita.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={() => deletingUser && deleteMutation.mutate(deletingUser.id)}
              className="bg-destructive hover:bg-destructive/90"
            >
              Deletar
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  );
}
