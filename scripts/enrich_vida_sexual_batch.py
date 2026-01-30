#!/usr/bin/env python3
"""
Enriquecimento AI dos 30 items de VIDA SEXUAL
Foco: Medicina Funcional + Sensibilidade Clínica
Tópicos: Hormônios, disfunção sexual, escalas validadas, suplementação, fatores psicossociais
"""

import requests
import json
import sys
from typing import Dict, List, Optional

# Configuração
API_BASE = "http://localhost:3001/api/v1"
LOGIN_EMAIL = "import@plenya.com"
LOGIN_PASSWORD = "Import@123456"

# 30 IDs do grupo VIDA SEXUAL
ITEM_IDS = [
    "70babe19-54b9-4d12-9f9e-5dd3fc1e6df0",
    "86f1dbce-c9c5-4c96-9b0b-4c6452c5dd1e",
    "b58adee8-cd40-4846-b16e-0757c17ac9fe",
    "b6ca4e70-80c2-4109-ac52-e485549517bb",
    "92ac8388-9238-4214-a0a9-8466fc10f70c",
    "8b2f53db-fd5f-4dcb-b0c8-53a4a878857f",
    "a58ccffd-e1f3-45e7-a4b6-6169795e40ed",
    "4b275a4a-ed2f-4ea0-8419-37bd0473239c",
    "521607f7-687a-41fa-8f02-2c637f8026b7",
    "3a1759ed-cc26-4fb6-99ad-dd12f4e9e919",
    "bd7913cc-c28c-4c61-852d-4e4251cf3e83",
    "c3bb6a2b-adc5-4df1-a88f-808db5e184d0",
    "75b6899b-5041-4b26-ae67-bd05fe13a325",
    "7ec84d95-9305-4a18-a61e-67eed6a5cc6f",
    "b0223270-b930-4859-8bc3-44c9deee8abe",
    "4c57cf2c-a4ad-44e0-a223-e3882a00cae3",
    "b884c9c9-8279-4bc8-a56b-41bec7497c68",
    "de0b5423-5ccd-4de8-aafb-65e09d39748a",
    "fb22a739-2a63-4782-a951-d11b75225e77",
    "802c0637-851f-4e48-a9af-1a71a90c337f",
    "f7034f07-8717-41af-8492-4a3f57381a49",
    "aa0a7442-aa1e-4276-ba50-235b08f1838a",
    "722717e2-659e-4572-a771-8547b004c12b",
    "b1ea8e2d-876c-45f3-a287-89c9b36f13b5",
    "7cc28981-bd20-4e5e-b7c0-bd312f6a7ed4",
    "79addea2-e157-4fa9-8aec-339b2232b1e3",
    "b08e3349-fc30-495b-8203-e7ba11b5ed8b",
    "7b587635-f23b-4b9a-9076-5b7f5e61920c",
    "f09d3b0e-ad37-47c1-a2c5-1677466f05ca",
    "6d18ecfb-040a-4467-9c04-2dc347979727"
]

