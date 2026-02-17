package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/services"
)

// Script para buscar artigos científicos no PubMed sobre adesão terapêutica
// e linkar ao ScoreItem de Adesão
func main() {
	scoreItemID := uuid.MustParse("c77cedd3-2800-759a-bdea-45cd69d48dad")

	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Conectar ao banco
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	db := database.DB
	ctx := context.Background()

	log.Println("🔍 Searching PubMed for therapeutic adherence articles...")

	pubmedService := services.NewPubMedService(db)

	// Queries específicas para adesão terapêutica
	queries := []string{
		"(medication adherence[MeSH] OR treatment adherence[MeSH]) AND (health coaching[Title/Abstract]) AND (Review[ptyp] OR Meta-Analysis[ptyp]) AND 2018:2024[dp]",
		"(patient compliance[MeSH]) AND (motivational interviewing[Title/Abstract]) AND (Review[ptyp]) AND 2018:2024[dp]",
		"(behavioral change[MeSH] OR behavior therapy[MeSH]) AND (self-efficacy[Title/Abstract]) AND (chronic disease[MeSH]) AND 2019:2024[dp]",
		"(health behavior[MeSH]) AND (adherence intervention[Title/Abstract]) AND (personalized medicine[Title/Abstract]) AND 2020:2024[dp]",
	}

	totalFound := 0
	totalLinked := 0

	for i, query := range queries {
		log.Printf("\n📋 Query %d/%d: %s", i+1, len(queries), query)

		// Buscar PMIDs
		pmids, err := pubmedService.RateLimitedSearch(ctx, query, 10)
		if err != nil {
			log.Printf("⚠️  Search failed: %v", err)
			continue
		}

		log.Printf("   Found %d articles", len(pmids))
		totalFound += len(pmids)

		if len(pmids) == 0 {
			continue
		}

		// Buscar metadata (simplificado - parser XML completo seria implementado)
		for _, pmid := range pmids {
			// Verificar se já existe
			var count int64
			db.Table("articles").Where("pm_id = ?", pmid).Count(&count)
			if count > 0 {
				log.Printf("   ⏭️  PMID %s already exists, linking...", pmid)

				// Buscar ID do artigo
				var existingArticle struct {
					ID uuid.UUID
				}
				db.Table("articles").Select("id").Where("pm_id = ?", pmid).First(&existingArticle)

				// Linkar ao ScoreItem
				err = db.Exec(`
					INSERT INTO article_score_items (article_id, score_item_id, confidence_score, auto_linked)
					VALUES (?, ?, 0.85, true)
					ON CONFLICT (score_item_id, article_id) DO NOTHING
				`, existingArticle.ID, scoreItemID).Error

				if err != nil {
					log.Printf("⚠️  Failed to link article: %v", err)
				} else {
					totalLinked++
				}
				continue
			}

			// Criar artigo básico (metadata completa seria extraída do XML)
			articleID := uuid.Must(uuid.NewV7())
			publishDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC) // Default - seria extraído do XML

			err = db.Exec(`
				INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, embedding_status)
				VALUES (?, ?, ?, ?, ?, ?, 'pending')
			`, articleID, fmt.Sprintf("Article PMID: %s (metadata to be fetched)", pmid), "Authors TBD", "Journal TBD", publishDate, pmid).Error

			if err != nil {
				log.Printf("⚠️  Failed to create article for PMID %s: %v", pmid, err)
				continue
			}

			// Linkar ao ScoreItem
			err = db.Exec(`
				INSERT INTO article_score_items (article_id, score_item_id, confidence_score, auto_linked)
				VALUES (?, ?, 0.85, true)
				ON CONFLICT (score_item_id, article_id) DO NOTHING
			`, articleID, scoreItemID).Error

			if err != nil {
				log.Printf("⚠️  Failed to link article: %v", err)
			} else {
				log.Printf("   ✅ Created and linked PMID: %s", pmid)
				totalLinked++
			}

			// Rate limiting
			time.Sleep(100 * time.Millisecond)
		}
	}

	log.Printf("\n\n📊 Summary:")
	log.Printf("   Total articles found: %d", totalFound)
	log.Printf("   Total articles linked: %d", totalLinked)

	// Contar total agora linkado
	var totalNow int64
	db.Table("article_score_items").Where("score_item_id = ?", scoreItemID).Count(&totalNow)
	log.Printf("   Total articles linked to ScoreItem: %d", totalNow)

	log.Println("\n✅ PubMed enrichment completed!")
}
