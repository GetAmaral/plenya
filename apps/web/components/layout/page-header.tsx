'use client'

import { useEffect, ReactNode } from 'react'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { usePageHeader } from '@/lib/page-context'

// Types kept for backwards-compat with all existing pages
interface Breadcrumb {
  label: string
  href?: string
  icon?: ReactNode
}

interface PageHeaderAction {
  label: string
  icon: ReactNode
  onClick: () => void
  variant?: 'default' | 'outline' | 'ghost'
  disabled?: boolean
  loading?: boolean
}

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
  breadcrumbs?: Breadcrumb[]  // still accepted, used to set TopBar title
  title?: string
  description?: string
  icon?: ReactNode | React.ComponentType<any>  // aceito p/ compat (ícone exibido pelo TopBar)
  // array estruturado (renderizado pelo componente) OU JSX livre (legado)
  actions?: (PageHeaderAction | PageHeaderDropdown)[] | ReactNode
  children?: ReactNode
}

function isDropdown(action: PageHeaderAction | PageHeaderDropdown): action is PageHeaderDropdown {
  return 'items' in action
}

export function PageHeader({
  breadcrumbs,
  title,
  description,
  actions,
  children,
}: PageHeaderProps) {
  const { setPageInfo } = usePageHeader()

  // actions pode ser o array estruturado ou JSX livre (legado)
  const actionItems = Array.isArray(actions) ? actions : []
  const customActions = Array.isArray(actions) ? null : actions

  // Push title + breadcrumbs to TopBar
  useEffect(() => {
    setPageInfo({ title: title ?? '', breadcrumbs: breadcrumbs ?? [] })
    return () => setPageInfo({ title: '', breadcrumbs: [] })
  }, [title, breadcrumbs, setPageInfo])

  // If nothing to render in body, return null
  if (!description && actionItems.length === 0 && !customActions && !children) return null

  return (
    <div className="flex items-center justify-between gap-4">
      {/* Description (count, context info) */}
      <div className="min-w-0">
        {description && (
          <p className="text-sm text-muted-foreground">{description}</p>
        )}
      </div>

      {/* Actions */}
      {(actionItems.length > 0 || customActions || children) && (
        <div className="flex items-center gap-3 shrink-0">
          {actionItems.map((action, index) => {
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

          {customActions}
          {children}
        </div>
      )}
    </div>
  )
}
