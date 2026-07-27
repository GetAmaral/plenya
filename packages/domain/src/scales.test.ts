import { describe, expect, it } from 'vitest';
import { SCALE_REGISTRY, scaleTotal } from './scales';

const DUBOIS_IMEDIATO = SCALE_REGISTRY['5_PALAVRAS_DE_DUBOIS_IMEDIATO_5'];
const DUBOIS_TARDIO = SCALE_REGISTRY['5_PALAVRAS_DE_DUBOIS_TARDIO_5'];

describe('Dubois — score total ponderado', () => {
  it('vale 10 por fase quando as 5 palavras vêm espontaneamente', () => {
    const todasEspontaneas = { 0: 2, 1: 2, 2: 2, 3: 2, 4: 2 };
    expect(scaleTotal(DUBOIS_IMEDIATO, todasEspontaneas)).toBe(10);
    expect(DUBOIS_IMEDIATO.maxScore).toBe(10);
    expect(DUBOIS_TARDIO.maxScore).toBe(10);
  });

  it('pontua a evocação com dica pela METADE da espontânea', () => {
    // 4 espontâneas + 1 com dica = 8 + 1 = 9 (antes dava 5, igual a 5 espontâneas)
    expect(scaleTotal(DUBOIS_IMEDIATO, { 0: 2, 1: 2, 2: 2, 3: 2, 4: 1 })).toBe(9);
    // 5 com dica = 5, contra 10 se todas fossem espontâneas
    expect(scaleTotal(DUBOIS_TARDIO, { 0: 1, 1: 1, 2: 1, 3: 1, 4: 1 })).toBe(5);
  });

  it('não pontua a palavra não evocada nem com dica', () => {
    expect(scaleTotal(DUBOIS_TARDIO, { 0: 2, 1: 2, 2: 2, 3: 2, 4: 0 })).toBe(8);
    expect(scaleTotal(DUBOIS_TARDIO, { 0: 0, 1: 0, 2: 0, 3: 0, 4: 0 })).toBe(0);
  });
});

describe('Span de dígitos', () => {
  it('pontua o maior comprimento acertado', () => {
    const direto = SCALE_REGISTRY['SPAN_DE_DIGITOS_DIRETO_8'];
    expect(scaleTotal(direto, { 3: 1, 4: 1, 5: 1, 6: 0 })).toBe(5);
    expect(scaleTotal(direto, { 3: 0 })).toBe(0);
  });

  it('cada comprimento traz 2 tentativas', () => {
    const direto = SCALE_REGISTRY['SPAN_DE_DIGITOS_DIRETO_8'];
    const inverso = SCALE_REGISTRY['SPAN_DE_DIGITOS_INVERSO_7'];
    for (const def of [direto, inverso]) {
      const spec = def.administration;
      if (spec?.type !== 'digit_span') throw new Error('esperado digit_span');
      for (const [len, seqs] of Object.entries(spec.sequencesByLength)) {
        expect(seqs).toHaveLength(2);
        // o rótulo "N dígitos" tem de bater com o tamanho real de cada sequência
        for (const seq of seqs) expect(seq).toHaveLength(Number(len));
      }
    }
  });
});
