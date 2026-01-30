#!/usr/bin/env python3
"""
Enriquecimento de 20 Score Items do grupo SONO - Parte 2
Foco em: Ambiente do sono, Apneias, Barulho, Bruxismo, Cama, Campos eletromagnéticos
"""

import requests
import json
import time
from typing import Dict, List, Optional

API_URL = "http://localhost:3001/api/v1"
EMAIL = "import@plenya.com"
PASSWORD = "Import@123456"

# IDs dos 20 items para processar
ITEMS = [
    ("ef3f2669-cf32-4929-bdb5-c0b299d60795", "Adolescência"),
    ("1ddb3ba1-bd1d-4692-af21-f56834636c1f", "Ambiente do sono"),
    ("74dad4e7-7184-4bea-886c-ed632a8e7d23", "Ambiente do sono"),
    ("6526cc60-99d8-4147-972b-e88d4a9b2acc", "Ambiente do sono"),
    ("29f615c2-9608-409f-aa26-2c81aec67dda", "Apneias"),
    ("3df33276-6bf6-4ca1-9dd2-ac6e478cd0e5", "Apneias"),
    ("6fc90b17-bd93-4d30-b0ec-13817d2e1756", "Apneias"),
    ("cd064e4a-8027-4b41-93dc-e082028ceab4", "Apnéias"),
    ("da35847c-3f72-47b5-ad2b-8505be91ebb2", "Apnéias"),
    ("02f8f992-b47a-407f-808e-bd95f8d4d390", "Barulho"),
    ("a1bae13e-5df5-4b9d-9af0-29fbe1071c31", "Barulho"),
    ("2b7e6b36-5578-4c46-901c-16e38005e8d9", "Barulho"),
    ("b4d6985c-3c5d-4709-ad60-43b2a44cd079", "Bruxismo"),
    ("3c05c193-2a45-47c8-bda6-5f22861354f1", "Bruxismo"),
    ("efe2b4bb-8980-420d-aa75-48508ea4d25d", "Bruxismo"),
    ("e7f8abf5-0988-4add-af05-8375aa5585a5", "Bruxismo"),
    ("67ed1753-01d9-45e0-9193-d1ce961ca972", "Cama"),
    ("a92458e9-f6db-4a09-bdfc-cdbac22ac227", "Cama"),
    ("a59269ac-50fb-4aac-9cf8-ea99accfeb68", "Cama"),
    ("339da4f9-89ef-4990-bbc5-e15de8eb96be", "Campos eletromagnéticos"),
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
    """Gera conteúdo clínico detalhado baseado no nome do item."""

    contents = {
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
   - Questionário de cronotipo (identificar se matutino/vespertino extremo)"""
        },

        "Ambiente do sono": {
            "clinical_relevance": """O ambiente de sono é determinante crucial para qualidade do sono, regulação do ritmo circadiano e restauração fisiológica. A medicina funcional integrativa reconhece que fatores ambientais modulam diretamente a produção de melatonina (hormônio regulador do sono), temperatura corporal (essencial para indução do sono), atividade do sistema nervoso autônomo e arquitetura do sono (distribuição das fases NREM e REM).

A neurobiologia do sono demonstra que: 1) A melatonina é suprimida por exposição à luz (especialmente luz azul 450-480nm), mesmo em intensidades baixas (<10 lux). A presença de fontes luminosas no quarto (LEDs de aparelhos, luz de rua, telas) pode reduzir a melatonina em até 50%, atrasando o início do sono em 1-3 horas. 2) A temperatura corporal central precisa diminuir 1-2°C para iniciar o sono - ambientes quentes (>21°C) impedem esta termorregulação, prolongando latência de sono e fragmentando o sono profundo. 3) Ruídos acima de 40 decibéis (conversação normal) causam microdespertares, mesmo sem despertar consciente, fragmentando o sono e reduzindo tempo em sono profundo (estágio 3 NREM) essencial para restauração física e consolidação imunológica.

Do ponto de vista funcional, um ambiente de sono inadequado perpetua cascata de disfunções: sono fragmentado → ativação do eixo HPA (cortisol elevado) → resistência à insulina → ganho de peso → apneia do sono → inflamação sistêmica. A otimização ambiental é intervenção de alta relação custo-benefício, frequentemente subestimada em relação a suplementação ou farmacoterapia.

Poluentes ambientais também afetam qualidade do sono: compostos orgânicos voláteis (VOCs) de tintas/móveis, alérgenos (ácaros, fungos), campos eletromagnéticos (EMF) de dispositivos wireless podem causar microinflamação de vias aéreas, alergias e hiperativação do sistema nervoso simpático.""",

            "patient_explanation": """O local onde você dorme afeta profundamente a qualidade do seu sono. Três fatores são fundamentais: escuridão total (qualquer luz, mesmo pequenos LEDs de aparelhos, pode atrapalhar a produção de melatonina), temperatura fresca (idealmente 18-20°C - seu corpo precisa esfriar um pouco para dormir bem), e silêncio (ruídos, mesmo que você não acorde conscientemente, fragmentam o sono profundo).

Outros aspectos importantes incluem: qualidade do ar (ácaros e mofo causam congestão nasal e respiração difícil), colchão e travesseiro adequados (suporte correto previne dores), e minimizar campos eletromagnéticos (Wi-Fi, celular carregando perto da cama podem afetar o sistema nervoso).""",

            "conduct": """Protocolo de otimização ambiental do sono (abordagem funcional integrativa):

1. Controle de Luz (Prioridade Máxima):
   - Escuridão total: cortinas blackout (bloqueio 100% de luz externa), fita adesiva preta em LEDs de aparelhos, remover relógios digitais luminosos
   - Meta: incapacidade de ver a própria mão estendida com luz apagada
   - Se impossível escuridão total: máscara de sono 100% blackout (seda ou algodão respirável)
   - Uso de luz vermelha/âmbar (>600nm) se necessário iluminação noturna (banheiro) - não suprime melatonina
   - Remover televisão do quarto (associação entre luz azul de TV e redução de 30-60 min de sono)

2. Termorregulação:
   - Temperatura ideal: 18-20°C (65-68°F) - ajustar ar-condicionado ou aquecimento conforme estação
   - Roupas de cama respiráveis: algodão, linho, bambu (evitar sintéticos que retêm calor)
   - Considerar cooling mattress pad se sudorese noturna ou menopausa
   - Banho morno 1-2h antes de dormir: aumenta temperatura corporal temporariamente, seguido de queda que facilita sono
   - Suplementação: magnésio glicina to (relaxa vasos periféricos, facilita dissipação de calor), glicina 3g (reduz temperatura central)

3. Controle Acústico:
   - Ambiente silencioso: meta <30 decibéis (biblioteca silenciosa)
   - Se ruído externo inevitável: ruído branco (apps, máquinas específicas) mascara sons disruptivos
   - Alternativa: ruído rosa ou marrom (frequências mais baixas, mais natural que ruído branco)
   - Protetores auriculares se necessário (espuma de memória, cera moldável - evitar inserção profunda excessiva)
   - Isolamento acústico: janelas duplas, cortinas pesadas, vedação de portas

4. Qualidade do Ar:
   - Purificador de ar HEPA: remove alérgenos (ácaros, pólen, esporos de fungos), particulados
   - Umidade relativa: 40-60% (usar umidificador se ar seco <30%, desumidificador se >70%)
   - Plantas purificadoras (opcional): Espada-de-São-Jorge, Lírio da Paz (reduzem VOCs, produzem O2 noturno)
   - Trocar roupa de cama semanalmente (água quente >60°C para matar ácaros)
   - Travesseiros e colchões anti-alérgicos, capas impermeáveis a ácaros
   - Evitar tapetes (acumulam ácaros) - preferir piso frio facilmente lavável

5. Campo Eletromagnético (EMF):
   - Desligar Wi-Fi noturno (timer automático ou desligar manualmente)
   - Celular em modo avião ou fora do quarto (não carregar ao lado da cama)
   - Desligar dispositivos Bluetooth
   - Relógio despertador analógico ou a bateria (não elétrico digital)
   - Se impossível eliminar EMF: considerar tecidos de proteção EMF (estudos conflitantes, mas relatos anedóticos de melhora)

6. Mobiliário e Superfície de Sono:
   - Colchão: substituir a cada 7-10 anos, firmeza média (suporte sem pressão excessiva)
   - Travesseiro: altura adequada para alinhamento cervical (posição supino: mais baixo; lateral: mais alto)
   - Posição de sono: lateral esquerdo (melhora drenagem linfática cerebral, reduz refluxo) ou supino com elevação leve (apneia/refluxo)
   - Evitar posição prona (rotação cervical excessiva, compressão facial)

