'use client'

import { ReactNode } from 'react'
import { useRouter } from 'next/navigation'
import { ChevronRight, Home } from 'lucide-react'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'

// Tipos para breadcrumbs
interface Breadcrumb {
  label: string
  href?: string
  icon?: ReactNode
}

// Tipos para ações simples (botões)
interface PageHeaderAction {
  label: string
  icon: ReactNode
  onClick: () => void
  variant?: 'default' | 'outline' | 'ghost'
  disabled?: boolean
  loading?: boolean
}

// Tipos para ações com dropdown
interface PageHeaderDropdown {
  label: string
  icon: ReactNode
  variant?: 'default' | 'outline' | 'ghost'
  disabled?: boolean
  items: {
    label: string
    icon?: ReactNode
    onClick: () => void
    disabled?: boolean
  }[]
}

interface PageHeaderProps {
  // Navegação
  breadcrumbs?: Breadcrumb[]

  // Título e contexto
  title: string
  description?: string

  // Ações (escolha uma das opções)
  actions?: (PageHeaderAction | PageHeaderDropdown)[]
  children?: ReactNode // Para dropdowns customizados complexos
}

// Type guard para verificar se é dropdown
function isDropdown(action: PageHeaderAction | PageHeaderDropdown): action is PageHeaderDropdown {
  return 'items' in action
}

export function PageHeader({
  breadcrumbs,
  title,
  description,
  actions = [],
  children,
}: PageHeaderProps) {
  const router = useRouter()

  return (
    <div className="space-y-8">
      {/* Breadcrumbs - Acima do header principal */}
      {breadcrumbs && breadcrumbs.length > 0 && (
        <nav className="flex items-center gap-2 text-xs text-muted-foreground">
          {/* Home sempre primeiro */}
          <button
            onClick={() => router.push('/dashboard')}
            className="hover:text-foreground transition-colors"
            aria-label="Voltar para Dashboard"
          >
            <Home className="h-3.5 w-3.5" />
          </button>

          {breadcrumbs.map((crumb, index) => (
            <div key={index} className="flex items-center gap-2">
              <ChevronRight className="h-3 w-3" />
              {crumb.href ? (
                <button
                  onClick={() => router.push(crumb.href!)}
                  className="hover:text-foreground transition-colors flex items-center gap-1.5"
                >
                  {crumb.icon}
                  <span>{crumb.label}</span>
                </button>
              ) : (
                <span className="text-foreground flex items-center gap-1.5">
                  {crumb.icon}
                  <span>{crumb.label}</span>
                </span>
              )}
            </div>
          ))}
        </nav>
      )}

      {/* Header Principal - Layout Distribuído Esquerda-Direita */}
      <div className="flex items-center justify-between gap-8">
        {/* LEFT SIDE - Título e Contexto */}
        <div className="flex-1 min-w-0">
          <h1 className="text-4xl font-bold tracking-tight">{title}</h1>
          {description && (
            <p className="text-muted-foreground mt-2">{description}</p>
          )}
        </div>

        {/* RIGHT SIDE - Ações */}
        {(actions.length > 0 || children) && (
          <div className="flex items-center gap-3 flex-shrink-0">
            {/* Renderizar ações fornecidas via prop */}
            {actions.map((action, index) => {
              // Se for dropdown
              if (isDropdown(action)) {
                return (
                  <DropdownMenu key={index}>
                    <DropdownMenuTrigger asChild>
                      <Button
                        variant={action.variant || 'ghost'}
                        size="default"
                        disabled={action.disabled}
                        className="gap-2"
                      >
                        {action.icon}
                        <span className="hidden sm:inline">{action.label}</span>
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      {action.items.map((item, itemIndex) => (
                        <DropdownMenuItem
                          key={itemIndex}
                          onClick={item.onClick}
                          disabled={item.disabled}
                          className="cursor-pointer"
                        >
                          {item.icon && <span className="mr-2">{item.icon}</span>}
                          {item.label}
                        </DropdownMenuItem>
                      ))}
                    </DropdownMenuContent>
                  </DropdownMenu>
                )
              }

              // Se for botão simples
              return (
                <Button
                  key={index}
                  onClick={action.onClick}
                  variant={action.variant || 'ghost'}
                  size="default"
                  disabled={action.disabled || action.loading}
                  className="gap-2"
                >
                  {action.icon}
                  <span className="hidden sm:inline">{action.label}</span>
                </Button>
              )
            })}

            {/* Renderizar children (para casos complexos customizados) */}
            {children}
          </div>
        )}
      </div>
    </div>
  )
}
