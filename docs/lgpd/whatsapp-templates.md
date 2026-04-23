# Templates WhatsApp Business — registro versionado

Registro de todos os templates submetidos à Meta WhatsApp Cloud API e seu status de aprovação.
Este arquivo é a **fonte de verdade humana** — quando atualizar texto, bumpe a versão e mantenha o histórico.

---

## Template: `magic_link`

| Campo | Valor |
|-------|-------|
| **Nome técnico** | `magic_link` |
| **Categoria** | `utility` |
| **Idioma** | `pt_BR` |
| **Variáveis** | `{{1}}` = URL do magic link |
| **Sub-tipo botões** | nenhum (link no body) |
| **Status atual** | ⏳ a submeter |
| **Justificativa categoria** | Mensagem transacional vinculada a ação iniciada pelo usuário (claim do Escore Light). Sem cunho promocional. |

### Versão 1.0 (2026-04-23)

```
Olá! Aqui é da Plenya. Seu resultado do Escore Plenya Light está pronto. Acesse pelo link seguro: {{1}}. Esse link é único e expira em 7 dias. Se não foi você, ignore esta mensagem.
```

**ATENÇÃO:** o "7 dias" precisa bater com o JWT TTL em
`apps/api/internal/services/anonymous_score_service.go` (busca por `7 * 24 * time.Hour`).
Se mudar um, mudar o outro **e** ressubmeter o template à Meta (texto novo = nova aprovação).

### Histórico de versões

| Versão | Data | Mudança | Submetido | Aprovado |
|--------|------|---------|-----------|----------|
| 1.0    | 2026-04-23 | Criação inicial (7 dias) | - | - |

---

## Processo pra adicionar/alterar template

1. Editar o arquivo aqui primeiro (versão + texto + categoria + justificativa).
2. Acessar Meta Business Manager → WhatsApp Manager → Message templates → Create template.
3. Colar o texto exato. Selecionar categoria + idioma `pt_BR`.
4. Submeter. Aprovação Meta leva até 24h.
5. Quando aprovado, atualizar status nesta tabela.
6. Se rejeitado, anotar motivo aqui e iterar (categoria errada é o motivo mais comum — utility ↔ marketing).

## Princípios

- **Categoria correta importa MUITO.** Marketing custa 5×–10× mais que utility e exige opt-in explícito por mensagem (não basta opt-in geral).
- **Conteúdo proibido:** ofertas, descontos, "promoção", linguagem persuasiva forte → vira marketing automaticamente.
- **Variáveis:** usar `{{N}}` numéricas, nunca nomeadas.
- **Botões:** evitar URL_BUTTON com variável dinâmica; preferir link no body (mais simples, aprovado mais rápido).
- **Mudança de texto** após aprovação requer **nova submissão**. Meta NÃO permite editar in-place.

## Aprovações pendentes

- [ ] `magic_link` v1.0 (utility, pt_BR) — submeter quando Business Verification concluir
