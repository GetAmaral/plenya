package services

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

// DuplicateFileError é retornado quando o arquivo já foi importado (mesmo SHA-256)
type DuplicateFileError struct {
	ExistingTitle string
	ExistingID    uuid.UUID
}

func (e *DuplicateFileError) Error() string {
	return "arquivo já importado: " + e.ExistingTitle
}

// ChapterContent representa o conteúdo de um capítulo extraído de um livro
type ChapterContent struct {
	Number  int
	Title   string
	Content string
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

// UploadFile faz upload de um arquivo (PDF, EPUB, TXT, MD) e extrai metadados automaticamente
func (s *ArticleService) UploadFile(fileReader io.Reader, filename string, userID uuid.UUID) (*models.Article, error) {
	// Criar pasta de uploads se não existir
	if err := os.MkdirAll(s.uploadFolder, 0755); err != nil {
		return nil, fmt.Errorf("erro ao criar pasta de uploads: %w", err)
	}

	// Gerar ID único para o arquivo
	articleID := uuid.New()
	ext := filepath.Ext(filename)
	newFilename := articleID.String() + ext
	filePath := filepath.Join(s.uploadFolder, newFilename)

	// Criar arquivo temporário preservando a extensão (necessário para detecção de formato)
	tempFile, err := os.CreateTemp("", "article-*"+ext)
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
		return nil, &DuplicateFileError{ExistingTitle: existingArticle.Title, ExistingID: existingArticle.ID}
	}

	// Extrair metadados do arquivo
	metadata, err := s.ExtractFileMetadata(tempFile.Name())
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

// UploadPDF mantém compatibilidade com scripts existentes
func (s *ArticleService) UploadPDF(fileReader io.Reader, filename string, userID uuid.UUID) (*models.Article, error) {
	return s.UploadFile(fileReader, filename, userID)
}

// UploadBook faz upload de um arquivo como livro, dividindo em capítulos automaticamente
func (s *ArticleService) UploadBook(fileReader io.Reader, filename string, userID uuid.UUID) (*models.Article, error) {
	if err := os.MkdirAll(s.uploadFolder, 0755); err != nil {
		return nil, fmt.Errorf("erro ao criar pasta de uploads: %w", err)
	}

	bookID := uuid.New()
	ext := filepath.Ext(filename)
	newFilename := bookID.String() + ext
	filePath := filepath.Join(s.uploadFolder, newFilename)

	tempFile, err := os.CreateTemp("", "book-*"+ext)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar arquivo temporário: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(fileReader)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo: %w", err)
	}
	if _, err := tempFile.Write(fileBytes); err != nil {
		return nil, fmt.Errorf("erro ao escrever arquivo temporário: %w", err)
	}

	hash := sha256.Sum256(fileBytes)
	fileHash := hex.EncodeToString(hash[:])

	existingArticle, _ := s.repo.FindByFileHash(fileHash)
	if existingArticle != nil {
		return nil, &DuplicateFileError{ExistingTitle: existingArticle.Title, ExistingID: existingArticle.ID}
	}

	var chapters []ChapterContent
	extLower := strings.ToLower(ext)
	switch extLower {
	case ".epub":
		fmt.Println("📚 Dividindo EPUB em capítulos...")
		chapters, err = s.splitEPUBIntoChapters(tempFile.Name())
	case ".pdf":
		fmt.Println("📚 Dividindo PDF em capítulos...")
		chapters, err = s.splitPDFChapters(tempFile.Name())
	case ".txt", ".md":
		fmt.Println("📚 Dividindo texto em capítulos...")
		var content string
		content, err = s.extractTextFileContent(tempFile.Name())
		if err == nil {
			chapters, err = s.splitTextIntoChapters(content)
		}
	default:
		return nil, fmt.Errorf("formato não suportado para importação como livro: %s", ext)
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao dividir arquivo em capítulos: %w", err)
	}
	if len(chapters) == 0 {
		return nil, fmt.Errorf("nenhum capítulo encontrado no arquivo")
	}

	// Extrair metadados do livro a partir do início do conteúdo
	var bookTitle, bookAuthors string
	firstContent := ""
	if len(chapters) > 0 {
		firstContent = chapters[0].Content
		if len(firstContent) > 4000 {
			firstContent = firstContent[:4000]
		}
	}
	if s.aiService != nil && len(firstContent) > 500 {
		fmt.Println("🤖 Extraindo metadados do livro via Claude...")
		aiMetadata, metaErr := s.aiService.ExtractArticleMetadata(firstContent)
		if metaErr == nil {
			if title, ok := aiMetadata["title"].(string); ok && title != "" {
				bookTitle = title
			}
			if authors, ok := aiMetadata["authors"].(string); ok && authors != "" {
				bookAuthors = authors
			}
		}
	}
	if bookTitle == "" {
		bookTitle = strings.TrimSuffix(filename, ext)
	}
	if bookAuthors == "" {
		bookAuthors = "Autor desconhecido"
	}

	destFile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar arquivo de destino: %w", err)
	}
	defer destFile.Close()
	if _, err := destFile.Write(fileBytes); err != nil {
		os.Remove(filePath)
		return nil, fmt.Errorf("erro ao salvar arquivo: %w", err)
	}

	fileSize := int64(len(fileBytes))
	internalLink := "./uploads/articles/" + newFilename
	totalChapters := len(chapters)
	sourceTypeBook := "book"

	cleanTitle := sanitizeUTF8Text(bookTitle)
	cleanAuthors := sanitizeUTF8Text(bookAuthors)

	bookArticle := &models.Article{
		ID:            bookID,
		Title:         cleanTitle,
		Authors:       cleanAuthors,
		Journal:       "Livro",
		PublishDate:   time.Now(),
		Language:      "pt",
		ArticleType:   "research_article",
		SourceType:    sourceTypeBook,
		TotalChapters: &totalChapters,
		InternalLink:  &internalLink,
		FileHash:      &fileHash,
		FileSize:      &fileSize,
		CreatedBy:     &userID,
		UpdatedBy:     &userID,
	}

	bookArticle, err = s.repo.Create(bookArticle)
	if err != nil {
		os.Remove(filePath)
		return nil, fmt.Errorf("erro ao criar registro do livro: %w", err)
	}

	fmt.Printf("📖 Criando %d capítulos...\n", len(chapters))
	sourceTypeChapter := "book_chapter"

	for i, chapter := range chapters {
		chapterNum := chapter.Number
		cleanContent := sanitizeUTF8Text(chapter.Content)

		chapterTitle := chapter.Title
		if chapterTitle == "" {
			chapterTitle = fmt.Sprintf("Capítulo %d", chapter.Number)
		}
		cleanChapterTitle := sanitizeUTF8Text(chapterTitle)

		chapterArticleTitle := fmt.Sprintf("%s — %s", cleanTitle, cleanChapterTitle)
		if len(chapterArticleTitle) > 1000 {
			chapterArticleTitle = chapterArticleTitle[:997] + "..."
		}

		chapterArticle := &models.Article{
			Title:           chapterArticleTitle,
			Authors:         cleanAuthors,
			Journal:         "Livro",
			PublishDate:     time.Now(),
			Language:        "pt",
			ArticleType:     "research_article",
			SourceType:      sourceTypeChapter,
			ParentArticleID: &bookArticle.ID,
			ChapterNumber:   &chapterNum,
			ChapterTitle:    &cleanChapterTitle,
			FullContent:     &cleanContent,
			CreatedBy:       &userID,
			UpdatedBy:       &userID,
		}

		chapterArticle, err = s.repo.Create(chapterArticle)
		if err != nil {
			fmt.Printf("⚠️  Erro ao criar capítulo %d: %v\n", i+1, err)
			continue
		}

		if s.queueService != nil {
			go s.queueService.QueueArticle(chapterArticle.ID)
		}
	}

	fmt.Printf("✅ Livro importado: %d capítulos criados\n", len(chapters))
	return bookArticle, nil
}

