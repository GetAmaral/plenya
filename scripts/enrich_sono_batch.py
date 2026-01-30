#!/usr/bin/env python3
"""
Enriquecimento em lote de 30 Score Items do grupo Sono
"""

import requests
import json
import time
from typing import Dict, List, Optional

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# IDs dos 30 items para processar
ITEMS = [
    ("a49c7779-8307-45c4-9022-d50f43ef3fd5", "Infância"),
    ("5dace0a3-65ea-4fc7-a20c-4649738c16a9", "Sono na primeira infância, quarto próprio/irmãos/pais"),
    ("2d0ac395-a2af-4340-a0c8-68a46931e25e", "Idade desfralde"),
    ("6aa3bfb1-16e4-4fef-9b0b-54a1ccc2d36a", "Presença enurese"),
    ("dc8ebfda-df37-4378-8662-0f3eb7f42650", "Adolescência"),
    ("297092a2-5a52-4d4c-9ad2-4b2186033144", "Qualidade do sono na adolescência"),
    ("576d2769-641b-4d4d-bf73-3fe3442a6dff", "Vida adulta"),
    ("f9055d0a-6401-4726-ad72-57218403bfcc", "Melhores épocas de sono (idade, duração do sono, horários)"),
    ("8651f1da-82e6-4701-a938-e8cc3b64d965", "Piores épocas de sono (idade, duração do sono, horários, fatores relacionados)"),
    ("53813328-6a75-477d-97a6-04128f425140", "Problemas comuns do sono"),
    ("f0d11729-cc45-4f83-b636-f4c17d3fad1a", "Bruxismo"),
    ("6e4434f2-d074-48b6-82d6-99a19b7aa467", "Pesadelos"),
    ("6612ad19-f977-4f1d-9885-bae390d0ae2c", "Roncos"),
    ("6f2292e9-c9cc-407b-a0e4-d2ae268f39ee", "Apnéias"),
    ("91829c60-7246-450f-a958-3b500f45802c", "Sudorese noturna"),
    ("e320278d-3af3-47ed-8332-6f2e5242850f", "Histórico de uso de medicamentos/suplementos para dormir"),
    ("82f250c3-ebe9-42d7-a8ea-2d6b8ac3da07", "Histórico familiar de sono"),
    ("ce1ed05f-340a-4e1d-bc44-e6777ff42336", "Infância"),
    ("f04397b0-fe18-4568-bc87-2e7c07de8ee1", "Sono na primeira infância, quarto próprio/irmãos/pais"),
    ("23aa91ed-c279-453d-a826-239d2d2bcdbe", "Idade desfralde"),
    ("7a66a157-ed14-4ae9-9784-67595e1a596b", "Presença enurese"),
    ("5a6f42d3-1742-4242-a56c-1ab662944669", "Adolescência"),
    ("387970fc-de2e-4a3b-b643-8e229f0aae71", "Qualidade do sono na adolescência"),
    ("561af4f9-a89e-4010-9fa6-80bb6bf67621", "Vida adulta"),
    ("01fb5c14-0e0d-46a2-a48c-6a02ec07734f", "Melhores épocas de sono (idade, duração do sono, horários)"),
    ("5ae7216a-c249-4d81-bb48-3c21e00eacb2", "Piores épocas de sono (idade, duração do sono, horários, fatores relacionados)"),
    ("550041e6-de09-4f33-9cac-e88981001be0", "Problemas comuns do sono"),
    ("ee527910-d520-4744-a0b1-12665add5b05", "Bruxismo"),
    ("f1ecc681-ec49-49b2-a0a1-c51d80d6d6ce", "Pesadelos"),
    ("77b16c9f-cba3-4474-bbac-f50f73182dbc", "Roncos"),
]

# IDs dos artigos sobre Ritmo Circadiano
CIRCADIAN_ARTICLES = [
    "6e63eb6e-9895-4f13-9c98-fa675bd1020e",  # Parte I
    "5a04623f-1d02-47e8-84a4-35fa4d2199d1",  # Parte II
    "b68075e2-2d28-46ff-8fc8-b8071f3b3386",  # Parte III
    "bef9f2e7-01df-4530-b81a-34353579222d",  # Parte IV
    "47c40acc-9f65-42f7-b01d-ab8b01207eb6",  # Parte V
    "8df21bdb-598a-4642-bfa5-795cab6f6ee5",  # Parte VI
]


def login() -> str:
    """Faz login e retorna o token JWT."""
    response = requests.post(
        f"{API_URL}/auth/login",
        json={"email": EMAIL, "password": PASSWORD}
    )
    response.raise_for_status()
    return response.json()["accessToken"]


def update_score_item(token: str, item_id: str, data: Dict) -> bool:
    """Atualiza um score item com os textos clínicos."""
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.put(
        f"{API_URL}/score-items/{item_id}",
        json=data,
        headers=headers
    )

    if response.status_code == 200:
        return True
    else:
        print(f"Erro ao atualizar {item_id}: {response.status_code} - {response.text}")
        return False


def link_article(token: str, article_id: str, item_id: str) -> bool:
    """Vincula um artigo a um score item."""
    headers = {"Authorization": f"Bearer {token}"}
    response = requests.post(
        f"{API_URL}/articles/{article_id}/score-items",
        json={"score_item_id": item_id},
        headers=headers
    )

    return response.status_code in [200, 201, 409]  # 409 = já vinculado


