-- fig-caption-numero.lua — a legenda impressa é só o número da figura.
--
-- O texto alternativo do markdown foi escrito para DESCREVER o desenho a quem ia
-- desenhá-lo ("Fluxograma da decisão clínica: HPB sintomática = indicação
-- aceita; alopecia cosmética = não prescrever..."). Ele serve como alt no EPUB,
-- mas impresso embaixo de um desenho que o leitor está olhando não serve. O
-- título continua dentro da arte; aqui fica só "Figura 8.1".
--
-- \caption{} vazio mantém o contador, que é o que garante numeração na ordem de
-- leitura (a numeração gravada dentro da arte era a ordem em que a figura foi
-- feita, e divergia).

function Figure(fig)
  fig.caption.long = pandoc.Blocks({})
  return fig
end
