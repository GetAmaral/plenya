import { Shield, Stethoscope, Heart, User as UserIcon, Apple, Brain, Activity, Briefcase, Users } from 'lucide-react'
import type { LucideIcon } from 'lucide-react'
import type { UserRole } from './auth-store'

export interface RoleDefinition {
  value: UserRole
  label: string
  icon: LucideIcon
  color: 'red' | 'purple' | 'blue' | 'green' | 'yellow' | 'orange' | 'pink' | 'cyan' | 'indigo'
  description: string
}

/**
 * Definição central de todos os roles do sistema
 * IMPORTANTE: Usar este objeto em todas as partes do sistema que precisam exibir roles
 */
export const ROLES: Record<UserRole, RoleDefinition> = {
  admin: {
    value: 'admin',
    label: 'Administrador',
    icon: Shield,
    color: 'red',
    description: 'Acesso total ao sistema',
  },
  manager: {
    value: 'manager',
    label: 'Gerente',
    icon: Briefcase,
    color: 'purple',
    description: 'Gerenciamento de equipe e relatórios',
  },
  doctor: {
    value: 'doctor',
    label: 'Médico',
    icon: Stethoscope,
    color: 'blue',
    description: 'Atendimento médico e prescrições',
  },
  nurse: {
    value: 'nurse',
    label: 'Enfermeiro',
    icon: Heart,
    color: 'pink',
    description: 'Cuidados de enfermagem',
  },
  nutritionist: {
    value: 'nutritionist',
    label: 'Nutricionista',
    icon: Apple,
    color: 'green',
    description: 'Orientação nutricional',
  },
  psychologist: {
    value: 'psychologist',
    label: 'Psicólogo',
    icon: Brain,
    color: 'purple',
    description: 'Atendimento psicológico',
  },
  physicalEducator: {
    value: 'physicalEducator',
    label: 'Educador Físico',
    icon: Activity,
    color: 'orange',
    description: 'Orientação de atividades físicas',
  },
  secretary: {
    value: 'secretary',
    label: 'Secretário',
    icon: Users,
    color: 'cyan',
    description: 'Atendimento e agendamento',
  },
  patient: {
    value: 'patient',
    label: 'Paciente',
    icon: UserIcon,
    color: 'indigo',
    description: 'Acesso aos próprios dados',
  },
}

/**
 * Array de roles na ordem de prioridade/privilégio
 */
export const ROLES_ARRAY: RoleDefinition[] = [
  ROLES.admin,
  ROLES.manager,
  ROLES.doctor,
  ROLES.nurse,
  ROLES.nutritionist,
  ROLES.psychologist,
  ROLES.physicalEducator,
  ROLES.secretary,
  ROLES.patient,
]

/**
 * Retorna a definição de um role
 */
export function getRoleDefinition(role: UserRole): RoleDefinition {
  return ROLES[role]
}

/**
 * Retorna o label de um role
 */
export function getRoleLabel(role: UserRole): string {
  return ROLES[role]?.label || role
}

/**
 * Retorna a cor de um role
 */
export function getRoleColor(role: UserRole): string {
  return ROLES[role]?.color || 'slate'
}

/**
 * Classes Tailwind para cores de badges
 */
export const ROLE_BADGE_COLORS = {
  red: {
    active: 'bg-red-100 text-red-700 dark:bg-red-950 dark:text-red-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-red-300 dark:border-red-800',
  },
  purple: {
    active: 'bg-purple-100 text-purple-700 dark:bg-purple-950 dark:text-purple-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-purple-300 dark:border-purple-800',
  },
  blue: {
    active: 'bg-blue-100 text-blue-700 dark:bg-blue-950 dark:text-blue-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-blue-300 dark:border-blue-800',
  },
  green: {
    active: 'bg-green-100 text-green-700 dark:bg-green-950 dark:text-green-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-green-300 dark:border-green-800',
  },
  yellow: {
    active: 'bg-yellow-100 text-yellow-700 dark:bg-yellow-950 dark:text-yellow-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-yellow-300 dark:border-yellow-800',
  },
  orange: {
    active: 'bg-orange-100 text-orange-700 dark:bg-orange-950 dark:text-orange-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-orange-300 dark:border-orange-800',
  },
  pink: {
    active: 'bg-pink-100 text-pink-700 dark:bg-pink-950 dark:text-pink-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-pink-300 dark:border-pink-800',
  },
  cyan: {
    active: 'bg-cyan-100 text-cyan-700 dark:bg-cyan-950 dark:text-cyan-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-cyan-300 dark:border-cyan-800',
  },
  indigo: {
    active: 'bg-indigo-100 text-indigo-700 dark:bg-indigo-950 dark:text-indigo-300',
    inactive: 'bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-400',
    border: 'border-indigo-300 dark:border-indigo-800',
  },
}
