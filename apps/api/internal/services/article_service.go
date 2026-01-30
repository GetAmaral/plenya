package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"

	"github.com/google/uuid"
	"github.com/ledongthuc/pdf"
	"gorm.io/gorm"
)

// ArticleService gerencia a lógica de negócio de artigos
type ArticleService struct {
	repo         *repository.ArticleRepository
	uploadFolder string
}

// NewArticleService cria uma nova instância do serviço
func NewArticleService(db *gorm.DB, uploadFolder string) *ArticleService {
	return &ArticleService{
		repo:         repository.NewArticleRepository(db),
		uploadFolder: uploadFolder,
	}
}

// CreateArticleDTO representa os dados para criar um artigo
type CreateArticleDTO struct {
	Title        string    `json:"title" validate:"required,min=3,max=1000"`
	Authors      string    `json:"authors" validate:"required,min=1,max=2000"`
	Journal      string    `json:"journal" validate:"required,min=1,max=500"`
	PublishDate  time.Time `json:"publishDate" validate:"required"`
	Language     string    `json:"language" validate:"required,min=2,max=10"`
	DOI          *string   `json:"doi,omitempty" validate:"omitempty,max=255"`
	PMID         *string   `json:"pmid,omitempty" validate:"omitempty,max=50"`
	ISSN         *string   `json:"issn,omitempty" validate:"omitempty,max=20"`
	Abstract     *string   `json:"abstract,omitempty"`
	FullContent  *string   `json:"fullContent,omitempty"`
	Notes        *string   `json:"notes,omitempty"`
	OriginalLink *string   `json:"originalLink,omitempty" validate:"omitempty,url,max=2048"`
	InternalLink *string   `json:"internalLink,omitempty" validate:"omitempty,max=2048"`
	ArticleType  string    `json:"articleType" validate:"required,oneof=research_article review meta_analysis case_study clinical_trial editorial letter protocol lecture"`
	Keywords     []string  `json:"keywords,omitempty"`
	MeshTerms    []string  `json:"meshTerms,omitempty"`
	Specialty    *string   `json:"specialty,omitempty" validate:"omitempty,max=200"`
	Favorite     bool      `json:"favorite"`
	Rating       *int      `json:"rating,omitempty" validate:"omitempty,min=0,max=5"`
}

// UpdateArticleDTO representa os dados para atualizar um artigo
type UpdateArticleDTO struct {
	Title        *string    `json:"title,omitempty" validate:"omitempty,min=3,max=1000"`
	Authors      *string    `json:"authors,omitempty" validate:"omitempty,min=1,max=2000"`
	Journal      *string    `json:"journal,omitempty" validate:"omitempty,min=1,max=500"`
	PublishDate  *time.Time `json:"publishDate,omitempty"`
	Language     *string    `json:"language,omitempty" validate:"omitempty,min=2,max=10"`
	DOI          *string    `json:"doi,omitempty" validate:"omitempty,max=255"`
	PMID         *string    `json:"pmid,omitempty" validate:"omitempty,max=50"`
	ISSN         *string    `json:"issn,omitempty" validate:"omitempty,max=20"`
	Abstract     *string    `json:"abstract,omitempty"`
	FullContent  *string    `json:"fullContent,omitempty"`
	Notes        *string    `json:"notes,omitempty"`
	OriginalLink *string    `json:"originalLink,omitempty" validate:"omitempty,url,max=2048"`
	ArticleType  *string    `json:"articleType,omitempty" validate:"omitempty,oneof=research_article review meta_analysis case_study clinical_trial editorial letter protocol lecture"`
	Keywords     []string   `json:"keywords,omitempty"`
	MeshTerms    []string   `json:"meshTerms,omitempty"`
	Specialty    *string    `json:"specialty,omitempty" validate:"omitempty,max=200"`
	Favorite     *bool      `json:"favorite,omitempty"`
	Rating       *int       `json:"rating,omitempty" validate:"omitempty,min=0,max=5"`
}

// PDFMetadata representa metadados extraídos de um PDF
type PDFMetadata struct {
	Title       string
	Authors     string
	Abstract    string
	FullContent string
	DOI         string
	PMID        string
	Keywords    []string
}

