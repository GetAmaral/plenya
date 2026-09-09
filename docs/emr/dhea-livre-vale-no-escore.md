# DHEA (não sulfatado) vale um item de escore?

**Não.** Vale um cadastro no catálogo de exames, sem item de escore. Este documento é o porquê,
escrito quando um laudo da paciente trouxe "DHEA - DEHIDROEPIANDROSTERONA 1,4 ng/mL" no lugar do
DHEA-S que tinha sido pedido, e a carga ficou sem onde pousar.

## O que o escore precisa de um marcador

Um item de escore compara um número contra uma escala fixa e devolve um nível. Para isso funcionar,
o mesmo estado clínico tem de produzir aproximadamente o mesmo número em coletas diferentes. É essa
propriedade, e não a importância biológica da molécula, que decide se algo pode virar item.

## Por que o DHEA não tem essa propriedade e o DHEA-S tem

| | DHEA | DHEA-S |
|---|---|---|
| meia-vida circulante | cerca de 25 minutos | 10 a 16 horas |
| ritmo circadiano | sim, acoplado ao ACTH | não |
| secreção | em pulsos | estável ao longo do dia |
| concentração relativa | cerca de 1% | cerca de 99% do androgênio adrenal circulante |

Duas coletas do mesmo paciente, uma às 8h e outra às 11h, dão DHEA diferentes sem que nada tenha
mudado nele. Num escore isso não vira sinal, vira ruído com aparência de sinal: o paciente veria o
nível subir e descer entre consultas por causa do horário da punção. O DHEA-S não tem esse problema,
e é por isso que ele é o marcador de androgênio adrenal na prática clínica.

## O que já existe no catálogo

`PLN82DBE091` (DHEA-S, µg/dL) tem **doze itens de escore**, um por sexo e década de vida, cada um com
seis faixas. É uma curadoria feita, calibrada e em uso. Duplicá-la para o DHEA exigiria construir do
zero doze escalas para um marcador que oscila dentro do próprio dia, e o resultado seria pior que o
que já está lá.

## O que fazer então

1. **Cadastrar `DHEA` no catálogo de exames, sem item de escore.** O resultado fica arquivado no
   prontuário e visível na linha do tempo, sem entrar na conta do escore. É o mesmo tratamento que a
   glicemia média estimada e o HOMA-BETA já recebem.
2. **Proteger o mapeamento.** O `alt_names` do DHEA-S contém "dehidroepiandrosterona sulfato" e
   "dheas". Um laudo escrito "DHEA - DEHIDROEPIANDROSTERONA" está a uma aproximação de distância de
   cair no item errado, e cairia em µg/dL contra uma escala de µg/dL, com o número de ng/mL: 1,4 seria
   lido como profundamente baixo numa escala em que a faixa de mulher de 60-69 anos começa em 10.
   Nenhuma rede de plausibilidade pega isso, porque 1,4 µg/dL é um valor possível. O cadastro
   próprio, com `alt_names` que reivindica as formas sem "sulfato", é o que evita a colisão.
3. **Repetir o DHEA-S quando o laudo vier trocado.** É o exame que responde à pergunta clínica, e é
   o que o escore lê.

## Referência

- Endotext, *Adrenal Androgens*: DHEA com meia-vida de cerca de 25 minutos e secreção circadiana
  semelhante à do ACTH; DHEA-S sem ritmo circadiano pela meia-vida longa do esteroide sulfatado.
  <https://www.ncbi.nlm.nih.gov/sites/books/NBK278929/>
- *Low DHEAS: A Sensitive and Specific Test for the Detection of Subclinical Hypercortisolism*,
  JCEM 2017: meia-vida sérica do DHEA-S de 10 a 16 horas, com níveis estáveis ao longo do dia.
  <https://academic.oup.com/jcem/article/102/3/786/3061900>
- Mayo Clinic Laboratories, DHEA sérico: na maioria das situações clínicas DHEA e DHEA-S são
  intercambiáveis, e o sulfato circula em concentração mais de cem vezes maior.
  <https://www.mayocliniclabs.com/test-catalog/overview/81405>