// splitEPUBIntoChapters extrai capítulos de um arquivo EPUB (um spine item = um capítulo)
func (s *ArticleService) splitEPUBIntoChapters(filePath string) ([]ChapterContent, error) {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir EPUB: %w", err)
	}
	defer r.Close()

	fileIndex := map[string]*zip.File{}
	for _, f := range r.File {
		fileIndex[f.Name] = f
	}

	// Localizar OPF
	opfPath := ""
	for _, f := range r.File {
		if strings.HasSuffix(f.Name, "container.xml") {
			rc, err := f.Open()
			if err != nil {
				continue
			}
			data, _ := io.ReadAll(rc)
			rc.Close()
			re := regexp.MustCompile(`full-path="([^"]+)"`)
			if m := re.FindSubmatch(data); len(m) > 1 {
				opfPath = string(m[1])
			}
			break
		}
	}

	// Parsear OPF para obter spine
	var contentOrder []string
	if opfPath != "" {
		if opfFile, ok := fileIndex[opfPath]; ok {
			rc, err := opfFile.Open()
			if err == nil {
				data, _ := io.ReadAll(rc)
				rc.Close()

				opfDir := ""
				if idx := strings.LastIndex(opfPath, "/"); idx >= 0 {
					opfDir = opfPath[:idx]
				}

				manifestMap := map[string]string{}
				itemRe := regexp.MustCompile(`(?s)<item\s([^>]+)>`)
				idAttrRe := regexp.MustCompile(`\bid="([^"]+)"`)
				hrefAttrRe := regexp.MustCompile(`\bhref="([^"]+)"`)
				for _, m := range itemRe.FindAllSubmatch(data, -1) {
					attrs := string(m[1])
					idM := idAttrRe.FindStringSubmatch(attrs)
					hrefM := hrefAttrRe.FindStringSubmatch(attrs)
					if len(idM) > 1 && len(hrefM) > 1 {
						manifestMap[idM[1]] = hrefM[1]
					}
				}

				spineRe := regexp.MustCompile(`<itemref\s[^>]*\bidref="([^"]+)"`)
				for _, m := range spineRe.FindAllSubmatch(data, -1) {
					idref := string(m[1])
					if href, ok := manifestMap[idref]; ok {
						if opfDir != "" {
							contentOrder = append(contentOrder, opfDir+"/"+href)
						} else {
							contentOrder = append(contentOrder, href)
						}
					}
				}
			}
		}
	}

	// Fallback: todos os XHTML exceto TOC/nav
	if len(contentOrder) == 0 {
		for _, f := range r.File {
			name := strings.ToLower(f.Name)
			if (strings.HasSuffix(name, ".html") || strings.HasSuffix(name, ".xhtml") || strings.HasSuffix(name, ".htm")) &&
				!strings.Contains(name, "toc") && !strings.Contains(name, "nav") {
				contentOrder = append(contentOrder, f.Name)
			}
		}
	}

	var chapters []ChapterContent
	chapterNum := 0
	for _, path := range contentOrder {
		f, ok := fileIndex[path]
		if !ok {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			continue
		}
		data, _ := io.ReadAll(rc)
		rc.Close()

		title := extractHTMLTitle(string(data))
		text := strings.TrimSpace(stripHTMLTags(string(data)))

		// Ignorar spine items muito pequenos (TOC, copyright pages, etc.)
		if len(text) < 200 {
			continue
		}

		chapterNum++
		chapters = append(chapters, ChapterContent{
			Number:  chapterNum,
			Title:   title,
			Content: text,
		})
	}

	if len(chapters) == 0 {
		return nil, fmt.Errorf("nenhum capítulo encontrado no EPUB")
	}
	return chapters, nil
}