// CreateArticle cria um novo artigo
func (s *ArticleService) CreateArticle(dto *CreateArticleDTO, userID uuid.UUID) (*models.Article, error) {
	article := &models.Article{
		Title:        dto.Title,
		Authors:      dto.Authors,
		Journal:      dto.Journal,
		PublishDate:  dto.PublishDate,
		Language:     dto.Language,
		DOI:          dto.DOI,
		PMID:         dto.PMID,
		ISSN:         dto.ISSN,
		Abstract:     dto.Abstract,
		FullContent:  dto.FullContent,
		Notes:        dto.Notes,
		OriginalLink: dto.OriginalLink,
		InternalLink: dto.InternalLink,
		ArticleType:  dto.ArticleType,
		Keywords:     dto.Keywords,
		MeshTerms:    dto.MeshTerms,
		Specialty:    dto.Specialty,
		Favorite:     dto.Favorite,
		Rating:       dto.Rating,
		CreatedBy:    &userID,
		UpdatedBy:    &userID,
	}

	return s.repo.Create(article)
}

// GetArticleByID busca um artigo por ID
func (s *ArticleService) GetArticleByID(id uuid.UUID) (*models.Article, error) {
	return s.repo.FindByID(id)
}

// ListArticles lista artigos com paginação e filtros
func (s *ArticleService) ListArticles(page, pageSize int, filters map[string]interface{}) ([]*models.Article, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	return s.repo.List(page, pageSize, filters)
}

// UpdateArticle atualiza um artigo
func (s *ArticleService) UpdateArticle(id uuid.UUID, dto *UpdateArticleDTO, userID uuid.UUID) (*models.Article, error) {
	article, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Atualizar campos apenas se fornecidos
	if dto.Title != nil {
		article.Title = *dto.Title
	}
	if dto.Authors != nil {
		article.Authors = *dto.Authors
	}
	if dto.Journal != nil {
		article.Journal = *dto.Journal
	}
	if dto.PublishDate != nil {
		article.PublishDate = *dto.PublishDate
	}
	if dto.Language != nil {
		article.Language = *dto.Language
	}
	if dto.DOI != nil {
		article.DOI = dto.DOI
	}
	if dto.PMID != nil {
		article.PMID = dto.PMID
	}
	if dto.ISSN != nil {
		article.ISSN = dto.ISSN
	}
	if dto.Abstract != nil {
		article.Abstract = dto.Abstract
	}
	if dto.FullContent != nil {
		article.FullContent = dto.FullContent
	}
	if dto.Notes != nil {
		article.Notes = dto.Notes
	}
	if dto.OriginalLink != nil {
		article.OriginalLink = dto.OriginalLink
	}
	if dto.ArticleType != nil {
		article.ArticleType = *dto.ArticleType
	}
	if dto.Keywords != nil {
		article.Keywords = dto.Keywords
	}
	if dto.MeshTerms != nil {
		article.MeshTerms = dto.MeshTerms
	}
	if dto.Specialty != nil {
		article.Specialty = dto.Specialty
	}
	if dto.Favorite != nil {
		article.Favorite = *dto.Favorite
	}
	if dto.Rating != nil {
		article.Rating = dto.Rating
	}

	article.UpdatedBy = &userID

	return s.repo.Update(article)
}

// DeleteArticle deleta um artigo (soft delete)
func (s *ArticleService) DeleteArticle(id uuid.UUID) error {
	article, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Deletar arquivo PDF se existir
	if article.InternalLink != nil && *article.InternalLink != "" {
		filePath := filepath.Join(s.uploadFolder, filepath.Base(*article.InternalLink))
		os.Remove(filePath) // Ignora erro se arquivo não existir
	}

	return s.repo.Delete(id)
}

