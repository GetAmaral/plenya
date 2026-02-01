package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ScoreRepository handles database operations for score-related entities
type ScoreRepository struct {
	db *gorm.DB
}

// NewScoreRepository creates a new score repository instance
func NewScoreRepository(db *gorm.DB) *ScoreRepository {
	return &ScoreRepository{db: db}
}

// ============================================================
// ScoreGroup Operations
// ============================================================

// CreateScoreGroup creates a new score group
func (r *ScoreRepository) CreateScoreGroup(group *models.ScoreGroup) error {
	return r.db.Create(group).Error
}

// GetScoreGroupByID retrieves a score group by ID
func (r *ScoreRepository) GetScoreGroupByID(id uuid.UUID) (*models.ScoreGroup, error) {
	var group models.ScoreGroup
	err := r.db.First(&group, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("score group not found")
		}
		return nil, err
	}
	return &group, nil
}

// GetAllScoreGroups retrieves all score groups ordered by order field
func (r *ScoreRepository) GetAllScoreGroups() ([]models.ScoreGroup, error) {
	var groups []models.ScoreGroup
	err := r.db.Order("\"order\" ASC, name ASC").Find(&groups).Error
	return groups, err
}

// GetScoreGroupTree retrieves a score group with all nested data
func (r *ScoreRepository) GetScoreGroupTree(id uuid.UUID) (*models.ScoreGroup, error) {
	var group models.ScoreGroup
	err := r.db.
		Preload("Subgroups", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC, name ASC")
		}).
		Preload("Subgroups.Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC, name ASC")
		}).
		Preload("Subgroups.Items.Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		Preload("Subgroups.Items.ParentItem").
		Preload("Subgroups.Items.ChildItems", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC, name ASC")
		}).
		Preload("Subgroups.Items.ChildItems.Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		First(&group, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("score group not found")
		}
		return nil, err
	}
	return &group, nil
}

// GetAllScoreGroupTrees retrieves all score groups with nested data
func (r *ScoreRepository) GetAllScoreGroupTrees() ([]models.ScoreGroup, error) {
	var groups []models.ScoreGroup
	err := r.db.
		Preload("Subgroups", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC, name ASC")
		}).
		Preload("Subgroups.Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC, name ASC")
		}).
		Preload("Subgroups.Items.Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		Preload("Subgroups.Items.ParentItem").
		Preload("Subgroups.Items.ChildItems", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC, name ASC")
		}).
		Preload("Subgroups.Items.ChildItems.Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		Order("\"order\" ASC, name ASC").
		Find(&groups).Error

	return groups, err
}

// UpdateScoreGroup updates an existing score group
func (r *ScoreRepository) UpdateScoreGroup(group *models.ScoreGroup) error {
	return r.db.Save(group).Error
}

// DeleteScoreGroup soft deletes a score group
func (r *ScoreRepository) DeleteScoreGroup(id uuid.UUID) error {
	result := r.db.Delete(&models.ScoreGroup{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("score group not found")
	}
	return nil
}

// ============================================================
// ScoreSubgroup Operations
// ============================================================

// CreateScoreSubgroup creates a new score subgroup
func (r *ScoreRepository) CreateScoreSubgroup(subgroup *models.ScoreSubgroup) error {
	return r.db.Create(subgroup).Error
}

// GetScoreSubgroupByID retrieves a score subgroup by ID
func (r *ScoreRepository) GetScoreSubgroupByID(id uuid.UUID) (*models.ScoreSubgroup, error) {
	var subgroup models.ScoreSubgroup
	err := r.db.Preload("Group").First(&subgroup, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("score subgroup not found")
		}
		return nil, err
	}
	return &subgroup, nil
}

// GetSubgroupsByGroupID retrieves all subgroups for a specific group
func (r *ScoreRepository) GetSubgroupsByGroupID(groupID uuid.UUID) ([]models.ScoreSubgroup, error) {
	var subgroups []models.ScoreSubgroup
	err := r.db.
		Where("group_id = ?", groupID).
		Order("\"order\" ASC, name ASC").
		Find(&subgroups).Error
	return subgroups, err
}

// UpdateScoreSubgroup updates an existing score subgroup
func (r *ScoreRepository) UpdateScoreSubgroup(subgroup *models.ScoreSubgroup) error {
	return r.db.Save(subgroup).Error
}

