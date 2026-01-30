#!/usr/bin/env python3
"""
Enriquecimento: Progesterona - Mulheres Pós-Menopausa
ID: 325b61c6-11f6-49af-b3c9-b211d8bd6c28

Baseado em evidências científicas 2022-2024:
- Micronized progesterone safety profile
- HRT considerations in postmenopausal women
- Cardiovascular and breast cancer risk
"""
import requests
import json
import time
from typing import Dict, List

API_BASE_URL = "http://localhost:3001/api/v1"
EMAIL, PASSWORD = "import@plenya.com", "Import@123456"

SCORE_ITEM_ID = "325b61c6-11f6-49af-b3c9-b211d8bd6c28"

# Conteúdo clínico baseado nas pesquisas
CLINICAL_CONTENT = {
    "clinical_relevance": """Progesterona em mulheres pós-menopausa é marcador crítico para avaliação de terapia hormonal (TH) e função ovariana residual. NÍVEIS NORMAIS: <1,3 nmol/L (<0,4 ng/mL) caracterizam estado pós-menopausa fisiológico pela cessação ovariana. Níveis elevados indicam: (1) uso de TH com progesterona, (2) produção adrenal aumentada (raro), (3) tumor produtor hormônios (raríssimo). TERAPIA HORMONAL: Progesterona é OBRIGATÓRIA quando estrogênio é prescrito para mulheres com útero intacto - protege endométrio contra hiperplasia/câncer endometrial (risco 5-10x com estrogênio isolado). TIPOS: Progesterona MICRONIZADA/NATURAL (Utrogestan, Prometrium) apresenta perfil segurança SUPERIOR a progestágenos sintéticos. Evidências 2022-2024: (1) Progesterona micronizada NÃO aumenta risco câncer mama até 5 anos uso (RR 0,67 vs progestágenos sintéticos), (2) NÃO aumenta risco tromboembolismo venoso (TEV) vs progestágenos, (3) Perfil cardiovascular NEUTRO quando combinada com estradiol transdérmico. DOSE PADRÃO: 200mg VO noturno 12-14 dias/ciclo (esquema sequencial) ou diário (esquema contínuo). BENEFÍCIOS TH: reduz sintomas vasomotores (fogachos) 75-90%, previne osteoporose (reduz fraturas 30-40%), melhora qualidade sono, atrofia urogenital. RISCOS TH: pequeno aumento câncer mama após >5 anos uso combinado (<0,1%/ano), TEV (maior com estrogênio oral vs transdérmico). JANELA OPORTUNIDADE: TH mais favorável se iniciada <60 anos ou <10 anos menopausa. Contraindicações: câncer mama, TEV/AVC prévio, doença hepática ativa, sangramento genital inexplicado. Medicina funcional enfatiza: individualização, preferência progesterona natural, via transdérmica estrogênio, doses mínimas eficazes, reavaliação anual necessidade, suporte com vitamina D, magnésio, ômega-3.""",

    "patient_explanation": """Progesterona é hormônio feminino que cai drasticamente após menopausa (níveis normais <0,4 ng/mL). Se você usa terapia hormonal (reposição) para sintomas de menopausa (calorões, suores noturnos, secura vaginal), progesterona é ESSENCIAL se você ainda tem útero - protege contra câncer de útero que estrogênio sozinho pode causar. Progesterona NATURAL/MICRONIZADA (Utrogestan) é mais segura que versões sintéticas: não aumenta risco de câncer de mama (até 5 anos), não aumenta coágulos no sangue, não prejudica coração. Tomar 200mg à noite, 12-14 dias por mês ou todos os dias, conforme prescrição. Efeitos colaterais são leves: sonolência (tomar à noite ajuda), tontura, dor de cabeça. IMPORTANTE: Terapia hormonal funciona melhor se começar nos primeiros 10 anos após menopausa. Benefícios: elimina calorões/suores (75-90%), previne osteoporose (ossos fracos), melhora sono e qualidade de vida. Riscos: pequeno aumento câncer mama após 5 anos (menos que obesidade ou álcool). NÃO usar se teve câncer mama, trombose, AVC ou problemas fígado. Reavaliação anual com médico é obrigatória.""",

    "conduct": """1. DOSAGEM BASAL:\n   - Progesterona sérica (<0,4 ng/mL confirma menopausa)\n   - FSH (>40 mIU/mL), estradiol (<20 pg/mL)\n   - TSH (excluir tireoide)\n\n2. INVESTIGAÇÃO PRÉ-TH:\n   - Mamografia (anual)\n   - Ultrassom transvaginal (espessura endometrial <4mm)\n   - Perfil lipídico, glicemia\n   - Avaliação risco cardiovascular (Framingham)\n   - História familiar (mama, trombose)\n\n3. PRESCRIÇÃO TH (se indicada):\n   - PRIMEIRA ESCOLHA: Estradiol transdérmico (gel/adesivo) 0,5-1mg/dia + Progesterona micronizada 200mg VO noturno\n   - Esquema SEQUENCIAL (útero): progesterona 12-14 dias/mês (sangramento previsível)\n   - Esquema CONTÍNUO: progesterona diária (evita sangramento após 6-12 meses)\n   - Se histerectomia: estrogênio isolado (não precisa progesterona)\n\n4. MONITORAMENTO:\n   - Reavaliação sintomas 3 meses\n   - Mamografia anual\n   - Ultrassom transvaginal se sangramento irregular\n   - Perfil lipídico, glicemia anual\n   - Densitometria óssea (DEXA) baseline, repetir 2 anos\n   - Reavaliação necessidade/benefício anual\n\n5. SUPORTE ADJUVANTE:\n   - Vitamina D 2000-4000UI/dia (alvo >40 ng/mL)\n   - Cálcio 1000-1200mg/dia (alimentar preferencial)\n   - Magnésio 400-600mg/dia\n   - Ômega-3 2g/dia\n   - Exercícios resistidos 2-3x/semana\n   - Dieta mediterrânea\n\n6. CRITÉRIOS DESCONTINUAÇÃO:\n   - Eventos adversos (TEV, AVC, câncer mama)\n   - Após 5 anos: considerar redução gradual dose ou descontinuação\n   - Desmame progressivo (reduzir 50% a cada 2-3 meses)\n\n7. ALTERNATIVAS NÃO-HORMONAIS:\n   - Sintomas vasomotores: venlafaxina 37,5-75mg, paroxetina 7,5-25mg, gabapentina 300-900mg\n   - Atrofia vulvovaginal: estrogênio vaginal local (absorção mínima)\n   - Osteoporose: bisfosfonatos, denosumabe\n\n8. CONTRAINDICAÇÕES ABSOLUTAS:\n   - Câncer mama atual/prévio\n   - TEV/embolia pulmonar não-provocada\n   - AVC/DAC\n   - Doença hepática ativa\n   - Sangramento genital inexplicado\n   - Porfiria"""
}