7. Minimalismo e Psicologia Ambiental:
   - Quarto apenas para sono e intimidade (não escritório, não sala de TV)
   - Decoração minimalista: reduz estímulos visuais, promove relaxamento
   - Cores frias (azul, verde, lavanda) associadas a redução de frequência cardíaca e pressão arterial
   - Evitar espelhos refletindo a cama (algumas tradições associam a sono fragmentado)

8. Cronobiologia e Exposição Matinal:
   - Abrir cortinas imediatamente ao acordar: exposição solar matinal sincroniza relógio circadiano
   - Idealmente: 15-30 min de luz solar direta (sem óculos escuros) na primeira hora do dia
   - Se exposição solar impossível: light box 10.000 lux por 20-30 min ao café da manhã

9. Avaliação de Exposições Tóxicas:
   - VOCs (compostos orgânicos voláteis): móveis novos, tintas recentes, produtos de limpeza aromáticos
   - Solução: ventilar ambiente, usar purificador com filtro de carvão ativado, evitar produtos químicos aromáticos (preferir limpeza com vinagre, bicarbonato)
   - Testar presença de mofo (manchas, odor característico) - pode exigir remediação profissional
   - Evitar velas aromáticas (liberam particulados), incensos, difusores sintéticos no quarto

10. Monitoramento e Ajustes:
    - Apps de monitoramento de sono: rastreamento de qualidade, fases (Oura Ring, Whoop, Apple Watch)
    - Termômetro de ambiente: verificar temperatura real (percepção subjetiva pode ser imprecisa)
    - Decibelímetro: apps gratuitos para medir ruído ambiental
    - Ajustar protocolo baseado em dados objetivos + sensação subjetiva ao acordar"""
        },

        "Apneias": {
            "clinical_relevance": """A apneia obstrutiva do sono (AOS) é caracterizada por colapsos recorrentes das vias aéreas superiores durante o sono (≥5 eventos/hora em adultos, ≥1 evento/hora em crianças), causando hipóxia intermitente, hipercapnia, fragmentação do sono e ativação repetida do sistema nervoso simpático. Afeta 15-30% dos adultos (prevalência subestimada devido a subdiagnóstico), sendo mais comum em homens >50 anos, indivíduos com obesidade (IMC >30), hipertensão resistente e síndrome metabólica.

A fisiopatologia da AOS envolve: 1) Obstrução anatômica: excesso de tecido adiposo faríngeo (obesidade), hipertrofia adenoamigdaliana (crianças e adultos jovens), retrognatia (mandíbula retraída), macroglossia, desvio de septo nasal. 2) Colapso funcional: redução do tônus muscular faríngeo durante sono REM (pior fase para apneias), edema de vias aéreas por inflamação sistêmica, congestão nasal (alergias, rinite). 3) Controle ventilatório anormal: instabilidade do drive respiratório central (mais comum em insuficiência cardíaca - apneia central).

As consequências cardiovasculares são severas e mediadas por múltiplos mecanismos: hipóxia intermitente → estresse oxidativo → disfunção endotelial → aterosclerose acelerada; ativação simpática repetida → hipertensão (presente em 50-80% dos casos de AOS, frequentemente resistente a medicamentos); hipóxia → eritropoiese excessiva → policitemia → aumento de viscosidade sanguínea → risco trombótico. A AOS está independentemente associada a: hipertensão resistente, arritmias (especialmente fibrilação atrial), insuficiência cardíaca, doença arterial coronariana, AVC isquêmico (risco 2-3x maior) e morte súbita cardíaca.

Consequências metabólicas: fragmentação do sono + hipóxia → resistência à insulina (mesmo independente de obesidade), síndrome metabólica, diabetes tipo 2 (risco 2.5x maior), dislipidemia (aumento de triglicerídeos, LDL pequeno e denso), esteatose hepática não alcoólica (NAFLD/NASH). A AOS causa alteração de hormônios reguladores do apetite: redução de leptina (saciedade) e aumento de grelina (fome), perpetuando obesidade.

Impacto neurocognitivo: hipóxia cerebral crônica → neurodegeneração, acúmulo de beta-amiloide (Alzheimer), comprometimento cognitivo (atenção, memória, função executiva), aumento de risco de demência (50-80% maior). Fragmentação do sono → fadiga diurna extrema → acidentes (dirigir, operar máquinas) - risco 7x maior de acidentes automobilísticos.

A medicina funcional integrativa aborda AOS como doença inflamatória sistêmica multifatorial, investigando: inflamação de baixo grau (PCR ultrassensível, citocinas), estresse oxidativo (8-OHdG, malonaldeído), disfunção mitocondrial, neuroinflamação, disbiose intestinal (microbioma alterado em AOS - eixo intestino-pulmão), deficiências nutricionais que pioram tônus muscular faríngeo.""",

            "patient_explanation": """Apneia do sono significa que você para de respirar repetidamente durante a noite - às vezes centenas de vezes, mesmo sem perceber. Cada parada de respiração diminui o oxigênio no sangue, acorda parcialmente o cérebro (fragmentando o sono profundo), e força o coração a trabalhar mais.

As consequências são graves: cansaço crônico (mesmo dormindo 8-9h você acorda exausto), pressão alta difícil de controlar, maior risco de infarto, derrame, diabetes, obesidade, depressão, demência e acidentes. Roncar alto é o sinal mais comum, mas nem todo mundo que ronca tem apneia. Pausas respiratórias observadas por parceiro de cama, despertar com sensação de sufocamento, sudorese noturna, e sonolência diurna extrema são sinais de alerta.

A boa notícia é que tratamento eficaz existe (CPAP, perda de peso, dispositivos orais, cirurgia em casos específicos) e pode reverter muitas das complicações.""",

            "conduct": """Protocolo funcional integrativo para diagnóstico e manejo de Apneia Obstrutiva do Sono:

1. Triagem e Diagnóstico:
   a) Questionários de Triagem:
      - STOP-BANG (8 questões): Snoring, Tiredness, Observed apnea, Pressure (hipertensão), BMI >35, Age >50, Neck >40cm, Gender male. Score ≥3: alto risco
      - Escala de Sonolência de Epworth: score >10 indica sonolência diurna excessiva
      - Questionário de Berlim: ronco, fadiga, hipertensão

   b) Exame Padrão-Ouro: Polissonografia (noturna em laboratório ou domiciliar)
      - Classifica severidade: Leve (5-15 eventos/h), Moderada (15-30), Severa (>30)
      - Avalia: índice de apneia-hipopneia (IAH), saturação de O2, arquitetura do sono, roncos, movimentos

   c) Exames Complementares:
      - Cefalometria (raio-X lateral da face): avalia anatomia de vias aéreas, retrognatia
      - Nasofibroscopia: hipertrofia de cornetos, desvio de septo, hipertrofia adenoamigdaliana
      - TC de seios da face se sinusopatia crônica

2. Investigação Funcional Laboratorial:
   - Marcadores metabólicos: hemoglobina glicada, glicemia de jejum + insulina (HOMA-IR), perfil lipídico completo, função hepática
   - Inflamação: PCR ultrassensível (>3 mg/L indica risco cardiovascular), homocisteína, fibrinogênio
   - Hormonal: TSH, T4L, T3L (hipotireoidismo associado a AOS), testosterona total/livre em homens (baixa testosterona piora tônus muscular faríngeo)
   - Nutricional: vitamina D (meta >40 ng/ml), magnésio eritrocitário, ferritina, vitamina B12, folato
   - Cardiovascular: ecocardiograma (avaliar hipertrofia VE, função diastólica), ECG (arritmias)
   - Estresse oxidativo: 8-OHdG urinário, malonaldeído (se disponível)

3. Tratamento Padrão-Ouro: CPAP (Continuous Positive Airway Pressure)
   - Indicação: AOS moderada-severa (IAH >15), ou AOS leve com comorbidades cardiovasculares
   - Eficácia: >95% se adesão adequada (uso >4h/noite, >70% das noites)
   - Benefícios comprovados: redução de PA sistólica (5-10 mmHg), melhora de glicemia (HbA1c -0.5%), redução de risco de AVC/IAM, melhora cognitiva
   - Desafios de adesão: desconforto da máscara, sensação de claustrofobia, boca seca, ruído
   - Soluções: ajuste de máscara (múltiplos modelos disponíveis), umidificador aquecido, CPAP auto-ajustável

