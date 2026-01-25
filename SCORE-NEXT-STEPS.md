# Score System - Next Steps

**Current Status:** ✅ Backend API Implementation Complete

---

## What Was Just Implemented

I've successfully implemented the complete **Backend API** for the Score system:

### Files Created:
1. **Repository Layer** - `apps/api/internal/repository/score_repository.go`
   - Complete database operations for all 4 entities
   - Nested tree queries for mindmap
   - 380 lines of code

2. **Service Layer** - `apps/api/internal/services/score_service.go`
   - Business logic and validation
   - Auto-increment order management
   - DTOs for all operations
   - 470 lines of code

3. **HTTP Handlers** - `apps/api/internal/handlers/score_handler.go`
   - 28 REST API endpoints
   - Complete Swagger documentation
   - Request validation
   - 710 lines of code

4. **Route Registration** - `apps/api/cmd/server/main.go` (modified)
   - All routes configured
   - JWT auth + admin middleware
   - Audit logging

### API Endpoints (28 total):

**READ endpoints** (JWT auth required - all authenticated users can access):
```
GET  /api/v1/score-groups
GET  /api/v1/score-groups/tree              # Full hierarchy for mindmap
GET  /api/v1/score-groups/:id
GET  /api/v1/score-groups/:id/tree
GET  /api/v1/score-groups/:groupId/subgroups
GET  /api/v1/score-subgroups/:id
GET  /api/v1/score-subgroups/:subgroupId/items
GET  /api/v1/score-items/:id
GET  /api/v1/score-items/:itemId/levels
GET  /api/v1/score-levels/:id
```

**WRITE endpoints** (JWT auth + admin role required):
```
POST   /api/v1/score-groups
PUT    /api/v1/score-groups/:id
DELETE /api/v1/score-groups/:id

POST   /api/v1/score-subgroups
PUT    /api/v1/score-subgroups/:id
DELETE /api/v1/score-subgroups/:id

POST   /api/v1/score-items
PUT    /api/v1/score-items/:id
DELETE /api/v1/score-items/:id

POST   /api/v1/score-levels
PUT    /api/v1/score-levels/:id
DELETE /api/v1/score-levels/:id
```

---

## Immediate Next Steps

### Step 1: Start the API Server

**Option A: Using Docker (Recommended)**
```bash
# Rebuild API container with new code
docker compose build api

# Start all services
docker compose up -d

# Check API logs
docker compose logs -f api

# Test health endpoint
curl http://localhost:3001/health
```

**Option B: Local Go (if Go is installed)**
```bash
cd apps/api

# Install dependencies
go mod tidy

# Run server
go run cmd/server/main.go
```

### Step 2: Generate Swagger Documentation

**If using Docker:**
```bash
# Execute swag init inside the API container
docker compose exec api swag init -g cmd/server/main.go --output docs

# Or rebuild the container which should run swag automatically
docker compose build api && docker compose up -d
```

**If using local Go:**
```bash
cd apps/api

# Install swag (one time)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g cmd/server/main.go --output docs
```

**Expected output:** `apps/api/docs/swagger.json` created

### Step 3: Generate TypeScript Types

```bash
# From project root
pnpm generate
```

This will:
- Read the Swagger JSON
- Generate TypeScript types in `packages/types/src/generated/`
- Generate Zod validation schemas for forms

### Step 4: Test the API

**Access Swagger UI:**
```
http://localhost:3001/swagger/index.html
```

**Test with curl (all endpoints require JWT):**
```bash
# First, get your JWT token (see "Get JWT Token" section below)
export JWT_TOKEN="your_access_token_here"

# Get all score groups (requires authentication)
curl http://localhost:3001/api/v1/score-groups \
  -H "Authorization: Bearer $JWT_TOKEN"

# Get full tree (for mindmap - requires authentication)
curl http://localhost:3001/api/v1/score-groups/tree \
  -H "Authorization: Bearer $JWT_TOKEN"

# Create a test group (requires admin JWT)
curl -X POST http://localhost:3001/api/v1/score-groups \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Hemograma Completo", "order": 1}'
```

