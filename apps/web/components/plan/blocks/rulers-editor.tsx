'use client';

import { useMemo, useState } from 'react';
import { Check, ChevronsUpDown, Lock, RotateCcw } from 'lucide-react';
import type { DeckRuler, PlanDossier, PlanRuler } from '@plenya/types';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command';
import { Input } from '@/components/ui/input';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip';
import { Field } from './field';
import { ListEditor } from './list-editor';
import { RichTextField } from './rich-text-field';

/**
 * A régua é o átomo visual do deck, e o bloco onde a distinção entre DADO e AUTORIA fica visível.
 *
 * Do dossiê, travados: `code`, `unit`, `segments` e `history`. São a escala do escore aplicável a
 * este paciente e as medidas dele; reescrever qualquer um seria inventar exame.
 *
 * Autorais, os três únicos: `display` (o nome que o paciente reconhece, não o do catálogo), `sub`
 * (o que o exame mede, em poucas palavras) e `note` — que é onde mora o rótulo avaliativo quando o
 * título do slide não o traz. Barra colorida sem rótulo tem desempenho pior do que com rótulo.
 *
 * O eixo é o caso do meio: sai do dossiê, mas é o único número legitimamente afinado à mão, quando
 * um valor extremo esmaga a escala. Por isso vem travado com um botão para soltar, mostrando o
 * valor de origem ao lado — e o servidor recusa eixo que corte faixa ou ponto medido.
 */

/** Converte a régua do dossiê no bloco do deck, trazendo travado o que é dado. */
function doDossie(code: string, r: PlanRuler): DeckRuler {
  return {
    code,
    // "Ferritina - Mulheres Pós-Menopausa" é nome de catálogo; o paciente lê "Ferritina".
    display: (r.name ?? '').split(' - ')[0].trim(),
    sub: '',
    unit: r.unit ?? '',
    axis: (r.axis ?? [0, 1]) as number[],
    segments: (r.segments ?? []).map((s) => ({ level: s.level, a: s.a, b: s.b })),
    history: (r.history ?? []).map((p) => ({ value: p.value, text: p.text, date: p.date })),
  };
}

export function RulersEditor({
  rulers,
  onChange,
  dossier,
  teto,
}: {
  rulers: DeckRuler[];
  onChange: (v: DeckRuler[]) => void;
  dossier?: PlanDossier;
  teto?: number;
}) {
  return (
    <ListEditor<DeckRuler>
      itens={rulers}
      onChange={onChange}
      chave={(r, i) => `${r.code ?? 'sem'}-${i}`}
      teto={teto}
      motivoDoTeto="Quatro é o teto comprovado: com oito réguas o slide transborda, e há teste no backend provando."
      semAdicionar
      vazio={<p className="py-2 text-center text-[11px] text-muted-foreground">Nenhuma régua ainda.</p>}
      render={(r, _i, atualiza) => <ItemDeRegua ruler={r} onChange={atualiza} dossier={dossier} />}
    />
  );
}

