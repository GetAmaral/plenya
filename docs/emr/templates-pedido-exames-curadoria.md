# Curadoria dos 3 painéis macro — Pedido de Exames

> Proposta de população dos templates `lab_request_templates`. **Para revisão do Dr. item a item.**
> Mecânica já em prod (commit `68f5303d`, migration 00038). Aqui é só o passo de DADOS.
> Legenda: **C** = Painel Plenya Completo · **I** = Painel Inicial · **A** = Painel Acompanhamento.
> Completo = todos os itens não-lixo desta lista. Sexo: `all` salvo marcado ♂/♀.

## 0. Decisões pendentes de aprovação (estruturais)

### 0.1 — Criar 2 exames avulsos que faltam (recomendado)
Não existe Glicemia de jejum nem Insulina basal como item requisitável — só como **componentes da
curva** (`GLICOSE 0 MIN`/`INSULINA 0 MIN`, `is_requestable=false`, **com `parent_test_id` apontando
pro TTOG e sem TUSS** → não dá pra reaproveitar). São base de I e A (HOMA-IR). Criar avulsos limpos.
TUSS oficiais confirmados online (iClinic/Tabela 22 + PROSAUDE + TRF5 + Diagnósticos do Brasil + IASSEPE):
- **Glicemia de jejum** — TUSS **40302040** ("Glicose - pesquisa e/ou dosagem"), biochemistry, `all`.
- **Insulina (jejum)** — TUSS **40316360** ("Insulina - pesquisa e/ou dosagem"), biochemistry, `all`.

### 0.1b — Gaps revelados pelos 4 pedidos reais do Dr. (TUSS confirmados online)
Análise de 4 pedidos reais do Dr. (1 paciente ♂, 3 ♀; PDFs não versionados, fora do repo por LGPD).
Exames que ele pede e o catálogo não cobre como requisitável:
- **Albumina** — já existe (`PLN6885D35A`, TUSS **40301222** ✓ confirmado), mas `is_requestable=false`
  e com `parent_test_id`. **Ativar** (set requestable=true). Aparece nos **4/4**.
- **Eritropoietina sérica** — criar. TUSS **40305295** (iClinic+PROSAUDE+BoaConsulta). hematology. 1/4.
- **Ultrassonografia de tireoide** — criar. TUSS **40901203** ("US - Órgãos superficiais (tireóide…)",
  iClinic). imaging. 2/4 (♀). *(Nota: 40901130=US abdome superior e 40901190=US dermatológico — NÃO usar.)*
- **Radiografia de tórax PA + perfil** — criar. TUSS **40805026** ("RX - Tórax - 2 incidências",
  iClinic+MS+TJAM+codigotuss). imaging. 1/4.
- Fora de painel (bespoke, texto livre): RM crânio protocolo demência / Espectroscopia RM / Tractografia RM
  (1 caso, com justificativa CID-10 escrita à mão — a caixa livre já cobre).

### 0.2 — Itens-lixo a deixar de fora (sub-campos do hemograma)
`Blastos`, `Mielócitos`, `Metamielócitos`, `Promielócitos`, `PDW`, `VPM`, `IST` (hematology, sem TUSS,
duplicado do biochem). Fora dos 3 painéis. Sugestão paralela: `is_requestable=false` neles.

### 0.3 — Flags de sexo (conservador)
- **♂ male:** PSA Total, PSA Livre/Total, USG próstata.
- **♀ female:** CA-125, Mamografia bilateral, USG mamas, USG transvaginal.
- Hormônios (Testo, Estradiol, FSH/LH, Progesterona, DHT, SHBG) ficam `all` — o Dr. pede conforme caso.

---

## 1. Bioquímica