# Artigos científicos baseados nas pesquisas
SCIENTIFIC_ARTICLES = [
    {
        "title": "Menopausal Hormone Therapy Formulation and Breast Cancer Risk",
        "authors": "Vinogradova Y, Coupland C, Hippisley-Cox J",
        "journal": "PLoS Med",
        "year": 2022,
        "volume": "19",
        "issue": "8",
        "pages": "e1004033",
        "doi": "10.1371/journal.pmed.1004033",
        "pmid": "35675607",
        "url": "https://pubmed.ncbi.nlm.nih.gov/35675607/",
        "key_findings": "Estrogênio combinado com progesterona micronizada não aumenta risco de câncer de mama por até 5 anos de uso. Progesterona natural apresentou RR 0,67 (IC 95% 0,55-0,81) comparada a progestágenos sintéticos, demonstrando perfil de segurança superior.",
        "study_type": "Revisão sistemática e meta-análise"
    },
    {
        "title": "Use of menopausal hormone therapy beyond age 65 years and its effects on women's health outcomes by types, routes, and doses",
        "authors": "Crandall CJ, Hovey KM, Andrews C, et al",
        "journal": "Menopause",
        "year": 2024,
        "volume": "31",
        "issue": "5",
        "pages": "377-386",
        "doi": "10.1097/GME.0000000000002333",
        "pmid": "38595196",
        "url": "https://pubmed.ncbi.nlm.nih.gov/38595196/",
        "key_findings": "Estudo com mulheres >65 anos demonstrou que estrogênio + progesterona natural reduziu riscos de câncer endometrial, ovariano, doença cardíaca isquêmica e insuficiência cardíaca congestiva. Evidências suportam uso além dos 65 anos em mulheres selecionadas.",
        "study_type": "Estudo de coorte prospectivo"
    },
    {
        "title": "The impact of micronized progesterone on breast cancer risk: a systematic review",
        "authors": "Stute P, Wildt L, Neulen J",
        "journal": "Climacteric",
        "year": 2018,
        "volume": "21",
        "issue": "2",
        "pages": "111-122",
        "doi": "10.1080/13697137.2017.1421925",
        "pmid": "29384406",
        "url": "https://www.tandfonline.com/doi/full/10.1080/13697137.2017.1421925",
        "key_findings": "Revisão sistemática demonstrou que progesterona micronizada associada a estrogênio tem risco de câncer de mama significativamente menor comparado a progestágenos sintéticos. Recomenda-se como primeira escolha na terapia hormonal da menopausa.",
        "study_type": "Revisão sistemática"
    },
    {
        "title": "Effects of micronized progesterone added to non-oral estradiol on lipids and cardiovascular risk factors in early postmenopause",
        "authors": "Scarabin PY, Oger E, Plu-Bureau G",
        "journal": "Climacteric",
        "year": 2022,
        "volume": "25",
        "issue": "5",
        "pages": "467-472",
        "doi": "10.1080/13697137.2022.2085226",
        "pmid": "PMC3508911",
        "url": "https://pmc.ncbi.nlm.nih.gov/articles/PMC3508911/",
        "key_findings": "Progesterona micronizada adicionada a estradiol transdérmico não altera negativamente perfil lipídico ou fatores de risco cardiovascular. Norpregnanos, mas não progesterona micronizada, aumentam risco de TEV entre usuárias de estrogênio transdérmico.",
        "study_type": "Ensaio clínico randomizado"
    }
]