// UploadPDF faz upload de um PDF e extrai metadados automaticamente
func (s *ArticleService) UploadPDF(fileReader io.Reader, filename string, userID uuid.UUID) (*models.Article, error) {
	// Criar pasta de uploads se não existir
	if err := os.MkdirAll(s.uploadFolder, 0755); err != nil {
		return nil, fmt.Errorf("erro ao criar pasta de uploads: %w", err)
	}

	// Gerar ID único para o arquivo
	articleID := uuid.New()
	ext := filepath.Ext(filename)
	newFilename := articleID.String() + ext
	filePath := filepath.Join(s.uploadFolder, newFilename)

	// Criar arquivo temporário para salvar
	tempFile, err := os.CreateTemp("", "article-*.pdf")
	if err != nil {
		return nil, fmt.Errorf("erro ao criar arquivo temporário: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Copiar dados do reader para arquivo temporário
	fileBytes, err := io.ReadAll(fileReader)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	if _, err := tempFile.Write(fileBytes); err != nil {
		return nil, fmt.Errorf("erro ao escrever arquivo temporário: %w", err)
	}

	// Calcular hash SHA-256
	hash := sha256.Sum256(fileBytes)
	fileHash := hex.EncodeToString(hash[:])

	// Verificar se já existe artigo com este hash (anti-duplicação)
	existingArticle, _ := s.repo.FindByFileHash(fileHash)
	if existingArticle != nil {
		return nil, errors.New("este arquivo PDF já foi importado anteriormente")
	}

	// Extrair metadados do PDF
	metadata, err := s.ExtractPDFMetadata(tempFile.Name())
	if err != nil {
		// Continua mesmo se extração falhar
		metadata = &PDFMetadata{}
	}

	// Se DOI foi encontrado, tentar buscar metadados adicionais via API
	if metadata.DOI != "" {
		pubmedData, err := s.FetchPubMedMetadata(metadata.DOI)
		if err == nil && pubmedData != nil {
			// Enriquecer metadados com dados do PubMed
			if metadata.Title == "" {
				metadata.Title = pubmedData.Title
			}
			if metadata.Authors == "" {
				metadata.Authors = pubmedData.Authors
			}
			if metadata.Abstract == "" {
				metadata.Abstract = pubmedData.Abstract
			}
			metadata.PMID = pubmedData.PMID
			metadata.Keywords = append(metadata.Keywords, pubmedData.Keywords...)
		}
	}

	// Copiar arquivo para destino final
	destFile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar arquivo de destino: %w", err)
	}
	defer destFile.Close()

	if _, err := destFile.Write(fileBytes); err != nil {
		os.Remove(filePath)
		return nil, fmt.Errorf("erro ao salvar arquivo: %w", err)
	}

	// Criar artigo no banco de dados
	fileSize := int64(len(fileBytes))
	internalLink := "./uploads/articles/" + newFilename

	// Valores padrão se extração falhar
	title := metadata.Title
	if title == "" {
		title = strings.TrimSuffix(filename, ext)
	}

	authors := metadata.Authors
	if authors == "" {
		authors = "Autor desconhecido"
	}

	journal := "Revista não identificada"
	publishDate := time.Now()

	article := &models.Article{
		ID:           articleID,
		Title:        title,
		Authors:      authors,
		Journal:      journal,
		PublishDate:  publishDate,
		Language:     "en",
		Abstract:     &metadata.Abstract,
		FullContent:  &metadata.FullContent,
		ArticleType:  "research_article",
		InternalLink: &internalLink,
		FileHash:     &fileHash,
		FileSize:     &fileSize,
		CreatedBy:    &userID,
		UpdatedBy:    &userID,
	}

	if metadata.DOI != "" {
		article.DOI = &metadata.DOI
	}
	if metadata.PMID != "" {
		article.PMID = &metadata.PMID
	}
	if len(metadata.Keywords) > 0 {
		article.Keywords = metadata.Keywords
	}

	return s.repo.Create(article)
}

// ExtractPDFMetadata extrai metadados de um arquivo PDF
func (s *ArticleService) ExtractPDFMetadata(filepath string) (*PDFMetadata, error) {
	file, reader, err := pdf.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir PDF: %w", err)
	}
	defer file.Close()

	metadata := &PDFMetadata{
		Keywords: []string{},
	}

	// Extrair texto completo do PDF
	var fullText strings.Builder
	numPages := reader.NumPage()

	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page := reader.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}

		fullText.WriteString(text)
		fullText.WriteString("\n")
	}

	content := fullText.String()
	metadata.FullContent = strings.TrimSpace(content)

	// Extrair DOI usando regex
	doiRegex := regexp.MustCompile(`(?i)doi[:\s]*([10]\.\d{4,}/[^\s]+)`)
	if matches := doiRegex.FindStringSubmatch(content); len(matches) > 1 {
		metadata.DOI = strings.TrimSpace(matches[1])
	}

	// Extrair PMID
	pmidRegex := regexp.MustCompile(`(?i)pmid[:\s]*(\d{7,})`)
	if matches := pmidRegex.FindStringSubmatch(content); len(matches) > 1 {
		metadata.PMID = strings.TrimSpace(matches[1])
	}

	// Tentar extrair título (primeira linha significativa)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) > 20 && len(trimmed) < 500 {
			metadata.Title = trimmed
			break
		}
	}

	// Tentar extrair abstract (texto após palavra "abstract")
	abstractRegex := regexp.MustCompile(`(?is)abstract[:\s]+(.*?)(?:introduction|keywords|background|\n\n)`)
	if matches := abstractRegex.FindStringSubmatch(content); len(matches) > 1 {
		abstract := strings.TrimSpace(matches[1])
		if len(abstract) > 50 {
			metadata.Abstract = abstract
		}
	}

	return metadata, nil
}

