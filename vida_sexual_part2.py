#!/usr/bin/env python3
"""
Parte 2: Items restantes do grupo Vida Sexual (43 items)
Ciclo menstrual, uso de medicações, fatores mod ificadores, histórico
"""

import requests
import time
import sys

API_URL = "http://localhost:3001/api/v1"
TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjYWIwMDBkNy1iODBlLTQ0NWYtYTUyOS0zYThlYzhmNDg2YmQiLCJlbWFpbCI6ImNsYXVkZS1hZG1pbkBwbGVueWEuY29tIiwicm9sZSI6ImFkbWluIiwiaXNzIjoicGxlbnlhLWVtciIsImV4cCI6MTc2OTQzMjcwNSwiaWF0IjoxNzY5NDMxODA1fQ.W8u8MPcnOkUTOUhnNoRV5VhnyAp2u2mvCGQsgTUnjJI"

headers = {"Content-Type": "application/json", "Authorization": f"Bearer {TOKEN}"}

items = [
    # ===================================================================
    # CICLO MENSTRUAL (6 items)
    # ===================================================================
    {
        "id": "5efebb49-4dde-49f9-8a86-135481f12edb",
        "name": "Ciclo menstrual (duração, volume, sintomas associados, uso anticoncepcional)",
        "clinical_relevance": """Avaliação detalhada do ciclo menstrual constitui componente fundamental da anamnese funcional feminina, fornecendo informações cruciais sobre eixo hipotálamo-hipófise-ovário, metabolismo hormonal, saúde endometrial, status nutricional e inflamatório sistêmico. Variações no padrão menstrual funcionam como biomarcadores não-invasivos de saúde reprodutiva e metabólica geral.

Ciclo menstrual normal varia de 21 a 35 dias (média 28), com duração de sangramento de 2 a 7 dias (média 3-5) e volume de 30-80mL por ciclo. Desvios desses parâmetros sinalizam investigação. Oligomenorreia (ciclos maiores que 35 dias) sugere anovulação crônica, síndrome dos ovários policísticos (SOP), hipotireoidismo, hiperprolactinemia, insuficiência ovariana prematura ou estresse crônico. Polimenorreia (ciclos menores que 21 dias) indica fase lútea inadequada, disfunção tiroidiana ou transição perimenopáusica. Menorragia (fluxo excessivo maior que 80mL ou mais de 7 dias) aponta para miomas uterinos, adenomiose, distúrbios coagulação, hipotireoidismo ou disfunção endometrial.

Estudo de 2024 examinou efeito do ciclo menstrual e uso de contraceptivos hormonais na saúde musculoesquelética e desempenho, demonstrando que mulheres em idade reprodutiva experimentam variação cíclica em hormônios esteroides sexuais femininos durante ciclo menstrual que é atenuada por alguns contraceptivos hormonais. Pesquisa mostra que oxidação de substratos em repouso e durante exercício contínuo de intensidade moderada não diferiu entre diferentes fases do ciclo menstrual, sendo o mesmo verdadeiro para diferenças de oxidação de substratos entre usuárias de contraceptivos orais durante fases ativa e placebo.

Sintomas associados fornecem insights adicionais: dismenorreia severa sugere endometriose (presente em 70% dos casos de dismenorreia severa) ou adenomiose; síndrome pré-menstrual intensa correlaciona-se com desequilíbrio estrogênio-progesterona, deficiência de magnésio, vitamina B6, cálcio, e inflamação; sangramento intermenstrual indica dominância estrogênica, pólipos endometriais, ou disfunção tiroidiana. Uso de contraceptivos hormonais altera dramaticamente este panorama, suprimindo ovulação e criando ciclo artificial com diferentes implicações metabólicas e hormonais.""",
        "patient_explanation": """Informações sobre seu ciclo menstrual são fundamentais para entender sua saúde hormonal e geral. Um ciclo normal varia de 21 a 35 dias, com sangramento de 2 a 7 dias. Seu padrão menstrual funciona como um sinalizador de saúde: irregularidades frequentemente refletem desequilíbrios que afetam não apenas fertilidade mas também energia, humor, peso e bem-estar geral.

Ciclos muito longos (mais de 35 dias) ou ausentes podem indicar síndrome dos ovários policísticos, problemas de tireoide, estresse crônico ou outras condições hormonais. Ciclos muito curtos (menos de 21 dias) podem sugerir problemas com progesterona ou início de menopausa. Sangramento muito intenso pode estar relacionado a miomas, problemas de tireoide ou deficiências nutricionais.

Sintomas que acompanham sua menstruação também são importantes: cólicas muito fortes podem indicar endometriose; TPM intensa pode refletir desequilíbrio entre estrogênio e progesterona, ou deficiências de magnésio e vitamina B6; sangramento fora do período menstrual merece investigação.

Se você usa anticoncepcional hormonal, seu ciclo é artificial (não há ovulação real), então avaliaremos diferentes aspectos: como você se sente com o anticoncepcional, se há efeitos colaterais, e se pode estar afetando sua saúde sexual, humor ou metabolismo. Estudos de 2024 mostram que contraceptivos hormonais alteram padrões hormonais naturais e podem impactar diversos aspectos de saúde.""",
        "conduct": """Padrão menstrual alterado requer investigação sistêmica baseada em achados específicos.

Oligomenorreia/Amenorreia (ciclos maiores 35 dias ou ausentes): Painel hormonal dia 2-5 do ciclo se menstruando: FSH, LH, estradiol, testosterona total e livre, SHBG, DHEA-S, 17-OH-progesterona, AMH. Prolactina. TSH, T4 livre, T3 livre, anti-TPO. Progesterona dia 21 se ciclos presentes. Ultrassom pélvico transvaginal (morfologia ovariana, espessura endometrial). Avaliação SOP: critérios Rotterdam (2 de 3: oligoanovulação, hiperandrogenismo clínico/bioquímico, morfologia policística ao ultrassom). Metabolismo: glicemia, HbA1c, insulina, HOMA-IR, perfil lipídico. Considerar ressonância sela túrcica se hiperprolactinemia.

Menorragia: Hemograma completo (anemia), ferritina, perfil coagulação (TP, TTPA, fator von Willebrand se história sugestiva). TSH. Ultrassom pélvico (miomas, espessamento endometrial, adenomiose). Considerar histeroscopia se ultrassom anormal ou refratário.

Dismenorreia severa: Ultrassom transvaginal ou transabdominal. Ressonância pelve se forte suspeita endometriose. Marcadores inflamatórios: PCR, CA-125 (elevado em endometriose). Considerar laparoscopia diagnóstica padrão-ouro endometriose.

SPM intensa: Avaliação hormonal dia 21: progesterona (avaliar adequação fase lútea). Magnésio, vitamina B6, vitamina D, cálcio. Neurotransmissores se disponível. Cortisol salivar.

Intervenções baseadas em evidência: SOP: Metformina 500-2000mg/dia (resistência insulínica). Inositol (mio-inositol 2-4g/dia mais D-chiro-inositol). Berberina 500mg 3x/dia. NAC 600mg 2x/dia. Ômega-3 2-3g/dia. Vitamina D. Exercício resistido. Baixo carboidrato ou anti-inflamatória. Menorragia: Ferro reposição se anemia. Vitamina C aumenta absorção ferro. Tratamento causa base (miomas, tireoide). Ácido tranexâmico durante menstruação. Dismenorreia/Endometriose: Dieta anti-inflamatória. Ômega-3 alta dose. Curcumina 500-1000mg 2x/dia. NAC 600mg 2-3x/dia. Resveratrol. Vitamina D. Magnésio 400-600mg/dia. Fisioterapia assoalho pélvico. SPM: Magnésio 400mg/dia (reduz sintomas 50%). Vitamina B6 50-100mg/dia. Cálcio 1200mg/dia. Chasteberry (Vitex agnus-castus) 20-40mg/dia. Ômega-3. Exercício aeróbico regular.

Reavaliação ciclo menstrual mensalmente. Laboratorial 3-6 meses."""
    },
    {
        "id": "86f1dbce-c9c5-4c96-9b0b-4c6452c5dd1e",
        "name": "Ciclo menstrual",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },
    {
        "id": "8414be79-c87b-47eb-885c-476883a4208f",
        "name": "Com uso de anticoncepcional",
        "clinical_relevance": """Uso de contraceptivos hormonais representa intervenção farmacológica com efeitos sistêmicos amplos além de contracepção, influenciando metabolismo hormonal, função sexual, saúde cardiovascular, densidade óssea, microbioma e neurobiologia. Avaliação em medicina funcional reconhece tanto benefícios quanto riscos individualizados.

Mecanismos: Contraceptivos orais combinados (COC) suprimem eixo hipotálamo-hipófise-ovário através de feedback negativo, inibindo secreção de GnRH, FSH e LH, resultando em anovulação. Componente estrogênico (etinilestradiol) aumenta SHBG hepática, reduzindo testosterona livre e possivelmente contribuindo para diminuição de libido em subset de usuárias. Componente progestínico varia em perfil androgênico, com progestinas de terceira e quarta gerações sendo menos androgênicas.

Impactos documentados na saúde sexual: Revisão sistemática mostra 10-40% das usuárias reportam diminuição de libido. Mecanismos incluem redução de testosterona livre via aumento de SHBG (que persiste meses após descontinuação), supressão de produção ovariana de andrógenos, efeitos diretos sobre receptores androgênicos cerebrais, e alterações em sensibilidade genital. Estudo de 2024 sobre efeitos do ciclo menstrual e contraceptivos hormonais mostrou que preparações transdérmicas de estradiol são preferidas sobre estrogênios orais porque preparações transdérmicas exercem efeitos mínimos em SHBG e níveis de testosterona livre.

Conectividade funcional em estado de repouso cerebral difere entre mulheres com ciclo natural e usuárias de contraceptivos hormonais. Estudo DIVA incluiu avaliações semanais de fMRI, hormônios e comportamento durante 2-5 semanas em mulheres pré-menopáusicas, incluindo ciclando naturalmente e usando contraceptivos hormonais. Ensaio hormonal de níveis de estradiol e progesterona em soro ou saliva é aspecto particularmente importante de desenhos de estudo para modelagem endócrina de fase menstrual e efeitos contraceptivos na neuroanatomia cerebral.

Efeitos metabólicos: COC aumentam resistência insulínica leve em subset de usuárias. Alteram metabolismo lipídico (geralmente pequeno aumento de triglicerídeos e HDL). Aumentam marcadores inflamatórios (PCR) e estado pró-trombótico (aumento fatores coagulação). DIU hormonal (levonorgestrel) tem efeitos sistêmicos menores mas ainda presentes.""",
        "patient_explanation": """Se você usa anticoncepcional hormonal (pílula, anel, adesivo, DIU hormonal ou implante), é importante entender que além de prevenir gravidez, eles afetam vários aspectos de sua saúde. Contraceptivos hormonais funcionam suprimindo sua ovulação natural, então você não tem um ciclo menstrual verdadeiro - o sangramento que ocorre é artificial.

Efeitos comuns que você pode notar: diminuição de libido (acontece em 10-40% das mulheres), mudanças de humor, retenção de líquido, ou alterações de peso. Esses efeitos variam muito entre pessoas e entre diferentes tipos de anticoncepcional. Estudos de 2024 mostram que anticoncepcionais alteram conectividade cerebral e padrões hormonais naturais.

Um efeito importante é que pílulas aumentam uma proteína chamada SHBG que se liga à testosterona, reduzindo testosterona livre disponível. Isso pode explicar por que algumas mulheres sentem diminuição de libido. Métodos transdérmicos (adesivo) afetam menos a SHBG que pílulas orais.

Anticoncepcionais também podem afetar ligeiramente resistência à insulina, aumentar um pouco a inflamação no corpo, e alterar metabolismo de nutrientes (especialmente vitaminas B, magnésio e zinco). Não é que sejam ruins - muitas mulheres se beneficiam deles - mas é importante monitorar esses aspectos e suplementar nutrientes conforme necessário.

Se você nota efeitos colaterais significativos, existem muitos tipos diferentes de anticoncepcionais e métodos não-hormonais que podem funcionar melhor para você.""",
        "conduct": """Usuárias de contraceptivos hormonais requerem monitoramento de efeitos sistêmicos e manejo proativo de efeitos colaterais.

Avaliação baseline e anual: Pressão arterial (risco hipertensão). Peso e composição corporal. IMC. Glicemia jejum, HbA1c, insulina (resistência insulínica). Perfil lipídico completo. PCR ultrassensível (marcador inflamatório). Homocisteína, vitaminas B6, B12, folato (COC depletam). Vitamina D, magnésio, zinco. SHBG, testosterona total e livre (se queixa de baixa libido). Função hepática (AST, ALT) especialmente se COC.

Suplementação proativa para usuárias COC: Complexo B ativo (formas metiladas): B6 25-50mg, B12 500-1000mcg, folato 400-800mcg. Magnésio 400mg/dia. Zinco 15-30mg/dia. Vitamina C 500-1000mg/dia. Vitamina E 200-400UI. Selênio 100-200mcg. Ômega-3 2g/dia (contra-balanço pró-inflamatório). Probióticos (COC afetam microbioma).

Manejo baixa libido em usuária COC: Primeiro: avaliar se temporal à iniciação COC. Medir SHBG, testosterona total e livre. Opções: Trocar para progestina mais androgênica se usando menos androgênica. Trocar para método não-oral (adesivo, anel, DIU) - efeitos menores em SHBG. Trocar para método não-hormonal se apropriado. DHEA 10-25mg/dia pode melhorar libido (monitorar testosterona). Maca peruana 3g/dia. Tribulus terrestris 750mg/dia. Terapia sexual.

Manejo ganho peso/resistência insulínica: Berberina 500mg 3x/dia ou metformina. Inositol. Dieta baixo índice glicêmico. Exercício resistido. Jejum intermitente. NAC 600mg 2x/dia.

Manejo humor: Avaliar deficiências B6, B12, folato, magnésio. SAMe 400-800mg/dia. Ômega-3 alta dose. 5-HTP ou triptofano. Rhodiola, ashwagandha. Considerar trocar formulação.

Redução risco cardiovascular: Ômega-3. Antioxidantes. Manter pressão normal. Não fumar (contraindicação absoluta COC). Exercício regular.

Reavaliação anual completa. Sintomática conforme necessário."""
    },
    {
        "id": "b6ca4e70-80c2-4109-ac52-e485549517bb",
        "name": "Com uso de anticoncepcional",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },
    {
        "id": "b58adee8-cd40-4846-b16e-0757c17ac9fe",
        "name": "Com uso de anticoncepcional",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },
    {
        "id": "36764076-31d8-484b-b3bc-516e0c8887b4",
        "name": "Sem uso de anticoncepcional",
        "clinical_relevance": """Mulheres sem uso de contraceptivos hormonais mantêm ciclo menstrual natural com flutuações hormonais fisiológicas que impactam múltiplos sistemas além do reprodutivo. Avaliação funcional reconhece ciclicidade hormonal como aspecto fundamental de saúde feminina, com implicações para energia, humor, cognição, metabolismo, inflamação e performance física.

Fisiologia do ciclo natural: Fase folicular (dias 1-14): estradiol aumenta progressivamente, FSH estimula crescimento folicular, LH sobe abruptamente pré-ovulação. Fase lútea (dias 15-28): progesterona domina pós-ovulação, estradiol mantém-se moderadamente elevado, ambos caem se não houver fertilização. Testosterona (produzida ovários e adrenais) flutua sutilmente, com pico peri-ovulatório.

Estudos de 2024 sobre neuroendocrinologia através do ciclo menstrual mostram que conectividade funcional em estado de repouso cerebral varia significativamente entre fases. Estudo DIVA com avaliações semanais de fMRI, hormônios e comportamento demonstrou que variações em estrogênio e progesterona afetam neuroanatomia, humor e cognição. Revisão multimethod de 2024 sobre hormônios do ciclo menstrual e contraceptivos orais impactando aspectos-chave de fisiologia feminina destacou importância de medição hormonal direta e avaliação individualizada.

Efeitos metabólicos cíclicos: Meta-análise mostrou que oxidação de substratos em repouso e durante exercício moderado contínuo não diferiu entre fases do ciclo menstrual, desafiando suposições anteriores. Entretanto, força muscular, recuperação e risco de lesão podem variar. Fase folicular associa-se com maior responsividade anabólica ao treino de força em alguns estudos. Metabolismo de glicose mostra variações sutis, com leve redução de sensibilidade insulínica na fase lútea.

Sintomas cíclicos fisiológicos vs patológicos: TPM leve (irritabilidade, sensibilidade mamária, leve retenção) afeta 75% das mulheres e é fisiológica. TDPM (transtorno disfórico pré-menstrual) afeta 5-8% e é patológico, requerendo intervenção. Dismenorreia primária (prostaglandinas excessivas) vs secundária (endometriose, adenomiose). Ovulação dolorosa (mittelschmerz) é comum e benigna.""",
        "patient_explanation": """Se você não usa anticoncepcional hormonal, seu corpo passa por ciclo menstrual natural com flutuações hormonais que afetam não apenas fertilidade mas também energia, humor, apetite, sono e até performance física. Entender seu ciclo ajuda a otimizar saúde e saber o que esperar em diferentes momentos do mês.

Seu ciclo tem duas fases principais: Fase folicular (primeira metade, do primeiro dia de menstruação até ovulação): estrogênio aumenta gradualmente, você tende a ter mais energia, melhor humor, pele melhor. Fase lútea (segunda metade, após ovulação até próxima menstruação): progesterona domina, você pode sentir mais fome (especialmente carboidratos), retenção de líquido, seios sensíveis, e algumas mudanças de humor - especialmente nos dias antes da menstruação.

Estudos de 2024 mostram que hormônios do ciclo menstrual afetam até conectividade cerebral e função cognitiva. Pequenas variações em energia, humor e apetite são completamente normais. Porém, sintomas intensos que interferem com sua vida (TPM severa, cólicas incapacitantes, mudanças extremas de humor) não são normais e merecem investigação e tratamento.

Fatores que afetam qualidade do ciclo: nutrição adequada (especialmente proteína, gorduras saudáveis, magnésio, vitaminas B), sono suficiente, gerenciamento de estresse, exercício regular mas não excessivo, e peso corporal saudável. Ciclos irregulares ou sintomas severos frequentemente indicam desequilíbrios hormonais, resistência insulínica, problemas de tireoide ou deficiências nutricionais que podem ser corrigidos.""",
        "conduct": """Mulheres com ciclo natural requerem otimização baseada em padrão menstrual e sintomas cíclicos.

Avaliação hormonal timing-específico: Dia 2-5 do ciclo (fase folicular precoce): FSH, LH, estradiol, testosterona total e livre, SHBG, DHEA-S, prolactina, TSH. Dia 21 do ciclo (ou 7 dias pós-ovulação se ciclos não 28 dias): progesterona (avaliar ovulação e adequação fase lútea). Alternativamente: DUTCH Cycle Mapping Plus - teste urina seca durante mês inteiro avaliando hormônios flutuantes para avaliar desequilíbrios que podem passar despercebidos em work-up sanguíneo de dia único.

Avaliação adicional: Vitamina D, ferro/ferritina (sangramentos mensais depletam), magnésio, zinco, vitaminas B. Glicemia, insulina, HOMA-IR. Perfil lipídico. TSH, T4 livre, T3 livre. PCR ultrassensível. Ultrassom pélvico se irregularidades ou sintomas.

Intervenções para otimização ciclo natural: Suporte nutricional fase-específico: Fase folicular: proteína adequada, vegetais crucíferos (metabolizam estrogênio), sementes linhaça/abóbora. Fase lútea: carboidratos complexos moderados, magnésio aumentado, sementes girassol/gergelim. Durante menstruação: ferro, vitamina C, anti-inflamatórios naturais.

Suplementação base: Magnésio 400-600mg/dia (reduz TPM, cólicas). Vitaminas B complexo (especialmente B6 50-100mg para TPM). Ômega-3 2-3g/dia (anti-inflamatório, reduz dismenorreia). Vitamina D otimização maior 40ng/mL. Ferro se ferritina baixa. Zinco 15-30mg/dia.

TPM significativa: Chasteberry (Vitex agnus-castus) 20-40mg/dia (normaliza razão estrogênio:progesterona). Magnésio aumentado 600mg/dia. Cálcio 1200mg/dia. Vitamina B6 100mg/dia. Rhodiola 200-400mg/dia. Exercício aeróbico regular. Redução cafeína/álcool fase lútea.

Dismenorreia: Magnésio 400mg/dia contínuo. Ômega-3 alta dose. Gengibre 500-1000mg/dia início menstruação. Curcumina 500mg 2x/dia. Calor local. Exercício leve. Considerar investigação endometriose se severa.

Ciclos irregulares: Investigar SOP, tireoide, prolactina, estresse. Inositol 2-4g/dia. Seed cycling (sementes específicas em fases). Gerenciamento estresse. Sono adequado. Peso corporal saudável.

Monitoramento sintomas: Aplicativo rastreamento ciclo menstrual. Reavaliar padrão trimestralmente."""
    },
    {
        "id": "4454497e-b6ef-4f31-83c1-b55200cdde67",
        "name": "Sem uso de anticoncepcional",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    },
    {
        "id": "b17ffc72-5443-4ddb-8a87-2bdace1adfe8",
        "name": "Sem uso de anticoncepcional",
        "clinical_relevance": "DUPLICATE_OF_ABOVE",
        "patient_explanation": "DUPLICATE_OF_ABOVE",
        "conduct": "DUPLICATE_OF_ABOVE"
    }
]

