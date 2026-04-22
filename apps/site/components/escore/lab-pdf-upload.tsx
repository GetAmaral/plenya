'use client';

/**
 * Componente de upload de PDF de exames para o Escore Light.
 *
 * Fluxo:
 *   1. Usuário escolhe um PDF (≤10MB)
 *   2. Upload síncrono para /api/v1/score-light/extract-labs
 *   3. Backend faz OCR + IA + match contra subset Light
 *   4. Resultado: lista de items preenchidos automaticamente + lista de "não reconhecidos"
 *   5. Pai (form) injeta os values em `responses`
 *
 * Privacidade: o PDF é processado em segundos e descartado pelo backend (LGPD).
 * Apenas os valores numéricos extraídos persistem na sessão (anônima, expira 90d).
 */
import { useRef, useState } from 'react';
import { extractLabsFromPDF, type LightLabExtractionResult, type LightLabMatch } from '@/lib/score-light/api';

const MAX_MB = 10;

export function LabPDFUpload({
  onExtracted,
}: {
  onExtracted: (matched: LightLabMatch[]) => void;
}) {
  const inputRef = useRef<HTMLInputElement>(null);
  const [status, setStatus] = useState<'idle' | 'uploading' | 'done' | 'error'>('idle');
  const [error, setError] = useState<string | null>(null);
  const [result, setResult] = useState<LightLabExtractionResult | null>(null);
  const [filename, setFilename] = useState<string | null>(null);

  async function handleFile(file: File) {
    setError(null);
    setResult(null);

    if (file.type !== 'application/pdf') {
      setError('Envie um arquivo PDF.');
      setStatus('error');
      return;
    }
    if (file.size > MAX_MB * 1024 * 1024) {
      setError(`PDF deve ter no máximo ${MAX_MB}MB.`);
      setStatus('error');
      return;
    }

    setFilename(file.name);
    setStatus('uploading');

    try {
      const r = await extractLabsFromPDF(file);
      setResult(r);
      setStatus('done');
      if (r.matched.length > 0) {
        onExtracted(r.matched);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Falha ao processar o PDF.');
      setStatus('error');
    }
  }

  function reset() {
    setStatus('idle');
    setResult(null);
    setError(null);
    setFilename(null);
    if (inputRef.current) inputRef.current.value = '';
  }

  return (
    <div className="border border-petrol/15 rounded-md bg-paper p-5 my-3">
      <div className="flex items-baseline justify-between gap-3 mb-2">
        <p className="label-upper text-gold text-[10px]">Atalho — importar PDF de exames</p>
        {status !== 'idle' && (
          <button
            type="button"
            onClick={reset}
            className="text-xs text-petrol/50 hover:text-petrol/80 underline underline-offset-4"
          >
            limpar
          </button>
        )}
      </div>
      <p className="text-petrol/70 text-sm leading-relaxed mb-4">
        Tem o PDF dos seus exames laboratoriais? Faça upload e nossa IA preenche
        automaticamente os campos de exames abaixo — você pode revisar e ajustar
        antes de enviar.
      </p>
      <div className="text-petrol/55 text-xs leading-relaxed mb-4 bg-cream border-l-2 border-gold pl-3 py-2 space-y-1">
        <p>
          <strong className="text-petrol/80">Como funciona:</strong> seu PDF é enviado para a{' '}
          <strong>Anthropic Claude</strong> (inteligência artificial nos EUA, sob cláusulas-padrão
          de proteção de dados) para extrair os valores dos exames.
        </p>
        <p>
          <strong className="text-petrol/80">Privacidade:</strong> o arquivo é descartado em
          segundos pelo nosso servidor e a Anthropic não retém o conteúdo. Apenas os valores
          numéricos extraídos ficam guardados, anônimos, na sua sessão (90 dias). Detalhes na{' '}
          <a href="/pt/privacidade" target="_blank" rel="noopener noreferrer" className="text-gold underline underline-offset-4">
            Política de Privacidade
          </a>.
        </p>
        <p className="italic">
          Se preferir, pule este atalho e preencha manualmente abaixo — funciona da mesma forma.
        </p>
      </div>

      {status === 'idle' && (
        <>
          <input
            ref={inputRef}
            type="file"
            accept="application/pdf,.pdf"
            onChange={(e) => {
              const f = e.target.files?.[0];
              if (f) void handleFile(f);
            }}
            className="hidden"
            id="lab-pdf-upload"
          />
          <label
            htmlFor="lab-pdf-upload"
            className="inline-block px-5 py-2.5 bg-gold text-petrol cursor-pointer hover:opacity-90 transition text-sm rounded-md"
          >
            Escolher PDF
          </label>
        </>
      )}

      {status === 'uploading' && (
        <div className="flex items-center gap-3 text-petrol">
          <span className="inline-block w-4 h-4 border-2 border-petrol/30 border-t-petrol rounded-full animate-spin" />
          <span className="text-sm">
            Processando <strong>{filename}</strong>… isso leva ~10 segundos.
          </span>
        </div>
      )}

      {status === 'error' && error && (
        <div className="border border-red-200 bg-red-50 text-red-700 text-sm p-3 rounded-md">
          {error}
        </div>
      )}

      {status === 'done' && result && (
        <div className="space-y-3">
          {result.matched.length > 0 ? (
            <div className="border border-emerald-200 bg-emerald-50 text-emerald-900 text-sm p-3 rounded-md">
              <p className="font-medium mb-1">
                ✓ {result.matched.length}{' '}
                {result.matched.length === 1 ? 'exame reconhecido' : 'exames reconhecidos'} e preenchido
                {result.matched.length > 1 ? 's' : ''} abaixo.
              </p>
              <ul className="text-xs text-emerald-800/80 space-y-0.5">
                {result.matched.map((m) => (
                  <li key={m.scoreItemId}>
                    — {m.itemName}: <strong>{m.numericValue}</strong>{' '}
                    {m.unit && <span className="opacity-60">{m.unit}</span>}
                  </li>
                ))}
              </ul>
            </div>
          ) : (
            <div className="border border-amber-200 bg-amber-50 text-amber-900 text-sm p-3 rounded-md">
              Não consegui reconhecer nenhum exame relevante para o Escore Light
              neste PDF. Você pode preencher manualmente abaixo.
            </div>
          )}

          {result.unmatched.length > 0 && (
            <details className="text-xs text-petrol/60">
              <summary className="cursor-pointer hover:text-petrol/80">
                {result.unmatched.length} exame{result.unmatched.length > 1 ? 's' : ''}{' '}
                do PDF não usado{result.unmatched.length > 1 ? 's' : ''} no Escore
                Light
              </summary>
              <ul className="mt-2 ml-3 space-y-0.5">
                {result.unmatched.slice(0, 20).map((u, i) => (
                  <li key={i}>
                    — {u.nomeExame}: {u.resultado} {u.unidade && <span>{u.unidade}</span>}
                  </li>
                ))}
                {result.unmatched.length > 20 && (
                  <li className="opacity-60">
                    …e mais {result.unmatched.length - 20}
                  </li>
                )}
              </ul>
            </details>
          )}

          {result.warnings && result.warnings.length > 0 && (
            <div className="text-xs text-amber-700">
              {result.warnings.map((w, i) => (
                <p key={i}>⚠ {w}</p>
              ))}
            </div>
          )}
        </div>
      )}
    </div>
  );
}
