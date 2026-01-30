#!/usr/bin/env python3
"""
Verify character counts for LH - Homens enrichment
"""

# Clinical content extracted from SQL
clinical_relevance = """O hormônio luteinizante (LH) é uma gonadotrofina hipofisária essencial para a função reprodutiva masculina, atuando nas células de Leydig testiculares para estimular a produção de testosterona. A dosagem sérica de LH é fundamental na investigação diagnóstica do hipogonadismo masculino, permitindo diferenciar entre hipogonadismo primário (testicular) e secundário (hipotálamo-hipofisário). Valores de referência normais para homens adultos situam-se entre 1,5-9,3 mIU/mL, embora possam variar entre laboratórios. Níveis elevados de LH (>9,3 mIU/mL) associados a testosterona baixa caracterizam hipogonadismo hipergonadotrófico primário, indicando falência testicular, enquanto LH baixo ou inapropriadamente normal (<1,5 mIU/mL) com testosterona reduzida sugere hipogonadismo hipogonadotrófico secundário, de origem central. As diretrizes internacionais de 2024 da Endocrine Society e European Association of Urology recomendam fortemente a dosagem de LH e FSH em todos os homens com testosterona baixa confirmada, classificação de evidência A. A relação LH/testosterona possui valor diagnóstico superior ao LH isolado, sendo particularmente útil na identificação de disfunção subclínica de células de Leydig. Condições clínicas associadas incluem síndrome de Kallmann (hipogonadismo congênito com anosmia e LH muito baixo), tumores hipofisários, hiperprolactinemia, uso de esteroides anabolizantes (supressão do eixo), insuficiência renal crônica, obesidade mórbida e envelhecimento masculino. O LH elevado constitui marcador inequívoco de disfunção testicular primária, excluindo causas fisiológicas de testosterona baixa transitória."""

patient_explanation = """O LH (hormônio luteinizante) é produzido pela hipófise, uma glândula localizada na base do cérebro, e funciona como um mensageiro que ordena aos testículos a produção de testosterona. Quando seus testículos não estão funcionando bem, a hipófise tenta compensar aumentando muito o LH, como se estivesse "gritando mais alto" para tentar obter resposta. Por outro lado, quando o problema está na própria hipófise ou no hipotálamo cerebral, o LH fica baixo porque o "comando central" não está enviando o sinal adequado. O exame de LH é essencial para o médico descobrir onde está o problema: se nos testículos (LH alto) ou no cérebro/hipófise (LH baixo). Valores normais ficam entre 1,5 e 9,3 mIU/mL. Este exame é especialmente importante quando você apresenta sintomas como diminuição da libido, fadiga, perda de massa muscular, infertilidade ou disfunção erétil. O resultado do LH sempre deve ser interpretado junto com o nível de testosterona total e livre. Em homens jovens com LH muito baixo, pode indicar síndrome de Kallmann, uma condição genética rara. Se você usa anabolizantes, o LH ficará muito baixo porque o cérebro entende que não precisa mais estimular os testículos, podendo causar atrofia testicular permanente."""

conduct = """A conduta clínica baseia-se na interpretação conjunta dos níveis de LH, testosterona total, testosterona livre e FSH, conforme algoritmo diagnóstico das diretrizes BSSM 2023 e AUA 2024. HIPOGONADISMO PRIMÁRIO (LH elevado >9,3 mIU/mL + testosterona baixa <300 ng/dL): investigar cariótipo para síndrome de Klinefelter (47,XXY), história de criptorquidia, orquite viral (caxumba), trauma testicular, torção, radiação ou quimioterapia. Tratamento: reposição de testosterona (gel, injeções ou adesivos) é a única opção eficaz, pois clomifeno não responde quando LH já está elevado. Monitorar hemograma (policitemia), PSA, função hepática e densidade óssea. HIPOGONADISMO SECUNDÁRIO (LH baixo/normal <4 mIU/mL + testosterona baixa): dosagem de prolactina sérica obrigatória; se elevada (>20 ng/mL), solicitar ressonância magnética de sela túrcica para investigar prolactinoma ou outros tumores hipofisários. Avaliar uso de medicações supressoras (opioides, glicocorticoides, antipsicóticos). Em jovens com LH <1 mIU/mL, considerar síndrome de Kallmann e solicitar teste olfatório e ressonância magnética hipofisária. Tratamento preservador de fertilidade: gonadotrofinas (hCG 1000-2000 UI 2-3x/semana ± FSH recombinante 75-150 UI 3x/semana) são superiores à testosterona exógena quando desejo de paternidade. Estudo de 2023 demonstrou que terapia combinada hCG+HMG desde o início alcança melhores taxas de espermatogênese que monoterapia. Fatores de mau prognóstico: IMC >30 kg/m², volume testicular inicial <5 mL, duração de tratamento <13 meses. VALORES LIMÍTROFES OU DISCORDANTES: repetir dosagens em 2-4 semanas, coletadas entre 7-11h da manhã (ritmo circadiano do LH). Solicitar LH com testosterona livre (diálise de equilíbrio ou espectrometria de massa), SHBG e albumina. Se testosterona total entre 250-350 ng/dL com LH normal-alto, calcular índice androgênico livre. SEGUIMENTO: reavaliar LH, testosterona e sintomas a cada 3-6 meses no primeiro ano de tratamento, depois anualmente. Homens em reposição de testosterona terão LH suprimido (<0,5 mIU/mL) por feedback negativo, o que é esperado e não requer intervenção."""

# Count characters
relevance_len = len(clinical_relevance)
explanation_len = len(patient_explanation)
conduct_len = len(conduct)

print("=" * 70)
print("LH - Homens Content Verification")
print("=" * 70)
print()

print(f"Clinical Relevance: {relevance_len} characters")
print(f"  Target range: 1500-2000 characters")
print(f"  Status: {'✓ PASS' if 1500 <= relevance_len <= 2000 else '✗ FAIL'}")
print()

print(f"Patient Explanation: {explanation_len} characters")
print(f"  Target range: 1000-1500 characters")
print(f"  Status: {'✓ PASS' if 1000 <= explanation_len <= 1500 else '✗ FAIL'}")
print()

print(f"Conduct: {conduct_len} characters")
print(f"  Target range: 1500-2500 characters")
print(f"  Status: {'✓ PASS' if 1500 <= conduct_len <= 2500 else '✗ FAIL'}")
print()

print("=" * 70)
print("Scientific Articles")
print("=" * 70)
articles = [
    {"pmid": "29562364", "year": 2018, "journal": "J Clin Endocrinol Metab"},
    {"pmid": "36876744", "year": 2023, "journal": "World J Mens Health"},
    {"pmid": "37007338", "year": 2023, "journal": "Cureus"},
    {"pmid": "30855798", "year": 2024, "journal": "StatPearls"}
]

for i, article in enumerate(articles, 1):
    print(f"{i}. PMID: {article['pmid']} ({article['year']}) - {article['journal']}")

print()
print(f"Total articles: {len(articles)}")
print()

# Overall status
all_pass = (
    1500 <= relevance_len <= 2000 and
    1000 <= explanation_len <= 1500 and
    1500 <= conduct_len <= 2500
)

print("=" * 70)
if all_pass:
    print("✓ ALL CHECKS PASSED - Ready for database insertion")
else:
    print("✗ SOME CHECKS FAILED - Review content before insertion")
print("=" * 70)
