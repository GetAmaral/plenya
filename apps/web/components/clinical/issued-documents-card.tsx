'use client';

/**
 * IssuedDocumentsCard — documentos clínicos emitidos (atestado/declaração/laudo).
 * Usado na capa do paciente e no workspace de consulta. Emite rascunho e assina
 * (ICP-Brasil quando o médico tem certificado; senão gera com assinatura manual).
 * Publicado automaticamente no portal do paciente.
 */

import { useState } from 'react';
import { FileSignature, Plus, Loader2, Download, ShieldCheck, PenLine, Trash2 } from 'lucide-react';
import { toast } from 'sonner';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Checkbox } from '@/components/ui/checkbox';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import {
  useIssuedDocuments,
  useCreateIssuedDocument,
  useSignIssuedDocument,
  useDeleteIssuedDocument,
  openIssuedDocumentPDF,
  type IssuedDocumentType,
} from '@/lib/api/issued-documents';

const selectCls =
  'w-full rounded-md border border-input bg-background px-3 py-2 text-sm outline-none focus:border-primary';

const TYPE_LABEL: Record<IssuedDocumentType, string> = {
  certificate: 'Atestado',
  declaration: 'Declaração',
  report: 'Laudo/Relatório',
};

export function IssuedDocumentsCard({
  patientId,
  appointmentId,
  className,
}: {
  patientId: string;
  appointmentId?: string;
  className?: string;
}) {
  const { data = [], isLoading } = useIssuedDocuments(patientId);
  const createDoc = useCreateIssuedDocument(patientId);
  const signDoc = useSignIssuedDocument(patientId);
  const deleteDoc = useDeleteIssuedDocument(patientId);

  const [open, setOpen] = useState(false);
  const [type, setType] = useState<IssuedDocumentType>('certificate');
  const [title, setTitle] = useState('');
  const [body, setBody] = useState('');
  const [purpose, setPurpose] = useState('');
  const [daysOff, setDaysOff] = useState('');
  const [includeCid, setIncludeCid] = useState(false);
  const [cidCode, setCidCode] = useState('');
  const [cidConsent, setCidConsent] = useState(false);
  const [busy, setBusy] = useState(false);

  const reset = () => {
    setType('certificate');
    setTitle('');
    setBody('');
    setPurpose('');
    setDaysOff('');
    setIncludeCid(false);
    setCidCode('');
    setCidConsent(false);
  };

  const buildPayload = () => ({
    appointmentId,
    type,
    title: title.trim(),
    body: body.trim(),
    purpose: purpose.trim() || undefined,
    daysOff: type === 'certificate' && daysOff ? parseInt(daysOff, 10) : undefined,
    includesCid: includeCid,
    cidCode: includeCid ? cidCode.trim() || undefined : undefined,
    cidConsent: includeCid ? cidConsent : false,
  });

  const validate = () => {
    if (!title.trim()) {
      toast.error('Informe um título');
      return false;
    }
    if (includeCid && !cidConsent) {
      toast.error('Incluir CID exige consentimento do paciente');
      return false;
    }
    return true;
  };

  const handleSaveDraft = async () => {
    if (!validate()) return;
    setBusy(true);
    try {
      await createDoc.mutateAsync(buildPayload());
      toast.success('Rascunho salvo');
      setOpen(false);
      reset();
    } catch (err) {
      toast.error('Falha ao salvar', { description: err instanceof Error ? err.message : undefined });
    } finally {
      setBusy(false);
    }
  };

  const handleSignAndIssue = async () => {
    if (!validate()) return;
    setBusy(true);
    try {
      const created = await createDoc.mutateAsync(buildPayload());
      const signed = await signDoc.mutateAsync(created.id);
      toast.success(
        signed.hasDigitalSignature ? 'Documento assinado (ICP-Brasil)' : 'Documento emitido (assinatura manual)',
        { description: 'Disponível no portal do paciente' },
      );
      setOpen(false);
      reset();
    } catch (err) {
      toast.error('Falha ao emitir', { description: err instanceof Error ? err.message : undefined });
    } finally {
      setBusy(false);
    }
  };

  const handleSign = async (docId: string) => {
    try {
      const signed = await signDoc.mutateAsync(docId);
      toast.success(
        signed.hasDigitalSignature ? 'Documento assinado (ICP-Brasil)' : 'Documento emitido (assinatura manual)',
      );
    } catch (err) {
      toast.error('Falha ao assinar', { description: err instanceof Error ? err.message : undefined });
    }
  };

  const handleDownload = async (docId: string) => {
    try {
      await openIssuedDocumentPDF(docId);
    } catch (err) {
      toast.error('Falha ao baixar PDF', { description: err instanceof Error ? err.message : undefined });
    }
  };

  return (
    <Card className={className}>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="flex items-center gap-2 text-base">
          <FileSignature className="h-4 w-4 text-primary" />
          Documentos emitidos
        </CardTitle>
        <Button variant="outline" size="sm" onClick={() => setOpen(true)}>
          <Plus className="mr-1 h-4 w-4" />
          Emitir documento
        </Button>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <Loader2 className="h-4 w-4 animate-spin text-muted-foreground" />
        ) : data.length === 0 ? (
          <p className="text-sm text-muted-foreground">Nenhum documento emitido.</p>
        ) : (
          <ul className="space-y-2">
            {data.map((doc) => (
              <li key={doc.id} className="flex items-center gap-2 rounded-md border p-2">
                <div className="min-w-0 flex-1">
                  <div className="flex flex-wrap items-center gap-2">
                    <span className="truncate text-sm font-medium">{doc.title}</span>
                    <Badge variant="secondary" className="text-[10px]">
                      {TYPE_LABEL[doc.type]}
                    </Badge>
                    {doc.status === 'signed' ? (
                      doc.hasDigitalSignature ? (
                        <Badge variant="secondary" className="border-emerald-200 bg-emerald-100 text-emerald-900 text-[10px]">
                          <ShieldCheck className="mr-1 h-3 w-3" /> Assinado ICP
                        </Badge>
                      ) : (
                        <Badge variant="secondary" className="text-[10px]">
                          <PenLine className="mr-1 h-3 w-3" /> Assinatura manual
                        </Badge>
                      )
                    ) : (
                      <Badge variant="outline" className="text-[10px]">
                        Rascunho
                      </Badge>
                    )}
                  </div>
                  <p className="text-xs text-muted-foreground">
                    {new Date(doc.issuedAt).toLocaleDateString('pt-BR')}
                    {doc.daysOff ? ` · ${doc.daysOff} dia(s) de afastamento` : ''}
                  </p>
                </div>
                <div className="flex shrink-0 items-center gap-1">
                  {doc.status === 'draft' ? (
                    <>
                      <Button size="sm" onClick={() => handleSign(doc.id)} disabled={signDoc.isPending}>
                        Assinar
                      </Button>
                      <Button
                        variant="ghost"
                        size="icon"
                        className="h-8 w-8"
                        onClick={() => deleteDoc.mutate(doc.id)}
                        title="Excluir rascunho"
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </>
                  ) : (
                    <Button variant="outline" size="sm" onClick={() => handleDownload(doc.id)}>
                      <Download className="mr-1 h-4 w-4" /> PDF
                    </Button>
                  )}
                </div>
              </li>
            ))}
          </ul>
        )}
      </CardContent>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="max-h-[90vh] overflow-y-auto sm:max-w-lg">
          <DialogHeader>
            <DialogTitle>Emitir documento</DialogTitle>
          </DialogHeader>
          <div className="space-y-3">
            <div>
              <Label htmlFor="doc-type">Tipo</Label>
              <select
                id="doc-type"
                className={selectCls}
                value={type}
                onChange={(e) => setType(e.target.value as IssuedDocumentType)}
              >
                <option value="certificate">Atestado</option>
                <option value="declaration">Declaração</option>
                <option value="report">Laudo/Relatório</option>
              </select>
            </div>
            <div>
              <Label htmlFor="doc-title">Título</Label>
              <Input id="doc-title" value={title} onChange={(e) => setTitle(e.target.value)} placeholder="Ex.: Atestado médico" />
            </div>
            <div>
              <Label htmlFor="doc-body">Corpo do documento</Label>
              <Textarea
                id="doc-body"
                rows={5}
                value={body}
                onChange={(e) => setBody(e.target.value)}
                placeholder="Texto do atestado/declaração/laudo…"
              />
            </div>
            {type === 'certificate' && (
              <div>
                <Label htmlFor="doc-days">Dias de afastamento (opcional)</Label>
                <Input
                  id="doc-days"
                  type="number"
                  min={0}
                  max={365}
                  value={daysOff}
                  onChange={(e) => setDaysOff(e.target.value)}
                />
              </div>
            )}
            <div>
              <Label htmlFor="doc-purpose">Finalidade (opcional)</Label>
              <Input id="doc-purpose" value={purpose} onChange={(e) => setPurpose(e.target.value)} />
            </div>
            <div className="space-y-2 rounded-md border p-3">
              <div className="flex items-center gap-2">
                <Checkbox id="doc-cid" checked={includeCid} onCheckedChange={(v) => setIncludeCid(!!v)} />
                <Label htmlFor="doc-cid" className="cursor-pointer">
                  Incluir CID no documento
                </Label>
              </div>
              {includeCid && (
                <div className="space-y-2 pl-6">
                  <Input value={cidCode} onChange={(e) => setCidCode(e.target.value)} placeholder="Código CID-10 (ex.: J11)" />
                  <div className="flex items-center gap-2">
                    <Checkbox id="doc-consent" checked={cidConsent} onCheckedChange={(v) => setCidConsent(!!v)} />
                    <Label htmlFor="doc-consent" className="cursor-pointer text-xs">
                      O paciente consente com a inclusão do CID (sigilo)
                    </Label>
                  </div>
                </div>
              )}
            </div>
          </div>
          <DialogFooter className="gap-2">
            <Button variant="outline" onClick={handleSaveDraft} disabled={busy}>
              {busy && <Loader2 className="mr-1 h-4 w-4 animate-spin" />}
              Salvar rascunho
            </Button>
            <Button onClick={handleSignAndIssue} disabled={busy}>
              {busy && <Loader2 className="mr-1 h-4 w-4 animate-spin" />}
              Assinar e emitir
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </Card>
  );
}