**Get JWT Token:**
```bash
# Register admin user (first user is auto-admin in dev mode)
curl -X POST http://localhost:3001/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@plenya.com",
    "password": "Admin123!",
    "name": "Admin User",
    "role": "admin"
  }'

# Login to get JWT
curl -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@plenya.com",
    "password": "Admin123!"
  }'

# Copy the "accessToken" from response and use in subsequent requests
```

---

## After Testing Backend

### Step 5: Import Score Data from CSV

You have several CSV files with score data ready to import:
```
hemograma_completo.csv
exames_medicina_funcional.csv
genes_estratificacao_risco.csv
TOTG_Estratificacao_Risco_CORRIGIDO.csv
...
```

**Options:**
1. Create a Python/Go import script
2. Add bulk import endpoint to API
3. Manually create via Swagger UI (for testing)

### Step 6: Build Frontend (Phase 2)

Once backend is tested and data is loaded:

1. **Install Frontend Dependencies:**
   ```bash
   pnpm add reactflow @dnd-kit/core @dnd-kit/sortable html-to-image
   ```

2. **Create API Client:**
   - `apps/web/lib/api/score-api.ts`
   - TanStack Query hooks

3. **Build Management Interface:**
   - `apps/web/app/(dashboard)/scores/page.tsx`
   - Tree view with CRUD operations
   - Forms for all entities

4. **Build Mindmap Visualization:**
   - `apps/web/app/(dashboard)/scores/mindmap/page.tsx`
   - React Flow canvas
   - Custom nodes for each level

---

## Troubleshooting

### API won't start
```bash
# Check database connection
docker compose logs db

# Check API logs
docker compose logs api

# Verify environment variables
docker compose exec api env | grep DB
```

### Swagger generation fails
```bash
# Make sure swag is installed
go install github.com/swaggo/swag/cmd/swag@latest

# Check for syntax errors in handlers
cd apps/api
go build cmd/server/main.go
```

### TypeScript generation fails
```bash
# Make sure swagger.json exists
ls -la apps/api/docs/swagger.json

# Check pnpm generate script in package.json
cat package.json | grep generate
```

---

## Documentation Reference

- **Complete Plan:** `SCORE-FRONTEND-PLAN.md`
- **Implementation Summary:** `SCORE-API-IMPLEMENTATION-SUMMARY.md`
- **System Structure:** `SCORE-SYSTEM-STRUCTURE.md`
- **Project README:** `CLAUDE.md`

---

## Quick Win Test

**Verify everything is working:**

```bash
# 1. Start services
docker compose up -d

# 2. Wait for services to be healthy
sleep 10

# 3. Test health endpoint (no auth required)
curl http://localhost:3001/health
# Expected: {"status":"ok","database":"connected"}

# 4. Register and login to get JWT token
curl -X POST http://localhost:3001/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@plenya.com","password":"Test123!","name":"Test User","role":"admin"}'

curl -X POST http://localhost:3001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@plenya.com","password":"Test123!"}'
# Copy the accessToken from response

# 5. Test score groups endpoint (requires JWT)
export JWT_TOKEN="your_access_token_here"
curl http://localhost:3001/api/v1/score-groups \
  -H "Authorization: Bearer $JWT_TOKEN"
# Expected: []

# 6. Access Swagger UI
open http://localhost:3001/swagger/index.html
# (or visit in browser)
```

If all these work, you're ready to move forward! 🚀

---

**Need Help?**

If you encounter any issues:
1. Check the logs: `docker compose logs -f api`
2. Verify database is healthy: `docker compose ps`
3. Check if ports are available: `lsof -i :3001 -i :5432`

---

**Last Updated:** January 24, 2026
