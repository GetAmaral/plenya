"""
Geradores adicionais - Triglicérides e Creatinina
"""

from typing import Dict


def generate_triglicerides_texts() -> Dict[str, str]:
    """Gera textos para Triglicérides"""

    clinical_relevance = """Os triglicerídeos são a principal forma de armazenamento de energia no organismo. Na medicina funcional integrativa, triglicerídeos elevados são reconhecidos como um dos marcadores mais sensíveis e precoces de resistência insulínica, síndrome metabólica e disfunção metabólica - frequentemente alterando-se ANTES da glicemia de jejum.

**Valores e Interpretação:**
- Ótimo funcional: <100 mg/dL (idealmente <80 mg/dL)
- Limítrofe: 100-150 mg/dL (resistência insulínica incipiente)
- Alto: 150-200 mg/dL (resistência insulínica estabelecida)
- Muito alto: 200-500 mg/dL (risco cardiovascular significativo)
- Extremamente alto: >500 mg/dL (risco de pancreatite aguda)

**Relação TG/HDL:** Um dos melhores preditores de resistência insulínica e risco cardiovascular. Ideal <2, melhor <1. Valores >3 indicam resistência insulínica grave.

**Fisiopatologia:** Triglicerídeos elevam-se principalmente por excesso de carboidratos refinados e açúcares. Frutose industrial é especialmente lipogênica, convertida diretamente em triglicerídeos no fígado via lipogênese de novo. Resistência insulínica, álcool, sedentarismo e obesidade visceral são outros fatores importantes."""

    patient_explanation = """Triglicerídeos são a gordura que circula no sangue para ser usada como energia ou guardada. Quando você come, especialmente carboidratos e açúcares, o fígado transforma o excesso em triglicerídeos.

Ter triglicerídeos altos significa que seu corpo está com dificuldade para processar açúcar e gordura, geralmente porque está comendo muito carboidrato (pão, massa, doce, refrigerante) ou bebendo álcool demais.

O ideal é abaixo de 100 mg/dL. Acima de 150 já mostra resistência à insulina. Acima de 500 é perigoso e pode causar pancreatite.

A boa notícia: triglicerídeos respondem MUITO rápido a mudanças na alimentação. Se você cortar açúcares, pães, massas e álcool, em 2-4 semanas podem cair pela metade."""

    conduct = """**Abordagem Terapêutica:**

**Triglicerídeos 100-200 mg/dL:**

**1. Intervenção Nutricional (mais efetiva que medicamentos):**
- Redução DRÁSTICA de carboidratos refinados: eliminar açúcar, refrigerantes, sucos, doces, pães brancos, massas
- Dieta Low Carb ou Cetogênica: redução de 50-70% em 8-12 semanas
- Eliminar ou reduzir álcool drasticamente
- Aumentar ômega-3: peixes gordos 3-4x/semana
- Aumentar fibras solúveis: aveia, linhaça, psyllium

**2. Exercício Físico:**
- Aeróbico 150-300 min/semana
- HIIT 2-3x/semana
- Musculação 3-4x/semana

**3. Suplementação:**
- Ômega-3 (EPA+DHA) 2-4g/dia: redução de 20-50% (MAIS EFETIVO)
- Bergamota (Vasguard) 1.000 mg/dia
- Berberina 500 mg 2-3x/dia (melhora resistência insulínica)
- Magnésio 400-600 mg/dia
- Cromo picolinato 200-400 mcg/dia

**4. Jejum Intermitente 16/8**

**Triglicerídeos >500 mg/dL (URGÊNCIA - risco pancreatite):**
- Dieta extremamente restrita em gorduras temporariamente
- Zero álcool
- Fenofibrato 160 mg/dia
- Ômega-3 4g/dia
- Reavaliar em 1-2 semanas

**Meta Terapêutica:**
- Ideal: <80 mg/dL
- Aceitável: <100 mg/dL
- Relação TG/HDL: <2 (idealmente <1)

**Monitoramento:** Triglicerídeos a cada 4-8 semanas. Resposta esperada: redução de 50-70% em 8-12 semanas com dieta low carb + ômega-3."""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }


