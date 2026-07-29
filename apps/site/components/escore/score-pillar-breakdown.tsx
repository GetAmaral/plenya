'use client';

/**
 * Detalhamento por pilar e subpilar do Escore Plenya (Light/Triagem), na página de resultado.
 *
 * Mesma estrutura e mesmas cores do EMR — a lógica de agregação e a paleta vêm de
 * `@plenya/ui/score-radar`, então os números batem com a legenda do radar e a cor de cada
 * letra é a mesma do arco.
 *
 * A apresentação é do site (petrol/cream, tipografia institucional), não a do prontuário.
 */

import { useMemo, useState } from 'react';
import { useLocale, useTranslations } from 'next-intl';
import {
  buildPillarBreakdown,
  letterColor,
  type BreakdownItemResult,
  type BreakdownPilar,
} from '@plenya/ui/score-radar';
import type { PublicSnapshot } from '@/lib/score-light/types';

// Display code AGIR (PT) → ACTS (EN), igual ao radar.
const EN_DISPLAY: Record<string, string> = { A: 'A', G: 'C', I: 'T', R: 'S' };

/**
 * Rótulos de grupo/subgrupo carregam roteiro de anamnese entre parênteses. É instrução para o
 * profissional, não texto de tela.
 */
function limpaRotulo(s: string) {
  return s
    .replace(/\s*\(([^)]{21,})\)/g, '')
    .replace(/[\s:·-]+$/, '')
    .trim();
}

export function ScorePillarBreakdown({ snapshot }: { snapshot: PublicSnapshot }) {
  const t = useTranslations('escoreLight');
  const isEn = useLocale() === 'en';
  const [aberto, setAberto] = useState<string | null>(null);

  const pilares = useMemo<BreakdownPilar[]>(() => {
    const entrada: BreakdownItemResult[] = (snapshot.itemResults ?? []).map((r) => ({
      status: r.status,
      actualPoints: r.actualPoints,
      maxPoints: r.maxPoints,
      levelNumber: r.levelNumber ?? null,
      levelName: r.levelName,
      bloco: limpaRotulo(r.bloco ?? ''),
      item: { name: r.item?.name, methodPillars: r.item?.methodPillars },
    }));
    return buildPillarBreakdown(entrada);
  }, [snapshot]);

  if (pilares.length === 0) return null;

  return (
    <div className="space-y-4">
      {pilares.map((p) => {
        const cor = letterColor(p.code).light;
        return (
          <section
            key={p.code}
            className="rounded-lg border p-4 sm:p-6"
            style={{ borderColor: cor }}
          >
            <div className="mb-2 flex flex-wrap items-baseline gap-x-3 gap-y-1">
              <span
                className="flex h-7 w-7 shrink-0 items-center justify-center rounded-full text-xs font-semibold text-white"
                style={{ background: cor }}
              >
                {isEn ? (EN_DISPLAY[p.code] ?? p.code) : p.code}
              </span>
              <h3 className="text-petrol min-w-0 flex-1 text-base">{p.name}</h3>
              <span className="text-petrol/60 shrink-0 text-sm">
                <span className="text-petrol text-xl font-light tabular-nums">
                  {Math.round(p.pct)}
                </span>
                <span className="ml-1">{t('resultHeroOf100')}</span>
              </span>
            </div>

            {/* Barra do PILAR */}
            <span
              className="mb-4 block h-3.5 w-full overflow-hidden rounded-r"
              style={{ background: `${cor}1f` }}
            >
              <span
                className="block h-full rounded-r-[4px] transition-[width] duration-500"
                style={{ width: `${Math.max(1.5, p.pct)}%`, background: cor }}
              />
            </span>

            {/* Barras dos SUBPILARES — os pontos do radar, na mesma ordem */}
            <ul className="space-y-2">
              {p.subpilares.map((s) => {
                const isOpen = aberto === s.key;
                return (
                  <li key={s.key}>
                    <button
                      type="button"
                      onClick={() => setAberto(isOpen ? null : s.key)}
                      aria-expanded={isOpen}
                      className="w-full text-left"
                    >
                      <span className="flex items-baseline justify-between gap-2">
                        <span className="text-petrol/70 hover:text-petrol min-w-0 truncate text-sm">
                          {isOpen ? '−' : '+'} {s.name}
                        </span>
                        <span className="text-petrol/60 shrink-0 text-sm tabular-nums">
                          {Math.round(s.pct)}
                        </span>
                      </span>
                      <span
                        className="mt-1 block h-2 w-full overflow-hidden rounded-r"
                        style={{ background: `${cor}1a` }}
                      >
                        <span
                          className="block h-full rounded-r-[4px] transition-[width] duration-500"
                          style={{ width: `${Math.max(1.5, s.pct)}%`, background: `${cor}b3` }}
                        />
                      </span>
                    </button>

                    {isOpen && s.itens.length > 0 && (
                      <div className="border-petrol/15 mt-3 space-y-3 border-l pl-3 sm:pl-4">
                        {agrupaPorBloco(s.itens).map(([bloco, lista]) => (
                          <div key={bloco}>
                            {bloco && (
                              <p className="label-upper text-petrol/45 text-[10px]">{bloco}</p>
                            )}
                            <ul className="mt-1 space-y-1.5">
                              {lista.map((it) => (
                                <li
                                  key={it.id}
                                  className="border-petrol/10 flex flex-col gap-1 border-b pb-1.5 last:border-b-0 sm:flex-row sm:items-baseline sm:justify-between sm:gap-4"
                                >
                                  <span className="text-petrol min-w-0 text-sm">{it.name}</span>
                                  {it.levelName && (
                                    <span className="text-petrol/60 shrink-0 text-sm">
                                      {it.levelName}
                                    </span>
                                  )}
                                </li>
                              ))}
                            </ul>
                          </div>
                        ))}
                      </div>
                    )}
                  </li>
                );
              })}
            </ul>
          </section>
        );
      })}
    </div>
  );
}

function agrupaPorBloco(itens: BreakdownPilar['subpilares'][number]['itens']) {
  const m = new Map<string, typeof itens>();
  for (const it of itens) {
    const arr = m.get(it.bloco) ?? [];
    arr.push(it);
    m.set(it.bloco, arr);
  }
  return [...m.entries()];
}
