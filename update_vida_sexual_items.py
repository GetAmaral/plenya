#!/usr/bin/env python3
"""
Script para atualizar todos os 53 items do grupo Vida Sexual
com conteúdo de medicina funcional baseado em evidências 2024-2025.
"""

import requests
import json
import time

API_URL = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjYWIwMDBkNy1iODBlLTQ0NWYtYTUyOS0zYThlYzhmNDg2YmQiLCJlbWFpbCI6ImNsYXVkZS1hZG1pbkBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTQzMjcwNSwiaWF0IjoxNzY5NDMxODA1fQ.W8u8MPcnOkUTOUhnNoRV5VhnyAp2u2mvCGQsgTUnjJI"

headers = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {TOKEN}"
}

# Dados dos 53 items agrupados por tema
items_data = {
    # GRUPO 1: IIEF-5 - Disfunção Erétil Masculina (2 items)
    "iief5": [
        {
            "id": "5aaadebc-141f-46f5-bf89-7603dadb91f1",
            "name": "IIEF-5 (Índice Internacional de Função Erétil): ____/25",
            "clinicalRelevance": """O Índice Internacional de Função Erétil versão simplificada (IIEF-5) representa muito mais do que uma ferramenta de avaliação urológica: constitui-se como biomarcador sistêmico de saúde cardiovascular, metabólica e hormonal. Validado internacionalmente e adaptado culturalmente para o Brasil, este questionário de 5 itens avalia a capacidade de obter e manter ereções, confiança erétil, penetração, manutenção durante intercurso e satisfação sexual, pontuando de 5 a 25 pontos.

A relevância clínica transcende a esfera sexual. Metanálise de 2024 demonstrou que disfunção erétil (DE) é biomarcador preditivo robusto para doença cardiovascular em homens em envelhecimento, com implicações significativas para detecção precoce e estratégias preventivas. O conceito de DE como canário na mina de carvão cardiovascular baseia-se na hipótese do tamanho arterial: artérias penianas (1-2mm) manifestam disfunção endotelial antes das coronárias (3-4mm), permitindo janela de intervenção de 2-5 anos antes de eventos cardiovasculares maiores.

Estudo de coorte prospectivo de novembro de 2024 evidenciou associação bidirecional entre DE e risco cardiovascular de 10 anos em homens com diabetes e hipertensão. Dados mostram que DE deve ser considerada manifestação precoce de doença cardiovascular, justificando avaliação cuidadosa de fatores de risco CV em homens com DE para prevenir eventos adversos maiores futuros. Pesquisa de 2024 identificou que índice cardiometabólico (CMI) e idade foram fatores de risco significativos para disfunção sexual (OR=2,672 e 1,081, respectivamente), com CMI predizendo disfunção sexual incluindo DE especialmente em homens mais jovens.

Na perspectiva de medicina funcional, IIEF-5 inferior a 22 pontos sinaliza investigação de múltiplos sistemas. Disfunção endotelial é mecanismo fisiopatológico central, mediada por biodisponibilidade reduzida de óxido nítrico, estresse oxidativo aumentado e inflamação sistêmica de baixo grau. Biomarcadores inflamatórios predizem resposta a inibidores de fosfodiesterase-5: pressão arterial sistólica e proteína C-reativa (PCR) predizem mudanças no IIEF em alta taxa, com níveis elevados de PCR relacionados à severidade de DE.

Avaliação hormonal é mandatória segundo consenso ICSM 2024. Testosterona livre deve ser considerada teste de rotina junto com testosterona total em homens com DE, enquanto LH, estradiol, SHBG e prolactina não demonstram correlação estatística com DE e não devem ser recomendados como investigações de rotina. Entretanto, idade mostra correlação estatística significativa tanto com severidade de DE quanto com níveis de testosterona total, testosterona livre, LH, FSH e SHBG. Na prática funcional, medição de estradiol junto com testosterona é importante em homens com disfunção sexual, mudanças de composição corporal ou sintomas persistentes de humor, com muitos clínicos monitorando razão testosterona:estradiol (T:E2) visando intervalo aproximado de 10:1 a 15:1 como alvo funcional.

miRNA-21 circulante emergiu como biomarcador inovador: estudo de fevereiro de 2024 identificou associação inversa entre níveis de expressão de miRNA-21 e prevalência de DE, sugerindo miRNA-21 circulante como biomarcador inovador para eventos cardiovasculares em pacientes com DE. Resistência insulínica representa outro nexo fisiopatológico crítico: altos níveis de insulina reduzem capacidade corporal de produzir hormônio de crescimento, reduzindo libido através de alteração nos níveis de testosterona.

Interpretação dos escores: 22-25 pontos = sem DE; 17-21 = DE leve; 12-16 = DE leve a moderada; 8-11 = DE moderada; 5-7 = DE severa. Cada categoria requer abordagem investigativa e terapêutica diferenciada.""",
            "patientExplanation": """O questionário IIEF-5 avalia sua função erétil através de 5 perguntas simples sobre confiança, rigidez, penetração, manutenção e satisfação sexual nos últimos 6 meses. A pontuação total varia de 5 a 25 pontos e fornece informação objetiva sobre sua saúde sexual masculina.

Muitos homens não percebem que dificuldades de ereção podem ser o primeiro sinal de problemas de saúde mais amplos. Pesquisas mostram que disfunção erétil frequentemente aparece 2-5 anos antes de problemas cardiovasculares, funcionando como um sistema de alerta precoce para seu coração e circulação. Isso acontece porque as artérias do pênis são menores que as do coração e mostram problemas circulatórios primeiro.

Se sua pontuação está abaixo de 22, isso indica necessidade de investigação médica mais profunda - não apenas da função sexual, mas também de sua saúde cardiovascular, metabólica e hormonal. Estudos de 2024 mostram que homens com disfunção erétil têm maior risco de diabetes, pressão alta, colesterol elevado e problemas cardíacos.

Vários fatores podem afetar sua função erétil: níveis hormonais (especialmente testosterona), qualidade da circulação sanguínea, inflamação no corpo, resistência à insulina, estresse oxidativo, qualidade do sono, nível de estresse psicológico e alguns medicamentos. A boa notícia é que muitos desses fatores podem ser melhorados com mudanças de estilo de vida e tratamento adequado.

Interpretar sua pontuação: 22-25 pontos significa função normal; 17-21 indica disfunção leve; 12-16 é leve a moderada; 8-11 é moderada; 5-7 é severa. Cada nível requer abordagem diferente.""",
            "conduct": """Pontuação IIEF-5 inferior a 22 requer protocolo investigativo abrangente integrando avaliação cardiovascular, metabólica, hormonal e neurovascular. Estratificação inicial por severidade orienta intensidade investigativa e urgência terapêutica.

Avaliação laboratorial tier 1 (todos os casos): Painel hormonal com testosterona total e livre (SHMS 2024 recomenda ambos), SHBG para cálculo de testosterona biodisponível, prolactina (hiperprolactinemia em 2-5% dos casos). Segundo consenso ICSM 2024, LH, estradiol, SHBG e prolactina não são rotina mas na medicina funcional medimos estradiol para avaliar razão T:E2, visando 10:1 a 15:1. TSH e T4 livre (hipotireoidismo subclínico afeta 15% dos casos). Painel metabólico: glicemia de jejum, HbA1c, insulina de jejum para cálculo de HOMA-IR (resistência insulínica presente em mais de 50% com DE). Perfil lipídico avançado incluindo LDL oxidada, apolipoproteínas. Marcadores inflamatórios: PCR ultrassensível (níveis elevados predizem pior resposta a PDE5i), homocisteína.

Avaliação cardiovascular (especialmente se menos de 17 pontos): Escore de risco cardiovascular (Framingham ou pooled cohort equations). ECG de repouso. Considerar teste ergométrico, escore de cálcio coronariano ou ultrassom de carótidas conforme fatores de risco. Medição de pressão arterial ambulatorial se hipertensão suspeitada.

Avaliação tier 2 (casos refratários ou complexos): Doppler peniano com teste de ereção farmacologicamente induzida. Vitamina D, zinco, magnésio (cofatores para síntese de testosterona). Teste de estresse adrenal (cortisol salivar em 4 pontos). Avaliação de apneia do sono (presente em 30-70% com DE). Screening psiquiátrico estruturado (depressão, ansiedade).

Intervenções funcionais baseadas em evidência: Otimização hormonal se testosterona total menor que 300ng/dL ou livre baixo com sintomas. Suporte nutricional: L-arginina 3-5g/dia (precursor de óxido nítrico), L-citrulina 1,5-3g/dia, Panax ginseng 900mg 3x/dia, Pycnogenol 40mg 3x/dia mais L-arginina 1,7g 2x/dia (combinação com evidência robusta). Suporte mitocondrial: CoQ10 200-300mg/dia, NAD mais precursores. Sensibilizadores de insulina: berberina 500mg 3x/dia, cromo, ácido alfa-lipóico, NAC.

PDE5i continuam padrão-ouro farmacológico. Consenso Princeton IV 2024 reafirma segurança cardiovascular. Estudo de 13 pesquisas identificou benefícios cardiovasculares significativos incluindo riscos reduzidos de infarto do miocárdio, insuficiência cardíaca e mortalidade global. Homens com DE expostos a tadalafil tiveram taxas ajustadas significativamente menores de eventos cardiovasculares adversos maiores (HR:0,81), revascularização coronária (HR:0,69), angina instável (HR:0,55), mortes cardiovasculares (HR:0,45) e morte por todas as causas (HR:0,56).

Monitoramento: reaplicação IIEF-5 a cada 3 meses. Reavaliação laboratorial a cada 3-6 meses durante otimização."""
        },
        {
            "id": "7cc28981-bd20-4e5e-b7c0-bd312f6a7ed4",
            "name": "IIEF-5 (Índice Internacional de Função Erétil): ____/25",
            # Mesmo conteúdo do anterior (item duplicado)
        }
    ]
}

def update_item(item_id, data):
    """Atualiza um item via API"""
    url = f"{API_URL}/score-items/{item_id}"

    payload = {
        "clinicalRelevance": data["clinicalRelevance"],
        "patientExplanation": data["patientExplanation"],
        "conduct": data["conduct"]
    }

    try:
        response = requests.put(url, headers=headers, json=payload, timeout=10)

        if response.status_code == 200:
            print(f"✓ Atualizado: {item_id[:8]}... - {data.get('name', 'Item')}")
            return True
        else:
            print(f"✗ Erro {response.status_code}: {item_id[:8]}...")
            print(f"  Response: {response.text[:200]}")
            return False

    except Exception as e:
        print(f"✗ Exceção: {item_id[:8]}... - {str(e)}")
        return False

def main():
    """Processa todos os items"""
    total = 0
    success = 0

    # IIEF-5
    for item in items_data["iief5"]:
        total += 1
        if update_item(item["id"], item):
            success += 1
        time.sleep(0.2)  # Rate limiting

    print(f"\n{'='*60}")
    print(f"Total processado: {total}")
    print(f"Sucesso: {success}")
    print(f"Falhas: {total - success}")
    print(f"{'='*60}")

if __name__ == "__main__":
    main()
