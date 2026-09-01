'use client';

import type { DeckSummaryBlock, DeckSummaryCard, DeckSummaryLine, PlanDossier } from '@plenya/types';

import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Field } from './field';
import { ListEditor } from './list-editor';
import { AdicionarReguaDoDossie } from './rulers-editor';

/**
 * O resumo é o bloco mais fundo do contrato: cartão → linha → régua completa embutida. É também o
 * slide que o paciente mais relê, e onde o deck concentra mais número por centímetro.
 *
 * Por isso a linha tem `code`: sem essa âncora, o `value` é string solta e conferir a origem do
 * número passa a ser impossível. Escolher a régua do prontuário preenche o código junto — é o
 * caminho certo, e o campo de código fica visível para que a falta salte aos olhos.
 */
export function SummaryEditor({
  summary,
  onChange,
  dossier,
  tetoCartoes = 2,
  tetoLinhas = 4,
}: {
  summary: DeckSummaryBlock;
  onChange: (v: DeckSummaryBlock) => void;
  dossier?: PlanDossier;
  tetoCartoes?: number;
  tetoLinhas?: number;
}) {
  const cards = summary.cards ?? [];

  return (
    <div className="space-y-3">
      <ListEditor<DeckSummaryCard>
        itens={cards}
        onChange={(c) => onChange({ ...summary, cards: c })}
        chave={(_, i) => `sc-${i}`}
        teto={tetoCartoes}
        motivoDoTeto="Dois cartões: um do que está forte, outro do que está se movendo."
        novoItem={() => ({ title: '', tone: 'bom', lines: [] })}
        rotuloAdicionar="Adicionar cartão"
        render={(card, _i, atualiza) => (
          <div className="space-y-2">
            <div className="flex items-end gap-2">
              <div className="min-w-0 flex-1">
                <Field label="Título do cartão" valor={card.title ?? ''} limite={26}>
                  <Input
                    value={card.title ?? ''}
                    onChange={(e) => atualiza({ ...card, title: e.target.value })}
                    className="h-8 text-sm"
                  />
                </Field>
              </div>
              <Select value={card.tone || 'bom'} onValueChange={(v) => atualiza({ ...card, tone: v })}>
                <SelectTrigger className="h-8 w-28 text-xs">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="bom" className="text-xs">Bom</SelectItem>
                  <SelectItem value="ruim" className="text-xs">Se movendo</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <ListEditor<DeckSummaryLine>
              itens={card.lines ?? []}
              onChange={(lines) => atualiza({ ...card, lines })}
              chave={(_, j) => `sl-${j}`}
              teto={tetoLinhas}
              motivoDoTeto="Quatro linhas por cartão é o teto confortável."
              semAdicionar
              render={(l, _j, up) => (
                <div className="space-y-1.5">
                  <div className="grid gap-1.5 sm:grid-cols-[1fr_5rem_4rem]">
                    <Input
                      value={l.name ?? ''}
                      onChange={(e) => up({ ...l, name: e.target.value })}
                      placeholder="Rim"
                      className="h-7 text-xs"
                    />
                    <Input
                      value={l.value ?? ''}
                      onChange={(e) => up({ ...l, value: e.target.value })}
                      placeholder="8,57"
                      className="h-7 text-xs"
                    />
                    <Input
                      value={l.unit ?? ''}
                      onChange={(e) => up({ ...l, unit: e.target.value })}
                      placeholder="mg/g"
                      className="h-7 text-xs"
                    />
                  </div>
                  <Input
                    value={l.sub ?? ''}
                    onChange={(e) => up({ ...l, sub: e.target.value })}
                    placeholder="sem perda de proteína"
                    className="h-7 text-xs"
                  />
                  <p className="text-[10px] text-muted-foreground">
                    {l.code ? (
                      <>exame de origem: {l.code}</>
                    ) : (
                      <span className="text-amber-700">
                        sem exame de origem: o número desta linha não tem como ser conferido
                      </span>
                    )}
                  </p>
                </div>
              )}
            />

            <AdicionarReguaDoDossie
              dossier={dossier}
              desabilitado={(card.lines?.length ?? 0) >= tetoLinhas}
              motivo={`Quatro linhas por cartão é o teto.`}
              onAdd={(r) => {
                const ultimo = r.history?.[r.history.length - 1];
                const linha: DeckSummaryLine = {
                  name: r.display,
                  code: r.code,
                  value: ultimo?.text ?? '',
                  unit: r.unit,
                  ruler: r,
                };
                atualiza({ ...card, lines: [...(card.lines ?? []), linha] });
              }}
            />
          </div>
        )}
      />

      <Field label="Título dos passos" valor={summary.stepsTitle ?? ''} limite={26}>
        <Input
          value={summary.stepsTitle ?? ''}
          onChange={(e) => onChange({ ...summary, stepsTitle: e.target.value })}
          placeholder="O que vamos fazer"
          className="h-8 text-sm"
        />
      </Field>

      <ListEditor<string>
        itens={summary.steps ?? []}
        onChange={(steps) => onChange({ ...summary, steps })}
        chave={(_, i) => `st-${i}`}
        teto={4}
        motivoDoTeto="Quatro passos é o teto confortável."
        novoItem={() => ''}
        rotuloAdicionar="Adicionar passo"
        render={(s, _i, up) => (
          <Input value={s} onChange={(e) => up(e.target.value)} className="h-7 text-xs" />
        )}
      />
    </div>
  );
}
