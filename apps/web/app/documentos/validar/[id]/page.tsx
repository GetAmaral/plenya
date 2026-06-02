'use client';

import { useQuery } from '@tanstack/react-query';
import { format } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { CheckCircle2, AlertCircle, FileCheck, User, Stethoscope, Shield, ExternalLink } from 'lucide-react';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001';

const TYPE_LABEL: Record<string, string> = {
  certificate: 'Atestado médico',
  declaration: 'Declaração',
  report: 'Laudo/Relatório',
};

interface ValidationResult {
  valid: boolean;
  pdfIntact: boolean;
  document: {
    id: string;
    type: string;
    title: string;
    status: string;
    hasDigitalSignature: boolean;
    issuedAt: string;
    signedAt?: string;
  };
  patient: { name: string };
  doctor: { name: string; crm: string };
}

async function validateDocument(id: string): Promise<ValidationResult> {
  const res = await fetch(`${API_URL}/api/v1/documents/validate/${id}`);
  if (!res.ok) throw new Error('not found');
  return res.json();
}

export default function ValidateDocumentPage({ params }: { params: { id: string } }) {
  const { data, isLoading, error } = useQuery({
    queryKey: ['document-validation', params.id],
    queryFn: () => validateDocument(params.id),
    retry: false,
  });

  if (isLoading) {
    return (
      <div className="container mx-auto max-w-2xl py-12">
        <Card>
          <CardContent className="flex items-center justify-center py-12">
            <div className="space-y-4 text-center">
              <FileCheck className="mx-auto h-12 w-12 animate-pulse text-muted-foreground" />
              <p className="text-muted-foreground">Validando documento…</p>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  if (error || !data) {
    return (
      <div className="container mx-auto max-w-2xl py-12">
        <Alert variant="destructive">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Documento não encontrado</AlertTitle>
          <AlertDescription>
            Não foi possível encontrar ou validar este documento. Verifique se o QR Code está correto.
          </AlertDescription>
        </Alert>
      </div>
    );
  }

  const isValid = data.valid && data.pdfIntact;

  return (
    <div className="container mx-auto max-w-3xl py-12">
      <div className="mb-8 text-center">
        <h1 className="mb-2 text-3xl font-bold">Validação de Documento Clínico</h1>
        <p className="text-muted-foreground">Plenya — verificação de autenticidade por hash</p>
      </div>

      <Alert
        variant={isValid ? 'default' : 'destructive'}
        className={isValid ? 'mb-8 border-green-500 bg-green-50 dark:bg-green-950' : 'mb-8'}
      >
        {isValid ? <CheckCircle2 className="h-5 w-5 text-green-600" /> : <AlertCircle className="h-5 w-5" />}
        <AlertTitle className={isValid ? 'text-xl text-green-600' : 'text-xl'}>
          {isValid ? 'DOCUMENTO VÁLIDO' : 'DOCUMENTO INVÁLIDO'}
        </AlertTitle>
        <AlertDescription className={isValid ? 'text-green-700' : ''}>
          {isValid
            ? 'Este documento é autêntico e não foi alterado.'
            : !data.pdfIntact
              ? 'O PDF foi alterado — integridade comprometida.'
              : 'Documento ainda não assinado/emitido.'}
        </AlertDescription>
      </Alert>

      <div className="grid gap-6">
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <FileCheck className="h-5 w-5" />
              Documento
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-3">
            <div className="flex items-center gap-2">
              <Badge variant="secondary">{TYPE_LABEL[data.document.type] ?? data.document.type}</Badge>
              {data.document.hasDigitalSignature ? (
                <Badge variant="secondary" className="border-emerald-200 bg-emerald-100 text-emerald-900">
                  <Shield className="mr-1 h-3 w-3" /> Assinatura ICP-Brasil
                </Badge>
              ) : (
                <Badge variant="outline">Assinatura manual</Badge>
              )}
            </div>
            <div>
              <p className="text-sm text-muted-foreground">Título</p>
              <p className="font-medium">{data.document.title}</p>
            </div>
            <div className="grid grid-cols-2 gap-4">
              <div>
                <p className="text-sm text-muted-foreground">Emitido em</p>
                <p className="font-medium">
                  {format(new Date(data.document.issuedAt), "dd 'de' MMMM 'de' yyyy", { locale: ptBR })}
                </p>
              </div>
              {data.document.signedAt && (
                <div>
                  <p className="text-sm text-muted-foreground">Assinado em</p>
                  <p className="font-medium">
                    {format(new Date(data.document.signedAt), "dd/MM/yyyy 'às' HH:mm", { locale: ptBR })}
                  </p>
                </div>
              )}
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <User className="h-5 w-5" />
              Paciente
            </CardTitle>
          </CardHeader>
          <CardContent>
            <p className="font-medium">{data.patient.name}</p>
            <p className="mt-1 text-xs text-muted-foreground">(Nome mascarado por segurança — LGPD)</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Stethoscope className="h-5 w-5" />
              Médico responsável
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-1">
            <p className="font-medium">{data.doctor.name}</p>
            <p className="text-sm text-muted-foreground">{data.doctor.crm}</p>
          </CardContent>
        </Card>

        {data.document.hasDigitalSignature && (
          <Button
            variant="outline"
            size="lg"
            onClick={() => window.open('https://validar.iti.gov.br/', '_blank')}
          >
            <ExternalLink className="mr-2 h-4 w-4" />
            Verificar assinatura no ITI
          </Button>
        )}
      </div>
    </div>
  );
}
