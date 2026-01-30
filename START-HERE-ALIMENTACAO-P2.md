# 🚀 START HERE - Batch ALIMENTAÇÃO Parte 2

## Você tem 2 minutos? Leia isto primeiro.

```
┌─────────────────────────────────────────────────────────┐
│                                                         │
│  MISSÃO: Enriquecer 20 items do grupo ALIMENTAÇÃO      │
│  TEMPO: ~12 minutos de execução                         │
│  CUSTO: ~$6 USD                                         │
│  RESULTADO: 20 items com conteúdo clínico robusto      │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

---

## ⚡ Execução Rápida (3 comandos)

```bash
# 1. Configure sua API key Anthropic
export ANTHROPIC_API_KEY='sk-ant-api03-...'

# 2. Execute o script
./scripts/execute_batch_alimentacao_p2.sh

# 3. Aguarde ~12 minutos ☕
```

**Pronto!** Arquivos gerados:
- `batch_alimentacao_parte2_results.json` (dados completos)
- `batch_alimentacao_parte2.sql` (migration para PostgreSQL)

---

## 📚 Qual documentação ler?

### Quero executar AGORA (2 min)
👉 `INSTRUCOES-RAPIDAS-ALIMENTACAO-P2.md`

### Sou gestor/coordenador (15 min)
👉 `BATCH-ALIMENTACAO-PARTE2-EXECUTIVE-SUMMARY.md`

### Sou desenvolvedor (30 min)
👉 `BATCH-ALIMENTACAO-P2-TECHNICAL-SPEC.md`

### Preciso acompanhar progresso (5 min)
👉 `ALIMENTACAO-PROGRESS-TRACKER.md`

### Quero ver tudo organizado
👉 `BATCH-ALIMENTACAO-P2-INDEX.md`

---

## 🎯 O Que Vai Acontecer

1. **Script processa 20 items**
   - Cada item do grupo ALIMENTAÇÃO
   - Com contexto clínico específico

2. **Claude Opus 4.5 gera conteúdo**
   - 8 campos por item
   - Baseado em evidências científicas
   - Todo em Português

3. **Outputs são criados**
   - JSON com dados completos
   - SQL pronto para aplicar no banco

4. **Você aplica no banco**
   - Items ficam enriquecidos
   - Prontos para uso clínico

---

## 📋 20 Items que Serão Processados

```
✓ Qualidade alimentação parentes    ✓ Perfil metabólico parental
✓ Infância                           ✓ Piores períodos
✓ Intolerâncias                      ✓ Preferências
✓ Introdução alimentar               ✓ Proteína do leite
✓ Lactose                            ✓ Quanto come
✓ Alimentação parental pré-gestação  ✓ Quem cozinha
✓ Recordatório (24h)                 ✓ Regras alimentares
✓ Onde e como come                   ✓ Relação com comida
✓ Ordem alimentos                    ✓ Restrições pessoais
✓ Outros                             ✓ Satisfação pós-refeição
```

---

## 🔧 Pré-requisitos

```bash
# 1. Python 3.8+
python3 --version

# 2. Dependência anthropic
pip install anthropic

