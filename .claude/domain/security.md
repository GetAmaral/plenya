# Domain - Security & LGPD

## ⚠️ Dev Bypass Auth

**CRÍTICO:** Em desenvolvimento, existe modo de bypass de autenticação (`DEV_BYPASS_AUTH=true`) que injeta automaticamente `admin@plenya.com`.

- ✅ **Usar apenas em dev local**
- ❌ **NUNCA ativar em staging/produção**
- ❌ **NUNCA commitar .env com bypass=true**
- 📖 Ver: [dev-bypass-auth.md](../workflows/dev-bypass-auth.md)

## Autenticação

### JWT Tokens

- **Access Token:** 15 minutos
- **Refresh Token:** 7 dias
- **Algoritmo:** HS256
- **Secret:** 256 bits mínimo

```go
type AuthClaims struct {
    UserID string   `json:"userId"`
    Email  string   `json:"email"`
    Roles  []string `json:"roles"`
    jwt.RegisteredClaims
}
```

### Password Hashing

```go
// Bcrypt (cost 12+)
hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// Verificar
err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
```

## Autorização (RBAC)

### Roles

```go
const (
    RoleAdmin        = "admin"
    RoleDoctor       = "doctor"
    RoleNurse        = "nurse"
    RolePsychologist = "psychologist"
    RoleNutritionist = "nutritionist"
    RoleSecretary    = "secretary"
    RolePatient      = "patient"
)
```

### Middleware

```go
func RequireRole(allowedRoles ...models.Role) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userRoles := c.Locals("userRoles").([]string)
        
        for _, userRole := range userRoles {
            for _, allowedRole := range allowedRoles {
                if userRole == string(allowedRole) {
                    return c.Next()
                }
            }
        }
        
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "insufficient permissions",
        })
    }
}
```

## Criptografia (LGPD)

### Dados Sensíveis

**SEMPRE criptografar:**
- CPF
- RG
- Documentos
- Dados bancários

**Implementação via Hooks:**

```go
func (p *Patient) BeforeSave(tx *gorm.DB) error {
    if p.CPF != nil && *p.CPF != "" && !isEncrypted(*p.CPF) {
        encrypted, _ := crypto.EncryptWithDefaultKey(*p.CPF)
        p.CPF = &encrypted
    }
    return nil
}

func (p *Patient) AfterFind(tx *gorm.DB) error {
    if p.CPF != nil && *p.CPF != "" && isEncrypted(*p.CPF) {
        decrypted, _ := crypto.DecryptWithDefaultKey(*p.CPF)
        p.CPF = &decrypted
    }
    return nil
}
```

### PostgreSQL pgcrypto

```sql
-- Já habilitado no init.sql
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Função de criptografia
CREATE FUNCTION encrypt_data(data TEXT, key TEXT)
RETURNS TEXT AS $$
BEGIN
  RETURN encode(pgp_sym_encrypt(data, key), 'base64');
END;
$$ LANGUAGE plpgsql;

-- Função de descriptografia
CREATE FUNCTION decrypt_data(encrypted TEXT, key TEXT)
RETURNS TEXT AS $$
BEGIN
  RETURN pgp_sym_decrypt(decode(encrypted, 'base64'), key);
END;
$$ LANGUAGE plpgsql;
```

## Audit Logging

### Model

```go
type AuditLog struct {
    ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
    UserID     uuid.UUID `gorm:"type:uuid;not null;index"`
    Action     string    `gorm:"type:varchar(50);not null"` // CREATE, READ, UPDATE, DELETE
    Resource   string    `gorm:"type:varchar(100);not null"` // Patient, Anamnesis, etc.
    ResourceID string    `gorm:"type:varchar(255)"`
    IPAddress  string    `gorm:"type:varchar(45)"`
    UserAgent  string    `gorm:"type:text"`
    CreatedAt  time.Time `gorm:"autoCreateTime;index"`
}
```

### Middleware

```go
func AuditLogger() fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("userID").(uuid.UUID)
        
        auditLog := &models.AuditLog{
            UserID:     userID,
            Action:     c.Method(),
            Resource:   c.Path(),
            IPAddress:  c.IP(),
            UserAgent:  c.Get("User-Agent"),
        }
        
        db.Create(auditLog)
        
        return c.Next()
    }
}
```

### Retenção

- **Mínimo:** 5 anos (LGPD Art. 16)
- **Imutável:** Append-only
- **Backups:** Criptografados

## Rate Limiting

```go
func RateLimiter() fiber.Handler {
    return limiter.New(limiter.Config{
        Max:        100,              // 100 requests
        Expiration: 1 * time.Minute,  // por minuto
        KeyGenerator: func(c *fiber.Ctx) string {
            return c.IP()
        },
        LimitReached: func(c *fiber.Ctx) error {
            return c.Status(429).JSON(fiber.Map{
                "error": "too many requests",
            })
        },
    })
}
```

## CORS

```go
func CORS() fiber.Handler {
    return cors.New(cors.Config{
        AllowOrigins:     "http://localhost:3000,https://app.plenya.com.br",
        AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
        AllowCredentials: true,
        MaxAge:           86400,
    })
}
```

## Compliance LGPD

### Direitos do Titular

1. **Confirmação e Acesso** (Art. 18, I e II)
   - Endpoint: `GET /api/v1/patients/:id/data`
   
2. **Correção** (Art. 18, III)
   - Endpoint: `PUT /api/v1/patients/:id`
   
3. **Anonimização/Bloqueio/Eliminação** (Art. 18, IV)
   - Endpoint: `DELETE /api/v1/patients/:id`
   - Soft delete + anonimização após período
   
4. **Portabilidade** (Art. 18, V)
   - Endpoint: `GET /api/v1/patients/:id/export` (JSON/PDF)

### Consentimento

```go
type Consent struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    PatientID uuid.UUID `gorm:"type:uuid;not null"`
    Purpose   string    `gorm:"type:text;not null"`
    Granted   bool      `gorm:"type:boolean;not null"`
    GrantedAt time.Time `gorm:"type:timestamp"`
    RevokedAt *time.Time `gorm:"type:timestamp"`
}
```

### Backups Criptografados

```bash
# Backup diário com criptografia
pg_dump plenya_db | gpg --encrypt --recipient admin@plenya.com.br > backup_$(date +%Y%m%d).sql.gpg

# Upload para S3 Glacier
aws s3 cp backup_*.sql.gpg s3://plenya-backups/ --storage-class GLACIER
```

Ver: `apps/api/internal/middleware/auth.go`
