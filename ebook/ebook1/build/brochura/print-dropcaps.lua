-- print-dropcaps.lua — CÓPIA DA BROCHURA
--
-- Insere \lettrine{primeira-letra}{resto-da-palavra} no primeiro parágrafo depois
-- da abertura de capítulo.
--
-- Diverge de versaoImpressa/print-dropcaps.lua em dois pontos pedidos pelo editor
-- em 2026-08-26 ("talvez a tua intenção era colocar uma letra inicial maior, acho
-- que fica bonito; talvez tenha faltado ali"):
--   1. A Introdução também ganha capitular. Ela chega como RawBlock \chapter*{...}
--      (o Header nunca dispara), então há um hook de RawBlock. Agradecimentos,
--      Sobre o Autor e Referências seguem sem capitular.
--   2. Parágrafo de abertura curto não é descartado.
--
-- Revisão de 2026-09-17: a varredura das 348 páginas achou CINCO capítulos ainda
-- sem capitular, por três causas distintas, todas aqui dentro:
--   3, "Eu tenho 45 anos." — o § abre com aspas, o pandoc entrega um Quoted e o
--      filtro exigia Str. Agora o Quoted é desmontado e as aspas viram `ante` da
--      capitular, que é para o que a opção existe.
--   8 e 9, "O capítulo anterior..." / "A gestão clínica..." — a primeira palavra
--      tem uma letra só e havia um `if #word < 2 then return`. \lettrine{O}{}
--      com segundo argumento vazio é perfeitamente válido.
--   13 e 17, §§ de 62 e 52 caracteres — o corte era `total < 100`, e o capítulo
--      ficava sem capitular. Baixar para capitular de duas linhas não resolve:
--      testado, o rabo da letra desce por dentro do parágrafo seguinte e come o
--      recuo dele. Um § de uma linha não comporta capitular. Então a capitular
--      PULA para o primeiro § que comporta, e a linha curta fica por cima, como
--      entrada. É o que o livro já faz de fato: "Ricardo voltou ao consultório
--      dezoito meses depois do infarto." é uma deixa, não uma abertura.

local awaiting_dropcap = false
local skipped = 0
local MAX_SKIP = 2   -- entradas curtas seguidas; mais que isso, desiste

local SKIP_TITLES = {
  ["Agradecimentos"] = true,
  ["Sobre o Autor"] = true,
  ["Sobre o autor"] = true,
}

-- Caracteres que podem vir colados antes da primeira letra e que devem ficar
-- FORA da capitular (ela é uma letra, não um sinal).
local ANTE = {
  ["\u{201C}"] = true, ["\u{201D}"] = true,   -- “ ”
  ["\u{2018}"] = true, ["\u{2019}"] = true,   -- ‘ ’
  ['"'] = true, ["'"] = true, ["("] = true,
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
  skipped = 0
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
    skipped = 0
  end
  return el
end

-- Primeiro caractere UTF-8 de uma string, e o resto.
local function utf8_head(s)
  local b = string.byte(s, 1)
  local n
  if b < 0x80 then n = 1
  elseif b < 0xE0 then n = 2
  elseif b < 0xF0 then n = 3
  else n = 4 end
  return s:sub(1, n), s:sub(n + 1)
end

function Para(el)
  if not awaiting_dropcap then
    return el
  end
  local content = el.content
  if #content == 0 then awaiting_dropcap = false; return el end
  local first = content[1]

  -- Parágrafo que abre com fala entre aspas chega como Quoted. Desmonta para Str
  -- + conteúdo + Str, para que as aspas possam ir no `ante` logo abaixo.
  if first.t == "Quoted" then
    local abre, fecha = "\u{2018}", "\u{2019}"
    if first.quotetype == "DoubleQuote" then abre, fecha = "\u{201C}", "\u{201D}" end
    local inner = first.content
    if #inner == 0 or inner[1].t ~= "Str" then awaiting_dropcap = false; return el end
    -- a aspa entra COLADA na primeira palavra; como Str isolada ela viraria o
    -- conteúdo inteiro do primeiro inline e a palavra sumiria na limpeza do ante.
    local novo = pandoc.List()
    novo:insert(pandoc.Str(abre .. inner[1].text))
    for i = 2, #inner do novo:insert(inner[i]) end
    novo:insert(pandoc.Str(fecha))
    for i = 2, #content do novo:insert(content[i]) end
    content = novo
    first = content[1]
  end

  if first.t ~= "Str" then awaiting_dropcap = false; return el end

  -- Separa o que vem antes da letra (aspas, parêntese) do corpo da palavra.
  local word = first.text
  local ante = ""
  while #word > 0 do
    local ch, rest = utf8_head(word)
    if ANTE[ch] then ante = ante .. ch; word = rest else break end
  end
  if #word < 1 then awaiting_dropcap = false; return el end

  -- A capitular ocupa um bloco vertical de N linhas: se o parágrafo tiver menos
  -- texto que isso, o rabo da letra desce por dentro do parágrafo seguinte. O
  -- número de linhas acompanha o tamanho do parágrafo. A medida aqui dá ~70
  -- caracteres por linha.
  local total = 0
  for _, inline in ipairs(content) do
    if inline.t == "Str" then total = total + #inline.text
    elseif inline.t == "Space" or inline.t == "SoftBreak" then total = total + 1
    end
  end
  -- Parágrafo de abertura curto demais para caber a capitular: deixa a linha
  -- como entrada e tenta o parágrafo seguinte.
  if total < 100 then
    if skipped < MAX_SKIP then
      skipped = skipped + 1
      return el                     -- awaiting_dropcap continua ligado
    end
    awaiting_dropcap = false
    return el
  end
  awaiting_dropcap = false

  local lines = (total < 200) and 2 or 3

  local first_char, rest = utf8_head(word)

  local opts = "lines=" .. lines .. ",findent=2pt,nindent=0pt"
  if ante ~= "" then
    opts = opts .. ",ante=" .. ante
  end

  local letrine = pandoc.RawInline("latex",
    "\\lettrine[" .. opts .. "]{" .. first_char .. "}{" .. rest .. "}")

  local new_inlines = pandoc.List()
  new_inlines:insert(letrine)
  for i = 2, #content do
    new_inlines:insert(content[i])
  end
  return pandoc.Para(new_inlines)
end
