# Enriquecimento de Score Items - Genética (Batch 1)

**Data:** 2026-01-26
**Processador:** Claude Sonnet 4.5
**Total de genes processados:** 6 de 20 (30% completo)

## Progresso

### Genes Completados (6/20)

1. **FTO rs9939609 (Obesidade)** - ID: da4da973-821f-4859-bc4f-c913a8b0b8d3
   - Clinical relevance: 380 palavras
   - Patient explanation: 165 palavras
   - Conduct: 310 palavras
   - Fontes: PubMed, Nutrigenomics reviews

2. **MC4R rs17782313 (Obesidade)** - ID: de659fee-1ab3-437f-b971-a3b4e251a599
   - Clinical relevance: 395 palavras
   - Patient explanation: 155 palavras
   - Conduct: 425 palavras
   - Fontes: MC4R polymorphisms research, Melanocortin pathway

3. **TCF7L2 rs7903146 (Diabetes)** - ID: a979d8f3-7e94-4ad9-bc86-d120765eef2f
   - Clinical relevance: 398 palavras
   - Patient explanation: 168 palavras
   - Conduct: 550 palavras
   - Fontes: TCF7L2 diabetes genetics, GLP-1 response studies

4. **PPARG Pro12Ala rs1801282 (Diabetes)** - ID: 944287b2-de63-4b77-8260-551db06da9bb
   - Clinical relevance: 385 palavras
   - Patient explanation: 185 palavras
   - Conduct: 520 palavras
   - Fontes científicas: MDPI, Nature Scientific Reports, PubMed
   - Evidências de interação gene-nutriente (vitamina E)

5. **MTHFR C677T rs1801133 (Homocisteína)** - ID: 93ffb2ea-0a41-428c-b481-13f55a321e85
   - Clinical relevance: 420 palavras
   - Patient explanation: 185 palavras
   - Conduct: 780 palavras (extenso - incluindo orientações CDC/ACMG)
   - Fontes científicas: CDC, RACGP, MDPI
   - Ênfase: teste NÃO recomendado rotineiramente

6. **APOE (Alzheimer e Lipídios)** - ID: 160b5986-1045-4ae7-85c7-8dfe058c4df7
   - Clinical relevance: 445 palavras
   - Patient explanation: 230 palavras
   - Conduct: 980 palavras (mais extenso - protocolo preventivo completo)
   - Fontes científicas: Springer Nature, MIT Picower Institute
   - Intervenções nutrigenômicas: colina, DHA, dieta mediterrânea

### Genes Pendentes (14/20)

7. KCNJ11 rs5219 (Diabetes)
8. IRS1 rs1801278 (Resistência Insulina)
9. SLC30A8 rs13266634 (Diabetes)
10. LDLR rs688 (Colesterol LDL)
11. PCSK9 rs11591147 R46L (Colesterol)
12. FADS1 rs174546 (Ômega-3)
13. FADS2 rs174575 (Ômega-3/DHA)
14. FABP2 Ala54Thr rs1799883 (Gordura)
15. ACE I/D rs4646994 (Hipertensão)
16. AGT rs699 M235T (Hipertensão)
17. NOS3 rs1799983 Glu298Asp (Hipertensão)
18. MCM6 rs4988235 (Lactose)
19. ADH1B rs1229984 Arg48His (Álcool)
20. ALDH2 rs671 Glu504Lys (Álcool)

## Evidências Coletadas

### Artigos já em base MFI (247 total)
- Apenas 1 article diretamente sobre nutrição genética identificado
- Necessário buscar evidências externas (PubMed, Scholar)

### Buscas Online Realizadas
1. PPARG Pro12Ala diabetes nutrigenomics 2024-2025
2. MTHFR C677T homocysteine folate clinical implications
3. APOE epsilon4 Alzheimer lipids cardiovascular nutrigenomics
4. FADS1 FADS2 omega-3 conversion genetics 2024
5. ACE gene polymorphism hypertension salt sensitivity

### Fontes Científicas Incorporadas
- MDPI Biomolecules (2023)
- Nature Scientific Reports (2020)
- PubMed reviews
- CDC guidelines
- RACGP Australian guidelines
- Springer Nature Molecular Neurodegeneration (2022)
- MIT Picower Institute (2023)
- Genes & Nutrition (2024)
- Frontiers in Nutrition (2023-2024)
- American Heart Association journals

## Padrão de Qualidade

### Clinical Relevance (200-400 palavras)
- Mecanismo molecular do polimorfismo
- Impacto funcional na fisiologia
- Associações clínicas baseadas em evidências
- Magnitude do efeito (odds ratios, riscos relativos)
- Interações gene-ambiente
- Contexto de medicina funcional

### Patient Explanation (100-200 palavras)
- Linguagem acessível, sem jargões técnicos
- Explicação do que o gene faz
- Possíveis genótipos e seus significados
- Ênfase: gene NÃO é destino
- Fatores ambientais são modificadores importantes
- Tom empoderador, não alarmista

### Conduct (150-300 palavras, alguns >500 se complexo)
- Quando testar (indicações clínicas)
- Interpretação por genótipo
- Intervenções nutrigenômicas personalizadas
- Suplementação baseada em evidências
- Monitoramento laboratorial
- Metas terapêuticas específicas
- Considerações farmacogenômicas quando relevante
- Evitar determinismo genético

## Princípios de Nutrigenômica Aplicados

1. **Não-determinismo:** Gene aumenta risco, não determina destino
2. **Modificabilidade:** Intervenções ambientais podem atenuar risco genético
3. **Personalização:** Genótipo informa estratégias individualizadas
4. **Evidências:** Recomendações baseadas em estudos científicos
5. **Integração:** Genótipo é UMA das ferramentas, não a única
6. **Empoderamento:** Informação genética como ferramenta de ação
7. **Segurança:** Evitar testes desnecessários (ex: MTHFR como screening)

## Desafios Identificados

1. **Base de articles limitada:** Apenas 1 article MFI sobre genética
   - Solução: Buscar evidências em PubMed, Google Scholar

2. **Estrutura do banco:** Genes importados como "grupos" ao invés de items
   - Impacto: Precisará de ajuste na estrutura após enriquecimento

3. **Acesso à API bloqueado:** Bash permission denied para chamadas API
   - Solução: Gerar JSON para importação manual posterior

4. **Volume de conteúdo:** Cada gene requer 3-5 horas de pesquisa e redação
   - Estimativa: 60-100 horas para completar 80 genes
   - Recomendação: Processar em batches de 10-20

## Próximos Passos

1. Completar genes 7-20 (14 genes restantes)
2. Gerar arquivo consolidado JSON com todos os 20 genes
3. Criar script de importação para atualizar banco
4. Linkar articles relevantes quando disponíveis
5. Validar textos com revisão médica
6. Processar próximo batch (genes 21-40)

## Métricas de Qualidade

- **Média de palavras clinical_relevance:** 404 palavras
- **Média de palavras patient_explanation:** 181 palavras
- **Média de palavras conduct:** 594 palavras
- **Total de fontes citadas:** 12 artigos científicos
- **Genes com intervenções nutrigenômicas específicas:** 6/6 (100%)
- **Genes com fontes científicas 2023-2024:** 4/6 (67%)

## Arquivo de Saída

`/home/user/plenya/genetic_enrichment_batch1.json`

Contém array JSON com 6 objetos completos, prontos para importação via API PUT /api/v1/score-items/{id}
