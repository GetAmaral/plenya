'use client';

import type { DeckDose, DeckDoseGroup, DeckTakeawayBox } from '@plenya/types';

import { Input } from '@/components/ui/input';
import { Field } from './field';
import { ListEditor } from './list-editor';
import { RichTextField } from './rich-text-field';

/**
 * "Para levar" — o que o paciente começa agora.
 *
 * `highlight` é o que MUDA o tratamento, e é um só: dois destaques não destacam nada. Os grupos
 * saem lado a lado num grid fixo de três colunas, então o teto de três não é conselho, é o layout.
 */
export function TakeawayEditor({
  take,
  onChange,
  tetoGrupos = 3,
}: {
  take: DeckTakeawayBox;
  onChange: (v: DeckTakeawayBox) => void;
  tetoGrupos?: number;
}) {
  const h = take.highlight ?? {};
  const grupos = take.groups ?? [];

  return (
    <div className="space-y-3">
      <div className="rounded border border-dashed p-2">
        <p className="mb-1.5 text-[11px] font-medium">Destaque</p>
        <div className="grid gap-2 sm:grid-cols-2">
          <Field label="O quê" valor={h.name ?? ''} limite={26}>
            <Input
              value={h.name ?? ''}
              onChange={(e) => onChange({ ...take, highlight: { ...h, name: e.target.value } })}
              className="h-8 text-sm"
            />
          </Field>
          <Field label="Quando" valor={h.when ?? ''} limite={26}>
            <Input
              value={h.when ?? ''}
              onChange={(e) => onChange({ ...take, highlight: { ...h, when: e.target.value } })}
              placeholder="Uma vez por semana"
              className="h-8 text-sm"
            />
          </Field>
          <Field label="Dose" valor={h.dose ?? ''} limite={12}>
            <Input
              value={h.dose ?? ''}
              onChange={(e) => onChange({ ...take, highlight: { ...h, dose: e.target.value } })}
              className="h-8 text-sm"
            />
          </Field>
          <Field label="Unidade" valor={h.unit ?? ''} limite={18}>
            <Input
              value={h.unit ?? ''}
              onChange={(e) => onChange({ ...take, highlight: { ...h, unit: e.target.value } })}
              placeholder="mg por semana"
              className="h-8 text-sm"
            />
          </Field>
        </div>
        <div className="mt-2">
          <Field label="Observação" valor={h.obs ?? ''} limite={80}>
            <Input
              value={h.obs ?? ''}
              onChange={(e) => onChange({ ...take, highlight: { ...h, obs: e.target.value } })}
              className="h-8 text-sm"
            />
          </Field>
        </div>
      </div>

      <div>
        <p className="mb-1 text-xs font-medium">Grupos</p>
        <ListEditor<DeckDoseGroup>
          itens={grupos}
          onChange={(g) => onChange({ ...take, groups: g })}
          chave={(_, i) => `grupo-${i}`}
          teto={tetoGrupos}
          motivoDoTeto="Três grupos é o limite do layout: eles saem lado a lado num grid de três colunas."
          novoItem={() => ({ title: '', items: [] })}
          rotuloAdicionar="Adicionar grupo"
          render={(g, _i, atualiza) => (
            <div className="space-y-2">
              <Field label="Título do grupo" valor={g.title ?? ''} limite={22}>
                <Input
                  value={g.title ?? ''}
                  onChange={(e) => atualiza({ ...g, title: e.target.value })}
                  placeholder="Todo dia"
                  className="h-8 text-sm"
                />
              </Field>
              <ListEditor<DeckDose>
                itens={g.items ?? []}
                onChange={(items) => atualiza({ ...g, items })}
                chave={(_, j) => `item-${j}`}
                novoItem={() => ({ name: '' })}
                rotuloAdicionar="Adicionar item"
                render={(it, _j, up) => (
                  <div className="grid gap-1.5 sm:grid-cols-[1fr_auto]">
                    <Input
                      value={it.name ?? ''}
                      onChange={(e) => up({ ...it, name: e.target.value })}
                      placeholder="Creatina"
                      className="h-7 text-xs"
                    />
                    <Input
                      value={it.dose ?? ''}
                      onChange={(e) => up({ ...it, dose: e.target.value })}
                      placeholder="5 g"
                      className="h-7 w-24 text-xs"
                    />
                    <Input
                      value={it.sub ?? ''}
                      onChange={(e) => up({ ...it, sub: e.target.value })}
                      placeholder="observação"
                      className="col-span-full h-7 text-xs"
                    />
                  </div>
                )}
              />
            </div>
          )}
        />
      </div>

      <RichTextField
        label="Nota"
        limite={160}
        valor={take.note ?? ''}
        onChange={(v) => onChange({ ...take, note: v })}
        linhas={2}
      />
    </div>
  );
}