// DeleteScoreSubgroup soft deletes a score subgroup
func (r *ScoreRepository) DeleteScoreSubgroup(id uuid.UUID) error {
	result := r.db.Delete(&models.ScoreSubgroup{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("score subgroup not found")
	}
	return nil
}

// ============================================================
// ScoreItem Operations
// ============================================================

// CreateScoreItem creates a new score item
func (r *ScoreRepository) CreateScoreItem(item *models.ScoreItem) error {
	return r.db.Create(item).Error
}

// GetScoreItemByID retrieves a score item by ID with levels
func (r *ScoreRepository) GetScoreItemByID(id uuid.UUID) (*models.ScoreItem, error) {
	var item models.ScoreItem
	err := r.db.
		Preload("Subgroup").
		Preload("Subgroup.Group").
		Preload("ParentItem").
		Preload("ChildItems").
		Preload("Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		First(&item, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("score item not found")
		}
		return nil, err
	}
	return &item, nil
}

// GetItemsBySubgroupID retrieves all items for a specific subgroup
func (r *ScoreRepository) GetItemsBySubgroupID(subgroupID uuid.UUID) ([]models.ScoreItem, error) {
	var items []models.ScoreItem
	err := r.db.
		Where("subgroup_id = ?", subgroupID).
		Preload("Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		Preload("ParentItem").
		Preload("ChildItems").
		Order("\"order\" ASC, name ASC").
		Find(&items).Error
	return items, err
}

// GetAllScoreItems retrieves all score items with group/subgroup info, ordered by hierarchy
func (r *ScoreRepository) GetAllScoreItems() ([]models.ScoreItem, error) {
	var items []models.ScoreItem
	err := r.db.
		Preload("Subgroup").
		Preload("Subgroup.Group").
		Joins("LEFT JOIN score_subgroups ON score_items.subgroup_id = score_subgroups.id").
		Joins("LEFT JOIN score_groups ON score_subgroups.group_id = score_groups.id").
		Order("score_groups.\"order\" ASC").
		Order("score_subgroups.\"order\" ASC").
		Order("score_items.\"order\" ASC").
		Order("score_items.name ASC").
		Find(&items).Error
	return items, err
}

// UpdateScoreItem updates an existing score item
func (r *ScoreRepository) UpdateScoreItem(item *models.ScoreItem) error {
	return r.db.Save(item).Error
}

// DeleteScoreItem soft deletes a score item
func (r *ScoreRepository) DeleteScoreItem(id uuid.UUID) error {
	result := r.db.Delete(&models.ScoreItem{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("score item not found")
	}
	return nil
}

// ============================================================
// ScoreLevel Operations
// ============================================================

// CreateScoreLevel creates a new score level
func (r *ScoreRepository) CreateScoreLevel(level *models.ScoreLevel) error {
	return r.db.Create(level).Error
}

// GetScoreLevelByID retrieves a score level by ID
func (r *ScoreRepository) GetScoreLevelByID(id uuid.UUID) (*models.ScoreLevel, error) {
	var level models.ScoreLevel
	err := r.db.Preload("Item").First(&level, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("score level not found")
		}
		return nil, err
	}
	return &level, nil
}

// GetLevelsByItemID retrieves all levels for a specific item
func (r *ScoreRepository) GetLevelsByItemID(itemID uuid.UUID) ([]models.ScoreLevel, error) {
	var levels []models.ScoreLevel
	err := r.db.
		Where("item_id = ?", itemID).
		Order("level ASC").
		Find(&levels).Error
	return levels, err
}

// UpdateScoreLevel updates an existing score level
func (r *ScoreRepository) UpdateScoreLevel(level *models.ScoreLevel) error {
	return r.db.Save(level).Error
}

// DeleteScoreLevel soft deletes a score level
func (r *ScoreRepository) DeleteScoreLevel(id uuid.UUID) error {
	result := r.db.Delete(&models.ScoreLevel{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("score level not found")
	}
	return nil
}

// ============================================================
// Utility Operations
// ============================================================

// GetMaxOrderForGroup returns the maximum order value for groups
func (r *ScoreRepository) GetMaxOrderForGroup() (int, error) {
	var maxOrder int
	err := r.db.Model(&models.ScoreGroup{}).
		Select("COALESCE(MAX(\"order\"), -1)").
		Scan(&maxOrder).Error
	return maxOrder, err
}

// GetMaxOrderForSubgroup returns the maximum order value for subgroups in a group
func (r *ScoreRepository) GetMaxOrderForSubgroup(groupID uuid.UUID) (int, error) {
	var maxOrder int
	err := r.db.Model(&models.ScoreSubgroup{}).
		Where("group_id = ?", groupID).
		Select("COALESCE(MAX(\"order\"), -1)").
		Scan(&maxOrder).Error
	return maxOrder, err
}

// GetMaxOrderForItem returns the maximum order value for items in a subgroup
func (r *ScoreRepository) GetMaxOrderForItem(subgroupID uuid.UUID) (int, error) {
	var maxOrder int
	err := r.db.Model(&models.ScoreItem{}).
		Where("subgroup_id = ?", subgroupID).
		Select("COALESCE(MAX(\"order\"), -1)").
		Scan(&maxOrder).Error
	return maxOrder, err
}
