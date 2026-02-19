package repository

import (
	"time"

	"github.com/plenya/api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ArticleRepository gerencia operações de banco de dados para artigos
type ArticleRepository struct {
	db *gorm.DB
}

// NewArticleRepository cria uma nova instância do repositório
func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

// Create cria um novo artigo
func (r *ArticleRepository) Create(article *models.Article) (*models.Article, error) {
	if err := r.db.Create(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

// FindByID busca um artigo por ID
func (r *ArticleRepository) FindByID(id uuid.UUID) (*models.Article, error) {
	var article models.Article
	if err := r.db.
		Preload("ScoreItems.Subgroup.Group").
		First(&article, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// FindByFileHash busca um artigo por hash do arquivo (anti-duplicação)
func (r *ArticleRepository) FindByFileHash(hash string) (*models.Article, error) {
	var article models.Article
	if err := r.db.First(&article, "file_hash = ?", hash).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

// List lista artigos com paginação e filtros
func (r *ArticleRepository) List(page, pageSize int, filters map[string]interface{}) ([]*models.Article, int64, error) {
	var articles []*models.Article
	var total int64

	query := r.db.Model(&models.Article{})

	// Por padrão, esconde capítulos de livro da listagem principal
	// Capítulos são exibidos via GET /articles/:id/chapters
	query = query.Where("source_type != 'book_chapter'")

	// Aplicar filtros
	if journal, ok := filters["journal"].(string); ok && journal != "" {
		query = query.Where("journal ILIKE ?", "%"+journal+"%")
	}

	if specialty, ok := filters["specialty"].(string); ok && specialty != "" {
		query = query.Where("specialty = ?", specialty)
	}

	if articleType, ok := filters["articleType"].(string); ok && articleType != "" {
		query = query.Where("article_type = ?", articleType)
	}

	if favorite, ok := filters["favorite"].(bool); ok {
		query = query.Where("favorite = ?", favorite)
	}

	if rating, ok := filters["rating"].(int); ok && rating >= 0 && rating <= 5 {
		query = query.Where("rating = ?", rating)
	}

	// Filtro por data (publishedAfter, publishedBefore)
	if publishedAfter, ok := filters["publishedAfter"].(time.Time); ok {
		query = query.Where("publish_date >= ?", publishedAfter)
	}

	if publishedBefore, ok := filters["publishedBefore"].(time.Time); ok {
		query = query.Where("publish_date <= ?", publishedBefore)
	}

	// Contar total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Paginação
	offset := (page - 1) * pageSize
	if err := query.
		Order("publish_date DESC, created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// Update atualiza um artigo
func (r *ArticleRepository) Update(article *models.Article) (*models.Article, error) {
	if err := r.db.Save(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

// Delete deleta um artigo (soft delete)
func (r *ArticleRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Article{}, "id = ?", id).Error
}

// Search busca artigos por texto com suporte a acentos (unaccent + pg_trgm)
// Ignora acentos em ambas as direções: "nutrição" encontra "nutricao" e vice-versa
func (r *ArticleRepository) Search(query string, page, pageSize int) ([]*models.Article, int64, error) {
	var articles []*models.Article
	var total int64

	// Usa immutable_unaccent para ignorar acentos na query e nos campos
	// Mantém ILIKE para compatibilidade, mas agora accent-insensitive
	// Exclui capítulos de livro — busca retorna artigos e livros apenas
	normalizedQuery := "%" + query + "%"
	searchQuery := r.db.Model(&models.Article{}).
		Where("source_type != 'book_chapter'").
		Where(
		`lower(immutable_unaccent(title)) ILIKE lower(immutable_unaccent(?))
		OR lower(immutable_unaccent(authors)) ILIKE lower(immutable_unaccent(?))
		OR lower(immutable_unaccent(COALESCE(abstract, ''))) ILIKE lower(immutable_unaccent(?))
		OR lower(immutable_unaccent(journal)) ILIKE lower(immutable_unaccent(?))`,
		normalizedQuery, normalizedQuery, normalizedQuery, normalizedQuery,
	)

	// Contar total
	if err := searchQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Paginação
	offset := (page - 1) * pageSize
	if err := searchQuery.
		Order("publish_date DESC, created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// UpdateLastAccessed atualiza o timestamp de último acesso
func (r *ArticleRepository) UpdateLastAccessed(id uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&models.Article{}).
		Where("id = ?", id).
		Update("last_accessed_at", now).Error
}

// GetByDOI busca artigo por DOI
func (r *ArticleRepository) GetByDOI(doi string) (*models.Article, error) {
	var article models.Article
	if err := r.db.First(&article, "doi = ?", doi).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

// GetByPMID busca artigo por PMID
func (r *ArticleRepository) GetByPMID(pmid string) (*models.Article, error) {
	var article models.Article
	if err := r.db.First(&article, "pmid = ?", pmid).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

// GetRecentArticles retorna artigos recentes
func (r *ArticleRepository) GetRecentArticles(limit int) ([]*models.Article, error) {
	var articles []*models.Article
	if err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// GetFavorites retorna artigos favoritos
func (r *ArticleRepository) GetFavorites(page, pageSize int) ([]*models.Article, int64, error) {
	var articles []*models.Article
	var total int64

	query := r.db.Model(&models.Article{}).
		Where("source_type != 'book_chapter'").
		Where("favorite = ?", true)

	// Contar total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Paginação
	offset := (page - 1) * pageSize
	if err := query.
		Order("publish_date DESC, created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// GetStatistics retorna estatísticas gerais
func (r *ArticleRepository) GetStatistics() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total de artigos
	var total int64
	if err := r.db.Model(&models.Article{}).Count(&total).Error; err != nil {
		return nil, err
	}
	stats["total"] = total

	// Total de favoritos
	var favorites int64
	if err := r.db.Model(&models.Article{}).Where("favorite = ?", true).Count(&favorites).Error; err != nil {
		return nil, err
	}
	stats["favorites"] = favorites

	// Contagem por tipo
	var typeStats []struct {
		ArticleType string
		Count       int64
	}
	if err := r.db.Model(&models.Article{}).
		Select("article_type, COUNT(*) as count").
		Group("article_type").
		Scan(&typeStats).Error; err != nil {
		return nil, err
	}

	typeCounts := make(map[string]int64)
	for _, stat := range typeStats {
		typeCounts[stat.ArticleType] = stat.Count
	}
	stats["byType"] = typeCounts

	// Journals mais comuns
	var journalStats []struct {
		Journal string
		Count   int64
	}
	if err := r.db.Model(&models.Article{}).
		Select("journal, COUNT(*) as count").
		Group("journal").
		Order("count DESC").
		Limit(10).
		Scan(&journalStats).Error; err != nil {
		return nil, err
	}

	journals := make([]map[string]interface{}, len(journalStats))
	for i, stat := range journalStats {
		journals[i] = map[string]interface{}{
			"journal": stat.Journal,
			"count":   stat.Count,
		}
	}
	stats["topJournals"] = journals

	return stats, nil
}

// Batch creates multiple articles
func (r *ArticleRepository) BatchCreate(articles []*models.Article) error {
	return r.db.CreateInBatches(articles, 100).Error
}

// GetArticlesBySpecialty busca artigos por especialidade
func (r *ArticleRepository) GetArticlesBySpecialty(specialty string, page, pageSize int) ([]*models.Article, int64, error) {
	var articles []*models.Article
	var total int64

	query := r.db.Model(&models.Article{}).Where("specialty = ?", specialty)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.
		Order("publish_date DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// AddScoreItems adiciona itens de escore a um artigo (many-to-many)
func (r *ArticleRepository) AddScoreItems(articleID uuid.UUID, scoreItemIDs []uuid.UUID) error {
	var article models.Article
	if err := r.db.First(&article, articleID).Error; err != nil {
		return err
	}

	// Buscar os ScoreItems pelos IDs
	var scoreItems []models.ScoreItem
	if err := r.db.Find(&scoreItems, scoreItemIDs).Error; err != nil {
		return err
	}

	// Adicionar associação usando GORM Association
	return r.db.Model(&article).Association("ScoreItems").Append(&scoreItems)
}

// RemoveScoreItems remove itens de escore de um artigo (many-to-many)
func (r *ArticleRepository) RemoveScoreItems(articleID uuid.UUID, scoreItemIDs []uuid.UUID) error {
	var article models.Article
	if err := r.db.First(&article, articleID).Error; err != nil {
		return err
	}

	// Buscar os ScoreItems pelos IDs
	var scoreItems []models.ScoreItem
	if err := r.db.Find(&scoreItems, scoreItemIDs).Error; err != nil {
		return err
	}

	// Remover associação usando GORM Association
	return r.db.Model(&article).Association("ScoreItems").Delete(&scoreItems)
}

// GetScoreItemsForArticle retorna todos os itens de escore associados a um artigo
func (r *ArticleRepository) GetScoreItemsForArticle(articleID uuid.UUID) ([]models.ScoreItem, error) {
	var article models.Article
	if err := r.db.Preload("ScoreItems").First(&article, articleID).Error; err != nil {
		return nil, err
	}

	return article.ScoreItems, nil
}

// GetArticlesForScoreItem retorna todos os artigos associados a um item de escore
func (r *ArticleRepository) GetArticlesForScoreItem(scoreItemID uuid.UUID) ([]models.Article, error) {
	var scoreItem models.ScoreItem
	if err := r.db.Preload("Articles").First(&scoreItem, scoreItemID).Error; err != nil {
		return nil, err
	}

	return scoreItem.Articles, nil
}

// GetChaptersByBookID retorna os capítulos de um livro, ordenados por chapter_number
func (r *ArticleRepository) GetChaptersByBookID(bookID uuid.UUID) ([]*models.Article, error) {
	var chapters []*models.Article
	if err := r.db.
		Where("parent_article_id = ? AND source_type = 'book_chapter'", bookID).
		Order("chapter_number ASC").
		Find(&chapters).Error; err != nil {
		return nil, err
	}
	return chapters, nil
}

// UpdateScoreItemsLastReview atualiza o campo last_review dos ScoreItems especificados
func (r *ArticleRepository) UpdateScoreItemsLastReview(scoreItemIDs []uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&models.ScoreItem{}).
		Where("id IN ?", scoreItemIDs).
		Update("last_review", now).Error
}