| Exame | TUSS | I | A | Sexo |
|---|---|:-:|:-:|:-:|
| Glicemia de jejum *(criar)* | 40302040 | ✓ | ✓ | |
| Insulina (jejum) *(criar)* | 40316360 | ✓ | ✓ | |
| Perfil lipídico / Lipidograma | 40302750 | ✓ | ✓ | |
| ApoB | 40301362 | ✓ | ✓ | |
| Lipoproteína A (Lp(a)) | 40301354 | ✓ | | |
| ApoA1 | 40301354 | | | |
| ApoB / ApoA1 | 40301362 | | | |
| LDL oxidada | — | | | |
| Creatinina | 40301630 | ✓ | ✓ | |
| Ureia | 40302580 | ✓ | ✓ | |
| Ácido úrico | 40301150 | ✓ | ✓ | |
| Cistatina C | 40321100 | ✓ | | |
| Sódio | 40302423 | ✓ | ✓ | |
| Potássio | 40302318 | ✓ | ✓ | |
| Cálcio total | 40301400 | ✓ | ✓ | |
| Cálcio iônico | 40301419 | ✓ | | |
| Magnésio sérico | 40302237 | ✓ | ✓ | |
| Magnésio RBC | 40302237 | | | |
| Fósforo | 40301931 | ✓ | | |
| AST (TGO) | 40302504 | ✓ | ✓ | |
| ALT (TGP) | 40302512 | ✓ | ✓ | |
| Gama GT | 40301990 | ✓ | ✓ | |
| Fosfatase alcalina | 40301885 | ✓ | | |
| Bilirrubinas totais | 40301397 | ✓ | | |
| Bilirrubina direta | 40301397 | ✓ | | |
| Bilirrubina indireta | 40301397 | | | |
| Ferritina | 40316270 | ✓ | ✓ | |
| Ferro | 40301842 | ✓ | | |
| Capacidade de fixação de ferro / IST | 40301427 | ✓ | | |
| 25-hidroxivitamina D | 40302830 | ✓ | ✓ | |
| Vitamina B12 (Cobalamina) | 40316572 | ✓ | ✓ | |
| Ácido fólico plasmático | 40301087 | ✓ | | |
| Ácido fólico eritrocitário | 40301087 | | | |
| Homocisteína | 40302113 | ✓ | | |
| Ácido Metilmalônico | 40313301 | | | |
| Zinco | 40313590 | ✓ | | |
| Cobre | 40301567 | | | |
| Selênio | 40313514 | | | |
| Coenzima Q10 | 40321126 | | | |
| Vitamina A | 40302601 | | | |
| Vitamina B1 (HPLC) | 40304302 | | | |
| Vitamina B2 (HPLC) | 40304310 | | | |
| Piridoxal-5-fosfato (B6) | — | | | |
| Vitamina C | 40301060 | | | |
| Vitamina E | 40302610 | | | |
| Adiponectina | 40321010 | | | |
| Leptina | 40321339 | | | |
| PSA Total | 40316149 | ✓ | | ♂ |
| PSA Livre/Total (%fPSA) | 40316130 | | | ♂ |
| CA-125 | 40316025 | | | ♀ |
| CEA | 40316041 | | | |
| Alfafetoproteína (AFP) | 40316068 | | | |
| NT-proBNP | 40302776 | | | |
| Troponina I Ultrassensível | 40302571 | | | |
| CPK | 40301648 | | | |
| Lipase | 40302199 | | | |
| Amilase | 40301281 | | | |
| Eletroforese de proteínas | 40301761 | | | |
| Gasometria venosa | 40302016 | | | |
| Curva insulina+glicose 75g (TTOG) | 40317293 | | | |
| Calprotectina fecal | 40303330 | | | |
| Gliadina Deaminada IgA (DGP-IgA) | 40306356 | | | |
| Rotina de urina (EAS) | 40311210 | ✓ | ✓ | |
| Chumbo | 40313140 | | | |
| Mercúrio | 40313360 | | | |
| Arsênio total urina | 40313069 | | | |
| Alumínio | 40301273 | | | |
| Cromo | 40313158 | | | |
| Manganês | 40313344 | | | |

## 2. Hematologia / Coagulação

| Exame | TUSS | I | A | Sexo |
|---|---|:-:|:-:|:-:|
| Hemograma completo | 40305627 | ✓ | ✓ | |
| Hemoglobina glicada (HbA1c) | 40302075 | ✓ | ✓ | |
| VHS (ESR) | 40305120 | ✓ | | |
| Reticulócitos | 40305350 | | | |
| Tempo de Protrombina (INR) | 40305570 | ✓ | | |
| Fibrinogênio | 40305260 | ✓ | | |
| D-dímero | 40304892 | | | |

## 3. Hormônios

