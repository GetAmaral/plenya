#!/usr/bin/env python3
"""
Migra genes que estão como GRUPOS para ITEMS nos subgrupos apropriados de Genética
"""

# Mapeamento manual: gene → subgrupo
GENE_TO_SUBGROUP = {
    # METABOLISMO
    'Metabolismo': [
        'MTHFR C677T rs1801133 (Homocisteína)',
        'FTO rs9939609 (Obesidade)',
        'MC4R rs17782313 (Obesidade)',
        'TCF7L2 rs7903146 (Diabetes)',
        'PPARG Pro12Ala rs1801282 (Diabetes)',
        'VDR FokI rs2228570 (Vitamina D)',
        'ABCC8 rs757110 (Diabetes)',
        'CDKAL1 rs7754840 (Diabetes)',
        'HHEX rs1111875 (Diabetes)',
        'IGF2BP2 rs4402960 (Diabetes)',
        'INS VNTR (Diabetes Tipo 1)',
        'IRS1 rs1801278 (Resistência Insulina)',
        'KCNJ11 rs5219 (Diabetes)',
        'SLC30A8 rs13266634 (Diabetes)',
        'PPARGC1A rs8192678 Gly482Ser (Metabolismo)',
        'LEPR rs1137101 Gln223Arg (Obesidade)',
        'POMC rs6713532 (Obesidade)',
        'PPARA rs4253778 (Resistência)',
        'GCK (MODY2)',
        'HNF1A (MODY3)',
        'HNF1B (MODY5)',
        'HNF4A (MODY1)',
        'BCO1 rs6564851 (Vitamina A)',
        'FABP2 Ala54Thr rs1799883 (Gordura)',
        'FADS1 rs174546 (Ômega-3)',
        'FADS2 rs174575 (Ômega-3/DHA)',
        'MCM6 rs4988235 (Lactose)',
        'SLC23A1 rs33972313 (Vitamina C)',
    ],
    # CARDIOVASCULAR
    'Cardiovascular': [
        'APOE (Alzheimer e Lipídios)',
        'ACE I/D rs4646994 (Hipertensão)',
        'AGT rs699 M235T (Hipertensão)',
        'AGTR1 rs5186 A1166C (Hipertensão)',
        'NOS3 rs1799983 Glu298Asp (Hipertensão)',
        'ADD1 rs4961 Gly460Trp (Hipertensão)',
        'CYP11B2 rs1799998 (Hipertensão)',
        'GNB3 rs5443 C825T (Hipertensão)',
        'LDLR rs688 (Colesterol LDL)',
        'PCSK9 rs11591147 R46L (Colesterol)',
        'ABCA1 rs9282541 (HDL)',
        'APOA1 rs670 (HDL)',
        'APOA5 rs662799 (Triglicerídeos)',
        'LCAT rs5923 (HDL)',
        'LIPC rs1800588 (HDL)',
        'LPL rs328 (Triglicerídeos)',
        'APOC2 (Deficiência)',
        'GPIHBP1 (Quilomicronemia)',
        'LMF1 (Hipertrigliceridemia)',
    ],
    # NEURODEGENERAÇÃO
    'Neurodegeneração': [
        'APP A673T rs63750847 (Alzheimer)',
        'PSEN1 E280A (Alzheimer Familial)',
        'PSEN2 (Alzheimer Familial)',
        'LRRK2 G2019S rs34637584 (Parkinson)',
        'SNCA rs356219 (Parkinson)',
        'PARK2 (Parkinson Precoce)',
        'PARK7/DJ-1 (Parkinson Precoce)',
        'PINK1 (Parkinson Precoce)',
        'C9orf72 (DFT/ELA)',
        'MAPT rs1467967 H1/H2 (Demência FTD)',
        'GRN rs5848 (Demência FTD)',
    ],
    # DETOXIFICAÇÃO
    'Detoxificação': [
        'CYP1A1 rs4646903 MspI (Detoxificação)',
        'CYP1A2 rs762551 (Cafeína)',
        'CYP2A6 rs1801272 (Nicotina)',
        'GSTM1 (Detoxificação)',
        'GSTP1 rs1695 Ile105Val (Detoxificação)',
        'GSTT1 (Detoxificação)',
        'NAT2 (Acetilador)',
        'EPHX1 rs1051740 Tyr113His (Detoxificação)',
        'ADH1B rs1229984 Arg48His (Álcool)',
        'ALDH2 rs671 Glu504Lys (Álcool)',
        'GPX1 rs1050450 Pro198Leu (Glutationa)',
        'SOD2 rs4880 Ala16Val (Antioxidante)',
        'CAT rs1001179 -262C>T (Catalase)',
    ],
    # IMUNIDADE
    'Imunidade': [
        'HLA-DQ2/DQ8 (Doença Celíaca)',
        'IL1B rs16944 -511C>T (Inflamação)',
        'IL6 rs1800795 -174G>C (Inflamação)',
        'TNF rs1800629 -308G>A (Inflamação)',
        'CRP rs1130864 (Proteína C Reativa)',
    ],
    # PERFORMANCE
    'Performance': [
        'ACTN3 rs1815739 R577X (Performance)',
        'COL5A1 rs12722 (Lesão Tendão)',
        'COL1A1 rs1800012 Sp1 (Osteoporose)',
        'ESR1 rs2234693 PvuII (Osteoporose)',
    ],
    # OUTROS
    'Outros': [
        'ALPL (Hipofosfatasia)',
    ],
}

print("# MAPEAMENTO DE GENES → SUBGRUPOS")
print()
for subgroup, genes in GENE_TO_SUBGROUP.items():
    print(f"## {subgroup} ({len(genes)} genes)")
    for gene in genes:
        print(f"  - {gene}")
    print()

total_genes = sum(len(genes) for genes in GENE_TO_SUBGROUP.values())
print(f"**Total mapeado:** {total_genes} genes")
print()
print("⚠️ Nota: Migração automática requer script SQL complexo.")
print("Sugestão: Fazer manualmente via interface ou criar script dedicado.")