// PubMedMetadata representa dados do PubMed API
type PubMedMetadata struct {
	Title    string
	Authors  string
	Abstract string
	PMID     string
	Keywords []string
}

// FetchPubMedMetadata busca metadados via DOI usando CrossRef API
func (s *ArticleService) FetchPubMedMetadata(doi string) (*PubMedMetadata, error) {
	// Usar CrossRef API (gratuita, sem necessidade de API key)
	url := fmt.Sprintf("https://api.crossref.org/works/%s", doi)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("CrossRef API retornou status %d", resp.StatusCode)
	}

	var result struct {
		Message struct {
			Title    []string `json:"title"`
			Abstract string   `json:"abstract"`
			Author   []struct {
				Given  string `json:"given"`
				Family string `json:"family"`
			} `json:"author"`
			Subject []string `json:"subject"`
		} `json:"message"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	metadata := &PubMedMetadata{
		Keywords: result.Message.Subject,
	}

	if len(result.Message.Title) > 0 {
		metadata.Title = result.Message.Title[0]
	}

	metadata.Abstract = result.Message.Abstract

	// Construir lista de autores
	var authors []string
	for _, author := range result.Message.Author {
		name := fmt.Sprintf("%s %s", author.Given, author.Family)
		authors = append(authors, strings.TrimSpace(name))
	}
	metadata.Authors = strings.Join(authors, ", ")

	return metadata, nil
}

// SearchArticles busca artigos por texto (full-text search)
func (s *ArticleService) SearchArticles(query string, page, pageSize int) ([]*models.Article, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	return s.repo.Search(query, page, pageSize)
}

// UpdateLastAccessed atualiza o timestamp de último acesso
func (s *ArticleService) UpdateLastAccessed(id uuid.UUID) error {
	return s.repo.UpdateLastAccessed(id)
}

// AddScoreItemsToArticle adiciona itens de escore a um artigo (many-to-many)
func (s *ArticleService) AddScoreItemsToArticle(articleID uuid.UUID, scoreItemIDs []uuid.UUID) error {
	if err := s.repo.AddScoreItems(articleID, scoreItemIDs); err != nil {
		return err
	}

	// Atualizar LastReview dos ScoreItems afetados
	return s.repo.UpdateScoreItemsLastReview(scoreItemIDs)
}

// RemoveScoreItemsFromArticle remove itens de escore de um artigo (many-to-many)
func (s *ArticleService) RemoveScoreItemsFromArticle(articleID uuid.UUID, scoreItemIDs []uuid.UUID) error {
	if err := s.repo.RemoveScoreItems(articleID, scoreItemIDs); err != nil {
		return err
	}

	// Atualizar LastReview dos ScoreItems afetados
	return s.repo.UpdateScoreItemsLastReview(scoreItemIDs)
}

// GetScoreItemsForArticle retorna todos os itens de escore associados a um artigo
func (s *ArticleService) GetScoreItemsForArticle(articleID uuid.UUID) ([]models.ScoreItem, error) {
	return s.repo.GetScoreItemsForArticle(articleID)
}