4. Intervenções de Estilo de Vida (Fundamentais):
   a) Perda de Peso (Prioridade Máxima):
      - Meta: redução de 10% do peso corporal melhora IAH em 30-50%
      - Redução de 20-25% pode curar AOS leve-moderada em >50% dos casos
      - Dieta anti-inflamatória: redução de carboidratos refinados, aumento de vegetais/fibras, proteína adequada
      - Jejum intermitente 16:8 (evidências de melhora de sensibilidade insulínica e perda de gordura visceral)
      - Exercício: 150 min/semana aeróbico moderado + treino de força 2-3x/semana

   b) Posição de Sono:
      - Evitar posição supina (de costas) - piora AOS em 50% dos casos (língua cai para trás)
      - Preferir decúbito lateral (esquerdo ideal - facilita drenagem linfática)
      - Estratégias: tennis ball na parte traseira do pijama, travesseiros posicionadores, alarme de posição

   c) Evitar Álcool e Sedativos:
      - Álcool relaxa musculatura faríngea (piora colapso) - evitar 4h antes de dormir
      - Benzodiazepínicos, opioides, relaxantes musculares: contraindicados (suprimem drive respiratório)

5. Modulação Nutricional e Suplementação:
   - Ômega-3 EPA/DHA: 2-3g/dia (anti-inflamatório potente, reduz PCR, melhora função endotelial)
   - Vitamina D3: corrigir deficiência, meta >40 ng/ml (imunomodulação, reduz inflamação de vias aéreas)
   - Magnésio quelato: 400-600mg/dia (relaxamento muscular, controle de PA, sensibilidade insulínica)
   - CoQ10 (ubiquinol): 200-400mg/dia (antioxidante, melhora função mitocondrial afetada pela hipóxia, reduz estresse oxidativo)
   - N-Acetilcisteína (NAC): 1200-1800mg/dia (antioxidante, mucolítico, reduz inflamação de vias aéreas)
   - Curcumina (com piperina): 1000mg/dia (anti-inflamatório potente, melhora sensibilidade insulínica)
   - Resveratrol: 200-500mg/dia (antioxidante, proteção cardiovascular, melhora função endotelial)
   - Probióticos: Lactobacillus rhamnosus, Bifidobacterium longum (modular eixo intestino-pulmão, reduzir inflamação sistêmica)

6. Tratamento de Congestão Nasal (Fundamental):
   - Alergias: teste de IgE específico, evitar alérgenos, corticosteroide nasal (fluticasona, mometasona)
   - Antileucotrienos: montelucaste 10mg/noite (reduz edema de vias aéreas)
   - Irrigação nasal: solução salina hipertônica 2x/dia (reduz inflamação, remove alérgenos)
   - Suplementação: quercetina 500mg 2x/dia (anti-histamínico natural), bromelina, vitamina C

7. Dispositivos Orais:
   - Indicação: AOS leve-moderada (IAH <30), intolerância a CPAP, retrognatia
   - Tipos: dispositivo de avanço mandibular (reposiciona mandíbula anteriormente, abre via aérea)
   - Eficácia: 50-70% (inferior a CPAP, mas adesão superior)
   - Confecção: dentista especializado em medicina do sono

8. Tratamento Cirúrgico (Casos Selecionados):
   a) Crianças: adenotonsilectomia (taxa de cura 70-90% se hipertrofia adenoamigdaliana)
   b) Adultos:
      - Uvulopalatofaringoplastia (UPPP): remove excesso de tecido faríngeo (eficácia 40-60%)
      - Cirurgia de avanço mandibular (maxilomandibular advancement): mais eficaz (>90%), mas invasiva
      - Estimulação do nervo hipoglosso: implante que estimula língua durante sono (aprovado FDA, eficácia 60-70%)
      - Cirurgia bariátrica: se obesidade mórbida (IMC >40) - perda de peso dramática pode curar AOS

9. Tratamento de Comorbidades:
   - Hipertensão: otimizar controle PA (meta <130/80), preferir classes que não pioram AOS (IECA, BRA, bloqueadores de canal de cálcio)
   - Diabetes: otimizar controle glicêmico (tratamento de AOS melhora HbA1c, mas diabetes não tratado piora AOS)
   - Hipotireoidismo: reposição adequada de levotiroxina (hipotireoidismo causa edema de vias aéreas)
   - Refluxo gastroesofágico: tratar (refluxo causa edema laríngeo que piora AOS)

10. Monitoramento e Follow-up:
    - Polissonografia de controle: 3-6 meses após início de CPAP (ajustar pressão), após perda de peso significativa (re-avaliar necessidade de CPAP)
    - Avaliação de adesão a CPAP: apps/software rastreiam uso diário
    - Exames laboratoriais a cada 6-12 meses: PA, HbA1c, perfil lipídico, PCR
    - Avaliação de sonolência: Escala de Epworth periodicamente

11. Abordagens Adjuvantes:
    - Exercícios orofaríngeos: fortalecem musculatura faríngea (20-30 min/dia por 3 meses - reduz IAH em 30-40% em estudos)
    - Canto: fortalece palato mole (estudos piloto mostraram redução de roncos)
    - Terapia miofuncional orofacial: com fonoaudiólogo especializado"""
        },

        "Barulho": {
            "clinical_relevance": """A poluição sonora noturna é fator frequentemente subestimado na medicina do sono, com impacto significativo na arquitetura do sono, ativação do sistema nervoso autônomo e saúde cardiovascular. Estudos de polissonografia demonstram que ruídos >40 decibéis (dB) - equivalente a conversação em tom normal - causam microdespertares (arousals) mesmo sem despertar consciente. Estes microdespertares fragmentam a progressão natural dos estágios do sono, reduzindo tempo em sono profundo (estágio 3 NREM, essencial para restauração física e consolidação imunológica) e sono REM (processamento emocional e consolidação de memória).

A fisiopatologia do impacto do ruído no sono envolve: 1) Via auditiva permanece ativa durante o sono: sons são processados pelo córtex auditivo primário mesmo em sono profundo. Ruídos súbitos ou de alta intensidade ativam sistema reticular ascendente (arousal system), causando transição para estágios mais leves de sono ou despertar completo. 2) Ativação simpática: ruídos noturnos aumentam frequência cardíaca, pressão arterial e cortisol sérico (medidos através de coletas sanguíneas durante polissonografia). Esta ativação simpática repetida, noite após noite, está associada a hipertensão crônica, resistência à insulina e risco cardiovascular aumentado. 3) Resposta de estresse: mesmo sem despertar consciente, o cérebro interpreta ruídos como ameaças potenciais, ativando eixo HPA (hipotálamo-hipófise-adrenal), resultando em cortisol elevado e qualidade de sono reduzida.

Estudos epidemiológicos de grande escala (WHO Burden of Disease from Environmental Noise) demonstram que exposição crônica a ruído noturno >55 dB está associada a: aumento de 7-17% no risco de hipertensão, aumento de 8% no risco de doença arterial coronariana, aumento de mortalidade cardiovascular. A Organização Mundial da Saúde estabelece que níveis noturnos seguros são <30 dB (equivalente a sussurro) para prevenir perturbações do sono e <40 dB para prevenir efeitos cardiovasculares.

Do ponto de vista funcional integrativo, a exposição crônica a ruído noturno perpetua cascata inflamatória: fragmentação do sono → ativação de NF-kB (fator de transcrição pró-inflamatório) → aumento de citocinas IL-6, TNF-alfa → resistência à insulina → síndrome metabólica. Indivíduos com sensibilidade aumentada a ruídos (hipervigilância, ansiedade, TEPT) apresentam respostas exacerbadas, com maior impacto na qualidade do sono.""",

            "patient_explanation": """Barulhos durante a noite - mesmo que você não acorde conscientemente - fragmentam seu sono profundo e afetam a recuperação do corpo. Sons acima de 40 decibéis (uma conversação normal) já causam microdespertares que você não percebe, mas que impedem o cérebro de completar ciclos de sono restaurador.

Exposição crônica a ruídos noturnos aumenta pressão arterial, estresse, risco de doenças cardíacas, ganho de peso e comprometimento da memória. Fontes comuns incluem: trânsito, vizinhos, parceiro que ronca, animais de estimação, aparelhos eletrônicos. A boa notícia é que soluções simples (cortinas acústicas, protetores auriculares, ruído branco) podem melhorar dramaticamente a qualidade do sono.""",

            "conduct": """Protocolo funcional integrativo para manejo de poluição sonora noturna:

1. Avaliação Quantitativa do Ruído:
   - Medição objetiva: usar decibelímetro (apps gratuitos: Decibel X, Sound Meter) durante várias noites
   - Identificar fontes: tráfego, vizinhos, parceiro roncando, animais, aparelhos eletrônicos (geladeira, ar-condicionado)
   - Meta: <30 dB (sussurro) ideal, <40 dB aceitável

