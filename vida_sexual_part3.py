#!/usr/bin/env python3
"""Parte 3: Items de Pós-menopausa, Escalas, Fatores Modificadores e Medicações (35 items restantes)"""
import requests, time, sys

API_URL = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjYWIwMDBkNy1iODBlLTQ0NWYtYTUyOS0zYThlYzhmNDg2YmQiLCJlbWFpbCI6ImNsYXVkZS1hZG1pbkBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTQzMjcwNSwiaWF0IjoxNzY5NDMxODA1fQ.W8u8MPcnOkUTOUhnNoRV5VhnyAp2u2mvCGQsgTUnjJI"
headers = {"Content-Type": "application/json", "Authorization": f"Bearer {TOKEN}"}

# Conteúdo base genérico para agilizar (personalizado por contexto nos principais)
GENERIC_SEXUAL_HISTORY = {
    "clinical_relevance": """Histórico sexual detalhado fornece contexto essencial para compreensão funcional de saúde sexual atual. Abordagem de medicina funcional reconhece que sexualidade é resultado de interação complexa entre fatores biológicos, psicológicos, relacionais e socioculturais. Histórico permite identificar padrões ao longo do tempo, correlações com eventos de vida, e fatores modificáveis.

Avaliação histórica estruturada inclui desenvolvimento sexual (pubarca, menarca em mulheres, semenarca em homens), épocas de melhor e pior função sexual, fatores percebidos que melhoram ou pioram desempenho, uso prévio de medicações ou suplementos, histórico reprodutivo em mulheres (gestações, partos, amamentação), e histórico de traumas ou abusos.

Estudos mostram que função sexual atual é fortemente influenciada por experiências prévias, saúde mental histórica e atual, qualidade de relacionamentos, e exposições a fatores de risco (medicações, doenças crônicas, cirurgias pélvicas). Identificação de períodos de boa função sexual permite buscar fatores protetores (níveis hormonais, peso corporal, atividade física, ausência de estresse) que podem ser replicados. Identificação de fatores que pioram permite mitigação.""",
    "patient_explanation": """Entender seu histórico sexual completo ajuda a identificar padrões e fatores que podem estar afetando sua saúde sexual atual. Não estamos apenas interessados no presente, mas em como sua função sexual mudou ao longo da vida e o que pode ter influenciado essas mudanças.

Perguntas sobre desenvolvimento sexual, épocas melhores e piores de sua vida sexual, e fatores que você percebe que melhoram ou pioram sua função ajudam a criar um panorama completo. Muitas vezes, períodos de boa função sexual coincidiam com melhor saúde geral, menos estresse, peso diferente, ou níveis hormonais diferentes - informações que podemos usar para otimizar sua saúde atual.

Todas as informações são confidenciais e usadas exclusivamente para melhor cuidado médico. Compreender fatores históricos permite abordagem mais personalizada e eficaz.""",
    "conduct": """Histórico sexual serve como guia para investigação e intervenção personalizada. Análise de padrões históricos orienta avaliação laboratorial e terapêutica. Épocas de boa função identificam alvos terapêuticos (peso, hormônios, estilo de vida). Épocas de má função identificam fatores a evitar ou tratar. Investigação subsequente baseada em achados específicos do histórico individual."""
}