// extractHTMLTitle extrai o título principal de um arquivo XHTML (h1 > h2 > title)
func extractHTMLTitle(html string) string {
	h1Re := regexp.MustCompile(`(?i)<h1[^>]*>([^<]{1,300})</h1>`)
	if m := h1Re.FindStringSubmatch(html); len(m) > 1 {
		if t := strings.TrimSpace(m[1]); len(t) > 0 {
			return t
		}
	}
	h2Re := regexp.MustCompile(`(?i)<h2[^>]*>([^<]{1,300})</h2>`)
	if m := h2Re.FindStringSubmatch(html); len(m) > 1 {
		if t := strings.TrimSpace(m[1]); len(t) > 0 {
			return t
		}
	}
	titleRe := regexp.MustCompile(`(?i)<title[^>]*>([^<]{1,300})</title>`)
	if m := titleRe.FindStringSubmatch(html); len(m) > 1 {
		if t := strings.TrimSpace(m[1]); len(t) > 0 {
			return t
		}
	}
	return ""
}

// cleanPDFText limpa artefatos comuns de encoding de PDF:
// - Liga­duras Unicode (ﬁ→fi, ﬀ→ff, etc.)
// - Soft hyphens (\xad)
// - Caracteres da Private Use Area (U+E000–U+F8FF) que não têm mapeamento
func cleanPDFText(text string) string {
	// Ligaduras padrão Unicode
	ligReplacer := strings.NewReplacer(
		"ﬁ", "fi", "ﬀ", "ff", "ﬂ", "fl", "ﬃ", "ffi", "ﬄ", "ffl",
		"\ufb0b", "fi", "\ufb0c", "fi",
	)
	text = ligReplacer.Replace(text)

	// Soft hyphens: substituir por hífen real (não remover, evita palavras grudadas)
	text = strings.ReplaceAll(text, "\xad", "-")

	// Remover caracteres da Private Use Area (garbled PDF artifacts)
	var cleaned strings.Builder
	cleaned.Grow(len(text))
	for _, r := range text {
		if (r >= 0xE000 && r <= 0xF8FF) || (r >= 0xE0000 && r <= 0xEFFFF) {
			continue
		}
		cleaned.WriteRune(r)
	}
	return cleaned.String()
}

// splitByRunningHeader detecta capítulos via cabeçalho repetido em cada página.
// Livros como "Functional Medicine Laboratory" têm "Chapter N" no topo de toda página.
// Estratégia: split em páginas (form feed \x0c), detecta o número do capítulo
// nos primeiros 80 chars de cada página, propaga para páginas sem cabeçalho.
func splitByRunningHeader(pages []string) []ChapterContent {
	// Padrão: "Chapter 3" ou "CAPÍTULO 3" como primeira linha da página
	headerRe := regexp.MustCompile(`(?im)^(?:chapter|cap[íi]tulo)\s+(\d+)\s*$`)

	// Identificar páginas de TOC: páginas com ≥3 capítulos diferentes no cabeçalho
	// são provavelmente sumários, não conteúdo real
	isTOCPage := make([]bool, len(pages))
	for i, p := range pages {
		matches := headerRe.FindAllStringSubmatch(p, -1)
		chapNums := map[int]bool{}
		for _, m := range matches {
			n := 0
			fmt.Sscanf(m[1], "%d", &n)
			if n > 0 {
				chapNums[n] = true
			}
		}
		if len(chapNums) >= 3 {
			isTOCPage[i] = true
		}
	}

	type pageInfo struct {
		text      string
		chapterNo int // 0 = não detectado
	}

	infos := make([]pageInfo, len(pages))
	lastChapter := 0
	detected := 0

	for i, p := range pages {
		if isTOCPage[i] {
			infos[i] = pageInfo{text: p, chapterNo: 0}
			continue
		}
		head := p
		if len(head) > 120 {
			head = head[:120]
		}
		if m := headerRe.FindStringSubmatch(head); len(m) > 1 {
			num := 0
			fmt.Sscanf(m[1], "%d", &num)
			if num > 0 && num != lastChapter {
				lastChapter = num
				detected++
			}
		}
		infos[i] = pageInfo{text: p, chapterNo: lastChapter}
	}

	if detected < 3 {
		return nil
	}

	// Agrupar páginas por capítulo
	type chGroup struct {
		num   int
		pages []string
	}
	var groups []chGroup
	for _, info := range infos {
		if info.chapterNo == 0 {
			continue
		}
		if len(groups) == 0 || groups[len(groups)-1].num != info.chapterNo {
			groups = append(groups, chGroup{num: info.chapterNo})
		}
		groups[len(groups)-1].pages = append(groups[len(groups)-1].pages, info.text)
	}

	var chapters []ChapterContent
	for _, g := range groups {
		combined := cleanPDFText(strings.Join(g.pages, "\n"))
		// Remover números de página isolados e linhas muito curtas
		lines := strings.Split(combined, "\n")
		var kept []string
		for _, l := range lines {
			stripped := strings.TrimSpace(l)
			if regexp.MustCompile(`^\d{1,4}$`).MatchString(stripped) {
				continue
			}
			kept = append(kept, l)
		}
		combined = strings.TrimSpace(strings.Join(kept, "\n"))
		if len(combined) < 500 {
			continue
		}

		// Título: primeira linha não-vazia depois do cabeçalho "Chapter N"
		title := fmt.Sprintf("Capítulo %d", g.num)
		if len(g.pages) > 0 {
			firstPageLines := strings.Split(strings.TrimSpace(g.pages[0]), "\n")
			for _, l := range firstPageLines {
				l = strings.TrimSpace(l)
				if l == "" {
					continue
				}
				// Pular o próprio cabeçalho "Chapter N"
				if headerRe.MatchString(l) {
					continue
				}
				if len(l) > 5 && len(l) < 200 {
					title = l
					break
				}
			}
		}

		chapters = append(chapters, ChapterContent{
			Number:  g.num,
			Title:   title,
			Content: combined,
		})
	}
	return chapters
}

