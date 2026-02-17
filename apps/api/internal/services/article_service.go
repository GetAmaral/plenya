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
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

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
	queueService *EmbeddingQueueService
	aiService    *AIService
}

// sanitizeUTF8Text remove caracteres inválidos UTF-8 e NULL bytes do texto
// CRÍTICO para evitar erros de encoding ao salvar no PostgreSQL
func sanitizeUTF8Text(s string) string {
	// Primeiro, remover NULL bytes (0x00) que são inválidos no PostgreSQL
	s = strings.ReplaceAll(s, "\x00", "")

	// Se já é UTF-8 válido, retornar
	if utf8.ValidString(s) {
		return s
	}

	// Remove caracteres UTF-8 inválidos
	var cleaned strings.Builder
	cleaned.Grow(len(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size == 1 {
			// Caractere inválido - pula
			i++
			continue
		}
		cleaned.WriteRune(r)
		i += size
	}

	return cleaned.String()
}

// NewArticleService cria uma nova instância do serviço
func NewArticleService(db *gorm.DB, uploadFolder string, queueService *EmbeddingQueueService, aiService *AIService) *ArticleService {
	return &ArticleService{
		repo:         repository.NewArticleRepository(db),
		uploadFolder: uploadFolder,
		queueService: queueService,
		aiService:    aiService,
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

	// Criar artigo
	article, err := s.repo.Create(article)
	if err != nil {
		return nil, err
	}

	// Auto-queue para embedding (assíncrono, não bloqueia request)
	if s.queueService != nil {
		go s.queueService.QueueArticle(article.ID)
	}

	return article, nil
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

	// Sanitizar texto para evitar erros de encoding UTF-8
	cleanAbstract := sanitizeUTF8Text(metadata.Abstract)
	cleanFullContent := sanitizeUTF8Text(metadata.FullContent)
	cleanTitle := sanitizeUTF8Text(title)
	cleanAuthors := sanitizeUTF8Text(authors)

	article := &models.Article{
		ID:           articleID,
		Title:        cleanTitle,
		Authors:      cleanAuthors,
		Journal:      journal,
		PublishDate:  publishDate,
		Language:     "en",
		Abstract:     &cleanAbstract,
		FullContent:  &cleanFullContent,
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

	// Criar artigo
	article, err = s.repo.Create(article)
	if err != nil {
		return nil, err
	}

	// Auto-queue para embedding (assíncrono, não bloqueia request)
	if s.queueService != nil {
		go s.queueService.QueueArticle(article.ID)
	}

	return article, nil
}

// ExtractPDFMetadata extrai metadados de um arquivo PDF
func (s *ArticleService) ExtractPDFMetadata(filepath string) (*PDFMetadata, error) {
	metadata := &PDFMetadata{
		Keywords: []string{},
	}

	var content string
	var extractedWithGo bool

	// TENTATIVA 1: Extrair com biblioteca Go (ledongthuc/pdf)
	file, reader, err := pdf.Open(filepath)
	if err == nil {
		defer file.Close()

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

		content = strings.TrimSpace(fullText.String())
		extractedWithGo = true
	} else {
		fmt.Printf("⚠️  Biblioteca Go falhou ao abrir PDF: %v\n", err)
	}

	// FALLBACK: Se biblioteca Go falhou OU extraiu pouco conteúdo, usar pdftotext
	if !extractedWithGo || len(content) < 100 {
		if extractedWithGo {
			fmt.Printf("⚠️  Biblioteca Go extraiu apenas %d chars, tentando pdftotext...\n", len(content))
		} else {
			fmt.Println("⚙️  Tentando extração com pdftotext...")
		}

		cmd := exec.Command("pdftotext", filepath, "-")
		output, err := cmd.Output()

		if err == nil && len(output) > 100 {
			content = strings.TrimSpace(string(output))
			fmt.Printf("✅ pdftotext extraiu %d caracteres com sucesso\n", len(content))
		} else {
			// Ambos falharam
			if err != nil {
				fmt.Printf("⚠️  pdftotext também falhou: %v\n", err)
			} else {
				fmt.Printf("⚠️  pdftotext extraiu apenas %d chars\n", len(output))
			}

			// Se nenhum método funcionou, retornar erro
			if len(content) < 100 {
				return nil, fmt.Errorf("falha ao extrair texto do PDF (Go e pdftotext falharam)")
			}
		}
	}

	metadata.FullContent = content

	// EXTRAÇÃO INTELIGENTE: Usar Claude para extrair metadados da primeira página
	firstPageText := s.extractFirstPage(content)

	if s.aiService != nil && len(firstPageText) > 500 {
		fmt.Println("🤖 Usando Claude Haiku para extração inteligente de metadados...")

		aiMetadata, err := s.aiService.ExtractArticleMetadata(firstPageText)
		if err != nil {
			fmt.Printf("⚠️  Extração com Claude falhou: %v\n", err)
			fmt.Println("   Usando fallback regex...")
		} else {
			// Aplicar metadados extraídos pelo Claude
			if title, ok := aiMetadata["title"].(string); ok && title != "" {
				metadata.Title = title
				fmt.Printf("   ✓ Title: %s\n", truncateString(title, 80))
			}

			if authors, ok := aiMetadata["authors"].(string); ok && authors != "" {
				metadata.Authors = authors
				fmt.Printf("   ✓ Authors: %s\n", truncateString(authors, 80))
			}

			if doi, ok := aiMetadata["doi"].(string); ok && doi != "" {
				metadata.DOI = doi
				fmt.Printf("   ✓ DOI: %s\n", doi)
			}

			if pmid, ok := aiMetadata["pmid"].(string); ok && pmid != "" {
				metadata.PMID = pmid
				fmt.Printf("   ✓ PMID: %s\n", pmid)
			}

			if abstract, ok := aiMetadata["abstract"].(string); ok && abstract != "" {
				metadata.Abstract = abstract
				fmt.Printf("   ✓ Abstract: %d chars\n", len(abstract))
			}

			if keywords, ok := aiMetadata["keywords"].([]interface{}); ok && len(keywords) > 0 {
				metadata.Keywords = make([]string, len(keywords))
				for i, kw := range keywords {
					if kwStr, ok := kw.(string); ok {
						metadata.Keywords[i] = kwStr
					}
				}
				fmt.Printf("   ✓ Keywords: %d found\n", len(metadata.Keywords))
			}

			// Se Claude extraiu metadados, retornar agora (pular regex fallback)
			if metadata.Title != "" && metadata.Authors != "" {
				return metadata, nil
			}
		}
	}

	// FALLBACK REGEX: Se Claude não disponível ou falhou, usar extração regex
	fmt.Println("🔍 Usando extração regex (fallback)...")

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
	if metadata.Title == "" {
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if len(trimmed) > 20 && len(trimmed) < 500 {
				metadata.Title = trimmed
				break
			}
		}
	}

	// Tentar extrair abstract (texto após palavra "abstract")
	if metadata.Abstract == "" {
		abstractRegex := regexp.MustCompile(`(?is)abstract[:\s]+(.*?)(?:introduction|keywords|background|\n\n)`)
		if matches := abstractRegex.FindStringSubmatch(content); len(matches) > 1 {
			abstract := strings.TrimSpace(matches[1])
			if len(abstract) > 50 {
				metadata.Abstract = abstract
			}
		}
	}

	return metadata, nil
}

// extractFirstPage - extrai aproximadamente a primeira página do texto (primeiros 4000 chars)
func (s *ArticleService) extractFirstPage(fullText string) string {
	maxChars := 4000
	if len(fullText) < maxChars {
		return fullText
	}
	return fullText[:maxChars]
}

// truncateString - trunca string para exibição
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
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

// GetArticlesForScoreItem retorna todos os artigos associados a um item de escore
func (s *ArticleService) GetArticlesForScoreItem(scoreItemID uuid.UUID) ([]models.Article, error) {
	return s.repo.GetArticlesForScoreItem(scoreItemID)
}
