package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// ============================================================
// Service
// ============================================================

// AnonymousScoreService gerencia o fluxo público do Escore Plenya Light:
// (1) exporta a configuração dos itens marcados como IsLightVersion para o site,
// (2) recebe respostas anônimas, persiste sessão + snapshot e
// (3) permite o claim posterior via magic link.
type AnonymousScoreService struct {
	db           *gorm.DB
	scoreRepo    *repository.ScoreRepository
	cfg          *config.Config
	emailService *EmailService
}

// NewAnonymousScoreService cria uma nova instância do service
func NewAnonymousScoreService(db *gorm.DB, scoreRepo *repository.ScoreRepository, cfg *config.Config, emailService *EmailService) *AnonymousScoreService {
	return &AnonymousScoreService{
		db:           db,
		scoreRepo:    scoreRepo,
		cfg:          cfg,
		emailService: emailService,
	}
}

// ============================================================
// Claim / Magic Link
// ============================================================

// magicLinkClaims são os claims do JWT usado no magic link.
type magicLinkClaims struct {
	SessionCode string `json:"sessionCode"`
	Email       string `json:"email"`
	jwt.RegisteredClaims
}

// RequestClaim gera um magic link JWT (15min), envia por email
// e retorna true se a operação foi iniciada. Não revela se a sessão existe
// ou não — previne enumeração de códigos públicos.
func (s *AnonymousScoreService) RequestClaim(code, email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" {
		return errors.New("email is required")
	}

	session, err := s.loadSessionByPublicCode(code)
	if err != nil {
		// Propositalmente retorna nil pra não revelar existência
		return nil
	}
	if session.ClaimedByPatientID != nil {
		// Já claimed — não reenvia
		return nil
	}

	claims := magicLinkClaims{
		SessionCode: code,
		Email:       email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "plenya-score-light",
			Subject:   "magic-link",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return fmt.Errorf("failed to sign magic token: %w", err)
	}

	link := fmt.Sprintf("%s/pt/escore-plenya/claim/%s", strings.TrimRight(s.cfg.Site.PublicURL, "/"), signed)

	if s.emailService != nil {
		if err := s.emailService.SendMagicLink(email, link); err != nil {
			return fmt.Errorf("failed to send magic link: %w", err)
		}
	}
	return nil
}