def generate_creatinina_texts() -> Dict[str, str]:
    """Gera textos para Creatinina"""

    clinical_relevance = """A creatinina é um produto final do metabolismo da creatina muscular, excretada quase exclusivamente pelos rins via filtração glomerular. É o marcador mais utilizado para avaliar a função renal.

**Valores Funcionais Ótimos:**
- Homens: 0,8-1,1 mg/dL
- Mulheres: 0,6-0,9 mg/dL
- TFGe (taxa de filtração glomerular estimada): >90 mL/min/1,73m²

**Limitações da Creatinina:**
1. Insensibilidade precoce: permanece normal até perda de 40-50% da função renal
2. Influência da massa muscular: atletas têm valores mais altos; idosos sarcopênicos têm valores falsamente normais
3. Variabilidade dietética: consumo de carne eleva transitoriamente

**Marcadores Complementares Essenciais:**
- Taxa de Filtração Glomerular estimada (TFGe CKD-EPI): mais sensível que creatinina isolada
- Cistatina C: não depende de massa muscular, detecta disfunção precoce
- Microalbuminúria: marcador precoce de lesão renal (especialmente em diabetes/hipertensão)
- Ácido úrico: frequentemente elevado em doença renal incipiente

**Causas de Creatinina Elevada:**
- Pré-renais: desidratação, hipovolemia
- Renais: doença renal crônica (diabetes, hipertensão), nefrite, nefrotóxicos (AINEs, aminoglicosídeos)
- Pós-renais: obstrução urinária"""

    patient_explanation = """A creatinina é uma "sujeira" que os músculos produzem o tempo todo e que os rins filtram e eliminam pela urina. Medir a creatinina no sangue mostra se os rins estão funcionando bem.

Quando a creatinina está alta, significa que os rins não estão conseguindo filtrar direito. Isso pode acontecer por diabetes, pressão alta, desidratação ou uso de remédios que prejudicam os rins (anti-inflamatórios).

O problema é que a creatinina só sobe quando você já perdeu quase metade da função dos rins. Por isso é importante fazer outros exames junto (como a TFG - taxa de filtração) para detectar problemas mais cedo.

Pessoas com mais músculo (atletas) têm creatinina naturalmente mais alta. Pessoas com pouco músculo (idosos) podem ter creatinina normal mesmo com os rins já funcionando mal."""

    conduct = """**Interpretação Funcional:**

**Creatinina Normal:**
- Homens: 0,8-1,1 mg/dL
- Mulheres: 0,6-0,9 mg/dL
- TFGe: >90 mL/min/1,73m²

**Creatinina Limítrofe (investigar):**
- Homens: 1,1-1,3 mg/dL
- Mulheres: 0,9-1,1 mg/dL
- TFGe: 60-89 mL/min/1,73m² (DRC estágio 2)

**Creatinina Elevada:**
- >1,3 mg/dL (homens), >1,1 mg/dL (mulheres)
- TFGe: <60 mL/min/1,73m² (DRC estágio 3+)

**Investigação Complementar:**
- TFGe (CKD-EPI): obrigatório
- Cistatina C: detecta disfunção precoce
- Microalbuminúria (albumina urinária): marcador precoce
- Exame de urina (EAS): proteinúria, hematúria
- Glicemia, HbA1c: diabetes é principal causa de DRC
- Pressão arterial: hipertensão é segunda causa
- Ultrassom renal: avaliar tamanho, obstrução

**Abordagem Terapêutica:**

**Creatinina Limítrofe ou TFGe 60-89 (DRC estágio 2):**

**1. Proteção Renal - Controle de Fatores de Risco:**
- Controle glicêmico rigoroso: HbA1c <7% (diabéticos)
- Controle pressórico: Meta <130/80 mmHg, preferir IECA ou BRA (nefroprotetores)
- Controle lipídico: LDL <100 mg/dL

**2. Hidratação Adequada:**
- 30-35 mL/kg/dia de água (individualizar)

**3. Nutrição:**
- Proteínas: Moderar ~0,8-1g/kg/dia
- Sódio: <2.000-2.300 mg/dia
- Monitorar potássio e fósforo

**4. Evitar Nefrotóxicos:**
- AINEs (ibuprofeno, diclofenaco): usar com cautela
- Contraste iodado: hidratar bem antes e depois
- Aminoglicosídeos, vancomicina: monitorar níveis

**5. Suplementação:**
- Magnésio: 400-600 mg/dia (deficiência piora função renal)
- Ômega-3: 2-3g/dia (anti-inflamatório renal)
- Antioxidantes: NAC 600-1.200 mg/dia, ácido alfa-lipoico
- Vitamina D: corrigir deficiência (alvo 30-50 ng/mL)
- Coenzima Q10: 100-200 mg/dia

**Creatinina Elevada ou TFGe <60 (DRC estágio 3+):**
- Encaminhamento ao nefrologista obrigatório
- Restrição proteica: 0,6-0,8 g/kg/dia
- Controle rigoroso de fósforo e potássio
- Monitoramento frequente

**Monitoramento:**
- DRC estágio 2: creatinina, TFGe semestral
- DRC estágio 3: trimestral
- Eletrólitos, hemograma conforme estágio"""

    return {
        "clinicalRelevance": clinical_relevance,
        "patientExplanation": patient_explanation,
        "conduct": conduct
    }
