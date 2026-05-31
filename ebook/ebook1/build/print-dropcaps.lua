-- print-dropcaps.lua
-- Pandoc Lua filter: insert \lettrine{first-letter}{first-word-rest} on the
-- first paragraph after each numbered chapter heading. Skips Introdução,
-- Agradecimentos, Sobre o Autor, Referências (those are \chapter*{} so they
-- arrive as RawBlock LaTeX and don't trigger Header at all).

local awaiting_dropcap = false

local SKIP_TITLES = {
  ["Introdução"] = true,
  ["Agradecimentos"] = true,
  ["Sobre o Autor"] = true,
  ["Sobre o autor"] = true,
}

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

  -- Skip lettrine if the first paragraph is too short to absorb a 3-line drop
  -- cap. Below ~200 chars the lettrine descender hangs into the next paragraph
  -- (see Cap 13: "Ricardo voltou ao consultório dezoito meses depois do
  -- infarto." — 62 chars; Cap 17: "Eu tenho a mesma idade biológica que meus
  -- pacientes." — 52 chars). The 3-line vertical box needs ~3 lines × ~70
  -- chars/line ≈ 210 chars of running text to look right.
  local total = 0
  for _, inline in ipairs(el.content) do
    if inline.t == "Str" then total = total + #inline.text
    elseif inline.t == "Space" or inline.t == "SoftBreak" then total = total + 1
    end
  end
  if total < 200 then return el end

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
    "\\lettrine[lines=3,findent=2pt,nindent=0pt]{" .. first_char .. "}{" .. rest .. "}")

  -- Replace first inline with raw lettrine + the rest of the inlines (drop original first Str)
  local new_inlines = pandoc.List()
  new_inlines:insert(letrine)
  for i = 2, #el.content do
    new_inlines:insert(el.content[i])
  end
  return pandoc.Para(new_inlines)
end
