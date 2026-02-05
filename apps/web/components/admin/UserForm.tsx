'use client'

import { useState, useRef } from 'react'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { useMutation } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { createUser, updateUser, type User, type CreateUserRequest, type UpdateUserRequest } from '@/lib/api/users'
import { toast } from 'sonner'
import { ROLES_ARRAY, ROLE_BADGE_COLORS, getRoleColor } from '@/lib/roles'
import { cn } from '@/lib/utils'
import { useAuthStore, type UserRole } from '@/lib/auth-store'
import { formatCPF, sanitizeCPF } from '@/lib/utils/cpf'

interface CreateUserFormProps {
  onSuccess: () => void
  onCancel: () => void
}

export function CreateUserForm({ onSuccess, onCancel }: CreateUserFormProps) {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [cpf, setCpf] = useState('')
  const [selectedRoles, setSelectedRoles] = useState<string[]>(['doctor'])

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const createMutation = useMutation({
    mutationFn: createUser,
    onSuccess: () => {
      toast.success('Usuário criado com sucesso!')
      onSuccess()
    },
    onError: (error: any) => {
      if (error.response?.status === 409) {
        toast.error('Este email já está em uso')
      } else {
        toast.error(error.response?.data?.error || 'Erro ao criar usuário')
      }
    },
  })

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    // Validation
    if (!name.trim() || name.length < 3) {
      toast.error('Nome deve ter pelo menos 3 caracteres')
      return
    }

    if (!email.trim() || !email.includes('@')) {
      toast.error('Email inválido')
      return
    }

    if (!password.trim() || password.length < 8) {
      toast.error('Senha deve ter pelo menos 8 caracteres')
      return
    }

    if (password !== confirmPassword) {
      toast.error('As senhas não coincidem')
      return
    }

    if (selectedRoles.length === 0) {
      toast.error('Selecione pelo menos um role')
      return
    }

    const payload: CreateUserRequest = {
      name: name.trim(),
      email: email.trim().toLowerCase(),
      password,
      cpf: sanitizeCPF(cpf) || undefined,  // Remove formatação
      roles: selectedRoles,
    }

    createMutation.mutate(payload)
  }

  const toggleRole = (roleValue: string) => {
    setSelectedRoles(prev =>
      prev.includes(roleValue)
        ? prev.filter(r => r !== roleValue)
        : [...prev, roleValue]
    )
  }

  return (
    <Card className="p-6">
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-xl font-semibold">Criar Novo Usuário</h2>
      </div>

      <form ref={formRef} onSubmit={handleSubmit} className="space-y-4">
        <div>
          <Label htmlFor="name">Nome Completo *</Label>
          <Input
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Dr. João da Silva"
            required
            minLength={3}
            maxLength={200}
          />
        </div>

        <div>
          <Label htmlFor="email">Email *</Label>
          <Input
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="joao@plenya.com"
            required
          />
        </div>

        <div>
          <Label htmlFor="cpf">CPF</Label>
          <Input
            id="cpf"
            value={cpf}
            onChange={(e) => setCpf(formatCPF(e.target.value))}
            placeholder="123.456.789-00"
            maxLength={14}
          />
          <p className="text-xs text-muted-foreground mt-1">
            Usado para validação de certificado digital
          </p>
        </div>

        <div>
          <Label htmlFor="password">Senha *</Label>
          <Input
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Mínimo 8 caracteres"
            required
            minLength={8}
          />
        </div>

        <div>
          <Label htmlFor="confirm-password">Confirmar Senha *</Label>
          <Input
            id="confirm-password"
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            placeholder="Digite a senha novamente"
            required
            minLength={8}
          />
          <p className="text-xs text-muted-foreground mt-1">
            A senha deve ter pelo menos 8 caracteres
          </p>
        </div>

        <div>
          <Label>Roles * (selecione um ou mais)</Label>
          <p className="text-xs text-muted-foreground mt-1 mb-3">
            Clique nos badges para selecionar/desselecionar
          </p>
          <div className="flex flex-wrap gap-2">
            {ROLES_ARRAY.map((role) => {
              const Icon = role.icon
              const isSelected = selectedRoles.includes(role.value)
              const colorScheme = getRoleColor(role.value as UserRole)
              const badgeClasses = ROLE_BADGE_COLORS[colorScheme as keyof typeof ROLE_BADGE_COLORS]

              return (
                <button
                  key={role.value}
                  type="button"
                  onClick={() => toggleRole(role.value)}
                  className={cn(
                    "inline-flex items-center gap-1.5 px-3 py-1.5 rounded-full text-sm font-medium transition-all cursor-pointer hover:scale-105",
                    isSelected
                      ? "border-2 border-lime-500" + " " + badgeClasses.active
                      : "border border-gray-300" + " " + badgeClasses.inactive
                  )}
                >
                  <Icon className="h-3.5 w-3.5" />
                  {role.label}
                </button>
              )
            })}
          </div>
        </div>

        <div className="flex justify-end gap-2 pt-4 border-t">
          <Button type="button" variant="outline" onClick={onCancel}>
            Cancelar
          </Button>
          <Button type="submit" disabled={createMutation.isPending}>
            {createMutation.isPending ? 'Criando...' : 'Criar Usuário'}
          </Button>
        </div>
      </form>
    </Card>
  )
}