/** O seletor fica fora da lista para poder aparecer mesmo com a lista vazia. */
export function AdicionarReguaDoDossie({
  dossier,
  onAdd,
  desabilitado,
  motivo,
}: {
  dossier?: PlanDossier;
  onAdd: (r: DeckRuler) => void;
  desabilitado?: boolean;
  motivo?: string;
}) {
  const [aberto, setAberto] = useState(false);

  const opcoes = useMemo(() => {
    const rs = Object.entries(dossier?.rulers ?? {}) as [string, PlanRuler][];
    return rs
      .filter(([, r]) => (r.history?.length ?? 0) > 0)
      .sort((a, b) => (a[1].name ?? '').localeCompare(b[1].name ?? '', 'pt-BR'));
  }, [dossier?.rulers]);

  const botao = (
    <Button
      type="button"
      size="sm"
      variant="outline"
      className="h-7 w-full justify-between text-xs"
      disabled={desabilitado}
    >
      Adicionar régua do prontuário
      <ChevronsUpDown className="ml-1 h-3 w-3 opacity-50" />
    </Button>
  );

  if (desabilitado) {
    return (
      <Tooltip>
        <TooltipTrigger asChild>
          <span className="block">{botao}</span>
        </TooltipTrigger>
        <TooltipContent className="max-w-xs text-xs">{motivo}</TooltipContent>
      </Tooltip>
    );
  }

  return (
    <Popover open={aberto} onOpenChange={setAberto}>
      <PopoverTrigger asChild>{botao}</PopoverTrigger>
      <PopoverContent className="w-80 p-0" align="start">
        <Command
          filter={(value, search) => {
            // Busca sem acento: quem digita "creatinina" não deveria depender do til.
            const norm = (s: string) => s.normalize('NFD').replace(/[̀-ͯ]/g, '').toLowerCase();
            return norm(value).includes(norm(search)) ? 1 : 0;
          }}
        >
          <CommandInput placeholder="Buscar exame do paciente" className="h-9 text-xs" />
          <CommandList>
            <CommandEmpty className="py-4 text-center text-xs text-muted-foreground">
              Nenhum exame com resultado para esta busca.
            </CommandEmpty>
            <CommandGroup>
              {opcoes.map(([code, r]) => {
                const ultimo = r.history?.[r.history.length - 1];
                return (
                  <CommandItem
                    key={code}
                    value={`${r.name} ${code}`}
                    onSelect={() => {
                      onAdd(doDossie(code, r));
                      setAberto(false);
                    }}
                    className="text-xs"
                  >
                    <span className="min-w-0 flex-1 truncate">{r.name}</span>
                    <span className="ml-2 shrink-0 text-[10px] text-muted-foreground">
                      {ultimo?.text ?? ''} {r.unit ?? ''}
                    </span>
                  </CommandItem>
                );
              })}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}

function ItemDeRegua({
  ruler,
  onChange,
  dossier,
}: {
  ruler: DeckRuler;
  onChange: (r: DeckRuler) => void;
  dossier?: PlanDossier;
}) {
  const [eixoLiberado, setEixoLiberado] = useState(false);
  const origem = ruler.code ? dossier?.rulers?.[ruler.code] : undefined;
  const eixoDoDossie = origem?.axis;
  const desviouDoEixo =
    eixoDoDossie != null &&
    ruler.axis != null &&
    (eixoDoDossie[0] !== ruler.axis[0] || eixoDoDossie[1] !== ruler.axis[1]);

  const ressincroniza = () => {
    if (!origem || !ruler.code) return;
    const novo = doDossie(ruler.code, origem);
    // Preserva o que é autoral: ressincronizar é atualizar o DADO, não apagar o texto.
    onChange({ ...novo, display: ruler.display, sub: ruler.sub, note: ruler.note });
  };

  return (
    <div className="space-y-2">
      <div className="flex flex-wrap items-center gap-1">
        <Badge variant="outline" className="h-5 gap-1 px-1.5 text-[10px] font-normal">
          <Lock className="h-2.5 w-2.5" />
          {ruler.code ?? 'sem código'}
        </Badge>
        {ruler.unit && (
          <Badge variant="outline" className="h-5 px-1.5 text-[10px] font-normal">
            {ruler.unit}
          </Badge>
        )}
        <Badge variant="outline" className="h-5 px-1.5 text-[10px] font-normal">
          {ruler.segments?.length ?? 0} faixas
        </Badge>
        <Badge variant="outline" className="h-5 px-1.5 text-[10px] font-normal">
          {ruler.history?.length ?? 0} medidas
        </Badge>
        {origem && (
          <Button
            type="button"
            size="sm"
            variant="ghost"
            className="ml-auto h-5 px-1 text-[10px] text-muted-foreground"
            onClick={ressincroniza}
            title="Traz o dado atual do prontuário e preserva o texto que você escreveu"
          >
            <RotateCcw className="mr-1 h-2.5 w-2.5" />
            Ressincronizar
          </Button>
        )}
      </div>
      <p className="text-[10px] text-muted-foreground">
        Código, unidade, faixas e medidas vêm do prontuário e não são editáveis aqui.
      </p>

      <Field label="Nome que o paciente lê" valor={ruler.display ?? ''} limite={28}>
        <Input
          value={ruler.display ?? ''}
          onChange={(e) => onChange({ ...ruler, display: e.target.value })}
          placeholder="Ferritina"
          className="h-8 text-sm"
        />
      </Field>
      <Field label="O que o exame mede" hint="Em poucas palavras: “estoque de ferro”." valor={ruler.sub ?? ''} limite={34}>
        <Input
          value={ruler.sub ?? ''}
          onChange={(e) => onChange({ ...ruler, sub: e.target.value })}
          className="h-8 text-sm"
        />
      </Field>
      <RichTextField
        label="Leitura"
        hint="É aqui que mora o rótulo avaliativo quando o título do slide não o traz."
        limite={70}
        valor={ruler.note ?? ''}
        onChange={(v) => onChange({ ...ruler, note: v })}
        linhas={1}
      />

      <div className="rounded border border-dashed p-2">
        <div className="flex items-center justify-between gap-2">
          <span className="text-[11px] font-medium">Eixo</span>
          <label className="flex items-center gap-1 text-[10px] text-muted-foreground">
            <input type="checkbox" checked={eixoLiberado} onChange={(e) => setEixoLiberado(e.target.checked)} />
            afinar
          </label>
        </div>
        <div className="mt-1 flex items-center gap-1">
          {[0, 1].map((i) => (
            <Input
              key={i}
              type="number"
              value={ruler.axis?.[i] ?? 0}
              disabled={!eixoLiberado}
              onChange={(e) => {
                const eixo = [...(ruler.axis ?? [0, 1])];
                eixo[i] = Number(e.target.value);
                onChange({ ...ruler, axis: eixo });
              }}
              className="h-7 text-xs"
            />
          ))}
        </div>
        {eixoDoDossie && (
          <p className="mt-1 flex items-center gap-1 text-[10px] text-muted-foreground">
            {desviouDoEixo ? (
              <>
                No prontuário: {eixoDoDossie[0]} a {eixoDoDossie[1]}
                <Button
                  type="button"
                  size="sm"
                  variant="ghost"
                  className="h-4 px-1 text-[10px]"
                  onClick={() => onChange({ ...ruler, axis: [...eixoDoDossie] })}
                >
                  voltar
                </Button>
              </>
            ) : (
              <>
                <Check className="h-2.5 w-2.5" /> igual ao do prontuário
              </>
            )}
          </p>
        )}
        <p className="mt-1 text-[10px] text-muted-foreground">
          Aperte só quando um valor extremo esmagar a escala. Eixo que deixa um ponto de fora esconde
          o dado sem avisar, e a publicação recusa.
        </p>
      </div>
    </div>
  );
}

export { doDossie as reguaDoDossie };