# 3. API key Anthropic
# Obtenha em: https://console.anthropic.com/
```

---

## ✅ Checklist Rápido

**Antes de executar**:
- [ ] Tenho API key Anthropic
- [ ] Instalei `pip install anthropic`
- [ ] Li instruções rápidas

**Durante execução**:
- [ ] Aguardar ~12 minutos
- [ ] Não interromper processo
- [ ] Verificar mensagens de erro

**Após execução**:
- [ ] Revisar JSON gerado
- [ ] Revisar SQL gerado
- [ ] Aplicar SQL no banco
- [ ] Validar resultados

---

## 🚨 Problemas Comuns

### "ANTHROPIC_API_KEY not found"
```bash
export ANTHROPIC_API_KEY='sk-ant-...'
# ou
echo 'sk-ant-...' > ~/.anthropic_key
```

### "anthropic package not installed"
```bash
pip install anthropic
```

### "Rate limit exceeded"
Script já tem delay de 2s. Aguarde e tente novamente.

---

## 📊 O Que Cada Item Ganha

8 campos enriquecidos:
1. ✍️ **question** - Pergunta clínica melhorada
2. 🏥 **clinical_relevance** - Por que isso importa
3. 📖 **interpretation_guide** - Como interpretar
4. 💊 **health_implications** - Implicações de saúde
5. ❓ **followup_questions** - Perguntas de follow-up
6. 🚨 **red_flags** - Sinais de alerta
7. 💡 **recommendations** - Recomendações práticas
8. 🔬 **scientific_background** - Contexto científico

---

## 💰 Custo e Tempo

| Métrica | Valor |
|---------|-------|
| Tempo de execução | ~12 minutos |
| Custo estimado | ~$6 USD |
| Items processados | 20 |
| Campos enriquecidos | 160 (20×8) |

---

## 📁 Arquivos Criados (Para Você)

### Documentação
```
INSTRUCOES-RAPIDAS-ALIMENTACAO-P2.md              (quick start)
BATCH-ALIMENTACAO-PARTE2-EXECUTIVE-SUMMARY.md     (resumo executivo)
BATCH-ALIMENTACAO-PARTE2-README.md                (referência)
ALIMENTACAO-PROGRESS-TRACKER.md                   (progresso)
BATCH-ALIMENTACAO-P2-TECHNICAL-SPEC.md            (técnico)
BATCH-ALIMENTACAO-P2-INDEX.md                     (índice)
BATCH-ALIMENTACAO-P2-DELIVERABLES.md              (deliverables)
START-HERE-ALIMENTACAO-P2.md                      (este arquivo)
```

### Scripts
```
scripts/batch_alimentacao_parte2.py               (processador)
scripts/execute_batch_alimentacao_p2.sh           (executor)
```

### Exemplo
```
BATCH-ALIMENTACAO-PARTE2-EXAMPLE.json             (amostra)
```

---

## 🎬 Próximos Passos (Após Execução)

### 1. Revisar Outputs
```bash
cat batch_alimentacao_parte2_results.json | jq '.[0]'
head -n 50 batch_alimentacao_parte2.sql
```

### 2. Aplicar no Banco
```bash
docker compose cp batch_alimentacao_parte2.sql db:/tmp/
docker compose exec db psql -U plenya_user -d plenya_db -f /tmp/batch_alimentacao_parte2.sql
```

### 3. Validar
```bash
docker compose exec db psql -U plenya_user -d plenya_db -c "
SELECT COUNT(*) FROM score_items
WHERE clinical_relevance IS NOT NULL
  AND group_name = 'Alimentação';
"
```

---

## 🆘 Precisa de Ajuda?

1. **Problemas técnicos**: Ver `TECHNICAL-SPEC.md`
2. **Dúvidas de execução**: Ver `EXECUTIVE-SUMMARY.md`
3. **Comandos rápidos**: Ver `INSTRUCOES-RAPIDAS-ALIMENTACAO-P2.md`
4. **Navegação geral**: Ver `INDEX.md`

---

## 🏁 Comece Agora

```bash
# Comando único para começar
./scripts/execute_batch_alimentacao_p2.sh
```

Ou siga o passo a passo em:
👉 **INSTRUCOES-RAPIDAS-ALIMENTACAO-P2.md**

---

**Criado**: 2026-01-27
**Status**: ✅ Pronto para uso
**Tempo para começar**: < 5 minutos

---

```
┌─────────────────────────────────────────────────────────┐
│                                                         │
│              🚀 VOCÊ ESTÁ PRONTO!                       │
│                                                         │
│  1. Configure API key                                   │
│  2. Execute o script                                    │
│  3. Aguarde ~12 minutos                                 │
│  4. Aplique no banco                                    │
│                                                         │
│              É isso. Simples assim.                     │
│                                                         │
└─────────────────────────────────────────────────────────┘
```
