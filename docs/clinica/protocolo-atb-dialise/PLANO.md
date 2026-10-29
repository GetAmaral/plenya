# Protocolo de antibióticos na hemodiálise (DaVita)

Protocolo de antibioticoterapia para a unidade de hemodiálise. Entregas, nesta ordem:
1. **Card-resumo visual** (ilustração por gpt-image-2.5; texto, doses e prazos sempre no HTML).
2. **Parte teórica** com fundamentação científica e referências para cada conduta.

Plano vivo: decisões entram aqui na hora. Padrão de produção igual ao de
`scripts/deck-builder/biopsia-renal/` (HTML -> PDF/PNG por Playwright).

## Proposta da equipe (discussão original, 2026-10-29, transcrita sem alteração)

### Infecção de cateter
Vanco / ceftazidima empírica - 14 dias

Vanco ataque 35mg/kg - mínimo 2g (90mg) e máximo 3g (90min)
- diluir 10mg/ml

Vanco Manutencao 15mg/kg - mínimo 1g (60min) e máximo 1.5g (60min)
- diluir 10mg/ml

Ceftazidima
- 2g fds (>= 72)
- 1g dia de semana
- Diluição 100ml
- Infusão 30min

Lock
- quando: toda infecção que vai manter o permcath
- Duração do tratamento
- Quais: vanco 1 ml+ cefta 0.5ml + heparina 0.5ml + soro qsp volume das vias

Coleta hemocultura 2 amostras - cateter e linha venosa
- coletar D0
- Repetir após 72h independente do resultado
- Se primeira + e segunda negativa D1 atb fica no 72h - fazer mais 14
- Se duas positivas = tirar cateter, repetir cada 72h até negativo - sempre mais 14d após a cultura negativa
- Repetir no final sempre

Fav nativa = 2 semanas, Doppler (inicial pelo menos)
Prótese = 4 semanas + Doppler (semanal)
Endocardite = 4 semanas
Procurar acessos antigos, stent venoso, marcapasso, prótese ortopédica
Novo permcath 72h após hemo negativa
Tentar ficar 72h sem acesso quando tirar por hemo positiva
Pra trocar em falha - pelo menos 72h de atb pra trocar

Óstio cateter - mupirocina, curativo gaze e hipafix (não tegaderm). 14 dias
- sem mupirocina = cefazolina
- Tunelite = ceftazidima (+ vanco se bacteremia) + hemocultura (retirar se positivo)

### Infecção urinária
Paciente com condições sociais ok: Ciprofloxacino 500mg/dia VO por 7 dias
Paciente sem social bom, que não vai ser aderente ao VO: Ceftazidima (preferência) ou Amica (ototóxica) EV por 7 dias
Urocultura + antibiograma para todos. Sem URC de controle.

### Pneumonia
Paciente em geral com condição: amoxicilina-clavulanato + azitromicina VO
Paciente com comprometimento pulmonar prévio, com condições: levofloxacino VO
5 dias
Pacientes sem condições sociais: levofloxacino EV por 3 doses
Pedir raio X para casos arrastados ou mais graves

### Celulite
Geral: cefalexina ou cefazolina EV
MRSA: sulfametoxazol-trimetoprima ou clindamicina
7 dias corridos
Casos mais graves: vancomicina, 14 dias corridos

## Decisões tomadas

- **Item 1, vancomicina (decidido):** ataque 35 mg/kg (2-3 g) e manutenção 15 mg/kg (1-1,5 g), **a 10 mg/ml**
  (bula FDA, restrição de volume), **nos 90 min finais da sessão** (ataque em 90 min, manutenção em 60 min),
  pela linha da diálise (El Nekidy 2022).
  - **Nível pré-HD 15-20 mg/l:** a unidade não tem a dosagem. Entra **só no texto teórico**, como sugestão
    "quando disponível". **Fora do card.**
  - **Pré-medicação H1 + H2** (Renz 1998, Wallace 1991): **só no texto teórico**, como sugestão. **Fora do card.**
- **Item 2, retirada do cateter (decidido, conservador):**
  - Germe que não seja S. aureus nem Candida: manter o cateter com lock (regra da discussão).
  - **S. aureus:** tentar lock se acesso precioso e paciente estável. **Parar e trocar por fio-guia** se
    febre após 48-72 h de antibiótico ou hemocultura de controle (72 h) positiva. (Maya 2007: cura 41%;
    complicação grave 25% com febre persistente vs 1,4%. Aslam 2014: fio-guia > lock no S. aureus.)
  - **Candida:** retirar sempre (sem estudo de lock; IDSA rec. 57).
  - **Qualquer germe:** retirar/trocar se febre persistente > 72 h, instabilidade ou infecção metastática.
  - Base para ser conservador: KDOQI 2019 st. 25.2 e IDSA rec. 57 (fio-guia sem sítio alternativo).