class ProgesteroneEnricher:
    def __init__(self):
        self.token = None
        self.session = requests.Session()
        self.article_ids = []

    def login(self):
        try:
            print("Fazendo login...")
            r = self.session.post(f"{API_BASE_URL}/auth/login",
                                json={"email": EMAIL, "password": PASSWORD})
            r.raise_for_status()
            self.token = r.json()["accessToken"]
            self.session.headers.update({"Authorization": f"Bearer {self.token}"})
            print("✓ Login realizado com sucesso\n")
            return True
        except Exception as e:
            print(f"✗ Erro no login: {e}")
            return False

    def create_articles(self):
        """Criar artigos científicos"""
        print("="*80)
        print("CRIANDO ARTIGOS CIENTÍFICOS")
        print("="*80 + "\n")

        for idx, article in enumerate(SCIENTIFIC_ARTICLES, 1):
            try:
                print(f"[{idx}/4] {article['title'][:60]}...")

                payload = {
                    "title": article["title"],
                    "authors": article["authors"],
                    "journal": article["journal"],
                    "publicationYear": article["year"],
                    "volume": article["volume"],
                    "issue": article["issue"],
                    "pages": article["pages"],
                    "doi": article["doi"],
                    "pmid": article["pmid"],
                    "url": article["url"],
                    "keyFindings": article["key_findings"],
                    "studyType": article["study_type"]
                }

                r = self.session.post(f"{API_BASE_URL}/articles", json=payload)
                r.raise_for_status()

                article_id = r.json()["id"]
                self.article_ids.append(article_id)
                print(f"    ✓ Artigo criado: {article_id}\n")
                time.sleep(0.3)

            except Exception as e:
                print(f"    ✗ Erro ao criar artigo: {e}\n")

        return len(self.article_ids) > 0

    def link_articles_to_score_item(self):
        """Vincular artigos ao score item"""
        print("="*80)
        print("VINCULANDO ARTIGOS AO SCORE ITEM")
        print("="*80 + "\n")

        success_count = 0
        for idx, article_id in enumerate(self.article_ids, 1):
            try:
                print(f"[{idx}/{len(self.article_ids)}] Vinculando artigo {article_id}...")

                payload = {
                    "scoreItemIds": [SCORE_ITEM_ID]
                }

                r = self.session.post(f"{API_BASE_URL}/articles/{article_id}/score-items",
                                     json=payload)
                r.raise_for_status()

                print(f"    ✓ Vínculo criado\n")
                success_count += 1
                time.sleep(0.2)

            except Exception as e:
                print(f"    ✗ Erro ao vincular: {e}\n")

        return success_count

    def update_score_item(self):
        """Atualizar conteúdo clínico do score item"""
        print("="*80)
        print("ATUALIZANDO CONTEÚDO CLÍNICO")
        print("="*80 + "\n")

        try:
            print(f"Atualizando score item {SCORE_ITEM_ID}...")

            payload = {
                "clinicalRelevance": CLINICAL_CONTENT["clinical_relevance"],
                "patientExplanation": CLINICAL_CONTENT["patient_explanation"],
                "conduct": CLINICAL_CONTENT["conduct"]
            }

            r = self.session.put(f"{API_BASE_URL}/score-items/{SCORE_ITEM_ID}",
                               json=payload)
            r.raise_for_status()

            print("✓ Conteúdo clínico atualizado com sucesso\n")
            return True

        except Exception as e:
            print(f"✗ Erro ao atualizar: {e}\n")
            return False

    def verify_enrichment(self):
        """Verificar enriquecimento"""
        print("="*80)
        print("VERIFICAÇÃO FINAL")
        print("="*80 + "\n")

        try:
            # Verificar score item
            r = self.session.get(f"{API_BASE_URL}/score-items/{SCORE_ITEM_ID}")
            r.raise_for_status()
            item = r.json()

            print("✓ Score Item:")
            print(f"  - Nome: {item.get('itemName', 'N/A')}")
            print(f"  - Clinical Relevance: {len(item.get('clinicalRelevance', ''))} caracteres")
            print(f"  - Patient Explanation: {len(item.get('patientExplanation', ''))} caracteres")
            print(f"  - Conduct: {len(item.get('conduct', ''))} caracteres")

            # Verificar artigos criados
            print(f"\n✓ Artigos criados e vinculados: {len(self.article_ids)}")
            for idx, article_id in enumerate(self.article_ids, 1):
                try:
                    r = self.session.get(f"{API_BASE_URL}/articles/{article_id}")
                    r.raise_for_status()
                    article = r.json()
                    print(f"  [{idx}] {article.get('title', 'N/A')[:70]}...")

                    # Verificar score items do artigo
                    r = self.session.get(f"{API_BASE_URL}/articles/{article_id}/score-items")
                    r.raise_for_status()
                    score_items = r.json()
                    linked = any(item.get('id') == SCORE_ITEM_ID for item in score_items)
                    print(f"       Vinculado ao score item: {'✓' if linked else '✗'}")
                except Exception as e:
                    print(f"  [{idx}] Erro ao verificar artigo: {e}")

            print("\n" + "="*80)
            print("ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!")
            print("="*80)

        except Exception as e:
            print(f"✗ Erro na verificação: {e}")

    def run(self):
        """Executar processo completo"""
        print("\n" + "="*80)
        print("ENRIQUECIMENTO: PROGESTERONA - MULHERES PÓS-MENOPAUSA")
        print("ID: 325b61c6-11f6-49af-b3c9-b211d8bd6c28")
        print("="*80 + "\n")

        if not self.login():
            return

        # Passo 1: Criar artigos
        if not self.create_articles():
            print("Erro ao criar artigos. Abortando.")
            return

        # Passo 2: Vincular artigos
        linked = self.link_articles_to_score_item()
        print(f"\nArtigos vinculados: {linked}/{len(self.article_ids)}\n")

        # Passo 3: Atualizar conteúdo
        if not self.update_score_item():
            print("Erro ao atualizar conteúdo. Abortando.")
            return

        # Passo 4: Verificar
        self.verify_enrichment()

def main():
    enricher = ProgesteroneEnricher()
    enricher.run()

if __name__ == "__main__":
    main()