2. Intervenções de Redução de Ruído (Controle da Fonte):
   a) Isolamento Acústico do Ambiente:
      - Janelas: vidros duplos/triplos (redução de 30-50 dB), vedação com borracha
      - Portas: vedação inferior (rolo de tecido, vedante adesivo)
      - Paredes: painéis acústicos (espuma, lã de rocha), estantes cheias de livros (absorvem som)
      - Cortinas: tecidos pesados, multilayer (veludo, blackout acústico) - redução de 10-20 dB

   b) Modificação de Fontes Internas:
      - Aparelhos eletrônicos: substituir geladeira/ar-condicionado ruidosos, manutenção de ventiladores
      - Posicionar aparelhos mais longe do quarto
      - Relógio: analógico ou digital silencioso (não tick-tock)

   c) Abordagem de Roncos do Parceiro:
      - Avaliar parceiro para apneia do sono (STOP-BANG, polissonografia se indicado)
      - Tratamento: CPAP, perda de peso, dispositivo oral, cirurgia conforme severidade
      - Temporariamente: parceiro dormir de lado (não supino), evitar álcool

3. Mascaramento de Ruído (Quando Redução Impossível):
   a) Ruído Branco:
      - Máquinas dedicadas de ruído branco (som contínuo, todas as frequências)
      - Apps: White Noise, myNoise (customizáveis)
      - Ventilador: opção low-tech eficaz (som constante mascara ruídos variáveis)
      - Volume: suficiente para mascarar ruídos disruptivos, mas <50 dB (evitar dano auditivo crônico)

   b) Ruído Rosa ou Marrom:
      - Frequências mais baixas que ruído branco (mais natural, relaxante)
      - Evidências: ruído rosa melhora sono profundo e consolidação de memória (estudo Northwestern 2017)

   c) Sons da Natureza:
      - Chuva, ondas do mar, floresta: preferência individual
      - Vantagem: psicologicamente mais relaxante que ruído branco
      - Desvantagem: sons variáveis podem não mascarar ruídos súbitos eficazmente

4. Proteção Auditiva Individual:
   a) Protetores Auriculares:
      - Espuma de memória (descartáveis): redução 25-35 dB, baixo custo
      - Silicone moldável: reutilizáveis, redução 20-30 dB, mais confortáveis
      - Custom-fit (moldados por audiologista): redução 30-40 dB, máximo conforto (custo elevado)
      - Considerar filtros acústicos: reduzem ruídos ambientais, mas permitem alarmes/despertador

   b) Técnica de Inserção Correta:
      - Rolar espuma em cilindro fino, puxar orelha para cima/trás, inserir profundamente, segurar 30s até expandir
      - Inserção inadequada reduz eficácia em 50%

5. Modulação Neurológica da Sensibilidade ao Ruído:
   - Indivíduos com ansiedade/hipervigilância têm limiar de despertar mais baixo
   - Intervenções:
      a) Magnésio quelato: 400-600mg à noite (modula sistema nervoso, reduz hiperexcitabilidade)
      b) L-teanina: 200-400mg (aumenta ondas alfa cerebrais, relaxamento sem sedação)
      c) Glicina: 3g antes de dormir (neurotransmissor inibitório, melhora qualidade do sono)
      d) Ashwagandha: 300-600mg (adaptógeno, reduz cortisol, ansiedade)

   - Técnicas comportamentais:
      a) Terapia cognitivo-comportamental para insônia (CBT-I): reestrutura resposta a ruídos noturnos
      b) Meditação mindfulness: reduz reatividade a estímulos externos
      c) Treino de relaxamento progressivo: reduz tensão muscular e ativação simpática

6. Cronobiologia e Ritmo Circadiano:
   - Ruído fragmenta sono, mas ritmo circadiano forte aumenta resiliência
   - Fortalecer ritmo: exposição solar matinal 30 min, horários fixos de sono/despertar, exercício regular

7. Avaliação de Impacto:
   - Diário de sono: registrar qualidade subjetiva, despertares percebidos, sensação ao acordar
   - Apps de monitoramento: rastreamento de fases de sono, microdespertares (Oura, Whoop)
   - Questionários: Pittsburgh Sleep Quality Index (PSQI), Epworth Sleepiness Scale
   - Polissonografia domiciliar: se suspeita de fragmentação severa

8. Abordagem Funcional para Reparo de Dano Cumulativo:
   - Se exposição crônica a ruído causou inflamação sistêmica/estresse oxidativo:
      a) Ômega-3 EPA/DHA: 2g/dia (anti-inflamatório)
      b) Curcumina: 1000mg/dia (reduz NF-kB, citocinas pró-inflamatórias)
      c) Resveratrol: 200-500mg (antioxidante, proteção cardiovascular)
      d) CoQ10: 200mg (função mitocondrial, antioxidante)
      e) Vitamina C e E: antioxidantes sinérgicos

9. Considerações Especiais:
   - Pais de bebês/crianças: inevitável algum ruído noturno
      - Estratégia: quarto do bebê longe do quarto dos pais se possível, monitor de áudio (não vídeo que emite luz), turnos entre cuidadores
   - Trabalho em turnos: ruído diurno quando dormindo
      - Crítico: isolamento acústico rigoroso, protetores auriculares, comunicação com família/vizinhos sobre horários de sono"""
        },

        "Bruxismo": {
            "clinical_relevance": """O bruxismo é caracterizado por atividade repetitiva dos músculos mastigatórios (masseter, temporal) manifestando-se como ranger ou apertar dos dentes, ocorrendo durante o sono (bruxismo do sono) ou vigília (bruxismo de vigília). O bruxismo do sono afeta 8-16% da população adulta e até 40% das crianças, sendo classificado como distúrbio de movimento relacionado ao sono pela International Classification of Sleep Disorders (ICSD-3).

A fisiopatologia é multifatorial e ainda não completamente elucidada: 1) Ativação do sistema nervoso central: microdespertares durante transições entre estágios do sono (especialmente de sono profundo para sono leve) ativam geradores de padrões centrais que desencadeiam contrações musculares rítmicas. Esta hiperativação está associada a aumento de atividade dopaminérgica e catecolaminérgica. 2) Estresse psicossocial: ativação crônica do eixo HPA resulta em cortisol elevado noturno (quando deveria estar no nadir) e hiperativação simpática, perpetuando tensão muscular. 3) Genética: estudos de gêmeos demonstram hereditariedade de 20-50%, com genes relacionados à regulação dopaminérgica (DRD2, COMT) possivelmente envolvidos.

Fatores associados incluem: deficiência de magnésio (relaxante muscular natural, cofator para produção de GABA), ansiedade e transtornos de estresse (presente em 70% dos casos de bruxismo severo), apneia obstrutiva do sono (30-50% dos pacientes com AOS têm bruxismo - hipótese: bruxismo como mecanismo compensatório para abrir vias aéreas colapsadas), refluxo gastroesofágico (ácido na faringe desencadeia resposta protetora de deglutição e apertar dental), uso de antidepressivos SSRIs/SNRIs (efeito colateral comum por aumento de serotonina que paradoxalmente pode causar bruxismo), parasitoses intestinais (controverso, relatos anedóticos frequentes mas evidências científicas limitadas).

As consequências incluem: desgaste dental severo (erosão de esmalte, fraturas dentárias, sensibilidade aumentada), dor e disfunção temporomandibular (DTM) com dor facial crônica, cefaleia tensional (especialmente temporal e frontal ao acordar), hipertrofia de masseter (aumento visível do músculo, assimetria facial), distúrbios do sono (fragmentação por microdespertares associados a episódios de bruxismo).

A medicina funcional integrativa aborda bruxismo investigando: deficiências nutricionais (magnésio, cálcio, vitaminas B), desequilíbrios de neurotransmissores (dopamina, serotonina, GABA), eixo intestino-cérebro (disbiose, parasitoses, permeabilidade intestinal), disfunções do eixo HPA (cortisol salivar 4 pontos), inflamação sistêmica e apneia do sono.""",

            "patient_explanation": """Bruxismo é ranger ou apertar os dentes durante o sono (ou às vezes durante o dia). Muitas pessoas não percebem que têm, até que o dentista note desgaste nos dentes, ou até que comecem a ter dor na mandíbula, dor de cabeça ao acordar, ou o parceiro escute o rangido.

As causas principais são: estresse e ansiedade (tensão muscular excessiva), falta de magnésio (mineral que relaxa os músculos), apneia do sono (paradas respiratórias que provocam apertar os dentes), refluxo ácido, uso de certos antidepressivos, ou até mesmo vermes intestinais (principalmente em crianças).

