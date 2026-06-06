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

  // Esta superfície é tela-cheia: cancela o padding do layout autenticado
  // (p-4 pt-6 sm:p-6 lg:p-8) com margens negativas e ocupa toda a altura abaixo do
  // TopBar (h-16 = 4rem; a barra de paciente é null em /conversas). O PageHeader fica
  // só pelo efeito colateral de setar o título do TopBar (sem description → corpo vazio).
  return (
    <div className="-mx-4 -mb-4 -mt-6 flex h-[calc(100vh-4rem)] flex-col gap-2 p-2 sm:-m-6 sm:p-3 lg:-m-8">
      <PageHeader title="WhatsApp" />
      <ReceptionMetricsBar />
      <WhatsAppChat variant="page" selected={selected} onSelect={setSelected} />
    </div>
  );
}
