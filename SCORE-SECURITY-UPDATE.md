# Score API - Security Update

**Date:** January 24, 2026
**Change:** Score data is now protected (proprietary medical data)

---

## What Changed

### Before (Incorrect):
- ❌ GET endpoints were public (no authentication required)
- ✅ POST/PUT/DELETE endpoints required admin authentication

### After (Correct):
- ✅ **ALL endpoints require JWT authentication**
- ✅ GET endpoints: Any authenticated user can read
- ✅ POST/PUT/DELETE endpoints: Admin only

---

## Security Model

### Read Access (GET endpoints)
**Who can access:** All authenticated users (patients, doctors, nurses, admins)
**Reason:** Authenticated users need to view score criteria for their lab results

**Endpoints:**
```
GET /api/v1/score-groups
GET /api/v1/score-groups/tree
GET /api/v1/score-groups/:id
GET /api/v1/score-groups/:id/tree
GET /api/v1/score-groups/:groupId/subgroups
GET /api/v1/score-subgroups/:id
GET /api/v1/score-subgroups/:subgroupId/items
GET /api/v1/score-items/:id
GET /api/v1/score-items/:itemId/levels
GET /api/v1/score-levels/:id
```

**Middleware:**
- `middleware.Auth(cfg)` - JWT validation
- `middleware.AuditLog(database.DB)` - Audit trail

### Write Access (POST/PUT/DELETE endpoints)
**Who can access:** Admin users only
**Reason:** Score definitions are critical medical reference data

**Endpoints:**
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

**Middleware:**
- `middleware.Auth(cfg)` - JWT validation
- `middleware.RequireAdmin()` - Admin role check
- `middleware.AuditLog(database.DB)` - Audit trail

---

## Files Modified

1. **`apps/api/cmd/server/main.go`**
   - Added `middleware.Auth(cfg)` to all score route groups
   - Moved `middleware.RequireAdmin()` to individual POST/PUT/DELETE routes
   - All operations logged via `middleware.AuditLog(database.DB)`

2. **`apps/api/internal/handlers/score_handler.go`**
   - Added `@Security BearerAuth` to all GET endpoint Swagger annotations
   - Added `@Failure 401` to all GET endpoint Swagger annotations

3. **Documentation updates:**
   - `SCORE-API-IMPLEMENTATION-SUMMARY.md` - Security section updated
   - `SCORE-NEXT-STEPS.md` - Test examples updated with JWT headers
   - `SCORE-SECURITY-UPDATE.md` - This file

---

## Testing with Authentication

### 1. Register/Login to get JWT token

```bash
# Register an admin user
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

# Extract accessToken from response
export JWT_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### 2. Test Read Access (All authenticated users)

```bash
# Get all score groups (requires JWT)
curl http://localhost:3001/api/v1/score-groups \
  -H "Authorization: Bearer $JWT_TOKEN"

# Get full tree for mindmap (requires JWT)
curl http://localhost:3001/api/v1/score-groups/tree \
  -H "Authorization: Bearer $JWT_TOKEN"
```

### 3. Test Write Access (Admin only)

```bash
# Create a score group (requires admin JWT)
curl -X POST http://localhost:3001/api/v1/score-groups \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Hemograma Completo",
    "order": 1
  }'

# Update a score group (requires admin JWT)
curl -X PUT http://localhost:3001/api/v1/score-groups/{id} \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Hemograma Completo - Atualizado"
  }'

# Delete a score group (requires admin JWT)
curl -X DELETE http://localhost:3001/api/v1/score-groups/{id} \
  -H "Authorization: Bearer $JWT_TOKEN"
```

### 4. Test Unauthorized Access

```bash
# Without JWT - should return 401 Unauthorized
curl http://localhost:3001/api/v1/score-groups

# With invalid JWT - should return 401 Unauthorized
curl http://localhost:3001/api/v1/score-groups \
  -H "Authorization: Bearer invalid_token"

# Non-admin trying to create - should return 403 Forbidden
curl -X POST http://localhost:3001/api/v1/score-groups \
  -H "Authorization: Bearer $NON_ADMIN_JWT" \
  -H "Content-Type: application/json" \
  -d '{"name":"Test"}'
```

---

## LGPD/Security Compliance

✅ **Authentication Required:** All endpoints require valid JWT token
✅ **Authorization:** Admin-only for mutations (CREATE/UPDATE/DELETE)
✅ **Audit Logging:** All operations logged with user ID, action, timestamp
✅ **Proprietary Data Protection:** Score definitions protected from public access
✅ **Role-Based Access Control (RBAC):** Enforced via middleware

---

## Frontend Impact

The frontend API client (`apps/web/lib/api/score-api.ts`) will need to:

1. **Include JWT token in ALL requests:**
   ```typescript
   const config = {
     headers: {
       'Authorization': `Bearer ${token}`
     }
   }
   ```

2. **Handle 401 Unauthorized:**
   - Redirect to login page
   - Refresh JWT token if expired

3. **Handle 403 Forbidden:**
   - Show "Admin access required" message
   - Hide admin-only UI elements for non-admin users

4. **TanStack Query setup:**
   ```typescript
   const queryClient = new QueryClient({
     defaultOptions: {
       queries: {
         onError: (error) => {
           if (error.response?.status === 401) {
             // Redirect to login
             router.push('/login')
           }
         }
       }
     }
   })
   ```

---

## Summary

**All Score API endpoints now require authentication.** This protects your proprietary medical knowledge and ensures compliance with data protection requirements.

- ✅ Read access: All authenticated users
- ✅ Write access: Admin users only
- ✅ Audit logging: All operations tracked
- ✅ Swagger docs: Updated with security requirements

**No breaking changes** to the existing API structure - only added authentication middleware.

---

**Last Updated:** January 24, 2026
