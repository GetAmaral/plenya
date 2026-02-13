package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PDFSource representa uma fonte de PDF
type PDFSource struct {
	Name string
	URL  string
}

// FetchResult representa o resultado da busca de um artigo
type FetchResult struct {
	ArticleID   string
	Title       string
	Status      string // "found", "not_found", "error"
	Source      string // "pubmed", "unpaywall", "scihub"
	Error       string
	PDFPath     string
	ProcessedAt time.Time
}

func main() {
	fmt.Println("🔍 PDF Fetcher - Buscando PDFs faltantes")
	fmt.Println("=" + strings.Repeat("=", 60))

	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env file not found (using environment variables)")
	}

	// Carregar configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Conectar ao banco
	dsn := cfg.Database.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database\n")

	// Buscar artigos falhados
	var failedArticles []models.Article
	err = db.Raw(`
		SELECT a.*
		FROM articles a
		JOIN embedding_queue eq ON a.id = eq.entity_id
		WHERE eq.status = 'failed'
		  AND a.deleted_at IS NULL
		  AND (a.internal_link IS NULL OR a.internal_link = '')
		ORDER BY a.publish_date DESC
	`).Scan(&failedArticles).Error

	if err != nil {
		log.Fatalf("❌ Failed to fetch failed articles: %v", err)
	}

	fmt.Printf("📚 Found %d articles without PDFs\n\n", len(failedArticles))

	// Estatísticas
	results := []FetchResult{}
	successCount := 0
	notFoundCount := 0
	errorCount := 0

	// Processar cada artigo
	for i, article := range failedArticles {
		fmt.Printf("[%d/%d] Processing: %s\n", i+1, len(failedArticles), truncate(article.Title, 80))

		result := FetchResult{
			ArticleID:   article.ID.String(),
			Title:       article.Title,
			ProcessedAt: time.Now(),
		}

		// Tentar buscar PDF
		pdfURL, source := findPDFURL(article)

		if pdfURL == "" {
			result.Status = "not_found"
			result.Error = "No PDF found in any source"
			notFoundCount++
			fmt.Printf("   ❌ Not found\n\n")
		} else {
			fmt.Printf("   ✅ Found on %s: %s\n", source, truncate(pdfURL, 60))
			result.Source = source

			// Baixar PDF
			pdfPath, err := downloadPDF(pdfURL, article.ID.String())
			if err != nil {
				result.Status = "error"
				result.Error = fmt.Sprintf("Download failed: %v", err)
				errorCount++
				fmt.Printf("   ❌ Download failed: %v\n\n", err)
			} else {
				result.Status = "found"
				result.PDFPath = pdfPath
				successCount++
				fmt.Printf("   📥 Downloaded to: %s\n", pdfPath)

				// Atualizar artigo no banco
				db.Model(&article).Update("internal_link", pdfPath)

				// TODO: Processar PDF via Claude AI e fazer embedding
				// Por ora, apenas marcar como pendente para reprocessamento
				db.Exec("UPDATE embedding_queue SET status = 'pending', retry_count = 0 WHERE entity_id = ?", article.ID)

				fmt.Printf("   ✅ Queued for processing\n\n")
			}
		}

		results = append(results, result)

		// Rate limiting para evitar bloqueios
		time.Sleep(2 * time.Second)
	}

	// Resumo final
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 Summary:")
	fmt.Printf("   ✅ Found and downloaded: %d\n", successCount)
	fmt.Printf("   ❌ Not found: %d\n", notFoundCount)
	fmt.Printf("   ⚠️  Errors: %d\n", errorCount)
	fmt.Println(strings.Repeat("=", 60))

	// Listar artigos não encontrados
	if notFoundCount > 0 {
		fmt.Println("\n📋 Articles not found (manual search needed):")
		for _, result := range results {
			if result.Status == "not_found" {
				fmt.Printf("   - %s\n", result.Title)
			}
		}
	}
}

// findPDFURL tenta encontrar URL do PDF em diferentes fontes
func findPDFURL(article models.Article) (string, string) {
	// 1. Tentar PubMed Central (se tiver PMID)
	if article.PMID != nil && *article.PMID != "" {
		url := tryPubMedCentral(*article.PMID)
		if url != "" {
			return url, "PubMed Central"
		}
	}

	// 2. Tentar Unpaywall (se tiver DOI)
	if article.DOI != nil && *article.DOI != "" {
		url := tryUnpaywall(*article.DOI)
		if url != "" {
			return url, "Unpaywall (Open Access)"
		}
	}

	// 3. Tentar Sci-Hub (último recurso, se tiver DOI)
	if article.DOI != nil && *article.DOI != "" {
		url := trySciHub(*article.DOI)
		if url != "" {
			return url, "Sci-Hub"
		}
	}

	return "", ""
}

// tryPubMedCentral busca PDF no PubMed Central
func tryPubMedCentral(pmid string) string {
	// PMC API: https://www.ncbi.nlm.nih.gov/pmc/tools/oa-service/
	url := fmt.Sprintf("https://www.ncbi.nlm.nih.gov/pmc/articles/PMC%s/pdf/", pmid)

	// Verificar se existe
	resp, err := http.Head(url)
	if err == nil && resp.StatusCode == 200 {
		return url
	}

	// Tentar formato alternativo
	url = fmt.Sprintf("https://europepmc.org/backend/ptpmcrender.fcgi?accid=PMC%s&blobtype=pdf", pmid)
	resp, err = http.Head(url)
	if err == nil && resp.StatusCode == 200 {
		return url
	}

	return ""
}

// tryUnpaywall busca PDF via Unpaywall (Open Access)
func tryUnpaywall(doi string) string {
	// Unpaywall API: https://unpaywall.org/products/api
	// Requer email (usar genérico)
	email := "research@plenya.com"
	url := fmt.Sprintf("https://api.unpaywall.org/v2/%s?email=%s", doi, email)

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ""
	}

	// Parse JSON response (simplificado)
	// TODO: Parse JSON corretamente
	// Por ora, retorna vazio
	return ""
}

// trySciHub busca PDF no Sci-Hub
func trySciHub(doi string) string {
	// Sci-Hub mirrors (atualizar conforme disponibilidade)
	mirrors := []string{
		"https://sci-hub.se/",
		"https://sci-hub.st/",
		"https://sci-hub.ru/",
	}

	for _, mirror := range mirrors {
		url := mirror + doi
		resp, err := http.Head(url)
		if err == nil && resp.StatusCode == 200 {
			return url
		}
	}

	return ""
}

// downloadPDF baixa o PDF da URL
func downloadPDF(url, articleID string) (string, error) {
	// Criar diretório se não existir
	uploadDir := "./uploads/articles"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}

	// Nome do arquivo
	filename := fmt.Sprintf("%s.pdf", articleID)
	filepath := filepath.Join(uploadDir, filename)

	// Download
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	// Salvar arquivo
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return "/uploads/articles/" + filename, nil
}

// truncate trunca string para tamanho máximo
func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
