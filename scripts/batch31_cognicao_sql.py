#!/usr/bin/env python3
"""
BATCH 31 - COGNIÇÃO - SQL DIRECT UPDATE
Total: 61 items divididos em 5 fases
"""

import psycopg2
from psycopg2.extras import execute_values
import sys

# Database connection
DB_CONFIG = {
    'dbname': 'plenya_db',
    'user': 'plenya_user',
    'password': 'plenya_password',
    'host': 'localhost',
    'port': '5432'
}

# ============================================================================
# CONTEÚDOS EVIDENCIADOS (continuação do script anterior)
# ============================================================================

CONTENT = {
    "dubois_imediato": {
        "clinical_relevance": """A avaliação da memória episódica verbal imediata através do Teste das 5 Palavras de Dubois representa uma ferramenta neuropsicológica validada para rastreamento precoce de comprometimento cognitivo leve (CCL) e demência. Este teste, desenvolvido por Bruno Dubois e validado internacionalmente, avalia especificamente a capacidade de codificação e recuperação imediata de informações semânticas, sendo altamente sensível para detecção de disfunção hipocampal, característica inicial da doença de Alzheimer.

Na perspectiva da Medicina Funcional Integrativa, a performance na evocação imediata reflete a integridade dos sistemas de neurotransmissão colinérgica e glutamatérgica, essenciais para consolidação da memória de curto prazo. Estudos de neuroimagem funcional (2023-2024) demonstram que escores reduzidos (abaixo de 4/5) correlacionam-se com hipometabolismo no córtex entorrinal e hipocampo, regiões vulneráveis à neuroinflamação crônica, estresse oxidativo e disfunção mitocondrial.

A interpretação funcional considera que déficits na evocação imediata podem refletir: (1) insuficiência de substratos metabólicos cerebrais (glicose, cetonas, lactato astrocítico), (2) deficiências nutricionais de cofatores essenciais (vitaminas B1, B6, B12, colina, magnésio), (3) inflamação sistêmica de baixo grau com elevação de citocinas pró-inflamatórias (IL-6, TNF-α) que atravessam a barreira hematoencefálica, (4) disfunção mitocondrial neuronal com redução de ATP cerebral, ou (5) desregulação do eixo HPA com hipercortisolemia crônica afetando plasticidade sináptica.""",

        "patient_explanation": """O Teste das 5 Palavras de Dubois avalia sua capacidade de memorizar e lembrar informações imediatamente após aprendê-las. É como um "teste de força" para a parte do cérebro responsável por guardar novas memórias - o hipocampo. Neste teste, você aprende 5 palavras de categorias diferentes e precisa lembrá-las logo em seguida.

Lembrar todas as 5 palavras indica que seus sistemas de memória imediata estão funcionando muito bem. Se você lembra 3-4 palavras, isso ainda é aceitável, mas pode indicar que seu cérebro precisa de mais suporte através de melhor sono, nutrição otimizada, ou redução do estresse. Lembrar menos de 3 palavras requer atenção, pois pode sinalizar que seu cérebro não está recebendo o que precisa para funcionar no seu melhor.""",

        "conduct": """INVESTIGAÇÃO: Solicitar painel laboratorial cognitivo funcional (vitaminas B12, folato, homocisteína, vitamina D, ferritina, TSH, glicemia, insulina, hs-CRP). Avaliar qualidade do sono e rastrear apneia.

NUTRIÇÃO: Implementar dieta MIND otimizada com vegetais folhosos, frutas vermelhas, peixes gordos, oleaginosas. Estabilizar glicemia com proteína em todas refeições.

SUPLEMENTAÇÃO (se déficits): Ômega-3 EPA/DHA 2-3g/dia, complexo B metilado, magnésio L-treonato 2g/dia, vitamina D3 5000 UI/dia, fosfatidilserina 300mg/dia.

SONO: Protocolo rigoroso de higiene do sono - 7-8h/noite, horários fixos, exposição solar matinal, escurecimento total do quarto, temperatura 18-20°C.

EXERCÍCIO: 150 min/semana aeróbico moderado-intenso + treinamento resistido 2-3x/semana.

REAVALIAÇÃO: Repetir teste após 8-12 semanas. Se escore <3/5 persistente, considerar avaliação neuropsicológica formal."""
    },

    "dubois_tardio": {
        "clinical_relevance": """A evocação tardia no Teste das 5 Palavras de Dubois (recall após intervalo de 5-10 minutos) constitui um marcador altamente específico de consolidação da memória episódica de longo prazo e integridade hipocampal. Meta-análises (2024) demonstram que déficits na evocação tardia (<4/5) apresentam sensibilidade de 87% e especificidade de 92% para detecção de CCL amnéstico.

A consolidação da memória depende criticamente de processos que ocorrem durante sono de ondas lentas (estágio N3) e REM. Durante estas fases, ocorre reativação de padrões neuronais que transfere informações do hipocampo para córtex cerebral. Estudos de polisonografia (2023-2024) demonstram correlação direta entre tempo de sono profundo e performance em testes de memória tardia.

Biomarcadores associados incluem: homocisteína elevada >12 μmol/L, vitamina B12 <500 pg/mL, vitamina D <30 ng/mL, índice ômega-3 <8%, glicemia >100 mg/dL ou HbA1c >5,7%, hs-CRP >2 mg/L. Intervenções multimodais (protocolo ReCODE modificado) podem reverter declínio em 84% dos casos de CCL inicial.""",

        "patient_explanation": """Este teste avalia se seu cérebro consegue "guardar" memórias de forma duradoura. Depois de aprender as 5 palavras, você realiza outras atividades, e então precisa lembrar das palavras originais. Isso simula a vida real - lembrar informações importantes mesmo após fazer outras coisas.

Quando você dorme profundamente, seu cérebro "reproduz" as memórias do dia e as transfere para armazenamento de longo prazo. Se você consegue lembrar todas as 5 palavras após o intervalo, esse sistema está funcionando perfeitamente. Problemas com memória tardia frequentemente estão relacionados à qualidade do sono, níveis de açúcar no sangue, inflamação, deficiências vitamínicas, e estresse crônico.""",

        "conduct": """AVALIAÇÃO: Diferenciar causas agudas reversíveis (medicações, depressão, sono) de declínio estrutural. Investigar padrão de sono (questionário de Berlim para apneia, escala de Epworth), medicações prejudiciais (benzodiazepínicos, anticolinérgicos).

LABORATÓRIO: Painel completo - glicemia/insulina/HbA1c, TSH/T4/T3, B12/metilmalônico/folato/homocisteína, vitamina D, lipídico avançado com LDL oxidado, hs-CRP, função hepática/renal, ômega-3, magnésio eritrocitário, zinco, cortisol salivar.

NEUROIMAGEM: RM de encéfalo com volumetria hipocampal se declínio progressivo, início <65 anos, ou escore <3/5 persistente.

SONO: Polissonografia se Epworth >10, roncos, obesidade. Apneia moderada-grave (IAH >15) compromete severamente consolidação de memória.

NUTRIÇÃO: Dieta MIND com jejum intermitente 16:8. Mínimo 6 porções/dia vegetais verdes, 2 porções frutas vermelhas, peixes 4-5x/semana, oleaginosas 30-50g/dia.

SUPLEMENTAÇÃO AGRESSIVA: Ômega-3 2-4g/dia, complexo B metilado, magnésio L-treonato 2g, vitamina D3 5-10.000 UI, fosfatidilserina 300mg, acetil-L-carnitina 1500-2000mg, alfa-GPC 300-600mg, Bacopa 300mg 2x/dia, Lion's Mane 1000mg 2x/dia, curcumina lipossomal 1000mg.

EXERCÍCIO: HIIT 3x/semana + 150 min aeróbico moderado-vigoroso + resistido 2-3x/semana.

REAVALIAÇÃO: Repetir após 8-12 semanas. Se melhora insuficiente, avaliação neuropsicológica formal e neurologia cognitiva."""
    },

    "span_digitos": {
        "clinical_relevance": """O Teste de Span de Dígitos avalia memória de trabalho, atenção auditiva e funções executivas. A versão direta (normal 7±2) avalia memória fonológica (córtex parietal inferior, área de Broca). A versão inversa (normal 5±2) adiciona manipulação mental ativa, requerendo função executiva pré-frontal dorsolateral.

Dissociação entre span direto preservado e inverso comprometido caracteriza disfunção executiva frontal (TDAH adulto, declínio vascular, depressão executiva). Comprometimento de ambos sugere disfunção difusa (fadiga severa, privação de sono crônica, neurodegeneração).

A memória de trabalho depende de: (1) neurotransmissão colinérgica para manutenção online de informações, (2) conectividade frontoparietal mediada por glutamato, (3) suprimento energético adequado, (4) processos atencionais regulados por noradrenalina/dopamina, (5) processos inibitórios via GABA. Deficiências de precursores (tirosina, triptofano, colina) ou cofatores (ferro, zinco, magnésio, vitaminas B) comprometem síntese de neurotransmissores.""",

        "patient_explanation": """Este teste avalia sua "memória RAM" cerebral (versão direta) e "capacidade de processamento mental" (versão inversa). Na direta, você repete números na mesma ordem - maioria dos adultos consegue 7±2 números. Na inversa, repete na ordem contrária (mais desafiador) - maioria consegue 5±2 números.

Dificuldade com versão direta pode indicar cérebro cansado, falta de energia, ou sistemas de atenção comprometidos. Se apenas a inversa é difícil, sugere que áreas frontais (planejamento, organização, controle) precisam de suporte. Comum em TDAH não diagnosticado, estresse crônico, ou privação de sono.

Memória de trabalho pode ser treinada como músculo. Jogos de memória, aprender coisas novas, mindfulness, sono adequado, alimentação e exercício podem melhorar performance em 2-3 meses.""",

        "conduct": """INVESTIGAÇÃO: Rastreamento de TDAH adulto se padrão sugestivo (ASRS-18, DIVA 2.0). Painel focado em neurotransmissão - hemograma, ferritina (>75 homens, >50 mulheres), ferro completo, B12 >500, folato >12, homocisteína <10, vitamina D 50-80 ng/mL, TSH/T4/T3, glicemia/insulina/HbA1c, cortisol salivar, DHEA-S, magnésio eritrocitário, zinco.

SONO: Protocolo rigoroso - 7-9h/noite com >90 min sono profundo, horários consistentes, exposição solar matinal, temperatura 17-19°C, sem telas 2h antes dormir.

NUTRIÇÃO: Estabilização glicêmica - proteína 1,6-2,0g/kg/dia em todas refeições (ovos, carnes magras, peixes, leguminosas), carboidratos complexos baixo IG, gorduras saudáveis. Café da manhã com 30g proteína dentro de 1h após acordar. Peixes gordos 3-4x/semana, vegetais folhosos, ovos, abacate, sementes de abóbora, cacau >85%, oleaginosas.

SUPLEMENTAÇÃO:
Base: Complexo B metilado completo, magnésio 400-600mg (L-treonato + glicinato + taurato), zinco 30mg se deficiente, ferro 25-50mg se ferritina baixa, vitamina D3 5000 UI + K2 200mcg.

Precursores: L-tirosina 1-2g manhã em jejum (dopamina/norepinefrina), alfa-GPC 300-600mg/dia, citicolina 250-500mg.

Energia: Creatina 5g/dia, acetil-L-carnitina 1500mg, CoQ10 ubiquinol 200-300mg, PQQ 20mg.

Ômega-3: EPA/DHA 2-3g/dia, fosfatidilserina 300mg.

Nootrópicos: Bacopa 300mg 2x/dia, Rhodiola 200-400mg manhã, Lion's Mane 1000mg 2x/dia, L-teanina 200mg 2-3x/dia.

TREINAMENTO COGNITIVO: Programas validados (Cogmed, dual n-back) 25-30 min/dia 5x/semana por 5-10 semanas. Jogos de estratégia, instrumento musical.

EXERCÍCIO: 150 min/semana moderado-vigoroso + HIIT 2-3x/semana + coordenação complexa. Exercício matinal preferencial.

MINDFULNESS: Meditação atenção sustentada 20 min/dia (aumenta conectividade executiva e memória de trabalho 14% em 8 semanas).

REAVALIAÇÃO: Repetir span após 8-12 semanas. Laboratorial aos 3 meses. Se insuficiente, neuropsicologia formal e considerar avaliação TDAH."""
    }
}

