'use client';

import type { DeckTable, DeckTableCol, DeckTableRow } from '@plenya/types';

import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Field } from './field';
import { ListEditor } from './list-editor';

/**
 * A tabela — o bloco mais usado do deck real (8 dos 20 slides de um, 9 dos 21 do outro), e por
 * isso o editor que mais merece cuidado.
 *
 * O modelo é por COLUNA e não por campos fixos, porque as combinações mudam de slide para slide:
 * "exame | decide o quê | prioridade" num, "o quê | quanto | por quê" no outro.
 *
 * O estilo da coluna não é cosmético e o editor precisa dizer isso: `dose` NÃO quebra linha, então
 * prosa ali vaza o slide para fora da moldura. E é o estilo da coluna que faz a célula ser tratada
 * como número pela triagem do servidor.
 */
const ESTILOS: { valor: string; rotulo: string; ajuda: string }[] = [
  { valor: 'texto', rotulo: 'Texto', ajuda: 'quebra linha normalmente' },
  { valor: 'why', rotulo: 'Explicação', ajuda: 'cinza, menor — a coluna que explica' },
  { valor: 'dose', rotulo: 'Dose', ajuda: 'NÃO quebra linha: número e dose, nunca prosa' },
  { valor: 'tag', rotulo: 'Selo', ajuda: 'rótulo curto, tipo "prioridade"' },
];

export function TableEditor({
  table,
  onChange,
  tetoLinhas,
  tetoColunas = 3,
}: {
  table: DeckTable;
  onChange: (v: DeckTable) => void;
  tetoLinhas?: number;
  tetoColunas?: number;
}) {
  const colunas = table.columns ?? [];
  const linhas = table.rows ?? [];

  const setColunas = (cols: DeckTableCol[]) => {
    // Mexer nas colunas tem que mexer nas células, senão a linha fica com mais ou menos células
    // do que colunas e o render descarta em silêncio.
    const rows = linhas.map((r) => {
      const cells = (r.cells ?? []).slice(0, cols.length);
      while (cells.length < cols.length) cells.push('');
      return { ...r, cells };
    });
    onChange({ ...table, columns: cols, rows });
  };

  return (
    <div className="space-y-3">
      <div>
        <p className="mb-1 text-xs font-medium">Colunas</p>
        <ListEditor<DeckTableCol>
          itens={colunas}
          onChange={setColunas}
          chave={(_, i) => `col-${i}`}
          teto={tetoColunas}
          motivoDoTeto={`Três colunas é o limite: nenhum deck real passou disso.`}
          novoItem={() => ({ label: '' })}
          rotuloAdicionar="Adicionar coluna"
          render={(col, i, atualiza) => (
            <div className="flex items-end gap-2">
              <div className="min-w-0 flex-1">
                <Field label={`Coluna ${i + 1}`} valor={col.label ?? ''} limite={26}>
                  <Input
                    value={col.label ?? ''}
                    onChange={(e) => atualiza({ ...col, label: e.target.value })}
                    className="h-8 text-sm"
                  />
                </Field>
              </div>
              <Select
                value={col.style || 'texto'}
                onValueChange={(v) => atualiza({ ...col, style: (v === 'texto' ? '' : v) as DeckTableCol['style'] })}
              >
                <SelectTrigger className="h-8 w-32 text-xs">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  {ESTILOS.map((e) => (
                    <SelectItem key={e.valor} value={e.valor} className="text-xs">
                      <span className="font-medium">{e.rotulo}</span>
                      <span className="block text-[10px] text-muted-foreground">{e.ajuda}</span>
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          )}
        />
      </div>

      <div>
        <p className="mb-1 text-xs font-medium">Linhas</p>
        <ListEditor<DeckTableRow>
          itens={linhas}
          onChange={(rows) => onChange({ ...table, rows })}
          chave={(_, i) => `row-${i}`}
          teto={tetoLinhas}
          motivoDoTeto={`Nenhuma tabela real passou de oito linhas; acima disso o slide costuma transbordar.`}
          novoItem={() => ({ cells: colunas.map(() => '') })}
          rotuloAdicionar="Adicionar linha"
          vazio={<p className="py-3 text-center text-[11px] text-muted-foreground">Sem linhas ainda.</p>}
          render={(row, ri, atualiza) => (
            <div className="space-y-1.5">
              <div className="grid gap-1.5" style={{ gridTemplateColumns: `repeat(${Math.max(colunas.length, 1)}, minmax(0,1fr))` }}>
                {colunas.map((col, ci) => (
                  <Input
                    key={ci}
                    value={row.cells?.[ci] ?? ''}
                    onChange={(e) => {
                      const cells = (row.cells ?? []).slice();
                      while (cells.length < colunas.length) cells.push('');
                      cells[ci] = e.target.value;
                      atualiza({ ...row, cells });
                    }}
                    placeholder={col.label ?? ''}
                    className="h-8 text-xs"
                    title={col.style === 'dose' ? 'coluna de dose: não quebra linha' : undefined}
                  />
                ))}
              </div>
              <label className="flex items-center gap-1.5 text-[11px] text-muted-foreground">
                <input
                  type="checkbox"
                  checked={Boolean(row.muted)}
                  onChange={(e) => atualiza({ ...row, muted: e.target.checked })}
                />
                Riscada
                <span>(o item considerado e descartado)</span>
              </label>
              <span className="sr-only">linha {ri + 1}</span>
            </div>
          )}
        />
      </div>

      <label className="flex items-center gap-1.5 text-[11px]">
        <input
          type="checkbox"
          checked={Boolean(table.dense)}
          onChange={(e) => onChange({ ...table, dense: e.target.checked })}
        />
        Compacta
        <span className="text-muted-foreground">(aperta o espaçamento quando a tabela é longa)</span>
      </label>
    </div>
  );
}