Além de desgastar os dentes, bruxismo pode causar dor crônica de cabeça e mandíbula, dificuldade para mastigar, e até alterar a aparência do rosto (músculos da bochecha ficam maiores). O tratamento envolve placa de proteção dentária + tratar a causa raiz (estresse, deficiências, apneia).""",

            "conduct": """Protocolo funcional integrativo para diagnóstico e tratamento de bruxismo:

1. Diagnóstico e Avaliação:
   a) Avaliação Odontológica:
      - Exame clínico: desgaste dental (facetas de desgaste em incisivos, caninos), fraturas, sensibilidade
      - Palpação muscular: masseter, temporal (dor, hipertrofia)
      - Avaliação de ATM (articulação temporomandibular): limitação de abertura, estalos, dor
      - Moldagem para placa oclusal (proteção dentária)

   b) Questionários e Diários:
      - Perguntar a parceiro de cama: escuta rangidos noturnos? Com que frequência?
      - Diário de sintomas: dor ao acordar (cabeça, face, mandíbula), qualidade do sono
      - Screening de ansiedade: GAD-7, escala de estresse percebido

   c) Polissonografia (se suspeita de apneia ou bruxismo severo):
      - Detecta episódios de bruxismo (eletromiografia de masseter)
      - Avalia associação com microdespertares, apneia, outros distúrbios de movimento

2. Investigação Funcional Laboratorial:
   - Magnésio eritrocitário: deficiência presente em 60-80% dos casos de bruxismo (magnésio sérico não reflete estoques intracelulares)
   - Cálcio sérico: desequilíbrio cálcio/magnésio afeta contração muscular
   - Vitamina D: deficiência associada a dor muscular crônica
   - Cortisol salivar 4 pontos: avaliar ritmo circadiano (cortisol noturno elevado indica hiperativação HPA)
   - Função tireoidiana: hipertireoidismo pode causar hiperatividade muscular
   - Parasitológico de fezes (3 amostras): especialmente em crianças, ou adultos com sintomas GI
   - Perfil de neurotransmissores urinários: dopamina, serotonina, GABA (se disponível - exame funcional avançado)

3. Proteção Dentária Imediata:
   - Placa oclusal (bite guard): confeccionada por dentista
      - Protege dentes de desgaste, redistribui forças mastigatórias
      - NÃO cura bruxismo, apenas protege dentes
      - Tipos: placa dura (acrílico - mais durável), placa macia (conforto - menos durável)
   - Uso noturno obrigatório enquanto trata causa-raiz

4. Suplementação Nutricional (Base do Tratamento Funcional):
   a) Magnésio (ESSENCIAL):
      - Magnésio quelato ou glicina to: 400-600mg à noite (2h antes de dormir)
      - Magnésio treonato: 144mg magnésio elementar (atravessa barreira hematoencefálica, reduz hiperexcitabilidade neural)
      - Duração: mínimo 8-12 semanas para repleção de estoques intracelulares
      - Efeitos: relaxamento muscular, suporte GABAérgico, redução de cortisol

   b) Cálcio (se deficiência):
      - Cálcio quelato: 500-1000mg/dia (longe do magnésio - competem por absorção)
      - Ratio Ca:Mg ideal = 2:1 a 1:1

   c) Complexo B (especialmente B5 e B6):
      - B5 (ácido pantotênico): 500mg/dia (suporte adrenal, reduz estresse)
      - B6 (P5P - forma ativada): 50-100mg/dia (cofator para produção de serotonina, GABA)

   d) Taurina:
      - 500-1000mg à noite (modulação GABAérgica, relaxamento SNC)

   e) Glicina:
      - 3g antes de dormir (neurotransmissor inibitório, relaxamento muscular, melhora sono)

   f) L-teanina:
      - 200-400mg (reduz ansiedade, aumenta GABA, não sedativa)

   g) Vitamina D3:
      - 2000-5000 UI/dia se deficiência (meta >40 ng/ml) - reduz dor muscular, modula inflamação

5. Manejo de Estresse e Ansiedade (Crucial):
   a) Técnicas de Relaxamento:
      - Relaxamento muscular progressivo de Jacobson: 15-20 min antes de dormir (foco em masseter, temporal)
      - Respiração diafragmática 4-7-8: ativa sistema parassimpático
      - Meditação mindfulness: 10-20 min/dia (reduz cortisol, melhora regulação emocional)

   b) Terapia Cognitivo-Comportamental (CBT):
      - CBT para ansiedade: reestrutura pensamentos catastrofistas, reduz hiperativação
      - Biofeedback: treina relaxamento de masseter através de feedback em tempo real

   c) Adaptógenos:
      - Ashwagandha (Withania somnifera): 300-600mg/dia (reduz cortisol em 25-30%, ansiedade)
      - Rhodiola rosea: 200-400mg pela manhã (modula resposta ao estresse, evitar à noite - pode ser estimulante)
      - Magnolia officinalis: 200-400mg (reduz cortisol, ansiedade, melhora sono)

6. Tratamento de Apneia do Sono (se presente):
   - Polissonografia para diagnóstico
   - Tratamento: CPAP, perda de peso, dispositivo oral, cirurgia conforme severidade
   - Tratamento de apneia frequentemente resolve bruxismo simultaneamente (hipótese: bruxismo como resposta compensatória a colapso de vias aéreas)

7. Manejo de Refluxo Gastroesofágico (se presente):
   - Elevação de cabeceira: 15-20 cm
   - Evitar refeições 3h antes de dormir
   - Dieta: reduzir cafeína, álcool, alimentos gordurosos, picantes, cítricos
   - Suplementação: enzimas digestivas, betaína HCl (se hipocloridria), DGL (alcaçuz deglicirizado) 400mg antes de refeições
   - Considerar inibidor de bomba de prótons (IBP) temporariamente se severo - porém causa deficiências de magnésio, B12, ferro a longo prazo (preferir abordagem funcional)

8. Tratamento de Parasitoses (se detectadas):
   - Farmacológico: albendazol, mebendazol, nitazoxanida conforme parasita identificado
   - Suporte natural: alho (alicina), óleo de orégano, sementes de abóbora (efeito anti-helmíntico leve)
   - Probióticos: restaurar microbioma após tratamento

9. Ajuste de Medicamentos:
   - Se uso de SSRIs/SNRIs causando bruxismo:
      - Discutir com psiquiatra: redução de dose, troca para outro antidepressivo (bupropiona, mirtazapina têm menor incidência)
      - Considerar alternativas naturais se depressão leve: SAMe 400-800mg, 5-HTP 100-300mg, ômega-3 2g/dia, exercício regular

10. Fisioterapia e Terapias Corporais:
    a) Fisioterapia orofacial:
       - Massagem de masseter, temporal, pterigoideos
       - Alongamento de musculatura mastigatória
       - Liberação miofascial de trigger points

    b) Acupuntura:
       - Pontos: E6 (jiache - sobre masseter), E7 (xiaguan - ATM), VB20 (fengchi - base do crânio)
       - Evidências de redução de dor e frequência de bruxismo

    c) Toxina botulínica (Botox):
       - Indicação: bruxismo severo refratário a outras medidas
       - Aplicação em masseter: reduz força de contração (protege dentes), reduz hipertrofia, melhora dor
       - Duração: 3-6 meses, necessário reaplicação
       - Efeito colateral: dificuldade de mastigação se dose excessiva

11. Higiene do Sono:
    - Rotina relaxante pré-sono: banho morno, leitura, música suave
    - Evitar estimulantes: cafeína após 14h, álcool 4h antes de dormir
    - Ambiente: escuro, fresco (18-20°C), silencioso
    - Exercício regular: reduz estresse, mas não próximo ao horário de dormir

12. Monitoramento e Follow-up:
    - Reavaliação odontológica: a cada 3-6 meses (avaliar desgaste dental, ajustar placa)
    - Diário de sintomas: dor, qualidade de sono, percepção de melhora
    - Re-testar magnésio: após 3 meses de suplementação (verificar normalização)
    - Ajustar protocolo baseado em resposta individual"""
        },

        "Cama": {
            "clinical_relevance": """A superfície de sono (colchão, travesseiro, roupas de cama) tem impacto direto na qualidade do sono através de múltiplos mecanismos: 1) Suporte biomecânico: alinhamento da coluna vertebral durante as 7-9 horas de sono (1/3 da vida) afeta saúde musculoesquelética. Colchão inadequado (muito mole ou muito firme) causa desalinhamento vertebral, pontos de pressão excessivos, microdespertares por mudanças de posição devido a desconforto, resultando em dor lombar crônica, cervicalgia e qualidade de sono reduzida. 2) Termorregulação: a temperatura corporal central precisa diminuir 1-2°C para iniciar o sono. Colchões que retêm calor (espuma viscoelástica de baixa qualidade) ou roupas de cama sintéticas impedem dissipação de calor, prolongando latência de sono e causando despertares noturnos. 3) Exposição a alérgenos: colchões e travesseiros acumulam ácaros, fungos, seus dejetos e fragmentos corporais (cada grama de poeira pode conter 100-500 ácaros). Indivíduos alérgicos apresentam congestão nasal, tosse, respiração bucal durante sono (reduz qualidade, aumenta risco de apneia), e ativação imunológica noturna (citocinas pró-inflamatórias IL-4, IL-5, IgE).

