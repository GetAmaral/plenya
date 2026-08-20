'use client';

/**
 * Botão "Enviar por e-mail" dos documentos clínicos (pedido de exames, emitido, receita).
 *
 * Vai o LINK seguro, não o PDF anexo: o arquivo traz dado clínico identificável, e-mail comum
 * atravessa servidor que não controlamos e fica na caixa do paciente para sempre. É o mesmo link
 * assinado e com prazo que o envio por WhatsApp usa.
 *
 * Envio é sempre ato explícito. Assinar uma receita gera o PDF e publica no portal do paciente —
 * não manda e-mail nem WhatsApp.
 */

import { useState } from 'react';
import { toast } from 'sonner';
import { Loader2, Mail } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { BotaoEnvioTooltip } from '@/components/clinical/botao-envio-tooltip';
import {
  useDocumentChannels,
  useSendClinicalDocEmail,
  type ClinicalDocType,
} from '@/lib/api/clinical-documents';

interface Props {
  patientId: string;
  docType: ClinicalDocType;
  docId: string;
  /** desabilita quando o documento ainda não tem PDF pronto. */
  disabled?: boolean;
}

export function SendDocumentEmailButton({ patientId, docType, docId, disabled }: Props) {
  const { data: canais } = useDocumentChannels(patientId);
  const send = useSendClinicalDocEmail(patientId);
  const [busy, setBusy] = useState(false);
  // Enquanto os canais não chegam, o botão fica habilitado: desabilitar por dado ausente que ainda
  // está carregando faria o botão piscar de desligado para ligado a cada abertura da tela.
  const semEmail = canais != null && !canais.hasEmail;

  async function enviar() {
    setBusy(true);
    try {
      await send.mutateAsync({ docType, docId });
      toast.success('Link do documento enviado por e-mail');
    } catch (e: any) {
      // Paciente sem e-mail é o caso comum: a recepção cadastra só pelo nome.
      if (e?.data?.message === 'no_email') {
        toast.error('Paciente sem e-mail cadastrado', {
          description: 'Cadastre o e-mail na ficha do paciente para enviar por aqui.',
        });
      } else {
        toast.error('Não deu para enviar', { description: e?.message });
      }
    } finally {
      setBusy(false);
    }
  }

  return (
    <BotaoEnvioTooltip
      motivo={semEmail ? 'Cadastre o e-mail do paciente para ativar o envio por e-mail' : undefined}
    >
      <Button
        variant="outline"
        size="sm"
        onClick={enviar}
        disabled={disabled || busy || semEmail}
      >
        {busy ? (
          <Loader2 className="mr-2 h-4 w-4 animate-spin" />
        ) : (
          <Mail className="mr-2 h-4 w-4" />
        )}
        E-mail
      </Button>
    </BotaoEnvioTooltip>
  );
}
