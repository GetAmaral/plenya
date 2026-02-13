# Backend - GORM Hooks

## Visão Geral

GORM hooks são callbacks executados automaticamente em momentos específicos do ciclo de vida de um model.

## Hooks Disponíveis (Ordem de Execução)

### Create Operation

```
BeforeSave
  ↓
BeforeCreate
  ↓
[INSERT SQL]
  ↓
AfterCreate
  ↓
AfterSave
```

### Update Operation

```
BeforeSave
  ↓
BeforeUpdate
  ↓
[UPDATE SQL]
  ↓
AfterUpdate
  ↓
AfterSave
```

### Query Operation

```
[SELECT SQL]
  ↓
AfterFind
```

### Delete Operation

```
BeforeDelete
  ↓
[DELETE/UPDATE SQL]  (soft delete = UPDATE deleted_at)
  ↓
AfterDelete
```

## Hook Obrigatório: BeforeCreate (UUID v7)

**TODOS os models DEVEM ter:**

```go
func (m *MyModel) BeforeCreate(tx *gorm.DB) error {
    if m.ID == uuid.Nil {
        m.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}
```

**Por quê?**
- Garante UUID v7 (time-ordered)
- Gerado pela aplicação (não pelo DB)
- Consistente em todos os models

## Hook Pattern: BeforeSave (Validation/Transformation)

### Caso 1: Criptografia de Dados Sensíveis

**Patient model - Criptografar CPF/RG:**

```go
func (p *Patient) BeforeSave(tx *gorm.DB) error {
    // Criptografar CPF se não estiver criptografado
    if p.CPF != nil && *p.CPF != "" && !isEncrypted(*p.CPF) {
        encrypted, err := crypto.EncryptWithDefaultKey(*p.CPF)
        if err != nil {
            return fmt.Errorf("failed to encrypt CPF: %w", err)
        }
        p.CPF = &encrypted
    }

    // Criptografar RG
    if p.RG != nil && *p.RG != "" && !isEncrypted(*p.RG) {
        encrypted, err := crypto.EncryptWithDefaultKey(*p.RG)
        if err != nil {
            return fmt.Errorf("failed to encrypt RG: %w", err)
        }
        p.RG = &encrypted
    }

    // Recalcular idade antes de salvar
    p.CalculateAge()

    return nil
}

// Helper para detectar se string está criptografada
func isEncrypted(s string) bool {
    // Assumindo que string criptografada tem formato específico
    // Ex: começa com "ENC:" ou é base64 longo
    return strings.HasPrefix(s, "ENC:") || len(s) > 100
}
```

### Caso 2: Normalização de Dados

```go
func (u *User) BeforeSave(tx *gorm.DB) error {
    // Normalizar email para lowercase
    if u.Email != "" {
        u.Email = strings.ToLower(strings.TrimSpace(u.Email))
    }

    // Normalizar telefone (remover caracteres especiais)
    if u.Phone != nil && *u.Phone != "" {
        normalized := strings.ReplaceAll(*u.Phone, "-", "")
        normalized = strings.ReplaceAll(normalized, " ", "")
        normalized = strings.ReplaceAll(normalized, "(", "")
        normalized = strings.ReplaceAll(normalized, ")", "")
        u.Phone = &normalized
    }

    return nil
}
```

### Caso 3: Hash de Senha

```go
func (u *User) BeforeSave(tx *gorm.DB) error {
    // Apenas hash se senha mudou
    if tx.Statement.Changed("Password") {
        hashedPassword, err := bcrypt.GenerateFromPassword(
            []byte(u.Password),
            bcrypt.DefaultCost,
        )
        if err != nil {
            return fmt.Errorf("failed to hash password: %w", err)
        }
        u.Password = string(hashedPassword)
    }

    return nil
}
```

## Hook Pattern: BeforeUpdate (Auto-Update Fields)

### Caso 1: LastReview Automático (ScoreItem, ScoreLevel)

**Atualiza LastReview quando campos clínicos mudam:**

```go
func (si *ScoreItem) BeforeUpdate(tx *gorm.DB) error {
    // Verificar se campos clínicos mudaram
    if tx.Statement.Changed("ClinicalRelevance") ||
       tx.Statement.Changed("PatientExplanation") ||
       tx.Statement.Changed("Conduct") {
        now := time.Now()
        si.LastReview = &now
    }

    return nil
}
```

**ScoreLevel - mesmo pattern:**

```go
func (sl *ScoreLevel) BeforeUpdate(tx *gorm.DB) error {
    if tx.Statement.Changed("ClinicalRelevance") ||
       tx.Statement.Changed("PatientExplanation") ||
       tx.Statement.Changed("Conduct") {
        now := time.Now()
        sl.LastReview = &now
    }

    return nil
}
```

### Caso 2: Versioning

```go
type Document struct {
    ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
    Content string    `gorm:"type:text"`
    Version int       `gorm:"type:integer;not null;default:1"`
}

func (d *Document) BeforeUpdate(tx *gorm.DB) error {
    // Incrementar versão se conteúdo mudou
    if tx.Statement.Changed("Content") {
        d.Version++
    }

    return nil
}
```