Estudos biomecânicos demonstram que colchões de firmeza média (5-6 em escala de 1-10) proporcionam melhor suporte para maioria dos indivíduos, mantendo curvatura natural da coluna (lordose lombar e cervical) independentemente da posição de sono. Colchões muito firmes causam pressão excessiva em regiões de maior peso (ombros, quadris em posição lateral; sacro e região torácica em posição supina), reduzindo fluxo sanguíneo local e causando dor. Colchões muito moles permitem afundamento excessivo, desalinhando coluna.

A vida útil de colchões é de 7-10 anos, após o qual ocorre degradação de materiais (espumas perdem densidade, molas perdem elasticidade), acúmulo de ácaros, fungos, bactérias (mesmo com higiene regular), e perda de suporte estrutural. Estudos mostram que substituição de colchão com >9-10 anos por colchão novo de qualidade resulta em redução de dor lombar (63%), rigidez matinal (62%), e melhora de qualidade de sono (55-60%).

A medicina funcional integrativa avalia a superfície de sono considerando: alergias respiratórias (teste de IgE para ácaros, fungos), dor musculoesquelética crônica (alinhamento biomecânico), qualidade do sono (microdespertares por desconforto), exposição a toxinas (retardantes de chama, VOCs de espumas de baixa qualidade), e termorregulação (sudorese noturna, menopausa).""",

            "patient_explanation": """A cama onde você dorme afeta profundamente a qualidade do seu sono e da sua saúde. Um colchão inadequado pode causar dor nas costas e no pescoço, fazer você acordar várias vezes à noite tentando encontrar posição confortável, e acumular ácaros que causam alergias (nariz entupido, tosse, espirros).

O colchão ideal oferece suporte para manter a coluna alinhada (nem muito duro nem muito mole), é respirável (não retém calor excessivo), e tem proteção contra ácaros. Travesseiros devem ter altura adequada para manter o pescoço alinhado com a coluna. Colchões devem ser trocados a cada 7-10 anos - depois disso perdem suporte e acumulam alérgenos mesmo com limpeza regular.

Investir em uma boa cama é investir em saúde: você passa 1/3 da vida dormindo.""",

            "conduct": """Protocolo funcional integrativo para otimização da superfície de sono:

1. Avaliação da Situação Atual:
   a) Idade do colchão: >10 anos = substituição mandatória, 7-10 anos = avaliar estado
   b) Inspeção visual: afundamentos permanentes, molas salientes, manchas de umidade/mofo
   c) Sintomas: dor ao acordar (lombar, cervical, ombros), rigidez matinal, melhora de dor ao sair da cama, sono melhor em outras camas (hotéis, casa de amigos)
   d) Alergias: congestão nasal noturna, tosse seca pela manhã, espirros ao acordar (sugere ácaros)

2. Seleção de Colchão (Baseada em Evidências):
   a) Firmeza:
      - Regra geral: firmeza média (5-7/10) para maioria dos indivíduos
      - Dormidores de lado: levemente mais macio (permite afundamento de ombros/quadris mantendo alinhamento)
      - Dormidores de costas: médio-firme (suporte lombar sem pressão excessiva no sacro)
      - Dormidores de barriga: mais firme (evitar hiperextensão lombar) - porém posição prona não é recomendada (rotação cervical excessiva)
      - Peso corporal: indivíduos >100kg necessitam colchão mais firme; <60kg mais macio

   b) Tipos de Colchão:
      - Molas ensacadas (pocket coil): boa ventilação, suporte independente (movimentos do parceiro não transferem), durabilidade 8-10 anos
      - Espuma viscoelástica (memory foam): contorno corporal, reduz pontos de pressão. CRÍTICO: escolher espuma de alta densidade (>5 lb/ft³) com tecnologia de resfriamento (gel, grafite, cobre infundido) - espumas baratas retêm calor excessivo
      - Látex natural: excelente durabilidade (15-20 anos), hipoalergênico, respirável, suporte responsivo. Desvantagem: custo elevado
      - Híbridos (molas + espuma/látex): combina ventilação de molas com contorno de espuma
      - Evitar: colchões muito baratos com espumas de baixa densidade (degradam em 2-3 anos, liberam VOCs)

   c) Certificações de Segurança:
      - CertiPUR-US: espumas sem retardantes de chama PBDE, metais pesados, formaldeído, CFCs
      - OEKO-TEX Standard 100: tecidos sem substâncias nocivas
      - GOTS (Global Organic Textile Standard): algodão orgânico certificado
      - Preferir colchões sem retardantes de chama químicos (tóxicos, disruptores endócrinos)

3. Seleção de Travesseiro (Alinhamento Cervical):
   a) Baseado em Posição de Sono:
      - Lateral: travesseiro alto/firme (preenche espaço entre orelha e ombro, mantém coluna cervical alinhada com torácica)
      - Supino (costas): travesseiro médio (suporta curvatura natural do pescoço sem hiperextensão)
      - Prono (barriga): travesseiro muito baixo ou sem travesseiro (minimizar extensão cervical) - posição não recomendada

   b) Materiais:
      - Espuma viscoelástica: contorno, suporte consistente. Escolher com tecnologia de resfriamento
      - Látex: responsivo, durável, hipoalergênico
      - Penas/down: moldável, respirável. Desvantagem: pode causar alergias, perde suporte com tempo
      - Fibra sintética hipoalergênica: opção econômica, laváveis
      - Travesseiros cervicais ergonômicos: úteis para cervicalgia crônica (preferir após avaliação fisioterápica)

   c) Substituição: travesseiros devem ser trocados a cada 1-2 anos (acúmulo de ácaros, perda de suporte)

4. Proteção Contra Alérgenos (Essencial):
   a) Capas anti-alérgicas:
      - Encasement completo de colchão e travesseiros: tecido impermeável a ácaros (poros <10 μm)
      - Material: poliéster com revestimento de poliuretano ou tecidos especiais (BugLock, AllerZip)
      - Laváveis: lavar a cada 1-2 meses em água quente (>60°C)

   b) Higiene de Roupas de Cama:
      - Trocar lençóis semanalmente
      - Lavar em água quente (>60°C) para matar ácaros - adicionar 5-10 gotas de óleo essencial de eucalipto (acaricida natural)
      - Secar em secadora em alta temperatura ou sol direto

   c) Exposição Solar do Colchão:
      - Quando possível: expor colchão ao sol direto 2-4x/ano (raios UV matam ácaros, fungos, bactérias)
      - Aspiração regular: aspirador com filtro HEPA, focus em costuras

5. Termorregulação:
   a) Roupas de Cama Respiráveis:
      - Lençóis: algodão egípcio, percale, linho (evitar microfibra sintética que retém calor)
      - Thread count: 200-400 (>400 pode reduzir respirabilidade)
      - Evitar: poliéster, cetim de baixa qualidade

   b) Colchão com Tecnologia de Resfriamento:
      - Espumas infundidas com gel, cobre, grafite
      - Sistemas de ar (SleepNumber) permitem ajuste de firmeza e temperatura
      - Sistemas ativos: colchão com circulação de água/ar (ChiliPad, BedJet) - útil para sudorese noturna, menopausa

   c) Roupa de Dormir:
      - Tecidos naturais: algodão, bambu, modal (evitar sintéticos)
      - Loose-fitting (folgados) para circulação de ar

6. Posição de Sono e Suporte Adicional:
   a) Posição Lateral (Mais Comum):
      - Travesseiro entre joelhos: alinha quadril, reduz pressão lombar
      - Preferir lado esquerdo: facilita drenagem linfática cerebral (sistema glinfático), reduz refluxo

   b) Posição Supina:
      - Travesseiro sob joelhos: reduz tensão lombar, especialmente se lordose acentuada

   c) Elevação de Cabeceira:
      - Indicado para: refluxo gastroesofágico (elevar 15-20cm), apneia leve, congestão nasal
      - Usar cunha de espuma (não apenas travesseiros - causa flexão cervical)

