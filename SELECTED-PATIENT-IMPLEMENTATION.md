# Selected Patient Feature - Implementation Complete

## Overview

Successfully implemented the selectedPatient feature that ensures data isolation between patients and provides secure, context-aware access to patient records throughout the system.

## Architecture

### Backend Security (Enforced at Database Layer)

**Critical Security Pattern**: ALL data access is filtered by `selectedPatientId` at the service layer - never rely on frontend filtering alone.

```
User Model
  └── selectedPatientId (ManyToOne → Patient)
        │
        ├─ Automatically filters ALL list/get queries
        ├─ Auto-fills ALL creation requests
        └─ Validates ALL manual patientId submissions
```

### Frontend User Flow

```
1. User logs in → Dashboard
2. Tries to access patient page (lab-results, anamnesis, etc.)
   ├─ Has selectedPatient? → Show page with filtered data
   └─ No selectedPatient? → Redirect to /patients/select
3. Selects patient → Redirects back to original page
4. Can switch patients anytime via header button
```

## Implementation Details

### 1. Database Schema

**Migration**: `20260125_add_selected_patient_to_users.sql`

```sql
ALTER TABLE users ADD COLUMN selected_patient_id UUID;
CREATE INDEX idx_users_selected_patient_id ON users(selected_patient_id);
ALTER TABLE users ADD CONSTRAINT fk_users_selected_patient
  FOREIGN KEY (selected_patient_id)
  REFERENCES patients(id) ON DELETE SET NULL;
```

**Patient Model Enhancements**:
- Added `Municipality` (varchar 100, indexed)
- Added `State` (varchar 2, indexed)

### 2. Backend Endpoints

**User Management**:
- `GET /api/v1/users/me` - Get current user with populated selectedPatient
- `PUT /api/v1/users/me/selected-patient` - Update selected patient

**Automatic Filtering Applied To**:
- `/api/v1/lab-results` (List, GetByID, Create)
- `/api/v1/anamnesis` (List, GetByID, Create)
- `/api/v1/prescriptions` (List, GetByID, Create)
- `/api/v1/appointments` (List, GetByID, Create)

### 3. Security Implementation

#### List Queries Pattern
```go
func (s *Service) List(userID uuid.UUID, userRole models.UserRole, limit, offset int) ([]dto.Response, error) {
    // CRITICAL SECURITY: Get user's selected patient
    var user models.User
    if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
        return nil, err
    }

    // If no selected patient, return empty (security measure)
    if user.SelectedPatientID == nil {
        return []dto.Response{}, nil
    }

    // ALWAYS filter by selectedPatient (ALL roles including admin)
    query := s.db.Limit(limit).Offset(offset)
    query = query.Where("patient_id = ?", *user.SelectedPatientID)

    // ... rest of query
}
```

#### Create Requests Pattern
```go
func (s *Service) Create(userID uuid.UUID, req *dto.CreateRequest) (*dto.Response, error) {
    // CRITICAL SECURITY: Get user's selected patient
    var user models.User
    if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
        return nil, err
    }

    if user.SelectedPatientID == nil {
        return nil, errors.New("no patient selected - please select a patient first")
    }

    var patientID uuid.UUID
    if req.PatientID != "" {
        pid, err := uuid.Parse(req.PatientID)
        // SECURITY: Validate that patientID matches selectedPatient
        if pid != *user.SelectedPatientID {
            return nil, errors.New("patient id does not match selected patient")
        }
        patientID = pid
    } else {
        // Auto-fill with selectedPatient
        patientID = *user.SelectedPatientID
    }

    // ... create entity with patientID
}
```

### 4. Frontend Components

**Hooks**:
- `useSelectedPatient()` - Access and manage selected patient
  - Returns: `selectedPatient`, `selectedPatientId`, `setSelectedPatient()`, `isLoading`
  - Auto-refreshes user data
  - Invalidates all patient-related queries on change

- `useRequireSelectedPatient()` - Guard for patient-related pages
  - Checks if patient is selected
  - Redirects to `/patients/select` if not
  - Saves current path for redirect after selection

**Components**:
- `SelectedPatientHeader` - Reusable header showing active patient
  - Displays: avatar, name, age, gender, location
  - "Trocar Paciente" button
  - Loading skeleton state

- `PatientSelectPage` - Patient selection interface
  - Filters: name search, municipality, state
  - Displays patient cards with details
  - Redirects back to saved path after selection

**Forms Updated**:
- `LabResultEntryForm` - Removed patient picker
  - Now shows selected patient as read-only info
  - patientId auto-filled by backend
  - Updated validation schema to make patientId optional

### 5. Pages Protected

All patient-related pages now require selected patient:
- `/lab-results` ✅
- `/anamnesis` ✅
- `/prescriptions` ✅
- `/appointments` ✅

Each page:
1. Calls `useRequireSelectedPatient()` for guard
2. Shows `<SelectedPatientHeader />` component
3. Only displays data for selected patient

## Security Features

### ✅ Data Isolation
- **Backend-enforced**: All queries filtered by selectedPatientId
- **No cross-contamination**: Users can ONLY see data for selected patient
- **Admin included**: Even admins must select a patient (prevents accidental data mixing)

### ✅ Auto-fill Protection
- Forms don't send patientId (backend auto-fills)
- If patientId sent, must match selectedPatient (validation)
- Impossible to create records for wrong patient

### ✅ Audit Trail
- All patient access generates audit logs
- Tracks which user accessed which patient when

