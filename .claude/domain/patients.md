# Domain - Patients

## Patient Model

```go
type Patient struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name      string    `gorm:"type:varchar(200);not null" json:"name"`
    BirthDate time.Time `gorm:"type:date;not null" json:"birthDate"`
    Age       int       `gorm:"-" json:"age"` // Calculado
    
    // @enum male,female,other
    Gender string `gorm:"type:varchar(10);not null;check:gender IN ('male','female','other')" json:"gender"`
    
    // Dados sensíveis (criptografados)
    CPF *string `gorm:"type:text" json:"cpf,omitempty"`
    RG  *string `gorm:"type:text" json:"rg,omitempty"`
    
    // Contato
    Email *string `gorm:"type:varchar(100);unique" json:"email,omitempty"`
    Phone *string `gorm:"type:varchar(20)" json:"phone,omitempty"`
    
    // Menopausa (apenas mulheres)
    Menopause *bool `gorm:"type:boolean" json:"menopause,omitempty"`
    
    // Relações
    Anamneses     []Anamnesis    `gorm:"foreignKey:PatientID" json:"anamneses,omitempty"`
    LabResults    []LabResult    `gorm:"foreignKey:PatientID" json:"labResults,omitempty"`
    Prescriptions []Prescription `gorm:"foreignKey:PatientID" json:"prescriptions,omitempty"`
    
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

## Hooks

```go
// UUID v7 (obrigatório)
func (p *Patient) BeforeCreate(tx *gorm.DB) error {
    if p.ID == uuid.Nil {
        p.ID = uuid.Must(uuid.NewV7())
    }
    return nil
}

// Criptografar CPF/RG + calcular idade
func (p *Patient) BeforeSave(tx *gorm.DB) error {
    if p.CPF != nil && *p.CPF != "" && !isEncrypted(*p.CPF) {
        encrypted, _ := crypto.EncryptWithDefaultKey(*p.CPF)
        p.CPF = &encrypted
    }
    if p.RG != nil && *p.RG != "" && !isEncrypted(*p.RG) {
        encrypted, _ := crypto.EncryptWithDefaultKey(*p.RG)
        p.RG = &encrypted
    }
    p.CalculateAge()
    return nil
}

// Descriptografar após buscar
func (p *Patient) AfterFind(tx *gorm.DB) error {
    if p.CPF != nil && *p.CPF != "" && isEncrypted(*p.CPF) {
        decrypted, _ := crypto.DecryptWithDefaultKey(*p.CPF)
        p.CPF = &decrypted
    }
    if p.RG != nil && *p.RG != "" && isEncrypted(*p.RG) {
        decrypted, _ := crypto.DecryptWithDefaultKey(*p.RG)
        p.RG = &decrypted
    }
    return nil
}

// Calcular idade
func (p *Patient) CalculateAge() {
    if !p.BirthDate.IsZero() {
        years := time.Since(p.BirthDate).Hours() / 24 / 365.25
        p.Age = int(years)
    }
}
```

## Workflows

### Criar Paciente

```go
func (s *PatientService) Create(dto CreatePatientDTO) (*models.Patient, error) {
    // Validar CPF único
    if dto.CPF != nil {
        exists, _ := s.repo.ExistsByCPF(*dto.CPF)
        if exists {
            return nil, errors.New("CPF already registered")
        }
    }
    
    patient := &models.Patient{
        Name:      dto.Name,
        BirthDate: dto.BirthDate,
        Gender:    dto.Gender,
        CPF:       dto.CPF,
        Email:     dto.Email,
    }
    
    if err := s.repo.Create(patient); err != nil {
        return nil, err
    }
    
    return patient, nil
}
```

### Buscar Aplicável a ScoreItems

```go
func (s *PatientService) GetApplicableScoreItems(patientID uuid.UUID) ([]models.ScoreItem, error) {
    patient, err := s.repo.GetByID(patientID)
    if err != nil {
        return nil, err
    }
    
    allItems, err := s.scoreRepo.ListItems()
    if err != nil {
        return nil, err
    }
    
    var applicable []models.ScoreItem
    for _, item := range allItems {
        if item.AppliesToPatient(patient) {
            applicable = append(applicable, item)
        }
    }
    
    return applicable, nil
}
```

Ver: `apps/api/internal/models/patient.go`