- **Item 3, ajuste pela cultura (decidido):** MSSA -> trocar vanco por **cefazolina 20 mg/kg, mínimo 1 g,
  máximo 2 g**, diluída em **SF 0,9%** (~10 ml por grama), **bolus de 3-5 min pela linha venosa no fim da
  sessão**, 3x/semana. MRSA/SCN resistente: manter vanco, suspender ceftazidima. Gram-negativo: suspender vanco
  e ajustar pelo antibiograma. (IDSA rec. 62-63; bula; Sowinski 2001; Duke 2024.)

## Conferência com a literatura (2026-10-29)

### Confirmado na fonte primária
- **Vanco ataque 35 mg/kg:** é a dose do consenso ASHP/IDSA 2020 (Rybak, rec. 13) para dose
  **intradialítica com dialisador de alta permeabilidade**. Após o fim da diálise: 25 mg/kg. Teto 3 g.
- **Vanco manutenção 15 mg/kg:** topo da faixa intradialítica alta permeabilidade (10-15 mg/kg, 3x/sem).
  Após a diálise: 10 mg/kg (alta) / 7,5 (baixa).
- **Lock vanco 1 ml + cefta 0,5 ml + heparina 0,5 ml:** idêntico ao protocolo de Maya 2008
  (Semin Dial 21:539-41, grupo do Allon/UAB), com **estoques vanco 5 mg/ml, cefta 5 mg/ml, heparina
  5.000 U/ml**. IDSA 2009 rec. 60: lock após cada sessão por 10-14 dias se o cateter fica.
- **Empírico vanco + gram-negativo** (IDSA 2009 rec. 61). Ceftazidima 1 g pós-HD / 2 g antes do intervalo longo.
- **Novo cateter de longa permanência após hemocultura negativa** (IDSA rec. 58).
- **Troca por fio-guia se sintomas resolvem em 2-3 dias** (IDSA rec. 59), coerente com "72h de atb pra trocar".
- **Mupirocina no óstio** (IDSA rec. 46, infecção de saída não complicada).

### Divergências a decidir (uma por vez)
1. ~~**Velocidade e diluição da vanco.**~~ DECIDIDO (ver acima). ASHP 2020: ≤ 5 mg/ml e ≥ 60 min por 1 g (10-15 mg/min);
   ataque de 25-35 mg/kg exige 2-3 h. Proposta: 2-3 g em 90 min a 10 mg/ml. 35 mg/kg só vale se
   intradialítico em alto fluxo. Faltam no protocolo: **nível pré-diálise 15-20 mg/l, semanal** (ASHP).
2. ~~**Retirada obrigatória**~~ DECIDIDO (conservador, ver acima). do cateter em *S. aureus*, *Pseudomonas* e *Candida* (IDSA rec. 57, A-II).
   Proposta só retira com 2 hemoculturas positivas.
3. ~~**Descalonamento pela cultura:**~~ DECIDIDO. MSSA -> cefazolina 20 mg/kg pós-HD (IDSA rec. 62-63). Ausente.
4. ***S. aureus*: 4-6 semanas** (IDSA rec. 78, 64), 14 dias só com critérios (rec. 80: não diabético,
   sem prótese, sem endocardite, bacteremia resolvida em 72 h...). Proposta: 14 dias após cultura negativa.
5. **Endocardite:** proposta 4 semanas; IDSA 4-6 semanas (S. aureus em valva nativa = 6, AHA 2015).
6. **Tunelite:** IDSA rec. 44 = retirar o cateter + 7-10 dias. Proposta mantém se hemocultura negativa e
   usa ceftazidima isolada (não cobre estafilococo, principal agente).
7. **FAV/prótese:** KDOQI 2019 não fixa duração; prótese com infecção extensa = excisão cirúrgica.
   Proposta não menciona cirurgia.
8. **Pneumonia:** ATS/IDSA 2019 exige imagem para o diagnóstico. Proposta: RX só em caso arrastado/grave.
9. **Doses orais ajustadas à HD** não estão escritas (cipro, amox-clav, levo, cefalexina, SMX-TMP).
10. **Celulite:** IDSA 2014 = 5 dias, estender se não melhora. Proposta 7 dias (aceitável, só registrar).

### Item 1 (vanco): restrições da equipe e literatura de apoio
Restrições da equipe (2026-10-29): **volume mínimo** e **tempo máximo de 90 min, nos 90 min finais da
sessão** (não há como infundir após a diálise, por logística).
- **10 mg/ml:** bula FDA (Pfizer, vancomycin HCl): "In selected patients in need of fluid restriction,
  a concentration up to 10 mg/mL may be used; use of such higher concentrations may increase the risk
  of infusion-related events." Paciente em HD é restrito de volume por definição.
- **Última hora pela linha da diálise:** El Nekidy WS, Cha R, Ghazi IM. Clin Nephrol 2022;97(2):111-20
  (revisão do farmacêutico de nefrologia): ataque 18-23 mg/kg e manutenção 8-13 mg/kg na última hora de
  HD de alto fluxo, nível pré-HD >= 13 e bons desfechos em bacteremia e pele/partes moles.