| Exame | TUSS | I | A | Sexo |
|---|---|:-:|:-:|:-:|
| TSH | 40316521 | ✓ | ✓ | |
| T4 Livre | 40316491 | ✓ | ✓ | |
| T3 Livre | 40316467 | ✓ | | |
| T3 Reverso | — | | | |
| Testosterona Total | 40316513 | ✓ | | |
| Testosterona Livre | 40316505 | ✓ | | |
| SHBG | 40316580 | ✓ | | |
| DHT | 40316220 | | | |
| DHEA-S | 40316351 | ✓ | | |
| Estradiol | 40316246 | ✓ | | |
| Estrona (E1) | 40316262 | | | |
| Progesterona | 40316408 | | | |
| FSH | 40316289 | ✓ | | |
| LH | 40316351 | ✓ | | |
| Prolactina | — | ✓ | | |
| IGF-1 | 40316440 | ✓ | | |
| Cortisol plasmático basal | 40316173 | ✓ | | |
| ACTH | 40316009 | | | |
| PTH | 40316424 | ✓ | | |
| TRAb | 40323170 | | | |

## 4. Imunologia / Inflamação / Sorologias

| Exame | TUSS | I | A | Sexo |
|---|---|:-:|:-:|:-:|
| PCR ultrassensível | 40308391 | ✓ | ✓ | |
| Anti-TPO | 40306933 | ✓ | | |
| Anti-Tireoglobulina | 40307034 | | | |
| LDH | 40301729 | ✓ | | |
| Hepatite B - HBsAg | 40307603 | ✓ | | |
| Hepatite B - Anti-HBs | 40307581 | ✓ | | |
| Hepatite B - Anti-HBc | 40307565 | ✓ | | |
| Hepatite C - Anti-HCV | 40307620 | ✓ | | |
| HIV 1+2 | 40307735 | ✓ | | |
| VDRL | 40310345 | ✓ | | |
| Urocultura + antibiograma | 40310116 | | | |
| FAN | 40306941 | | | |
| Complemento C3 | 40308006 | | | |
| Complemento C4 | 40308006 | | | |
| Anti-RO (SSA) | 40306810 | | | |
| Anti-LA (SSB) | 40306780 | | | |
| Anti-endomísio IgA | 40306259 | | | |
| Anti-transglutaminase IgA | 40306470 | | | |
| Anti-transglutaminase IgG | 40306488 | | | |
| IgA / IgG / IgM / IgE | 40307166… | | | |
| IL-6 | 40308766 | | | |
| Microalbuminúria/creatinina (RAC) | 40311112 | ✓ | ✓ | |

## 5. Imagem / Métodos gráficos

| Exame | TUSS | I | A | Sexo |
|---|---|:-:|:-:|:-:|
| Eletrocardiograma (ECG) | 40101010 | ✓ | | |
| Ecodopplercardiograma transtorácico (ETT) | 40901106 | ✓ | | |
| TC escore de cálcio coronariano | 41001087 | ✓ | | |
| Angiotomografia coronariana | 41001230 | | | |
| Doppler carótidas/vertebrais | 40901360 | ✓ | | |
| Doppler aorta/artérias renais | 40901394 | | | |
| USG abdome total | 40901122 | ✓ | | |
| DXA composição corporal | 40808149 | ✓ | | |
| Fundoscopia sob midríase | 41301439 | | | |
| Mamografia digital bilateral | 40808041 | ✓ | | ♀ |
| USG mamas | 40901114 | | | ♀ |
| USG transvaginal | 40901300 | | | ♀ |
| USG próstata (abdominal) | 40901750 | | | ♂ |
| TC tórax | 41001079 | | | |
| Endoscopia digestiva alta (EDA) | 40201120 | | | |
| Colonoscopia | 40201082 | | | |
| Radiografia panorâmica mandíbula | 81000405 | | | |
| Teste H2/metano expirado (SIBO) | — | | | |

---

## 6. Painel Inicial ancorado nos 4 pedidos reais do Dr. (recomendado)

Os 4 pedidos são quase o mesmo painel-espinha. **Presente nos 4/4** (= base do Painel Inicial real,
muito mais amplo que o rascunho original de 68; substitui a coluna "I" das tabelas acima onde divergir):

