# Instruções Rápidas - Batch ALIMENTAÇÃO Parte 2

## TL;DR - Execute Agora

```bash
# 1. Configure a API key (escolha um método)
export ANTHROPIC_API_KEY='sk-ant-api03-...'
# OU
echo 'sk-ant-api03-...' > ~/.anthropic_key

# 2. Execute o script
./scripts/execute_batch_alimentacao_p2.sh

# 3. Aguarde ~12 minutos

# 4. Aplique no banco
docker compose cp batch_alimentacao_parte2.sql db:/tmp/
docker compose exec db psql -U plenya_user -d plenya_db -f /tmp/batch_alimentacao_parte2.sql
```

## O Que Vai Acontecer

1. Script processa 20 items do grupo ALIMENTAÇÃO
2. Claude Opus 4.5 gera conteúdo clínico robusto
3. Cria 2 arquivos:
   - `batch_alimentacao_parte2_results.json` (dados completos)
   - `batch_alimentacao_parte2.sql` (migration)
4. Você aplica o SQL no banco
5. Items ficam enriquecidos no sistema

## Items Processados

```
1.  Qualidade alimentação parentes
2.  Infância
3.  Intolerâncias
4.  Introdução alimentar
5.  Lactose
6.  Alimentação parental pré/durante gestação
7.  Recordatório (24h)
8.  Onde e como come
9.  Ordem alimentos
10. Outros
11. Perfil metabólico parental
12. Piores períodos
13. Preferências
14. Proteína do leite
15. Quanto come
16. Quem cozinha
17. Regras alimentares
18. Relação com comida
19. Restrições pessoais
20. Satisfação pós-refeição
```

## Campos Enriquecidos (8 por item)

- question (pergunta melhorada)
- clinical_relevance
- interpretation_guide
- health_implications (array)
- followup_questions (array)
- red_flags (array)
- recommendations (array)
- scientific_background

## Custo e Tempo

- **Tempo**: 12-15 minutos
- **Custo**: ~$7 USD
- **Qualidade**: Conteúdo clínico baseado em evidências

## Validação Rápida

```bash
# Verificar quantos items foram processados
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT COUNT(*) as total_enriquecidos
FROM score_items
WHERE group_name = 'Alimentação'
  AND clinical_relevance IS NOT NULL;
"

# Ver exemplo de um item enriquecido
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT
  title,
  question,
  clinical_relevance
FROM score_items
WHERE title = 'Intolerâncias'
  AND clinical_relevance IS NOT NULL;
"
```

## Arquivos Criados

```
/home/user/plenya/
├── scripts/
│   ├── batch_alimentacao_parte2.py          (script principal)
│   └── execute_batch_alimentacao_p2.sh      (executor)
├── batch_alimentacao_parte2_results.json    (output JSON)
├── batch_alimentacao_parte2.sql             (output SQL)
├── BATCH-ALIMENTACAO-PARTE2-README.md       (doc completa)
├── BATCH-ALIMENTACAO-PARTE2-EXAMPLE.json    (exemplo)
├── BATCH-ALIMENTACAO-PARTE2-EXECUTIVE-SUMMARY.md
└── INSTRUCOES-RAPIDAS-ALIMENTACAO-P2.md     (este arquivo)
```

## Precisa de Ajuda?

Consulte:
- `BATCH-ALIMENTACAO-PARTE2-EXECUTIVE-SUMMARY.md` (resumo executivo)
- `BATCH-ALIMENTACAO-PARTE2-README.md` (documentação detalhada)
- `BATCH-ALIMENTACAO-PARTE2-EXAMPLE.json` (exemplo de output)

---

**Ready to go!** Configure a API key e execute o script.