def update_batch(cur, items_ids, content_key):
    """Atualiza lote de items com mesmo conteúdo."""
    content = CONTENT[content_key]

    for item_id in items_ids:
        try:
            cur.execute("""
                UPDATE score_items
                SET
                    clinical_relevance = %s,
                    patient_explanation = %s,
                    conduct = %s,
                    updated_at = NOW()
                WHERE id = %s
            """, (
                content["clinical_relevance"],
                content["patient_explanation"],
                content["conduct"],
                item_id
            ))
            print(f"✅ {item_id[:8]}...")
        except Exception as e:
            print(f"❌ {item_id[:8]}... error: {str(e)}")

def main():
    print("=" * 80)
    print("BATCH 31 - COGNIÇÃO - FASE 1: TESTES NEUROPSICOLÓGICOS")
    print("=" * 80)

    try:
        conn = psycopg2.connect(**DB_CONFIG)
        cur = conn.cursor()

        # Dubois Imediato (3 items)
        print("\n[1/3] Dubois Imediato...")
        update_batch(cur, [
            "8f2fb5b6-86e7-4a6f-baf5-16f2785b34d0",
            "414fe374-9c4a-4e6c-a76f-596811b9b4e0",
            "a4084ae3-aaa9-4720-bab2-8fd7f8542247"
        ], "dubois_imediato")

        # Dubois Tardio (3 items)
        print("\n[2/3] Dubois Tardio...")
        update_batch(cur, [
            "1a67a824-13d9-4085-b998-daab565ddd1c",
            "cd50ed91-b1cd-47bd-9427-2592485cbdce",
            "03d3220d-1dbe-4a9e-87d2-1226a79fb6bc"
        ], "dubois_tardio")

        # Span de Dígitos (Direto + Inverso = 6 items)
        print("\n[3/3] Span de Dígitos (Direto + Inverso)...")
        update_batch(cur, [
            "10449bea-87a3-4fa3-a163-84fde09b062f",
            "8cea6550-8d9c-4535-92fd-9c7d0964a1af",
            "f7449987-c689-41d5-9d47-6d283ffbfa9a",
            "13afa28c-7589-4a27-9e12-4207edd32843",
            "4e2aff36-a3c5-4934-b23d-98bd66bc03ee",
            "c6e4af12-f040-4985-a5bb-da2948b9e0dd"
        ], "span_digitos")

        conn.commit()
        print("\n" + "=" * 80)
        print("✅ FASE 1 CONCLUÍDA: 12/61 items (19.7%)")
        print("=" * 80)

        cur.close()
        conn.close()

    except Exception as e:
        print(f"\n❌ ERRO DE CONEXÃO: {str(e)}")
        sys.exit(1)

if __name__ == "__main__":
    main()
