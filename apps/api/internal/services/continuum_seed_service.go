// Package services — ContinuumSeedService cria os templates oficiais
// (Continuum Semestral, Anual + 3 boxes padrão) na primeira inicialização.
//
// Idempotente: se já existir template com o mesmo Name, pula.
package services

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type ContinuumSeedService struct {
	db *gorm.DB
}

func NewContinuumSeedService(db *gorm.DB) *ContinuumSeedService {
	return &ContinuumSeedService{db: db}
}

// SeedOfficialTemplates cria templates oficiais se ainda não existirem.
// Roda no boot. Usa um system user UUID (00000000-0000-0000-0000-000000000001)
// pra CreatedByUserID quando não há um usuário humano disponível — quando
// qualquer admin editar via UI, esse campo é sobrescrito.
func (s *ContinuumSeedService) SeedOfficialTemplates() error {
	systemUser := uuid.MustParse("00000000-0000-0000-0000-000000000001")

	// Boxes oficiais primeiro (templates de programa referenciam eles)
	boxBoasVindas, err := s.upsertBox("Box Boas-vindas", "Caixa de boas-vindas enviada ao paciente no início do Continuum.", `## Conteúdo padrão
- Carta de boas-vindas personalizada
- Garrafa Plenya
- Folder do Método AGIR
- Brindes Plenya selecionados`, systemUser)
	if err != nil {
		return err
	}
	boxMensal, err := s.upsertBox("Box Mensal", "Caixa de reposição mensal — suplementos e manipulados conforme protocolo.", `## Conteúdo padrão
- Suplementos do protocolo (revisar mês a mês)
- Manipulados específicos
- Brinde Plenya do mês`, systemUser)
	if err != nil {
		return err
	}
	boxReavaliacao, err := s.upsertBox("Box Reavaliação", "Caixa enviada após reavaliação trimestral com ajustes do protocolo.", `## Conteúdo padrão
- Suplementos ajustados pós-reavaliação
- Carta com leitura do trimestre
- Brindes Plenya`, systemUser)
	if err != nil {
		return err
	}

	// Continuum Semestral (26 semanas)
	if err := s.upsertProgramTemplate("Continuum Semestral",
		"Programa de 6 meses com acompanhamento semanal multidisciplinar (Médico, Nutricionista, Psicólogo, Educador Físico), reavaliação trimestral e Box Plenya.",
		26, systemUser, buildSemestralItems(boxBoasVindas.ID, boxMensal.ID, boxReavaliacao.ID)); err != nil {
		return err
	}

	// Continuum Anual (52 semanas)
	if err := s.upsertProgramTemplate("Continuum Anual",
		"Programa de 12 meses com mesma estrutura do Semestral, com reavaliações trimestrais adicionais e maior consolidação dos hábitos.",
		52, systemUser, buildAnualItems(boxBoasVindas.ID, boxMensal.ID, boxReavaliacao.ID)); err != nil {
		return err
	}

	return nil
}

func (s *ContinuumSeedService) upsertBox(name, desc, contents string, by uuid.UUID) (*models.ContinuumBoxTemplate, error) {
	var existing models.ContinuumBoxTemplate
	err := s.db.Where("name = ?", name).First(&existing).Error
	if err == nil {
		return &existing, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}
	row := models.ContinuumBoxTemplate{
		Name:            name,
		Description:     desc,
		Contents:        contents,
		Status:          "active",
		CreatedByUserID: by,
	}
	if err := s.db.Create(&row).Error; err != nil {
		return nil, err
	}
	log.Printf("✅ Continuum seed: created box template %q", name)
	return &row, nil
}

