-- linha-de-destaque.lua — linha de destaque não hifeniza.
--
-- As oito linhas "***Pilar X — ...***" que fecham a transição de cada capítulo
-- são títulos, não texto corrido, e estavam hifenizando como texto corrido:
-- "Atividade Física, Alimentação e Suplementação Inteli-" / "gente" (p.120 e
-- p.137), com a palavra "gente" sozinha na linha de baixo. O template já proíbe
-- hífen em \part, \chapter, \section, legenda e sumário via \notitlehyphen; aqui
-- é a mesma regra para o parágrafo que é só um destaque.
--
-- O grupo em volta é necessário: \hyphenpenalty é lido na hora de quebrar o
-- parágrafo, então o \par tem de ficar DENTRO do grupo, senão o ajuste vaza para
-- os parágrafos seguintes.

local function unico(inlines, tipo)
  return inlines and #inlines == 1 and inlines[1].t == tipo
end

function Para(el)
  local c = el.content
  local destaque =
    (unico(c, "Strong") and unico(c[1].content, "Emph")) or
    (unico(c, "Emph") and unico(c[1].content, "Strong"))
  if not destaque then return el end

  local novo = pandoc.List()
  novo:insert(pandoc.RawInline("latex", "{\\notitlehyphen "))
  for _, x in ipairs(c) do novo:insert(x) end
  novo:insert(pandoc.RawInline("latex", "\\par}"))
  return pandoc.Para(novo)
end
