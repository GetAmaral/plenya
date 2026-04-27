# NAP Master — Plenya + Dr. Getúlio

> **NAP** = Name / Address / Phone. Cada caractere precisa ser idêntico em **todos** os listings (Google Business Profile, Doctoralia, BoaConsulta, CatalogoMed, JSON-LD do site, footer). Divergência = perda de sinal de autoridade local.

## Como usar

1. Preencha os campos `[PREENCHER]` abaixo.
2. Use **exatamente** o mesmo texto (acentos, caixa, abreviações) em todo lugar.
3. Quando algum dado mudar (mudança de telefone, sede), atualize aqui PRIMEIRO e depois propague.

---

## Plenya (clínica)

| Campo | Valor |
|---|---|
| **Nome legal** | Plenya Saúde |
| **Nome curto** | Plenya |
| **Categoria primária** | Clínica médica · Medicina funcional integrativa |
| **Categorias secundárias** | Nefrologia · Medicina preventiva · Longevidade |
| **CNPJ** | `[PREENCHER]` |
| **Endereço completo** | `[PREENCHER — Rua, número, bairro, Londrina, PR, CEP]` |
| **Telefone fixo** | `[PREENCHER — formato (43) 9999-9999]` |
| **WhatsApp** | `[PREENCHER — formato +55 43 99999-9999]` |
| **Email** | contato@plenyasaude.com.br |
| **Site** | https://plenyasaude.com.br |
| **Horário** | Segunda a sexta, 8h às 18h |
| **Instagram** | https://instagram.com/plenyaSaude |
| **Diretor Técnico** | Dr. Getúlio José Mattos do Amaral Filho · CRM-PR 21.876 · RQE 16.038 |

### Descrição curta (até 250 caracteres — usar em GBP, BoaConsulta short bio)
```
Clínica de medicina funcional integrativa em Londrina-PR. Programa Continuum Plenya — saúde, performance e longevidade com equipe multidisciplinar (médico, nutricionista, psicólogo, educador físico) e Método AGIR.
```

### Descrição longa (até 750 caracteres — GBP descrição completa)
```
Plenya é uma clínica de medicina funcional integrativa em Londrina, PR. Atende em formato particular, presencial e por telemedicina, com foco em saúde preventiva, performance e longevidade — não em doença instalada.

O cuidado é organizado pelo Método AGIR: Atividade Física & Alimentação, Gestão Clínica & Metabólica, Integração Mente-Corpo e Ritmo Circadiano. A medida é o Escore Plenya, com mais de 800 itens longitudinais.

Direção clínica do Dr. Getúlio Amaral Filho — nefrologista, professor, autor do livro ANTES.
```

---

## Dr. Getúlio Amaral Filho (médico — perfil pessoal)

| Campo | Valor |
|---|---|
| **Nome completo** | Dr. Getúlio José Mattos do Amaral Filho |
| **Nome curto** | Dr. Getúlio Amaral Filho |
| **CRM** | CRM-PR 21.876 |
| **RQE** | 16.038 (Nefrologia) |
| **Especialidade primária** | Nefrologia |
| **Especialidade secundária** | Clínica Médica |
| **Pós-graduação** | Medicina Funcional Integrativa (ABMFI) |
| **Formado em** | 2004 — Universidade Estadual de Londrina (UEL) |
| **Cidades atende** | Londrina-PR (presencial); Brasil (online) |
| **Email pessoal** | contato@drgetulioamaralfilho.com.br |
| **Site** | https://drgetulioamaralfilho.com.br |
| **Instagram** | https://instagram.com/drGetulioAmaralFilho |
| **Atende por** | Plenya (Londrina) · Nefroclínica Londrina · Telemedicina |
| **Tempo de consulta** | 60-90 minutos |
| **Modalidade** | Particular (não atende convênio) |

### Bio curta (até 250 caracteres)
```
Médico nefrologista (CRM-PR 21.876 · RQE 16.038), formado pela UEL em 2004. Especialista em medicina interna e nefrologia pela Santa Casa de Londrina. Pós-graduação em medicina funcional integrativa pela ABMFI. Direção clínica da Plenya. Autor do livro ANTES.
```

### Bio longa (Doctoralia, BoaConsulta — até 1500 caracteres)
```
Médico nefrologista com vinte anos de prática clínica em Londrina-PR. Formado em medicina pela Universidade Estadual de Londrina (UEL) em 2004, com especialização em Clínica Médica (2006) e Nefrologia (2008) pela Santa Casa de Londrina, onde hoje coordena a Residência Médica em Nefrologia.

Pós-graduado em Medicina Funcional Integrativa pela ABMFI. Sócio da Nefroclínica Londrina e Direção Clínica da Plenya — clínica de longevidade e medicina preventiva.

Responsável técnico das unidades DaVita Intra Hospitalar Santa Casa de Londrina e DaVita Londrina (hemodiálise).

Autor do livro "ANTES — A Janela Silenciosa entre o Normal e o Ótimo" (2026), sobre os anos em que a longevidade é construída ou perdida.

Atende em formato particular, presencial em Londrina ou online por telemedicina, com foco em prevenção, longevidade e medicina funcional integrativa.

CRM-PR 21.876 · RQE 16.038
```

---

## Padronização (use copy-paste)

| Tipo | String exata |
|---|---|
| Cidade | `Londrina` |
| Estado (sigla) | `PR` |
| Estado (extenso) | `Paraná` |
| País | `Brasil` |
| Idioma | `pt-BR` |
| Categoria GBP Plenya | `Clínica médica` |
| Categoria GBP Médico | `Nefrologista` |
| Subcategoria adicional | `Médico` |

---

## Checklist de propagação (após preencher dados acima)

- [ ] Atualizar `packages/brand/src/brand.ts` adicionando `address`, `phone`, `whatsapp`
- [ ] Atualizar `apps/site/components/layout/site-footer.tsx` com endereço/telefone visível
- [ ] Atualizar `apps/site/app/[locale]/contato/page.tsx` com endereço completo + `<address>` semântico
- [ ] Atualizar `apps/site/components/seo/medical-clinic-schema.tsx` adicionando `address.streetAddress` + `telephone` + `geo.latitude/longitude`
- [ ] Atualizar `apps/site-getulio/components/seo/clinics-schema.tsx` com endereços reais
- [ ] Criar Google Business Profile (ver `02-...md`)
- [ ] Cadastrar Doctoralia + BoaConsulta + CatalogoMed (ver `03` e `04`)

> Quando você preencher os campos acima, me avise — eu propago automático nos arquivos de código (passos 1-5) e te entrego prontas as cartas de cadastro nas plataformas.
