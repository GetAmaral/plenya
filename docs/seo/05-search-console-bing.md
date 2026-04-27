# Google Search Console + Bing Webmaster

Sem isso, você está cego: não sabe quais queries levam ao site, nem se há erros de indexação. Tudo o que fizemos nas Ondas 1-2 é invisível pra você.

**Tempo total**: ~30 min para os 4 properties (2 sites × 2 ferramentas).

---

## Google Search Console (GSC)

### 1. Adicionar properties

Vá em https://search.google.com/search-console e clique **"Adicionar propriedade"**.

Crie **4 properties separadas** (recomendado):

| Property | Tipo | Valor |
|---|---|---|
| 1. Plenya domain | Domínio | `plenyasaude.com.br` |
| 2. Plenya prefixo | Prefixo URL | `https://plenyasaude.com.br` |
| 3. Getúlio domain | Domínio | `drgetulioamaralfilho.com.br` |
| 4. Getúlio prefixo | Prefixo URL | `https://drgetulioamaralfilho.com.br` |

> **Por que dois tipos por site?** Property "Domínio" é mais completa (cobre subdomínios, http/https, www/sem-www) mas precisa verificação DNS. Property "Prefixo URL" é mais simples (só HTML tag ou meta), serve como backup. Tenha as duas.

### 2. Verificação DNS (property tipo Domínio)

Para cada domínio, GSC vai pedir um TXT record. Adicionar no Cloudflare:

```
Tipo: TXT
Nome: @
Valor: google-site-verification=<código_que_gsc_te_dá>
TTL: Auto
Proxy: Desativado (DNS only ⚪)
```

Após adicionar, voltar no GSC e clicar **"Verificar"**. Demora ~5min pra propagar.

### 3. Verificação HTML (property tipo Prefixo URL)

Mais simples — escolher **"Tag HTML"** e colar a meta tag em ambos os layouts.

**Plenya** — `apps/site/app/[locale]/layout.tsx` linha 82-84 (dentro de `<head>`):

```tsx
<meta name="google-site-verification" content="<código_que_gsc_te_dá>" />
```

**Getúlio** — `apps/site-getulio/app/[locale]/layout.tsx` (dentro de `<head>`):

```tsx
<meta name="google-site-verification" content="<código_que_gsc_te_dá>" />
```

> **Quando você me passar os dois códigos, eu adiciono via Edit + commit + deploy.**

### 4. Submeter sitemaps

Após verificação:

| Property | Sitemap a submeter |
|---|---|
| Plenya | `https://plenyasaude.com.br/sitemap.xml` |
| Getúlio | `https://drgetulioamaralfilho.com.br/sitemap.xml` |

Em GSC: **Sitemaps** → cole a URL → **Enviar**.

Resultado esperado em 24-72h: "Sucesso · X URLs descobertas".

### 5. Configurações iniciais

Para cada property:

- [ ] **Endereço**: associar ao Brasil (define geo-targeting)
- [ ] **Configurações > Usuários**: adicionar `getamaralb002@gmail.com` como Owner se não for sua conta principal
- [ ] **Verificar Indexing > Cobertura** após 1 semana — confirmar que páginas estão sendo indexadas (esperado: ~50 Plenya, ~17 Getúlio)

---

## Bing Webmaster Tools

Pequeno trabalho, ROI marginal mas vale (Bing alimenta DuckDuckGo, ChatGPT search, parte do Yahoo).

### Cadastro
1. https://www.bing.com/webmasters
2. Login com mesma conta Google
3. **Importar do Google Search Console** — Bing copia properties + sitemaps automaticamente

Se você cadastrar GSC antes, o trabalho do Bing fica em **2 cliques**.

### Verificação
- Mesma meta tag `<meta name="msvalidate.01" ...>`. **Quando me passar o código, adiciono.**

### Sitemaps
- Auto-submetidos via importação GSC.

---

## Indexnow (bonus — 5 min)

Indexnow é um protocolo que **avisa o Bing/Yandex** quando seu site muda — em vez de esperar o crawler descobrir. Google ainda não usa, mas Bing/Yandex sim.

Implementação: já temos sitemap dinâmico, então é apenas questão de adicionar uma chave em `apps/site/public/<chave>.txt` e enviar pings via webhook após cada deploy. **Implemento depois que você cadastrar Bing.**

---

## Como ler GSC depois (mensal)

Após 30 dias com tráfego, abrir **Performance > Searches** e responder:

1. **Quais queries levam pessoas ao site?** — confirmam (ou não) a tese de SEO
2. **CTR alto + posição baixa** — sinal pra subir o ranking via melhoria de title/description
3. **Posição alta + CTR baixo** — title/description ruim, melhorar
4. **Páginas com cliques zero** — investigar se merecem ficar ou ser removidas

Métricas baseline depois das Ondas 1-2 (estimativa, ainda sem dados):
- 0 → 200-500 cliques/mês (Plenya) em 60 dias
- 0 → 50-150 cliques/mês (Getúlio) em 60 dias

Se em 90 dias estivermos abaixo disso, revisitamos a estratégia.

---

## Checklist execução

- [ ] GSC property Plenya domain (DNS) — código TXT: `_______`
- [ ] GSC property Plenya prefixo — código meta: `_______`
- [ ] GSC property Getúlio domain (DNS) — código TXT: `_______`
- [ ] GSC property Getúlio prefixo — código meta: `_______`
- [ ] GSC sitemap Plenya submetido
- [ ] GSC sitemap Getúlio submetido
- [ ] Bing import GSC
- [ ] Bing verificação meta — código: `_______`

> Me passe os códigos quando GSC/Bing te derem. Faço commits dos meta tags + adicionar TXT no Cloudflare se você liberar acesso.
