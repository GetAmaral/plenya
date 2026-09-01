'use client';

import { useState } from 'react';
import { ChevronDown } from 'lucide-react';

import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Field } from './field';
import { RichTextField } from './rich-text-field';
import type { DeckSlide } from '@plenya/types';

/**
 * O envelope do slide.
 *
 * `eyebrow`, `title` e `punch` ficam sempre visíveis, e não por gosto: contados nos dois decks v2
 * reais, aparecem em 100%, ~90% e ~85% dos slides. `lede`, `kicker`, `source` e `legend` aparecem
 * em menos de um terço e ficam atrás de "mais opções" — mostrar sete campos onde três bastam faz
 * o cartão parecer trabalho.
 */
export function SlideHeaderFields({
  slide,
  onChange,
}: {
  slide: DeckSlide;
  onChange: (patch: Partial<DeckSlide>) => void;
}) {
  const [abertos, setAbertos] = useState(
    Boolean(slide.lede || slide.kicker || slide.source || slide.legend),
  );

  return (
    <div className="space-y-3">
      <Field label="Chapéu" hint="O bloco a que o slide pertence, e a contagem." valor={slide.eyebrow ?? ''} limite={54}>
        <Input
          value={slide.eyebrow ?? ''}
          onChange={(e) => onChange({ eyebrow: e.target.value })}
          placeholder="O que está se movendo · 1 de 3"
          className="h-8 text-sm"
        />
      </Field>

      <RichTextField
        label="Título"
        hint="Uma AFIRMAÇÃO, não um rótulo: “A ferritina dobrou em dois anos”, não “Ferritina”."
        limite={64}
        valor={slide.title ?? ''}
        onChange={(v) => onChange({ title: v })}
        linhas={2}
      />

      <RichTextField
        label="Fecho"
        hint="A frase que fecha o slide com a consequência."
        limite={130}
        valor={slide.punch ?? ''}
        onChange={(v) => onChange({ punch: v })}
        linhas={2}
        destaqueDourado
      />

      <Button
        type="button"
        variant="ghost"
        size="sm"
        className="h-6 px-1 text-[11px] text-muted-foreground"
        onClick={() => setAbertos((v) => !v)}
      >
        <ChevronDown className={`mr-1 h-3 w-3 transition-transform ${abertos ? 'rotate-180' : ''}`} />
        Campos raros
      </Button>

      {abertos && (
        <div className="space-y-3 border-l-2 pl-3">
          <RichTextField
            label="Abertura"
            hint="Parágrafo de entrada. Raro fora da capa e do resumo."
            limite={200}
            valor={slide.lede ?? ''}
            onChange={(v) => onChange({ lede: v })}
          />
          <RichTextField
            label="Apoio"
            limite={200}
            valor={slide.kicker ?? ''}
            onChange={(v) => onChange({ kicker: v })}
          />
          <Field label="Rodapé" hint="Nota de fonte ou ressalva." valor={slide.source ?? ''} limite={160}>
            <Input
              value={slide.source ?? ''}
              onChange={(e) => onChange({ source: e.target.value })}
              className="h-8 text-sm"
            />
          </Field>
          <label className="flex items-center gap-2 text-xs">
            <input
              type="checkbox"
              checked={Boolean(slide.legend)}
              onChange={(e) => onChange({ legend: e.target.checked })}
            />
            Mostrar a legenda da rampa de cores
            <span className="text-[10px] text-muted-foreground">(só no primeiro slide de régua do bloco)</span>
          </label>
        </div>
      )}
    </div>
  );
}
