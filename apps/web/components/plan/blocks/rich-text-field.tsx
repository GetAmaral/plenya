'use client';

import { useRef } from 'react';
import { Bold, CornerDownLeft, Italic } from 'lucide-react';

import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { Field } from './field';

/**
 * Campo de texto do deck, com as TRÊS marcações que o deck aceita.
 *
 * Deliberadamente não é o `rich-text-editor` (Tiptap) que o resto do EMR usa para nota clínica.
 * Aquele produz HTML de bloco — parágrafos, listas, divs — e o renderizador do deck aceita apenas
 * `em`, `strong`, `b`, `i`, `small` e `br` inline. Qualquer outra tag é escapada e apareceria
 * LITERAL na tela do paciente, como `<p>` escrito por extenso no meio da frase.
 *
 * Daí a barra de três botões: é exatamente o vocabulário disponível, e nada além dele.
 */
export interface RichTextFieldProps {
  label: string;
  hint?: string;
  limite?: number;
  valor: string;
  onChange: (v: string) => void;
  linhas?: number;
  placeholder?: string;
  /** O `em` do punch sai dourado no deck; vale dizer isso a quem escreve. */
  destaqueDourado?: boolean;
}

export function RichTextField({
  label,
  hint,
  limite,
  valor,
  onChange,
  linhas = 2,
  placeholder,
  destaqueDourado,
}: RichTextFieldProps) {
  const ref = useRef<HTMLTextAreaElement>(null);

  const envolve = (tag: string) => {
    const el = ref.current;
    if (!el) return;
    const ini = el.selectionStart ?? 0;
    const fim = el.selectionEnd ?? 0;
    const dentro = valor.slice(ini, fim);
    const novo = `${valor.slice(0, ini)}<${tag}>${dentro}</${tag}>${valor.slice(fim)}`;
    onChange(novo);
    // Devolve o cursor para dentro da marcação, senão continuar digitando escreve fora dela.
    requestAnimationFrame(() => {
      el.focus();
      const pos = ini + tag.length + 2;
      el.setSelectionRange(pos, pos + dentro.length);
    });
  };

  const insere = (texto: string) => {
    const el = ref.current;
    const ini = el?.selectionStart ?? valor.length;
    onChange(`${valor.slice(0, ini)}${texto}${valor.slice(ini)}`);
  };

  return (
    <Field label={label} hint={hint} limite={limite} valor={valor}>
      <div className="space-y-1">
        <Textarea
          ref={ref}
          value={valor}
          onChange={(e) => onChange(e.target.value)}
          rows={linhas}
          placeholder={placeholder}
          className="resize-y text-sm"
        />
        <div className="flex items-center gap-1">
          <Button type="button" size="icon" variant="ghost" className="h-6 w-6" onClick={() => envolve('em')}>
            <Italic className="h-3 w-3" />
            <span className="sr-only">ênfase</span>
          </Button>
          <Button type="button" size="icon" variant="ghost" className="h-6 w-6" onClick={() => envolve('strong')}>
            <Bold className="h-3 w-3" />
            <span className="sr-only">negrito</span>
          </Button>
          <Button type="button" size="icon" variant="ghost" className="h-6 w-6" onClick={() => insere('<br>')}>
            <CornerDownLeft className="h-3 w-3" />
            <span className="sr-only">quebra de linha</span>
          </Button>
          <span className="ml-1 text-[10px] text-muted-foreground">
            {destaqueDourado ? 'a ênfase sai dourada' : 'só ênfase, negrito e quebra'}
          </span>
        </div>
      </div>
    </Field>
  );
}