interface EditUserFormProps {
  user: User
  onSuccess: () => void
  onCancel: () => void
}

export function EditUserForm({ user, onSuccess, onCancel }: EditUserFormProps) {
  const [name, setName] = useState(user.name)
  const [email, setEmail] = useState(user.email)
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [cpf, setCpf] = useState(formatCPF((user as any).cpf) || '')
  const [showPasswordFields, setShowPasswordFields] = useState(false)
  const [selectedRoles, setSelectedRoles] = useState<string[]>(user.roles || [])

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  // Get logged user and updateUser from Zustand
  const loggedUser = useAuthStore((state) => state.user)
  const updateUserStore = useAuthStore((state) => state.updateUser)

  const updateMutation = useMutation({
    mutationFn: (data: UpdateUserRequest) => updateUser(user.id, data),
    onSuccess: (updatedUser) => {
      toast.success('Usuário atualizado com sucesso!')

      // Se editou o próprio usuário, atualizar Zustand store
      if (loggedUser && user.id === loggedUser.id) {
        updateUserStore(updatedUser)
      }

      onSuccess()
    },
    onError: (error: any) => {
      if (error.response?.status === 409) {
        toast.error('Este email já está em uso')
      } else if (error.response?.status === 403) {
        toast.error('Você não tem permissão para editar este usuário')
      } else {
        toast.error(error.response?.data?.error || 'Erro ao atualizar usuário')
      }
    },
  })

  const toggleRole = (roleValue: string) => {
    setSelectedRoles(prev =>
      prev.includes(roleValue)
        ? prev.filter(r => r !== roleValue)
        : [...prev, roleValue]
    )
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    // Validation
    if (!name.trim() || name.length < 3) {
      toast.error('Nome deve ter pelo menos 3 caracteres')
      return
    }

    if (!email.trim() || !email.includes('@')) {
      toast.error('Email inválido')
      return
    }

    if (password && password.length < 8) {
      toast.error('Senha deve ter pelo menos 8 caracteres')
      return
    }

    if (password && password !== confirmPassword) {
      toast.error('As senhas não coincidem')
      return
    }

    if (selectedRoles.length === 0) {
      toast.error('Selecione pelo menos um role')
      return
    }

    const payload: UpdateUserRequest = {}

    // Only send changed fields
    if (name.trim() !== user.name) {
      payload.name = name.trim()
    }

    if (email.trim().toLowerCase() !== user.email) {
      payload.email = email.trim().toLowerCase()
    }

    if (password.trim()) {
      payload.password = password.trim()
    }

    // Compare CPF sanitizado (sem formatação)
    const sanitizedCurrentCPF = sanitizeCPF(cpf)
    const sanitizedUserCPF = sanitizeCPF((user as any).cpf)
    if (sanitizedCurrentCPF !== sanitizedUserCPF) {
      payload.cpf = sanitizedCurrentCPF || null
    }

    // Check if roles changed
    const rolesChanged = JSON.stringify(selectedRoles.sort()) !== JSON.stringify((user.roles || []).sort())
    if (rolesChanged) {
      payload.roles = selectedRoles
    }

    // Check if anything changed
    if (Object.keys(payload).length === 0) {
      toast.info('Nenhuma alteração detectada')
      return
    }

    updateMutation.mutate(payload)
  }

  return (
    <Card className="p-6">
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-xl font-semibold">Editar Usuário</h2>
      </div>

      <form ref={formRef} onSubmit={handleSubmit} className="space-y-4">
        <div>
          <Label htmlFor="edit-name">Nome Completo *</Label>
          <Input
            id="edit-name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Dr. João da Silva"
            required
            minLength={3}
            maxLength={200}
          />
        </div>

        <div>
          <Label htmlFor="edit-email">Email *</Label>
          <Input
            id="edit-email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="joao@plenya.com"
            required
          />
        </div>

        <div>
          <Label htmlFor="edit-cpf">CPF</Label>
          <Input
            id="edit-cpf"
            value={cpf}
            onChange={(e) => setCpf(formatCPF(e.target.value))}
            placeholder="123.456.789-00"
            maxLength={14}
          />
          <p className="text-xs text-muted-foreground mt-1">
            Usado para validação de certificado digital
          </p>
        </div>

        <div className="border rounded-lg p-4">
          <button
            type="button"
            onClick={() => setShowPasswordFields(!showPasswordFields)}
            className="flex items-center gap-2 text-sm font-medium hover:text-primary transition-colors"
          >
            <svg
              className={cn("h-4 w-4 transition-transform", showPasswordFields && "rotate-90")}
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
            </svg>
            Alterar senha
          </button>

          {showPasswordFields && (
            <div className="mt-4 space-y-4">
              <div>
                <Label htmlFor="edit-password">Nova Senha *</Label>
                <Input
                  id="edit-password"
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  placeholder="Mínimo 8 caracteres"
                  minLength={8}
                />
              </div>

              <div>
                <Label htmlFor="edit-confirm-password">Confirmar Nova Senha *</Label>
                <Input
                  id="edit-confirm-password"
                  type="password"
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                  placeholder="Digite a senha novamente"
                  minLength={8}
                />
              </div>
            </div>
          )}
        </div>

        <div>
          <Label>Roles * (selecione um ou mais)</Label>
          <p className="text-xs text-muted-foreground mt-1 mb-3">
            Clique nos badges para selecionar/desselecionar
          </p>
          <div className="flex flex-wrap gap-2">
            {ROLES_ARRAY.map((role) => {
              const Icon = role.icon
              const isSelected = selectedRoles.includes(role.value)
              const colorScheme = getRoleColor(role.value as UserRole)
              const badgeClasses = ROLE_BADGE_COLORS[colorScheme as keyof typeof ROLE_BADGE_COLORS]

              return (
                <button
                  key={role.value}
                  type="button"
                  onClick={() => toggleRole(role.value)}
                  className={cn(
                    "inline-flex items-center gap-1.5 px-3 py-1.5 rounded-full text-sm font-medium transition-all cursor-pointer hover:scale-105",
                    isSelected
                      ? "border-2 border-lime-500" + " " + badgeClasses.active
                      : "border border-gray-300" + " " + badgeClasses.inactive
                  )}
                >
                  <Icon className="h-3.5 w-3.5" />
                  {role.label}
                </button>
              )
            })}
          </div>
        </div>

        <div className="flex justify-end gap-2 pt-4 border-t">
          <Button type="button" variant="outline" onClick={onCancel}>
            Cancelar
          </Button>
          <Button type="submit" disabled={updateMutation.isPending}>
            {updateMutation.isPending ? 'Salvando...' : 'Salvar Alterações'}
          </Button>
        </div>
      </form>
    </Card>
  )
}