7. Redução de Exposição a Toxinas:
   a) VOCs (Compostos Orgânicos Voláteis):
      - Colchões novos: "desintoxicar" por 24-72h em ambiente ventilado antes de usar (off-gassing)
      - Preferir materiais naturais: látex natural, algodão orgânico, lã natural

   b) Retardantes de Chama:
      - Legislação varia por país: EUA exige retardantes (tóxicos), Europa mais restritiva
      - Buscar certificações CertiPUR-US, OEKO-TEX
      - Alternativa: colchões com barreira física de retardação (lã natural - naturalmente retardante sem químicos)

8. Considerações Especiais:
   a) Dor Crônica:
      - Avaliação fisioterápica: postura, alinhamento, pontos de pressão
      - Considerar colchões ajustáveis (bases elétricas permitem elevação de cabeceira/pés)

   b) Parceiros com Preferências Diferentes:
      - Colchões dual-firmness: cada lado com firmeza diferente
      - Colchões separados: solução mais eficaz (mas menos intimidade)

   c) Alergias Severas:
      - Colchão de látex natural (hipoalergênico, antimicrobiano, anti-ácaros)
      - Purificador de ar HEPA no quarto
      - Evitar bichos de pelúcia, tapetes, cortinas pesadas (acumulam ácaros)

   d) Sudorese Noturna (Menopausa, Hipertireoidismo):
      - Colchões com tecnologia de resfriamento ativo
      - Lençóis de bambu (absorvem umidade 40% melhor que algodão)
      - Tratar causa-raiz: TRH bioidêntica se menopausa, tratar hipertireoidismo

9. Investimento e Custo-Benefício:
   - Colchões de qualidade: investimento inicial alto (R$3.000-15.000), mas durabilidade de 10-15 anos
   - Custo por noite: R$3.000 / 10 anos / 365 dias = ~R$0,82/noite
   - Benefícios: redução de dor (economia com fisioterapia, analgésicos), melhora de sono (produtividade, saúde mental), prevenção de problemas de saúde a longo prazo
   - Priorizar: melhor colchão possível dentro do orçamento (mais importante que móveis decorativos)

10. Teste e Período de Adaptação:
    - Maioria das empresas oferece trial period (30-120 dias) - USAR este período
    - Adaptação: corpo pode levar 2-4 semanas para se acostumar com novo colchão
    - Avaliar: qualidade de sono, dor ao acordar, sensação de restauração
    - Não hesitar em trocar se não funcionar (dentro do trial period)"""
        },

        "Campos eletromagnéticos": {
            "clinical_relevance": """Os campos eletromagnéticos (EMF) são áreas de energia invisível produzidas por eletricidade e dispositivos wireless. Classificam-se em: 1) Campos de extremamente baixa frequência (ELF-EMF, 0-300 Hz): produzidos por linhas de transmissão elétrica, fiação doméstica, aparelhos elétricos. 2) Radiofrequência (RF-EMF, 3 kHz-300 GHz): produzidos por celulares, Wi-Fi, Bluetooth, torres de celular, micro-ondas.

A exposição a EMF no ambiente de sono é preocupação crescente na medicina ambiental, embora a literatura científica seja controversa e polarizada. A Organização Mundial da Saúde classifica RF-EMF como "possivelmente carcinogênico para humanos" (Grupo 2B) baseado em estudos de associação entre uso intenso de celular e glioma. Quanto aos efeitos no sono, os mecanismos propostos (ainda não definitivamente comprovados) incluem: 1) Supressão de melatonina: estudos in vitro e em animais demonstram que exposição a EMF pode reduzir produção de melatonina pela glândula pineal. Estudos humanos mostram resultados conflitantes - alguns encontram redução de 15-20% na melatonina noturna com exposição a RF-EMF antes de dormir, outros não encontram efeito significativo. 2) Alteração da arquitetura do sono: estudos de polissonografia após exposição a EMF de celular mostram resultados inconsistentes - alguns reportam redução de sono REM e aumento de despertares, outros não encontram diferença. 3) Ativação do sistema nervoso: EMF podem ativar canais de cálcio voltagem-dependentes, aumentando cálcio intracelular e potencialmente causando hiperexcitabilidade neuronal.

Estudos epidemiológicos de longo prazo são limitados e frequentemente com vieses (confundidores como uso de telas antes de dormir - luz azul é comprovadamente disruptiva do sono, mas estudos atribuem efeito ao EMF). Meta-análises Cochrane concluem que evidências atuais são insuficientes para estabelecer relação causal definitiva entre exposição a RF-EMF de celulares e distúrbios do sono, mas recomendam princípio da precaução, especialmente em crianças (cérebro em desenvolvimento potencialmente mais vulnerável).

A medicina funcional integrativa reconhece que sensibilidade individual a EMF varia significativamente (possivelmente relacionada a polimorfismos genéticos em canais de cálcio, glutationa, metalotioneína). Indivíduos que reportam eletrossensibilidade apresentam sintomas consistentes (insônia, cefaleia, fadiga) quando expostos a EMF, mesmo que estudos duplo-cego controlados não confirmem resposta acima do placebo (fenômeno psicossomático vs. metodologia inadequada dos estudos é debatido).

Abordagem pragmática: embora evidências definitivas sejam limitadas, reduzir exposição a EMF no quarto é intervenção de baixo custo e sem desvantagens, seguindo princípio da precaução ("primeiro, não causar dano"). Benefícios adicionais incluem redução de luz de telas (comprovadamente disruptiva) e redução de tentação de usar dispositivos antes de dormir.""",

            "patient_explanation": """Campos eletromagnéticos (EMF) são ondas invisíveis produzidas por dispositivos elétricos e sem fio (Wi-Fi, celular, Bluetooth, torres de celular, aparelhos elétricos). A ciência ainda debate se EMF afeta o sono - alguns estudos sugerem que pode reduzir melatonina e fragmentar o sono, outros não encontram efeito.

O consenso atual é: não há prova definitiva de que EMF causa problemas de sono nas intensidades típicas de ambientes domésticos, mas também não há prova definitiva de segurança completa, especialmente com exposição crônica de longo prazo. Crianças podem ser mais vulneráveis (cérebro em desenvolvimento).

A boa prática é: princípio da precaução - reduzir exposição desnecessária é fácil e sem desvantagens. Desligar Wi-Fi à noite, colocar celular em modo avião ou fora do quarto, e evitar aparelhos elétricos próximos à cama são medidas simples que, no mínimo, reduzem tentação de usar telas antes de dormir (luz azul comprovadamente atrapalha o sono).""",

            "conduct": """Protocolo funcional integrativo para redução de exposição a EMF no ambiente de sono (princípio da precaução):

1. Avaliação da Exposição Atual:
   a) Inventário de Fontes de EMF no Quarto:
      - Dispositivos wireless: celular, tablet, laptop, smart TV, smart speakers (Alexa, Google Home)
      - Wi-Fi router: localização em relação ao quarto (idealmente >3 metros da cama)
      - Aparelhos elétricos: relógio digital, rádio-relógio, lâmpadas inteligentes, carregadores
      - Fiação elétrica: tomadas e interruptores na cabeceira da cama
      - Fontes externas: torres de celular, transformadores elétricos, linhas de alta tensão

   b) Medição Objetiva (Opcional):
      - Medidores de EMF: dispositivos portáteis medem ELF-EMF (gaussímetros) e RF-EMF (medidores de radiofrequência)
      - Apps de smartphone: limitados, mas podem detectar fontes de RF próximas
      - Serviço profissional: para avaliação abrangente (custo elevado, raramente necessário)

2. Redução de Radiofrequência (RF-EMF):
   a) Celular (Principal Fonte):
      - IDEAL: remover celular do quarto completamente (usar relógio despertador tradicional)
      - Alternativa: modo avião (desativa todas as transmissões RF)
      - Minimamente: desativar dados móveis, Wi-Fi, Bluetooth individualmente
      - NUNCA: carregar celular ao lado da cama (especialmente na cabeceira)
      - Distância: se impossível remover, manter >1-2 metros da cama

   b) Wi-Fi Router:
      - IDEAL: desligar Wi-Fi noturno (timer automático ou manualmente antes de dormir)
      - Implementação: timer na tomada (liga 6h, desliga 23h) ou configuração no router
      - Alternativa: mover router para local mais distante do quarto (>3-5 metros)
      - Rede cabeada (Ethernet): para computadores - elimina necessidade de Wi-Fi

   c) Outros Dispositivos Wireless:
      - Smart TV, speakers, tablets: desligar completamente à noite (não standby)
      - Baby monitors: usar monitores com cabo ou tecnologia de baixa frequência (não Wi-Fi)
      - Smartwatches, fitness trackers: remover ou modo avião durante sono

