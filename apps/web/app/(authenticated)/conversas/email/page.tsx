'use client';

import { useState } from 'react';
import { Mail } from 'lucide-react';

import { useRequireAuth } from '@/lib/use-auth';
import { PageHeader } from '@/components/layout/page-header';
import { NewEmailDialog } from '@/components/conversations/new-email-dialog';
import { EmailInbox, type EmailSelection } from '@/components/conversations/email-inbox';

/**
 * Caixa de e-mail dedicada (webmail). Separada do WhatsApp: assíncrona, com assunto e
 * anexos. Sem barra de paciente (inbox de CRM).
 */
export default function ConversasEmailPage() {
  useRequireAuth();
  const [selected, setSelected] = useState<EmailSelection>(null);
  const [newEmailOpen, setNewEmailOpen] = useState(false);

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
          setSelected({ type: ownerType, id: ownerId });
        }}
      />

      <EmailInbox selected={selected} onSelect={setSelected} />
    </div>
  );
}