// splitByObjectivesMarker detecta capítulos via marcador de objetivos no início.
// Livros como McArdle usam "OBJETIVOS DO CAPÍTULO" ou "CHAPTER OBJECTIVES"
// nos primeiros parágrafos de cada capítulo.
func splitByObjectivesMarker(pages []string) []ChapterContent {
	// Marcadores em PT e EN (case-insensitive)
	markerRe := regexp.MustCompile(`(?i)OBJETIVOS\s+DO\s+CAP[ÍI]TULO|CHAPTER\s+OBJECTIVES|LEARNING\s+OBJECTIVES`)

	type chapterStart struct {
		pageIdx   int
		title     string
		chapterNo int
	}
	var starts []chapterStart

	for i, p := range pages {
		head := p
		if len(head) > 600 {
			head = p[:600]
		}
		if !markerRe.MatchString(head) {
			continue
		}

		// Título: primeira linha não-vazia da página atual (se não for o marcador),
		// ou última linha significativa da página anterior
		title := ""
		pageLines := strings.Split(strings.TrimSpace(p), "\n")
		firstMeaningful := ""
		for _, l := range pageLines {
			l = strings.TrimSpace(l)
			if l == "" {
				continue
			}
			if markerRe.MatchString(l) {
				break // chegou no marcador antes de achar título
			}
			if len(l) > 5 {
				firstMeaningful = l
				break
			}
		}

		if firstMeaningful != "" {
			title = firstMeaningful
		} else if i > 0 {
			// Título está na página anterior
			prevLines := strings.Split(strings.TrimSpace(pages[i-1]), "\n")
			for j := len(prevLines) - 1; j >= 0; j-- {
				l := strings.TrimSpace(prevLines[j])
				if len(l) > 5 && len(l) < 300 {
					title = l
					break
				}
			}
		}

		if title == "" {
			title = fmt.Sprintf("Capítulo %d", len(starts)+1)
		}

		starts = append(starts, chapterStart{
			pageIdx:   i,
			title:     title,
			chapterNo: len(starts) + 1,
		})
	}

	if len(starts) < 3 {
		return nil
	}

	var chapters []ChapterContent
	for i, s := range starts {
		endPage := len(pages)
		if i+1 < len(starts) {
			// O capítulo seguinte começa na página do marcador ou na anterior (onde está o título)
			endPage = starts[i+1].pageIdx
			if endPage > 0 && starts[i+1].pageIdx > 0 {
				// Incluir a página de título do próximo capítulo no slice do próximo
				endPage = starts[i+1].pageIdx
			}
		}

		// Incluir página com o título (pode ser i-1 quando título está na prev page)
		startPage := s.pageIdx
		if startPage > 0 {
			prevTitle := strings.TrimSpace(pages[startPage-1])
			if strings.Contains(prevTitle, s.title) {
				startPage = startPage - 1
			}
		}

		chPages := pages[startPage:endPage]
		combined := cleanPDFText(strings.Join(chPages, "\n"))

		// Remover linhas com apenas números de página
		lines := strings.Split(combined, "\n")
		var kept []string
		for _, l := range lines {
			if regexp.MustCompile(`^\s*\d{1,4}\s*$`).MatchString(l) {
				continue
			}
			kept = append(kept, l)
		}
		combined = strings.TrimSpace(strings.Join(kept, "\n"))

		if len(combined) < 500 {
			continue
		}

		chapters = append(chapters, ChapterContent{
			Number:  s.chapterNo,
			Title:   s.title,
			Content: combined,
		})
	}
	return chapters
}

