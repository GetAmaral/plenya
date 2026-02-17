package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

// PubMedService integra com PubMed E-utilities API para buscar artigos científicos
type PubMedService struct {
	db         *gorm.DB
	httpClient *http.Client
	baseURL    string
	apiKey     string // Opcional - aumenta rate limit de 3/s para 10/s
	email      string // Requerido por NCBI
}

// PubMedArticle representa metadata de artigo retornado pelo PubMed
type PubMedArticle struct {
	PMID        string
	DOI         string
	Title       string
	Abstract    string
	Authors     []string
	Journal     string
	Year        int
	Volume      string
	Issue       string
	Pages       string
	PubType     []string
	MeshTerms   []string
	Keywords    []string
}

// NewPubMedService cria nova instância do serviço
func NewPubMedService(db *gorm.DB) *PubMedService {
	return &PubMedService{
		db: db,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: "https://eutils.ncbi.nlm.nih.gov/entrez/eutils",
		apiKey:  os.Getenv("PUBMED_API_KEY"),  // Opcional
		email:   os.Getenv("PUBMED_EMAIL"),     // Requerido
	}
}

// SearchArticles busca artigos no PubMed usando query
// Retorna lista de PMIDs
func (s *PubMedService) SearchArticles(ctx context.Context, query string, maxResults int) ([]string, error) {
	if s.email == "" {
		return nil, fmt.Errorf("PUBMED_EMAIL environment variable is required")
	}

	if maxResults <= 0 || maxResults > 100 {
		maxResults = 10
	}

	// Construir URL de busca
	params := url.Values{}
	params.Set("db", "pubmed")
	params.Set("term", query)
	params.Set("retmax", fmt.Sprintf("%d", maxResults))
	params.Set("retmode", "json")
	params.Set("tool", "plenya")
	params.Set("email", s.email)
	if s.apiKey != "" {
		params.Set("api_key", s.apiKey)
	}

	searchURL := fmt.Sprintf("%s/esearch.fcgi?%s", s.baseURL, params.Encode())

	// Fazer request
	req, err := http.NewRequestWithContext(ctx, "GET", searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("PubMed API returned status %d", resp.StatusCode)
	}

	// Parse resposta JSON
	var result struct {
		ESearchResult struct {
			IDList []string `json:"idlist"`
			Count  string   `json:"count"`
		} `json:"esearchresult"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.ESearchResult.IDList, nil
}

// FetchArticleDetails busca metadata completa de artigos pelos PMIDs
func (s *PubMedService) FetchArticleDetails(ctx context.Context, pmids []string) ([]*PubMedArticle, error) {
	if len(pmids) == 0 {
		return []*PubMedArticle{}, nil
	}

	// Construir URL de fetch
	params := url.Values{}
	params.Set("db", "pubmed")
	params.Set("id", strings.Join(pmids, ","))
	params.Set("retmode", "xml")
	params.Set("tool", "plenya")
	params.Set("email", s.email)
	if s.apiKey != "" {
		params.Set("api_key", s.apiKey)
	}

	fetchURL := fmt.Sprintf("%s/efetch.fcgi?%s", s.baseURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", fetchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute fetch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("PubMed API returned status %d", resp.StatusCode)
	}

	// Parse XML (simplificado - em produção usar encoding/xml)
	// Por ora, retornar estrutura básica
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// TODO: Implementar parser XML completo
	// Por ora, criar artigos com PMID apenas
	articles := make([]*PubMedArticle, 0, len(pmids))
	for _, pmid := range pmids {
		articles = append(articles, &PubMedArticle{
			PMID: pmid,
			// Metadata será extraída do XML em implementação futura
		})
	}

	_ = body // Usar body quando implementar parser XML

	return articles, nil
}

// DownloadPDF tenta baixar PDF de artigo via Unpaywall (open access)
// Retorna caminho do arquivo salvo ou erro
func (s *PubMedService) DownloadPDF(ctx context.Context, doi string, articleID uuid.UUID) (string, error) {
	if doi == "" {
		return "", fmt.Errorf("DOI is required")
	}

	// Tentar Unpaywall primeiro (gratuito, legal)
	unpaywallURL := fmt.Sprintf("https://api.unpaywall.org/v2/%s?email=%s",
		url.QueryEscape(doi),
		url.QueryEscape(s.email))

	req, err := http.NewRequestWithContext(ctx, "GET", unpaywallURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create unpaywall request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to query unpaywall: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unpaywall returned status %d (article may not be open access)", resp.StatusCode)
	}

	// Parse resposta Unpaywall
	var unpaywallResp struct {
		BestOALocation struct {
			URLForPDF string `json:"url_for_pdf"`
		} `json:"best_oa_location"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&unpaywallResp); err != nil {
		return "", fmt.Errorf("failed to decode unpaywall response: %w", err)
	}

	pdfURL := unpaywallResp.BestOALocation.URLForPDF
	if pdfURL == "" {
		return "", fmt.Errorf("no open access PDF available")
	}

	// Baixar PDF
	pdfReq, err := http.NewRequestWithContext(ctx, "GET", pdfURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create PDF request: %w", err)
	}

	pdfResp, err := s.httpClient.Do(pdfReq)
	if err != nil {
		return "", fmt.Errorf("failed to download PDF: %w", err)
	}
	defer pdfResp.Body.Close()

	if pdfResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("PDF download returned status %d", pdfResp.StatusCode)
	}

	// Salvar PDF
	uploadsDir := os.Getenv("UPLOADS_DIR")
	if uploadsDir == "" {
		uploadsDir = "/app/uploads"
	}

	articlesDir := filepath.Join(uploadsDir, "articles")
	if err := os.MkdirAll(articlesDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create articles directory: %w", err)
	}

	filename := fmt.Sprintf("%s.pdf", articleID.String())
	filepath := filepath.Join(articlesDir, filename)

	outFile, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, pdfResp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write PDF: %w", err)
	}

	return filepath, nil
}

