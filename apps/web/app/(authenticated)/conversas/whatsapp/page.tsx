'use client';

import { useState } from 'react';

import { useRequireAuth } from '@/lib/use-auth';
import { PageHeader } from '@/components/layout/page-header';
import { ReceptionMetricsBar } from '@/components/conversations/reception-metrics-bar';
import { WhatsAppChat, type ChatSelection } from '@/components/conversations/whatsapp-chat';

/**
 * Superfície dedicada de WhatsApp (chat de verdade). Sem barra de paciente selecionado —
 * é uma inbox de CRM, não uma tela de dados clínicos.
 */
export default function ConversasWhatsAppPage() {
  useRequireAuth();
  const [selected, setSelected] = useState<ChatSelection>(null);

  return (
    <div className="flex h-[calc(100vh-4rem)] flex-col gap-3 p-3 sm:p-4">
      <PageHeader
        title="WhatsApp"
        description="Conversas de WhatsApp com leads e pacientes — chat e respostas em um só lugar."
      />
      <ReceptionMetricsBar />
      <WhatsAppChat variant="page" selected={selected} onSelect={setSelected} />
    </div>
  );
}