// splitBySectionNumbering detecta capítulos em livros com seções numeradas no formato N.M.
// Quando o número do capítulo (N) muda em linhas isoladas como "1.1", "2.3", "10.2",
// isso indica uma nova fronteira de capítulo.
// Exemplo: "Bioquímica Básica" (Marzzoco) — 22 capítulos com seções 1.1–22.x.
//
// Para evitar falso-positivos do sumário, identifica páginas de TOC pela
// presença de múltiplos capítulos diferentes na mesma página.
func splitBySectionNumbering(pages []string) []ChapterContent {
	// Linha com exatamente "N.M" (número de capítulo.número de seção)
	secLineRe := regexp.MustCompile(`(?m)^(\d{1,2})\.(\d{1,2})$`)

	// Identificar páginas de TOC: página com ≥3 capítulos diferentes = sumário
	isTOCPage := make([]bool, len(pages))
	for i, p := range pages {
		matches := secLineRe.FindAllStringSubmatch(p, -1)
		chapNums := map[int]bool{}
		for _, m := range matches {
			ch := 0
			fmt.Sscanf(m[1], "%d", &ch)
			if ch > 0 {
				chapNums[ch] = true
			}
		}
		if len(chapNums) >= 3 {
			isTOCPage[i] = true
		}
	}

	// Coletar primeira ocorrência de cada capítulo fora de páginas de TOC
	type chStart struct {
		chapterNo int
		pageIdx   int
	}
	var chStarts []chStart
	seenChapter := map[int]bool{}
	lastCh := 0

	for i, p := range pages {
		if isTOCPage[i] {
			continue
		}
		for _, m := range secLineRe.FindAllStringSubmatch(p, -1) {
			ch, sec := 0, 0
			fmt.Sscanf(m[1], "%d", &ch)
			fmt.Sscanf(m[2], "%d", &sec)
			if ch < 1 || ch > 30 || sec < 1 || sec > 30 {
				continue
			}
			if ch != lastCh && !seenChapter[ch] {
				seenChapter[ch] = true
				chStarts = append(chStarts, chStart{chapterNo: ch, pageIdx: i})
				lastCh = ch
			}
		}
	}

	if len(chStarts) < 3 {
		return nil
	}

	// Verificar que os capítulos formam uma sequência razoável (sem saltos grandes)
	for i := 1; i < len(chStarts); i++ {
		if chStarts[i].chapterNo-chStarts[i-1].chapterNo > 5 {
			return nil // muito espaçado — provavelmente falso-positivo
		}
	}

	// Calcular tocSkip baseado na primeira detecção válida (para boundary de startPg)
	firstValidPage := 0
	if len(chStarts) > 0 {
		firstValidPage = chStarts[0].pageIdx
	}

	var chapters []ChapterContent
	for i, cs := range chStarts {
		endPg := len(pages)
		if i+1 < len(chStarts) {
			endPg = chStarts[i+1].pageIdx
		}

		// Incluir uma página antes (pode conter o separador/título do capítulo)
		startPg := cs.pageIdx
		if startPg > firstValidPage {
			startPg--
		}

		chPages := pages[startPg:endPg]
		combined := cleanPDFText(strings.Join(chPages, "\n"))

		// Limpar linhas de número de página isolado
		lines := strings.Split(combined, "\n")
		var kept []string
		for _, l := range lines {
			if regexp.MustCompile(`^\s*\d{1,4}\s*$`).MatchString(l) {
				continue
			}
			kept = append(kept, l)
		}
		combined = strings.TrimSpace(strings.Join(kept, "\n"))

		if len(combined) < 300 {
			continue
		}

		chapters = append(chapters, ChapterContent{
			Number:  cs.chapterNo,
			Title:   fmt.Sprintf("Capítulo %d", cs.chapterNo),
			Content: combined,
		})
	}
	return chapters
}

// splitPDFChapters divide um PDF em capítulos usando estratégias em cascata.
//
// Estratégias (em ordem de preferência):
//  1. Form-feed + running header ("Chapter N" / "CAPÍTULO N" no topo de cada página)
//  2. Form-feed + marcador de objetivos ("OBJETIVOS DO CAPÍTULO" / "CHAPTER OBJECTIVES")
//  3. Claude Haiku extrai TOC do início e localiza títulos no texto
//  4. Regex com padrões de heading comuns em PDFs acadêmicos
//  5. Split por tamanho fixo (50K chars por parte) — último recurso
func (s *ArticleService) splitPDFChapters(filePath string) ([]ChapterContent, error) {
	// Usar pdftotext diretamente para preservar os form feeds (\x0c) entre páginas.
	// A biblioteca Go (ledongthuc/pdf) não preserva separadores de página.
	var pages []string
	var fullText string

	cmd := exec.Command("pdftotext", filePath, "-")
	if output, err := cmd.Output(); err == nil && len(output) > 100 {
		rawText := string(output)
		pages = strings.Split(rawText, "\x0c")
		fullText = cleanPDFText(rawText)
	}

	// Se pdftotext falhou, extrair com a biblioteca Go (sem form feeds)
	if fullText == "" {
		var err error
		fullText, err = s.extractPDFContent(filePath)
		if err != nil {
			return nil, fmt.Errorf("falha ao extrair texto PDF: %w", err)
		}
		fullText = cleanPDFText(fullText)
	}

	// ESTRATÉGIA 1a: Running header em cada página ("Chapter N" / "CAPÍTULO N")
	if len(pages) >= 5 {
		if chapters := splitByRunningHeader(pages); len(chapters) >= 3 {
			fmt.Printf("✅ PDF split por running header: %d capítulos\n", len(chapters))
			return chapters, nil
		}
	}

	// ESTRATÉGIA 1b: Marcador de objetivos por página
	if len(pages) >= 5 {
		if chapters := splitByObjectivesMarker(pages); len(chapters) >= 3 {
			fmt.Printf("✅ PDF split por marcador de objetivos: %d capítulos\n", len(chapters))
			return chapters, nil
		}
	}

	// ESTRATÉGIA 1c: Transição de numeração de seções (N.M como linha isolada)
	// Detecta capítulos em livros acadêmicos com seções numeradas como "1.1", "2.3"
	// quando o número do capítulo (N) muda, detectamos nova fronteira.
	if len(pages) >= 5 {
		if chapters := splitBySectionNumbering(pages); len(chapters) >= 3 {
			fmt.Printf("✅ PDF split por numeração de seções: %d capítulos\n", len(chapters))
			return chapters, nil
		}
	}

	// ESTRATÉGIA 2: Claude Haiku — extrai títulos do sumário e localiza no texto
	if s.aiService != nil && len(fullText) > 1000 {
		firstPart := fullText
		if len(firstPart) > 8000 {
			firstPart = fullText[:8000]
		}
		titles, tocErr := s.aiService.ExtractTableOfContents(firstPart)
		if tocErr == nil && len(titles) >= 3 {
			chapters := s.splitTextByTitles(fullText, titles)
			if len(chapters) >= 3 {
				fmt.Printf("✅ PDF split por TOC Claude: %d capítulos\n", len(chapters))
				return chapters, nil
			}
		}
	}

	// ESTRATÉGIA 3: Regex heading patterns
	chapters := s.splitByHeadingRegex(fullText)
	if len(chapters) >= 3 {
		fmt.Printf("✅ PDF split por regex: %d capítulos\n", len(chapters))
		return chapters, nil
	}

	// ESTRATÉGIA 4: Split por tamanho fixo (50K chars por parte)
	chapters = s.splitByFixedSize(fullText, 50000)
	fmt.Printf("✅ PDF split por tamanho fixo: %d partes\n", len(chapters))
	return chapters, nil
}

