-- +goose Up
-- O que a ferramenta escreveu e o médico nunca reescreveu, marcado no ato da publicação.
--
-- As revisões já respondem "quem escreveu esta frase": basta achar a revisão mais recente cujo
-- `ops` toca aquele caminho. Mas isso é arqueologia — exige percorrer a cadeia inteira e cruzar
-- JSONB, e ninguém faz isso na pressa de uma consulta.
--
-- A pergunta que precisa ser barata é outra: "esta devolutiva, que o paciente já leu, contém texto
-- gerado por ferramenta que ninguém reescreveu depois?". Calculada no momento da publicação, ela
-- vira um campo. Depois, vira pesquisa.
--
-- Não é divulgação ao paciente: o médico assina e responde pelo documento, e essa decisão está
-- registrada no plano da feature. É registro interno, para auditoria e para medir se a ferramenta
-- está sendo revisada de fato — a evidência sobre revisão de conteúdo redigido por IA é de que a
-- releitura falha com frequência, e a única forma de saber se está falhando AQUI é medir.

ALTER TABLE public.patient_plan_revisions
  ADD COLUMN IF NOT EXISTS ai_touched_paths jsonb;

COMMENT ON COLUMN public.patient_plan_revisions.ai_touched_paths IS
  'Só em revisão de publicação: os caminhos cujo ÚLTIMO escritor foi o assistente, dentro desta versão. Vazio significa que tudo o que o paciente leu passou pela mão do médico depois da ferramenta.';

-- +goose Down
ALTER TABLE public.patient_plan_revisions DROP COLUMN IF EXISTS ai_touched_paths;
