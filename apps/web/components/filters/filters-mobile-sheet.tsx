'use client';

import { Filter } from 'lucide-react';
import type { ReactNode } from 'react';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet';

interface FiltersMobileSheetProps {
  /** Conteúdo dos chips/inputs de filtro. */
  children: ReactNode;
  /** Quantos filtros estão ativos — exibe badge no botão. */
  activeCount: number;
  /** Trigger só aparece em telas <md (768px). */
  className?: string;
}

/**
 * Botão "Filtros" + Sheet pra mobile (<md). Em telas maiores, esconda este
 * wrapper e renderize os chips inline.
 */
export function FiltersMobileSheet({
  children,
  activeCount,
  className,
}: FiltersMobileSheetProps) {
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="outline" size="sm" className={className}>
          <Filter className="mr-2 h-4 w-4" />
          Filtros
          {activeCount > 0 && (
            <Badge variant="secondary" className="ml-2 rounded px-1.5 py-0 text-xs">
              {activeCount}
            </Badge>
          )}
        </Button>
      </SheetTrigger>
      <SheetContent side="right" className="w-full sm:max-w-md overflow-y-auto">
        <SheetHeader>
          <SheetTitle>Filtros</SheetTitle>
          <SheetDescription>
            Refine a lista combinando os filtros abaixo.
          </SheetDescription>
        </SheetHeader>
        <div className="mt-6 flex flex-col gap-3">{children}</div>
      </SheetContent>
    </Sheet>
  );
}