// ConfirmClaimResult é o retorno do claim bem-sucedido — Patient/User + JWTs do EMR.
type ConfirmClaimResult struct {
	SessionCode  string    `json:"sessionCode"`
	PatientID    uuid.UUID `json:"patientId"`
	UserID       uuid.UUID `json:"userId"`
	Email        string    `json:"email"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
}

// ConfirmClaim valida o magic token, cria/busca User+Patient e vincula a sessão.
// Retorna tokens do EMR para o frontend fazer login imediato.
func (s *AnonymousScoreService) ConfirmClaim(token string, authSvc *AuthService) (*ConfirmClaimResult, error) {
	parsed, err := jwt.ParseWithClaims(token, &magicLinkClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.Secret), nil
	})
	if err != nil || !parsed.Valid {
		return nil, errors.New("invalid or expired magic link")
	}
	claims, ok := parsed.Claims.(*magicLinkClaims)
	if !ok || claims.Subject != "magic-link" {
		return nil, errors.New("invalid magic link claims")
	}

	session, err := s.loadSessionByPublicCode(claims.SessionCode)
	if err != nil {
		return nil, fmt.Errorf("session not found or expired: %w", err)
	}

	email := strings.ToLower(strings.TrimSpace(claims.Email))

	// Busca ou cria User+Patient
	var user models.User
	err = s.db.Where("email = ?", email).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to lookup user: %w", err)
	}

	var patient *models.Patient
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Cria User + Patient novos dentro de transação
		err = s.db.Transaction(func(tx *gorm.DB) error {
			newUser := models.User{
				Name:  emailToDisplayName(email),
				Email: email,
			}
			if err := newUser.SetRoles([]string{string(models.RolePatient)}); err != nil {
				return err
			}
			if err := tx.Create(&newUser).Error; err != nil {
				return err
			}

			newPatient := models.Patient{
				UserID: newUser.ID,
				Name:   newUser.Name,
				Age:    session.Age,
				Gender: models.Gender(session.Gender),
			}
			if session.PostMenopause != nil {
				newPatient.Menopause = session.PostMenopause
			}
			if session.Height != nil {
				newPatient.Height = session.Height
			}
			if session.Weight != nil {
				newPatient.Weight = session.Weight
			}
			if err := tx.Create(&newPatient).Error; err != nil {
				return err
			}

			// Usa patient como paciente selecionado padrão
			newUser.SelectedPatientID = &newPatient.ID
			if err := tx.Save(&newUser).Error; err != nil {
				return err
			}

			user = newUser
			patient = &newPatient
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create user/patient: %w", err)
		}
	} else {
		// User existe — busca Patient vinculado
		var p models.Patient
		if err := s.db.Where("user_id = ?", user.ID).First(&p).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
			// User existe sem Patient (profissional) — cria Patient vinculado pra permitir claim
			newPatient := models.Patient{
				UserID: user.ID,
				Name:   user.Name,
				Age:    session.Age,
				Gender: models.Gender(session.Gender),
			}
			if err := s.db.Create(&newPatient).Error; err != nil {
				return nil, err
			}
			patient = &newPatient
		} else {
			patient = &p
		}
	}

	// Vincula sessão
	if _, err := s.ClaimSession(claims.SessionCode, patient.ID, email); err != nil {
		return nil, fmt.Errorf("failed to claim session: %w", err)
	}

	// Gera tokens do EMR pro frontend fazer login imediato
	authResp, err := authSvc.GenerateTokensForUser(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return &ConfirmClaimResult{
		SessionCode:  claims.SessionCode,
		PatientID:    patient.ID,
		UserID:       user.ID,
		Email:        email,
		AccessToken:  authResp.AccessToken,
		RefreshToken: authResp.RefreshToken,
	}, nil
}

// emailToDisplayName gera um nome placeholder a partir do email (user pode editar depois).
func emailToDisplayName(email string) string {
	localPart := email
	if idx := strings.Index(email, "@"); idx > 0 {
		localPart = email[:idx]
	}
	// Fallback mínimo — User.Name valida min=3
	if len(localPart) < 3 {
		return "Paciente Plenya"
	}
	return localPart
}

// ============================================================
// DTOs de resposta pública (sem ClinicalRelevance/Conduct, sem campos internos)
// ============================================================

// PublicAnonymousScoreItem é a versão pública de uma resposta — sem expor
// dados técnicos do ScoreItem subjacente (clinical_relevance, conduct).
type PublicAnonymousScoreItem struct {
	ScoreItemID   uuid.UUID `json:"scoreItemId"`
	Question      string    `json:"question"`
	NumericValue  *float64  `json:"numericValue,omitempty"`
	SelectedLevel *int      `json:"selectedLevel,omitempty"`
	TextValue     *string   `json:"textValue,omitempty"`
}

// PublicGroupResult representa o radar por grupo no payload público.
type PublicGroupResult struct {
	GroupID             uuid.UUID `json:"groupId"`
	GroupName           string    `json:"groupName"`
	ActualPoints        float64   `json:"actualPoints"`
	PossiblePoints      float64   `json:"possiblePoints"`
	ScorePercentage     float64   `json:"scorePercentage"`
	ItemsEvaluatedCount int       `json:"itemsEvaluatedCount"`
}

// PublicSnapshot é o radar resultante (sem snapshotId interno, sem timestamps).
type PublicSnapshot struct {
	TotalActualPoints      float64             `json:"totalActualPoints"`
	TotalPossiblePoints    float64             `json:"totalPossiblePoints"`
	TotalScorePercentage   float64             `json:"totalScorePercentage"`
	ItemsEvaluatedCount    int                 `json:"itemsEvaluatedCount"`
	ItemsNotEvaluatedCount int                 `json:"itemsNotEvaluatedCount"`
	GroupResults           []PublicGroupResult `json:"groupResults"`
}

// PublicSession é o payload público devolvido pelos endpoints /score-light/sessions.
// Não inclui email/claim (dados pessoais) nem ClinicalRelevance/Conduct (técnicos).
type PublicSession struct {
	PublicCode    string                     `json:"publicCode"`
	Age           int                        `json:"age"`
	Gender        string                     `json:"gender"`
	PostMenopause *bool                      `json:"postMenopause,omitempty"`
	Height        *float64                   `json:"height,omitempty"`
	Weight        *float64                   `json:"weight,omitempty"`
	CreatedAt     time.Time                  `json:"createdAt"`
	ExpiresAt     *time.Time                 `json:"expiresAt,omitempty"`
	IsClaimed     bool                       `json:"isClaimed"`
	Items         []PublicAnonymousScoreItem `json:"items"`
	Snapshot      *PublicSnapshot            `json:"snapshot,omitempty"`
}

// toPublicSession converte um modelo persistido em payload público seguro.
func toPublicSession(s *models.AnonymousScoreSession) *PublicSession {
	out := &PublicSession{
		PublicCode:    s.PublicCode,
		Age:           s.Age,
		Gender:        s.Gender,
		PostMenopause: s.PostMenopause,
		Height:        s.Height,
		Weight:        s.Weight,
		CreatedAt:     s.CreatedAt,
		ExpiresAt:     s.ExpiresAt,
		IsClaimed:     s.ClaimedByPatientID != nil,
		Items:         make([]PublicAnonymousScoreItem, 0, len(s.Items)),
	}
	for _, it := range s.Items {
		question := ""
		if it.ScoreItem != nil {
			if it.ScoreItem.LightQuestion != nil && *it.ScoreItem.LightQuestion != "" {
				question = *it.ScoreItem.LightQuestion
			} else {
				question = it.ScoreItem.Name
			}
		}
		out.Items = append(out.Items, PublicAnonymousScoreItem{
			ScoreItemID:   it.ScoreItemID,
			Question:      question,
			NumericValue:  it.NumericValue,
			SelectedLevel: it.SelectedLevel,
			TextValue:     it.TextValue,
		})
	}
	if s.Snapshot != nil {
		ps := &PublicSnapshot{
			TotalActualPoints:      s.Snapshot.TotalActualPoints,
			TotalPossiblePoints:    s.Snapshot.TotalPossiblePoints,
			TotalScorePercentage:   s.Snapshot.TotalScorePercentage,
			ItemsEvaluatedCount:    s.Snapshot.ItemsEvaluatedCount,
			ItemsNotEvaluatedCount: s.Snapshot.ItemsNotEvaluatedCount,
			GroupResults:           make([]PublicGroupResult, 0, len(s.Snapshot.GroupResults)),
		}
		for _, gr := range s.Snapshot.GroupResults {
			pr := PublicGroupResult{
				GroupID:             gr.GroupID,
				ActualPoints:        gr.ActualPoints,
				PossiblePoints:      gr.PossiblePoints,
				ScorePercentage:     gr.ScorePercentage,
				ItemsEvaluatedCount: gr.ItemsEvaluatedCount,
			}
			if gr.Group != nil {
				pr.GroupName = gr.Group.Name
			}
			ps.GroupResults = append(ps.GroupResults, pr)
		}
		out.Snapshot = ps
	}
	return out
}

// ============================================================
// DTOs / Tipos públicos (também serializados como JSON pro site)
// ============================================================

// LightLevelConfig é a representação enxuta de um ScoreLevel para o frontend público.
// Não inclui ClinicalRelevance/Conduct (somente PatientExplanation).
type LightLevelConfig struct {
	ID                 uuid.UUID `json:"id"`
	Level              int       `json:"level"`
	Name               string    `json:"name"`
	LowerLimit         *string   `json:"lowerLimit,omitempty"`
	UpperLimit         *string   `json:"upperLimit,omitempty"`
	Operator           string    `json:"operator"`
	PatientExplanation *string   `json:"patientExplanation,omitempty"`
}

// LightItemConfig representa um ScoreItem do Light já filtrado e empacotado.
type LightItemConfig struct {
	ID                 uuid.UUID          `json:"id"`
	Name               string             `json:"name"`
	LightQuestion      string             `json:"lightQuestion"` // resolvido via fallback chain
	Unit               *string            `json:"unit,omitempty"`
	Gender             *string            `json:"gender,omitempty"`
	AgeRangeMin        *int               `json:"ageRangeMin,omitempty"`
	AgeRangeMax        *int               `json:"ageRangeMax,omitempty"`
	PostMenopause      *bool              `json:"postMenopause,omitempty"`
	Points             *float64           `json:"points,omitempty"`
	LightOrder         int                `json:"lightOrder"`
	LabTestCode        *string            `json:"labTestCode,omitempty"`
	AnamneseItemCode   *string            `json:"anamneseItemCode,omitempty"`
	PatientExplanation *string            `json:"patientExplanation,omitempty"`
	Levels             []LightLevelConfig `json:"levels"`
}

// LightSubgroupConfig agrupa items por subgrupo
type LightSubgroupConfig struct {
	ID    uuid.UUID         `json:"id"`
	Name  string            `json:"name"`
	Order int               `json:"order"`
	Items []LightItemConfig `json:"items"`
}

// LightGroupConfig é o nível mais externo da árvore de configuração.
type LightGroupConfig struct {
	ID        uuid.UUID             `json:"id"`
	Name      string                `json:"name"`
	Order     int                   `json:"order"`
	Subgroups []LightSubgroupConfig `json:"subgroups"`
}

// LightConfig é o payload completo exportado para o site.
type LightConfig struct {
	Version     string             `json:"version"`     // hash baseado em contagem + última atualização
	GeneratedAt time.Time          `json:"generatedAt"`
	ItemCount   int                `json:"itemCount"`
	Groups      []LightGroupConfig `json:"groups"`
}

// SessionResponseDTO representa uma resposta individual no payload de criação
type SessionResponseDTO struct {
	ScoreItemID   uuid.UUID `json:"scoreItemId" validate:"required"`
	NumericValue  *float64  `json:"numericValue,omitempty"`
	SelectedLevel *int      `json:"selectedLevel,omitempty" validate:"omitempty,gte=0,lte=6"`
	TextValue     *string   `json:"textValue,omitempty"`
}

// CreateSessionRequest é o payload do POST /sessions
type CreateSessionRequest struct {
	Age           int                  `json:"age" validate:"required,gte=0,lte=150"`
	Gender        string               `json:"gender" validate:"required,oneof=male female other"`
	PostMenopause *bool                `json:"postMenopause,omitempty"`
	Height        *float64             `json:"height,omitempty"`
	Weight        *float64             `json:"weight,omitempty"`
	Responses     []SessionResponseDTO `json:"responses" validate:"required,min=1,dive"`
}

// ============================================================
// BuildLightConfig — exporta config do Light para o site
// ============================================================

// BuildLightConfig retorna toda a árvore de items marcados como IsLightVersion,
// preservando agrupamento Group → Subgroup → Item → Level.
// Esse é o JSON consumido pelo script de sync do site.
func (s *AnonymousScoreService) BuildLightConfig() (*LightConfig, error) {
	groups, err := s.scoreRepo.GetAllScoreGroupTrees()
	if err != nil {
		return nil, fmt.Errorf("failed to load score tree: %w", err)
	}

	out := &LightConfig{
		GeneratedAt: time.Now(),
		Groups:      []LightGroupConfig{},
	}
	itemCount := 0

	for _, g := range groups {
		gOut := LightGroupConfig{
			ID:        g.ID,
			Name:      g.Name,
			Order:     g.Order,
			Subgroups: []LightSubgroupConfig{},
		}

		for _, sg := range g.Subgroups {
			sgOut := LightSubgroupConfig{
				ID:    sg.ID,
				Name:  sg.Name,
				Order: sg.Order,
				Items: []LightItemConfig{},
			}

			seen := make(map[uuid.UUID]bool)
			for _, it := range sg.Items {
				if !it.IsLightVersion || seen[it.ID] {
					continue
				}
				sgOut.Items = append(sgOut.Items, mapItemToLightConfig(it))
				seen[it.ID] = true
				itemCount++

				// Child items também entram se marcados (e ainda não vistos)
				for _, child := range it.ChildItems {
					if !child.IsLightVersion || seen[child.ID] {
						continue
					}
					sgOut.Items = append(sgOut.Items, mapItemToLightConfig(child))
					seen[child.ID] = true
					itemCount++
				}
			}

			if len(sgOut.Items) > 0 {
				gOut.Subgroups = append(gOut.Subgroups, sgOut)
			}
		}

		if len(gOut.Subgroups) > 0 {
			out.Groups = append(out.Groups, gOut)
		}
	}

	out.ItemCount = itemCount
	out.Version = fmt.Sprintf("v%d-%d", itemCount, out.GeneratedAt.Unix())
	return out, nil
}

func mapItemToLightConfig(it models.ScoreItem) LightItemConfig {
	question := it.Name
	if it.LightQuestion != nil && *it.LightQuestion != "" {
		question = *it.LightQuestion
	}
	lightOrder := 0
	if it.LightOrder != nil {
		lightOrder = *it.LightOrder
	}

	levels := make([]LightLevelConfig, 0, len(it.Levels))
	for _, lv := range it.Levels {
		levels = append(levels, LightLevelConfig{
			ID:                 lv.ID,
			Level:              lv.Level,
			Name:               lv.Name,
			LowerLimit:         lv.LowerLimit,
			UpperLimit:         lv.UpperLimit,
			Operator:           lv.Operator,
			PatientExplanation: lv.PatientExplanation,
		})
	}

	return LightItemConfig{
		ID:                 it.ID,
		Name:               it.Name,
		LightQuestion:      question,
		Unit:               it.Unit,
		Gender:             it.Gender,
		AgeRangeMin:        it.AgeRangeMin,
		AgeRangeMax:        it.AgeRangeMax,
		PostMenopause:      it.PostMenopause,
		Points:             it.Points,
		LightOrder:         lightOrder,
		LabTestCode:        it.LabTestCode,
		AnamneseItemCode:   it.AnamneseItemCode,
		PatientExplanation: it.PatientExplanation,
		Levels:             levels,
	}
}

// ============================================================
// CreateSession — recebe respostas, persiste sessão + snapshot
// ============================================================

// CreateSession persiste uma nova sessão anônima, calcula o snapshot e retorna
// o payload público (sem campos técnicos). Operação atômica.
func (s *AnonymousScoreService) CreateSession(req CreateSessionRequest) (*PublicSession, error) {
	if len(req.Responses) == 0 {
		return nil, errors.New("at least one response is required")
	}

	session := &models.AnonymousScoreSession{
		Age:           req.Age,
		Gender:        req.Gender,
		PostMenopause: req.PostMenopause,
		Height:        req.Height,
		Weight:        req.Weight,
	}

	// Indexa respostas por ScoreItemID
	respByItem := make(map[uuid.UUID]SessionResponseDTO, len(req.Responses))
	for _, r := range req.Responses {
		respByItem[r.ScoreItemID] = r
	}

	// Carrega árvore de scores (para conseguir avaliar cada item Light)
	groups, err := s.scoreRepo.GetAllScoreGroupTrees()
	if err != nil {
		return nil, fmt.Errorf("failed to load score tree: %w", err)
	}

	pseudoPatient := session.ToPatientForApplies()

	err = s.db.Transaction(func(tx *gorm.DB) error {
		// 1. Cria sessão (gera PublicCode + ExpiresAt no BeforeCreate)
		if err := tx.Create(session).Error; err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}

		// 2. Persiste cada AnonymousScoreItem (resposta crua)
		for _, r := range req.Responses {
			item := models.AnonymousScoreItem{
				SessionID:     session.ID,
				ScoreItemID:   r.ScoreItemID,
				NumericValue:  r.NumericValue,
				TextValue:     r.TextValue,
				SelectedLevel: r.SelectedLevel,
			}
			if err := tx.Create(&item).Error; err != nil {
				return fmt.Errorf("failed to create response item: %w", err)
			}
		}

		// 3. Calcula snapshot (radar)
		snapshot := &models.AnonymousScoreSnapshot{
			SessionID: session.ID,
		}
		groupResultsMap := make(map[uuid.UUID]*models.AnonymousScoreGroupResult)

		for _, g := range groups {
			gr := &models.AnonymousScoreGroupResult{
				GroupID: g.ID,
			}
			groupResultsMap[g.ID] = gr

			for _, sg := range g.Subgroups {
				for _, it := range sg.Items {
					if it.IsLightVersion {
						evaluateLightItem(it, pseudoPatient, respByItem, snapshot, gr)
					}
					for _, child := range it.ChildItems {
						if child.IsLightVersion {
							evaluateLightItem(child, pseudoPatient, respByItem, snapshot, gr)
						}
					}
				}
			}

			if gr.PossiblePoints > 0 {
				gr.ScorePercentage = (gr.ActualPoints / gr.PossiblePoints) * 100
			}
		}

		if snapshot.TotalPossiblePoints > 0 {
			snapshot.TotalScorePercentage = (snapshot.TotalActualPoints / snapshot.TotalPossiblePoints) * 100
		}

		// 4. Persiste snapshot + group results
		if err := tx.Create(snapshot).Error; err != nil {
			return fmt.Errorf("failed to create snapshot: %w", err)
		}
		for _, gr := range groupResultsMap {
			if gr.PossiblePoints == 0 && gr.ActualPoints == 0 && gr.ItemsEvaluatedCount == 0 {
				continue // pula grupos sem nada avaliado (não polui radar)
			}
			gr.SnapshotID = snapshot.ID
			if err := tx.Create(gr).Error; err != nil {
				return fmt.Errorf("failed to create group result: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Recarrega com relations para retorno (payload público)
	full, err := s.loadSessionByPublicCode(session.PublicCode)
	if err != nil {
		return nil, err
	}
	return toPublicSession(full), nil
}

// evaluateLightItem avalia um único item Light e atualiza os agregados (snapshot + group result).
// Pré-condição: caller já confirmou item.IsLightVersion.
func evaluateLightItem(
	item models.ScoreItem,
	pseudoPatient *models.Patient,
	responses map[uuid.UUID]SessionResponseDTO,
	snapshot *models.AnonymousScoreSnapshot,
	groupResult *models.AnonymousScoreGroupResult,
) {
	if !item.AppliesToPatient(pseudoPatient) {
		snapshot.ItemsNotEvaluatedCount++
		return
	}

	resp, ok := responses[item.ID]
	if !ok {
		snapshot.ItemsNotEvaluatedCount++
		return
	}

	matched := matchLevelForResponse(item, resp)
	if matched == nil {
		snapshot.ItemsNotEvaluatedCount++
		return
	}

	maxPoints := 0.0
	if item.Points != nil {
		maxPoints = *item.Points
	}
	multiplier := GetLevelMultiplier(matched.Level)
	actual := maxPoints * multiplier

	snapshot.TotalActualPoints += actual
	snapshot.TotalPossiblePoints += maxPoints
	snapshot.ItemsEvaluatedCount++

	groupResult.ActualPoints += actual
	groupResult.PossiblePoints += maxPoints
	groupResult.ItemsEvaluatedCount++
}

// matchLevelForResponse encontra o ScoreLevel correspondente à resposta dada.
// Suporta dois modos:
//   1. SelectedLevel preenchido — encontra level por número
//   2. NumericValue preenchido — usa EvaluatesTrue (ordenado por level ASC)
func matchLevelForResponse(item models.ScoreItem, resp SessionResponseDTO) *models.ScoreLevel {
	if resp.SelectedLevel != nil {
		for i := range item.Levels {
			if item.Levels[i].Level == *resp.SelectedLevel {
				return &item.Levels[i]
			}
		}
		return nil
	}

	if resp.NumericValue != nil {
		levels := make([]models.ScoreLevel, len(item.Levels))
		copy(levels, item.Levels)
		SortLevelsAsc(levels) // mais crítico primeiro — primeiro match vence
		for i := range levels {
			if levels[i].EvaluatesTrue(*resp.NumericValue) {
				return &levels[i]
			}
		}
	}

	return nil
}

// ============================================================
// GetSessionByPublicCode — recupera sessão pra render do radar
// ============================================================

// GetSessionByPublicCode retorna o payload público da sessão (sem campos técnicos).
// Retorna ErrRecordNotFound se sessão expirou ou não existe.
func (s *AnonymousScoreService) GetSessionByPublicCode(code string) (*PublicSession, error) {
	session, err := s.loadSessionByPublicCode(code)
	if err != nil {
		return nil, err
	}
	return toPublicSession(session), nil
}

// loadSessionByPublicCode é o lookup interno — retorna o modelo completo (uso server-side).
func (s *AnonymousScoreService) loadSessionByPublicCode(code string) (*models.AnonymousScoreSession, error) {
	var session models.AnonymousScoreSession
	err := s.db.
		Preload("Snapshot.GroupResults.Group").
		Preload("Items.ScoreItem").
		Where("public_code = ?", code).
		First(&session).Error
	if err != nil {
		return nil, err
	}
	if session.ClaimedByPatientID == nil && session.ExpiresAt != nil && session.ExpiresAt.Before(time.Now()) {
		return nil, gorm.ErrRecordNotFound
	}
	return &session, nil
}

// ============================================================
// ClaimSession — vincula sessão a um Patient (após magic link)
// ============================================================

// ClaimSession vincula a sessão anônima a um PatientID já validado pelo magic link.
// Limpa ExpiresAt (sessões claimed não expiram). Retorna o modelo cru (uso interno EMR).
func (s *AnonymousScoreService) ClaimSession(code string, patientID uuid.UUID, email string) (*models.AnonymousScoreSession, error) {
	var session models.AnonymousScoreSession
	if err := s.db.Where("public_code = ?", code).First(&session).Error; err != nil {
		return nil, err
	}

	if session.ClaimedByPatientID != nil {
		return &session, nil // idempotente
	}

	now := time.Now()
	session.ClaimedByPatientID = &patientID
	session.ClaimedAt = &now
	session.Email = &email
	session.ExpiresAt = nil

	if err := s.db.Save(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

// GetSessionsByPatientID retorna todas as sessões claimed por um paciente,
// ordenadas por CreatedAt DESC. Uso interno EMR (área do paciente).
func (s *AnonymousScoreService) GetSessionsByPatientID(patientID uuid.UUID) ([]models.AnonymousScoreSession, error) {
	var sessions []models.AnonymousScoreSession
	err := s.db.
		Preload("Snapshot.GroupResults.Group").
		Where("claimed_by_patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&sessions).Error
	return sessions, err
}

// GetSessionsByCurrentUser busca o Patient vinculado ao userID e devolve as sessões
// claimed dele em payload público (PublicSession[]).
func (s *AnonymousScoreService) GetSessionsByCurrentUser(userID uuid.UUID) ([]PublicSession, error) {
	var patient models.Patient
	if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []PublicSession{}, nil
		}
		return nil, err
	}
	rawSessions, err := s.GetSessionsByPatientID(patient.ID)
	if err != nil {
		return nil, err
	}
	out := make([]PublicSession, 0, len(rawSessions))
	for i := range rawSessions {
		out = append(out, *toPublicSession(&rawSessions[i]))
	}
	return out, nil
}
