-- Expandir clinical_relevance para atender requisito de 1500-2000 caracteres
UPDATE score_items
SET clinical_relevance = 'A urocultura com contagem de colônias e antibiograma automatizado representa o padrão-ouro diagnóstico para infecções do trato urinário (ITU), permitindo identificação precisa do agente etiológico e seu perfil de susceptibilidade antimicrobiana. O limiar clássico de ≥100.000 UFC/mL (10⁵ CFU/mL) tem sido revisado por diretrizes contemporâneas da IDSA e AUA: em pacientes sintomáticos, contagens entre 10.000-100.000 UFC/mL (10⁴-10⁵) podem representar ITU verdadeira, especialmente em amostras obtidas por cateterismo ou punção suprapúbica. A automação por sistemas como VITEK 2 e Phoenix BD reduz o tempo de resposta (TAT) para 4-6 horas versus 24-48h dos métodos convencionais, crucial para iniciar terapia direcionada precocemente e reduzir morbimortalidade em sepse urinária. Em 2024, a resistência de E. coli produtora de ESBL (beta-lactamase de espectro estendido) atingiu níveis críticos globalmente, com 65.9% de resistência a cefepima e 56.1% a cefotaxima, tornando o antibiograma indispensável para evitar falha terapêutica. As diretrizes IDSA 2025 recomendam obtenção de cultura antes de iniciar antibioticoterapia em ITU complicada, cistite recorrente e pielonefrite, enfatizando de-escalation baseado em resultados para reduzir pressão seletiva. A interpretação adequada diferencia contaminação (múltiplos organismos, <10⁴ UFC/mL, células epiteliais >10/campo) de infecção verdadeira, reduzindo uso inapropriado de antimicrobianos e custos hospitalares. O antibiograma automatizado detecta resistências emergentes (carbapenemases, fluoroquinolonas, colistina) essenciais para vigilância epidemiológica institucional e controle de surtos nosocomiais.',
    updated_at = CURRENT_TIMESTAMP
WHERE id = '7ec43763-493a-481b-9865-4f793840ab20';

SELECT
    id,
    name,
    length(clinical_relevance) as clinical_len,
    length(patient_explanation) as patient_len,
    length(conduct) as conduct_len
FROM score_items
WHERE id = '7ec43763-493a-481b-9865-4f793840ab20';
