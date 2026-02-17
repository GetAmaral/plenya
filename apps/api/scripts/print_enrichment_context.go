package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/gorm"
)

// EnrichmentContext contexto para enrichment por agent externo (Claude Code)
type EnrichmentContext struct {
	ScoreItem ScoreItemData `json:"score_item"`
	Articles  []ArticleData `json:"articles"`
	Tier      string        `json:"tier"`
}

type ScoreItemData struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Unit                *string `json:"unit"`
	Gender              *string `json:"gender"`
	AgeRangeMin         *int    `json:"age_range_min"`
	AgeRangeMax         *int    `json:"age_range_max"`
	Points              *float64 `json:"points"`
	ClinicalRelevance   *string `json:"clinical_relevance"`
	PatientExplanation  *string `json:"patient_explanation"`
	Conduct             *string `json:"conduct"`
}

type ArticleData struct {
	Title      string  `json:"title"`
	Authors    string  `json:"authors"`
	Journal    string  `json:"journal"`
	Year       int     `json:"year"`
	Abstract   *string `json:"abstract"`
	DOI        *string `json:"doi"`
	PMID       *string `json:"pmid"`
	Similarity float64 `json:"similarity"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run scripts/print_enrichment_context.go <item_number_or_uuid>")
	}

	// Carregar config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Conectar banco
	if err = database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db := database.DB

	// Resolver ScoreItem ID
	scoreItemID, err := resolveScoreItemID(db, os.Args[1])
	if err != nil {
		log.Fatalf("Failed to resolve ScoreItem: %v", err)
	}

	// Buscar ScoreItem
	var item models.ScoreItem
	if err := db.First(&item, scoreItemID).Error; err != nil {
		log.Fatalf("Failed to find ScoreItem: %v", err)
	}

	// Buscar artigos linkados (top 10 por similaridade)
	vectorRepo := repository.NewArticleVectorRepository(db)
	articleResults, err := vectorRepo.FindSimilarArticlesForScoreItem(scoreItemID, 10, 0.7)
	if err != nil {
		log.Printf("Warning: RAG search failed: %v", err)
		articleResults = []repository.ArticleSearchResult{}
	}

	// Determinar tier
	tier := determineTier(&item)

	// Construir contexto
	context := EnrichmentContext{
		ScoreItem: ScoreItemData{
			ID:                  item.ID.String(),
			Name:                item.Name,
			Unit:                item.Unit,
			Gender:              item.Gender,
			AgeRangeMin:         item.AgeRangeMin,
			AgeRangeMax:         item.AgeRangeMax,
			Points:              item.Points,
			ClinicalRelevance:   item.ClinicalRelevance,
			PatientExplanation:  item.PatientExplanation,
			Conduct:             item.Conduct,
		},
		Articles: make([]ArticleData, 0, len(articleResults)),
		Tier:     tier,
	}

	for _, result := range articleResults {
		context.Articles = append(context.Articles, ArticleData{
			Title:      result.Article.Title,
			Authors:    result.Article.Authors,
			Journal:    result.Article.Journal,
			Year:       result.Article.PublishDate.Year(),
			Abstract:   result.Article.Abstract,
			DOI:        result.Article.DOI,
			PMID:       result.Article.PMID,
			Similarity: result.Similarity,
		})
	}

	// Output JSON
	jsonData, err := json.MarshalIndent(context, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}

func resolveScoreItemID(db *gorm.DB, input string) (uuid.UUID, error) {
	if id, err := uuid.Parse(input); err == nil {
		return id, nil
	}

	seqNum, err := strconv.Atoi(input)
	if err != nil {
		return uuid.Nil, fmt.Errorf("input must be a UUID or sequence number")
	}

	var item models.ScoreItem
	offset := seqNum - 1
	if offset < 0 {
		return uuid.Nil, fmt.Errorf("sequence number must be >= 1")
	}

	err = db.Order("created_at ASC").Offset(offset).Limit(1).First(&item).Error
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to find ScoreItem #%d: %w", seqNum, err)
	}

	return item.ID, nil
}

func determineTier(item *models.ScoreItem) string {
	crLen := len(stringValue(item.ClinicalRelevance))
	peLen := len(stringValue(item.PatientExplanation))
	condLen := len(stringValue(item.Conduct))

	// Tier 1: PRESERVAR
	if crLen >= 1500 && peLen >= 600 && condLen >= 800 {
		return "preserve"
	}

	// Tier 3: REESCREVER
	if crLen < 500 {
		return "rewrite"
	}

	// Tier 2: ENRIQUECER
	return "enrich"
}

func stringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
