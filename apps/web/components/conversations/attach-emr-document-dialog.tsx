'use client';

/**
 * Picker "Anexar arquivo do EMR" do compositor de conversa: lista os documentos clínicos já
 * gerados do paciente (pedido de exames, emitido, receita, prontuário) e envia o escolhido pelo
 * WhatsApp sem o round-trip de baixar+reanexar. Modo automático pela janela 24h (arquivo se
 * aberta; senão link via template), com fallback se a janela fechar no meio.
 */

import { useState } from 'react';
import { toast } from 'sonner';
import { FileText, Loader2 } from 'lucide-react';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {
  useClinicalDocuments,
  useWhatsAppWindow,
  useSendClinicalDocWhatsApp,
  type ClinicalDocItem,
  type ClinicalDocType,
} from '@/lib/api/clinical-documents';

const TYPE_LABEL: Record<ClinicalDocType, string> = {
  lab_request: 'Pedido de exames',
  issued_document: 'Documento',
  prescription: 'Receita',
  patient_document: 'Prontuário',
};

interface Props {
  patientId: string;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function AttachEmrDocumentDialog({ patientId, open, onOpenChange }: Props) {
  const { data: docs, isLoading } = useClinicalDocuments(patientId, open);
  const { data: win } = useWhatsAppWindow(patientId);
  const send = useSendClinicalDocWhatsApp(patientId);
  const [pendingId, setPendingId] = useState<string | null>(null);

  async function pick(doc: ClinicalDocItem) {
    const mode = win?.windowOpen ? 'file' : 'link';
    setPendingId(doc.docId);
    try {
      await send.mutateAsync({ docType: doc.docType, docId: doc.docId, mode });
      toast.success(mode === 'file' ? 'Documento enviado no WhatsApp' : 'Link enviado no WhatsApp');
      onOpenChange(false);
    } catch (e: any) {
      if (mode === 'file' && e?.data?.message === 'window_closed') {
        try {
          await send.mutateAsync({ docType: doc.docType, docId: doc.docId, mode: 'link' });
          toast.success('Link enviado no WhatsApp');
          onOpenChange(false);
        } catch (e2: any) {
          toast.error('Falha ao enviar', { description: e2?.message });
        }
      } else {
        toast.error('Falha ao enviar', { description: e?.message });
      }
    } finally {
      setPendingId(null);
    }
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-lg">
        <DialogHeader>
          <DialogTitle>Anexar arquivo do EMR</DialogTitle>
          <DialogDescription>
            {win?.windowOpen
              ? 'Envia o arquivo direto no WhatsApp do paciente.'
              : 'Fora da janela de 24h — envia o link seguro do documento (reabre a conversa).'}
          </DialogDescription>
        </DialogHeader>

        <div className="max-h-[50vh] overflow-y-auto">
          {isLoading ? (
            <p className="py-6 text-center text-sm text-muted-foreground">Carregando…</p>
          ) : !docs?.length ? (
            <p className="py-6 text-center text-sm text-muted-foreground">
              Nenhum documento gerado para este paciente.
            </p>
          ) : (
            <ul className="space-y-1">
              {docs.map((d) => (
                <li key={`${d.docType}:${d.docId}`}>
                  <button
                    type="button"
                    disabled={pendingId !== null}
                    onClick={() => pick(d)}
                    className="flex w-full items-center gap-2 rounded-md border border-border px-3 py-2 text-left text-sm hover:bg-muted disabled:cursor-not-allowed disabled:opacity-50"
                  >
                    <FileText className="h-4 w-4 shrink-0 text-muted-foreground" />
                    <span className="min-w-0 flex-1">
                      <span className="block truncate font-medium">{d.title}</span>
                      <span className="block text-xs text-muted-foreground">
                        {TYPE_LABEL[d.docType]} · {new Date(d.createdAt).toLocaleDateString('pt-BR')}
                      </span>
                    </span>
                    {pendingId === d.docId && <Loader2 className="h-4 w-4 shrink-0 animate-spin" />}
                  </button>
                </li>
              ))}
            </ul>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
}