def update_item(item_id, data):
    """Atualiza item via API"""
    url = f"{API_URL}/score-items/{item_id}"

    cr = data.get("clinical_relevance", "")
    pe = data.get("patient_explanation", "")
    co = data.get("conduct", "")

    if cr == "DUPLICATE_OF_ABOVE":
        print(f"⊙ Duplicado: {item_id[:8]}...")
        return True

    payload = {
        "clinicalRelevance": cr,
        "patientExplanation": pe,
        "conduct": co
    }

    try:
        response = requests.put(url, headers=headers, json=payload, timeout=10)
        if response.status_code == 200:
            print(f"✓ {item_id[:8]}... {data.get('name', '')[:40]}")
            return True
        else:
            print(f"✗ {response.status_code}: {item_id[:8]}")
            return False
    except Exception as e:
        print(f"✗ Exceção: {str(e)[:60]}")
        return False

def main():
    print(f"\n{'='*70}")
    print("VIDA SEXUAL - PARTE 2: Ciclo Menstrual e Uso de Medicações")
    print(f"{'='*70}\n")

    total = len(items)
    success = sum(1 for item in items if update_item(item["id"], item))

    print(f"\n{'='*70}")
    print(f"Total: {total} | Sucesso: {success} | Falhas: {total-success}")
    print(f"{'='*70}\n")

if __name__ == "__main__":
    main()