> 25-OH Vit D · Ácido fólico · Ácido úrico · ACTH · **Albumina** · ApoA1 · ApoB · Bilirrubinas T+frações ·
> Cálcio · Capacidade de fixação de ferro (IST) · Cobre · Cortisol basal · Creatinina · CPK ·
> **Glicose jejum** · **Insulina jejum** · DHT · DHEA-S · Estradiol · Ferritina · Ferro ·
> Fosfatase alcalina · Fósforo · FSH · Gama GT · Gasometria venosa · HbA1c · Hemograma · Homocisteína ·
> LH · Magnésio · PCR-us · Perfil lipídico · Potássio · PTH · Selênio · SHBG · Sódio · T3 livre · T4 livre ·
> Testosterona livre · Testosterona total · AST · ALT · TSH · Ureia · Vitamina A · Vitamina B12 ·
> Vitamina C · Vitamina E · Zinco · Rotina de urina · Microalbuminúria/creatinina · DXA composição corporal
> **(≈54 itens)**

**Imagem que ele acopla de rotina:** DXA (4/4) · USG abdome total (3/4) · TC escore de cálcio (2/4) ·
USG próstata ♂ · USG transvaginal + USG mamas + **USG tireoide** ♀.

**Variações por caso (vão pro Completo, não pro Inicial):** Cistatina C (2/4) · Cálcio iônico (2/4) ·
Prolactina (3/4) · Amilase · Fibrinogênio · T3 reverso · Eritropoietina · sorologias (HBsAg/anti-HBs/
anti-HCV/HIV) · Urocultura · Anti-TPO/Anti-TG/TRAb · IGF-1 · LDH · ETT · Rx tórax · EDA · Colonoscopia.

**Implicação:** o **Inicial** sobe dos ~68 originais para ~54 itens-núcleo dele + imagem por sexo;
itens que eu tinha posto como "Completo only" (ACTH, Albumina, Cobre, Selênio, Gasometria, CPK, Vit A/C/E,
DHT, DHEA-S, Estradiol) entram no Inicial. **Acompanhamento** permanece o enxuto (~24).

---

## 7. Decisões do Dr. (2026-06-11) sobre composição

| Item | Decisão | Notas |
|---|---|---|
| Sorologias (HBsAg, anti-HBs, anti-HCV, HIV, VDRL) | **Completo** apenas | fora do Inicial |
| Imagem de rotina (DXA, USG abdome, escore de cálcio, USG sexo-específico) | **Completo + Inicial** | |
| **Albumina** | **Inicial** (avulsa) | no **Completo** é coberta pela **Eletroforese de proteínas** (`PLNC2D0EC61`) |
| **Eritropoietina** (40305295) | **Completo** apenas | |
| **IGF-1 / Somatomedina C** | **Completo** apenas | (Dr. escreveu "somatostatina" = somatomedina C) |
| **Peptídeo C** (criar, 40316394) | **Completo + Inicial + Acompanhamento** | embasamento §8 |

## 8. Embasamento do Peptídeo C nos 3 painéis (estudo solicitado pelo Dr.)

**Por que medir Peptídeo C de jejum num programa de longevidade/performance (não-diabético):**

1. **Marcador mais fiel da secreção endógena de insulina.** Insulina e peptídeo C são co-secretados em
   quantidades equimolares pela célula β, mas a insulina sofre extração hepática de primeira passagem
   (~50%, variável) e tem meia-vida de 4–5 min; o peptídeo C **não é extraído pelo fígado**, é depurado
   pelo rim e tem meia-vida de **30–35 min** → nível mais estável, menos ruído pulsátil, reflete a
   produção integrada de insulina (StatPearls/NCBI NBK526026; *Diabetes Care* 2020;43:2296 sobre clearance).
2. **Melhor índice de resistência insulínica/hiperinsulinemia que a insulina isolada** (Patel et al.,
   *J Am Heart Assoc* 2012;1:e003152). Num programa que rastreia hiperinsulinemia subclínica antes da
   disglicemia, é um marcador precoce e reprodutível.
