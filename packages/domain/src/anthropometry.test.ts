import { describe, expect, it } from 'vitest';
import { DERIVED_METRICS, computeDerived, getDerivedMetric } from './anthropometry';

const homem = { PESO: 80, ALTURA: 180, ABDOMINAL_HOMEM: 90, QUADRIL: 100, PESCOCO_HOMEM: 40 };

describe('métricas antropométricas derivadas', () => {
  it('calcula o IMC a partir de peso e altura', () => {
    // 80 / 1,80² = 24,69
    expect(computeDerived(DERIVED_METRICS.IMC, homem)).toBe(24.7);
  });

  it('não calcula quando falta entrada', () => {
    expect(computeDerived(DERIVED_METRICS.IMC, { PESO: 80 })).toBeUndefined();
    expect(computeDerived(DERIVED_METRICS.IMC, {})).toBeUndefined();
  });

  it('aceita alternativas homem/mulher na mesma entrada', () => {
    const razao = DERIVED_METRICS.RAZAO_CINTURA_ALTURA;
    expect(computeDerived(razao, homem)).toBe(0.5);
    expect(computeDerived(razao, { ALTURA: 160, ABDOMINAL_MULHER: 72 })).toBe(0.45);
  });

  it('calcula razão cintura/quadril e relação pescoço-altura', () => {
    expect(computeDerived(DERIVED_METRICS.RAZAO_CINTURA_QUADRIL_HOMEM, homem)).toBe(0.9);
    // 40 cm / 1,80 m
    expect(computeDerived(DERIVED_METRICS.RELACAO_PESCOCO_ALTURA_HOMEM, homem)).toBe(22.2);
  });

  it('calcula o BRI pela fórmula de Thomas 2013', () => {
    // cintura 90, altura 180 → raio 14,32; semi-altura 90; e = 1 − (0,1591)² = 0,9747
    expect(computeDerived(DERIVED_METRICS.BRI, homem)).toBeCloseTo(3.32, 1);
  });

  it('não devolve NaN quando o BRI sai do domínio (cintura > π·altura)', () => {
    expect(computeDerived(DERIVED_METRICS.BRI, { ABDOMINAL_HOMEM: 600, ALTURA: 160 })).toBeUndefined();
  });

  it('calcula percentuais de massa e água', () => {
    expect(computeDerived(DERIVED_METRICS.MME_PESO, { MASSA_MUSCULAR_ESQUELETICA: 34, PESO: 80 })).toBe(42.5);
    expect(computeDerived(DERIVED_METRICS.GORDURA_CORPORAL_HOMEM, { MASSA_GORDA_TOTAL: 16, PESO: 80 })).toBe(20);
    expect(computeDerived(DERIVED_METRICS.RAZAO_AEC_ACT, { AGUA_EXTRACELULAR: 18, AGUA_CORPORAL_TOTAL: 48 })).toBe(37.5);
  });

  it('calcula índices de massa (FMI, ASMI, índice MME)', () => {
    expect(computeDerived(DERIVED_METRICS.FMI_FAT_MASS_INDEX_HOMEM, { MASSA_GORDA_TOTAL: 16, ALTURA: 180 })).toBe(4.9);
    expect(computeDerived(DERIVED_METRICS.ASMI_HOMEM, { MASSA_APENDICULAR: 24, ALTURA: 180 })).toBe(7.41);
  });

  it('protege contra divisão por zero', () => {
    expect(computeDerived(DERIVED_METRICS.IMC, { PESO: 80, ALTURA: 0 })).toBeUndefined();
    expect(computeDerived(DERIVED_METRICS.MME_PESO, { MASSA_MUSCULAR_ESQUELETICA: 34, PESO: 0 })).toBeUndefined();
  });

  it('só reconhece códigos registrados', () => {
    expect(getDerivedMetric('IMC')).toBeDefined();
    expect(getDerivedMetric('PESO')).toBeUndefined();
    expect(getDerivedMetric(null)).toBeUndefined();
  });
});