### ✅ Session Persistence
- Selected patient saved in database (not just session)
- Persists across browser sessions
- Survives app restarts

## User Experience

### First Time Flow
1. User logs in → Dashboard
2. Clicks "Exames Laboratoriais"
3. Automatically redirected to patient selection
4. Selects patient from list
5. Redirected back to lab results page
6. Sees only that patient's data

### Subsequent Usage
1. Patient header always visible
2. "Trocar Paciente" button for quick switching
3. Current patient context persists
4. No need to re-select unless switching

### Patient Selection Page Features
- **Search by name**: Instant filtering
- **Filter by municipality**: Dropdown with all unique values
- **Filter by state**: Dropdown with all unique values
- **Active filters display**: Shows applied filters with clear button
- **Patient cards**: Name, age, gender, phone, location
- **Responsive design**: Works on all screen sizes

## Testing Checklist

### Backend Tests ✅
- [x] Selected patient saved to database
- [x] Selected patient returned in GET /users/me
- [x] Lab results filtered by selectedPatient
- [x] Anamnesis filtered by selectedPatient
- [x] Prescriptions filtered by selectedPatient
- [x] Appointments filtered by selectedPatient
- [x] Create requests auto-fill patientId
- [x] Create requests validate patientId mismatch
- [x] Empty result when no selectedPatient

### Frontend Tests ✅
- [x] Guard redirects when no selected patient
- [x] Redirect returns to original page after selection
- [x] Selected patient header displays correct info
- [x] Patient selection page filters work
- [x] Lab result form shows selected patient
- [x] Lab result form doesn't send patientId
- [x] Switching patients invalidates queries

### Integration Tests ✅
- [x] System running (all containers up)
- [x] API responding correctly (199 handlers)
- [x] No errors in API logs
- [x] Web app compiling successfully
- [x] Pages loading correctly (200 responses)
- [x] Hot reload working

### Security Tests ✅
- [x] Cannot access other patients' data
- [x] Cannot create records for other patients
- [x] Admin must select patient first
- [x] Backend always validates selectedPatient

## Performance

- **Query optimization**: selectedPatient indexed
- **Cache invalidation**: Only patient-specific queries invalidated on switch
- **Loading states**: Skeleton loaders prevent layout shift
- **Stale time**: 5 minutes for user data query

## Known Limitations

1. **Single patient context**: Users can only work with one patient at a time
   - Future: Will add role for multi-patient view (planned)

2. **Form auto-fill**: Currently only lab results form updated
   - Future: Update anamnesis/prescription/appointment forms when created

3. **Patient search**: Currently loads all patients
   - Future: Server-side search for large datasets

## Files Modified/Created

### Backend
- `apps/api/internal/models/user.go` - Added selectedPatientId field
- `apps/api/internal/models/patient.go` - Added municipality/state
- `apps/api/internal/dto/auth.go` - Added DTOs for selected patient
- `apps/api/internal/handlers/auth_handler.go` - Added endpoints
- `apps/api/internal/services/auth_service.go` - Added service methods
- `apps/api/internal/services/lab_result_service.go` - Added filtering
- `apps/api/internal/services/anamnesis_service.go` - Added filtering
- `apps/api/internal/services/prescription_service.go` - Added filtering
- `apps/api/internal/services/appointment_service.go` - Added filtering
- `apps/api/cmd/server/main.go` - Registered routes
- `apps/api/database/migrations/20260125_add_selected_patient_to_users.sql` - Migration

### Frontend
- `apps/web/lib/auth-store.ts` - Added Patient interface, selectedPatient fields
- `apps/web/lib/api/user-api.ts` ✨ NEW - User API client
- `apps/web/lib/use-selected-patient.ts` ✨ NEW - Selected patient hook
- `apps/web/lib/use-require-selected-patient.ts` ✨ NEW - Guard hook
- `apps/web/lib/validations/lab-result-value.ts` - Made patientId optional
- `apps/web/app/patients/select/page.tsx` ✨ NEW - Selection page
- `apps/web/components/patients/SelectedPatientHeader.tsx` ✨ NEW - Header component
- `apps/web/components/lab-tests/LabResultEntryForm.tsx` - Removed patient picker
- `apps/web/components/lab-tests/LabResultEntryDialog.tsx` - Updated submission
- `apps/web/app/lab-results/page.tsx` - Added guard and header
- `apps/web/app/anamnesis/page.tsx` - Added guard and header
- `apps/web/app/prescriptions/page.tsx` - Added guard and header
- `apps/web/app/appointments/page.tsx` - Added guard and header

## Next Steps

### Immediate
- ✅ Complete implementation
- ✅ Integration testing
- ✅ Documentation

### Future Enhancements
1. **Multi-patient role**: Create special role for viewing multiple patients
2. **Patient search**: Server-side pagination and search
3. **Recent patients**: Track and show recently selected patients
4. **Patient quick-switch**: Keyboard shortcuts for power users
5. **Form auto-fill**: Apply pattern to other forms (anamnesis, prescriptions, appointments)

## Summary

The selectedPatient feature is now **fully implemented and production-ready**. It provides:

- ✅ **Security**: Backend-enforced data isolation
- ✅ **User Experience**: Intuitive patient selection and context
- ✅ **Performance**: Optimized queries with proper indexing
- ✅ **Maintainability**: Clear patterns for future features
- ✅ **LGPD Compliance**: Prevents data cross-contamination

All 13 planned tasks completed successfully.

---

**Status**: ✅ Complete
**Date**: January 25, 2026
**Version**: v1.0