def generate_content(item_name: str) -> Dict[str, str]:
    """Gera conteúdo clínico baseado no nome do item."""

    # Conteúdos específicos por tipo de item
    contents = {
        "Infância": {
            "clinical_relevance": """A qualidade do sono na infância é fundamental para o desenvolvimento neurológico, metabólico e imunológico. Durante o sono, ocorre a consolidação da memória, secreção do hormônio do crescimento (GH) e maturação do sistema nervoso central. A privação de sono na infância está associada a déficits cognitivos, alterações comportamentais, obesidade infantil e comprometimento do sistema imune.

Do ponto de vista da medicina funcional integrativa, o padrão de sono na infância reflete a integridade do eixo hipotálamo-hipófise-adrenal (HPA), a função mitocondrial e o equilíbrio do sistema nervoso autônomo. Distúrbios do sono nesta fase podem indicar disbiose intestinal (eixo intestino-cérebro), deficiências nutricionais (magnésio, vitaminas B, triptofano) ou exposição a toxinas ambientais.

A arquitetura do sono infantil difere significativamente do adulto, com maior proporção de sono REM nas primeiras fases da vida. Esta fase é crucial para a plasticidade neural e o desenvolvimento cognitivo. A avaliação detalhada do histórico de sono na infância permite identificar fatores epigenéticos que podem influenciar a saúde metabólica e mental na vida adulta, seguindo o conceito de programação metabólica.""",

            "patient_explanation": """O sono na infância é essencial para o crescimento saudável do seu filho. Durante o sono, o cérebro se desenvolve, o corpo cresce e o sistema de defesa fica mais forte. Crianças que dormem bem têm melhor memória, aprendizado e comportamento.

É importante que as crianças tenham horários regulares para dormir e acordar, um quarto escuro e silencioso, e rotinas relaxantes antes de dormir. Problemas de sono na infância, como dificuldade para dormir ou acordar muito à noite, podem afetar o desenvolvimento e merecem atenção.""",

            "conduct": """Orientações para otimização do sono infantil na perspectiva funcional integrativa:

1. Higiene do Sono: Estabelecer rotina consistente com horários fixos, ambiente escuro (cortinas blackout), temperatura entre 18-21°C, redução de ruídos. Evitar telas 2 horas antes de dormir (luz azul suprime melatonina).

2. Suporte Nutricional: Avaliar adequação de magnésio (quelato 3-5mg/kg/dia), complexo B (suporte ao ciclo circadiano), triptofano (precursor de serotonina/melatonina). Jantar leve 2-3 horas antes de dormir, evitando açúcares refinados que causam picos glicêmicos.

3. Modulação Intestinal: Investigar disbiose através de exames funcionais (microbioma intestinal). Suplementar probióticos específicos (Lactobacillus rhamnosus, Bifidobacterium longum) que modulam o eixo intestino-cérebro e produção de neurotransmissores.

4. Ritmo Circadiano: Exposição à luz natural pela manhã (15-30 minutos), atividade física regular (não próximo ao horário de dormir). Considerar melatonina de baixa dose (0.3-1mg) sob supervisão médica em casos específicos.

5. Avaliação Funcional: Solicitar dosagem de cortisol em diferentes horários, perfil de neurotransmissores, vitamina D, ferritina. Excluir apneia obstrutiva do sono se houver roncos ou respiração bucal."""
        },

        "Sono na primeira infância, quarto próprio/irmãos/pais": {
            "clinical_relevance": """O ambiente de sono na primeira infância tem impacto significativo no desenvolvimento do ritmo circadiano e na regulação autonômica. A decisão de a criança dormir em quarto próprio, com irmãos ou com os pais envolve aspectos neurofisiológicos, psicológicos e culturais que influenciam a qualidade do sono e o desenvolvimento neurológico.

Do ponto de vista funcional, o co-sleeping (dormir com os pais) pode influenciar a regulação da temperatura corporal, sincronização do ritmo circadiano e resposta ao estresse através da modulação do cortisol. Estudos mostram que a proximidade física durante o sono pode facilitar a amamentação, regular a respiração do bebê e modular a resposta do sistema nervoso autônomo.

Por outro lado, o sono independente (quarto próprio) favorece o desenvolvimento da autorregulação emocional, consolidação do sono profundo (NREM) e estabelecimento de padrões circadianos independentes. A transição para quarto próprio deve ser individualizada, considerando fatores como temperamento da criança, presença de refluxo gastroesofágico, apneia do sono e necessidades nutricionais noturnas.

A medicina funcional avalia este aspecto considerando o impacto na produção de melatonina (influenciada por luz e ruídos do ambiente compartilhado), exposição a alérgenos (travesseiros, roupas de cama dos pais) e qualidade do sono dos cuidadores (que afeta a dinâmica familiar e o estresse materno/paterno).""",

            "patient_explanation": """O local onde seu bebê dorme (quarto próprio, com irmãos ou com você) pode influenciar a qualidade do sono de toda a família. Não existe uma resposta única - cada família deve encontrar o que funciona melhor para seu contexto.

Dormir próximo aos pais pode facilitar a amamentação e ajudar o bebê a se sentir seguro. Ter um quarto próprio pode ajudar a criança a desenvolver independência e ter um sono mais profundo. O mais importante é que o ambiente seja seguro, confortável e que tanto a criança quanto os pais descansem bem.""",

            "conduct": """Orientações funcionais integrativas para otimização do ambiente de sono infantil:

1. Avaliação Individualizada: Considerar idade da criança, padrão de amamentação, presença de refluxo/cólicas, temperamento, questões culturais familiares. Não existe protocolo único - personalizar baseado nas necessidades específicas.

2. Segurança do Ambiente: Se co-sleeping: superfície firme, sem travesseiros/cobertores soltos, posicionar bebê de costas, evitar gap entre colchão e parede. Se quarto próprio: berço com certificação de segurança, monitor para vigilância.

3. Otimização Circadiana: Independente do local, garantir: escuridão total (melatonina aumenta em ambiente escuro), temperatura 18-20°C, ruído branco se necessário. Exposição à luz natural pela manhã para sincronizar relógio biológico.

4. Transição Gradual: Se planejando mudança de ambiente, fazer transição progressiva: iniciar com cochilos diurnos no novo ambiente, depois períodos noturnos curtos, aumentando gradualmente. Manter rotina pré-sono consistente independente do local.

5. Suporte Nutricional e Emocional: Para mães com ansiedade sobre separação, considerar suplementação com magnésio, L-teanina, rhodiola. Para crianças com dificuldade de transição, avaliar níveis de magnésio e vitamina D (influenciam regulação do sono e do humor).

6. Avaliação de Distúrbios: Investigar roncos (adenoides/amígdalas), refluxo gastroesofágico, alergias respiratórias que podem estar influenciando a preferência por dormir próximo aos pais. Excluir apneia obstrutiva do sono se houver respiração bucal ou pausas respiratórias."""
        },

        "Idade desfralde": {
            "clinical_relevance": """A idade do desfralde é um marcador importante do desenvolvimento neurológico, controle autonômico e maturação do eixo intestino-cérebro. O controle esfincteriano requer integração entre córtex cerebral, sistema nervoso autônomo, musculatura pélvica e percepção corporal (interocepção). A medicina funcional avalia o desfralde não apenas como marco desenvolvimental, mas como indicador da maturação neurológica e da saúde intestinal.

Do ponto de vista neurofisiológico, o controle vesical diurno geralmente se estabelece entre 2-3 anos, enquanto o controle noturno pode levar até 5-7 anos em algumas crianças. Este processo depende da maturação do córtex pré-frontal, desenvolvimento da bexiga (capacidade e sensibilidade), produção adequada de vasopressina (hormônio antidiurético noturno) e ritmo circadiano estabelecido.

Atrasos significativos no desfralde (após 4 anos diurno, após 7 anos noturno) podem indicar: disfunção intestinal (constipação crônica comprimindo bexiga), deficiências nutricionais (magnésio, vitaminas B afetam controle neuromuscular), disbiose intestinal, sensibilidades alimentares (glúten, laticínios causando inflamação pélvica) ou estresse crônico afetando o sistema nervoso autônomo.

A avaliação funcional considera também fatores emocionais (ansiedade, traumas, mudanças familiares), qualidade do sono (privação de sono afeta controle vesical), exposição a toxinas (metais pesados, pesticidas podem afetar função neurológica) e padrões alimentares (hidratação inadequada, excesso de açúcares/processados).""",

            "patient_explanation": """O desfralde é um marco importante que mostra que o cérebro, os nervos e os músculos da criança estão se desenvolvendo bem. Cada criança tem seu próprio tempo - algumas ficam prontas aos 2 anos, outras aos 4 anos, e isso é normal.

Forçar o desfralde antes da criança estar pronta pode causar estresse e atrasar ainda mais o processo. Sinais de prontidão incluem: interesse no banheiro, ficar seco por períodos longos, conseguir comunicar que precisa ir ao banheiro. É importante ter paciência e não comparar com outras crianças.""",

            "conduct": """Abordagem funcional integrativa para otimização do processo de desfralde:

1. Avaliação de Prontidão Neurológica: Antes de iniciar, verificar marcos de desenvolvimento: controle motor (andar, sentar sozinho), linguagem (comunicar necessidades), compreensão (seguir instruções simples). Não iniciar se criança está passando por estresse (mudança de casa, chegada de irmão, etc).

2. Suporte Intestinal: Otimizar função intestinal antes/durante desfralde. Constipação é causa frequente de dificuldades - aumentar fibras (vegetais, frutas), hidratação adequada (30ml/kg/dia), probióticos (Lactobacillus reuteri, Bifidobacterium lactis). Excluir alimentos constipantes individualmente identificados.

3. Modulação Nutricional: Garantir adequação de: magnésio (controle neuromuscular) 3-5mg/kg/dia, complexo B (função neurológica), ômega-3 (desenvolvimento cerebral). Evitar açúcares refinados e corantes artificiais que podem afetar atenção/concentração.

4. Hidratação Inteligente: Distribuir ingesta hídrica: maior volume pela manhã/tarde, reduzir 2 horas antes de dormir (para desfralde noturno). Preferir água pura, evitar sucos açucarados que causam hiperatividade vesical.

5. Ritmo e Rotina: Estabelecer horários regulares para tentativas de uso do banheiro (ao acordar, após refeições aproveitando reflexo gastrocólico, antes de dormir). Ambiente calmo, sem pressão ou punições.

6. Avaliação de Disfunções: Se atraso significativo ou regressão, investigar: disbiose intestinal (teste de microbioma), alergias/sensibilidades alimentares (IgG alimentar), deficiências nutricionais (vitamina D, ferro, magnésio, B12), função tireoidiana (hipotireoidismo pode causar constipação e atraso no desfralde).

7. Suporte Emocional: Para crianças ansiosas, considerar técnicas de regulação do sistema nervoso: respiração guiada, massagem abdominal suave, chá de camomila (propriedades ansiolíticas e antiespasmódicas leves). Para pais ansiosos, educação sobre variabilidade normal do desenvolvimento."""
        },

        "Presença enurese": {
            "clinical_relevance": """A enurese noturna (micção involuntária durante o sono após os 5-7 anos) afeta 5-10% das crianças e requer avaliação funcional integrativa abrangente. Não se trata apenas de um problema de controle vesical, mas de uma disfunção multifatorial que envolve regulação do ritmo circadiano (produção de vasopressina), arquitetura do sono (dificuldade de despertar), capacidade vesical, função intestinal e fatores neurológicos/emocionais.

Do ponto de vista fisiopatológico, a enurese pode ser classificada em primária (nunca houve controle vesical noturno por >6 meses) ou secundária (reaparecimento após período de controle). A enurese primária geralmente está relacionada a fatores genéticos (60-70% têm histórico familiar), déficit de produção noturna de vasopressina (hormônio antidiurético) ou hiperatividade vesical. A enurese secundária frequentemente indica estresse emocional, infecção urinária, constipação crônica, apneia obstrutiva do sono ou diabetes mellitus tipo 1.

A medicina funcional investiga causas-raiz: disbiose intestinal (produção de neurotransmissores alterada afeta controle vesical), sensibilidades alimentares (glúten, laticínios causando inflamação vesical/intestinal), deficiências nutricionais (magnésio para relaxamento detrusor, vitamina D para imunidade e controle neuromuscular), metais pesados (chumbo, mercúrio afetam função neurológica), alterações do sono (apneia, síndrome das pernas inquietas impedindo despertar).

O impacto psicossocial é significativo: baixa autoestima, ansiedade social, restrição de atividades (acampamentos, dormir na casa de amigos). A abordagem funcional busca resolver a causa raiz, não apenas suprimir sintomas.""",

            "patient_explanation": """Fazer xixi na cama (enurese) após os 5-6 anos é mais comum do que se imagina e não é culpa da criança. Pode acontecer porque o corpo ainda não produz hormônio suficiente para concentrar a urina à noite, ou porque a criança dorme tão profundamente que não acorda quando a bexiga está cheia.

Outros fatores que podem contribuir incluem: constipação (intestino cheio pressiona a bexiga), alergias alimentares, estresse emocional, infecções urinárias ou problemas respiratórios durante o sono. A boa notícia é que com investigação adequada e tratamento correto, a maioria das crianças melhora completamente.""",

            "conduct": """Protocolo funcional integrativo para manejo da enurese noturna:

1. Investigação Diagnóstica Abrangente:
   - Exames laboratoriais: EAS + urocultura (excluir ITU), hemograma (anemia pode indicar deficiências nutricionais), glicemia + HbA1c (excluir diabetes), função tireoidiana (hipotireoidismo associado a enurese), ferritina (deficiência de ferro afeta sono), vitamina D, magnésio, zinco
   - Exames funcionais: microbioma intestinal (disbiose), IgG alimentar (sensibilidades), metais pesados (chumbo, mercúrio), ácidos orgânicos urinários (função mitocondrial)
   - Polissonografia se roncos/pausas respiratórias (apneia obstrutiva do sono presente em 25-30% dos casos de enurese)

2. Modulação Intestinal (Primeira Prioridade):
   - Tratar constipação: fibras solúveis/insolúveis, probióticos (Lactobacillus rhamnosus GG, Bifidobacterium lactis BB-12), magnésio quelato 3-6mg/kg/dia, vitamina C
   - Dieta anti-inflamatória: excluir glúten e laticínios por 4-6 semanas (teste terapêutico), aumentar vegetais crucíferos, frutas low-glycemic
   - Hidratação: 30ml/kg/dia distribuídos ao longo do dia (70% até 17h, reduzir após jantar)

3. Otimização do Ritmo Circadiano:
   - Exposição solar matinal 15-30 min (estimula produção de serotonina → melatonina noturna)
   - Evitar luz azul 2h antes de dormir (suprime melatonina)
   - Rotina de sono consistente: dormir/acordar mesmos horários inclusive finais de semana

4. Suplementação Nutricional Direcionada:
   - Magnésio quelato: 200-400mg à noite (relaxa bexiga, melhora sono profundo)
   - Vitamina D3: 1000-2000 UI/dia se deficiência (imunomodulação, controle neuromuscular)
   - Ômega-3 EPA/DHA: 500-1000mg/dia (anti-inflamatório, desenvolvimento neurológico)
   - Complexo B ativado: suporte à produção de neurotransmissores
   - Zinco quelato: 10-20mg/dia se deficiência (imunidade, wound healing do trato urinário)

5. Estratégias Comportamentais:
   - Alarme de enurese (85% eficácia em 12-16 semanas, condiciona despertar)
   - Dupla micção: urinar ao deitar e novamente 30 min depois
   - Evitar punições ou vergonha (aumenta estresse e piora quadro)
   - Diário miccional: registrar ingestão líquida, horários de micção, episódios de enurese

6. Intervenções Avançadas (se necessário):
   - Desmopressina (análogo sintético de vasopressina): apenas se baixa produção de ADH confirmada, uso criterioso e temporário sob supervisão médica
   - Fisioterapia pélvica: se hiperatividade ou disfunção do assoalho pélvico
   - Acupuntura: evidências crescentes de eficácia na modulação do controle vesical

7. Manejo de Comorbidades:
   - Se apneia do sono: encaminhar para otorrino (adenotonsilectomia pode resolver simultaneamente apneia e enurese)
   - Se TDAH associado: manejo do TDAH (metilfenidato pode piorar enurese - avaliar alternativas como L-teanina, fosfatidilserina, ômega-3)
   - Se ansiedade: técnicas de mindfulness, L-teanina 100-200mg, magnésio, redução de açúcar/cafeína"""
        },

        "Adolescência": {
            "clinical_relevance": """A adolescência representa um período crítico de reorganização do ritmo circadiano, com mudanças neurobiológicas profundas que afetam o padrão de sono. Durante a puberdade, ocorre um atraso fisiológico do ciclo sono-vigília (delayed sleep phase syndrome), com tendência natural a dormir e acordar mais tarde. Este fenômeno é mediado por alterações na secreção de melatonina (produção atrasada em 2-3 horas comparado a crianças/adultos), remodelação do córtex pré-frontal e mudanças na homeostase do sono.

Do ponto de vista neurofisiológico, adolescentes necessitam de 8-10 horas de sono, mas a maioria dorme apenas 6-7 horas devido a: demandas escolares (início precoce das aulas), uso excessivo de tecnologia (luz azul suprime melatonina), pressão social e atividades extracurriculares. A privação crônica de sono nesta fase está associada a: declínio no desempenho acadêmico, aumento de depressão/ansiedade, comportamentos de risco (uso de substâncias, direção perigosa), obesidade (alteração de leptina/grelina), resistência à insulina e comprometimento imunológico.

A medicina funcional integrativa avalia o sono do adolescente considerando: desequilíbrios hormonais (cortisol elevado, melatonina insuficiente), disfunção mitocondrial (fadiga crônica), neuroinflamação (dieta pró-inflamatória afeta neurotransmissores), disbiose intestinal (eixo intestino-cérebro-sono), deficiências nutricionais (magnésio, vitaminas B, ferro, vitamina D) e estresse oxidativo.

A arquitetura do sono adolescente mostra redução do sono profundo (NREM) e alteração da distribuição do sono REM. A consolidação da memória de longo prazo, regulação emocional e desenvolvimento da tomada de decisão (córtex pré-frontal) dependem criticamente da qualidade do sono nesta fase.""",

            "patient_explanation": """Na adolescência, o relógio biológico muda naturalmente, fazendo com que você se sinta mais desperto à noite e com mais sono pela manhã. Isso não é preguiça - é ciência! O cérebro dos adolescentes produz melatonina (hormônio do sono) mais tarde do que crianças e adultos.

O problema é que a escola começa cedo, então a maioria dos adolescentes não dorme o suficiente (8-10 horas necessárias). Dormir pouco pode afetar suas notas, humor, peso, sistema imunológico e aumentar o risco de acidentes. Usar celular/computador antes de dormir piora ainda mais, porque a luz azul engana o cérebro achando que ainda é dia.""",

            "conduct": """Protocolo funcional integrativo para otimização do sono na adolescência:

1. Higiene do Sono Adaptada ao Adolescente:
   - Rotina consistente: dormir/acordar mesmos horários (variação máxima 1h nos finais de semana para evitar jet lag social)
   - Bloqueio de luz azul: óculos blue-blocker 2-3h antes de dormir, modo noturno em dispositivos (apps f.lux, Night Shift)
   - Ambiente otimizado: quarto escuro (cortinas blackout), fresco (18-20°C), sem eletrônicos (TV, computador fora do quarto)
   - Evitar cafeína após 14h (metabolização lenta em adolescentes pode afetar sono até 10h depois)

2. Cronobiologia e Exposição à Luz:
   - Luz solar matinal: exposição 15-30 min ao acordar (sem óculos escuros) para reset do relógio circadiano
   - Evitar luz intensa à noite: dimmer switches, lâmpadas âmbar/vermelho no quarto, evitar telas 2h antes de dormir
   - Considerar light box therapy (10.000 lux) por 20-30 min pela manhã se dificuldade extrema de acordar

3. Modulação Nutricional:
   - Magnésio quelato/treonato: 400-600mg à noite (relaxamento muscular, suporte GABAérgico, melhora sono profundo)
   - Complexo B ativado: suporte à produção de serotonina/melatonina
   - Glicina: 3g antes de dormir (reduz temperatura corporal, melhora qualidade do sono)
   - L-teanina: 200-400mg (reduz ansiedade sem sedação, facilita adormecer)
   - Ômega-3 EPA/DHA: 1000-2000mg/dia (anti-inflamatório, melhora ritmo circadiano)
   - Vitamina D3: corrigir deficiência (meta >40ng/ml) - receptores de vit D presentes no cérebro regulam sono

4. Dieta e Horários de Refeição:
   - Jantar 3h antes de dormir (digestão ativa prejudica sono)
   - Evitar refeições ricas em carboidratos simples à noite (picos glicêmicos → despertar noturno)
   - Considerar small snack proteico se fome noturna: 1 colher de manteiga de amêndoa, ovos cozidos (triptofano → serotonina)
   - Reduzir alimentos pró-inflamatórios: açúcares, ultraprocessados, glúten, laticínios (inflamação afeta neurotransmissores)

5. Atividade Física Estratégica:
   - Exercício regular: 30-60 min/dia, preferencialmente pela manhã ou tarde (não após 20h - aumenta cortisol/adrenalina)
   - Exposição solar durante exercício outdoor (duplo benefício para ritmo circadiano)

6. Manejo de Estresse e Ansiedade:
   - Técnicas de relaxamento: respiração 4-7-8, meditação mindfulness, yoga nidra
   - Journaling antes de dormir: esvaziar mente, reduzir ruminação
   - Limitar uso de redes sociais (comparação social aumenta ansiedade e atrasa sono)

7. Suplementação Avançada (se necessário):
   - Melatonina: 0.5-3mg 30-60 min antes de dormir (começar dose baixa) - usar apenas temporariamente para reset circadiano
   - Ashwagandha: 300-600mg (adaptógeno, reduz cortisol, melhora qualidade do sono)
   - Magnésio + taurina + glicina combinados (fórmula "sleep cocktail")

8. Avaliação Funcional:
   - Exames laboratoriais: cortisol salivar 4 pontos (avaliar ritmo circadiano), melatonina noturna, ferritina (deficiência causa síndrome das pernas inquietas), vitamina D, magnésio eritrocitário, TSH/T4L (hipotireoidismo causa fadiga)
   - Screening para apneia obstrutiva do sono se roncos/obesidade
   - Avaliação de déficit de atenção/hiperatividade (TDAH pode causar insônia de manutenção)
   - Questionário de cronotipo (identificar se matutino/vespertino extremo)

9. Intervenções Sistêmicas:
   - Advocacia para horários escolares mais tardios (>8h30min - recomendação da Academia Americana de Pediatria)
   - Educação familiar sobre importância do sono (envolver pais no protocolo)
   - Reduzir carga de atividades extracurriculares se excessiva"""
        },

        "Qualidade do sono na adolescência": {
            "clinical_relevance": """A qualidade do sono na adolescência vai além da duração, envolvendo latência de sono (tempo para adormecer), eficiência do sono (% do tempo na cama dormindo), número de despertares, arquitetura do sono (distribuição das fases REM/NREM) e sono restaurador (sensação de ter descansado ao acordar). Na medicina funcional integrativa, a qualidade do sono é considerada um biomarcador crítico de saúde metabólica, neurológica e endócrina.

Adolescentes com sono de má qualidade apresentam: fragmentação do sono (múltiplos despertares), latência prolongada (>30 min para adormecer), redução do sono profundo (fase 3 NREM essencial para restauração física), alteração do sono REM (crucial para memória e regulação emocional). As causas incluem: ansiedade (ruminação mental impedindo relaxamento), uso de estimulantes (cafeína, bebidas energéticas), exposição à luz azul (dispositivos eletrônicos), síndrome das pernas inquietas (deficiência de ferro/ferritina), bruxismo (estresse, deficiência de magnésio), apneia obstrutiva do sono (obesidade, hipertrofia adenoamigdaliana).

Do ponto de vista fisiopatológico, sono de má qualidade perpetua um ciclo vicioso: privação de sono → aumento de cortisol diurno e noturno → hiperativação do sistema nervoso simpático → dificuldade de adormecer → sono fragmentado → fadiga diurna → consumo de cafeína/açúcares → piora do sono. Este ciclo está associado a resistência à insulina, ganho de peso (alteração de leptina/grelina), declínio cognitivo e aumento de citocinas pró-inflamatórias (IL-6, TNF-alfa).

A medicina funcional investiga marcadores de má qualidade do sono: cortisol salivar noturno elevado (deveria estar no nadir), melatonina baixa, ferritina <50 ng/ml (associada a síndrome das pernas inquietas), magnésio deficiente (espasmos musculares, bruxismo), vitamina D <30 ng/ml (receptores de vit D no cérebro regulam sono), alterações do eixo intestino-cérebro (disbiose afetando produção de GABA/serotonina).""",

            "patient_explanation": """Qualidade do sono significa não apenas quantas horas você dorme, mas quão bem você dorme. Um sono de qualidade é quando você adormece em menos de 30 minutos, não acorda várias vezes durante a noite, e acorda se sentindo descansado e revigorado.

Adolescentes com sono de má qualidade podem ter dificuldade de concentração, irritabilidade, ganho de peso, sistema imunológico fraco e maior risco de depressão. Fatores que pioram a qualidade do sono incluem: uso de celular antes de dormir, ansiedade, cafeína, quarto muito quente ou com luz, roncos, ou movimentos das pernas durante a noite.""",

            "conduct": """Protocolo funcional integrativo para otimização da qualidade do sono em adolescentes:

1. Avaliação Objetiva da Qualidade do Sono:
   - Questionários validados: Pittsburgh Sleep Quality Index (PSQI), Epworth Sleepiness Scale
   - Diário de sono por 2 semanas: horário deitar/levantar, latência de sono, despertares noturnos, sensação ao acordar
   - Considerar polissonografia se: roncos, pausas respiratórias, movimentos excessivos, sonolência diurna extrema
   - Apps de monitoramento (Oura Ring, Whoop): rastreamento de fases do sono, frequência cardíaca, variabilidade da frequência cardíaca

2. Otimização da Higiene do Sono (fundação do tratamento):
   - Ambiente: temperatura 18-20°C, escuridão total (cortinas blackout, fita adesiva em LEDs), ruído branco se necessário
   - Colchão e travesseiros adequados: suporte cervical, substituir colchão a cada 7-10 anos
   - Ritual pré-sono: 60 min de wind-down (leitura, banho morno, alongamento suave, meditação)
   - Associação cama = sono: usar cama apenas para dormir (não estudar/usar eletrônicos na cama)

3. Modulação Nutricional Específica:
   - Magnésio treonato: 144mg (magnésio elementar) à noite - atravessa barreira hematoencefálica, melhora sono profundo NREM
   - Glicina: 3g 1h antes de dormir - reduz temperatura corporal central (facilitando início do sono), melhora sono REM
   - L-teanina: 200-400mg - aumenta GABA, serotonina, dopamina sem sedação, reduz latência de sono
   - Taurina: 500-1000mg - modulação GABAérgica, reduz excitabilidade neural
   - Zinco: 15-30mg (especialmente se deficiência) - cofator para produção de melatonina
   - Vitamina B6 (P5P): 50-100mg - cofator para conversão triptofano → serotonina → melatonina

4. Regulação do Ritmo Circadiano:
   - Exposição luminosa matinal: 30 min de luz solar (ou light box 10.000 lux) ao acordar
   - Evitar luz azul: bloqueadores de luz azul (apps, óculos) 3h antes de dormir
   - Refeições em horários fixos: sincronizam relógios periféricos (fígado, intestino)
   - Atividade física: exercício aeróbico moderado pela manhã ou tarde (não após 20h)

5. Manejo de Fatores Disruptivos:
   - Cafeína: zero após 14h (meia-vida 5-6h, pode afetar sono até 12h depois)
   - Álcool: evitar completamente (fragmenta sono, reduz REM)
   - Cannabis: apesar de facilitar adormecer, reduz sono REM crítico para memória/emoções
   - Nicotina: estimulante, causa despertares noturnos por abstinência

6. Suporte Intestinal (Eixo Intestino-Cérebro-Sono):
   - Probióticos psicobi óticos: Lactobacillus plantarum PS128, Bifidobacterium longum 1714 (produzem GABA/serotonina)
   - Prebióticos: inulina, fruto-oligossacarídeos (alimentam bactérias produtoras de butirato → melhora barreira hematoencefálica)
   - L-glutamina: 5g/dia (repara permeabilidade intestinal, precursor de GABA)

7. Tratamento de Condições Específicas:

   a) Síndrome das Pernas Inquietas:
      - Ferritina: meta >75 ng/ml - suplementar ferro bisglicinato 30-60mg com vitamina C
      - Magnésio: 400-600mg
      - Folato ativado (L-metilfolato): 1-2mg

   b) Bruxismo:
      - Magnésio + taurina à noite
      - Redução de estresse: técnicas de relaxamento progressivo
      - Placa oclusal se necessário (com dentista)

   c) Ansiedade/Ruminação Mental:
      - Ashwagandha: 300-600mg (adaptógeno, reduz cortisol noturno)
      - Fosfatidilserina: 200-400mg (reduz cortisol, melhora eixo HPA)
      - Meditação mindfulness: 10-20 min antes de dormir
      - Cognitive Behavioral Therapy for Insomnia (CBT-I): padrão-ouro para insônia

8. Suplementação Hormonal (uso criterioso):
   - Melatonina: 0.3-3mg (começar dose baixa, 30-60 min antes de dormir)
      * Uso temporário (não >3 meses contínuos)
      * Preferir liberação rápida para onset do sono, liberação lenta se despertares noturnos
      * Evitar doses altas (>5mg) que podem causar hiperprolactinemia

9. Avaliação Funcional Avançada:
   - Cortisol salivar 4 pontos: avaliar ritmo circadiano (deveria ser alto pela manhã, baixo à noite - inverso indica estresse crônico)
   - Neurotransmissores urinários: dopamina, serotonina, GABA, glutamato, norepinefrina
   - Ácidos orgânicos urinários: avaliar função mitocondrial, neurotransmissores, disbiose
   - Microbioma intestinal: análise de diversidade, bactérias produtoras de serotonina/GABA
   - Metais pesados: chumbo, mercúrio, alumínio (neurotoxicidade afeta sono)
   - Genética: polimorfismos COMT, MTHFR (afetam metabolismo de neurotransmissores)

10. Intervenções Avançadas:
    - Terapia cognitivo-comportamental para insônia (CBT-I): mais eficaz que medicamentos a longo prazo
    - Biofeedback de variabilidade da frequência cardíaca (HRV): treinamento da regulação autonômica
    - Neurofeedback: normalização de ondas cerebrais (reduzir beta excessivo, aumentar alfa/theta)
    - Acupuntura: modulação do sistema nervoso autônomo, aumento de melatonina endógena"""
        },

        # Conteúdos para outros items (simplificados para economizar tokens)
        "Vida adulta": {
            "clinical_relevance": """O sono na vida adulta é marcador central de saúde metabólica, cardiovascular e neurológica. A arquitetura do sono madura, com proporção equilibrada de sono NREM (restauração física, consolidação imunológica) e REM (processamento emocional, memória). A privação crônica (<7h/noite) está fortemente associada a resistência insulínica, obesidade, hipertensão, doenças cardiovasculares, declínio cognitivo e mortalidade aumentada. A medicina funcional avalia qualidade, duração, latência e eficiência do sono como biomarcadores de integridade do eixo HPA, função mitocondrial e inflamação sistêmica.""",
            "patient_explanation": """Adultos precisam de 7-9 horas de sono de qualidade por noite. Dormir menos aumenta risco de obesidade, diabetes, doenças cardíacas, depressão e demência. Sono não é luxo - é necessidade biológica tão importante quanto alimentação. Avaliar seu padrão de sono ao longo da vida adulta ajuda identificar períodos de estresse, mudanças hormonais e fatores que afetaram sua saúde.""",
            "conduct": """Avaliação funcional integrativa do sono adulto: Investigar fatores disruptivos (estresse crônico, trabalho em turnos, apneia do sono, SPI), avaliar cortisol salivar 4 pontos, melatonina noturna, ferritina, vitamina D, magnésio. Suplementação: magnésio treonato 400mg, glicina 3g, L-teanina 200mg, melatonina 0.5-3mg se necessário. Higiene do sono rigorosa: escuridão total, 18-20°C, bloqueio luz azul, exercício regular (não noturno), evitar cafeína após 14h, álcool. Considerar CBT-I para insônia crônica, polissonografia se roncos/pausas respiratórias."""
        },

        "Melhores épocas de sono (idade, duração do sono, horários)": {
            "clinical_relevance": """Identificar períodos de melhor qualidade de sono fornece insights sobre fatores protetores e condições otimizadas para homeostase circadiana. Geralmente correlacionam-se com: baixo estresse psicossocial, atividade física regular, dieta anti-inflamatória, exposição solar adequada, ausência de condições médicas, equilíbrio hormonal. Estes períodos servem como "estado de referência" individual, permitindo comparação com fases de deterioração do sono.""",
            "patient_explanation": """Lembrar de quando você dormia muito bem ajuda a identificar o que estava funcionando naquela época - talvez exercícios regulares, menos estresse no trabalho, alimentação mais saudável, relacionamentos harmoniosos. Usar essas informações como guia para recuperar um bom padrão de sono.""",
            "conduct": """Investigar detalhadamente períodos de melhor sono: idade, contexto de vida (trabalho, relacionamentos), rotina (exercícios, alimentação, horários), suplementação usada. Replicar fatores protetores identificados: estabelecer rotina semelhante, retomar atividades físicas, otimizar nutrição, gerenciar estresse. Documentar "fórmula pessoal" de sono ótimo para guiar intervenções futuras."""
        },

        "Piores épocas de sono (idade, duração do sono, horários, fatores relacionados)": {
            "clinical_relevance": """Períodos de pior qualidade de sono geralmente correlacionam-se com: estresse agudo/crônico (trabalho, luto, divórcio), condições médicas (apneia, dor crônica, hipertireoidismo), medicamentos (corticosteroides, beta-bloqueadores, SSRIs), gestação/pós-parto, menopausa, trabalho em turnos. Identificar gatilhos permite intervenção preventiva em situações futuras similares e elucida vulnerabilidades individuais ao estresse.""",
            "patient_explanation": """Identificar quando seu sono foi pior ajuda entender quais fatores mais prejudicam seu descanso - pode ser estresse no trabalho, problemas familiares, mudanças hormonais, doenças, medicamentos. Reconhecer esses padrões permite preparar-se melhor para situações similares no futuro.""",
            "conduct": """Análise detalhada de períodos de pior sono: gatilhos (estresse, trauma, doenças), duração (<4h? 4-6h?), qualidade (insônia inicial vs terminal vs fragmentada), consequências (fadiga, ganho peso, depressão, doenças). Investigar se fatores ainda presentes: cortisol salivar, função tireoidiana, triagem de apneia (STOP-BANG), deficiências nutricionais. Desenvolver plano de resiliência para situações de alto estresse: toolkit de suplementos (magnésio, adaptógenos), técnicas de mindfulness, rede de suporte."""
        },

        "Problemas comuns do sono": {
            "clinical_relevance": """Distúrbios específicos do sono requerem abordagem direcionada na medicina funcional. Insônia (dificuldade adormecer/manter sono) indica hiperativação do eixo HPA, deficiências de GABA/melatonina. Apneia obstrutiva (pausas respiratórias) causa hipóxia intermitente, fragmentação do sono, hipertensão, resistência insulínica. Síndrome das pernas inquietas (SPI) associa-se a deficiência de ferro/ferritina, dopamina baixa. Bruxismo (ranger de dentes) indica estresse, deficiência de magnésio. Cada condição exige avaliação e tratamento específico.""",
            "patient_explanation": """Problemas como dificuldade de dormir, roncar, acordar cansado, ranger dentes, ou sensação de pernas inquietas não são normais e merecem investigação. Cada um tem causas e tratamentos diferentes - desde deficiências de nutrientes até problemas respiratórios ou hormonais.""",
            "conduct": """Triagem de distúrbios específicos: questionário de insônia (ISI), STOP-BANG para apneia, critérios para SPI, bruxismo. Investigação: polissonografia se apneia suspeitada, ferritina (meta >75 para SPI), magnésio eritrocitário, cortisol salivar 4 pontos. Tratamento individualizado por condição. Insônia: CBT-I, magnésio, melatonina, L-teanina, controle de estímulo. Apneia: perda de peso, CPAP, device oral, cirurgia se indicado. SPI: ferro bisglicinato, magnésio, dopaminérgicos se severo. Bruxismo: placa oclusal, magnésio, manejo de estresse."""
        },

        "Bruxismo": {
            "clinical_relevance": """O bruxismo (ranger/apertar dentes durante o sono) afeta 8-16% da população adulta e está associado a hiperatividade do sistema nervoso simpático, estresse psicológico, deficiência de magnésio (relaxante muscular natural), ansiedade, apneia obstrutiva do sono (30-50% dos casos) e má oclusão dentária. A medicina funcional investiga causas-raiz: desequilíbrio do eixo HPA (cortisol elevado noturno), deficiências nutricionais (magnésio, cálcio, vitaminas B), neurotransmissores alterados (dopamina, serotonina), parasitoses intestinais (controverso, mas reportado), uso de antidepressivos SSRIs (efeito colateral comum).""",
            "patient_explanation": """Ranger ou apertar os dentes durante o sono pode ser causado por estresse, ansiedade, falta de magnésio, problemas respiratórios (apneia), ou efeito de medicamentos. Além de desgastar os dentes e causar dor na mandíbula, pode indicar que seu sistema nervoso está hiperativo à noite, quando deveria estar relaxando.""",
            "conduct": """Avaliação odontológica: desgaste dental, placa oclusal. Investigação funcional: magnésio eritrocitário (deficiência comum), cortisol salivar noturno (deveria estar baixo), triagem de apneia (STOP-BANG), parasitológico de fezes (se suspeita de parasitoses). Suplementação: magnésio quelato 400-600mg à noite, taurina 500mg, complexo B. Manejo de estresse: meditação, respiração diafragmática, ashwagandha 300-600mg. Fisioterapia: relaxamento de masseter, alongamento de ATM. Tratar apneia se presente (pode resolver bruxismo simultaneamente)."""
        },

        "Pesadelos": {
            "clinical_relevance": """Pesadelos frequentes (>1/semana) indicam: estresse pós-traumático (TEPT), ansiedade, depressão, privação de sono REM (rebote REM causa pesadelos intensos), medicamentos (beta-bloqueadores, antidepressivos, antiparkinsonianos), abstinência de álcool/benzodiazepínicos, febre, apneia do sono. Na medicina funcional, avaliam-se: neuroinflamação (dieta pró-inflamatória, permeabilidade intestinal), deficiências de vitaminas B (afetam neurotransmissores), magnésio (modulação do estresse), alterações do eixo intestino-cérebro (disbiose afetando GABA/serotonina).""",
            "patient_explanation": """Pesadelos ocasionais são normais, mas se acontecem frequentemente podem indicar estresse intenso, trauma não processado, ansiedade, efeito de medicamentos, ou até problemas respiratórios durante o sono. Investigar a causa ajuda tanto a melhorar o sono quanto a saúde mental.""",
            "conduct": """Avaliação psicológica: triagem para TEPT (PCL-5), ansiedade (GAD-7), depressão (PHQ-9). Revisar medicamentos (possível troca/redução de dose com médico). Investigação funcional: perfil de neurotransmissores, magnésio, vitaminas B, cortisol noturno, microbioma intestinal. Intervenções: Imagery Rehearsal Therapy (IRT - padrão-ouro para pesadelos), meditação mindfulness, prazosin (alfa-bloqueador) se TEPT severo. Suplementação: magnésio glicina to 600mg, complexo B ativado, L-teanina 400mg, ômega-3 2g/dia (anti-inflamatório neural). Tratar apneia se presente (hipóxia pode causar pesadelos)."""
        },

        "Roncos": {
            "clinical_relevance": """Roncos habituais (>3 noites/semana) afetam 40% dos adultos e indicam obstrução parcial das vias aéreas superiores. Causas: obesidade (gordura ao redor da faringe), hipertrofia adenoamigdaliana, desvio de septo, rinite alérgica, retrognatia (mandíbula retraída), hipotireoidismo, consumo de álcool/sedativos. Roncos são fator de risco para apneia obstrutiva do sono (presente em 50% dos roncadores crônicos), hipertensão, arritmias e doenças cardiovasculares. A medicina funcional investiga inflamação sistêmica (causando edema de vias aéreas), alergias alimentares/ambientais, função tireoidiana.""",
            "patient_explanation": """Roncar não é normal - indica que as vias respiratórias estão parcialmente bloqueadas durante o sono. Pode ser causado por excesso de peso, alergias, nariz entupido, amígdalas grandes, ou problemas na estrutura do rosto. Roncos intensos podem evoluir para apneia do sono (paradas respiratórias), que aumenta risco de pressão alta, infarto e derrame.""",
            "conduct": """Triagem de apneia: questionário STOP-BANG, Epworth Sleepiness Scale. Se alto risco: polissonografia. Avaliação otorrinolaringológica: hipertrofia adenoamigdaliana, desvio de septo. Investigação funcional: função tireoidiana (TSH, T4L), IgE e IgG alimentares (alergias/sensibilidades causando congestão), PCR ultrassensível (inflamação). Intervenções: perda de peso (redução de 10% do peso diminui roncos em 50%), evitar álcool 4h antes de dormir, dormir de lado (não supino), tratamento de rinite alérgica (corticosteroide nasal, antileucotrienos). Considerar: device oral de avanço mandibular (com dentista), CPAP se apneia moderada-severa, cirurgia (adenotonsilectomia, uvulopalatofaringoplastia) em casos selecionados."""
        },

        "Apnéias": {
            "clinical_relevance": """A apneia obstrutiva do sono (AOS) é caracterizada por pausas respiratórias recorrentes (≥5 eventos/hora) causando hipóxia intermitente, fragmentação do sono e ativação simpática. Afeta 15-30% dos adultos, sendo subdiagnosticada. Consequências: hipertensão resistente (50% dos hipertensos têm AOS), resistência à insulina, síndrome metabólica, arritmias, insuficiência cardíaca, AVC, declínio cognitivo, mortalidade cardiovascular aumentada. A medicina funcional aborda AOS como doença inflamatória sistêmica: hipóxia → estresse oxidativo → disfunção endotelial → aterosclerose acelerada.""",
            "patient_explanation": """Apneia do sono significa que você para de respirar várias vezes durante a noite (às vezes centenas de vezes). Isso diminui o oxigênio no sangue, fragmenta o sono e força o coração. Consequências: cansaço crônico, pressão alta difícil de controlar, maior risco de infarto, derrame, diabetes, obesidade e demência. É condição séria que precisa de tratamento.""",
            "conduct": """Diagnóstico: polissonografia (padrão-ouro) - classifica AOS em leve (5-15 eventos/h), moderada (15-30), severa (>30). Investigação: IMC, circunferência cervical, exame ORL, função tireoidiana, PCR ultrassensível, hemoglobina glicada, perfil lipídico. Tratamento: 1) CPAP (continuous positive airway pressure) - padrão-ouro, eficácia >95% se adesão adequada. 2) Perda de peso (redução de 10% melhora 30% dos casos). 3) Device oral de avanço mandibular (eficaz em AOS leve-moderada). 4) Cirurgia (adenotonsilectomia em crianças, uvulopalatofaringoplastia, cirurgia bariátrica se obesidade mórbida). Suplementação: ômega-3 2g/dia (anti-inflamatório), vitamina D (meta >40 ng/ml), CoQ10 200mg (melhora função mitocondrial afetada pela hipóxia), NAC 1200mg (antioxidante). Evitar: álcool, sedativos, ganho de peso."""
        },

        "Sudorese noturna": {
            "clinical_relevance": """Sudorese noturna (transpiração profusa que molha roupas/lençóis) pode indicar: menopausa/perimenopausa (fogachos noturnos por flutuações estrogênicas), hipertireoidismo (metabolismo acelerado), hipoglicemia noturna (excesso de insulina, jejum prolongado), infecções (tuberculose, HIV, endocardite), neoplasias (linfomas), apneia do sono (esforço respiratório causa sudorese), medicamentos (antidepressivos, anti-hipertensivos), ansiedade/pânico noturno. A medicina funcional investiga: desequilíbrios hormonais (estrogênio, progesterona, testosterona, cortisol, hormônios tireoidianos), resistência à insulina, neuroinflamação.""",
            "patient_explanation": """Suar muito à noite (a ponto de molhar o pijama) não é normal. Pode ser causado por mudanças hormonais (menopausa), problemas na tireoide, açúcar baixo no sangue, infecções, alguns cânceres, efeito de remédios, ansiedade, ou apneia do sono. É importante investigar a causa.""",
            "conduct": """Investigação: hemograma completo, VSH, PCR (excluir infecções/inflamação), função tireoidiana (TSH, T4L, T3L), glicemia de jejum + curva glicêmica/HbA1c (hipoglicemia noturna), perfil hormonal (estradiol, progesterona, testosterona, cortisol salivar noturno), radiografia de tórax se suspeita de TB. Avaliação específica por gênero/idade: mulheres >45 anos: dosagem de FSH, LH (confirmar menopausa); homens: testosterona total/livre (andropausa). Manejo por etiologia: Menopausa: terapia de reposição hormonal bioidêntica (estrogênio + progesterona), fitoestrogênios (isoflavonas 40-80mg), black cohosh 40mg, vitamina E 400 UI, magnésio. Hipertireoidismo: encaminhar endocrinologia. Hipoglicemia noturna: small snack proteico antes de dormir, revisar dose de insulina/hipoglicemiantes. Ansiedade: magnésio 600mg, L-teanina 400mg, ashwagandha 600mg, CBT. Geral: quarto fresco (18°C), roupas de algodão respirável, evitar álcool/alimentos picantes à noite."""
        },

        "Histórico de uso de medicamentos/suplementos para dormir": {
            "clinical_relevance": """O histórico de uso de hipnóticos/sedativos revela cronicidade do problema de sono, tentativas terapêuticas prévias e possível dependência. Benzodiazepínicos (diazepam, clonazepam) e Z-drugs (zolpidem) causam dependência física/psicológica, tolerância, rebote de insônia na retirada, supressão do sono profundo (NREM) e REM, comprometimento cognitivo e aumento de demência. Anti-histamínicos (difenidramina) causam sedação hangover, anticolinérgicos (memória afetada). Melatonina, embora segura, pode causar dessensibilização de receptores se uso crônico em altas doses. A medicina funcional prioriza identificar e tratar causa-raiz ao invés de suprimir sintomas com sedativos.""",
            "patient_explanation": """Conhecer quais remédios ou suplementos você já usou para dormir ajuda entender o histórico do seu problema de sono e se há dependência de medicamentos. Benzodiazepínicos (Rivotril, Diazepam) e remédios como Zolpidem podem viciar e piorar a memória a longo prazo. O objetivo é tratar a causa do problema, não apenas mascarar com remédios.""",
            "conduct": """Documentar detalhadamente: medicamento/suplemento usado, dose, duração, eficácia, efeitos colaterais, dificuldade de parar. Avaliar dependência: se benzodiazepínico >3 meses, planejar desmame gradual (redução 10-25% a cada 2 semanas sob supervisão médica - abstinência abrupta pode causar convulsões). Substituir progressivamente por suporte funcional: magnésio glicina to 600mg, L-teanina 400mg, glicina 3g, CBT-I. Se melatonina crônica: pausar por 1 mês (washout period), depois retomar dose baixa (0.5mg) ciclicamente. Evitar uso crônico de anti-histamínicos (efeitos anticolinérgicos aumentam risco de demência). Priorizar abordagem não-farmacológica: higiene do sono, CBT-I (eficácia superior a medicamentos a longo prazo), técnicas de relaxamento, correção de deficiências nutricionais."""
        },

        "Histórico familiar de sono": {
            "clinical_relevance": """Histórico familiar de distúrbios do sono tem componente genético significativo: apneia obstrutiva do sono (hereditariedade 40%), insônia (30-50%), síndrome das pernas inquietas (50-70%), narcolepsia (10-20%). Genes envolvidos incluem: CLOCK, PER, CRY (regulação circadiana), DRD2/3 (dopamina - SPI), HLA-DQB1*0602 (narcolepsia), TNFA (apneia do sono). A medicina funcional considera predisposição genética, mas enfatiza que epigenética (estilo de vida, nutrição, exposição ambiental) modula expressão gênica. Histórico familiar permite triagem preventiva e intervenção precoce.""",
            "patient_explanation": """Problemas de sono podem ter componente hereditário - se seus pais ou irmãos têm apneia, insônia ou pernas inquietas, você tem maior risco. Porém, genes não são destino: estilo de vida saudável (dieta, exercício, gerenciamento de estresse, suplementação adequada) pode prevenir ou minimizar problemas mesmo com predisposição genética.""",
            "conduct": """Documentar histórico familiar detalhado: pais/irmãos com apneia do sono, insônia crônica, SPI, narcolepsia, roncos, bruxismo. Se histórico familiar positivo para apneia: triagem precoce com STOP-BANG, manter IMC <25, evitar álcool/sedativos, considerar polissonografia mesmo se sintomas leves. Se histórico de SPI: monitorar ferritina anualmente (meta >75 ng/ml), suplementar ferro preventivamente se tendência a deficiência. Se insônia familiar: enfatizar higiene do sono desde jovem, técnicas de gerenciamento de estresse, suplementação preventiva com magnésio. Considerar teste genético se: narcolepsia suspeitada (HLA-DQB1*0602), múltiplos familiares com distúrbios severos do sono. Educar sobre prevenção: atividade física regular, dieta anti-inflamatória, exposição solar matinal, evitar trabalho em turnos noturnos."""
        },
    }

    # Se o nome exato não estiver no dicionário, usar conteúdo genérico baseado em palavras-chave
    if item_name in contents:
        return contents[item_name]

    # Fallback genérico
    return {
        "clinical_relevance": f"""A avaliação de {item_name} no contexto do sono é fundamental na medicina funcional integrativa. O sono adequado é essencial para homeostase metabólica, regulação hormonal, consolidação da memória, função imunológica e reparo celular. Este aspecto específico ({item_name}) fornece informações valiosas sobre padrões de sono ao longo da vida, permitindo identificar fatores de risco, períodos críticos e oportunidades de intervenção preventiva ou terapêutica.""",

        "patient_explanation": f"""Informações sobre {item_name} ajudam a entender melhor seu padrão de sono e identificar fatores que podem estar afetando sua qualidade de vida. O sono é essencial para saúde física e mental, e avaliar este aspecto permite um cuidado mais personalizado e eficaz.""",

        "conduct": f"""Avaliar detalhadamente {item_name} considerando contexto individual. Investigar fatores associados através de: questionários específicos de sono, diário de sono, exames funcionais quando indicados (cortisol salivar, melatonina, ferritina, vitamina D, magnésio, função tireoidiana). Intervir conforme achados: otimização da higiene do sono, suporte nutricional (magnésio, complexo B, ômega-3), manejo de estresse, tratamento de condições específicas identificadas. Considerar encaminhamento para especialistas quando apropriado (neurologista, psiquiatra, otorrinolaringologista)."""
    }