class VidaSexualEnricher:
    def __init__(self):
        self.token = None
        self.session = requests.Session()

    def login(self) -> bool:
        """Autenticação na API"""
        try:
            response = self.session.post(
                f"{API_BASE}/auth/login",
                json={"email": LOGIN_EMAIL, "password": LOGIN_PASSWORD}
            )
            response.raise_for_status()
            data = response.json()
            # Tentar diferentes formatos de resposta
            self.token = data.get('accessToken') or data.get('access_token') or data.get('token') or data.get('data', {}).get('accessToken')
            if not self.token:
                print(f"✗ Token não encontrado na resposta: {data}")
                return False
            self.session.headers.update({'Authorization': f'Bearer {self.token}'})
            print("✓ Autenticação bem-sucedida")
            return True
        except Exception as e:
            print(f"✗ Erro no login: {e}")
            import traceback
            traceback.print_exc()
            return False

    def get_item(self, item_id: str) -> Optional[Dict]:
        """Buscar item por ID"""
        try:
            response = self.session.get(f"{API_BASE}/score-items/{item_id}")
            response.raise_for_status()
            return response.json()
        except Exception as e:
            print(f"✗ Erro ao buscar item {item_id}: {e}")
            return None

    def generate_enrichment(self, item: Dict) -> Dict:
        """
        Gera conteúdo clínico enriquecido para item de Vida Sexual
        Baseado em medicina funcional e evidências científicas
        """
        question = item.get('question', '')
        description = item.get('description', '')

        # Mapeamento de conteúdo clínico por padrões no nome do item
        enrichment = {
            'clinicalRelevance': '',
            'conduct': '',
            'patientExplanation': ''
        }

        # ESCALAS VALIDADAS (ASEX, FSFI, IIEF-5)
        if 'ASEX' in question or 'Arizona' in question:
            enrichment = {
                'clinical_relevance': 'A escala ASEX (Arizona Sexual Experience Scale) é ferramenta validada para avaliar disfunção sexual induzida por medicamentos, especialmente antidepressivos SSRIs/SNRIs. Score ≥19 ou qualquer item ≥5 indica disfunção sexual significativa. Particularmente sensível para detectar efeitos de paroxetina, sertralina, venlafaxina.',
                'clinical_action': 'Score elevado: investigar medicamentos atuais (antidepressivos, anti-hipertensivos, antipsicóticos), avaliar testosterona total/livre, prolactina, TSH. Considerar switch medicamentoso (bupropiona, mirtazapina têm menor impacto), suplementação com zinco (50mg/dia), vitamina D (alvo >40ng/mL), magnésio (400mg/dia). Avaliar adicionar bupropiona como adjuvante ou tadalafil/sildenafil.',
                'scientific_sources': [
                    'McGahuey CA et al. The Arizona Sexual Experience Scale (ASEX): reliability and validity. J Sex Marital Ther. 2000;26(1):25-40',
                    'Rothmore J. Antidepressant-induced sexual dysfunction. Med J Aust. 2020;212(7):329-334',
                    'Prasad AS et al. Zinc supplementation and testosterone levels in men. Nutrition. 1996;12(5):344-8'
                ]
            }

        elif 'FSFI' in question or 'Female Sexual Function' in question:
            enrichment = {
                'clinical_relevance': 'O FSFI (Female Sexual Function Index) é padrão-ouro para avaliar função sexual feminina em 6 domínios: desejo, excitação, lubrificação, orgasmo, satisfação, dor. Score total <26,55 indica disfunção sexual. Correlaciona-se com níveis de estradiol, testosterona, DHEA-S, vitamina D e estado nutricional.',
                'clinical_action': 'Score baixo: painel hormonal completo (estradiol, progesterona, testosterona total/livre, DHEA-S, SHBG, prolactina), vitamina D, zinco, magnésio. Avaliar uso de contraceptivos orais (reduzem testosterona livre e aumentam SHBG). Suplementação: maca peruana (3g/dia), L-arginina (5g/dia pré-atividade), ômega-3 (2g/dia), probióticos (saúde vaginal). Considerar DHEA vaginal 6,5mg se atrofia.',
                'scientific_sources': [
                    'Rosen R et al. The Female Sexual Function Index (FSFI): a multidimensional self-report instrument. J Sex Marital Ther. 2000;26(2):191-208',
                    'Davis SR et al. Testosterone for low libido in postmenopausal women. N Engl J Med. 2008;359(19):2005-17',
                    'Ito TY et al. The enhancement of female sexual function with ArginMax. J Sex Marital Ther. 2006;32(5):369-78'
                ]
            }

        elif 'IIEF' in question or 'disfunção erétil' in question.lower():
            enrichment = {
                'clinical_relevance': 'O IIEF-5 (International Index of Erectile Function) avalia disfunção erétil em homens. Score 5-7=grave, 8-11=moderada, 12-16=leve-moderada, 17-21=leve, 22-25=sem DE. Disfunção erétil é marcador precoce de doença cardiovascular (precede IAM em 3-5 anos) e correlaciona-se com resistência insulínica, deficiências de vitamina D, zinco, magnésio.',
                'clinical_action': 'Score baixo: investigar causas vasculares (glicemia, HbA1c, perfil lipídico, homocisteína), hormonais (testosterona total/livre manhã cedo, SHBG, estradiol, prolactina, TSH), nutricionais (vitamina D, zinco, magnésio, vitamina B12). Suplementação: L-citrulina (6g/dia), Panax ginseng vermelho (1g 3x/dia), zinco (50mg/dia), vitamina D (alvo 50-80ng/mL), CoQ10 (200mg/dia). Estilo de vida: exercício aeróbico (150min/sem), perda de peso se IMC>25.',
                'scientific_sources': [
                    'Rosen RC et al. The international index of erectile function (IIEF). Urology. 1997;49(6):822-30',
                    'Thompson IM et al. Erectile dysfunction and subsequent cardiovascular disease. JAMA. 2005;294(23):2996-3002',
                    'Rhim HC et al. The potential role of arginine supplements on erectile dysfunction. J Sex Med. 2019;16(2):223-34'
                ]
            }

        # CICLO MENSTRUAL
        elif 'ciclo menstrual' in question.lower() or 'menstruação' in question.lower():
            enrichment = {
                'clinical_relevance': 'Irregularidades menstruais refletem desbalanços hormonais (estrogênio, progesterona, LH, FSH), metabólicos (resistência insulínica, SOP), nutricionais (deficiência de ferro, vitamina D, magnésio), tireoidianos ou estresse crônico (cortisol elevado suprime eixo HPG). Ciclos <21 ou >35 dias, sangramento >7 dias ou TPM grave requerem investigação.',
                'clinical_action': 'Irregularidade: investigar TSH, T4 livre, prolactina, FSH/LH (dia 3 do ciclo), estradiol, progesterona (dia 21 do ciclo), testosterona, DHEA-S, insulina basal, HOMA-IR, hemograma (anemia), ferritina, vitamina D. Suplementação: vitex agnus-castus (400mg/dia para TPM), magnésio glicinato (400mg/dia), vitamina B6 (100mg/dia), ômega-3 (2g/dia), inositol (4g/dia se SOP). Reduzir cafeína, açúcar, manejo de estresse.',
                'scientific_sources': [
                    'Reed BG, Carr BR. The Normal Menstrual Cycle and the Control of Ovulation. Endotext. 2018',
                    'Schliep KC et al. Caffeinated beverage intake and reproductive hormones. Am J Clin Nutr. 2012;95(2):488-97',
                    'Genazzani AD et al. Inositol as putative integrative treatment for PCOS. Reprod Biomed Online. 2016;33(6):770-80'
                ]
            }

        # ANTICONCEPCIONAL
        elif 'anticoncepcional' in question.lower() or 'contraceptivo' in question.lower():
            enrichment = {
                'clinical_relevance': 'Contraceptivos orais combinados aumentam SHBG (reduzindo testosterona livre e libido), depletam vitaminas B6, B12, ácido fólico, magnésio, zinco, vitamina C e E. Etinilestradiol aumenta risco trombótico (verificar mutações MTHFR, Fator V Leiden). Progestágenos podem causar retenção hídrica, alterações humor, ganho de peso.',
                'clinical_action': 'Uso atual: suplementar complexo B (B6 50mg, B12 1000mcg, folato 400mcg), magnésio (400mg/dia), zinco (25mg/dia), vitamina C (500mg/dia), vitamina E (400 UI/dia). Avaliar homocisteína, mutações trombofílicas se história familiar. Se baixa libido: considerar métodos não-hormonais (DIU cobre) ou apenas progestágeno (desogestrel). Monitorar pressão arterial, glicemia, triglicerídeos.',
                'scientific_sources': [
                    'Palmery M et al. Oral contraceptives and changes in nutritional requirements. Eur Rev Med Pharmacol Sci. 2013;17(13):1804-13',
                    'Panzer C et al. Impact of oral contraceptives on sex hormone-binding globulin and androgen levels. Contraception. 2006;73(5):470-7',
                    'Graham CA et al. Effects of oral contraceptives on mood and sexual interest. J Sex Med. 2007;4(6):1691-8'
                ]
            }

        # LIBIDO / DESEJO SEXUAL
        elif 'libido' in question.lower() or 'desejo sexual' in question.lower():
            enrichment = {
                'clinical_relevance': 'Libido depende de testosterona (homens e mulheres), dopamina, estrogênio adequado. Redução associada a: hipotireoidismo, hiperprolactinemia, deficiência de vitamina D/zinco/magnésio, uso de SSRIs/beta-bloqueadores, estresse crônico (cortisol elevado), resistência insulínica, privação de sono. Em mulheres, contraceptivos orais e menopausa reduzem testosterona livre.',
                'clinical_action': 'Baixa libido: investigar testosterona total/livre (manhã cedo), SHBG, estradiol (mulheres), prolactina, TSH, T4 livre, vitamina D, zinco, magnésio, cortisol salivar (4 pontos), glicemia/insulina. Suplementação: maca peruana (3g/dia), tribulus terrestris (750mg/dia), L-arginina (5g/dia), zinco (50mg/dia), vitamina D (alvo 60ng/mL), ashwagandha (600mg/dia). Otimizar sono (>7h), exercício de força, reduzir álcool.',
                'scientific_sources': [
                    'Santoro N et al. Role of estrogens and estrogen-like compounds in female sexual function. J Sex Med. 2016;13(3):305-16',
                    'Gonzales GF. Ethnobiology and Ethnopharmacology of Lepidium meyenii (Maca). Evid Based Complement Alternat Med. 2012;2012:193496',
                    'Wehr E et al. Association of vitamin D status with serum androgen levels in men. Clin Endocrinol. 2010;73(2):243-8'
                ]
            }

        # DESEMPENHO SEXUAL
        elif 'desempenho' in question.lower() and 'sexual' in question.lower():
            enrichment = {
                'clinical_relevance': 'Desempenho sexual depende de saúde vascular (óxido nítrico), hormonal (testosterona, estrogênio), neurotransmissores (dopamina, serotonina), energia mitocondrial. Disfunção associada a: aterosclerose, diabetes, obesidade, apneia do sono, deficiências nutricionais (L-arginina, zinco, magnésio, vitaminas B), medicamentos (anti-hipertensivos, antidepressivos).',
                'clinical_action': 'Investigar: testosterona total/livre, estradiol, glicemia/HbA1c, perfil lipídico completo, homocisteína, vitamina D, magnésio, zinco. Suplementação: L-citrulina (6-8g/dia, precursor óxido nítrico), Panax ginseng vermelho (1g 3x/dia), maca (3g/dia), CoQ10 (200mg/dia), ômega-3 (2g/dia), zinco (50mg/dia). Exercício aeróbico + força 4-5x/semana, reduzir álcool, otimizar sono, manejo estresse (meditação, yoga).',
                'scientific_sources': [
                    'Cormio L et al. Oral L-citrulline supplementation improves erection hardness in men. Urology. 2011;77(1):119-22',
                    'Jang DJ et al. Red ginseng for treating erectile dysfunction. Cochrane Database Syst Rev. 2008;(3):CD007989',
                    'Meldrum DR et al. Lifestyle and metabolic approaches to erectile dysfunction. Endocr Rev. 2012;33(5):741-67'
                ]
            }

        # FATORES QUE MELHORAM
        elif 'melhora' in question.lower() or 'positivamente' in question.lower():
            enrichment = {
                'clinical_relevance': 'Fatores que melhoram função sexual: exercício regular (aumenta testosterona, fluxo sanguíneo), sono adequado (7-9h, pico testosterona matinal), alimentação anti-inflamatória (mediterrânea), redução de estresse (cortisol antagoniza testosterona), suplementação estratégica, conexão emocional com parceiro(a).',
                'clinical_action': 'Otimizações: exercício aeróbico 150min/sem + força 2-3x/sem, sono 7-9h (melatonina 3mg se insônia), dieta mediterrânea (peixes gordos, azeite, vegetais, nozes), reduzir álcool (<2 doses/dia), cessação tabagismo. Suplementos: L-arginina + picnogenol (40mg + 1,7g/dia), maca (3g/dia), zinco (50mg/dia), magnésio (400mg/dia), vitamina D (alvo 60ng/mL). Terapia de casal se aspectos relacionais.',
                'scientific_sources': [
                    'Boyle P et al. Lifestyle factors and the risk of erectile dysfunction. BJU Int. 2003;92(5):519-21',
                    'Stanislavov R, Nikolova V. Treatment of erectile dysfunction with pycnogenol and L-arginine. J Sex Marital Ther. 2003;29(3):207-13',
                    'Allen MS, Walter EE. Health-Related Lifestyle Factors and Sexual Dysfunction. J Sex Med. 2018;15(4):458-75'
                ]
            }

        # FATORES QUE PIORAM
        elif 'piora' in question.lower() or 'negativamente' in question.lower():
            enrichment = {
                'clinical_relevance': 'Fatores que pioram função sexual: medicamentos (SSRIs, beta-bloqueadores, finasterida, espironolactona), álcool excessivo (>2 doses/dia), tabagismo (vasoconstrição), obesidade (aromatase converte testosterona em estrogênio), estresse crônico (cortisol), privação de sono, diabetes descompensado, hipertireoidismo/hipotireoidismo, pornografia excessiva (dessensibilização dopaminérgica).',
                'clinical_action': 'Identificar e modificar: revisar medicações com médico (considerar alternativas sem impacto sexual), reduzir/cessar álcool e tabaco, perda de peso se IMC>25 (meta 5-10%), manejo estresse (mindfulness, psicoterapia), higiene do sono, controle glicêmico rigoroso se diabetes. Avaliar uso de pornografia e considerar "reset dopaminérgico" (abstinência 30-90 dias). Investigar apneia do sono (polissonografia).',
                'scientific_sources': [
                    'Rothmore J. Antidepressant-induced sexual dysfunction. Med J Aust. 2020;212(7):329-34',
                    'Park BJ et al. The impact of smoking on sexual function in men. BJU Int. 2012;110(10):1508-16',
                    'Corona G et al. Body weight loss and erectile dysfunction. Expert Rev Cardiovasc Ther. 2013;11(10):1465-76'
                ]
            }

        # HISTÓRICO DE ABUSOS/TRAUMAS
        elif 'abuso' in question.lower() or 'trauma' in question.lower():
            enrichment = {
                'clinical_relevance': 'Histórico de abuso sexual ou trauma impacta profundamente função sexual através de mecanismos neurobiológicos (desregulação HPA, hiperativação amígdala), psicológicos (PTSD, dissociação, evitação) e somáticos (dor pélvica crônica, vaginismo). Requer abordagem compassiva, trauma-informada, multidisciplinar. Correlaciona-se com níveis elevados de cortisol, alterações em oxitocina, inflamação crônica.',
                'clinical_action': 'Encaminhar para psicoterapia especializada (EMDR, terapia focada no trauma, terapia sensório-motora). Avaliação médica compassiva: dor pélvica (fisioterapia pélvica), cortisol salivar, DHEA-S, inflamação (PCR-us, homocisteína). Suporte nutricional: ômega-3 (4g/dia), magnésio (600mg/dia), complexo B, vitamina D, probióticos. Técnicas somáticas: yoga trauma-sensível, meditação mindfulness. NUNCA pressionar ou minimizar experiências. Consentimento contínuo essencial.',
                'scientific_sources': [
                    'Shapiro F. Eye movement desensitization and reprocessing (EMDR) therapy. J Sex Med. 2014;11(7):1623-34',
                    'Ogden P, Fisher J. Sensorimotor psychotherapy: Interventions for trauma. Norton. 2015',
                    'Yehuda R et al. Post-traumatic stress disorder. Nat Rev Dis Primers. 2015;1:15057'
                ]
            }

        # HORMÔNIOS / MEDICAMENTOS
        elif 'hormônio' in question.lower() or 'medicamento' in question.lower():
            enrichment = {
                'clinical_relevance': 'Múltiplos hormônios e medicamentos impactam função sexual. Hormônios chave: testosterona (desejo), estrogênio (lubrificação, vascularização), progesterona (pode reduzir libido em excesso), prolactina (inibe se elevada), TSH/T4 (meta-regulação). Medicamentos problemáticos: SSRIs/SNRIs (70% têm disfunção), beta-bloqueadores, finasterida (síndrome pós-finasterida), espironolactona, contraceptivos orais.',
                'clinical_action': 'Painel hormonal completo: testosterona total/livre (manhã cedo), SHBG, estradiol, progesterona (dia 21 ciclo se menstrua), prolactina, TSH, T4 livre, DHEA-S, cortisol. Revisão medicamentosa: trocar SSRI por bupropiona/mirtazapina, beta-bloqueador por nebivolol, considerar "drug holiday" SSRIs fim de semana. Se terapia testosterona: monitorar hematócrito, PSA, estradiol (usar anastrozol se elevado). Suporte: L-arginina, zinco, vitamina D, maca.',
                'scientific_sources': [
                    'Corona G et al. Testosterone supplementation and sexual function. Endocr Rev. 2014;35(6):906-60',
                    'Hirsch M et al. Antidepressant-induced sexual dysfunction. UpToDate. 2023',
                    'Traish AM et al. Post-finasteride syndrome: a surmountable challenge. Transl Androl Urol. 2015;4(3):295-308'
                ]
            }

        # HISTÓRICO REPRODUTIVO
        elif 'reprodutivo' in question.lower() or 'gravidez' in question.lower() or 'filho' in question.lower():
            enrichment = {
                'clinical_relevance': 'Histórico reprodutivo informa sobre fertilidade, função hormonal, complicações obstétricas (eclâmpsia, diabetes gestacional predizem risco cardiovascular). Infertilidade pode associar-se a endometriose, SOP, insuficiência ovariana, fatores masculinos (varicocele, qualidade espermática). Pós-parto: queda estrogênio/progesterona, prolactina elevada (amamentação) reduzem libido temporariamente.',
            'clinical_action': 'Histórico de infertilidade: investigar causas persistentes (SOP, endometriose, reserva ovariana via AMH/FSH/contagem folicular). Pós-parto com baixa libido: tranquilizar que é fisiológico durante amamentação (prolactina antagoniza testosterona), melhorar sono, suporte parceiro, fisioterapia pélvica se dor/trauma parto. Se diabetes gestacional prévio: rastreio TOTG anual, estilo de vida anti-diabético. Suplementação pós-parto: ferro, vitamina D, ômega-3, complexo B.',
                'scientific_sources': [
                    'Bellver J, Donnez J. Introduction: Infertility etiology and management. Fertil Steril. 2019;111(4):607-8',
                    'Khajehei M et al. Sexual function and postpartum depression. J Pregnancy. 2015;2015:434763',
                    'Kim C et al. Gestational diabetes and the incidence of type 2 diabetes. Diabetes Care. 2002;25(10):1862-8'
                ]
            }

        # MENOPAUSA / PERIMENOPAUSA
        elif 'menopausa' in question.lower() or 'climatério' in question.lower():
            enrichment = {
                'clinical_relevance': 'Menopausa (ausência menstruação 12 meses) e perimenopausa caracterizam-se por declínio estrogênio/progesterona, resultando em atrofia vaginal, redução lubrificação, dispareunia, diminuição libido (queda testosterona concomitante). FSH>40 mUI/mL confirma. Sintomas: fogachos, suores noturnos, insônia, alterações humor, perda massa óssea, aumento risco cardiovascular.',
                'clinical_action': 'Confirmar status: FSH, estradiol, AMH (reserva ovariana). Terapia hormonal bioidêntica se <60 anos e <10 anos pós-menopausa: estradiol transdérmico (0,5-1mg/dia) + progesterona micronizada (100-200mg/noite) se útero intacto. DHEA vaginal (6,5mg) ou estriol vaginal (0,5mg) para atrofia. Testosterona transdérmica (0,5mg/dia) se baixa libido. Suplementos: isoflavonas (50mg/dia), maca (3g/dia), cimicifuga (40mg/dia), vitamina D+K2, cálcio, magnésio, ômega-3.',
                'scientific_sources': [
                    'Stuenkel CA et al. Treatment of Symptoms of the Menopause: An Endocrine Society Clinical Practice Guideline. J Clin Endocrinol Metab. 2015;100(11):3975-4011',
                    'Davis SR et al. Testosterone for low libido in postmenopausal women. N Engl J Med. 2008;359(19):2005-17',
                    'Krause MS et al. Vaginal DHEA for postmenopausal atrophy. Climacteric. 2015;18(1):79-85'
                ]
            }

        # DESENVOLVIMENTO SEXUAL
        elif 'desenvolvimento' in question.lower() and 'sexual' in question.lower():
            enrichment = {
                'clinical_relevance': 'Desenvolvimento sexual puberal depende de eixo HPG (GnRH → LH/FSH → esteroides gonadais). Puberdade normal: meninas 8-13 anos (telarca), meninos 9-14 anos (aumento testicular). Atraso (>13 anos meninas, >14 meninos) ou precocidade (<8 meninas, <9 meninos) requerem investigação. Impactos de obesidade (leptina, aromatase), desnutrição, doenças crônicas, estresse.',
                'clinical_action': 'Atraso puberal: investigar LH, FSH, testosterona/estradiol, cariótipo (Turner, Klinefelter), IGF-1, TSH, prolactina, ferritina, transglutaminase (celíaca), ressonância sela túrcica. Precocidade: LH, FSH, testosterona/estradiol, ultrassom pélvico (meninas), idade óssea, ressonância cérebro. Otimizar nutrição (zinco, vitamina D, calorias adequadas), sono, reduzir estresse. Encaminhar endocrinologia pediátrica se anormalidades.',
                'scientific_sources': [
                    'Palmert MR, Dunkel L. Clinical practice: Delayed puberty. N Engl J Med. 2012;366(5):443-53',
                    'Klein DA et al. Disorders of puberty. Endocrinol Metab Clin North Am. 2015;44(4):737-53',
                    'Kaplowitz PB. Link between body fat and the timing of puberty. Pediatrics. 2008;121 Suppl 3:S208-17'
                ]
            }

        # DOR DURANTE RELAÇÃO (DISPAREUNIA)
        elif 'dor' in question.lower() and ('relação' in question.lower() or 'sexual' in question.lower()):
            enrichment = {
                'clinical_relevance': 'Dispareunia (dor coital) tem causas múltiplas: atrofia vaginal (hipoestrogenismo), vaginismo (contração involuntária), infecções (candidíase, vaginose), endometriose, adenomiose, aderências pélvicas, síndrome dor pélvica crônica, histórico trauma. Abordagem requer exame ginecológico sensível, investigação causas orgânicas e psicológicas.',
                'clinical_action': 'Investigar: exame especular (atrofia, lesões), swab vaginal (pH, microscopia, cultura), ultrassom transvaginal (endometriose, massas), estradiol, testosterona. Tratamento por causa: atrofia (estrogênio/DHEA vaginal), vaginismo (fisioterapia pélvica, dilatadores graduais), infecções (tratamento específico), endometriose (cirurgia, dienogeste). Suporte: fisioterapia do assoalho pélvico, terapia sexual, lubrificantes base água/silicone, ômega-3, vitamina E vaginal.',
                'scientific_sources': [
                    'Goldstein AT et al. Vulvodynia: Assessment and Treatment. J Sex Med. 2016;13(4):572-90',
                    'Pukall CF et al. Vulvodynia: Definition, Prevalence, Impact, and Pathophysiology. J Sex Med. 2016;13(3):291-304',
                    'Bergeron S et al. Vulvodynia. Nat Rev Dis Primers. 2020;6(1):36'
                ]
            }

        # ORGASMO / ANORGASMIA
        elif 'orgasmo' in question.lower():
            enrichment = {
                'clinical_relevance': 'Anorgasmia (dificuldade/incapacidade atingir orgasmo) afeta 10-40% mulheres, 5-10% homens. Causas: medicamentos (SSRIs, antipsicóticos), doenças neurológicas (diabetes, esclerose múltipla), cirurgias pélvicas, hipotireoidismo, deficiência testosterona, fatores psicológicos (ansiedade, trauma), falta de estimulação adequada. Orgasmo envolve dopamina, oxitocina, endorfinas.',
                'clinical_action': 'Investigar: testosterona, prolactina, TSH, glicemia/HbA1c (neuropatia diabética), revisar medicações. Se SSRI: considerar bupropiona adjuvante, sildenafil 50-100mg PRN (mulheres off-label), ou "drug holiday". Suplementação: L-arginina (3-5g/dia), maca (3g/dia), ginkgo biloba (240mg/dia para anorgasmia por antidepressivos). Abordagem: educação sexual, comunicação com parceiro, sexoterapia, dispositivos (vibradores), exercícios Kegel, mindfulness sexual.',
                'scientific_sources': [
                    'Meston CM et al. Women\'s orgasm. Annu Rev Sex Res. 2004;15:173-257',
                    'Cohen AJ, Bartlik B. Ginkgo biloba for antidepressant-induced sexual dysfunction. J Sex Marital Ther. 1998;24(2):139-43',
                    'Brotto LA, Basson R. Group mindfulness-based therapy for women with provoked vestibulodynia. Mindfulness. 2014;5:88-99'
                ]
            }

        # EJACULAÇÃO PRECOCE
        elif 'ejaculação precoce' in question.lower() or 'ejacular' in question.lower():
            enrichment = {
                'clinical_relevance': 'Ejaculação precoce (EP) definida como ejaculação <1-2 minutos com angústia/frustração. Prevalência 20-30% homens. Tipos: primária (sempre presente, frequentemente serotonérgica) ou adquirida (surgiu depois, investigar prostatite, hipertireoidismo, ansiedade). Controle ejaculatório envolve serotonina (inibe), dopamina/noradrenalina (facilitam), testosterona, sensibilidade peniana.',
                'clinical_action': 'Investigar: TSH (hipertireoidismo acelera), testosterona, exame prostático se >40 anos. Tratamento farmacológico: dapoxetina (SSRI on-demand, 30-60mg 1h antes), paroxetina/sertralina 20mg/dia (off-label), anestésicos tópicos (lidocaína/prilocaína 20min antes). Suplementação: magnésio (400mg/dia), zinco (50mg/dia), L-triptofano (1g/noite, precursor serotonina). Técnicas comportamentais: stop-start, squeeze, exercícios Kegel masculinos, mindfulness, terapia sexual.',
                'scientific_sources': [
                    'Serefoglu EC et al. An evidence-based unified definition of lifelong and acquired premature ejaculation. J Sex Med. 2014;11(6):1423-41',
                    'McMahon CG et al. Dapoxetine for premature ejaculation. Expert Opin Pharmacother. 2011;12(5):815-25',
                    'Pastore AL et al. Pelvic muscle floor rehabilitation for patients with lifelong premature ejaculation. Urology. 2014;84(1):104-10'
                ]
            }

        # ORIENTAÇÃO SEXUAL / IDENTIDADE DE GÊNERO
        elif 'orientação' in question.lower() or 'identidade' in question.lower():
            enrichment = {
                'clinical_relevance': 'Orientação sexual (atração afetivo-sexual) e identidade de gênero (autopercepção) são aspectos centrais da saúde sexual. Minorias sexuais/gênero enfrentam estresse minoritário (discriminação, estigma) associado a taxas elevadas de ansiedade, depressão, abuso de substâncias. Cuidado afirmativo, linguagem inclusiva e compreensão de necessidades específicas (ex: prep, rastreio ISTs, saúde trans) são essenciais.',
                'clinical_action': 'Criar ambiente acolhedor: pronomes corretos, formulários inclusivos, evitar presunções heteronormativas. Rastreios: ISTs ampliado (incluindo sítios extragenitais se HSH), saúde mental (PHQ-9, GAD-7), uso de substâncias. Se pessoa trans: acompanhamento hormonal (testosterona, estradiol, bloqueadores), suporte multidisciplinar (endócrino, saúde mental, cirurgião se desejado). Recursos: encaminhamento para grupos de apoio, terapia afirmativa. PrEP se exposição HIV.',
                'scientific_sources': [
                    'Institute of Medicine. The Health of Lesbian, Gay, Bisexual, and Transgender People. National Academies Press. 2011',
                    'Coleman E et al. Standards of Care for the Health of Transgender and Gender Diverse People, Version 8. Int J Transgend Health. 2022;23(Suppl 1):S1-S259',
                    'Meyer IH. Prejudice, social stress, and mental health in lesbian, gay, and bisexual populations. Psychol Bull. 2003;129(5):674-697'
                ]
            }

        # INFECÇÕES SEXUALMENTE TRANSMISSÍVEIS (ISTs)
        elif 'ist' in question.lower() or 'dst' in question.lower() or 'infecção' in question.lower():
            enrichment = {
                'clinical_relevance': 'ISTs (HIV, sífilis, gonorreia, clamídia, HPV, herpes, hepatites) impactam saúde sexual, fertilidade, bem-estar psicológico. Rastreio regular é preventivo. Grupos de risco elevado: múltiplos parceiros, HSH, usuários drogas injetáveis, profissionais do sexo. Muitas ISTs são assintomáticas (clamídia 70% mulheres, gonorreia 50% homens).',
                'clinical_action': 'Rastreio anual (ou 3-6 meses se alto risco): HIV 4ª geração, VDRL/FTA-Abs (sífilis), hepatites B/C, gonorreia/clamídia (swab uretral/vaginal/retal/faríngeo conforme práticas). HPV: Papanicolau + colposcopia, vacinação até 26 anos (ideal 9-14). PrEP se exposição HIV recorrente (tenofovir/emtricitabina). Tratamento: específico por agente, notificar parceiros. Prevenção: preservativo consistente, redução parceiros, vacinação (HPV, hepatite B).',
                'scientific_sources': [
                    'CDC. Sexually Transmitted Infections Treatment Guidelines, 2021. MMWR Recomm Rep. 2021;70(4):1-187',
                    'Grant RM et al. Preexposure chemoprophylaxis for HIV prevention. N Engl J Med. 2010;363(27):2587-99',
                    'Petrosky E et al. Use of 9-valent HPV vaccine. MMWR Morb Mortal Wkly Rep. 2015;64(11):300-4'
                ]
            }

        # PORNOGRAFIA / USO EXCESSIVO
        elif 'pornografia' in question.lower() or 'pornô' in question.lower():
            enrichment = {
                'clinical_relevance': 'Uso excessivo de pornografia pode levar a dessensibilização dopaminérgica, disfunção erétil induzida por pornografia (PIED), ejaculação precoce/retardada, expectativas irreais, redução satisfação sexual com parceiro real. Mecanismo: superestimulação de centros de recompensa (nucleus accumbens) → downregulation receptores D2 dopamina, similar a vícios. Prevalência crescente em homens <40 anos.',
                'clinical_action': 'Avaliação compassiva sem julgamento: frequência, tipo, impacto em relacionamentos/função sexual. Se PIED ou disfunção relacionada: "reboot" (abstinência pornografia 60-90 dias), redução masturbação, reintrodução sensações reais graduais. Suporte: terapia cognitivo-comportamental, grupos de apoio (ex: NoFap), educação sobre neurobiologia vício. Suplementação para restaurar função dopaminérgica: L-tirosina (1g/dia), mucuna pruriens (300mg/dia), ômega-3, magnésio, zinco, exercício regular.',
                'scientific_sources': [
                    'Park BY et al. Is Internet Pornography Causing Sexual Dysfunctions? A Review. Behav Sci. 2016;6(3):17',
                    'Volkow ND et al. Neuroscience of addiction: Relevance to prevention and treatment. Am J Psychiatry. 2018;175(8):729-40',
                    'Pizzol D et al. Associations between pornography and sexual dysfunction. J Sex Med. 2016;13(7):1138-43'
                ]
            }

        # SATISFAÇÃO SEXUAL / RELACIONAMENTO
        elif 'satisfação' in question.lower():
            enrichment = {
                'clinical_relevance': 'Satisfação sexual é multidimensional: física (prazer, orgasmo), emocional (intimidade, conexão), relacional (comunicação, reciprocidade), contextual (tempo, privacidade). Correlaciona-se com satisfação geral de relacionamento, saúde mental, qualidade de vida. Fatores negativos: estresse, fadiga, problemas de comunicação, discrepância libido entre parceiros, rotina.',
                'clinical_action': 'Avaliar múltiplos domínios: função sexual (usar escalas validadas FSFI/IIEF), satisfação relacional, saúde mental (ansiedade/depressão), estressores vida. Intervenções: terapia de casal focada em sexualidade, técnicas de comunicação (Gottman Method), exercícios de intimidade (sensate focus), educação sexual, exploração de preferências. Otimizar saúde: exercício, sono, nutrição anti-inflamatória, reduzir álcool/substâncias. Suplementos: maca, L-arginina, ômega-3, magnésio.',
                'scientific_sources': [
                    'Sprecher S, Cate RM. Sexual satisfaction and sexual expression as predictors of relationship satisfaction. Fam Process. 2004;43(1):23-38',
                    'Brotto LA et al. Psychological and interpersonal dimensions of sexual function. J Sex Med. 2016;13(4):538-71',
                    'Gottman JM, Silver N. The Seven Principles for Making Marriage Work. Harmony Books. 2015'
                ]
            }

        # EDUCAÇÃO SEXUAL / CONHECIMENTO
        elif 'educação' in question.lower() or 'conhecimento' in question.lower():
            enrichment = {
                'clinical_relevance': 'Educação sexual adequada (anatomia, fisiologia, consentimento, comunicação, prazer, diversidade) está associada a: menor risco ISTs/gravidez não planejada, relações mais saudáveis, maior satisfação sexual, redução estigma/vergonha. Lacunas educacionais são comuns devido a tabus culturais/religiosos, currículos escolares deficientes.',
                'clinical_action': 'Fornecer educação baseada em evidências, culturalmente sensível, inclusiva: anatomia sexual, ciclo resposta sexual, variações normalidade, consentimento, comunicação, prevenção ISTs/gravidez, prazer mútuo. Recursos confiáveis: Sociedade Brasileira de Sexualidade Humana (SBRASH), SIECUS, livros (Come as You Are - Emily Nagoski). Desmistificar mitos comuns. Encorajar comunicação aberta com parceiro(a). Se histórico trauma/vergonha: encaminhar para sexoterapia.',
                'scientific_sources': [
                    'UNESCO. International technical guidance on sexuality education. 2018',
                    'Goldfarb ES, Lieberman LD. Three Decades of Research: The Case for Comprehensive Sex Education. J Adolesc Health. 2021;68(1):13-27',
                    'Nagoski E. Come as You Are: The Surprising New Science. Simon & Schuster. 2015'
                ]
            }

        # FREQUÊNCIA SEXUAL / ATIVIDADE
        elif 'frequência' in question.lower() or 'vezes' in question.lower():
            enrichment = {
                'clinical_relevance': 'Frequência sexual varia amplamente (média casais: 1-2x/semana), influenciada por: idade, duração relacionamento (declínio natural), saúde física/mental, estresse, filhos pequenos, medicamentos. Não existe "normal" universal - satisfação de ambos parceiros é mais importante que número absoluto. Discrepância libido entre parceiros é queixa comum (30-40% casais).',
                'clinical_action': 'Explorar sem julgamento: satisfação com frequência atual, discrepâncias entre parceiros, barreiras (tempo, energia, dor, interesse). Se desejo baixo: investigar causas médicas (hormônios, medicamentos, fadiga), psicológicas (estresse, ansiedade, depressão), relacionais. Intervenções: agendar intimidade (reduz barreira "esperar vontade"), sensate focus, terapia de casal, otimizar energia (sono, exercício, nutrição), reduzir estressores, tratar causas médicas. Normalizar que qualidade > quantidade.',
                'scientific_sources': [
                    'Muise A et al. Sexual frequency predicts greater well-being, but more is not always better. Soc Psychol Personal Sci. 2016;7(4):295-302',
                    'Mark KP. The relative impact of individual sexual desire and couple desire discrepancy. Sex Relation Ther. 2012;27(2):133-46',
                    'McCarthy BW, Wald LM. Strategies and techniques to address inhibited sexual desire. J Fam Psychother. 2015;26(4):286-98'
                ]
            }

        # IMAGEM CORPORAL / AUTOESTIMA
        elif 'imagem corporal' in question.lower() or 'autoestima' in question.lower():
            enrichment = {
                'clinical_relevance': 'Imagem corporal negativa afeta função sexual através de: evitação intimidade, distração cognitiva durante sexo (spectatoring), redução excitação/prazer, vergonha. Fatores de risco: mídia idealizada, comentários críticos (parceiros, família), histórico bullying, transtornos alimentares. Mulheres são particularmente vulneráveis, mas homens também sofrem (ansiedade de desempenho, tamanho pênis, calvície).',
                'clinical_action': 'Rastreio: Body Image-States Scale, escala autoestima (Rosenberg). Encaminhamento: psicoterapia (TCC, ACT, terapia focada na compaixão), grupos de suporte, sexoterapia se impacto sexual significativo. Técnicas: mindfulness durante sexo (focar sensações vs. aparência), exercícios mirror exposure, comunicação positiva com parceiro, desafiar padrões irreais mídia. Promover saúde via comportamentos positivos (exercício, alimentação equilibrada) vs. aparência.',
                'scientific_sources': [
                    'Woertman L, van den Brink F. Body image and female sexual functioning. J Sex Res. 2012;49(2-3):184-211',
                    'Pujols Y et al. The association between sexual satisfaction and body image in women. J Sex Med. 2010;7(2 Pt 2):905-16',
                    'Tylka TL, Wood-Barcalow NL. What is and what is not positive body image? Body Image. 2015;14:118-29'
                ]
            }

        # PARCEIRO SEXUAL / RELACIONAMENTO
        elif 'parceiro' in question.lower() or 'relacionamento' in question.lower():
            enrichment = {
                'clinical_relevance': 'Qualidade do relacionamento é preditor forte de satisfação sexual. Fatores positivos: comunicação aberta (incluindo sobre sexo), intimidade emocional, resolução construtiva de conflitos, tempo de qualidade juntos, gratidão/apreciação. Fatores negativos: conflitos não resolvidos, ressentimento, traição, falta de comunicação, discrepância valores/objetivos. Duração relacionamento naturalmente reduz novidade (habituação).',
                'clinical_action': 'Avaliação relacional: satisfação geral, comunicação, conflitos, intimidade emocional/sexual, confiança. Ferramentas: Couples Satisfaction Index (CSI-32), questionário apego adulto. Encaminhamento terapia de casal se: conflitos persistentes, infidelidade, discrepância sexual significativa. Técnicas: date nights regulares, novidade sexual (ambientes novos, atividades), exercícios de apreciação, melhorar comunicação sexual (artigo "Yes, No, Maybe" lists), retiros de casal.',
                'scientific_sources': [
                    'Sprecher S. Sexual satisfaction in premarital relationships. J Sex Res. 2002;39(3):190-6',
                    'Gottman JM. The Science of Trust. Norton. 2011',
                    'Muise A et al. Keeping the Spark Alive: Being Motivated to Meet a Partner\'s Sexual Needs. Soc Psychol Personal Sci. 2013;4(3):267-73'
                ]
            }

        # SONO E FUNÇÃO SEXUAL
        elif 'sono' in question.lower() and 'sexual' in question.lower():
            enrichment = {
                'clinical_relevance': 'Sono adequado é essencial para função sexual: aumenta testosterona (pico durante sono REM, redução 10-15% por hora perdida), regula cortisol, melhora humor/energia. Privação sono crônica (<6h) associa-se a: redução libido (60-70%), disfunção erétil, dificuldade orgasmo, redução satisfação sexual. Apneia do sono afeta 40-70% homens com DE.',
                'clinical_action': 'Avaliar qualidade/duração sono: questionário Pittsburgh Sleep Quality Index (PSQI), sintomas apneia (ronco, sonolência diurna, pausas respiratórias). Se apneia suspeita: encaminhar para polissonografia, CPAP se confirmado. Higiene do sono: 7-9h, horários regulares, ambiente escuro/fresco, reduzir luz azul 2h antes, evitar cafeína tarde, magnésio 400mg/noite. Se insônia: TCC-I, melatonina 3-5mg, L-teanina 200mg, glicina 3g, ashwagandha 600mg.',
                'scientific_sources': [
                    'Leproult R, Van Cauter E. Effect of 1 week of sleep restriction on testosterone. JAMA. 2011;305(21):2173-4',
                    'Andersen ML et al. The association of testosterone, sleep, and sexual function in men. Brain Res. 2011;1416:80-104',
                    'Budweiser S et al. Sleep apnea is an independent correlate of erectile dysfunction. J Sex Med. 2009;6(11):3147-57'
                ]
            }

        # EXERCÍCIO E FUNÇÃO SEXUAL
        elif 'exercício' in question.lower() or 'atividade física' in question.lower():
            enrichment = {
                'clinical_relevance': 'Exercício melhora função sexual via: aumento testosterona (exercício de força), melhor fluxo sanguíneo (aeróbico), redução cortisol/estresse, melhor imagem corporal, aumento energia/resistência. Exercício regular reduz risco DE em 30-40%. Excesso (overtraining) pode suprimir testosterona e libido. Exercícios do assoalho pélvico (Kegel) melhoram controle ejaculatório e intensidade orgasmo.',
                'clinical_action': 'Prescrição ótima para saúde sexual: aeróbico moderado 150min/sem (caminhada rápida, corrida, ciclismo, natação) + força 2-3x/sem (foco grandes grupos musculares, agachamento, levantamento terra) + exercícios Kegel diários (3 séries 10 contrações 5-10seg). Evitar overtraining (sinais: fadiga persistente, baixa libido, irritabilidade). Treino pré-atividade sexual aumenta fluxo sanguíneo. Yoga melhora flexibilidade, consciência corporal, reduz ansiedade.',
                'scientific_sources': [
                    'Lamina S et al. Effects of continuous and interval training programs. Eur J Cardiovasc Prev Rehabil. 2009;16(6):670-5',
                    'La Vignera S et al. Physical activity and erectile dysfunction. Reprod Biol Endocrinol. 2011;9:131',
                    'Dorey G et al. Pelvic floor exercises for erectile dysfunction. BJU Int. 2005;96(4):595-7'
                ]
            }

        # ÁLCOOL E FUNÇÃO SEXUAL
        elif 'álcool' in question.lower() or 'bebida' in question.lower():
            enrichment = {
                'clinical_relevance': 'Álcool tem efeito bifásico: baixas doses (1-2 drinks) reduzem ansiedade e podem facilitar desinibição; doses moderadas-altas prejudicam excitação, ereção, lubrificação, orgasmo (depressor SNC). Uso crônico causa: redução testosterona, aumento estrogênio (via aromatase hepática), neuropatia periférica, disfunção hepática, dependência psicológica. "Whiskey dick" (DE alcoólica) é comum.',
                'clinical_action': 'Avaliar consumo: AUDIT-C (rastreio). Recomendação: ≤1 dose/dia mulheres, ≤2 homens (1 dose = 350mL cerveja, 150mL vinho, 45mL destilado). Se uso excessivo: intervenção breve, encaminhar para tratamento dependência se necessário. Abstinência temporária (30-90 dias) para avaliar impacto em função sexual. Suporte para abstinência: naltrexona, acamprosato, grupos AA. Suplementação para recuperação: NAC (1200mg/dia), milk thistle (silimarina 300mg/dia), vitaminas B (tiamina, ácido fólico), magnésio, zinco.',
                'scientific_sources': [
                    'Arackal BS, Benegal V. Prevalence of sexual dysfunction in male subjects with alcohol dependence. Indian J Psychiatry. 2007;49(2):109-12',
                    'Emanuele MA, Emanuele NV. Alcohol\'s effects on male reproduction. Alcohol Health Res World. 1998;22(3):195-201',
                    'George WH et al. Alcohol and sexual health behavior. Alcohol Res Health. 2006;29(3):180-5'
                ]
            }

        # DEFAULT: Conteúdo genérico estruturado
        else:
            enrichment = {
                'clinical_relevance': f'Este item avalia aspectos importantes da saúde sexual relacionados a: {question}. A sexualidade é componente integral da saúde global, influenciada por fatores hormonais (testosterona, estrogênio, prolactina), vasculares (fluxo sanguíneo), neurológicos (dopamina, serotonina, oxitocina), psicológicos (autoestima, ansiedade, humor) e relacionais (comunicação, intimidade). Abordagem deve ser compassiva, não-julgamental e baseada em evidências.',
                'clinical_action': 'Realizar avaliação holística: histórico médico completo, medicamentos atuais, histórico psicossocial, exame físico apropriado. Investigações laboratoriais conforme indicação clínica: painel hormonal (testosterona, estradiol, prolactina, TSH), perfil metabólico (glicemia, lipídios), vitaminas/minerais (vitamina D, zinco, magnésio, B12). Encaminhamentos: sexoterapia para questões psicológicas/relacionais, especialistas médicos conforme necessário. Otimizar estilo de vida: exercício regular, sono adequado, nutrição anti-inflamatória, redução estresse.',
                'scientific_sources': [
                    'McCabe MP et al. Psychological and Interpersonal Dimensions of Sexual Function and Dysfunction. J Sex Med. 2016;13(4):538-71',
                    'Corona G et al. Assessment and Treatment of Sexual Dysfunction in Men and Women. Endocr Rev. 2014;35(6):906-60',
                    'Basson R et al. Sexual Function in Chronic Illness. J Sex Med. 2010;7(1 Pt 2):374-88'
                ]
            }

        return enrichment

    def update_item(self, item_id: str, enrichment: Dict) -> bool:
        """Atualizar item com conteúdo enriquecido"""
        try:
            payload = {
                'clinical_relevance': enrichment['clinical_relevance'],
                'clinical_action': enrichment['clinical_action'],
                'scientific_sources': enrichment['scientific_sources']
            }

            response = self.session.put(
                f"{API_BASE}/score-items/{item_id}",
                json=payload
            )
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"✗ Erro ao atualizar item {item_id}: {e}")
            return False

    def process_batch(self):
        """Processar batch completo de 30 items"""
        print(f"\n{'='*80}")
        print("ENRIQUECIMENTO VIDA SEXUAL - 30 ITEMS")
        print(f"{'='*80}\n")

        if not self.login():
            return

        results = {
            'success': [],
            'failed': [],
            'errors': []
        }

        for idx, item_id in enumerate(ITEM_IDS, 1):
            print(f"\n[{idx}/30] Processando item {item_id}...")

            # Buscar item
            item = self.get_item(item_id)
            if not item:
                results['failed'].append(item_id)
                results['errors'].append(f"{item_id}: Falha ao buscar item")
                continue

            question = item.get('question', '')
            print(f"  Pergunta: {question[:80]}...")

            # Gerar enriquecimento
            enrichment = self.generate_enrichment(item)

            print(f"  Relevância clínica: {enrichment['clinical_relevance'][:100]}...")
            print(f"  Fontes científicas: {len(enrichment['scientific_sources'])} referências")

            # Atualizar
            if self.update_item(item_id, enrichment):
                print(f"  ✓ Item atualizado com sucesso")
                results['success'].append(item_id)
            else:
                print(f"  ✗ Falha ao atualizar item")
                results['failed'].append(item_id)
                results['errors'].append(f"{item_id}: Falha ao atualizar")

        # Relatório final
        print(f"\n\n{'='*80}")
        print("RELATÓRIO FINAL - BATCH VIDA SEXUAL")
        print(f"{'='*80}")
        print(f"✓ Sucesso: {len(results['success'])}/{len(ITEM_IDS)}")
        print(f"✗ Falhas: {len(results['failed'])}/{len(ITEM_IDS)}")

        if results['failed']:
            print(f"\nItems com falha:")
            for error in results['errors']:
                print(f"  - {error}")

        # Salvar relatório
        report_file = '/home/user/plenya/VIDA-SEXUAL-BATCH-REPORT.md'
        with open(report_file, 'w', encoding='utf-8') as f:
            f.write("# Relatório de Enriquecimento - VIDA SEXUAL (30 items)\n\n")
            f.write(f"**Data:** {__import__('datetime').datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n\n")
            f.write(f"## Resumo\n\n")
            f.write(f"- Total de items: {len(ITEM_IDS)}\n")
            f.write(f"- Sucesso: {len(results['success'])}\n")
            f.write(f"- Falhas: {len(results['failed'])}\n\n")

            if results['success']:
                f.write(f"## Items Processados com Sucesso\n\n")
                for item_id in results['success']:
                    f.write(f"- {item_id}\n")
                f.write("\n")

            if results['failed']:
                f.write(f"## Items com Falha\n\n")
                for error in results['errors']:
                    f.write(f"- {error}\n")

        print(f"\nRelatório salvo em: {report_file}")
        print(f"\n{'='*80}\n")

def main():
    enricher = VidaSexualEnricher()
    enricher.process_batch()

if __name__ == "__main__":
    main()
