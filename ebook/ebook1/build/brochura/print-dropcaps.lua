-- print-dropcaps.lua — CÓPIA DA BROCHURA
-- Diverge de versaoImpressa/print-dropcaps.lua em dois pontos, pedidos pelo
-- editor em 2026-08-26 ("talvez a tua intenção era colocar uma letra inicial
-- maior, acho que fica bonito; talvez tenha faltado ali"):
--   1. A Introdução também ganha capitular. Ela chega como RawBlock
--      \chapter*{...} (o Header nunca dispara), então há um hook de RawBlock.
--      Agradecimentos, Sobre o Autor e Referências seguem sem capitular.
--   2. Parágrafo de abertura curto não é mais descartado: abaixo de 200
--      caracteres a capitular cai para 2 linhas em vez de sumir. Só abaixo de
--      100 é que não cabe capitular nenhuma.
--
-- Pandoc Lua filter: insere \lettrine{primeira-letra}{resto-da-palavra} no
-- primeiro parágrafo depois da abertura de capítulo.

local awaiting_dropcap = false

local SKIP_TITLES = {
  ["Agradecimentos"] = true,
  ["Sobre o Autor"] = true,
  ["Sobre o autor"] = true,
}

-- Aberturas que chegam como \chapter*{...} cru (o build monta o LaTeX direto
-- para não numerá-las). Sem este hook o Header nunca dispara e elas ficariam
-- sem capitular.
function RawBlock(el)
  if el.format ~= "latex" and el.format ~= "tex" then return el end
  local title = el.text:match("\\chapter%*{(.-)}")
  if not title then return el end
  if SKIP_TITLES[title] or title:match("^Referências") then
    awaiting_dropcap = false
  else
    awaiting_dropcap = true
  end
  return el
end

function Header(el)
  if el.level == 1 then
    local plain = pandoc.utils.stringify(el)
    if SKIP_TITLES[plain] or plain:match("^Referências") then
      awaiting_dropcap = false
    else
      awaiting_dropcap = true
    end
  end
  return el
end

function Para(el)
  if not awaiting_dropcap then
    return el
  end
  awaiting_dropcap = false

  -- Extract first letter and remainder of first word
  if #el.content == 0 then return el end
  local first = el.content[1]
  if first.t ~= "Str" then return el end

  local word = first.text
  if #word < 2 then return el end

  -- A capitular ocupa um bloco vertical de N linhas: se o parágrafo tiver menos
  -- texto que isso, o rabo da letra desce por dentro do parágrafo seguinte. Em
  -- vez de abrir mão da capitular (comportamento antigo), o número de linhas
  -- acompanha o tamanho do parágrafo. A medida de linha aqui dá ~70 caracteres.
  local total = 0
  for _, inline in ipairs(el.content) do
    if inline.t == "Str" then total = total + #inline.text
    elseif inline.t == "Space" or inline.t == "SoftBreak" then total = total + 1
    end
  end
  local lines = 3
  if total < 100 then return el
  elseif total < 200 then lines = 2 end

  -- Use UTF-8-aware length: take the first character (could be multi-byte for accented).
  local first_char_bytes
  local b = string.byte(word, 1)
  if b < 0x80 then first_char_bytes = 1
  elseif b < 0xE0 then first_char_bytes = 2
  elseif b < 0xF0 then first_char_bytes = 3
  else first_char_bytes = 4 end

  local first_char = word:sub(1, first_char_bytes)
  local rest = word:sub(first_char_bytes + 1)

  -- Build the LaTeX raw inline
  local letrine = pandoc.RawInline("latex",
    "\\lettrine[lines=" .. lines .. ",findent=2pt,nindent=0pt]{"
    .. first_char .. "}{" .. rest .. "}")

  -- Replace first inline with raw lettrine + the rest of the inlines (drop original first Str)
  local new_inlines = pandoc.List()
  new_inlines:insert(letrine)
  for i = 2, #el.content do
    new_inlines:insert(el.content[i])
  end
  return pandoc.Para(new_inlines)
end
