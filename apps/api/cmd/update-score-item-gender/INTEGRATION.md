# Integração do Script - Score Item Gender Update

Este documento descreve diferentes formas de integrar o script de atualização de gênero ao pipeline do Plenya EMR.

## Opção 1: Script Standalone (Atual)

**Status**: ✅ Implementado

Executado manualmente ou via CI/CD quando necessário.

**Prós:**
- Controle total sobre quando executar
- Fácil de testar e debugar
- Não interfere com deploy automático
- Pode ser executado em staging antes de produção

**Contras:**
- Requer execução manual
- Pode ser esquecido em novos ambientes

**Como usar:**
```bash
docker compose exec api make update-gender
```

---

## Opção 2: Seed/Bootstrap (Recomendado para setup inicial)

Adicionar ao processo de seed inicial do banco de dados.

### Implementação

```go
// apps/api/internal/database/seed.go

package database

import (
    "log"
    "strings"

    "github.com/plenya/api/internal/models"
    "gorm.io/gorm"
)

// SeedScoreItemGender atualiza gender de todos score_items
func SeedScoreItemGender(db *gorm.DB) error {
    log.Println("🔄 Seeding score_item gender fields...")

    var items []models.ScoreItem
    if err := db.Find(&items).Error; err != nil {
        return err
    }

    updated := 0
    for _, item := range items {
        gender := detectGenderFromName(item.Name)

        // Só atualizar se for diferente do atual
        currentGender := "not_applicable"
        if item.Gender != nil {
            currentGender = *item.Gender
        }

        if currentGender != gender {
            if err := db.Model(&item).Update("gender", gender).Error; err != nil {
                log.Printf("⚠️  Failed to update item %s: %v", item.ID, err)
                continue
            }
            updated++
        }
    }

    log.Printf("✅ Updated %d score_item gender fields", updated)
    return nil
}

func detectGenderFromName(name string) string {
    nameLower := strings.ToLower(name)

    maleKeywords := []string{
        "homem", "homens", "masculino", "masculina",
        "dos homens", "para homens", "no homem",
    }
    for _, kw := range maleKeywords {
        if strings.Contains(nameLower, kw) {
            return "male"
        }
    }

    femaleKeywords := []string{
        "mulher", "mulheres", "feminino", "feminina",
        "das mulheres", "para mulheres", "na mulher",
    }
    for _, kw := range femaleKeywords {
        if strings.Contains(nameLower, kw) {
            return "female"
        }
    }

    return "not_applicable"
}
```

### Chamar no bootstrap

```go
// apps/api/cmd/server/main.go

func main() {
    // ... código existente ...

    if cfg.Server.Environment == "development" {
        log.Println("🔄 Running auto migrations...")
        if err := database.AutoMigrate(); err != nil {
            log.Fatalf("❌ Failed to auto migrate: %v", err)
        }

        // Seed gender após migrations
        if err := database.SeedScoreItemGender(database.DB); err != nil {
            log.Printf("⚠️  Failed to seed gender: %v", err)
        }
    }

    // ... resto do código ...
}
```

---

## Opção 3: Migration SQL (Para ambientes de produção)

Criar uma migration Atlas que atualiza os dados.

### Criar migration

```bash
cd apps/api
atlas migrate new update_score_item_gender --env dev
```

### Conteúdo da migration

```sql
-- apps/api/database/migrations/20260208000001_update_score_item_gender.sql

-- Update male items
UPDATE score_items
SET gender = 'male'
WHERE deleted_at IS NULL
  AND (
    LOWER(name) LIKE '%homem%'
    OR LOWER(name) LIKE '%homens%'
    OR LOWER(name) LIKE '%masculin%'
    OR LOWER(name) LIKE '%dos homens%'
    OR LOWER(name) LIKE '%para homens%'
    OR LOWER(name) LIKE '%no homem%'
    OR LOWER(name) LIKE '%nos homens%'
  );

-- Update female items
UPDATE score_items
SET gender = 'female'
WHERE deleted_at IS NULL
  AND (
    LOWER(name) LIKE '%mulher%'
    OR LOWER(name) LIKE '%mulheres%'
    OR LOWER(name) LIKE '%feminin%'
    OR LOWER(name) LIKE '%das mulheres%'
    OR LOWER(name) LIKE '%para mulheres%'
    OR LOWER(name) LIKE '%na mulher%'
    OR LOWER(name) LIKE '%nas mulheres%'
  );

-- Ensure not_applicable for items without gender markers
UPDATE score_items
SET gender = 'not_applicable'
WHERE deleted_at IS NULL
  AND gender IS NULL;
```

### Aplicar migration

```bash
atlas migrate apply --env dev
```

---

## Opção 4: Hook GORM BeforeSave (Automático)

Atualizar automaticamente ao criar/editar um `ScoreItem`.

### Implementação

