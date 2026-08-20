package services

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var ErrMagistralComponentNotFound = errors.New("magistral component not found")

type MagistralComponentService struct {
	db *gorm.DB
}

func NewMagistralComponentService(db *gorm.DB) *MagistralComponentService {
	return &MagistralComponentService{db: db}
}

// Search — busca por nome/sinônimo. Ordena por uso: o repertório do próprio prescritor sobe
// primeiro, que é o que faz a busca acertar no primeiro item depois de algumas semanas.
func (s *MagistralComponentService) Search(query string, limit int) ([]models.MagistralComponent, error) {
	q := strings.TrimSpace(query)
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	var out []models.MagistralComponent

	// Busca vazia = listagem do repertório (mais prescritos primeiro). É o que a tela de
	// curadoria precisa mostrar ao abrir, e o que o autocomplete NÃO usa (lá o mínimo é 2
	// caracteres, senão a lista aparece antes de o médico dizer o que quer).
	if q == "" {
		err := s.db.Where("is_active = true AND deleted_at IS NULL").
			Order("usage_count DESC").Order("name").Limit(limit).Find(&out).Error
		return out, err
	}
	if len(q) < 2 {
		return []models.MagistralComponent{}, nil
	}

	like := "%" + strings.ToLower(q) + "%"
	err := s.db.
		Where("is_active = true AND deleted_at IS NULL").
		Where(`lower(public.immutable_unaccent(name)) LIKE lower(public.immutable_unaccent(?))
		       OR lower(public.immutable_unaccent(synonyms)) LIKE lower(public.immutable_unaccent(?))`, like, like).
		Order(clause.Expr{
			SQL:  `CASE WHEN lower(public.immutable_unaccent(name)) LIKE lower(public.immutable_unaccent(?)) THEN 0 ELSE 1 END`,
			Vars: []interface{}{strings.ToLower(q) + "%"},
		}).
		Order("usage_count DESC").
		Order("name").
		Limit(limit).
		Find(&out).Error
	return out, err
}

func (s *MagistralComponentService) GetByID(id uuid.UUID) (*models.MagistralComponent, error) {
	var c models.MagistralComponent
	if err := s.db.First(&c, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMagistralComponentNotFound
		}
		return nil, err
	}
	return &c, nil
}

