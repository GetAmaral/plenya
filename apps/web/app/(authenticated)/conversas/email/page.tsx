'use client';

import { useState } from 'react';
import { Mail } from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { cn } from '@/lib/utils';
import { PageHeader } from '@/components/layout/page-header';
import { NewEmailDialog } from '@/components/conversations/new-email-dialog';
import { EmailInbox, type EmailSelection } from '@/components/conversations/email-inbox';
import { NotificationsList } from '@/components/conversations/notifications-list';
import { useNotificationEmails } from '@/lib/api/conversations-api';

type Tab = 'caixa' | 'notificacoes';

/**
 * Caixa de e-mail dedicada (webmail). Separada do WhatsApp: assíncrona, com assunto e
 * anexos. Abas: Caixa (leads/pacientes reais) e Notificações (e-mails automáticos).
 */
export default function ConversasEmailPage() {
  useRequireAuth();
  const [tab, setTab] = useState<Tab>('caixa');
  const [selected, setSelected] = useState<EmailSelection>(null);
  const [newEmailOpen, setNewEmailOpen] = useState(false);
  const notifications = useNotificationEmails();
  const unreadNotifs = notifications.data?.unreadCount ?? 0;

  return (
    <div className="flex h-[calc(100vh-4rem)] flex-col gap-3 p-3 sm:p-4">
      <PageHeader
        title="E-mails"
        description="Caixa de e-mail com leads e pacientes — assunto, histórico e anexos."
        actions={[
          {
            label: 'Novo email',
            icon: <Mail className="h-4 w-4" />,
            onClick: () => setNewEmailOpen(true),
          },
        ]}
      />

      <NewEmailDialog
        open={newEmailOpen}
        onOpenChange={setNewEmailOpen}
        onSent={({ ownerType, ownerId }) => {
          setTab('caixa');
          setSelected({ type: ownerType, id: ownerId });
        }}
      />

      {/* Abas Caixa | Notificações */}
      <div className="flex shrink-0 items-center gap-1 rounded-lg border border-border bg-muted/40 p-1 self-start">
        <button
          type="button"
          onClick={() => setTab('caixa')}
          className={cn(
            'rounded-md px-3 py-1.5 text-sm font-medium transition-colors',
            tab === 'caixa' ? 'bg-background shadow-sm' : 'text-muted-foreground hover:text-foreground'
          )}
        >
          Caixa
        </button>
        <button
          type="button"
          onClick={() => setTab('notificacoes')}
          className={cn(
            'flex items-center gap-1.5 rounded-md px-3 py-1.5 text-sm font-medium transition-colors',
            tab === 'notificacoes' ? 'bg-background shadow-sm' : 'text-muted-foreground hover:text-foreground'
          )}
        >
          Notificações
          {unreadNotifs > 0 && (
            <span className="flex h-5 min-w-5 items-center justify-center rounded-full bg-amber-500 px-1 text-[11px] font-semibold text-white">
              {unreadNotifs > 99 ? '99+' : unreadNotifs}
            </span>
          )}
        </button>
      </div>

      {tab === 'caixa' ? (
        <EmailInbox selected={selected} onSelect={setSelected} />
      ) : (
        <NotificationsList />
      )}
    </div>
  );
}
