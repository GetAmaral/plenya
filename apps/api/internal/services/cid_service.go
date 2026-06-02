package services

import (
	"strings"

	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// CIDService — busca no catálogo curado de CID-10.
type CIDService struct {
	db *gorm.DB
}

func NewCIDService(db *gorm.DB) *CIDService {
	return &CIDService{db: db}
}

// Search casa por código (prefixo) ou descrição (unaccent ILIKE). Vazio = top N.
func (s *CIDService) Search(q string, limit int) ([]models.CIDCode, error) {
	if limit <= 0 || limit > 50 {
		limit = 20
	}
	var rows []models.CIDCode
	query := s.db.Model(&models.CIDCode{})
	q = strings.TrimSpace(q)
	if q != "" {
		like := "%" + q + "%"
		query = query.Where(
			"code ILIKE ? OR unaccent(description) ILIKE unaccent(?)",
			q+"%", like,
		)
	}
	if err := query.Order("code ASC").Limit(limit).Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
