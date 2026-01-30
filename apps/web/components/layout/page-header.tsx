'use client'

import { ReactNode } from 'react'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { MoreVertical } from 'lucide-react'

interface PageHeaderAction {
  label: string
  icon?: ReactNode
  onClick: () => void
  variant?: 'default' | 'outline' | 'ghost' | 'destructive'
  disabled?: boolean
}

interface PageHeaderProps {
  title: string
  description?: string
  primaryAction?: PageHeaderAction
  secondaryActions?: PageHeaderAction[]
  children?: ReactNode
}

export function PageHeader({
  title,
  description,
  primaryAction,
  secondaryActions = [],
  children,
}: PageHeaderProps) {
  return (
    <div className="mb-6">
      {/* Mobile/Tablet: Stack layout */}
      <div className="flex flex-col gap-4 lg:hidden">
        {/* Title section */}
        <div>
          <h1 className="text-2xl font-bold tracking-tight">{title}</h1>
          {description && (
            <p className="text-sm text-muted-foreground mt-1">{description}</p>
          )}
        </div>

        {/* Actions section */}
        {(primaryAction || secondaryActions.length > 0 || children) && (
          <div className="flex items-center gap-2 flex-wrap">
            {/* Primary action - always visible */}
            {primaryAction && (
              <Button
                onClick={primaryAction.onClick}
                variant={primaryAction.variant || 'default'}
                disabled={primaryAction.disabled}
                size="sm"
              >
                {primaryAction.icon}
                {primaryAction.label}
              </Button>
            )}

            {/* Custom children */}
            {children}

            {/* Secondary actions in dropdown for mobile */}
            {secondaryActions.length > 0 && (
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="outline" size="sm">
                    <MoreVertical className="h-4 w-4" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end" className="w-48">
                  {secondaryActions.map((action, index) => (
                    <DropdownMenuItem
                      key={index}
                      onClick={action.onClick}
                      disabled={action.disabled}
                      className="flex items-center gap-2"
                    >
                      {action.icon}
                      <span>{action.label}</span>
                    </DropdownMenuItem>
                  ))}
                </DropdownMenuContent>
              </DropdownMenu>
            )}
          </div>
        )}
      </div>

      {/* Desktop: Horizontal layout */}
      <div className="hidden lg:flex lg:items-start lg:justify-between lg:gap-4">
        {/* Title section */}
        <div className="min-w-0 flex-1">
          <h1 className="text-3xl font-bold tracking-tight">{title}</h1>
          {description && (
            <p className="text-sm text-muted-foreground mt-2">{description}</p>
          )}
        </div>

        {/* Actions section */}
        {(primaryAction || secondaryActions.length > 0 || children) && (
          <div className="flex items-center gap-2 flex-shrink-0 flex-wrap justify-end">
            {/* Custom children */}
            {children}

            {/* Secondary actions - all visible on desktop */}
            {secondaryActions.map((action, index) => (
              <Button
                key={index}
                onClick={action.onClick}
                variant={action.variant || 'outline'}
                disabled={action.disabled}
                size="sm"
              >
                {action.icon}
                {action.label}
              </Button>
            ))}

            {/* Primary action */}
            {primaryAction && (
              <Button
                onClick={primaryAction.onClick}
                variant={primaryAction.variant || 'default'}
                disabled={primaryAction.disabled}
                size="sm"
              >
                {primaryAction.icon}
                {primaryAction.label}
              </Button>
            )}
          </div>
        )}
      </div>

      {/* Divider */}
      <div className="border-b mt-4" />
    </div>
  )
}