// splitTextIntoChapters divide arquivo TXT/MD em capítulos
func (s *ArticleService) splitTextIntoChapters(content string) ([]ChapterContent, error) {
	// Tentar headings Markdown nível 1
	if strings.Contains(content, "\n# ") || strings.HasPrefix(content, "# ") {
		chapters := s.splitByMarkdownHeadings(content)
		if len(chapters) >= 2 {
			return chapters, nil
		}
	}
	// Tentar separadores ---
	if strings.Contains(content, "\n---\n") {
		chapters := s.splitBySeparator(content, "\n---\n")
		if len(chapters) >= 2 {
			return chapters, nil
		}
	}
	// Fallback: tamanho fixo
	return s.splitByFixedSize(content, 50000), nil
}

// splitByMarkdownHeadings divide texto MD por headings # ou ##
func (s *ArticleService) splitByMarkdownHeadings(content string) []ChapterContent {
	headingRe := regexp.MustCompile(`(?m)^#{1,2}\s+(.+)$`)
	matches := headingRe.FindAllStringIndex(content, -1)
	if len(matches) < 2 {
		return nil
	}

	var chapters []ChapterContent
	for i, match := range matches {
		titleMatch := headingRe.FindStringSubmatch(content[match[0]:match[1]])
		if len(titleMatch) < 2 {
			continue
		}
		title := strings.TrimSpace(titleMatch[1])

		start := match[1]
		end := len(content)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}

		text := strings.TrimSpace(content[start:end])
		if len(text) < 200 {
			continue
		}

		chapters = append(chapters, ChapterContent{
			Number:  len(chapters) + 1,
			Title:   title,
			Content: text,
		})
	}
	return chapters
}

// splitBySeparator divide texto por um separador explícito
func (s *ArticleService) splitBySeparator(content, separator string) []ChapterContent {
	parts := strings.Split(content, separator)
	var chapters []ChapterContent
	for i, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) < 200 {
			continue
		}
		lines := strings.SplitN(part, "\n", 2)
		title := fmt.Sprintf("Parte %d", i+1)
		body := part
		if len(lines) > 1 && len(lines[0]) < 200 {
			title = strings.TrimSpace(lines[0])
			body = strings.TrimSpace(lines[1])
		}
		chapters = append(chapters, ChapterContent{
			Number:  len(chapters) + 1,
			Title:   title,
			Content: body,
		})
	}
	return chapters
}

// splitByHeadingRegex divide texto por padrões de heading típicos de PDFs
func (s *ArticleService) splitByHeadingRegex(content string) []ChapterContent {
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`(?im)^(cap[íi]tulo|chapter)\s+(\d+|[IVXLC]+)[.:]\s*(.+)$`),
		regexp.MustCompile(`(?im)^(parte|part)\s+(\d+|[IVXLC]+)[.:]\s*(.+)$`),
		regexp.MustCompile(`(?m)^\d{1,2}\s{2,}[A-ZÁÀÂÃÉÊÍÓÔÕÚ][^\n.]{10,80}$`),
	}

	type matchPos struct {
		start int
		title string
	}
	var allMatches []matchPos

	for _, re := range patterns {
		for _, m := range re.FindAllStringIndex(content, -1) {
			title := strings.TrimSpace(content[m[0]:m[1]])
			allMatches = append(allMatches, matchPos{start: m[0], title: title})
		}
	}

	if len(allMatches) < 3 {
		return nil
	}

	// Ordenar por posição
	for i := 0; i < len(allMatches); i++ {
		for j := i + 1; j < len(allMatches); j++ {
			if allMatches[j].start < allMatches[i].start {
				allMatches[i], allMatches[j] = allMatches[j], allMatches[i]
			}
		}
	}

	var chapters []ChapterContent
	for i, m := range allMatches {
		start := m.start
		end := len(content)
		if i+1 < len(allMatches) {
			end = allMatches[i+1].start
		}
		text := strings.TrimSpace(content[start:end])
		if len(text) < 200 {
			continue
		}
		chapters = append(chapters, ChapterContent{
			Number:  len(chapters) + 1,
			Title:   m.title,
			Content: text,
		})
	}
	return chapters
}