```go
// apps/api/internal/models/score_item.go

// BeforeSave hook to auto-detect gender from name
func (si *ScoreItem) BeforeSave(tx *gorm.DB) error {
    // Auto-detect gender se não foi explicitamente definido
    if si.Gender == nil || *si.Gender == "" {
        gender := detectGenderFromName(si.Name)
        si.Gender = &gender
    }
    return nil
}

func detectGenderFromName(name string) string {
    nameLower := strings.ToLower(name)

    maleKeywords := []string{"homem", "homens", "masculino"}
    for _, kw := range maleKeywords {
        if strings.Contains(nameLower, kw) {
            return "male"
        }
    }

    femaleKeywords := []string{"mulher", "mulheres", "feminino"}
    for _, kw := range femaleKeywords {
        if strings.Contains(nameLower, kw) {
            return "female"
        }
    }

    return "not_applicable"
}
```

**Prós:**
- Automático para novos items
- Garante consistência

**Contras:**
- Pode sobrescrever valores definidos manualmente
- Adiciona lógica ao model

---

## Opção 5: Scheduled Job (Manutenção periódica)

Executar automaticamente em intervalos (ex: diariamente).

### Implementação

```go
// apps/api/internal/scheduler/gender_update_job.go

package scheduler

import (
    "log"
    "time"

    "github.com/plenya/api/internal/models"
    "gorm.io/gorm"
)

type GenderUpdateJob struct {
    db *gorm.DB
}

func NewGenderUpdateJob(db *gorm.DB) *GenderUpdateJob {
    return &GenderUpdateJob{db: db}
}

func (j *GenderUpdateJob) Start() {
    // Executar a cada 24 horas
    ticker := time.NewTicker(24 * time.Hour)
    go func() {
        // Executar imediatamente na inicialização
        j.Run()

        // Depois executar periodicamente
        for range ticker.C {
            j.Run()
        }
    }()
}

func (j *GenderUpdateJob) Run() {
    log.Println("🔄 Running scheduled gender update job...")

    var items []models.ScoreItem
    if err := j.db.Find(&items).Error; err != nil {
        log.Printf("⚠️  Failed to fetch items: %v", err)
        return
    }

    updated := 0
    for _, item := range items {
        gender := detectGenderFromName(item.Name)

        currentGender := "not_applicable"
        if item.Gender != nil {
            currentGender = *item.Gender
        }

        if currentGender != gender {
            if err := j.db.Model(&item).Update("gender", gender).Error; err != nil {
                continue
            }
            updated++
        }
    }

    log.Printf("✅ Gender update job completed: %d items updated", updated)
}
```

### Adicionar ao main.go

```go
// apps/api/cmd/server/main.go

func main() {
    // ... código existente ...

    // Iniciar scheduler de atualização de gênero
    genderUpdateJob := scheduler.NewGenderUpdateJob(database.DB)
    genderUpdateJob.Start()

    // ... resto do código ...
}
```

---

## Opção 6: API Endpoint (Admin)

Criar endpoint para admins executarem manualmente via UI.

### Implementação

```go
// apps/api/internal/handlers/admin_handler.go

func (h *AdminHandler) UpdateScoreItemGender(c *fiber.Ctx) error {
    // Buscar todos os items
    var items []models.ScoreItem
    if err := h.db.Find(&items).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{
            "error": "Failed to fetch items",
        })
    }

    updated := 0
    errors := 0

    for _, item := range items {
        gender := detectGenderFromName(item.Name)

        if err := h.db.Model(&item).Update("gender", gender).Error; err != nil {
            errors++
            continue
        }
        updated++
    }

    return c.JSON(fiber.Map{
        "message": "Gender update completed",
        "updated": updated,
        "errors":  errors,
        "total":   len(items),
    })
}
```

### Rota

```go
// apps/api/cmd/server/main.go

admin := v1.Group("/admin")
admin.Use(middleware.Auth(cfg))
admin.Use(middleware.RequireAdmin())
admin.Post("/score-items/update-gender", adminHandler.UpdateScoreItemGender)
```

---

## Recomendações

Para o Plenya EMR, recomendo a seguinte estratégia:

1. **Implementar Opção 2 (Seed)** para setup inicial de novos ambientes
2. **Manter Opção 1 (Script Standalone)** para execuções pontuais e troubleshooting
3. **Considerar Opção 4 (Hook)** para garantir consistência em novos items
4. **Usar Opção 3 (Migration SQL)** para deploy em produção

### Roadmap de Implementação

1. **Fase Atual**: ✅ Script standalone funcional
2. **Próximo passo**: Adicionar ao seed.go
3. **Futuro**: Hook automático no model
4. **Produção**: Migration SQL antes de deploy

---

## Testing

Independente da opção escolhida, sempre:

1. **Backup antes de executar**
2. **Testar em ambiente de desenvolvimento primeiro**
3. **Verificar resultados com queries SQL**
4. **Monitorar logs durante execução**

```bash
# Exemplo de teste completo
docker compose exec db pg_dump -U plenya_user plenya_db > backup.sql
docker compose exec api make update-gender
docker compose exec db psql -U plenya_user -d plenya_db -c \
  "SELECT gender, COUNT(*) FROM score_items GROUP BY gender;"
```