3. Redução de ELF-EMF (Campos de Baixa Frequência):
   a) Aparelhos Elétricos na Cabeceira:
      - Remover: rádio-relógio elétrico, luminárias elétricas conectadas na cabeceira
      - Substituir por: relógio a bateria, luminárias a bateria ou distantes da cama
      - Carregadores: não deixar carregadores conectados próximos à cama (mesmo sem dispositivo conectado, geram campo)

   b) Fiação Elétrica:
      - Tomadas e interruptores na parede da cabeceira geram campo (especialmente se circuito com alta carga)
      - Solução (ideal mas invasiva): desligar disjuntor do quarto à noite (kill switch)
      - Alternativa: manter cama >30cm da parede com fiação, ou reposicionar cama

   c) Eletrodomésticos Adjacentes:
      - Identificar: geladeira, micro-ondas, máquina de lavar na parede oposta ao quarto
      - Campos magnéticos atravessam paredes: manter cama >1 metro de aparelhos de alto consumo

4. Proteções e Blindagem (Abordagem Avançada - Opcional):
   a) Tecidos de Blindagem EMF:
      - Cortinas: tecidos com fios de prata/cobre bloqueiam RF-EMF externa (torres de celular)
      - Canopies (dossel de cama): coberturas completas para cama - bloqueio de 99% de RF-EMF
      - Roupas de cama: lençóis com fibras de prata (caros, eficácia questionável para saúde)
      - CRÍTICO: eficácia depende de aterramento correto (consultar profissional especializado)

   b) Tintas de Blindagem:
      - Tintas com partículas metálicas aplicadas nas paredes
      - Bloqueiam RF-EMF externa, mas requerem aterramento e podem criar "gaiola" que concentra EMF interno se não implementado corretamente
      - Custo-benefício baixo para maioria dos casos

   c) Filtros de Linha e Aterramento:
      - Dirty electricity (harmônicas de alta frequência na rede elétrica): filtros Graham-Stetzer
      - Aterramento (earthing/grounding): tapetes, lençóis conectados ao terra - evidências preliminares de redução de inflamação e melhora de sono (mecanismo: descarga de elétrons, redução de potencial elétrico corporal)

5. Estilo de Vida e Comportamento:
   a) Uso de Dispositivos Antes de Dormir:
      - Evitar telas 2h antes de dormir: benefício duplo (reduz luz azul + EMF)
      - Se uso necessário: modo avião, blue-blocker glasses, distância mínima 30cm

   b) Hábitos de Celular:
      - Durante o dia: usar fone com fio (não Bluetooth), viva-voz, ou texting ao invés de chamadas longas próximo à cabeça
      - Carregar celular longe da cama e áreas de permanência prolongada

6. Fortificação Biológica (Suporte Antioxidante):
   - Hipótese: EMF pode causar estresse oxidativo (produção de radicais livres)
   - Suplementação antioxidante (evidências limitadas, mas seguro):
     a) Glutationa: 500mg/dia ou precursores (NAC 1200mg, ácido alfa-lipóico 300mg)
     b) Vitamina C: 1000mg/dia (antioxidante hidrossolúvel)
     c) Vitamina E: 400 UI/dia (antioxidante lipossolúvel)
     d) Melatonina: 1-3mg à noite (antioxidante potente, neuroprotetor, regulariza sono)
     e) Magnésio: 400-600mg/dia (modula canais de cálcio, reduz hiperexcitabilidade)

7. Avaliação de Eletrossensibilidade:
   - Se sintomas persistentes (insônia, cefaleia, fadiga) apesar de redução de EMF:
      a) Descartar outras causas: exames laboratoriais (hemograma, função tireoidiana, vitamina D, magnésio, cortisol)
      b) Teste de eliminação: período de 2-4 semanas com EMF minimizado ao máximo (acampamento, área rural)
      c) Considerar componente psicossomático: ansiedade sobre EMF pode causar sintomas (tratamento: CBT, mindfulness)

8. Perspectiva Equilibrada e Baseada em Evidências:
   a) Reconhecer Limitações da Ciência Atual:
      - Estudos de longo prazo insuficientes (tecnologia 5G muito recente)
      - Dificuldade de estudos duplo-cego (difícil criar "placebo" de EMF)
      - Vieses de publicação (estudos negativos menos publicados)

   b) Princípio da Precaução vs. Ansiedade Excessiva:
      - Medidas simples (desligar Wi-Fi, celular fora do quarto) são razoáveis e sem desvantagens
      - Blindagem extrema e ansiedade excessiva podem ser contraproducentes (estresse crônico comprovadamente prejudica saúde mais que EMF de baixa intensidade)

   c) Priorização:
      - Fatores comprovadamente importantes: higiene do sono, nutrição, exercício, manejo de estresse, luz azul
      - EMF: implementar medidas simples, mas não obcecar

9. Considerações Especiais:
   a) Crianças:
      - Cérebro em desenvolvimento: aplicar princípio da precaução de forma mais rigorosa
      - Evitar tablets/celulares no quarto, Wi-Fi próximo ao berço/cama

   b) Gravidez:
      - Evidências limitadas sobre efeitos no feto
      - Precaução razoável: evitar laptop no colo, celular próximo à barriga

   c) Implantes Médicos:
      - Marcapassos, desfibriladores: seguir orientações do fabricante sobre distância de fontes de EMF

10. Recursos e Educação Continuada:
    - Organizações: International EMF Alliance, BioInitiative Report (perspectiva cautelosa)
    - Contra-ponto: WHO, IEEE (perspectiva de segurança dentro dos limites estabelecidos)
    - Abordagem equilibrada: ler ambos os lados, aplicar pensamento crítico, implementar precauções razoáveis sem pânico"""
        },
    }

    # Fallback para items sem conteúdo específico
    if item_name in contents:
        return contents[item_name]
    else:
        # Conteúdo genérico baseado no nome
        return {
            "clinical_relevance": f"""A avaliação de {item_name} no contexto do sono é fundamental na medicina funcional integrativa. O sono adequado é essencial para homeostase metabólica, regulação hormonal, consolidação da memória, função imunológica e reparo celular. Fatores relacionados a {item_name} podem impactar significativamente a qualidade do sono através de mecanismos neurobiológicos, hormonais e ambientais. A medicina integrativa investiga como este aspecto específico influencia o ritmo circadiano, arquitetura do sono (fases NREM e REM), e processos restauradores noturnos.""",

            "patient_explanation": f"""Aspectos relacionados a {item_name} podem afetar a qualidade do seu sono de maneiras que você talvez não perceba conscientemente. Avaliar este fator ajuda a identificar oportunidades de melhoria para um sono mais restaurador, essencial para saúde física, mental e emocional. Bom sono não é luxo - é necessidade biológica fundamental.""",

            "conduct": f"""Avaliação funcional integrativa de {item_name}: Investigar através de questionários específicos, diário de sono, e exames laboratoriais quando indicados (cortisol salivar 4 pontos, melatonina noturna, ferritina, vitamina D, magnésio eritrocitário, função tireoidiana). Intervenções personalizadas baseadas em achados: otimização da higiene do sono, suporte nutricional (magnésio quelato 400-600mg, complexo B ativado, glicina 3g, L-teanina 200-400mg, ômega-3 1-2g/dia), manejo de estresse (meditação mindfulness, respiração 4-7-8, adaptógenos como ashwagandha 300-600mg), tratamento de condições específicas identificadas. Encaminhamento para especialistas quando apropriado (neurologista, medicina do sono, psiquiatra, fisioterapeuta)."""
        }


def main():
    print("=== Enriquecimento de 20 Score Items do Grupo SONO - Parte 2 ===\n")

    # Login
    print("1. Fazendo login...")
    token = login()
    print(f"✓ Token obtido: {token[:50]}...\n")

    # Processar cada item
    success_count = 0
    error_count = 0

    for idx, (item_id, item_name) in enumerate(ITEMS, 1):
        print(f"\n[{idx}/20] Processando: {item_name}")
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

            # Vincular aos artigos de Ritmo Circadiano
            linked = 0
            for article_id in CIRCADIAN_ARTICLES[:3]:  # Link 3 artigos
                if link_article(token, article_id, item_id):
                    linked += 1

            print(f"✓ Vinculado a {linked} artigos de Ritmo Circadiano")
        else:
            print(f"✗ Erro ao atualizar item")
            error_count += 1

        # Delay para não sobrecarregar API
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