// splitByFixedSize divide texto em partes de tamanho fixo
func (s *ArticleService) splitByFixedSize(content string, chunkSize int) []ChapterContent {
	var chapters []ChapterContent
	total := len(content)
	partNum := 0
	for start := 0; start < total; start += chunkSize {
		end := start + chunkSize
		if end > total {
			end = total
		}
		part := strings.TrimSpace(content[start:end])
		if len(part) < 100 {
			continue
		}
		partNum++
		chapters = append(chapters, ChapterContent{
			Number:  partNum,
			Title:   fmt.Sprintf("Parte %d", partNum),
			Content: part,
		})
	}
	return chapters
}

// splitTextByTitles divide texto usando títulos como âncoras (após extração via Claude)
func (s *ArticleService) splitTextByTitles(content string, titles []string) []ChapterContent {
	type titlePos struct {
		title string
		pos   int
	}

	var positions []titlePos
	contentLower := strings.ToLower(content)

	for _, title := range titles {
		searchTitle := strings.ToLower(strings.TrimSpace(title))
		if len(searchTitle) < 5 {
			continue
		}
		idx := strings.Index(contentLower, searchTitle)
		if idx >= 0 {
			positions = append(positions, titlePos{title: title, pos: idx})
		}
	}

	if len(positions) < 3 {
		return nil
	}

	// Ordenar por posição
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			if positions[j].pos < positions[i].pos {
				positions[i], positions[j] = positions[j], positions[i]
			}
		}
	}

	var chapters []ChapterContent
	for i, pos := range positions {
		start := pos.pos
		end := len(content)
		if i+1 < len(positions) {
			end = positions[i+1].pos
		}
		text := strings.TrimSpace(content[start:end])
		if len(text) < 200 {
			continue
		}
		chapters = append(chapters, ChapterContent{
			Number:  len(chapters) + 1,
			Title:   pos.title,
			Content: text,
		})
	}
	return chapters
}

// GetChaptersByBookID retorna os capítulos de um livro
func (s *ArticleService) GetChaptersByBookID(bookID uuid.UUID) ([]*models.Article, error) {
	return s.repo.GetChaptersByBookID(bookID)
}

// extractPDFContent extrai texto bruto de um PDF usando biblioteca Go + pdftotext como fallback
func (s *ArticleService) extractPDFContent(filePath string) (string, error) {
	var content string
	var extractedWithGo bool

	// TENTATIVA 1: Extrair com biblioteca Go (ledongthuc/pdf)
	file, reader, err := pdf.Open(filePath)
	if err == nil {
		defer file.Close()

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

	// FALLBACK: pdftotext
	if !extractedWithGo || len(content) < 100 {
		if extractedWithGo {
			fmt.Printf("⚠️  Biblioteca Go extraiu apenas %d chars, tentando pdftotext...\n", len(content))
		} else {
			fmt.Println("⚙️  Tentando extração com pdftotext...")
		}

		cmd := exec.Command("pdftotext", filePath, "-")
		output, err := cmd.Output()

		if err == nil && len(output) > 100 {
			content = strings.TrimSpace(string(output))
			fmt.Printf("✅ pdftotext extraiu %d caracteres com sucesso\n", len(content))
		} else {
			if err != nil {
				fmt.Printf("⚠️  pdftotext também falhou: %v\n", err)
			} else {
				fmt.Printf("⚠️  pdftotext extraiu apenas %d chars\n", len(output))
			}
			if len(content) < 100 {
				return "", fmt.Errorf("falha ao extrair texto do PDF (Go e pdftotext falharam)")
			}
		}
	}

	return content, nil
}

// extractEPUBContent extrai texto de um arquivo EPUB (ZIP com XHTML)
func (s *ArticleService) extractEPUBContent(filePath string) (string, error) {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		return "", fmt.Errorf("falha ao abrir EPUB: %w", err)
	}
	defer r.Close()

	// Indexar todos os arquivos do ZIP
	fileIndex := map[string]*zip.File{}
	for _, f := range r.File {
		fileIndex[f.Name] = f
	}

	// Passo 1: localizar OPF via META-INF/container.xml
	opfPath := ""
	for _, f := range r.File {
		if strings.HasSuffix(f.Name, "container.xml") {
			rc, err := f.Open()
			if err != nil {
				continue
			}
			data, _ := io.ReadAll(rc)
			rc.Close()
			re := regexp.MustCompile(`full-path="([^"]+)"`)
			if m := re.FindSubmatch(data); len(m) > 1 {
				opfPath = string(m[1])
			}
			break
		}
	}

	// Passo 2: parsear OPF para obter ordem de leitura (spine)
	var contentOrder []string
	if opfPath != "" {
		if opfFile, ok := fileIndex[opfPath]; ok {
			rc, err := opfFile.Open()
			if err == nil {
				data, _ := io.ReadAll(rc)
				rc.Close()

				// Base dir do OPF para resolver paths relativos
				opfDir := ""
				if idx := strings.LastIndex(opfPath, "/"); idx >= 0 {
					opfDir = opfPath[:idx]
				}

				// Mapear id → href do manifest (atributos em qualquer ordem)
				manifestMap := map[string]string{}
				itemRe := regexp.MustCompile(`(?s)<item\s([^>]+)>`)
				idAttrRe := regexp.MustCompile(`\bid="([^"]+)"`)
				hrefAttrRe := regexp.MustCompile(`\bhref="([^"]+)"`)
				for _, m := range itemRe.FindAllSubmatch(data, -1) {
					attrs := string(m[1])
					idM := idAttrRe.FindStringSubmatch(attrs)
					hrefM := hrefAttrRe.FindStringSubmatch(attrs)
					if len(idM) > 1 && len(hrefM) > 1 {
						manifestMap[idM[1]] = hrefM[1]
					}
				}

				// Extrair spine em ordem
				spineRe := regexp.MustCompile(`<itemref\s[^>]*\bidref="([^"]+)"`)
				for _, m := range spineRe.FindAllSubmatch(data, -1) {
					idref := string(m[1])
					if href, ok := manifestMap[idref]; ok {
						if opfDir != "" {
							contentOrder = append(contentOrder, opfDir+"/"+href)
						} else {
							contentOrder = append(contentOrder, href)
						}
					}
				}
			}
		}
	}

	// Passo 3: extrair texto na ordem correta
	var content strings.Builder
	if len(contentOrder) > 0 {
		for _, path := range contentOrder {
			f, ok := fileIndex[path]
			if !ok {
				continue
			}
			rc, err := f.Open()
			if err != nil {
				continue
			}
			data, _ := io.ReadAll(rc)
			rc.Close()
			content.WriteString(stripHTMLTags(string(data)))
			content.WriteString("\n\n")
		}
	} else {
		// Fallback: extrair todos os HTML/XHTML (exceto nav/toc)
		for _, f := range r.File {
			name := strings.ToLower(f.Name)
			if !strings.HasSuffix(name, ".html") && !strings.HasSuffix(name, ".xhtml") && !strings.HasSuffix(name, ".htm") {
				continue
			}
			if strings.Contains(name, "toc") || strings.Contains(name, "nav") {
				continue
			}
			rc, err := f.Open()
			if err != nil {
				continue
			}
			data, _ := io.ReadAll(rc)
			rc.Close()
			content.WriteString(stripHTMLTags(string(data)))
			content.WriteString("\n\n")
		}
	}

	result := strings.TrimSpace(content.String())
	if len(result) < 100 {
		return "", fmt.Errorf("EPUB extraiu apenas %d caracteres (conteúdo insuficiente)", len(result))
	}
	return result, nil
}

