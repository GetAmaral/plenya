'use client';

import type { DeckCard } from '@plenya/types';

import { Input } from '@/components/ui/input';
import { Field } from './field';
import { ListEditor } from './list-editor';
import { RichTextField } from './rich-text-field';

/**
 * Cartões. Usado por `two-cards` (sempre dois) e por `plan-step`.
 *
 * Nenhum campo é obrigatório no contrato do Go, então o editor tolera cartão vazio: o render não
 * quebra, e forçar preenchimento aqui atrapalharia quem está montando o slide aos poucos.
 *
 * Em `two-cards` os dois cartões são fixos: é o contraste entre o caminho descartado (`dim`) e o
 * que vale (`focus`) que ensina, e um cartão sozinho não contrasta com nada.
 */
export function CardsEditor({
  cards,
  onChange,
  teto,
  fixo,
}: {
  cards: DeckCard[];
  onChange: (v: DeckCard[]) => void;
  teto?: number;
  fixo?: boolean;
}) {
  return (
    <ListEditor<DeckCard>
      itens={cards}
      onChange={onChange}
      chave={(_, i) => `card-${i}`}
      teto={teto}
      motivoDoTeto={fixo ? 'Este slide é sempre de dois cartões.' : `Até ${teto} cartões neste slide.`}
      semAdicionar={fixo && cards.length >= 2}
      semRemover={fixo}
      novoItem={() => ({ kicker: '', body: '' })}
      rotuloAdicionar="Adicionar cartão"
      render={(card, i, atualiza) => (
        <div className="space-y-2">
          <Field label={`Cartão ${i + 1} · rótulo`} valor={card.kicker ?? ''} limite={48}>
            <Input
              value={card.kicker ?? ''}
              onChange={(e) => atualiza({ ...card, kicker: e.target.value })}
              placeholder="Caminho 1 · sobrecarga genética"
              className="h-8 text-sm"
            />
          </Field>
          <RichTextField
            label="Corpo"
            limite={220}
            valor={card.body ?? ''}
            onChange={(v) => atualiza({ ...card, body: v })}
            linhas={3}
          />
          <div className="flex gap-4 text-[11px]">
            <label className="flex items-center gap-1.5">
              <input
                type="checkbox"
                checked={Boolean(card.focus)}
                onChange={(e) => atualiza({ ...card, focus: e.target.checked, dim: e.target.checked ? false : card.dim })}
              />
              Em destaque
            </label>
            <label className="flex items-center gap-1.5">
              <input
                type="checkbox"
                checked={Boolean(card.dim)}
                onChange={(e) => atualiza({ ...card, dim: e.target.checked, focus: e.target.checked ? false : card.focus })}
              />
              Apagado
              <span className="text-muted-foreground">(o caminho descartado)</span>
            </label>
          </div>
        </div>
      )}
    />
  );
}