- **Velocidade maior com pré-medicação:** Renz CL et al. Anesth Analg 1998;87(3):681-5 (ECR duplo-cego):
  1 g em 10 min; difenidramina <= 1 mg/kg + cimetidina <= 4 mg/kg VO 1 h antes -> hipotensão 0% vs 50%
  no placebo; infusão interrompida 5% vs 50%. Wallace MR et al. J Infect Dis 1991;164(6):1180-5 (ECR):
  1 g/60 min, difenidramina 50 mg -> síndrome do homem vermelho 0/16 vs 8/17 no placebo.
- Velocidades da proposta: 2 g/90 min = 22 mg/min; 3 g/90 min = 33 mg/min; 1-1,5 g/60 min = 17-25 mg/min
  (bula: <= 10 mg/min).

### Item 2 (retirada): restrição da equipe e literatura para salvar o cateter
Restrição (2026-10-29): falência de veias torna o acesso precioso; equipe quer ser conservadora na retirada.
- **KDOQI 2019, statement 25.2 (Expert Opinion):** abordagem individualizada; opções incluem troca por
  fio-guia, retirada e reinserção, salvamento do cateter e lock concomitante, "particularly if the CVC is
  deemed to be the patient's final access".
- **IDSA 2009 rec. 57 (B-II):** sem sítio alternativo -> troca do cateter infectado por fio-guia.
- **Aslam S et al. JASN 2014;25(12):2927-41** (meta-análise, 28 estudos, 1.596 pacientes): lock e troca por
  fio-guia têm cura semelhante, ambos superiores a só antibiótico sistêmico (OR 2,08 e 2,88). Cura maior em
  SCN > gram-negativos > S. aureus. **No S. aureus, troca por fio-guia > lock (OR 4,72).**
- **Poole CV et al. NDT 2004;19(5):1237-44:** cura com lock 87% gram-negativo, 75% S. epidermidis,
  40% S. aureus. Sobrevida do cateter sem infecção maior que com troca rotineira (154 vs 71 dias).
- **Maya ID et al. AJKD 2007;50(2):289-95** (113 S. aureus, lock 3 semanas): cura 41%; complicação grave
  em 25% dos com febre persistente vs 1,4% dos demais -> **febre persistente é o sinal de parar de tentar.**
- **Guía Española de Acceso Vascular (2017):** recomenda (fraca) retirada em S. aureus, Pseudomonas e
  Candida; não há estudo de lock em Candida.

### Item 3 (ajuste pela cultura) e cefazolina
- **IDSA 2009 rec. 62-63 (A-II):** MSSA -> trocar vanco por cefazolina, **20 mg/kg (peso real), arredondado
  para 500 mg, após a diálise**. Equipe aprovou a regra de ajuste (2026-10-29); pediu estudo de concentração/velocidade.
- **Bula (cefazolina, Hikma, rev. 2020):** "Direct (bolus) injection: ... further dilute vials with
  approximately 5 mL Sterile Water for Injection. Inject the solution slowly over 3 to 5 minutes, directly or
  through tubing". Infusão intermitente: 50-100 ml. -> **cerca de 100 mg/ml (1 g em ~10 ml), em 3-5 min.**
- **Remoção pela diálise:** Sowinski KM et al. AJKD 2001;37(4):766-76: razão de redução 62% (alto fluxo) e
  50% (alta eficiência) numa sessão; 15-20 mg/kg após cada sessão mantêm nível adequado em intervalos de 2-3 dias.
  -> **dar no fim da sessão** (últimos minutos, pela linha venosa) para não ser dialisada.
- **Dose fixa:** Duke C et al. J Antimicrob Chemother 2024;79(11):2980-9: **2 g 3x/semana pós-diálise**
  sustentado para MSSA (alto fluxo, pacientes infectados). Esquema 2 g / 2 g / 3 g (3 g antes do intervalo
  longo) aparece em guias institucionais (UCSF IDMP), sem estudo próprio localizado.
- Estudo em andamento: CEFAZODIAL (NCT06093269), 20 mg/kg no circuito ao fim da sessão de 4 h.

### Fontes primárias lidas
- Mermel LA et al. IDSA 2009 CRBSI. Clin Infect Dis 2009;49(1):1-45. PMC4039170.
- Rybak MJ et al. Vancomycin TDM revised consensus. Am J Health Syst Pharm 2020;77(11):835-64.
- Maya ID. Antibiotic lock for tunneled HD catheter bacteremia. Semin Dial 2008;21(6):539-41.
- KDOQI Vascular Access 2019, Supplement 2 (AV access infections).
- Bula FDA vancomicina (Pfizer), Dosage and Administration.
- El Nekidy 2022, Renz 1998, Wallace 1991 (acima).
