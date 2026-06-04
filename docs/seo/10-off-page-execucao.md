# SEO Off-page — runbook de execução (NAP Plenya + Dr. Getúlio)

> Estado em 2026-06-04. **On-page/NAP no código está COMPLETO e deployado** nos dois sites
> (`brand.ts`, schemas, footer, `/contato`, `legal.ts`, `clinics-schema.tsx` do Getúlio).
> Fonte canônica dos dados: [`01-NAP-master.md`](01-NAP-master.md). Auditoria de consistência
> 2026-06-04: telefone `+55 43 99974-8899` idêntico em todos os arquivos; CEP de atendimento
> `86050-460` só em NAP/atendimento; CEP fiscal `86058-100` só em `legal.ts`. Sem divergência.

O que falta é **off-page** (listings + verificação) — exige você logado nas contas. Abaixo,
separado por **quem faz**. Ordem pensada pelo caminho crítico (o que destrava mais primeiro).

---

## Dados canônicos para copiar (de `01-NAP-master.md`)

- **Nome:** Plenya Saúde · **Razão social:** Plenya Serviços de Saúde Ltda. · **CNPJ:** 66.991.259/0001-50
- **Endereço de atendimento (NAP):** Av. Ayrton Senna da Silva, 500 — Edifício Torre Pietra, sala 1402, Gleba Palhano, Londrina/PR, 86050-460
- **Geo:** -23.3296924, -51.1779253
- **Telefone/WhatsApp:** +55 43 99974-8899 · **Email:** contato@plenyasaude.com.br
- **Horário:** Seg a sex, 8h às 18h · **Site:** https://plenyasaude.com.br · **IG:** @plenyaSaude
- **Diretor técnico:** Dr. Getúlio José Mattos do Amaral Filho · CRM-PR 21.876 · RQE 16.038
- ⚠️ **NUNCA** usar o endereço fiscal (Av. Gil de Abreu e Souza, 2335) em nada de SEO local.

---

## Passo 1 — Google Business Profile (GBP) · maior impacto, fazer primeiro

**Quem:** você cria + verifica (precisa de telefone/cartão postal). Eu preparo todo o conteúdo.

1. (eu) Entrego o "kit GBP": nome, categorias, descrição curta/longa, horário, serviços,
   atributos, e a primeira leva de fotos a tirar/usar. Base: [`02-google-business-profile.md`](02-google-business-profile.md).
2. (você) business.google.com → criar perfil com o **endereço de atendimento** acima.
3. (você) Verificação (Google manda código por SMS/ligação ou cartão postal). É o gargalo de tempo.
4. (você) Após verificar, colar o conteúdo do kit e publicar.
5. (eu, depois) Confiro se o `MedicalClinic` do site e o GBP batem 100% (NAP citation).

> Sem endereço físico o GBP é rejeitado. Já temos o de atendimento, então está **desbloqueado**.

## Passo 2 — Search Console + Bing Webmaster (verificação) · destrava analytics

**Quem:** você cria conta e me passa os códigos; eu coloco no código/Coolify e dou deploy.

- **Site Getúlio:** já lê `GOOGLE_SITE_VERIFICATION` e `BING_SITE_VERIFICATION` do env. Você só
  cola os 2 códigos no Coolify (app `qkdzqaauicc001qfkghfur0s`) e redeploy. Zero código.
- **Site Plenya:** hoje **não tem** meta tag de verificação. ➜ tarefa minha: replicar o padrão env
  do Getúlio em `apps/site/app/[locale]/layout.tsx` (faço agora, fica pronto pro código chegar).
- (você) Criar 4 propriedades: GSC Plenya, GSC Getúlio, Bing Plenya, Bing Getúlio → me passar os
  4 códigos.
- (eu) Colar/escalar + submeter os sitemaps (`/sitemap.xml`) nos dois consoles. Roteiro:
  [`05-search-console-bing.md`](05-search-console-bing.md).

## Passo 3 — Diretórios médicos (NAP citations) · reforça sinal local

**Quem:** você cria os perfis (exigem confirmar identidade médica/CRM). Eu entrego copy pronta.

- **Doctoralia** ([`03-doctoralia.md`](03-doctoralia.md)) e **BoaConsulta + CatalogoMed**
  ([`04-boaconsulta-catalogomed.md`](04-boaconsulta-catalogomed.md)).
- (eu) Entrego: bio curta/longa do Dr. Getúlio (já prontas no NAP-master), endereço, telefone,
  especialidades, e o texto a colar campo a campo — **idêntico** ao NAP (cada caractere conta).
- (você) Criar perfil em cada um, validar CRM, colar.

## Passo 4 — Wikidata · entidade no grafo do Google

**Quem:** quase tudo eu. Drafts já existem (`wikidata-*.txt`, `05-wikidata-draft.md`).

- (eu) Reviso os QuickStatements (pessoa Dr. Getúlio + livro ANTES + clínica).
- (você) Logar no wikidata.org/QuickStatements (login Google/Wikimedia) e colar o batch, OU me
  autorizar a usar uma conta. Item já existente? Só completar propriedades.

## Passo 5 — IndexNow para o site Plenya · indexação instantânea (opcional)

**Quem:** eu, 100%. Hoje só o Getúlio tem.

- (eu) Replicar o endpoint `/api/indexnow` + chave pública + secret no `apps/site`, espelhando
  o padrão do Getúlio. Setar `INDEXNOW_TRIGGER_SECRET` no Coolify Plenya. Disparo inicial do sitemap.

## Passo 6 — LinkedIn Community Management API · destravado pelo CNPJ

**Quem:** você submete (precisa do login da Página). Eu já tenho o use-case redigido.

- Ver memória `linkedin_cma_pendente` + [`08-linkedin-upgrade.md`](08-linkedin-upgrade.md). Agora
  que CNPJ + razão social existem, o formulário pode ser preenchido.

---

## Resumo do caminho crítico

1. **GBP** (Passo 1) — verificação por correio/SMS leva dias, então começa primeiro.
2. Em paralelo: **códigos de verificação** (Passo 2) e **diretórios** (Passo 3).
3. Resto (Wikidata, IndexNow, LinkedIn) sem urgência.

**O que eu já posso fazer sem você:** scaffold de verificação no site Plenya (Passo 2),
IndexNow Plenya (Passo 5), e montar todos os "kits" de copy (Passos 1, 3, 4, 6).
</content>
</invoke>