// CreateArticleFromPubMed cria registro de Article no banco a partir de metadata PubMed
func (s *PubMedService) CreateArticleFromPubMed(pubmedArticle *PubMedArticle, pdfPath string) (*models.Article, error) {
	// Verificar se artigo já existe (por PMID ou DOI)
	var existing models.Article

	// Build query dynamically to avoid matching empty strings
	query := s.db.Model(&models.Article{})
	if pubmedArticle.PMID != "" {
		query = query.Where("pm_id = ?", pubmedArticle.PMID)
	}
	if pubmedArticle.DOI != "" {
		if pubmedArticle.PMID != "" {
			query = query.Or("doi = ?", pubmedArticle.DOI)
		} else {
			query = query.Where("doi = ?", pubmedArticle.DOI)
		}
	}

	// Only search if at least one identifier is present
	if pubmedArticle.PMID != "" || pubmedArticle.DOI != "" {
		err := query.First(&existing).Error
		if err == nil {
			return &existing, nil // Já existe
		}
		if err != gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("failed to check existing article: %w", err)
		}
	}

	// Criar publish date
	publishDate := time.Date(pubmedArticle.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	// Criar novo artigo
	article := &models.Article{
		Title:           pubmedArticle.Title,
		Authors:         strings.Join(pubmedArticle.Authors, ", "),
		Journal:         pubmedArticle.Journal,
		PublishDate:     publishDate,
		Language:        "en",
		ArticleType:     "research_article", // Default
		Keywords:        pubmedArticle.Keywords,
		MeshTerms:       pubmedArticle.MeshTerms,
		EmbeddingStatus: "pending", // Será processado para embeddings
	}

	// Set optional fields only if non-empty (avoid storing empty strings)
	if pubmedArticle.DOI != "" {
		article.DOI = &pubmedArticle.DOI
	}
	if pubmedArticle.PMID != "" {
		article.PMID = &pubmedArticle.PMID
	}
	if pubmedArticle.Abstract != "" {
		article.Abstract = &pubmedArticle.Abstract
	}
	if pdfPath != "" {
		article.InternalLink = &pdfPath
	}

	if err := s.db.Create(article).Error; err != nil {
		return nil, fmt.Errorf("failed to create article: %w", err)
	}

	return article, nil
}

// GenerateQueryForScoreItem usa LLM para gerar query PubMed otimizada
// Simplificado - em produção usaria Claude Haiku
func (s *PubMedService) GenerateQueryForScoreItem(item *models.ScoreItem) string {
	// Por ora, query simples baseada no nome
	// TODO: Integrar Claude Haiku para queries melhores

	query := item.Name

	// Adicionar MeSH terms comuns
	if item.Unit != nil {
		query += " AND (Reference Values[Mesh])"
	}

	// Filtrar por tipo de publicação
	query += " AND (Review[ptyp] OR Meta-Analysis[ptyp])"

	// Filtrar por data (últimos 6 anos)
	currentYear := time.Now().Year()
	query += fmt.Sprintf(" AND %d:%d[dp]", currentYear-6, currentYear)

	return query
}

// RateLimitedSearch executa busca com rate limiting automático
func (s *PubMedService) RateLimitedSearch(ctx context.Context, query string, maxResults int) ([]string, error) {
	// Rate limit: 3 req/s sem API key, 10 req/s com API key
	delay := 334 * time.Millisecond
	if s.apiKey != "" {
		delay = 100 * time.Millisecond
	}

	time.Sleep(delay)
	return s.SearchArticles(ctx, query, maxResults)
}

// Helper function
func stringPtr(s string) *string {
	return &s
}