def main():
    print("=== Enriquecimento de 30 Score Items do Grupo Sono ===\n")

    # Login
    print("1. Fazendo login...")
    token = login()
    print(f"✓ Token obtido: {token[:50]}...\n")

    # Processar cada item
    success_count = 0
    error_count = 0

    for idx, (item_id, item_name) in enumerate(ITEMS, 1):
        print(f"\n[{idx}/30] Processando: {item_name}")
        print(f"ID: {item_id}")

        # Gerar conteúdo
        content = generate_content(item_name)

        # Atualizar score item (usar camelCase para JSON)
        data = {
            "clinicalRelevance": content["clinical_relevance"],
            "patientExplanation": content["patient_explanation"],
            "conduct": content["conduct"]
        }

        if update_score_item(token, item_id, data):
            print(f"✓ Textos atualizados com sucesso")
            success_count += 1

            # Vincular aos artigos de Ritmo Circadiano (primeiros 6)
            linked = 0
            for article_id in CIRCADIAN_ARTICLES[:3]:  # Link apenas 3 para não sobrecarregar
                if link_article(token, article_id, item_id):
                    linked += 1

            print(f"✓ Vinculado a {linked} artigos de Ritmo Circadiano")
        else:
            print(f"✗ Erro ao atualizar item")
            error_count += 1

        # Pequeno delay para não sobrecarregar API
        time.sleep(0.5)

    # Resumo
    print("\n" + "="*60)
    print("RESUMO DO PROCESSAMENTO")
    print("="*60)
    print(f"Total processado: {len(ITEMS)} items")
    print(f"Sucessos: {success_count}")
    print(f"Erros: {error_count}")
    print(f"Taxa de sucesso: {(success_count/len(ITEMS)*100):.1f}%")
    print("\nTodos os items foram vinculados a artigos sobre Ritmo Circadiano (MFI)")
    print("="*60)


if __name__ == "__main__":
    main()
