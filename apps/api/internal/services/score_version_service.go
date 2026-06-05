package services

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// ScoreVersionService gerencia as versões públicas do escore (Triagem/Light/...) e gera o
// config estático que o site consome. A version só seleciona+ordena itens; a taxonomia,
// pilares, levels e campos de site vêm do próprio score_item (fonte de verdade).
type ScoreVersionService struct {
	db        *gorm.DB
	scoreRepo *repository.ScoreRepository
}

func NewScoreVersionService(db *gorm.DB, scoreRepo *repository.ScoreRepository) *ScoreVersionService {
	return &ScoreVersionService{db: db, scoreRepo: scoreRepo}
}

// ── DTOs ────────────────────────────────────────────────────────────────
type CreateScoreVersionDTO struct {
	Name        string  `json:"name" validate:"required,min=2,max=120"`
	Slug        string  `json:"slug" validate:"required,min=2,max=60"`
	Description *string `json:"description,omitempty"`
	SiteIntro   *string `json:"siteIntro,omitempty"`
	Order       *int    `json:"order,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}

type UpdateScoreVersionDTO struct {
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
	SiteIntro   *string `json:"siteIntro,omitempty"`
	Order       *int    `json:"order,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}

type VersionItemDTO struct {
	ScoreItemID  uuid.UUID `json:"scoreItemId" validate:"required"`
	DisplayOrder int       `json:"displayOrder"`
}

type SetVersionItemsDTO struct {
	Items []VersionItemDTO `json:"items" validate:"required"`
}

// ── CRUD ────────────────────────────────────────────────────────────────
func (s *ScoreVersionService) List() ([]models.ScoreVersion, error) {
	var versions []models.ScoreVersion
	err := s.db.Order(`"order" ASC`).Find(&versions).Error
	return versions, err
}

func (s *ScoreVersionService) GetWithItems(id uuid.UUID) (*models.ScoreVersion, error) {
	var v models.ScoreVersion
	err := s.db.
		Preload("Items", func(db *gorm.DB) *gorm.DB { return db.Order("display_order ASC") }).
		Preload("Items.ScoreItem").
		First(&v, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("score version not found")
	}
	return &v, err
}

func (s *ScoreVersionService) Create(dto CreateScoreVersionDTO) (*models.ScoreVersion, error) {
	v := &models.ScoreVersion{Name: dto.Name, Slug: dto.Slug, Description: dto.Description, SiteIntro: dto.SiteIntro, Active: true}
	if dto.Order != nil {
		v.Order = *dto.Order
	}
	if dto.Active != nil {
		v.Active = *dto.Active
	}
	if err := s.db.Create(v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

// isReservedVersionSlug marca as versões CORE de que o site e o scoring anônimo dependem.
// Não podem ser renomeadas, excluídas, nem esvaziadas pela UI de admin (senão o /escore-light
// e o radar público quebram silenciosamente).
func isReservedVersionSlug(slug string) bool {
	return slug == "light" || slug == "triagem"
}

func (s *ScoreVersionService) Update(id uuid.UUID, dto UpdateScoreVersionDTO) (*models.ScoreVersion, error) {
	var v models.ScoreVersion
	if err := s.db.First(&v, "id = ?", id).Error; err != nil {
		return nil, errors.New("score version not found")
	}
	if dto.Slug != nil && *dto.Slug != v.Slug && isReservedVersionSlug(v.Slug) {
		return nil, fmt.Errorf("a versão %q é reservada e o slug não pode ser alterado", v.Slug)
	}
	if dto.Name != nil {
		v.Name = *dto.Name
	}
	if dto.Slug != nil {
		v.Slug = *dto.Slug
	}
	if dto.Description != nil {
		v.Description = dto.Description
	}
	if dto.SiteIntro != nil {
		v.SiteIntro = dto.SiteIntro
	}
	if dto.Order != nil {
		v.Order = *dto.Order
	}
	if dto.Active != nil {
		v.Active = *dto.Active
	}
	if err := s.db.Save(&v).Error; err != nil {
		return nil, err
	}
	return &v, nil
}

func (s *ScoreVersionService) Delete(id uuid.UUID) error {
	var v models.ScoreVersion
	if err := s.db.First(&v, "id = ?", id).Error; err != nil {
		return errors.New("score version not found")
	}
	if isReservedVersionSlug(v.Slug) {
		return fmt.Errorf("a versão %q é reservada e não pode ser excluída", v.Slug)
	}
	return s.db.Delete(&models.ScoreVersion{}, "id = ?", id).Error
}

// SetItems substitui TODOS os itens da version (replace). Hard-delete + insert em transação.
func (s *ScoreVersionService) SetItems(versionID uuid.UUID, items []VersionItemDTO) error {
	var v models.ScoreVersion
	if err := s.db.First(&v, "id = ?", versionID).Error; err != nil {
		return errors.New("score version not found")
	}
	if isReservedVersionSlug(v.Slug) && len(items) == 0 {
		return fmt.Errorf("a versão %q é reservada e não pode ficar sem itens", v.Slug)
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("version_id = ?", versionID).Delete(&models.ScoreVersionItem{}).Error; err != nil {
			return err
		}
		if len(items) == 0 {
			return nil
		}
		rows := make([]models.ScoreVersionItem, 0, len(items))
		for _, it := range items {
			rows = append(rows, models.ScoreVersionItem{VersionID: versionID, ScoreItemID: it.ScoreItemID, DisplayOrder: it.DisplayOrder})
		}
		return tx.Create(&rows).Error
	})
}

// ── Geração do config para o site ───────────────────────────────────────
// BuildConfig monta a árvore Group→Subgroup→Item→Level apenas com os itens da version,
// ordenados por display_order, reusando LightConfig/mapItemToLightConfig. É a FONTE do
// gerador estático do site (substitui o BuildLightConfig filtrado por is_light_version).
func (s *ScoreVersionService) BuildConfig(slug string) (*LightConfig, error) {
	var v models.ScoreVersion
	if err := s.db.First(&v, "slug = ?", slug).Error; err != nil {
		return nil, errors.New("score version not found")
	}

	var vItems []models.ScoreVersionItem
	if err := s.db.Where("version_id = ?", v.ID).Find(&vItems).Error; err != nil {
		return nil, err
	}
	orderByItem := make(map[uuid.UUID]int, len(vItems))
	for _, vi := range vItems {
		orderByItem[vi.ScoreItemID] = vi.DisplayOrder
	}

	groups, err := s.scoreRepo.GetAllScoreGroupTrees()
	if err != nil {
		return nil, fmt.Errorf("failed to load score tree: %w", err)
	}

	out := &LightConfig{GeneratedAt: time.Now(), Groups: []LightGroupConfig{}}
	itemCount := 0

	appendIfInVersion := func(sgOut *LightSubgroupConfig, it models.ScoreItem, seen map[uuid.UUID]bool) {
		ord, ok := orderByItem[it.ID]
		if !ok || seen[it.ID] {
			return
		}
		cfg := mapItemToLightConfig(it)
		cfg.LightOrder = ord // ordem da version manda
		sgOut.Items = append(sgOut.Items, cfg)
		seen[it.ID] = true
		itemCount++
	}

	for _, g := range groups {
		gOut := LightGroupConfig{ID: g.ID, Name: g.Name, Order: g.Order, Subgroups: []LightSubgroupConfig{}}
		for _, sg := range g.Subgroups {
			sgOut := LightSubgroupConfig{ID: sg.ID, Name: sg.Name, Order: sg.Order, Items: []LightItemConfig{}}
			seen := make(map[uuid.UUID]bool)
			for _, it := range sg.Items {
				appendIfInVersion(&sgOut, it, seen)
				for _, child := range it.ChildItems {
					appendIfInVersion(&sgOut, child, seen)
				}
			}
			if len(sgOut.Items) > 0 {
				sort.SliceStable(sgOut.Items, func(a, b int) bool { return sgOut.Items[a].LightOrder < sgOut.Items[b].LightOrder })
				gOut.Subgroups = append(gOut.Subgroups, sgOut)
			}
		}
		if len(gOut.Subgroups) > 0 {
			out.Groups = append(out.Groups, gOut)
		}
	}

	out.ItemCount = itemCount
	out.Version = fmt.Sprintf("%s-%d", slug, itemCount)
	return out, nil
}
