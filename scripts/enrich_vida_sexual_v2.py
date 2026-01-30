#!/usr/bin/env python3
"""
Enriquecimento AI dos 30 items de VIDA SEXUAL - VERSÃO 2
Foco: Medicina Funcional + Sensibilidade Clínica
Campos: clinicalRelevance, conduct, patientExplanation
"""

import requests
import json
import sys
import time
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
            self.token = data.get('accessToken')
            if not self.token:
                print(f"✗ Token não encontrado")
                return False
            self.session.headers.update({'Authorization': f'Bearer {self.token}'})
            print("✓ Autenticação bem-sucedida")
            return True
        except Exception as e:
            print(f"✗ Erro no login: {e}")
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
        name = item.get('name', '')

        # ASEX (Arizona Sexual Experience Scale)
        if 'ASEX' in name or 'Arizona' in name:
            return {
                'clinicalRelevance': 'A escala ASEX (Arizona Sexual Experience Scale) é ferramenta validada para avaliar disfunção sexual induzida por medicamentos, especialmente antidepressivos SSRIs/SNRIs. Score ≥19 ou qualquer item ≥5 indica disfunção sexual significativa. Particularmente sensível para detectar efeitos de paroxetina, sertralina, venlafaxina. Correlaciona-se com níveis de testosterona, prolactina e uso crônico de psicotrópicos.',
                'conduct': 'Score elevado (≥19): investigar medicamentos atuais (antidepressivos, anti-hipertensivos, antipsicóticos), avaliar testosterona total/livre (manhã cedo), prolactina, TSH. Considerar switch medicamentoso (bupropiona, mirtazapina têm menor impacto), adicionar bupropiona como adjuvante ou tadalafil/sildenafil. Suplementação: zinco 50mg/dia, vitamina D (alvo >40ng/mL), magnésio 400mg/dia, L-arginina 5g/dia. Reavaliar ASEX após 4-6 semanas de intervenção.',
                'patientExplanation': 'Esta escala avalia como sua função sexual pode estar sendo afetada por medicamentos, especialmente antidepressivos. Pontuação alta indica que pode haver impacto significativo. Existem alternativas de tratamento e suplementos que podem ajudar a melhorar este aspecto sem comprometer o tratamento da sua condição de base.'
            }

        # FSFI (Female Sexual Function Index)
        elif 'FSFI' in name or 'Female Sexual Function' in name:
            return {
                'clinicalRelevance': 'O FSFI (Female Sexual Function Index) é padrão-ouro para avaliar função sexual feminina em 6 domínios: desejo, excitação, lubrificação, orgasmo, satisfação, dor. Score total <26,55 indica disfunção sexual. Correlaciona-se fortemente com níveis de estradiol, testosterona, DHEA-S, vitamina D e estado nutricional. Contraceptivos orais podem reduzir significativamente o score ao aumentar SHBG e reduzir testosterona livre.',
                'conduct': 'Score baixo (<26,55): painel hormonal completo (estradiol, progesterona, testosterona total/livre, DHEA-S, SHBG, prolactina, TSH), vitamina D, zinco sérico, magnésio. Avaliar uso de contraceptivos orais (considerar métodos não-hormonais se baixa libido). Suplementação: maca peruana 3g/dia, L-arginina 5g/dia pré-atividade, ômega-3 2g/dia, probióticos vaginais. Considerar DHEA vaginal 6,5mg se atrofia. Encaminhar para sexoterapia se componente relacional/psicológico.',
                'patientExplanation': 'Este questionário avalia vários aspectos da sua função sexual feminina, incluindo desejo, excitação, lubrificação e satisfação. Pontuações baixas podem estar relacionadas a desequilíbrios hormonais, nutricionais ou emocionais. Existem tratamentos naturais e hormonais que podem melhorar significativamente estes aspectos.'
            }

        # IIEF-5 (Disfunção Erétil)
        elif 'IIEF' in name or 'Índice Internacional' in name:
            return {
                'clinicalRelevance': 'O IIEF-5 (International Index of Erectile Function) avalia disfunção erétil em homens. Score 5-7=grave, 8-11=moderada, 12-16=leve-moderada, 17-21=leve, 22-25=sem DE. Disfunção erétil é marcador precoce de doença cardiovascular (precede IAM em 3-5 anos) e está associada a resistência insulínica, deficiências de vitamina D, zinco, magnésio e redução de testosterona.',
                'conduct': 'Score <22: investigar causas vasculares (glicemia, HbA1c, perfil lipídico, homocisteína), hormonais (testosterona total/livre manhã cedo, SHBG, estradiol, prolactina, TSH), nutricionais (vitamina D, zinco, magnésio, B12). Suplementação: L-citrulina 6g/dia, Panax ginseng vermelho 1g 3x/dia, zinco 50mg/dia, vitamina D (alvo 50-80ng/mL), CoQ10 200mg/dia. Estilo de vida: exercício aeróbico 150min/sem, perda de peso se IMC>25. Considerar inibidores PDE-5 se não-responsivo.',
                'patientExplanation': 'Este questionário avalia a função erétil. Dificuldades nesta área podem ser um sinal precoce de problemas circulatórios ou hormonais. A boa notícia é que mudanças no estilo de vida, suplementação e tratamentos específicos geralmente trazem melhora significativa.'
            }

        # Ciclo Menstrual
        elif 'ciclo menstrual' in name.lower() or 'menstruaç' in name.lower() or 'menstruou' in name.lower():
            return {
                'clinicalRelevance': 'Irregularidades menstruais refletem desbalanços hormonais (estrogênio, progesterona, LH, FSH), metabólicos (resistência insulínica, SOP), nutricionais (deficiência de ferro, vitamina D, magnésio), tireoidianos ou estresse crônico (cortisol elevado suprime eixo HPG). Ciclos <21 ou >35 dias, sangramento >7 dias ou TPM grave requerem investigação. Amenorreia >3 meses pode indicar SOP, disfunção tiroidiana ou hiperprolactinemia.',
                'conduct': 'Irregularidade: investigar TSH, T4 livre, prolactina, FSH/LH (dia 3 do ciclo), estradiol, progesterona (dia 21 do ciclo), testosterona, DHEA-S, insulina basal, HOMA-IR, hemograma completo, ferritina, vitamina D. Suplementação: vitex agnus-castus 400mg/dia (TPM), magnésio glicinato 400mg/dia, vitamina B6 100mg/dia, ômega-3 2g/dia, inositol 4g/dia (se SOP). Reduzir cafeína (<200mg/dia), açúcar refinado, manejo de estresse (cortisol salivar se indicado).',
                'patientExplanation': 'O ciclo menstrual regular é um sinal de equilíbrio hormonal e saúde geral. Irregularidades podem estar relacionadas a tireoide, nutrição, estresse ou síndrome dos ovários policísticos. Investigação e suplementação adequadas geralmente normalizam o ciclo.'
            }

        # Anticoncepcional
        elif 'anticoncepcional' in name.lower() or 'contraceptivo' in name.lower() or 'DIU' in name or 'pílula' in name.lower():
            return {
                'clinicalRelevance': 'Contraceptivos orais combinados aumentam SHBG (reduzindo testosterona livre e libido em até 60%), depletam vitaminas B6, B12, ácido fólico, magnésio, zinco, vitamina C e E. Etinilestradiol aumenta risco trombótico (verificar mutações MTHFR, Fator V Leiden). Progestágenos podem causar retenção hídrica, alterações de humor, ganho de peso. DIU hormonal tem efeitos sistêmicos menores mas também pode afetar libido.',
                'conduct': 'Uso de contraceptivo oral: suplementar complexo B (B6 50mg, B12 1000mcg, folato 400mcg), magnésio 400mg/dia, zinco 25mg/dia, vitamina C 500mg/dia, vitamina E 400 UI/dia. Avaliar homocisteína (alvo <7 µmol/L), mutações trombofílicas se história familiar. Se baixa libido: considerar métodos não-hormonais (DIU cobre) ou apenas progestágeno (desogestrel). Monitorar pressão arterial trimestral, glicemia e triglicerídeos anualmente.',
                'patientExplanation': 'Anticoncepcionais hormonais podem afetar seus níveis de vitaminas e hormônios como testosterona, impactando energia e libido. Suplementação adequada ajuda a minimizar estes efeitos. Se houver queda importante de libido, existem métodos contraceptivos alternativos sem hormônios.'
            }

        # Libido / Desejo Sexual
        elif 'libido' in name.lower() or 'desejo sexual' in name.lower() or 'interesse por sexo' in name.lower():
            return {
                'clinicalRelevance': 'Libido depende de testosterona (homens e mulheres), dopamina, estrogênio adequado e saúde geral. Redução associada a: hipotireoidismo, hiperprolactinemia, deficiência de vitamina D/zinco/magnésio, uso de SSRIs/beta-bloqueadores, estresse crônico (cortisol elevado antagoniza testosterona), resistência insulínica, privação de sono. Em mulheres, contraceptivos orais e menopausa reduzem testosterona livre em até 70%.',
                'conduct': 'Baixa libido: investigar testosterona total/livre (manhã cedo), SHBG, estradiol (mulheres), prolactina, TSH, T4 livre, vitamina D, zinco sérico, magnésio RBC, cortisol salivar (4 pontos), glicemia/insulina, hemograma. Suplementação: maca peruana 3g/dia, tribulus terrestris 750mg/dia, L-arginina 5g/dia, zinco 50mg/dia, vitamina D (alvo 60ng/mL), ashwagandha 600mg/dia (reduz cortisol). Otimizar sono (>7h), exercício de força 3x/sem, reduzir álcool (<2 doses/dia).',
                'patientExplanation': 'A libido pode ser afetada por vários fatores: hormônios, nutrição, sono, estresse e medicamentos. Identificar e corrigir a causa geralmente traz bons resultados. Suplementos naturais como maca e ashwagandha, combinados com mudanças no estilo de vida, costumam melhorar significativamente o desejo sexual.'
            }

        # Desempenho Sexual
        elif 'desempenho' in name.lower() or 'ereção' in name.lower() or 'penetração' in name.lower():
            return {
                'clinicalRelevance': 'Desempenho sexual depende de saúde vascular (produção de óxido nítrico), hormonal (testosterona, estrogênio), neurotransmissores (dopamina, serotonina), e energia mitocondrial. Disfunção associada a: aterosclerose, diabetes (neuropatia autonômica), obesidade (aromatase aumentada), apneia do sono, deficiências nutricionais (L-arginina, zinco, magnésio, vitaminas B), medicamentos (anti-hipertensivos, antidepressivos).',
                'conduct': 'Investigar: testosterona total/livre, estradiol, glicemia/HbA1c, perfil lipídico completo, homocisteína, vitamina D, magnésio, zinco, PSA (>40 anos). Suplementação: L-citrulina 6-8g/dia (precursor óxido nítrico), Panax ginseng vermelho 1g 3x/dia, maca 3g/dia, CoQ10 200mg/dia, ômega-3 2g/dia, zinco 50mg/dia, picnogenol 40mg + L-arginina 1,7g/dia. Exercício aeróbico + força 4-5x/semana, reduzir álcool, otimizar sono (avaliar apneia se roncos), manejo estresse.',
                'patientExplanation': 'O desempenho sexual depende de boa circulação, hormônios equilibrados e energia. Problemas nesta área podem indicar questões vasculares ou metabólicas que precisam ser investigadas. Suplementos que melhoram a circulação e exercícios regulares costumam trazer excelentes resultados.'
            }

        # Fatores que Melhoram
        elif 'melhora' in name.lower() or 'favorece' in name.lower() or 'auxilia' in name.lower():
            return {
                'clinicalRelevance': 'Fatores que melhoram função sexual: exercício regular (aumenta testosterona 15-30%, melhora fluxo sanguíneo), sono adequado (7-9h, pico testosterona matinal), alimentação anti-inflamatória (mediterrânea), redução de estresse (cortisol antagoniza testosterona), suplementação estratégica, conexão emocional com parceiro(a). Exercício de força especialmente efetivo para aumentar testosterona.',
                'conduct': 'Otimizações: exercício aeróbico 150min/sem + força 2-3x/sem (agachamento, levantamento terra), sono 7-9h (melatonina 3mg se insônia), dieta mediterrânea (peixes gordos 3x/sem, azeite extravirgem, vegetais coloridos, nozes, sementes), reduzir álcool (<2 doses/dia), cessação tabagismo. Suplementos: L-arginina + picnogenol (1,7g + 40mg/dia), maca 3g/dia, zinco 50mg/dia, magnésio 400mg/dia, vitamina D (alvo 60ng/mL), ômega-3 2g/dia. Terapia de casal se aspectos relacionais.',
                'patientExplanation': 'Vários hábitos melhoram a função sexual: exercícios regulares (especialmente musculação), sono adequado, alimentação saudável, redução de estresse e suplementos naturais. Cuidar da saúde geral se reflete positivamente na vida sexual. Conexão emocional com o parceiro também é fundamental.'
            }

        # Fatores que Pioram
        elif 'piora' in name.lower() or 'prejudica' in name.lower() or 'dificulta' in name.lower():
            return {
                'clinicalRelevance': 'Fatores que pioram função sexual: medicamentos (SSRIs reduzem libido em 70%, beta-bloqueadores causam DE em 25%, finasterida causa síndrome pós-finasterida), álcool excessivo (>2 doses/dia reduz testosterona 20%), tabagismo (vasoconstrição, DE aumenta 85%), obesidade (aromatase converte testosterona em estrogênio), estresse crônico (cortisol elevado), privação de sono, diabetes descompensado, hiper/hipotireoidismo, pornografia excessiva (dessensibilização dopaminérgica).',
                'conduct': 'Identificar e modificar: revisar medicações com médico (alternativas sem impacto sexual: nebivolol vs outros beta-bloqueadores, bupropiona vs SSRIs), reduzir/cessar álcool e tabaco (vareniclina para cessação), perda de peso se IMC>25 (meta 5-10% melhora testosterona 15%), manejo estresse (mindfulness 20min/dia, psicoterapia CBT), higiene do sono (polissonografia se roncos/apneia), controle glicêmico rigoroso (HbA1c <6,5%). Avaliar uso de pornografia (considerar "reset dopaminérgico" 30-90 dias abstinência).',
                'patientExplanation': 'Vários fatores podem prejudicar a função sexual: certos medicamentos, álcool, cigarro, excesso de peso, estresse, falta de sono e uso excessivo de pornografia. Identificar e modificar estes fatores geralmente traz melhora significativa. Sempre converse com seu médico antes de alterar medicações.'
            }

        # Histórico de Abuso/Trauma
        elif 'abuso' in name.lower() or 'trauma' in name.lower() or 'violência' in name.lower():
            return {
                'clinicalRelevance': 'Histórico de abuso sexual ou trauma impacta profundamente função sexual através de mecanismos neurobiológicos (desregulação eixo HPA, hiperativação amígdala), psicológicos (PTSD, dissociação, evitação) e somáticos (dor pélvica crônica, vaginismo). Requer abordagem compassiva, trauma-informada, multidisciplinar. Associa-se a níveis elevados de cortisol crônico, alterações em oxitocina, inflamação sistêmica aumentada (PCR elevado).',
                'conduct': 'ABORDAGEM SENSÍVEL: criar ambiente seguro, validar experiências, nunca pressionar. Encaminhar para psicoterapia especializada (EMDR, terapia focada no trauma, terapia sensório-motora). Avaliação médica compassiva: dor pélvica (fisioterapia pélvica especializada), cortisol salivar (4 pontos), DHEA-S, inflamação (PCR-us, homocisteína). Suporte nutricional: ômega-3 4g/dia (anti-inflamatório neural), magnésio 600mg/dia, complexo B, vitamina D, probióticos. Técnicas somáticas: yoga trauma-sensível, meditação mindfulness. Consentimento contínuo essencial.',
                'patientExplanation': 'Experiências de trauma podem afetar profundamente a saúde sexual e emocional. É importante trabalhar com profissionais especializados em trauma que podem ajudar através de terapias específicas e suporte nutricional. A cura é possível e você merece apoio compassivo e respeitoso nesta jornada.'
            }

        # Hormônios / Medicamentos
        elif 'hormônio' in name.lower() or 'medicamento' in name.lower() or 'remédio' in name.lower() or 'tratamento' in name.lower():
            return {
                'clinicalRelevance': 'Múltiplos hormônios e medicamentos impactam função sexual. Hormônios chave: testosterona (desejo em ambos sexos), estrogênio (lubrificação, vascularização), progesterona (excesso reduz libido), prolactina (inibe se elevada >25 ng/mL), TSH/T4 (meta-regulação). Medicamentos problemáticos: SSRIs/SNRIs (70% têm disfunção), beta-bloqueadores (25% DE), finasterida (síndrome pós-finasterida 5-10%), espironolactona (antiandrogênica), contraceptivos orais (aumentam SHBG 400%).',
                'conduct': 'Painel hormonal completo: testosterona total/livre (manhã 7-9h), SHBG, estradiol, progesterona (dia 21 se menstrua), prolactina, TSH, T4 livre, DHEA-S, cortisol salivar (4 pontos). Revisão medicamentosa criteriosa: trocar SSRI por bupropiona/mirtazapina, beta-bloqueador por nebivolol, considerar "drug holiday" SSRIs fim de semana. Se terapia testosterona: monitorar hematócrito, PSA, estradiol a cada 3 meses (usar anastrozol 0,25mg 2x/sem se estradiol >40 pg/mL). Suporte: L-arginina 5g/dia, zinco 50mg/dia, vitamina D, maca 3g/dia.',
                'patientExplanation': 'Hormônios e medicamentos podem afetar significativamente a função sexual. É importante investigar níveis hormonais e revisar medicações com seu médico. Muitas vezes é possível trocar por alternativas com menos efeitos sexuais, ou adicionar suplementos que ajudam a minimizar o impacto.'
            }

        # Histórico Reprodutivo / Gravidez
        elif 'gravidez' in name.lower() or 'filho' in name.lower() or 'gestação' in name.lower() or 'reprodutivo' in name.lower() or 'fertilidade' in name.lower():
            return {
                'clinicalRelevance': 'Histórico reprodutivo informa sobre fertilidade, função hormonal, complicações obstétricas que predizem riscos futuros (eclâmpsia e diabetes gestacional aumentam risco cardiovascular 2-4x). Infertilidade pode associar-se a endometriose (30% infertilidade), SOP (anovulação), insuficiência ovariana (FSH elevado), fatores masculinos (varicocele, qualidade espermática reduzida). Pós-parto: queda estrogênio/progesterona, prolactina elevada durante amamentação reduzem libido temporariamente (fisiológico).',
                'conduct': 'Histórico de infertilidade: investigar causas persistentes - SOP (insulina, HOMA-IR, testosterona, ultrassom), endometriose (CA-125, ultrassom especializado), reserva ovariana (AMH, FSH basal, contagem folicular antral), fator masculino (espermograma). Pós-parto com baixa libido: tranquilizar que é fisiológico durante amamentação (prolactina antagoniza testosterona), melhorar sono (apoio parceiro/família), fisioterapia pélvica se dor/trauma parto, suplementação (ferro, vitamina D, ômega-3, complexo B). Se diabetes gestacional prévio: TOTG 75g anual, estilo de vida anti-diabético (low-carb, exercício).',
                'patientExplanation': 'Histórico de gravidez e fertilidade fornece informações importantes sobre sua saúde hormonal e reprodutiva. Após o parto, é normal haver redução temporária de libido devido às mudanças hormonais e amamentação. Fisioterapia pélvica e suplementação adequada ajudam na recuperação. Se houve dificuldade para engravidar, vale investigar causas que podem ser tratadas.'
            }

        # Menopausa / Climatério
        elif 'menopausa' in name.lower() or 'climatério' in name.lower() or 'pós-menopausa' in name.lower():
            return {
                'clinicalRelevance': 'Menopausa (ausência menstruação 12 meses consecutivos) e perimenopausa caracterizam-se por declínio estrogênio/progesterona, resultando em atrofia vaginal (50-60% mulheres), redução lubrificação, dispareunia (dor coital 40%), diminuição libido (queda testosterona concomitante 50%). FSH>40 mUI/mL confirma menopausa. Sintomas: fogachos (75%), suores noturnos, insônia, alterações humor, perda massa óssea (osteopenia/osteoporose), aumento risco cardiovascular (perda efeito protetor estrogênio).',
                'conduct': 'Confirmar status: FSH, estradiol, AMH (reserva ovariana). Terapia hormonal bioidêntica se <60 anos e <10 anos pós-menopausa: estradiol transdérmico 0,5-1mg/dia + progesterona micronizada 100-200mg/noite (se útero intacto). DHEA vaginal 6,5mg ou estriol vaginal 0,5mg (atrofia) 3x/sem. Testosterona transdérmica 0,5mg/dia (baixa libido). Suplementos: isoflavonas 50mg/dia (fogachos), maca 3g/dia (libido), cimicifuga 40mg/dia, vitamina D 5000 UI + K2 100mcg, cálcio 500mg (ossos), magnésio 400mg, ômega-3 2g/dia. Densitometria óssea baseline.',
                'patientExplanation': 'Menopausa é uma fase natural que traz mudanças hormonais significativas, afetando vida sexual, ossos, humor e risco cardiovascular. Terapia hormonal bioidêntica, quando indicada e segura, pode melhorar muito a qualidade de vida. Suplementos naturais e exercícios também ajudam bastante com os sintomas.'
            }

        # Desenvolvimento Sexual / Puberdade
        elif 'desenvolvimento' in name.lower() or 'puberdade' in name.lower() or 'adolescência' in name.lower():
            return {
                'clinicalRelevance': 'Desenvolvimento sexual puberal depende de eixo HPG (GnRH → LH/FSH → esteroides gonadais). Puberdade normal: meninas 8-13 anos (telarca), meninos 9-14 anos (aumento testicular >4mL). Atraso (>13 anos meninas, >14 meninos) ou precocidade (<8 meninas, <9 meninos) requerem investigação. Fatores de impacto: obesidade (leptina aumentada, aromatase, puberdade precoce), desnutrição/exercício excessivo (atraso), doenças crônicas (Crohn, celíaca, fibrose cística), estresse psicológico.',
                'conduct': 'Atraso puberal: investigar LH, FSH, testosterona/estradiol, IGF-1, cariótipo (Turner 45X0, Klinefelter 47XXY), TSH, prolactina, ferritina, transglutaminase IgA (celíaca), ressonância sela túrcica (tumor?), idade óssea (raio-X mão esquerda). Precocidade: mesmos exames + ultrassom pélvico (meninas), GnRH stimulation test. Otimizar nutrição: zinco 25mg/dia, vitamina D (alvo 50ng/mL), calorias/proteínas adequadas, sono 9-10h, reduzir estresse. Encaminhar endocrinologia pediátrica se anormalidades (FSH/LH baixos = central, altos = gonadal).',
                'patientExplanation': 'O desenvolvimento sexual ocorre normalmente entre 8-13 anos em meninas e 9-14 em meninos. Atraso ou início muito precoce podem indicar questões hormonais ou nutricionais que precisam ser investigadas. Na maioria dos casos, otimização nutricional e sono ajudam. Quando necessário, tratamento hormonal pode ser indicado.'
            }

        # Dor durante Relação (Dispareunia)
        elif 'dor' in name.lower() and ('relação' in name.lower() or 'penetração' in name.lower() or 'coito' in name.lower()):
            return {
                'clinicalRelevance': 'Dispareunia (dor coital) tem causas múltiplas: atrofia vaginal (hipoestrogenismo, 50% pós-menopausa), vaginismo (contração involuntária assoalho pélvico, 15% mulheres), infecções (candidíase recorrente, vaginose bacteriana), endometriose (dor profunda), adenomiose, aderências pélvicas pós-cirúrgicas, síndrome dor pélvica crônica, vulvodínia, histórico trauma/abuso (componente psicológico). Abordagem requer exame ginecológico sensível, investigação causas orgânicas e psicológicas.',
                'conduct': 'Investigar: exame especular delicado (atrofia, lesões, eritema vulvar), swab vaginal (pH >4,5 = vaginose, microscopia, cultura), ultrassom transvaginal (endometriose, adenomiose, massas), estradiol, testosterona (atrofia). Tratamento por causa: atrofia (estrogênio/DHEA vaginal 3x/sem), vaginismo (fisioterapia pélvica especializada + dilatadores graduais + psicoterapia), infecções (tratamento específico: fluconazol, metronidazol), endometriose (laparoscopia diagnóstica/terapêutica, dienogeste). Suporte: fisioterapia assoalho pélvico (TODAS pacientes), terapia sexual, lubrificantes base água/silicone, ômega-3 2g/dia (anti-inflamatório), vitamina E vaginal 400 UI.',
                'patientExplanation': 'Dor durante relações sexuais não é normal e tem várias causas possíveis: desde atrofia vaginal (falta de estrogênio), infecções, endometriose ou tensão muscular. Investigação cuidadosa identifica a causa. Fisioterapia pélvica, tratamentos específicos e às vezes terapia sexual trazem excelente melhora.'
            }

        # Orgasmo / Anorgasmia
        elif 'orgasmo' in name.lower() or 'clímax' in name.lower():
            return {
                'clinicalRelevance': 'Anorgasmia (dificuldade/incapacidade atingir orgasmo) afeta 10-40% mulheres, 5-10% homens. Causas: medicamentos (SSRIs 30-70%, antipsicóticos), doenças neurológicas (diabetes - neuropatia autonômica, esclerose múltipla), cirurgias pélvicas (prostatectomia, histerectomia), hipotireoidismo, deficiência testosterona, fatores psicológicos (ansiedade de desempenho, histórico trauma), falta de estimulação adequada. Orgasmo envolve dopamina, oxitocina, endorfinas, integridade vias neurológicas.',
                'conduct': 'Investigar: testosterona total/livre, prolactina, TSH, glicemia/HbA1c (>6,5% risco neuropatia), revisar medicações. Se SSRI: considerar bupropiona 150mg adjuvante, sildenafil 50-100mg PRN 1h antes (mulheres off-label), ou "drug holiday" (suspender SSRI sexta-sábado). Suplementação: L-arginina 3-5g/dia, maca 3g/dia, ginkgo biloba 240mg/dia (especialmente para anorgasmia induzida por antidepressivos), zinco 50mg/dia. Abordagem: educação sexual, comunicação com parceiro (guiar estimulação), sexoterapia (sensate focus), dispositivos (vibradores terapêuticos), exercícios Kegel (50 contrações/dia), mindfulness sexual.',
                'patientExplanation': 'Dificuldade em atingir orgasmo pode estar relacionada a medicamentos (especialmente antidepressivos), hormônios, questões neurológicas ou falta de estimulação adequada. Investigação médica, ajustes de medicação, suplementos e orientação sexual costumam resolver. Exercícios do assoalho pélvico e comunicação com parceiro também são importantes.'
            }

        # Ejaculação Precoce
        elif 'ejaculação' in name.lower() or 'ejacular' in name.lower():
            return {
                'clinicalRelevance': 'Ejaculação precoce (EP) definida como ejaculação persistente <1-2 minutos após penetração com angústia/frustração. Prevalência 20-30% homens. Tipos: primária (sempre presente, frequentemente base serotonérgica/genética) ou adquirida (surgiu recentemente, investigar prostatite, hipertireoidismo, ansiedade, relacionamento). Controle ejaculatório envolve serotonina (inibe ejaculação), dopamina/noradrenalina (facilitam), testosterona, sensibilidade glande peniana (variável individual).',
                'conduct': 'Investigar: TSH (hipertireoidismo acelera tudo), testosterona (se baixa pode causar EP), exame prostático/PSA se >40 anos ou sintomas urinários (prostatite). Tratamento farmacológico: dapoxetina 30-60mg on-demand 1-3h antes (SSRI aprovado para EP), paroxetina 20mg/dia ou sertralina 50mg/dia (off-label, uso contínuo), anestésicos tópicos (lidocaína/prilocaína creme 20min antes, lavar antes penetração). Suplementação: magnésio 400mg/dia, zinco 50mg/dia, L-triptofano 1g/noite (precursor serotonina). Técnicas comportamentais: stop-start (parar quando próximo), squeeze (pressionar glande), exercícios Kegel masculinos (contrações 3x10/dia), mindfulness sexual, terapia sexual com parceira.',
                'patientExplanation': 'Ejaculação precoce é condição comum e tratável. Pode ser abordada com medicações específicas tomadas antes da relação ou diariamente, técnicas comportamentais e exercícios do assoalho pélvico. Comunicação com a parceira e terapia sexual também ajudam bastante. A maioria dos homens consegue melhora significativa.'
            }

        # Orientação Sexual / Identidade de Gênero
        elif 'orientação' in name.lower() or 'identidade de gênero' in name.lower() or 'transgênero' in name.lower() or 'LGBTQ' in name:
            return {
                'clinicalRelevance': 'Orientação sexual (atração afetivo-sexual) e identidade de gênero (autopercepção) são aspectos centrais da saúde sexual e bem-estar. Minorias sexuais e de gênero enfrentam estresse minoritário (discriminação, estigma, microagressões) associado a taxas 2-3x elevadas de ansiedade, depressão, ideação suicida, abuso de substâncias. Cuidado afirmativo, linguagem inclusiva, compreensão de necessidades específicas (PrEP, rastreio ISTs em múltiplos sítios anatômicos, saúde trans) são essenciais para qualidade assistencial.',
                'conduct': 'Criar ambiente acolhedor: usar pronomes corretos (perguntar preferência), formulários inclusivos (opções gênero/orientação), evitar presunções heteronormativas. Rastreios específicos: ISTs ampliado incluindo sítios extragenitais (faringe, reto) se HSH/práticas anais, HIV 3-6 meses se risco, saúde mental (PHQ-9, GAD-7, ideação suicida), uso de substâncias (chemsex). Se pessoa trans: acompanhamento hormonal regular (testosterona/estradiol/bloqueadores), densitometria óssea, monitorar hematócrito (testo), prolactina (estradiol), suporte multidisciplinar (endócrino, saúde mental, cirurgião se desejado). Recursos: encaminhar grupos apoio LGBTQIA+, terapia afirmativa. PrEP (tenofovir/emtricitabina) se exposição recorrente HIV.',
                'patientExplanation': 'Sua orientação sexual e identidade de gênero são aspectos importantes da sua saúde e bem-estar. Nosso objetivo é fornecer cuidado respeitoso, afirmativo e baseado em evidências, atendendo suas necessidades específicas de saúde. Você merece sentir-se seguro e acolhido no ambiente de saúde.'
            }

        # ISTs (Infecções Sexualmente Transmissíveis)
        elif 'IST' in name or 'DST' in name or 'infecção sexual' in name.lower() or 'doença sexual' in name.lower():
            return {
                'clinicalRelevance': 'ISTs (HIV, sífilis, gonorreia, clamídia, HPV, herpes HSV1/2, hepatites B/C) impactam saúde sexual, fertilidade (clamídia/gonorreia → doença inflamatória pélvica → infertilidade), bem-estar psicológico, e alguns têm potencial oncogênico (HPV → câncer colo/ânus/orofaringe, HBV → hepatocarcinoma). Rastreio regular é preventivo. Grupos risco elevado: múltiplos parceiros, HSH, PWID (usuários drogas injetáveis), profissionais do sexo, <25 anos sexualmente ativos. Muitas ISTs são assintomáticas (clamídia 70% mulheres, gonorreia 50% homens, HIV fase aguda pode ser flu-like inespecífico).',
                'conduct': 'Rastreio anual (ou 3-6 meses se alto risco): HIV 4ª geração (antígeno p24 + anticorpos, janela 2-4 semanas), VDRL/RPR + FTA-Abs/TPPA (sífilis), HBsAg + anti-HBs + anti-HBc, anti-HCV, gonorreia/clamídia NAAT (swab uretral/endocervical/retal/faríngeo conforme práticas sexuais). HPV: Papanicolau 21-65 anos, coteste (citologia + HPV) >30 anos, vacinação 9-valente até 26 anos (ideal 9-14 antes início atividade sexual). PrEP se exposição HIV recorrente: tenofovir/emtricitabina 300/200mg/dia ou on-demand (2cp 2-24h antes, 1cp 24h após, 1cp 48h após). Tratamento ISTs: ceftriaxona 500mg IM dose única (gonorreia) + azitromicina 1g VO (clamídia), penicilina benzatina (sífilis), notificação/tratamento parceiros.',
                'patientExplanation': 'ISTs são infecções comuns que podem afetar saúde sexual, fertilidade e até aumentar risco de câncer (HPV). Muitas são assintomáticas, por isso rastreio regular é importante. Uso consistente de preservativo, redução de parceiros, vacinação contra HPV e hepatite B, e se exposição frequente ao HIV, uso de PrEP (prevenção) são fundamentais.'
            }

        # Pornografia
        elif 'pornografia' in name.lower() or 'pornô' in name.lower():
            return {
                'clinicalRelevance': 'Uso excessivo de pornografia (>3x/semana ou >60min/sessão) pode levar a dessensibilização dopaminérgica (downregulation receptores D2), disfunção erétil induzida por pornografia (PIED - prevalência crescente homens <40 anos), ejaculação precoce ou retardada, expectativas irreais sobre sexo/corpos, redução satisfação sexual com parceiro real, relacionamentos afetados. Mecanismo: superestimulação centros de recompensa (nucleus accumbens), similar neurobiologia de vícios comportamentais (jogo, internet).',
                'conduct': 'Avaliação compassiva sem julgamento: frequência, tipo de conteúdo, impacto em função sexual/relacionamentos, tentativas prévias redução. Se PIED ou disfunção relacionada: "reboot" (abstinência pornografia + masturbação 60-90 dias, permite ressensibilização dopaminérgica), redução gradual, reintrodução sensações reais progressiva, bloqueadores de conteúdo se dificuldade autocontrole. Suporte: terapia cognitivo-comportamental (CBT), grupos apoio online (NoFap, r/pornfree), educação neurobiologia vício, tratamento comorbidades (ansiedade, depressão). Suplementação para restaurar função dopaminérgica: L-tirosina 1g/dia, mucuna pruriens 300mg/dia (L-DOPA natural), ômega-3 2g/dia, magnésio 400mg, zinco 50mg, exercício regular (aumenta receptores D2).',
                'patientExplanation': 'Uso excessivo de pornografia pode afetar sua resposta sexual com parceiros reais e relacionamentos, através de mudanças na química cerebral (dopamina). Reduzir ou parar temporariamente, junto com suporte psicológico e suplementação, costuma restaurar função sexual normal em 2-3 meses. Não há julgamento, é uma questão médica tratável.'
            }

        # Satisfação Sexual
        elif 'satisfação' in name.lower():
            return {
                'clinicalRelevance': 'Satisfação sexual é multidimensional: física (prazer, orgasmo, ausência dor), emocional (intimidade, conexão, sentir-se desejado), relacional (comunicação, reciprocidade, resolução conflitos), contextual (tempo adequado, privacidade, segurança). Correlaciona-se fortemente (r=0.6-0.7) com satisfação geral de relacionamento, saúde mental (menores taxas ansiedade/depressão), qualidade de vida. Fatores negativos: estresse crônico, fadiga, problemas comunicação, discrepância libido entre parceiros (30-40% casais), rotina/falta novidade, imagem corporal negativa.',
                'conduct': 'Avaliar múltiplos domínios: função sexual objetiva (usar escalas validadas FSFI/IIEF), satisfação relacional (CSI-32), saúde mental (PHQ-9 depressão, GAD-7 ansiedade), estressores vida (trabalho, filhos, finanças). Intervenções: terapia de casal focada sexualidade (Gottman Method, EFT), técnicas comunicação sexual (exercício "Yes, No, Maybe" lists), exercícios intimidade (sensate focus - focar sensações sem objetivo orgasmo), educação sexual (anatomia, fisiologia prazer), exploração preferências consensual. Otimizar saúde base: exercício regular (150min/sem), sono 7-9h, nutrição anti-inflamatória, reduzir álcool/substâncias. Suplementos: maca 3g/dia, L-arginina 3-5g/dia, ômega-3 2g/dia, magnésio 400mg/dia. Agendar intimidade (reduz barreira "esperar vontade espontânea").',
                'patientExplanation': 'Satisfação sexual envolve aspectos físicos, emocionais e do relacionamento. Não é só sobre orgasmo, mas conexão, comunicação e intimidade. Melhorar satisfação geralmente requer abordar múltiplos aspectos: saúde física (exercício, sono, nutrição), comunicação com parceiro, reduzir estresse, e às vezes terapia de casal. Qualidade é mais importante que quantidade.'
            }

        # Educação Sexual / Conhecimento
        elif 'educação' in name.lower() or 'conhecimento' in name.lower() or 'informação' in name.lower():
            return {
                'clinicalRelevance': 'Educação sexual abrangente (anatomia, fisiologia, consentimento, comunicação, prazer, diversidade, prevenção ISTs/gravidez) está associada a: menor risco ISTs e gravidez não planejada (30-50% redução), relações mais saudáveis e equitativas, maior satisfação sexual recíproca, redução estigma/vergonha em torno sexualidade, início sexual mais tardio e responsável. Lacunas educacionais são prevalentes devido a tabus culturais/religiosos, currículos escolares deficientes (abstinência-only ineficaz), desinformação internet.',
                'conduct': 'Fornecer educação baseada evidências, medicamente precisa, culturalmente sensível, inclusiva: anatomia sexual (clitóris tem 8000 terminações nervosas, apenas glande visível), ciclo resposta sexual (Masters & Johnson: excitação, platô, orgasmo, resolução), variações normalidade (tamanho, formato, timing), consentimento ativo e contínuo, comunicação assertiva sobre desejos/limites, prevenção ISTs (preservativo consistente, PrEP) e gravidez (múltiplos métodos), prazer mútuo e reciprocidade. Recursos confiáveis: Sociedade Brasileira de Sexualidade Humana (SBRASH), SIECUS, livros (Come as You Are - Emily Nagoski, Guide to Getting It On). Desmistificar mitos comuns (hímen "romper", tamanho pênis, orgasmo vaginal vs clitoriano). Encorajar comunicação aberta com parceiro. Se histórico trauma/vergonha: encaminhar sexoterapia.',
                'patientExplanation': 'Educação sexual adequada é fundamental para saúde sexual e relacionamentos saudáveis. Aprender sobre anatomia, comunicação, consentimento e prazer mútuo melhora satisfação e previne problemas. Não tenha vergonha de fazer perguntas - sexualidade é parte natural e saudável da vida humana, e informação correta empodera você a fazer escolhas conscientes.'
            }

        # Frequência Sexual
        elif 'frequência' in name.lower() or 'vezes' in name.lower() or 'quantas vezes' in name.lower():
            return {
                'clinicalRelevance': 'Frequência sexual varia amplamente entre indivíduos e relacionamentos (média casais adultos: 1-2x/semana, mas variação 0-7x/semana é normal), influenciada por: idade (declínio natural pós-30 anos), duração relacionamento (fase paixão inicial alta, depois normalização), saúde física/mental, níveis estresse, presença filhos pequenos (redução 40-60% primeiro ano pós-parto), medicamentos, disponibilidade tempo/energia. Estudos mostram frequência >1x/semana não aumenta necessariamente bem-estar (platô), mas satisfação com frequência é mais importante que número absoluto. Discrepância libido entre parceiros é queixa extremamente comum (30-40% casais).',
                'conduct': 'Explorar sem julgamento: satisfação com frequência atual (ambos parceiros), existência de discrepância desejo, barreiras específicas (tempo limitado, fadiga, dor, interesse baixo, estresse). Se desejo hipoativo: investigar causas médicas (hormônios - testosterona/estradiol/tireoide, medicamentos, condições crônicas), psicológicas (depressão, ansiedade, estresse crônico, trauma), relacionais (conflitos não-resolvidos, ressentimento, rotina). Intervenções: agendar intimidade propositalmente (remove barreira "esperar vontade espontânea", especialmente mulheres têm desejo mais responsivo), sensate focus (focar sensações sem pressão performance), terapia casal se discrepância significativa, otimizar energia (sono adequado, exercício regular, nutrição, reduzir estressores). Normalizar que qualidade > quantidade, e que desejo responsivo (surge durante atividade) é tão válido quanto espontâneo.',
                'patientExplanation': 'Não existe frequência "normal" de relações sexuais - varia muito entre pessoas e fases da vida. O importante é que ambos parceiros estejam satisfeitos. Se houver diferença grande de desejo, comunicação aberta, agendar momentos de intimidade e cuidar da saúde geral costumam ajudar. Qualidade da conexão é mais importante que quantidade de vezes.'
            }

        # Imagem Corporal / Autoestima
        elif 'imagem corporal' in name.lower() or 'autoestima' in name.lower() or 'aparência' in name.lower():
            return {
                'clinicalRelevance': 'Imagem corporal negativa afeta profundamente função sexual através de: evitação intimidade/exposição corporal, distração cognitiva durante sexo ("spectatoring" - observar-se vs. sentir prazer), redução excitação/lubrificação/ereção, dificuldade orgasmo, vergonha, ansiedade de desempenho. Fatores de risco: exposição mídia idealizada/redes sociais (comparação social), comentários críticos (parceiros, família, bullying histórico), transtornos alimentares (anorexia, bulimia, TANE), obesidade (estigma social). Mulheres são particularmente vulneráveis (pressão estética maior), mas homens também sofrem significativamente (ansiedade tamanho pênis, musculatura, calvície, gordura abdominal).',
                'conduct': 'Rastreio: Body Image-States Scale, Multidimensional Body-Self Relations Questionnaire, escala autoestima Rosenberg. Encaminhamento: psicoterapia (TCC - desafiar pensamentos automáticos negativos, ACT - aceitação/desfusão cognitiva, terapia focada compaixão), grupos suporte, sexoterapia se impacto sexual significativo. Técnicas específicas: mindfulness durante sexo (focar deliberadamente sensações físicas vs. aparência), exercícios mirror exposure (exposição gradual espelho sem julgamento), comunicação positiva com parceiro (verbalizar necessidade afirmação, pedir evitar críticas), desafiar padrões irreais mídia (awareness photoshop/filtros). Promover saúde via comportamentos saudáveis (exercício por saúde não aparência, alimentação equilibrada, sono) vs. busca estética inalcançável. Considerar impacto redes sociais (reduzir uso Instagram/TikTok se gatilho).',
                'patientExplanation': 'Preocupação excessiva com aparência durante sexo pode interferir muito no prazer e conexão. Seu corpo é capaz de sentir prazer independente de como você acha que se parece. Terapia, mindfulness e comunicação com parceiro ajudam a focar nas sensações e conexão ao invés de autocrítica. Seu valor não está na sua aparência.'
            }

        # Relacionamento com Parceiro
        elif 'parceiro' in name.lower() or 'relacionamento' in name.lower() or 'casal' in name.lower():
            return {
                'clinicalRelevance': 'Qualidade do relacionamento é preditor mais forte de satisfação sexual (r=0.65-0.75), superando fatores físicos isolados. Fatores positivos: comunicação aberta incluindo sobre sexo (capacidade discutir desejos, limites, problemas), intimidade emocional (vulnerabilidade, compartilhar sentimentos), resolução construtiva de conflitos (Gottman: reparação após brigas), tempo qualidade juntos (não só logística), gratidão/apreciação expressa, confiança. Fatores negativos: conflitos não-resolvidos crônicos, ressentimento acumulado, traição/infidelidade (quebra confiança), evitação comunicação ("stonewalling"), crítica/desprezo (Gottman: "4 cavaleiros apocalipse"), discrepância valores/objetivos vida. Duração relacionamento naturalmente reduz novidade (habituação, declínio dopamina), mas pode aumentar intimidade/segurança se bem cultivado.',
                'conduct': 'Avaliação relacional holística: satisfação geral relacionamento (CSI-32 Couples Satisfaction Index), padrões comunicação (Gottman Sound Relationship House), conflitos (frequência, intensidade, resolução), intimidade emocional/sexual, confiança/segurança apego, comprometimento. Ferramentas validadas: CSI-32, Dyadic Adjustment Scale, questionário estilos apego adulto (ansioso, evitativo, seguro). Encaminhamento terapia casal se: conflitos persistentes destrutivos, infidelidade (processo reparação possível com profissional), discrepância sexual significativa (libido, preferências), problemas comunicação crônicos. Técnicas baseadas evidências: Gottman Method (construir amor/intimidade, gerenciar conflito), EFT (Emotionally Focused Therapy - trabalhar apego), CBCT (terapia comportamental casal). Intervenções simples: date nights regulares sem falar trabalho/filhos (reconexão), introduzir novidade sexual consensual (ambientes novos, viagens, atividades diferentes), exercícios apreciação escrita diária (3 coisas que valoriza no parceiro), retiros de casal intensivos.',
                'patientExplanation': 'Qualidade do relacionamento é o fator mais importante para satisfação sexual. Comunicação aberta (incluindo sobre sexo), tempo de qualidade juntos, resolver conflitos de forma construtiva e expressar apreciação fortalecem conexão emocional e sexual. Se houver dificuldades persistentes, terapia de casal pode ajudar muito. Relacionamentos requerem cultivo ativo.'
            }

        # Sono e Sexualidade
        elif ('sono' in name.lower() or 'dormir' in name.lower()) and any(x in name.lower() for x in ['sexual', 'libido', 'desejo', 'função']):
            return {
                'clinicalRelevance': 'Sono adequado é essencial para função sexual saudável: testosterona tem pico durante sono REM (redução 10-15% por cada hora perdida <8h), regula cortisol (privação sono aumenta cortisol que antagoniza testosterona), melhora humor/energia/disposição. Privação sono crônica (<6h/noite) está associada a: redução libido (60-70% pessoas afetadas), disfunção erétil (homens com apneia têm 2-3x mais DE), dificuldade orgasmo, redução satisfação sexual, irritabilidade (afeta relacionamento). Apneia obstrutiva do sono afeta 40-70% homens com disfunção erétil, através de hipóxia intermitente (dano endotelial) e fragmentação sono (queda testosterona).',
                'conduct': 'Avaliar qualidade/duração sono: Pittsburgh Sleep Quality Index (PSQI), Epworth Sleepiness Scale, sintomas apneia (ronco alto, pausas respiratórias testemunhadas, sonolência diurna excessiva, acordar sufocado, noctúria >2x). Se apneia suspeita (Epworth >10, roncos + pausas): encaminhar para polissonografia, CPAP se IAH >15 (melhora DE em 60-70% casos). Higiene do sono: meta 7-9h consistentes, horários regulares deitar/acordar (inclusive fim de semana), ambiente escuro (cortinas blackout)/fresco (18-20°C)/silencioso, reduzir luz azul telas 2h antes (óculos bloqueadores ou apps f.lux), evitar cafeína após 14h, refeição leve noturna. Se insônia: TCC-I (Cognitive Behavioral Therapy for Insomnia, gold standard não-farmacológico), suplementos: melatonina 3-5mg 1h antes, magnésio glicinato 400mg/noite, L-teanina 200mg, glicina 3g, ashwagandha 600mg (reduz cortisol). Evitar benzodiazepínicos crônicos (tolerância, dependência, suprimem REM).',
                'patientExplanation': 'Sono adequado (7-9h) é fundamental para hormônios sexuais, energia e desejo. Falta de sono reduz testosterona, aumenta estresse e diminui libido. Se você ronca muito ou tem pausas respiratórias, pode ter apneia do sono, que afeta muito função sexual. Melhorar higiene do sono e tratar apneia (se presente) geralmente restaura função sexual.'
            }

        # Exercício e Sexualidade
        elif 'exercício' in name.lower() or 'atividade física' in name.lower():
            return {
                'clinicalRelevance': 'Exercício regular melhora função sexual através de múltiplos mecanismos: aumenta testosterona (exercício força aumenta 15-30% agudamente, crônico melhora baseline), melhora saúde vascular e fluxo sanguíneo (exercício aeróbico estimula óxido nítrico endotelial), reduz cortisol/estresse, melhora imagem corporal/autoestima, aumenta energia/resistência física, melhora sono. Estudos mostram exercício regular (≥150min/sem moderado) reduz risco disfunção erétil em 30-40% homens. Exercícios assoalho pélvico (Kegel) melhoram controle ejaculatório masculino e intensidade orgasmo (ambos sexos). CUIDADO: overtraining (>10h/sem alta intensidade sem recuperação adequada) pode suprimir eixo HPG e reduzir testosterona/libido (tríade atleta feminina: amenorreia, osteoporose, transtorno alimentar).',
                'conduct': 'Prescrição ótima para saúde sexual: AERÓBICO moderado 150min/sem ou vigoroso 75min/sem (caminhada rápida, corrida, ciclismo, natação, dança - melhora cardiovascular/óxido nítrico) + FORÇA 2-3x/sem focando grandes grupos musculares (agachamento, levantamento terra, supino, remada - aumenta testosterona, 3 séries 8-12 repetições) + EXERCÍCIOS KEGEL diários ambos sexos (identificar músculo: interromper jato urina, contrair esse músculo 5-10seg, relaxar 5-10seg, 3 séries 10 contrações/dia). Evitar overtraining: sinais (fadiga persistente, insônia paradoxal, baixa libido, irritabilidade, performance diminuída, infecções frequentes) - se presente, reduzir volume/intensidade 50%, priorizar recuperação/sono. Timing interessante: exercício leve-moderado 2-4h antes atividade sexual aumenta fluxo sanguíneo genital. Yoga específica melhora flexibilidade, consciência corporal, reduz ansiedade desempenho (Kundalini, tantra yoga).',
                'patientExplanation': 'Exercício regular é um dos melhores remédios naturais para função sexual: aumenta hormônios, melhora circulação, reduz estresse e aumenta confiança. Combine exercícios aeróbicos (caminhada, corrida, natação) com musculação e exercícios do assoalho pélvico (Kegel). Cuidado apenas para não exagerar - exercício demais pode ter efeito oposto.'
            }

        # Álcool e Sexualidade
        elif 'álcool' in name.lower() or 'bebida' in name.lower():
            return {
                'clinicalRelevance': 'Álcool tem efeito bifásico dose-dependente em função sexual: baixas doses (1-2 drinks, <0,05% BAC) reduzem inibições/ansiedade via GABA-érgico, podem facilitar relaxamento/desinibição temporários; doses moderadas-altas (>2 drinks) prejudicam significativamente excitação, ereção (vasoconstrição), lubrificação, orgasmo através de depressão SNC, supressão testosterona aguda, desidratação. Uso crônico excessivo causa: redução testosterona (aumento aromatase hepática converte testosterona→estrogênio), atrofia testicular, ginecomastia, neuropatia periférica (afeta sensibilidade genital), disfunção hepática (cirrose reduz metabolização estrogênio), dependência psicológica ("preciso álcool para relaxar sexualmente"). "Whiskey dick" (disfunção erétil alcoólica aguda) é fenômeno comum e bem documentado.',
                'conduct': 'Avaliar consumo: AUDIT-C (screening rápido 3 questões, score ≥4 homens/≥3 mulheres indica risco). Limites recomendados saúde: mulheres ≤1 dose/dia, homens ≤2 doses/dia (1 dose padrão = 350mL cerveja 5%, 150mL vinho 12%, 45mL destilado 40%). Se uso excessivo (>14 drinks/sem homens, >7 mulheres): intervenção breve motivacional (técnica FRAMES), encaminhar tratamento especializado dependência se AUDIT completo >15 ou sintomas abstinência. Abstinência temporária ("dry month" 30-90 dias) para avaliar impacto real em função sexual (geralmente melhora significativa). Suporte farmacológico para abstinência se dependência: naltrexona 50mg/dia (reduz craving bloqueando receptores opioides), acamprosato 666mg 3x/dia (reduz ansiedade abstinência), topiramato off-label, grupos apoio (AA, SMART Recovery). Suplementação para recuperação hepática/neurológica: NAC (N-acetilcisteína) 1200mg/dia (glutationa, detox hepático), milk thistle (silimarina) 300mg/dia, vitaminas B complexo (tiamina B1 100mg, ácido fólico 1mg, B12 1000mcg - prevenir neuropatia), magnésio 400mg, zinco 50mg.',
                'patientExplanation': 'Álcool em pequena quantidade pode relaxar, mas mais que 1-2 doses prejudica função sexual (dificulta ereção, lubrificação, orgasmo). Uso frequente e excessivo reduz hormônios sexuais e causa dependência psicológica. Se você percebe que precisa de álcool para relaxar sexualmente, ou bebe regularmente demais, vale tentar período sem beber e observar melhora. Tratamento para dependência está disponível se necessário.'
            }

        # DEFAULT: Conteúdo genérico estruturado baseado em medicina funcional
        else:
            return {
                'clinicalRelevance': f'Item relacionado a saúde sexual: {name}. A sexualidade é componente integral da saúde global, influenciada por fatores hormonais (testosterona, estrogênio, prolactina, hormônios tireoidianos), vasculares (produção óxido nítrico, saúde endotelial), neurológicos (dopamina, serotonina, oxitocina), psicológicos (autoestima, ansiedade, depressão, imagem corporal), relacionais (comunicação, intimidade, resolução conflitos) e comportamentais (sono, exercício, nutrição, uso substâncias). Abordagem deve ser compassiva, não-julgamental, culturalmente sensível e baseada em evidências científicas sólidas.',
                'conduct': 'Realizar avaliação holística biopsicossocial: histórico médico completo (condições crônicas, cirurgias prévias), medicamentos atuais (muitos têm efeitos sexuais: antidepressivos, anti-hipertensivos, antipsicóticos, finasterida), histórico psicossocial (relacionamento, saúde mental, trauma), exame físico apropriado. Investigações laboratoriais conforme indicação clínica: painel hormonal (testosterona total/livre manhã cedo, SHBG, estradiol, prolactina, TSH, T4 livre), perfil metabólico (glicemia, HbA1c, perfil lipídico, homocisteína), vitaminas/minerais (vitamina D alvo >50ng/mL, zinco sérico, magnésio RBC, B12). Encaminhamentos conforme necessário: sexoterapia para questões psicológicas/relacionais, urologista/ginecologista, endocrinologista, fisioterapeuta pélvico. Otimizar estilo de vida: exercício regular (aeróbico 150min/sem + força 2-3x/sem), sono 7-9h, nutrição anti-inflamatória (mediterrânea), redução estresse (mindfulness, yoga), limitar álcool (<2 doses/dia), cessação tabagismo. Suplementação básica saúde sexual: zinco 50mg/dia, vitamina D 5000 UI, magnésio 400mg, ômega-3 2g/dia, L-arginina 3-5g/dia.',
                'patientExplanation': 'A saúde sexual é parte importante do bem-estar geral e envolve aspectos físicos, emocionais e relacionais. Problemas nesta área podem ter várias causas, desde hormônios e medicamentos até questões emocionais ou de relacionamento. Com investigação adequada e abordagem personalizada, a maioria das dificuldades sexuais pode ser tratada com sucesso. Não tenha vergonha de discutir - é uma questão médica legítima e tratável.'
            }

    def update_item(self, item_id: str, enrichment: Dict) -> bool:
        """Atualizar item com conteúdo enriquecido"""
        try:
            payload = enrichment

            response = self.session.put(
                f"{API_BASE}/score-items/{item_id}",
                json=payload
            )
            response.raise_for_status()
            return True
        except Exception as e:
            print(f"✗ Erro ao atualizar item {item_id}: {e}")
            if hasattr(e, 'response') and e.response is not None:
                print(f"  Resposta: {e.response.text[:500]}")
            return False

    def process_batch(self):
        """Processar batch completo de 30 items"""
        print(f"\n{'='*80}")
        print("ENRIQUECIMENTO VIDA SEXUAL - 30 ITEMS (v2)")
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

            name = item.get('name', '')
            print(f"  Nome: {name[:80]}...")

            # Gerar enriquecimento
            enrichment = self.generate_enrichment(item)

            print(f"  Relevância clínica: {len(enrichment.get('clinicalRelevance', ''))} caracteres")
            print(f"  Conduta: {len(enrichment.get('conduct', ''))} caracteres")
            print(f"  Explicação paciente: {len(enrichment.get('patientExplanation', ''))} caracteres")

            # Atualizar
            if self.update_item(item_id, enrichment):
                print(f"  ✓ Item atualizado com sucesso")
                results['success'].append(item_id)
                time.sleep(0.5)  # Rate limiting gentil
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
        report_file = '/home/user/plenya/VIDA-SEXUAL-ENRICHMENT-COMPLETE.md'
        with open(report_file, 'w', encoding='utf-8') as f:
            f.write("# Relatório de Enriquecimento - VIDA SEXUAL (30 items)\n\n")
            f.write(f"**Data:** {time.strftime('%Y-%m-%d %H:%M:%S')}\n\n")
            f.write(f"## Resumo\n\n")
            f.write(f"- Total de items: {len(ITEM_IDS)}\n")
            f.write(f"- Sucesso: {len(results['success'])}\n")
            f.write(f"- Falhas: {len(results['failed'])}\n\n")

            f.write(f"## Metodologia\n\n")
            f.write(f"Enriquecimento baseado em medicina funcional com foco em:\n\n")
            f.write(f"- Escalas validadas (ASEX, FSFI, IIEF-5)\n")
            f.write(f"- Aspectos hormonais (testosterona, estrogênio, prolactina)\n")
            f.write(f"- Deficiências nutricionais (vitamina D, zinco, magnésio)\n")
            f.write(f"- Efeitos medicamentosos (SSRIs, beta-bloqueadores)\n")
            f.write(f"- Suplementação baseada em evidências\n")
            f.write(f"- Abordagem trauma-informada e compassiva\n")
            f.write(f"- Saúde LGBTQIA+ afirmativa\n\n")

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