## Hook Pattern: AfterFind (Post-Processing)

### Caso 1: Descriptografia

**Patient model - Descriptografar CPF/RG após buscar:**

```go
func (p *Patient) AfterFind(tx *gorm.DB) error {
    // Descriptografar CPF
    if p.CPF != nil && *p.CPF != "" && isEncrypted(*p.CPF) {
        decrypted, err := crypto.DecryptWithDefaultKey(*p.CPF)
        if err != nil {
            // Log erro mas não falha (manter criptografado)
            log.Warn().Err(err).Msg("Failed to decrypt CPF")
            return nil
        }
        p.CPF = &decrypted
    }

    // Descriptografar RG
    if p.RG != nil && *p.RG != "" && isEncrypted(*p.RG) {
        decrypted, err := crypto.DecryptWithDefaultKey(*p.RG)
        if err != nil {
            log.Warn().Err(err).Msg("Failed to decrypt RG")
            return nil
        }
        p.RG = &decrypted
    }

    return nil
}
```

### Caso 2: Calcular Campos Derivados

```go
func (p *Patient) AfterFind(tx *gorm.DB) error {
    // Recalcular idade
    p.CalculateAge()

    return nil
}

func (p *Patient) CalculateAge() {
    if !p.BirthDate.IsZero() {
        years := time.Since(p.BirthDate).Hours() / 24 / 365.25
        p.Age = int(years)
    }
}
```

### Caso 3: Carregar Dados Relacionados Customizados

```go
func (p *Patient) AfterFind(tx *gorm.DB) error {
    // Carregar estatísticas agregadas
    var count int64
    tx.Model(&Anamnesis{}).Where("patient_id = ?", p.ID).Count(&count)
    p.AnamnesisCount = int(count)

    return nil
}
```

## Hook Pattern: BeforeDelete (Cascade Manual)

```go
func (p *Patient) BeforeDelete(tx *gorm.DB) error {
    // Deletar dados relacionados que não têm CASCADE
    if err := tx.Where("patient_id = ?", p.ID).Delete(&CustomData{}).Error; err != nil {
        return err
    }

    // Audit log
    auditLog := &AuditLog{
        UserID:     getUserFromContext(tx), // Helper function
        Action:     "DELETE",
        Resource:   "Patient",
        ResourceID: p.ID.String(),
    }
    if err := tx.Create(auditLog).Error; err != nil {
        return err
    }

    return nil
}
```

## Hooks Combinados: Patient (Exemplo Completo)

```go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
    "github.com/plenya/api/pkg/crypto"
)

type Patient struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(200);not null" json:"name"`
    BirthDate time.Time `gorm:"type:date;not null" json:"birthDate"`
    CPF       *string   `gorm:"type:text" json:"cpf,omitempty"`
    RG        *string   `gorm:"type:text" json:"rg,omitempty"`
    Age       int       `gorm:"-" json:"age"` // Não persistido

    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Hook 1: Gerar UUID v7 (OBRIGATÓRIO)