// extractTextFileContent lê conteúdo bruto de arquivos TXT ou MD
func (s *ArticleService) extractTextFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("falha ao ler arquivo: %w", err)
	}
	content := strings.TrimSpace(string(data))
	if len(content) < 10 {
		return "", fmt.Errorf("arquivo vazio ou muito pequeno")
	}
	return content, nil
}

// stripHTMLTags remove tags HTML/XML e decodifica entidades básicas
func stripHTMLTags(s string) string {
	tagRe := regexp.MustCompile(`<[^>]+>`)
	text := tagRe.ReplaceAllString(s, " ")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&nbsp;", " ")
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&#39;", "'")
	text = strings.ReplaceAll(text, "&#160;", " ")
	spaceRe := regexp.MustCompile(`[ \t]+`)
	text = spaceRe.ReplaceAllString(text, " ")
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	blankRe := regexp.MustCompile(`\n{3,}`)
	text = blankRe.ReplaceAllString(text, "\n\n")
	return strings.TrimSpace(text)
}

// ExtractFileMetadata extrai metadados de um arquivo (PDF, EPUB, TXT ou MD)
func (s *ArticleService) ExtractFileMetadata(filePath string) (*PDFMetadata, error) {
	metadata := &PDFMetadata{
		Keywords: []string{},
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	var content string
	var err error

	switch ext {
	case ".epub":
		fmt.Println("📖 Extraindo conteúdo de EPUB...")
		content, err = s.extractEPUBContent(filePath)
	case ".txt", ".md":
		fmt.Printf("📄 Lendo arquivo %s...\n", strings.ToUpper(strings.TrimPrefix(ext, ".")))
		content, err = s.extractTextFileContent(filePath)
	default:
		fmt.Println("📄 Extraindo conteúdo de PDF...")
		content, err = s.extractPDFContent(filePath)
	}

	if err != nil {
		return nil, err
	}

	if len(content) < 100 {
		return nil, fmt.Errorf("conteúdo extraído muito pequeno (%d chars)", len(content))
	}

	metadata.FullContent = content

	// EXTRAÇÃO INTELIGENTE: Usar Claude para extrair metadados da primeira parte do texto
	firstPageText := s.extractFirstPage(content)

	if s.aiService != nil && len(firstPageText) > 500 {
		fmt.Println("🤖 Usando Claude Haiku para extração inteligente de metadados...")

		aiMetadata, err := s.aiService.ExtractArticleMetadata(firstPageText)
		if err != nil {
			fmt.Printf("⚠️  Extração com Claude falhou: %v\n", err)
			fmt.Println("   Usando fallback regex...")
		} else {
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
			// Se Claude extraiu metadados suficientes, pular regex
			if metadata.Title != "" && metadata.Authors != "" {
				return metadata, nil
			}
		}
	}

	// FALLBACK REGEX
	fmt.Println("🔍 Usando extração regex (fallback)...")

	doiRegex := regexp.MustCompile(`(?i)doi[:\s]*([10]\.\d{4,}/[^\s]+)`)
	if matches := doiRegex.FindStringSubmatch(content); len(matches) > 1 {
		metadata.DOI = strings.TrimSpace(matches[1])
	}

	pmidRegex := regexp.MustCompile(`(?i)pmid[:\s]*(\d{7,})`)
	if matches := pmidRegex.FindStringSubmatch(content); len(matches) > 1 {
		metadata.PMID = strings.TrimSpace(matches[1])
	}

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

// ExtractPDFMetadata mantém compatibilidade com scripts existentes
func (s *ArticleService) ExtractPDFMetadata(filePath string) (*PDFMetadata, error) {
	return s.ExtractFileMetadata(filePath)
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
