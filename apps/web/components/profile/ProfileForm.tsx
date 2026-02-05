'use client'

import { useState, useRef } from 'react'
import { useRouter } from 'next/navigation'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { useMutation } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import { toast } from 'sonner'
import { isGranted, type UserRole } from '@/lib/auth-store'
import type { User } from '@/lib/api/users'
import { updateProfile } from '@/lib/api/profile'
import { Badge } from '@/components/ui/badge'
import { ROLES, ROLE_BADGE_COLORS, getRoleColor, getRoleLabel } from '@/lib/roles'
import { cn } from '@/lib/utils'
import { formatCPF, sanitizeCPF } from '@/lib/utils/cpf'

interface ProfileFormProps {
  user: User
  onSuccess: (user: User) => void
}

export function ProfileForm({ user, onSuccess }: ProfileFormProps) {
  const router = useRouter()
  const [name, setName] = useState(user.name || '')
  const [cpf, setCpf] = useState(formatCPF((user as any).cpf) || '')
  const [professionalPhone, setProfessionalPhone] = useState((user as any).professionalPhone || '')
  const [professionalAddress, setProfessionalAddress] = useState((user as any).professionalAddress || '')
  const [specialty, setSpecialty] = useState((user as any).specialty || '')
  const [crm, setCrm] = useState((user as any).crm || '')
  const [crmUF, setCrmUF] = useState((user as any).crmUF || '')
  const [rqe, setRqe] = useState((user as any).rqe || '')

  const formRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef })

  const isDoctor = isGranted(user as any, 'doctor')

  const handleCancel = () => {
    router.back()
  }

  const updateMutation = useMutation({
    mutationFn: updateProfile,
    onSuccess: (updatedUser) => {
      toast.success('Perfil atualizado com sucesso!')
      onSuccess(updatedUser)
      router.push('/dashboard')
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao atualizar perfil')
    },
  })

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    // Validação
    if (!name.trim() || name.length < 3) {
      toast.error('Nome deve ter pelo menos 3 caracteres')
      return
    }


    if (isDoctor && crmUF && crmUF.length !== 2) {
      toast.error('UF do CRM deve ter 2 caracteres')
      return
    }

    const payload: any = {
      name: name.trim(),
      cpf: sanitizeCPF(cpf) || null,  // Remove formatação antes de enviar
      professionalPhone: professionalPhone.trim() || null,
      professionalAddress: professionalAddress.trim() || null,
    }

    if (isDoctor) {
      payload.specialty = specialty.trim() || null
      payload.crm = crm.trim() || null
      payload.crmUF = crmUF.trim().toUpperCase() || null
      payload.rqe = rqe.trim() || null
    }

    updateMutation.mutate(payload)
  }

  return (
    <Card className="p-6">
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-xl font-semibold">Meus Dados</h2>
      </div>

      <form ref={formRef} onSubmit={handleSubmit} className="space-y-4">
        {/* Email (readonly) */}
        <div>
          <Label htmlFor="email">Email</Label>
          <Input
            id="email"
            type="email"
            value={user.email}
            disabled
            className="bg-muted"
          />
          <p className="text-xs text-muted-foreground mt-1">
            O email não pode ser alterado
          </p>
        </div>

        {/* Roles (readonly) */}
        <div>
          <Label>Funções no Sistema</Label>
          <div className="flex flex-wrap gap-2 mt-2 p-3 rounded-md bg-muted">
            {(user as any).roles?.map((role: string) => {
              const roleData = ROLES[role as UserRole];
              if (!roleData) return null;
              const Icon = roleData.icon;
              const colorScheme = getRoleColor(role as UserRole);
              const badgeClasses = ROLE_BADGE_COLORS[colorScheme as keyof typeof ROLE_BADGE_COLORS];

              return (
                <Badge
                  key={role}
                  variant="outline"
                  className={cn("gap-1.5 border", badgeClasses.active, badgeClasses.border)}
                >
                  <Icon className="h-3.5 w-3.5" />
                  {getRoleLabel(role as UserRole)}
                </Badge>
              );
            })}
          </div>
          <p className="text-xs text-muted-foreground mt-1">
            Suas funções são definidas por um administrador
          </p>
        </div>

        {/* Nome */}
        <div>
          <Label htmlFor="name">Nome Completo *</Label>
          <Input
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Seu nome completo"
            required
            minLength={3}
            maxLength={200}
          />
        </div>

        {/* CPF */}
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

        {/* Telefone Profissional */}
        <div>
          <Label htmlFor="professionalPhone">Telefone Profissional</Label>
          <Input
            id="professionalPhone"
            type="tel"
            value={professionalPhone}
            onChange={(e) => setProfessionalPhone(e.target.value)}
            placeholder="(11) 98765-4321"
            maxLength={20}
          />
        </div>

        {/* Endereço Profissional */}
        <div>
          <Label htmlFor="professionalAddress">Endereço Profissional</Label>
          <Textarea
            id="professionalAddress"
            value={professionalAddress}
            onChange={(e) => setProfessionalAddress(e.target.value)}
            placeholder="Rua, número, complemento, bairro, cidade, estado, CEP"
            rows={3}
          />
        </div>

        {/* Campos específicos para médicos */}
        {isDoctor && (
          <>
            <div className="border-t pt-4 mt-6">
              <h3 className="text-lg font-medium mb-4">Dados Profissionais (Médico)</h3>
            </div>

            <div>
              <Label htmlFor="specialty">Especialidade</Label>
              <Input
                id="specialty"
                value={specialty}
                onChange={(e) => setSpecialty(e.target.value)}
                placeholder="Ex: Cardiologia"
                maxLength={100}
              />
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div className="md:col-span-2">
                <Label htmlFor="crm">CRM</Label>
                <Input
                  id="crm"
                  value={crm}
                  onChange={(e) => setCrm(e.target.value)}
                  placeholder="123456"
                  maxLength={20}
                />
              </div>

              <div>
                <Label htmlFor="crmUF">UF</Label>
                <Input
                  id="crmUF"
                  value={crmUF}
                  onChange={(e) => setCrmUF(e.target.value.toUpperCase())}
                  placeholder="SP"
                  maxLength={2}
                />
              </div>
            </div>

            <div>
              <Label htmlFor="rqe">RQE (Registro de Qualificação de Especialista)</Label>
              <Input
                id="rqe"
                value={rqe}
                onChange={(e) => setRqe(e.target.value)}
                placeholder="RQE-12345"
                maxLength={20}
              />
            </div>
          </>
        )}

        <div className="flex justify-end gap-2 pt-4 border-t">
          <Button type="button" variant="outline" onClick={handleCancel}>
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