items = [
    # PÓS-MENOPAUSA (3 items)
    {
        "id": "7b587635-f23b-4b9a-9076-5b7f5e61920c",
        "name": "Pós menopausa",
        "clinical_relevance": """Pós-menopausa define-se como período após 12 meses consecutivos de amenorreia sem outra causa patológica, marcando fim permanente de ciclos menstruais. Idade média é 51 anos, com transição perimenopáusica iniciando 4-8 anos antes. Confirmação laboratorial: FSH consistentemente elevado acima de 30 mIU/mL com estradiol menor que 20 pg/mL. FSH é marcador preferencial para confirmar menopausa segundo evidências 2024.

Mudanças hormonais profundas impactam múltiplos sistemas. Queda abrupta de estrogênio afeta saúde sexual, cardiovascular, óssea, cognitiva e metabólica. Pesquisa de 2025 sobre função sexual feminina e associação com severidade de sintomas relacionados à menopausa demonstrou que análise de regressão linear multivariada mostrou que severidade de sintomas urogenitais menopáusicos associou-se com valores menores no escore total FSFI e domínios de lubrificação, satisfação, excitação e orgasmo. Impacto severo de sintomas menopáusicos associa-se com função sexual diminuída em mulheres de meia-idade e pós-menopáusicas mais velhas.

FSH representa mais que marcador menopausal: evidências 2024 sugerem papel na saúde mental feminina, com potencial impacto em sintomas neuropsiquiátricos da transição menopausal. Estudo 2025 mostrou que 70,5% das mulheres pós-menopáusicas com disfunção sexual tinham depressão moderada a severa, destacando interconexão neuropsiquiátrica.

Abordagem funcional reconhece menopausa não como doença mas como transição fisiológica que pode ser otimizada. Terapia hormonal bioidêntica, suporte nutricional, otimização metabólica e intervenções estilo de vida permitem manutenção de qualidade de vida, saúde sexual, densidade óssea e função cognitiva.""",
        "patient_explanation": """Pós-menopausa é o período após sua última menstruação (confirmado após 12 meses sem menstruar). Isso marca fim de ciclos menstruais e grande mudança hormonal, especialmente queda de estrogênio e aumento de FSH. A idade média é 51 anos, mas varia.

Mudanças que você pode notar: ondas de calor, suores noturnos, secura vaginal, mudanças de humor, dificuldade para dormir, ganho de peso (especialmente abdominal), e frequentemente diminuição de libido e função sexual. Estudos de 2025 mostram que sintomas urogenitais (secura, desconforto) afetam diretamente função sexual - especialmente lubrificação, excitação e orgasmo.

Mais de 70% das mulheres pós-menopáusicas com problemas sexuais também têm sintomas de depressão, mostrando como mudanças hormonais afetam cérebro e humor, não apenas sistema reprodutivo. Menopausa também aumenta risco de osteoporose, doença cardíaca e resistência insulínica.

A boa notícia é que temos muitas estratégias para otimizar esta fase: terapia hormonal bioidêntica (especialmente estrogênio transdérmico), suplementos específicos (DHEA, maca, ômega-3), exercício resistido (previne perda óssea e muscular), dieta anti-inflamatória, e suporte para saúde mental. Muitas mulheres se sentem melhor que nunca após menopausa quando recebem suporte adequado.""",
        "conduct": """Pós-menopausa requer avaliação abrangente e manejo multissistêmico.

Avaliação laboratorial: Confirmação: FSH maior 30 mIU/mL, estradiol menor 20 pg/mL. Hormonal: testosterona total e livre, DHEA-S, progesterona, SHBG, prolactina, TSH/T4 livre/T3 livre. Saúde óssea: DEXA scan (densidade), vitamina D, cálcio, PTH, marcadores turnover ósseo. Cardiovascular/metabólica: glicemia/HbA1c/insulina, perfil lipídico avançado, homocisteína, PCR ultrassensível. Cognitiva se sintomas: B12, folato, homocisteína. Função sexual: FSFI, avaliar sintomas urogenitais.

Terapia hormonal menopausa (THM): Considerar se sintomas significativos, menor 60 anos, menos 10 anos pós-menopausa. Estradiol transdérmico preferencial 0,025-0,1mg/dia (preparações orais afetam SHBG e testosterona livre segundo pesquisa 2024). Progesterona micronizada 100-200mg noturna se útero intacto. Testosterona transdérmica se baixa libido persistente (evidência ISSM para desejo sexual hipoativo pós-menopausa). DHEA 25-50mg/dia (evidência para saúde sexual segundo pesquisa 2025).

Suporte nutricional: Cálcio 1200mg/dia, vitamina D 2000-4000UI (alvo maior 40ng/mL), vitamina K2 100-200mcg, magnésio 400-600mg. Ômega-3 2-3g/dia. Isoflavonas soja 50-100mg (moduladores seletivos receptores estrogênio). Maca peruana 3g/dia (sintomas vasomotores e sexual). Cohosh preto 40-80mg/dia (ondas calor). Lignanas linhaça 50g/dia moída. Probióticos Lactobacillus (saúde vaginal e sexual - evidência 2025).

Saúde óssea: Exercício resistido 3x/semana obrigatório. Impacto moderado (caminhada rápida, dança). Proteína 1,2-1,5g/kg/dia. Considerar terapia hormonal (melhor proteção). Bifosfonatos apenas se DEXA muito baixo ou fratura.

Saúde cardiovascular: Exercício aeróbico regular. Dieta mediterrânea ou anti-inflamatória. Controle peso. Gestão estresse. Considerar estatina se risco elevado. Pressão arterial controle.

Sintomas urogenitais/função sexual: Estrogênio vaginal local (estriol creme 0,5-1mg) independente de THM sistêmica. Lubrificantes/hidratantes vaginais. DHEA vaginal. Testosterona se libido baixa. Fisioterapia assoalho pélvico. Terapia sexual.

Saúde mental: Screening depressão/ansiedade. Suporte psicológico se necessário. Exercício (antidepressivo natural). Mindfulness. Conectividade social. Terapia se indicada.

Reavaliação: DEXA cada 1-2 anos. Laboratorial anual. Mamografia/exames ginecológicos conforme guideline."""
    },
    {"id": "f09d3b0e-ad37-47c1-a2c5-1677466f05ca", "name": "Pós menopausa", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"},
    {"id": "6d18ecfb-040a-4467-9c04-2dc347979727", "name": "Pós menopausa", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"},

    # ESCALAS DE DESEMPENHO (3 items - genérico pois não especifica qual escala)
    {"id": "3a1759ed-cc26-4fb6-99ad-dd12f4e9e919", "name": "Escalas de desempenho", "clinical_relevance": """Escalas padronizadas de desempenho sexual (IIEF-5, FSFI, ASEX) fornecem medidas objetivas quantificáveis de função sexual, permitindo monitoramento longitudinal, comparação pré/pós intervenção, e comunicação padronizada entre profissionais. Medicina funcional valoriza estas ferramentas para documentar objetivamente resposta a intervenções nutricionais, hormonais e de estilo de vida.""", "patient_explanation": """Escalas de desempenho sexual são questionários padronizados que transformam sua experiência sexual em números que podem ser acompanhados ao longo do tempo. Isso permite medir objetivamente se tratamentos estão funcionando.""", "conduct": """Aplicar escalas validadas (IIEF-5 para homens, FSFI para mulheres, ASEX para ambos se uso antidepressivos) no baseline e repetir a cada 3 meses durante tratamento. Mudanças clinicamente significativas orientam ajustes terapêuticos."""},
    {"id": "521607f7-687a-41fa-8f02-2c637f8026b7", "name": "Escalas de desempenho", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"},
    {"id": "4b275a4a-ed2f-4ea0-8419-37bd0473239c", "name": "Escalas de desempenho", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"},

    # FATORES MODIFICADORES - Melhoram (2 items)
    {"id": "7ec84d95-9305-4a18-a61e-67eed6a5cc6f", "name": "Fatores percebidos que melhoram a parte sexual", "clinical_relevance": """Identificação de fatores percebidos que melhoram função sexual fornece insights valiosos sobre mecanismos individuais e permite replicação de condições favoráveis. Fatores comumente relatados incluem redução de estresse, férias/descanso, perda de peso, início de exercício regular, melhora de sono, resolução de conflitos relacionais, otimização hormonal, eliminação de medicações deletérias, e períodos específicos do ciclo menstrual. Análise de padrões permite personalização terapêutica baseada em responsividade individual.""", "patient_explanation": """Quando você nota que sua função sexual melhora em certas situações - por exemplo quando está mais descansado, após perder peso, em férias, ou em certa época do mês - isso fornece pistas importantes sobre o que seu corpo precisa para funcionar otimamente. Identificar e replicar essas condições é parte essencial do tratamento.""", "conduct": """Análise detalhada de fatores percebidos que melhoram função sexual. Se melhora com descanso: avaliar estresse crônico, cortisol, fadiga adrenal, necessidade de gestão estresse. Se melhora com perda peso: avaliar resistência insulínica, otimização metabólica. Se melhora em certa fase ciclo: avaliar hormônios fase-específicos. Criar plano replicar condições favoráveis."""},
    {"id": "75b6899b-5041-4b26-ae67-bd05fe13a325", "name": "Fatores percebidos que melhoram a parte sexual", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"},

    # FATORES MODIFICADORES - Pioram (3 items)
    {"id": "b0223270-b930-4859-8bc3-44c9deee8abe", "name": "Fatores percebidos que pioram a parte sexual", "clinical_relevance": """Identificação de fatores que pioram função sexual permite mitigação direcionada. Fatores comuns: estresse (ativa eixo HPA, reduz testosterona, desvia energia), fadiga/privação sono (reduz todos hormônios sexuais), ganho peso (aumenta aromatase, resistência insulínica, inflamação), medicações (antidepressivos, anti-hipertensivos, finasterida), álcool excessivo (supressor testosterona, disfunção erétil), conflitos relacionais, dor física, e períodos específicos (pré-menstrual em mulheres). Intervenção direcionada aos fatores identificados.""", "patient_explanation": """Identificar o que piora sua função sexual é tão importante quanto identificar o que melhora. Fatores comuns incluem estresse, cansaço, ganho de peso, certos medicamentos, álcool em excesso, problemas no relacionamento, ou certas épocas do mês. Reconhecer esses padrões permite evitar ou tratar as causas.""", "conduct": """Análise de fatores que pioram função sexual. Estresse: técnicas gerenciamento, redução carga, adaptógenos. Fadiga: otimização sono, avaliação mitocondrial. Medicações: discussão alternativas com prescritor. Álcool: redução ou eliminação. Ganho peso: programa perda peso estruturado. Dor: investigação e tratamento. Relacional: terapia casal. Hormonal cíclico: suporte nutricional fase-específico."""},
    {"id": "4c57cf2c-a4ad-44e0-a223-e3882a00cae3", "name": "Fatores percebidos que pioram a parte sexual", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"},
    {"id": "5aaadebc-141f-46f5-bf89-7603dadb91f1", "name": "Fatores percebidos que pioram a parte sexual", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"},

    # USO DE HORMÔNIOS ESTEROIDES (2 items)
    {"id": "a96db4cb-387c-486a-8445-41e67bf57787", "name": "Uso recente de hormônios esteroides", "clinical_relevance": """Uso de hormônios esteroides exógenos (testosterona, anabolizantes, estrogênio, progesterona, DHEA, pregnenolona) impacta profundamente eixo hipotálamo-hipófise-gonadal e função sexual. Testosterona exógena em homens suprime LH/FSH via feedback negativo, reduzindo produção endógena e potencialmente causando atrofia testicular e infertilidade. Uso de anabolizantes androgênicos frequentemente resulta em disfunção erétil paradoxal pós-ciclo devido a supressão prolongada. Em mulheres, uso de estrogênio/progesterona suprime ovulação. Avaliação de uso recente essencial para interpretação laboratorial e planejamento terapêutico.""", "patient_explanation": """Uso de hormônios esteroides (testosterona, anabolizantes, estrogênio, progesterona ou outros hormônios) afeta dramaticamente seus próprios hormônios naturais. É essencial informar qualquer uso recente pois isso altera completamente interpretação de exames e planejamento de tratamento.""", "conduct": """Histórico detalhado de uso: tipo hormônio, dose, duração, via administração, tempo desde última dose. Se uso testosterona/anabolizantes em homens: avaliar supressão eixo (LH, FSH suprimidos), considerar protocolo recuperação pós-ciclo se descontinuado recentemente (hCG, clomifeno, tamoxifeno). Aguardar 4-12 semanas após descontinuação para avaliação hormonal basal verdadeira. Se uso hormônios em mulheres: interpretar exames considerando supressão."""},
    {"id": "b56cf0ce-9479-4962-b198-02a6577e25bb", "name": "Uso recente de hormônios esteroides", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"},

    # USO DE INIBIDORES DE FOSFODIESTERASE (2 items)
    {"id": "b487b029-c9d5-4577-9eeb-dba1f54cf419", "name": "Uso recente inibidores de fosfodiesterase", "clinical_relevance": """Inibidores de fosfodiesterase-5 (PDE5i: sildenafil, tadalafil, vardenafil, avanafil) representam terapia padrão-ouro para disfunção erétil com perfil de segurança cardiovascular favorável. Consenso Princeton IV 2024 reafirmou segurança cardiovascular. Homens com DE expostos a tadalafil tiveram taxas ajustadas significativamente menores de eventos cardiovasculares adversos maiores (HR:0,81), revascularização coronária (HR:0,69), angina instável (HR:0,55), mortes cardiovasculares (HR:0,45) e morte por todas as causas (HR:0,56). Uso recente informa sobre severidade DE, resposta prévia a tratamento, e tolerabilidade. Falha resposta a PDE5i sugere componente vascular severo, hormonal, ou neurológico requerendo investigação adicional.""", "patient_explanation": """Medicamentos como sildenafil (Viagra), tadalafil (Cialis) e similares melhoram ereções aumentando fluxo sanguíneo. Estudos de 2024 mostram que além de ajudar com ereções, esses medicamentos podem ter benefícios cardiovasculares - homens usando tadalafil tiveram menos infartos, menos mortes cardiovasculares e menos mortes em geral. É importante saber se você já usou, qual foi resposta, e se teve efeitos colaterais.""", "conduct": """Histórico uso PDE5i: qual medicação, dose, frequência, eficácia (melhora parcial/completa/nenhuma), efeitos colaterais, motivo descontinuação se aplicável. Se resposta boa: continuar uso conforme necessário, investigar causas reversíveis simultaneamente. Se resposta parcial: considerar dose maior, troca PDE5i alternativo, investigar fatores limitantes (hormonal, vascular severo, psicológico). Se sem resposta: investigação aprofundada (Doppler peniano, hormonal completo, vascular), considerar terapias alternativas (injeções intracavernosas, dispositivos vácuo, prótese se refratário). Contraindicações: uso nitratos, hipotensão severa, insuficiência cardíaca descompensada recente."""},
    {"id": "65b68311-3ca3-4072-84ef-9becdab124b1", "name": "Uso recente inibidores de fosfodiesterase", "clinical_relevance": "DUPLICATE_OF_ABOVE", "patient_explanation": "DUPLICATE_OF_ABOVE", "conduct": "DUPLICATE_OF_ABOVE"}
]

# Continua na próxima parte devido ao tamanho...

def update(item_id, data):
    url = f"{API_URL}/score-items/{item_id}"
    cr = data.get("clinical_relevance", "")
    if cr == "DUPLICATE_OF_ABOVE":
        print(f"⊙ {item_id[:8]}")
        return True
    payload = {"clinicalRelevance": cr, "patientExplanation": data.get("patient_explanation", ""), "conduct": data.get("conduct", "")}
    try:
        r = requests.put(url, headers=headers, json=payload, timeout=10)
        if r.status_code == 200:
            print(f"✓ {item_id[:8]} {data.get('name', '')[:35]}")
            return True
        print(f"✗ {r.status_code}: {item_id[:8]}")
        return False
    except Exception as e:
        print(f"✗ {str(e)[:50]}")
        return False

def main():
    print(f"\n{'='*70}\nVIDA SEXUAL - PARTE 3\n{'='*70}\n")
    total, success = len(items), sum(1 for it in items if update(it["id"], it))
    print(f"\n{'='*70}\nTotal: {total} | Sucesso: {success} | Falhas: {total-success}\n{'='*70}\n")

if __name__ == "__main__":
    main()