func (s *ContinuumSeedService) upsertProgramTemplate(name, desc string, weeks int, by uuid.UUID, items []models.ContinuumTemplateItem) error {
	var existing models.ContinuumTemplate
	err := s.db.Where("name = ?", name).First(&existing).Error
	if err == nil {
		return nil // já existe — não sobrescreve (admin pode ter editado)
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	t := models.ContinuumTemplate{
		Name:            name,
		Description:     desc,
		DurationWeeks:   weeks,
		Status:          "active",
		CreatedByUserID: by,
	}
	if err := s.db.Create(&t).Error; err != nil {
		return err
	}
	for i := range items {
		items[i].TemplateID = t.ID
		if err := s.db.Create(&items[i]).Error; err != nil {
			return err
		}
	}
	log.Printf("✅ Continuum seed: created program template %q with %d items", name, len(items))
	return nil
}

// buildSemestralItems gera 4 anamneses iniciais (sem 0) + 1 milestone reunião
// (sem 1) + rotação semanal Med/Nutri/Psico/EF (sem 2-25, 24 items) +
// 1 reavaliação (sem 12) + 3 boxes (sem 0, 12, 24) + 1 fechamento (sem 25).
func buildSemestralItems(boxBoasVindas, boxMensal, boxReav uuid.UUID) []models.ContinuumTemplateItem {
	items := []models.ContinuumTemplateItem{}

	// Semana 0 — 4 anamneses iniciais + box de boas-vindas
	items = append(items,
		appointmentItem(0, 1, "Anamnese inicial — Médico", models.ContinuumSpecialtyDoctor, 1),
		appointmentItem(0, 2, "Anamnese inicial — Nutricionista", models.ContinuumSpecialtyNutritionist, 2),
		appointmentItem(0, 3, "Anamnese inicial — Psicólogo", models.ContinuumSpecialtyPsychologist, 3),
		appointmentItem(0, 4, "Anamnese inicial — Educador Físico", models.ContinuumSpecialtyPhysicalEducator, 4),
		boxItem(0, 0, "Box de Boas-vindas", boxBoasVindas, 5),
	)
	// Semana 1 — Reunião da equipe + plano integrado
	items = append(items, milestoneItem(1, 2, "Reunião da equipe + Plano Integrado",
		"Equipe se reúne, integra o caso e devolve a leitura clínica ao paciente.", 1))

	// Semanas 2-25 — Rotação semanal (Med/Nutri/Psico/EF)
	specRotation := []models.ContinuumItemSpecialty{
		models.ContinuumSpecialtyDoctor,
		models.ContinuumSpecialtyNutritionist,
		models.ContinuumSpecialtyPsychologist,
		models.ContinuumSpecialtyPhysicalEducator,
	}
	titles := map[models.ContinuumItemSpecialty]string{
		models.ContinuumSpecialtyDoctor:           "Encontro semanal — Médico",
		models.ContinuumSpecialtyNutritionist:     "Encontro semanal — Nutricionista",
		models.ContinuumSpecialtyPsychologist:     "Encontro semanal — Psicólogo",
		models.ContinuumSpecialtyPhysicalEducator: "Encontro semanal — Educador Físico",
	}
	for w := 2; w <= 25; w++ {
		spec := specRotation[(w-2)%4]
		items = append(items, appointmentItem(w, 2, titles[spec], spec, 1))
	}

	// Semana 12 — Reavaliação trimestral + Box reavaliação
	items = append(items,
		reassessmentItem(12, 0, "Reavaliação trimestral (3 meses) — Escore + exames", 2),
		boxItem(12, 5, "Box Reavaliação (3 meses)", boxReav, 3),
	)
	// Semana 4, 8, 16, 20 — boxes mensais
	for _, w := range []int{4, 8, 16, 20} {
		items = append(items, boxItem(w, 0, "Box Mensal", boxMensal, 0))
	}

	// Semana 25 — Avaliação final
	items = append(items, milestoneItem(25, 5, "Avaliação Final do Continuum",
		"Fechamento estruturado do período. Equipe e paciente revisam jornada e decidem próximo ciclo.", 2))

	return items
}

// buildAnualItems estende o Semestral pra 52 semanas: continua a rotação
// semanal + adiciona 2 reavaliações trimestrais (sem 24 e 36) + boxes mensais.
func buildAnualItems(boxBoasVindas, boxMensal, boxReav uuid.UUID) []models.ContinuumTemplateItem {
	items := []models.ContinuumTemplateItem{}

	// Semana 0 — 4 anamneses + box boas-vindas
	items = append(items,
		appointmentItem(0, 1, "Anamnese inicial — Médico", models.ContinuumSpecialtyDoctor, 1),
		appointmentItem(0, 2, "Anamnese inicial — Nutricionista", models.ContinuumSpecialtyNutritionist, 2),
		appointmentItem(0, 3, "Anamnese inicial — Psicólogo", models.ContinuumSpecialtyPsychologist, 3),
		appointmentItem(0, 4, "Anamnese inicial — Educador Físico", models.ContinuumSpecialtyPhysicalEducator, 4),
		boxItem(0, 0, "Box de Boas-vindas", boxBoasVindas, 5),
	)
	items = append(items, milestoneItem(1, 2, "Reunião da equipe + Plano Integrado",
		"Equipe se reúne, integra o caso e devolve a leitura clínica ao paciente.", 1))

	specRotation := []models.ContinuumItemSpecialty{
		models.ContinuumSpecialtyDoctor,
		models.ContinuumSpecialtyNutritionist,
		models.ContinuumSpecialtyPsychologist,
		models.ContinuumSpecialtyPhysicalEducator,
	}
	titles := map[models.ContinuumItemSpecialty]string{
		models.ContinuumSpecialtyDoctor:           "Encontro semanal — Médico",
		models.ContinuumSpecialtyNutritionist:     "Encontro semanal — Nutricionista",
		models.ContinuumSpecialtyPsychologist:     "Encontro semanal — Psicólogo",
		models.ContinuumSpecialtyPhysicalEducator: "Encontro semanal — Educador Físico",
	}
	for w := 2; w <= 51; w++ {
		spec := specRotation[(w-2)%4]
		items = append(items, appointmentItem(w, 2, titles[spec], spec, 1))
	}

	// Reavaliações trimestrais — sem 12, 24, 36
	for _, w := range []int{12, 24, 36} {
		items = append(items,
			reassessmentItem(w, 0, "Reavaliação trimestral — Escore + exames", 2),
			boxItem(w, 5, "Box Reavaliação", boxReav, 3),
		)
	}
	// Boxes mensais — semanas 4, 8, 16, 20, 28, 32, 40, 44, 48
	for _, w := range []int{4, 8, 16, 20, 28, 32, 40, 44, 48} {
		items = append(items, boxItem(w, 0, "Box Mensal", boxMensal, 0))
	}

	// Avaliação final — sem 51
	items = append(items, milestoneItem(51, 5, "Avaliação Final do Continuum Anual",
		"Fechamento estruturado dos 12 meses. Equipe e paciente decidem próximo ciclo.", 2))

	return items
}

func appointmentItem(week, dayOffset int, title string, spec models.ContinuumItemSpecialty, position int) models.ContinuumTemplateItem {
	s := spec
	return models.ContinuumTemplateItem{
		Type:               models.ContinuumItemAppointment,
		Specialty:          &s,
		Title:              title,
		WeekOffset:         week,
		ExpectedOffsetDays: dayOffset,
		LateAfterDays:      7,
		Position:           position,
	}
}

func boxItem(week, dayOffset int, title string, boxTemplateID uuid.UUID, position int) models.ContinuumTemplateItem {
	id := boxTemplateID
	return models.ContinuumTemplateItem{
		Type:               models.ContinuumItemBox,
		Title:              title,
		WeekOffset:         week,
		ExpectedOffsetDays: dayOffset,
		LateAfterDays:      14,
		BoxTemplateID:      &id,
		Position:           position,
	}
}

func reassessmentItem(week, dayOffset int, title string, position int) models.ContinuumTemplateItem {
	return models.ContinuumTemplateItem{
		Type:               models.ContinuumItemReassessment,
		Title:              title,
		WeekOffset:         week,
		ExpectedOffsetDays: dayOffset,
		LateAfterDays:      14,
		Position:           position,
	}
}

func milestoneItem(week, dayOffset int, title, desc string, position int) models.ContinuumTemplateItem {
	return models.ContinuumTemplateItem{
		Type:               models.ContinuumItemMilestone,
		Title:              title,
		Description:        desc,
		WeekOffset:         week,
		ExpectedOffsetDays: dayOffset,
		LateAfterDays:      7,
		Position:           position,
	}
}