3. **Valor prognóstico em adultos SEM diabetes** (o nosso público). Coorte NHANES III (n≈9.211, sem DM):
   o maior quartil de peptídeo C de jejum vs. o menor teve **HR 1,80 (IC95% 1,33–2,43) p/ mortalidade
   geral**, **3,20 (2,07–4,93) p/ mortalidade cardiovascular** e **2,73 (1,55–4,82) p/ DAC**, prevendo
   melhor que HbA1c ou glicemia (Min & Min, *CMAJ* 2013;185:E402). Reforço: *Diabetes Care* 2013;36:708
   (associação independente com mortalidade total/CV e gravidade de DAC).
4. **Sinergia com o que já pedimos.** Com glicemia e insulina de jejum no mesmo painel, o peptídeo C
   permite separar **resistência insulínica com β preservada** (peptídeo C alto) de **falência β**
   (peptídeo C baixo) e calcular razões C-peptídeo/glicose — leitura que a insulina sozinha não dá.

**Caveat obrigatório (relevante p/ nefro):** o peptídeo C é **depurado pelo rim** → acumula na DRC e
**superestima** a secreção quando o RFG cai. Interpretar sempre junto da creatinina/cistatina C/RFG.
Em jejum, coletar com glicemia simultânea (a leitura é a dupla peptídeo C × glicose).

> Fontes: NCBI StatPearls NBK526026 · Diabetes Care 2020;43:2296 · J Am Heart Assoc 2012;1:e003152 ·
> CMAJ 2013;185:E402 (NHANES III) · Diabetes Care 2013;36:708. (Links nos resultados de busca da sessão.)

---

## Resumo de tamanho (proposta)
- **Completo:** ~141 itens (tudo menos os 7 lixo + os 2 a criar).
- **Inicial:** ~68 itens (baseline 1ª consulta).
- **Acompanhamento:** ~24 itens (rotina enxuta).

## Plano de semeadura (após aprovação)
1. Criar 5 exames avulsos — INSERT idempotente em `lab_test_definitions`:
   Glicose jejum (40302040), Insulina jejum (40316360), **Peptídeo C (40316394, hormones)**,
   Eritropoietina (40305295, hematology), USG tireoide (40901203, imaging),
   Rx tórax PA+perfil (40805026, imaging).
2. Ativar **Albumina** (`PLN6885D35A`): `is_requestable=true`.
3. UPDATE `sex_applicability` nos itens de 0.3.
4. INSERT 3 `lab_request_templates` + vínculos no m2m `lab_request_template_tests`
   (Inicial ancorado na §6).
5. Dry-run em clone do prod → aplicar dev → aplicar prod (banco direto, idempotente).
6. DELETE templates `teste`/`teste2`.

## Justificativa por exame (convenção `#` no texto livre) — 2026-06-12

O médico pode justificar um exame digitando, **logo abaixo dele**, uma linha iniciada por `#`.
Linhas `#` consecutivas concatenam (justificativa multi-linha; word-wrap automático no PDF).
`#` órfão (sem exame acima) ou vazio é ignorado. Exemplo no texto livre do pedido:

```
Ressonância de crânio
# Paciente com sinais de demência, investigação diagnóstica de causa estrutural.
# Afastar hidrocefalia de pressão normal.
Hemograma completo
```

**Por que `#` e não `>`/aspas:** o médico usa `>`/`<` como limiar ("clearance < 60") no início de
frases, e aspas sofrem com autocorreção (curvas) e delimitador não-fechado. `#` quase nunca abre
uma linha clínica → marcador menos ambíguo. Prefixo de linha (não wrapper aberto-fecha) é robusto:
só precisa do início, nunca colide com pontuação interna.

**Sem schema novo:** a justificativa vive dentro do próprio `LabRequest.Exams` (texto livre);
persiste de graça, zero migration. No render (`pdfdoc`), vira `ExamItem.Justification` e aparece
sob o nome do exame, em itálico com barra dourada. A paginação conta só exames (justificativa não
ocupa "vaga" das 40/página). No EMR, o contador de exames e o preview em badges ignoram linhas `#`.

Arquivos: `apps/api/internal/pdfdoc/exam_request.go` (parser + render), `pdfdoc/stationery.go`
(`.exjust`), `apps/web/app/(authenticated)/lab-requests/page.tsx` (placeholder + contador + preview).
Pendência engatilhada: no import/dedup externo, ao remover um exame casado, remover as linhas `#`
logo abaixo dele junto.
