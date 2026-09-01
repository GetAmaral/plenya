# Handoff — planos de paciente (Ana Cláudia e José Ricardo)

**Atualizado:** 2026-09-01 · Sessão anterior terminou aqui.
**Dados clínicos ficam em `pacs/<NOME>/`, que é gitignored. Este arquivo não repete valor de exame.**

---

## 1 · O que está FEITO e em produção

### Código (master, deployado nos 3 apps em `f46f67be`)

| Commit | O quê |
|---|---|
| `b56cf1c9` | Layout do receituário: ficha do paciente separada do miolo, fórmulas com respiro, painel de aviamento/posologia. Papelaria compartilhada (`pdfdoc/stationery.go`) — vale para TODO documento. |
| `708118a2` | Nome legível do PDF da receita, no padrão do pedido de exames. Padrão extraído para `internal/utils/docfilename.go` (`DocumentFileName`, `CompactName`, `ASCIIFallback`, `ContentDisposition`). |
| `06a23a33` | **fix do escore hormonal** + migrations 00085/00086. Banco de prod no goose **86**. |
| `fef35188` · `cc1aa9f8` · `f46f67be` | ebook, logs do Instagram, docs novos. |

### Banco de produção (aplicado direto, SQL versionado em `docs/emr/`)

- `magistral-veiculo-tcm.sql` — cadastra **TCM** como veículo oleoso (unidade `ml`, densidade 0,945, entra como qsp e fica fora do cálculo de volume de propósito).
- `magistral-densidade-vitamina-e.sql` — densidade da vitamina E de 0,45 (default de pó) para **0,95** (acetato líquido). O erro dobrava o volume e fazia o `CalculateCapsule` pedir 2 cápsulas onde cabia 1.

### Ana Cláudia — `01a025de-1071-70a1-851b-d2eacd5deef5`

- **Receita magistral em RASCUNHO** (`01a0592b-366c-751b-aa8b-663da085a25b`), sem assinatura, 2 fórmulas de 90 doses cada: lipossolúveis em TCM (almoço) e hidrossolúveis + antioxidantes (jantar).
- Notas internas de cada dose em `pacs/ANA .../formula-notas-internas.md`. **Nada de raciocínio clínico no campo `note` da receita**: ele é impresso no PDF assinado que a farmácia lê.
- Deck de **21 páginas**, `deck-ana.pdf` (16:9) e `deck-ana-A4.pdf`.

### José Ricardo — `01a04572-9787-7a7f-a9b2-5acf3f3f3060`

- **Pedido de exames ASSINADO** em 01/09, 49 exames (`01a04990-3610-71da-beff-8c9f735f8510`).
- Conduta: **tirzepatida + creatina + ômega-3 começam agora; fórmula manipulada espera os exames.**
- Sono: **adequado**, deixou de ser pendência. Medida de cintura: **removida do deck** a pedido.
- Deck de **20 páginas, SÓ 16:9** (`deck-jose-ricardo.pdf`). Ele não quer mais o A4 deste.

---

## 2 · O que ficou PENDENTE

| | |
|---|---|
| **PeptiStrong + HMB** | Sachê de preservação de massa magra para o Ricardo. Nunca despachado. Ele tem 64 anos, começa GLP-1 e não faz treino de carga: é o de maior risco de perder massa magra. |
| **`plano-final.md` do Ricardo** | Desatualizado em 3 pontos: sono como ponto cego, "vitamina D e ômega-3 na fórmula" na semana 0, e circunferência abdominal na semana 0. O deck está certo, o markdown não. |
| **Vitamina D do Ricardo** | Única do conjunto com medida recente (25 ng/mL, abril/2026). Poderia começar hoje; por decisão dele espera junto com o resto. |
| **Falha de catálogo, não corrigida** | Painel-pai (hemograma, lipidograma, rotina de urina, bilirrubinas) nunca tem resultado próprio: o laboratório reporta os analitos filhos. O cruzamento da tela de pedido não olha os filhos, então manda repetir exame de quem acabou de fazer. **É mudança de código.** |
| **Densitometria óssea e espirometria** | Não existem como exame pedível no catálogo. Entram como texto livre. |

---

## 3 · O que ele quer DISCUTIR na volta

### A · Registrar os planos no EMR, com visual em tela e em A4

Hoje o plano vive fora do EMR: `pacs/<NOME>/deck/build.py` → `deck.html` + `deck.css` → Playwright →
PNG por slide, PDF 16:9 e PDF A4 paisagem (escala 1122.52/1920 = 0.5846).

O que precisa ser decidido: onde isso mora no EMR (PatientDocument? entidade nova?), como o paciente
vê no portal, e se o A4 sai do mesmo HTML ou de um render separado.

Peças que já existem e podem ser reaproveitadas:
- `pdfdoc` é a **fonte única de layout** da papelaria (`renderDocument`, paginação in situ no
  Chromium, cabeçalho/rodapé repetidos, assinatura na última página).
- `PatientDocumentsService.CreateFromBytes` publica PDF como documento do paciente, com download
  autenticado.
- Já existe `care_plan_report.go` no `pdfdoc` — **conferir o que ele faz antes de desenhar coisa nova.**

### B · Esqueleto/prompt pré-formado para automatizar novos planos

Objetivo: que montar o plano do próximo paciente não seja recomeçar do zero.
Insumos que hoje são manuais e deveriam ser derivados: réguas (`reguas.json`), classificação dos
exames, escolha de quais achados viram slide, e o texto em voz de paciente.

---

## 4 · Regras aprendidas nesta sessão, para não repetir erro

1. **`prescription_formula_components.note` e `prescription_formulas.instructions` SAEM IMPRESSOS**
   no PDF assinado. Só instrução de manipulação ali. Raciocínio clínico vai para arquivo.
2. **Cruzar protocolo com prontuário exige olhar os analitos filhos**, senão painel-pai vira
   falso-positivo de "nunca feito".
3. **`sex_applicability` usa `'all'`**, não `'both'`.
4. **`MassToMg` devolve `ok=false` para UI, % e ml** — componente em UI sai do cálculo de volume da
   cápsula. Escrever a dose em mg quando o volume importa.
5. **Densidade `density_source='classe'` é default de pó** e erra feio em ativo líquido.
6. **Os scripts de render do deck têm o nome de saída HARDCODED** e são copiados entre pacientes.
   Conferir `render-16x9.js` e `render-a4.js` ao criar deck novo, senão ele grava por cima do
   deck de outro paciente (aconteceu duas vezes).
7. **Verificador de estouro** em `/tmp/chk-overflow.js`: mede o elemento mais baixo/à direita de cada
   `section.slide` contra 1920x1080. Rodar sempre antes de gerar PDF.
8. **Deploy é por-app e um de cada vez.** Matar o script local NÃO aborta o build do Coolify: ele roda
   no servidor. Conferir `application_deployment_queues` antes de concluir que travou.
