# Backend - API Endpoints

## Quando Usar API HTTP

| Contexto | Método | Exemplo |
|----------|--------|---------|
| **Apps (web/mobile)** | API HTTP | `POST /api/v1/patients` |
| **Desenvolvimento manual** | Docker/psql | `docker compose exec db psql...` |
| **Scripts Go** | Repository direto | `repo.CreatePatient(patient)` |
| **Testes** | Repository direto | `repo.GetByID(id)` |

**Regra:** API HTTP é para aplicações finais, NÃO para desenvolvimento manual.

## Endpoints Principais

### Auth

```
POST   /api/v1/auth/register      Criar conta
POST   /api/v1/auth/login         Login (JWT)
POST   /api/v1/auth/refresh       Renovar token
POST   /api/v1/auth/2fa/enable    Habilitar 2FA
POST   /api/v1/auth/2fa/verify    Verificar código 2FA
```

### Patients

```
GET    /api/v1/patients           Listar (paginado)
GET    /api/v1/patients/:id       Detalhes
POST   /api/v1/patients           Criar
PUT    /api/v1/patients/:id       Atualizar
DELETE /api/v1/patients/:id       Soft delete
```

### Score Groups

```
GET    /api/v1/score-groups              Listar todos
GET    /api/v1/score-groups/:id          Detalhes
GET    /api/v1/score-groups/:id/tree     Com nested data
GET    /api/v1/score-groups/tree         Todos com nested (mindmap)
POST   /api/v1/score-groups              Criar (admin)
PUT    /api/v1/score-groups/:id          Atualizar (admin)
DELETE /api/v1/score-groups/:id          Soft delete (admin)
```

### Score Subgroups

```
GET    /api/v1/score-groups/:groupId/subgroups   Por grupo
GET    /api/v1/score-subgroups/:id                Detalhes
POST   /api/v1/score-subgroups                    Criar (admin)
PUT    /api/v1/score-subgroups/:id                Atualizar (admin)
DELETE /api/v1/score-subgroups/:id                Soft delete (admin)
```

### Score Items

```
GET    /api/v1/score-items                        Listar todos
GET    /api/v1/score-subgroups/:subgroupId/items  Por subgrupo
GET    /api/v1/score-items/:id                    Detalhes
POST   /api/v1/score-items                        Criar (admin)
PUT    /api/v1/score-items/:id                    Atualizar (admin)
DELETE /api/v1/score-items/:id                    Soft delete (admin)
POST   /api/v1/score-items/:id/duplicate          Duplicar item
```

### Score Levels

```
GET    /api/v1/score-items/:itemId/levels   Por item
GET    /api/v1/score-levels/:id             Detalhes
POST   /api/v1/score-levels                 Criar (admin)
PUT    /api/v1/score-levels/:id             Atualizar (admin)
DELETE /api/v1/score-levels/:id             Soft delete (admin)
```

## Handler Pattern Completo

```go
// CreateScoreItem godoc
// @Summary Create a new score item
// @Description Create a new score item (admin only)
// @Tags Score Items
// @Accept json
// @Produce json
// @Param body body services.CreateScoreItemDTO true "Score Item Data"
// @Success 201 {object} models.ScoreItem
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items [post]
func (h *ScoreHandler) CreateScoreItem(c *fiber.Ctx) error {
    // 1. Parse request body
    var dto services.CreateScoreItemDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // 2. Validar DTO
    if err := h.validator.Struct(dto); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // 3. Chamar service
    item, err := h.service.CreateItem(dto)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // 4. Retornar 201 Created
    return c.Status(http.StatusCreated).JSON(item)
}
```

## Status Codes

- `200 OK` - Sucesso (GET, PUT)
- `201 Created` - Recurso criado (POST)
- `204 No Content` - Sucesso sem body (DELETE)
- `400 Bad Request` - Validação falhou
- `401 Unauthorized` - Não autenticado
- `403 Forbidden` - Sem permissão
- `404 Not Found` - Recurso não existe
- `500 Internal Server Error` - Erro do servidor

## Authentication

```go
// Middleware Auth
router.Post("/api/v1/score-items",
    middleware.Auth(cfg),
    middleware.RequireAdmin(),
    handler.CreateScoreItem)
```

## Swagger UI

Acessar: `http://localhost:3001/swagger/index.html`

Ver também: `database-ops.md` para operações manuais.
