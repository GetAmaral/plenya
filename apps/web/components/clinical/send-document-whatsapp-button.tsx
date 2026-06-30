'use client';

/**
 * Botão "Enviar por WhatsApp" compartilhado pelos cartões de documento clínico (pedido de exames,
 * emitido, receita). Smart por janela de 24h:
 *   - janela aberta  → menu: "Enviar arquivo" (mídia inline) + "Enviar link" (template)
 *   - janela fechada → botão único "Enviar link" (template reabre a conversa)
 * Fallback: se o arquivo falhar por janela fechada (race), reenvia automaticamente como link.
 */

import { useState } from 'react';
import { toast } from 'sonner';
import { MessageCircle, FileUp, Link2, Loader2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import {
  useWhatsAppWindow,
  useSendClinicalDocWhatsApp,
  type ClinicalDocType,
  type SendMode,
} from '@/lib/api/clinical-documents';

interface Props {
  patientId: string;
  docType: ClinicalDocType;
  docId: string;
  /** desabilita quando o documento ainda não tem PDF pronto. */
  disabled?: boolean;
}

export function SendDocumentWhatsAppButton({ patientId, docType, docId, disabled }: Props) {
  const { data: win } = useWhatsAppWindow(patientId);
  const send = useSendClinicalDocWhatsApp(patientId);
  const [pending, setPending] = useState<SendMode | null>(null);
  const busy = pending !== null;
  const noPhone = win != null && !win.hasPhone;

  async function doSend(mode: SendMode) {
    setPending(mode);
    try {
      await send.mutateAsync({ docType, docId, mode });
      toast.success(
        mode === 'file'
          ? 'Documento enviado no WhatsApp'
          : 'Link do documento enviado no WhatsApp',
      );
    } catch (e: any) {
      // Arquivo fora da janela 24h (race entre o estado e o envio) → cai pro link automaticamente.
      const windowClosed = e?.data?.message === 'window_closed';
      if (mode === 'file' && windowClosed) {
        toast.message('Janela de 24h fechada', {
          description: 'Enviando o link do documento (reabre a conversa).',
        });
        try {
          await send.mutateAsync({ docType, docId, mode: 'link' });
          toast.success('Link do documento enviado no WhatsApp');
        } catch (e2: any) {
          toast.error('Falha ao enviar', { description: e2?.message });
        }
      } else {
        toast.error('Falha ao enviar', { description: e?.message });
      }
    } finally {
      setPending(null);
    }
  }

  const icon = busy ? (
    <Loader2 className="animate-spin" />
  ) : (
    <MessageCircle />
  );

  // Janela aberta → menu com arquivo + link.
  if (win?.windowOpen) {
    return (
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            size="sm"
            disabled={disabled || busy || noPhone}
            title={noPhone ? 'Paciente sem telefone' : 'Enviar por WhatsApp'}
          >
            {icon}
            WhatsApp
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuItem onClick={() => doSend('file')} disabled={busy}>
            <FileUp className="mr-2 h-4 w-4" />
            Enviar arquivo
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => doSend('link')} disabled={busy}>
            <Link2 className="mr-2 h-4 w-4" />
            Enviar link
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    );
  }

  // Janela fechada (ou ainda desconhecida) → só link.
  return (
    <Button
      variant="outline"
      size="sm"
      disabled={disabled || busy || noPhone}
      onClick={() => doSend('link')}
      title={noPhone ? 'Paciente sem telefone' : 'Enviar link por WhatsApp (fora da janela de 24h)'}
    >
      {icon}
      WhatsApp
    </Button>
  );
}