func (p *Patient) BeforeCreate(tx *gorm.DB) error {
    if p.ID == uuid.Nil {
        p.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}

// Hook 2: Criptografar dados sensíveis + calcular idade
func (p *Patient) BeforeSave(tx *gorm.DB) error {
    // Criptografar CPF
    if p.CPF != nil && *p.CPF != "" && !isEncrypted(*p.CPF) {
        encrypted, err := crypto.EncryptWithDefaultKey(*p.CPF)
        if err != nil {
            return err
        }
        p.CPF = &encrypted
    }

    // Criptografar RG
    if p.RG != nil && *p.RG != "" && !isEncrypted(*p.RG) {
        encrypted, err := crypto.EncryptWithDefaultKey(*p.RG)
        if err != nil {
            return err
        }
        p.RG = &encrypted
    }

    // Recalcular idade
    p.CalculateAge()

    return nil
}

// Hook 3: Descriptografar após buscar
func (p *Patient) AfterFind(tx *gorm.DB) error {
    // Descriptografar CPF
    if p.CPF != nil && *p.CPF != "" && isEncrypted(*p.CPF) {
        decrypted, err := crypto.DecryptWithDefaultKey(*p.CPF)
        if err != nil {
            // Não falhar, apenas logar
            return nil
        }
        p.CPF = &decrypted
    }

    // Descriptografar RG
    if p.RG != nil && *p.RG != "" && isEncrypted(*p.RG) {
        decrypted, err := crypto.DecryptWithDefaultKey(*p.RG)
        if err != nil {
            return nil
        }
        p.RG = &decrypted
    }

    return nil
}

// Método auxiliar
func (p *Patient) CalculateAge() {
    if !p.BirthDate.IsZero() {
        years := time.Since(p.BirthDate).Hours() / 24 / 365.25
        p.Age = int(years)
    }
}

func isEncrypted(s string) bool {
    return strings.HasPrefix(s, "ENC:") || len(s) > 100
}
```

## Verificar Mudanças em Campos

```go
func (m *Model) BeforeUpdate(tx *gorm.DB) error {
    // Verificar se campo específico mudou
    if tx.Statement.Changed("FieldName") {
        // Fazer algo
    }

    // Verificar múltiplos campos
    if tx.Statement.Changed("Field1") || tx.Statement.Changed("Field2") {
        // Fazer algo
    }

    // Pegar valor antigo
    oldValue := tx.Statement.Dest.(*Model).FieldName

    // Comparar
    if tx.Statement.Changed("Status") && oldValue != m.Status {
        // Status mudou
    }

    return nil
}
```

## Acessar Contexto no Hook

```go
func (m *Model) BeforeCreate(tx *gorm.DB) error {
    // Pegar user ID do contexto (setado pelo middleware)
    if userID, ok := tx.Statement.Context.Value("userID").(uuid.UUID); ok {
        m.CreatedBy = userID
    }

    return nil
}
```

**Middleware que seta contexto:**

```go
func Auth(cfg *config.Config) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // ... validar JWT ...

        // Setar no contexto Fiber
        c.Locals("userID", userID)

        // Passar para GORM via context
        db := c.Locals("db").(*gorm.DB)
        db = db.WithContext(context.WithValue(c.Context(), "userID", userID))
        c.Locals("db", db)

        return c.Next()
    }
}
```

## Regras e Boas Práticas

### ✅ DO

- Usar BeforeCreate para UUID v7 (obrigatório)
- Usar BeforeSave para criptografia/normalização
- Usar BeforeUpdate para campos auto-atualizados
- Usar AfterFind para descriptografia/cálculos
- Retornar erros quando operação deve falhar
- Logar avisos quando descriptografia falha (mas não retornar erro)

### ❌ DON'T

- NÃO fazer queries pesadas em hooks (N+1 problem)
- NÃO fazer chamadas HTTP/externas (performance)
- NÃO modificar outros models (usar transactions explícitas)
- NÃO ignorar erros silenciosamente (exceto descriptografia)
- NÃO usar hooks para lógica de apresentação

### Performance

```go
// ❌ BAD: N+1 queries
func (p *Patient) AfterFind(tx *gorm.DB) error {
    // Executado para CADA patient!
    tx.Where("patient_id = ?", p.ID).Find(&p.Anamneses)
    return nil
}

// ✅ GOOD: Use Preload
db.Preload("Anamneses").Find(&patients)
```

## Hooks vs Service Layer

**Quando usar hooks:**
- Geração automática (UUID, timestamps)
- Criptografia/descriptografia
- Normalização de dados
- Campos calculados simples

**Quando usar service layer:**
- Validação de regras de negócio complexas
- Operações que envolvem múltiplos models
- Lógica que precisa ser testável isoladamente
- Operações que podem falhar de forma esperada

**Exemplo:**

```go
// ✅ HOOK: Geração automática
func (p *Patient) BeforeCreate(tx *gorm.DB) error {
    if p.ID == uuid.Nil {
        p.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}

// ✅ SERVICE: Validação de negócio
func (s *PatientService) Create(dto CreatePatientDTO) (*models.Patient, error) {
    // Verificar se CPF já existe
    exists, err := s.repo.ExistsByCPF(dto.CPF)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("CPF already registered")
    }

    // Criar patient (hooks executam automaticamente)
    patient := &models.Patient{
        Name: dto.Name,
        CPF:  &dto.CPF,
    }

    if err := s.repo.Create(patient); err != nil {
        return nil, err
    }

    return patient, nil
}
```

## Debugging Hooks

```go
func (m *Model) BeforeCreate(tx *gorm.DB) error {
    log.Debug().
        Str("model", "Model").
        Str("hook", "BeforeCreate").
        Interface("data", m).
        Msg("Hook executing")

    // ... lógica do hook ...

    log.Debug().
        Str("model", "Model").
        Str("hook", "BeforeCreate").
        Msg("Hook completed")

    return nil
}
```

## Desabilitar Hooks Temporariamente

```go
// Desabilitar hooks em operação específica
db.Session(&gorm.Session{SkipHooks: true}).Create(&patient)

// Desabilitar apenas um hook (não há suporte direto, usar flag)
type Patient struct {
    // ... campos ...
    skipEncryption bool `gorm:"-"` // Não persistido
}

func (p *Patient) BeforeSave(tx *gorm.DB) error {
    if p.skipEncryption {
        return nil
    }

    // ... criptografia ...
    return nil
}

// Uso
patient.skipEncryption = true
db.Create(&patient)
```

## Exemplos de Hooks no Projeto

Ver arquivos:
- `apps/api/internal/models/patient.go` - Criptografia
- `apps/api/internal/models/score_item.go` - LastReview auto-update
- `apps/api/internal/models/score_level.go` - LastReview auto-update
- Todos os models - BeforeCreate para UUID v7
