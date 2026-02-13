# Documentação Técnica Plenya

Esta pasta contém a documentação técnica modular do projeto Plenya EMR.

## Estrutura

```
.claude/
├── README.md                    # Este arquivo
├── 01-overview.md              # Visão geral do projeto
├── 02-stack.md                 # Stack técnica (a criar)
├── 03-architecture.md          # Arquitetura (a criar)
│
├── backend/                    # Documentação Backend Go
│   ├── models.md              # Padrões de Go models (a criar)
│   ├── database.md            # ⭐ Como trabalhar com banco direto
│   ├── hooks.md               # GORM hooks (a criar)
│   ├── service-layer.md       # Service patterns (a criar)
│   └── api-endpoints.md       # API endpoints (a criar)
│
├── frontend/                   # Documentação Frontend
│   ├── form-navigation.md     # useFormNavigation (a criar)
│   ├── patient-context.md     # useRequireSelectedPatient (a criar)
│   └── tanstack-query.md      # Query patterns (a criar)
│
├── domain/                     # Domínio de Negócio
│   ├── score-system.md        # ⭐ Sistema de Escores completo
│   ├── patients.md            # Workflows de pacientes (a criar)
│   └── security.md            # LGPD, segurança (a criar)
│
└── workflows/                  # Workflows Práticos
    ├── development.md          # Como desenvolver
    ├── database-ops.md         # ⭐ Operações diretas no banco
    └── adding-features.md      # Adicionar features (a criar)
```

## Arquivos Críticos (Já Criados)

### ⭐ database-ops.md
**Essencial para manipulação manual de dados durante desenvolvimento.**

Conteúdo:
- Como acessar banco via Docker
- SQL direto vs Go scripts vs API HTTP
- Exemplos práticos de CRUD
- Debugging de queries

**Quando ler:** Sempre que for manipular dados manualmente (adicionar score items, pacientes de teste, etc.)

### ⭐ score-system.md
**Documentação completa do sistema de escores (core feature).**

Conteúdo:
- Hierarquia de 4 níveis (Group → Subgroup → Item → Level)
- Filtros demográficos (gender, age, menopause)
- Operadores de comparação (=, >, >=, <, <=, between)
- Enriquecimento clínico (clinical_relevance, patient_explanation, conduct)
- Métodos de negócio (AppliesToPatient, EvaluatesTrue)
- Workflows práticos (duplicação, atualização)

**Quando ler:** Ao trabalhar com sistema de escores

### development.md
**Workflow diário de desenvolvimento.**

Conteúdo:
- Setup inicial
- Comandos Docker
- Hot reload
- Debugging
- Testes
- Migrations

**Quando ler:** Início de cada sessão de desenvolvimento

### 01-overview.md
**Visão geral do projeto, objetivos, stack resumida.**

Conteúdo:
- Objetivo do projeto
- Escala esperada
- Plataformas
- Diferenciais
- Roadmap
- Custos

**Quando ler:** Para entender o contexto geral

## Como Usar Esta Documentação

### Se você é Claude (IA)

1. **Leia CLAUDE.md primeiro** (raiz do projeto) para entender as regras de ouro
2. **Identifique o contexto da tarefa**:
   - Manipular dados? → `workflows/database-ops.md` + `domain/score-system.md`
   - Adicionar feature? → `backend/models.md` + `workflows/adding-features.md`
   - Trabalhar no frontend? → `frontend/*.md`
3. **Consulte arquivos específicos** conforme necessidade
4. **NUNCA adivinhe** - se não tiver certeza, leia a documentação relevante

### Se você é Humano

1. **CLAUDE.md** é o índice - comece por lá
2. **Esta pasta (.claude/)** tem detalhes técnicos
3. **Use busca** (`grep -r "termo" .claude/`) para encontrar informações
4. **Contribua** - se algo não está claro, melhore a documentação

## Princípios de Organização

Esta estrutura segue melhores práticas:

1. **Modularidade** - Arquivos pequenos e focados (< 500 linhas)
2. **Hierarquia** - Organização por área (backend, frontend, domain)
3. **Progressive Disclosure** - CLAUDE.md → overview → detalhes
4. **Searchability** - Fácil encontrar informações específicas
5. **Maintainability** - Fácil atualizar sem quebrar referências

## Referências

Baseado em:
- [Claude.md Guide](https://www.builder.io/blog/claude-md-guide)
- [Architecture Decision Records (ADR)](https://adr.github.io/)
- [Agents.md Pattern](https://layer5.io/blog/ai/agentsmd-one-file-to-guide-them-all/)

## Status de Implementação

✅ Criados:
- CLAUDE.md (raiz)
- .claude/README.md (este arquivo)
- .claude/01-overview.md
- .claude/workflows/development.md
- .claude/workflows/database-ops.md
- .claude/domain/score-system.md

🚧 A Criar:
- .claude/02-stack.md
- .claude/03-architecture.md
- .claude/backend/*.md (5 arquivos)
- .claude/frontend/*.md (3 arquivos)
- .claude/domain/patients.md
- .claude/domain/security.md
- .claude/workflows/adding-features.md

**Total:** 6/16 arquivos criados (37.5%)

## Migração do CLAUDE.md Original

O CLAUDE.md original (1000+ linhas) foi:
1. **Backup criado:** `CLAUDE.md.backup.<timestamp>`
2. **Substituído** por versão enxuta (índice)
3. **Conteúdo migrado** para arquivos modulares em `.claude/`

Prioridade de migração:
1. ✅ Operações de banco (database-ops.md)
2. ✅ Sistema de escores (score-system.md)
3. ✅ Development workflow
4. 🚧 Backend patterns (models, hooks, services)
5. 🚧 Frontend patterns (forms, context, queries)
6. 🚧 Security/LGPD

## Contribuindo

Ao adicionar/modificar documentação:

1. **Arquivos pequenos** - Se passar de 500 linhas, divida
2. **Links entre arquivos** - Use referências relativas `[texto](../path/file.md)`
3. **Exemplos práticos** - Sempre inclua código executável
4. **Atualizar índices** - CLAUDE.md e este README.md
5. **Testar links** - Garantir que referências não quebram

## Próximos Passos

1. Criar arquivos backend/ restantes
2. Criar arquivos frontend/ restantes
3. Completar domain/ (patients, security)
4. Criar workflows/adding-features.md
5. Revisar e consolidar