// findByName resolve pelo nome normalizado (sem acento, minúsculo) — é assim que a substância
// digitada à mão na receita reencontra a linha do catálogo.
func (s *MagistralComponentService) findByName(tx *gorm.DB, name string) (*models.MagistralComponent, error) {
	n := strings.TrimSpace(name)
	if n == "" {
		return nil, nil
	}

	// Escada de tentativas, da mais certa para a menos certa, parando na primeira que acerta.
	// Nenhum degrau é aproximado: casar por semelhança traria faixa de dose e densidade de OUTRA
	// substância para dentro de uma receita, o que é pior do que não casar nada.
	//
	// O que motivou a escada: as fórmulas do formulário escrevem "Vitamina B6
	// (piridoxal-5-fosfato)", "PQQ (pirroloquinolina quinona)", "Coenzima Q10 (ubiquinona) ou
	// ubiquinol". Todas existem no catálogo, e nenhuma casava — quatro fórmulas ficavam sem
	// cálculo de cápsula por causa da forma de escrever o nome.
	for _, termo := range nameCandidates(n) {
		var c models.MagistralComponent
		err := tx.Where(`lower(public.immutable_unaccent(name)) = lower(public.immutable_unaccent(?))
		                 OR EXISTS (
		                      SELECT 1 FROM unnest(string_to_array(coalesce(synonyms, ''), ',')) AS syn
		                       WHERE btrim(syn) <> ''
		                         AND lower(public.immutable_unaccent(btrim(syn)))
		                           = lower(public.immutable_unaccent(?))
		                 )`, termo, termo).
			First(&c).Error
		if err == nil {
			return &c, nil
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return nil, nil
}

// Qualificadores que descrevem o insumo sem mudar qual substância é. "anidra", "quelato" e afins
// ficam DE FORA de propósito: mudam o insumo, a densidade e o fator de correção.
var qualificadoresDeInsumo = []string{"padronizada", "padronizado", "micronizada", "micronizado", "puro", "pura"}

// nameCandidates devolve as formas do nome a tentar, na ordem.
func nameCandidates(n string) []string {
	out := []string{n}
	add := func(v string) {
		v = strings.TrimSpace(v)
		if v == "" {
			return
		}
		for _, j := range out {
			if strings.EqualFold(j, v) {
				return
			}
		}
		out = append(out, v)
	}

	// "Coenzima Q10 (ubiquinona) ou ubiquinol" → o que vem antes do "ou" é a substância principal.
	base := n
	if i := strings.Index(strings.ToLower(base), " ou "); i > 0 {
		base = base[:i]
		add(base)
	}

	// Tira o parêntese: "Vitamina B6 (piridoxal-5-fosfato)" → "Vitamina B6", e o conteúdo do
	// parêntese também vira candidato, porque às vezes é ELE o nome do catálogo.
	if abre := strings.Index(base, "("); abre >= 0 {
		fecha := strings.Index(base[abre:], ")")
		semParenteses := strings.TrimSpace(base[:abre])
		add(semParenteses)
		if fecha > 1 {
			add(base[abre+1 : abre+fecha])
		}
		base = semParenteses
	}

	// Por último, sem o qualificador de insumo: "Curcumina padronizada" → "Curcumina".
	lower := strings.ToLower(base)
	for _, q := range qualificadoresDeInsumo {
		if strings.HasSuffix(lower, " "+q) {
			add(base[:len(base)-len(q)-1])
			break
		}
	}
	return out
}

// Upsert cria ou atualiza um componente (curadoria manual).
func (s *MagistralComponentService) Upsert(id *uuid.UUID, req *dto.MagistralComponentRequest, userID uuid.UUID) (*models.MagistralComponent, error) {
	var c models.MagistralComponent
	if id != nil {
		existing, err := s.GetByID(*id)
		if err != nil {
			return nil, err
		}
		c = *existing
	}

	c.Name = strings.TrimSpace(req.Name)
	c.Synonyms = strings.TrimSpace(req.Synonyms)
	c.CAS = req.CAS
	c.DCBCode = req.DCBCode
	if req.DefaultUnit != "" {
		c.DefaultUnit = req.DefaultUnit
	}
	c.UsualDose = req.UsualDose
	c.MinDose = req.MinDose
	c.MaxDose = req.MaxDose
	c.BulkDensity = req.BulkDensity
	c.Brand = req.Brand
	c.ElementalPercent = req.ElementalPercent
	if req.DoseAsElemental != nil {
		c.DoseAsElemental = *req.DoseAsElemental
	}
	c.CorrectionNote = req.CorrectionNote
	// Editar a densidade à mão é informar medida, não classe: o rótulo acompanha.
	if req.DensitySource != nil {
		c.DensitySource = req.DensitySource
	} else if req.BulkDensity != nil && (c.DensitySource == nil || *c.DensitySource == "classe") {
		medida := "informada pelo médico"
		c.DensitySource = &medida
	}
	c.EutecticFormer = req.EutecticFormer
	c.Hygroscopic = req.Hygroscopic
	c.Oxidizing = req.Oxidizing
	c.OxidationSensitive = req.OxidationSensitive
	c.Photosensitive = req.Photosensitive
	c.Bitterness = req.Bitterness
	c.SachetOK = req.SachetOK
	c.Notes = req.Notes
	c.Indications = req.Indications
	c.DoseReference = req.DoseReference
	c.IndicationBullets = req.IndicationBullets
	c.DoseBullets = req.DoseBullets
	if req.Source != "" {
		c.Source = req.Source
	}
	// Editar à mão é conferir: o que veio do RAG deixa de ser sugestão quando o médico salva.
	if c.EvidenceStatus == "suggested" {
		c.EvidenceStatus = "confirmed"
	}
	c.ReviewedBy = &userID

	if id == nil {
		if err := s.db.Create(&c).Error; err != nil {
			return nil, err
		}
		return &c, nil
	}
	if err := s.db.Save(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// SaveDefaultDose é a curadoria OPORTUNISTA: o médico acabou de prescrever uma dose para uma
// substância sem faixa cadastrada e pediu para guardar como padrão. Cria a linha se não existir.
//
// Guarda só o que ele afirmou (dose usual e unidade). Densidade, amargor e flags continuam NULL —
// inferir densidade a partir de dose seria inventar.
func (s *MagistralComponentService) SaveDefaultDose(substance string, dose float64, unit string, userID uuid.UUID) (*models.MagistralComponent, error) {
	name := strings.TrimSpace(substance)
	if name == "" {
		return nil, errors.New("substância obrigatória")
	}
	if dose <= 0 {
		return nil, errors.New("dose precisa ser maior que zero")
	}

	var out *models.MagistralComponent
	err := s.db.Transaction(func(tx *gorm.DB) error {
		existing, err := s.findByName(tx, name)
		if err != nil {
			return err
		}
		if existing == nil {
			c := models.MagistralComponent{
				Name:        name,
				DefaultUnit: unit,
				UsualDose:   &dose,
				Source:      "curadoria",
				ReviewedBy:  &userID,
			}
			if err := tx.Create(&c).Error; err != nil {
				return err
			}
			out = &c
			return nil
		}
		existing.UsualDose = &dose
		existing.DefaultUnit = unit
		existing.ReviewedBy = &userID
		if existing.Source == "uso" {
			existing.Source = "curadoria"
		}
		if err := tx.Save(existing).Error; err != nil {
			return err
		}
		out = existing
		return nil
	})
	return out, err
}

// EnsureComponents é o outro lado da curadoria oportunista: toda substância prescrita que ainda
// não está no catálogo entra como esboço (só nome e unidade, sem nada clínico), e o que já está
// tem o contador de uso incrementado. Assim a busca vai ficando parecida com o repertório real
// sem ninguém abrir uma tela de cadastro.
func (s *MagistralComponentService) EnsureComponents(tx *gorm.DB, formulas []models.PrescriptionFormula) error {
	for fi := range formulas {
		for ci := range formulas[fi].Components {
			comp := &formulas[fi].Components[ci]
			if comp.MagistralComponentID != nil {
				if err := tx.Model(&models.MagistralComponent{}).
					Where("id = ?", *comp.MagistralComponentID).
					UpdateColumn("usage_count", gorm.Expr("usage_count + 1")).Error; err != nil {
					return err
				}
				continue
			}

			existing, err := s.findByName(tx, comp.Substance)
			if err != nil {
				return err
			}
			if existing != nil {
				comp.MagistralComponentID = &existing.ID
				if err := tx.Model(&models.MagistralComponent{}).
					Where("id = ?", existing.ID).
					UpdateColumn("usage_count", gorm.Expr("usage_count + 1")).Error; err != nil {
					return err
				}
				continue
			}

			draft := models.MagistralComponent{
				Name:        strings.TrimSpace(comp.Substance),
				DefaultUnit: comp.Unit,
				Source:      "uso",
				UsageCount:  1,
			}
			if draft.Name == "" {
				continue
			}
			if err := tx.Create(&draft).Error; err != nil {
				return err
			}
			comp.MagistralComponentID = &draft.ID
		}
	}
	return nil
}

// baseRules traz as regras de base ativas. Busca todas e filtra em memória: são poucas dezenas, e
// o casamento por texto sem acento é o mesmo que roda no motor puro.
func (s *MagistralComponentService) baseRules(vehicle string) []models.MagistralBaseIncompatibility {
	if strings.TrimSpace(vehicle) == "" {
		return nil
	}
	var rows []models.MagistralBaseIncompatibility
	if err := s.db.Where("is_active = true").Find(&rows).Error; err != nil {
		return nil
	}
	return rows
}

// in28Limits busca os tetos do Anexo IV só dos nutrientes que aparecem nesta fórmula.
func (s *MagistralComponentService) in28Limits(comps []FormulaCheckComponent) map[string]models.In28Limit {
	nutrientes := make([]string, 0, len(comps))
	for _, c := range comps {
		if c.Catalog != nil && c.Catalog.In28Nutrient != nil {
			nutrientes = append(nutrientes, *c.Catalog.In28Nutrient)
		}
	}
	if len(nutrientes) == 0 {
		return nil
	}
	var rows []models.In28Limit
	if err := s.db.Where("nutrient IN ?", nutrientes).Find(&rows).Error; err != nil {
		// Teto é conferência adicional: se a tabela falhar, o resto do painel continua útil.
		return nil
	}
	out := make(map[string]models.In28Limit, len(rows))
	for _, r := range rows {
		out[r.Nutrient] = r
	}
	return out
}

// CheckFormula resolve os componentes contra o catálogo, busca os pares curados entre eles e
// devolve avisos + cálculo de cápsula. É o endpoint que a tela chama enquanto o médico digita.
func (s *MagistralComponentService) CheckFormula(req *dto.MagistralCheckRequest) (*dto.MagistralCheckResponse, error) {
	check := FormulaCheckInput{PharmaceuticalForm: req.PharmaceuticalForm}
	var ids []uuid.UUID
	catalogByID := map[uuid.UUID]*models.MagistralComponent{}
	asElemental := map[string]bool{}

	for _, c := range req.Components {
		item := FormulaCheckComponent{
			Substance: strings.TrimSpace(c.Substance),
			Quantity:  c.Quantity,
			Unit:      c.Unit,
		}
		asElemental[strings.ToLower(item.Substance)] = c.AsElemental
		item.AsElemental = c.AsElemental

		var found *models.MagistralComponent
		if c.MagistralComponentID != nil && *c.MagistralComponentID != "" {
			if id, err := uuid.Parse(*c.MagistralComponentID); err == nil {
				if cc, err := s.GetByID(id); err == nil {
					found = cc
				}
			}
		}
		if found == nil {
			cc, err := s.findByName(s.db, item.Substance)
			if err != nil {
				return nil, err
			}
			found = cc
		}
		if found != nil {
			item.Catalog = found
			ids = append(ids, found.ID)
			catalogByID[found.ID] = found
		}
		check.Components = append(check.Components, item)
	}

	var pairs []models.MagistralIncompatibility
	if len(ids) > 1 {
		if err := s.db.Where("component_a_id IN ? AND component_b_id IN ?", ids, ids).
			Find(&pairs).Error; err != nil {
			return nil, err
		}
	}

	check.DosesPerDay = req.DosesPerDay
	if check.DosesPerDay <= 0 {
		check.DosesPerDay = DosesPorDia(req.Posology)
	}
	check.In28 = s.in28Limits(check.Components)
	check.Vehicle = strings.TrimSpace(req.Vehicle)
	check.BaseRules = s.baseRules(check.Vehicle)
	alerts := CheckFormula(check, pairs)

	// Cálculo de cápsula só faz sentido em forma sólida encapsulada.
	var capsule *CapsuleAdvice
	if isCapsule(req.PharmaceuticalForm) {
		inputs := make([]CapsuleInput, 0, len(check.Components))
		for _, c := range check.Components {
			mass, ok := MassToMg(c.Quantity, c.Unit)
			if !ok {
				continue
			}
			// Quem ocupa volume na cápsula é o INSUMO, não o elemento. Dose de 300 mg de magnésio
			// elementar com bisglicinato a 30% são 1.000 mg de pó — a cápsula é dimensionada por
			// esse número, e ignorá-lo subestimava o tamanho em toda fórmula com quelato.
			if asElemental[strings.ToLower(c.Substance)] && c.Catalog != nil &&
				c.Catalog.ElementalPercent != nil && *c.Catalog.ElementalPercent > 0 {
				mass = mass * 100 / *c.Catalog.ElementalPercent
			}
			in := CapsuleInput{Substance: c.Substance, MassMg: mass}
			if c.Catalog != nil && c.Catalog.BulkDensity != nil {
				in.BulkDensity = *c.Catalog.BulkDensity
				in.DensityApprox = c.Catalog.DensitySource != nil && *c.Catalog.DensitySource == "classe"
			}
			inputs = append(inputs, in)
		}
		advice := CalculateCapsule(inputs)
		capsule = &advice
	}

	resp := &dto.MagistralCheckResponse{Alerts: []dto.MagistralAlertResponse{}}
	for _, a := range alerts {
		resp.Alerts = append(resp.Alerts, dto.MagistralAlertResponse{
			Level:      string(a.Level),
			Kind:       a.Kind,
			Substances: a.Substances,
			Message:    a.Message,
		})
	}
	if capsule != nil {
		resp.Capsule = capsule
	}
	// A tela usa isto para oferecer "salvar como padrão" só onde falta dado.
	for _, c := range check.Components {
		known := c.Catalog != nil
		hasDose := known && c.Catalog.UsualDose != nil
		hasDensity := known && c.Catalog.BulkDensity != nil
		resp.Components = append(resp.Components, dto.MagistralComponentMatch{
			Substance:  c.Substance,
			Known:      known,
			HasDose:    hasDose,
			HasDensity: hasDensity,
		})
		if known {
			id := c.Catalog.ID.String()
			m := &resp.Components[len(resp.Components)-1]
			// Categoria de receita: é o que faz "testosterona" numa fórmula magistral sair como
			// Controle Especial em vez de receita simples.
			m.DefaultCategory = string(c.Catalog.DefaultCategory)
			if c.Catalog.RegulatoryNote != nil {
				m.RegulatoryNote = strings.TrimSpace(*c.Catalog.RegulatoryNote)
			}
			m.ElementalPercent = c.Catalog.ElementalPercent
			m.DoseAsElemental = c.Catalog.DoseAsElemental
			m.CorrectionNote = c.Catalog.CorrectionNote
			// Conversão pronta para a tela, só quando a dose declarada é do elemento.
			if asElemental[strings.ToLower(c.Substance)] && c.Catalog.ElementalPercent != nil &&
				*c.Catalog.ElementalPercent > 0 {
				if mg, ok := MassToMg(c.Quantity, c.Unit); ok {
					bruto := mg * 100 / *c.Catalog.ElementalPercent
					m.RawMaterialMg = &bruto
				}
			}
			if c.Catalog.PreferredAlternativeID != nil {
				var pref models.MagistralComponent
				if err := s.db.Select("name").First(&pref, "id = ?", *c.Catalog.PreferredAlternativeID).Error; err == nil {
					nome := pref.Name
					m.PreferredName = &nome
					m.PreferenceNote = c.Catalog.PreferenceNote
				}
			}
			m.ID = &id
			m.UsualDose = c.Catalog.UsualDose
			m.DefaultUnit = c.Catalog.DefaultUnit
			m.Indications = c.Catalog.Indications
			m.DoseReference = c.Catalog.DoseReference
			m.IndicationBullets = c.Catalog.IndicationBullets
			m.DoseBullets = c.Catalog.DoseBullets
			m.EvidenceStatus = c.Catalog.EvidenceStatus
		}
	}
	return resp, nil
}

// ListEvidence devolve os trechos do RAG ligados ao componente, fixados primeiro e depois por
// similaridade. É material de LEITURA: nenhum cálculo do sistema consome estes registros.
func (s *MagistralComponentService) ListEvidence(componentID uuid.UUID) ([]models.MagistralComponentArticle, error) {
	var out []models.MagistralComponentArticle
	err := s.db.Preload("Article").
		Where("component_id = ?", componentID).
		Order("pinned DESC").Order("similarity DESC").
		Limit(20).
		Find(&out).Error
	return out, err
}

// PinEvidence marca (ou desmarca) um trecho como o que sustenta a indicação.
func (s *MagistralComponentService) PinEvidence(evidenceID uuid.UUID, pinned bool, userID uuid.UUID) error {
	updates := map[string]any{"pinned": pinned}
	if pinned {
		updates["pinned_by"] = userID
	} else {
		updates["pinned_by"] = nil
	}
	return s.db.Model(&models.MagistralComponentArticle{}).Where("id = ?", evidenceID).Updates(updates).Error
}

// ConfirmEvidence é o aceite do médico: o que veio do RAG deixa de ser sugestão. A distinção
// existe para a tela nunca apresentar texto extraído por máquina como se fosse conferido.
func (s *MagistralComponentService) ConfirmEvidence(componentID, userID uuid.UUID) error {
	return s.db.Model(&models.MagistralComponent{}).Where("id = ?", componentID).
		Updates(map[string]any{
			"evidence_status": "confirmed",
			"reviewed_by":     userID,
			"last_review":     time.Now(),
		}).Error
}

// ListIncompatibilities devolve os pares curados (tela de curadoria).
func (s *MagistralComponentService) ListIncompatibilities() ([]models.MagistralIncompatibility, error) {
	var out []models.MagistralIncompatibility
	err := s.db.Preload("ComponentA").Preload("ComponentB").Order("created_at DESC").Find(&out).Error
	return out, err
}

// CreateIncompatibility cadastra um par. A ordem do par é normalizada no BeforeCreate.
func (s *MagistralComponentService) CreateIncompatibility(req *dto.MagistralIncompatibilityRequest, userID uuid.UUID) (*models.MagistralIncompatibility, error) {
	a, err := uuid.Parse(req.ComponentAID)
	if err != nil {
		return nil, errors.New("componentAId inválido")
	}
	b, err := uuid.Parse(req.ComponentBID)
	if err != nil {
		return nil, errors.New("componentBId inválido")
	}
	if a == b {
		return nil, errors.New("o par precisa ser de substâncias diferentes")
	}

	m := models.MagistralIncompatibility{
		ComponentAID: a,
		ComponentBID: b,
		Severity:     models.IncompatibilitySeverity(req.Severity),
		Mechanism:    strings.TrimSpace(req.Mechanism),
		Note:         req.Note,
		Source:       strings.TrimSpace(req.Source),
		ReviewedBy:   &userID,
	}
	if m.Severity == "" {
		m.Severity = models.IncompatWarn
	}
	if err := s.db.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *MagistralComponentService) DeleteIncompatibility(id uuid.UUID) error {
	return s.db.Delete(&models.MagistralIncompatibility{}, "id = ?", id).Error
}
