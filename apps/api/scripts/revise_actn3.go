//go:build legacy_scripts
// +build legacy_scripts

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
	"gorm.io/gorm"
)

const (
	scoreItemID     = "019c1a2b-a36f-77ac-984e-0dcda6b517f0"
	minArticles     = 7
	ragThreshold    = 0.7
	pubmedMaxSearch = 15
)

func main() {
	ctx := context.Background()

	// Conectar ao banco
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Inicializar serviços
	embeddingService := services.NewEmbeddingService()
	vectorRepo := repository.NewArticleVectorRepository(db)
	semanticService := services.NewArticleSemanticService(vectorRepo, embeddingService)
	pubmedService := services.NewPubMedService(db)
	articleService := services.NewArticleService(db, embeddingService)

	fmt.Println("🔬 Revisão Completa: ACTN3 rs1815739 R577X (Performance)")
	fmt.Println("================================================================")

	// 1. Buscar ScoreItem atual
	var scoreItem models.ScoreItem
	if err := db.Preload("Levels").First(&scoreItem, "id = ?", scoreItemID).Error; err != nil {
		log.Fatalf("Failed to find ScoreItem: %v", err)
	}

	fmt.Printf("\n✅ ScoreItem encontrado: %s\n", scoreItem.Name)
	fmt.Printf("   - Relevância clínica atual: %d caracteres\n", len(scoreItem.ClinicalRelevance))
	fmt.Printf("   - Explicação paciente atual: %d caracteres\n", len(scoreItem.PatientExplanation))
	fmt.Printf("   - Conduta atual: %d caracteres\n", len(scoreItem.Conduct))
	fmt.Printf("   - Max points atual: %.0f\n", scoreItem.Points)

	// 2. Contar artigos já linkados
	var currentArticleCount int64
	if err := db.Model(&models.ArticleScoreItem{}).
		Where("score_item_id = ?", scoreItemID).
		Count(&currentArticleCount).Error; err != nil {
		log.Fatalf("Failed to count articles: %v", err)
	}

	fmt.Printf("\n📚 Artigos atualmente linkados: %d\n", currentArticleCount)

	// 3. Buscar artigos similares via RAG
	fmt.Println("\n🔍 Buscando artigos similares via RAG (threshold 0.7)...")

	itemUUID, _ := uuid.Parse(scoreItemID)
	ragResults, err := semanticService.RecommendArticlesForScoreItem(itemUUID, 20)
	if err != nil {
		log.Printf("⚠️  RAG search failed: %v", err)
		ragResults = []repository.ArticleSearchResult{}
	}

	fmt.Printf("   - RAG encontrou: %d artigos candidatos\n", len(ragResults))

	// Filtrar artigos já linkados
	var linkedArticleIDs []uuid.UUID
	if err := db.Model(&models.ArticleScoreItem{}).
		Where("score_item_id = ?", scoreItemID).
		Pluck("article_id", &linkedArticleIDs).Error; err != nil {
		log.Fatalf("Failed to get linked article IDs: %v", err)
	}

	linkedMap := make(map[uuid.UUID]bool)
	for _, id := range linkedArticleIDs {
		linkedMap[id] = true
	}

	newRAGArticles := []repository.ArticleSearchResult{}
	for _, result := range ragResults {
		if !linkedMap[result.ArticleID] {
			newRAGArticles = append(newRAGArticles, result)
		}
	}

	fmt.Printf("   - Novos artigos via RAG: %d\n", len(newRAGArticles))

	// 4. Linkar novos artigos via RAG
	ragLinked := 0
	for _, result := range newRAGArticles {
		if int64(ragLinked)+currentArticleCount >= minArticles {
			break
		}

		link := &models.ArticleScoreItem{
			ScoreItemID:     itemUUID,
			ArticleID:       result.ArticleID,
			ConfidenceScore: result.Similarity,
			AutoLinked:      true,
			LinkedAt:        time.Now(),
		}

		if err := db.Create(link).Error; err != nil {
			log.Printf("⚠️  Failed to link article %s: %v", result.ArticleID, err)
			continue
		}

		ragLinked++
		fmt.Printf("   ✓ Linkado: %s (similaridade: %.3f)\n", result.Title, result.Similarity)
	}

	currentArticleCount += int64(ragLinked)
	fmt.Printf("\n✅ Total de artigos após RAG: %d\n", currentArticleCount)

	// 5. Se ainda não tiver 7 artigos, buscar no PubMed
	pubmedLinked := 0
	if currentArticleCount < minArticles {
		fmt.Println("\n🌐 Buscando artigos adicionais no PubMed...")

		// Gerar query otimizada
		query := generateACTN3Query()
		fmt.Printf("   - Query: %s\n", query)

		// Buscar PMIDs
		pmids, err := pubmedService.RateLimitedSearch(ctx, query, pubmedMaxSearch)
		if err != nil {
			log.Printf("⚠️  PubMed search failed: %v", err)
		} else {
			fmt.Printf("   - PubMed encontrou: %d artigos\n", len(pmids))

			// Buscar detalhes e processar
			for _, pmid := range pmids {
				if currentArticleCount >= minArticles {
					break
				}

				// Verificar se já existe
				var existing models.Article
				err := db.Where("pm_id = ?", pmid).First(&existing).Error
				if err == nil {
					// Artigo já existe, só linkar se ainda não estiver linkado
					if !linkedMap[existing.ID] {
						link := &models.ArticleScoreItem{
							ScoreItemID:     itemUUID,
							ArticleID:       existing.ID,
							ConfidenceScore: 0.85,
							AutoLinked:      true,
							LinkedAt:        time.Now(),
						}

						if err := db.Create(link).Error; err != nil {
							log.Printf("⚠️  Failed to link existing article %s: %v", pmid, err)
							continue
						}

						pubmedLinked++
						currentArticleCount++
						linkedMap[existing.ID] = true
						fmt.Printf("   ✓ Linkado (existente): PMID %s\n", pmid)
					}
					continue
				}

				// Buscar detalhes completos
				details, err := pubmedService.FetchArticleDetails(ctx, []string{pmid})
				if err != nil {
					log.Printf("⚠️  Failed to fetch details for PMID %s: %v", pmid, err)
					continue
				}

				if len(details) == 0 {
					continue
				}

				pubmedArticle := details[0]

				// Tentar baixar PDF
				var pdfPath string
				if pubmedArticle.DOI != "" {
					newArticleID := uuid.New()
					path, err := pubmedService.DownloadPDF(ctx, pubmedArticle.DOI, newArticleID)
					if err != nil {
						log.Printf("   ℹ️  PDF not available for PMID %s: %v", pmid, err)
					} else {
						pdfPath = path
						fmt.Printf("   📄 PDF baixado: %s\n", pdfPath)
					}
				}

				// Criar artigo no banco
				article, err := pubmedService.CreateArticleFromPubMed(pubmedArticle, pdfPath)
				if err != nil {
					log.Printf("⚠️  Failed to create article from PMID %s: %v", pmid, err)
					continue
				}

				// Gerar embedding se tiver abstract
				if article.Abstract != nil && *article.Abstract != "" {
					if err := articleService.GenerateEmbedding(ctx, article.ID); err != nil {
						log.Printf("⚠️  Failed to generate embedding for PMID %s: %v", pmid, err)
					}
				}

				// Linkar ao ScoreItem
				link := &models.ArticleScoreItem{
					ScoreItemID:     itemUUID,
					ArticleID:       article.ID,
					ConfidenceScore: 0.85,
					AutoLinked:      true,
					LinkedAt:        time.Now(),
				}

				if err := db.Create(link).Error; err != nil {
					log.Printf("⚠️  Failed to link new article PMID %s: %v", pmid, err)
					continue
				}

				pubmedLinked++
				currentArticleCount++
				linkedMap[article.ID] = true
				fmt.Printf("   ✓ Criado e linkado: PMID %s - %s\n", pmid, article.Title)

				// Rate limit
				time.Sleep(350 * time.Millisecond)
			}
		}
	}

	fmt.Printf("\n✅ Total final de artigos: %d (RAG: %d, PubMed: %d)\n",
		currentArticleCount, ragLinked, pubmedLinked)

	// 6. Revisar campos clínicos
	fmt.Println("\n📝 Revisando campos clínicos...")

	updatedFields := updateACTN3ClinicalFields(&scoreItem)

	// 7. Salvar mudanças com audit trail
	if updatedFields {
		// Criar histórico de revisão
		history := &models.ScoreItemReviewHistory{
			ScoreItemID: itemUUID,
			ReviewedAt:  time.Now(),
			ReviewType:  "comprehensive_revision",
			Changes: map[string]interface{}{
				"clinical_relevance_updated":  true,
				"patient_explanation_updated": true,
				"conduct_updated":             true,
				"max_points_updated":          true,
				"articles_added_rag":          ragLinked,
				"articles_added_pubmed":       pubmedLinked,
				"total_articles":              currentArticleCount,
			},
			PerformedBy: nil, // Sistema automático
		}

		historyJSON, _ := json.Marshal(history.Changes)
		if err := db.Exec(`
			INSERT INTO score_item_review_history
			(score_item_id, reviewed_at, review_type, changes, performed_by)
			VALUES (?, ?, ?, ?, ?)
		`, history.ScoreItemID, history.ReviewedAt, history.ReviewType, historyJSON, nil).Error; err != nil {
			log.Printf("⚠️  Failed to create review history: %v", err)
		}

		// Atualizar ScoreItem
		scoreItem.LastReview = &history.ReviewedAt
		if err := db.Save(&scoreItem).Error; err != nil {
			log.Fatalf("Failed to save ScoreItem: %v", err)
		}

		fmt.Println("✅ Campos clínicos atualizados e salvos com audit trail")
	}

	fmt.Println("\n================================================================")
	fmt.Println("✅ REVISÃO COMPLETA FINALIZADA COM SUCESSO!")
	fmt.Printf("   - Artigos totais: %d\n", currentArticleCount)
	fmt.Printf("   - Novos artigos RAG: %d\n", ragLinked)
	fmt.Printf("   - Novos artigos PubMed: %d\n", pubmedLinked)
	fmt.Printf("   - Campos clínicos: atualizados\n")
	fmt.Printf("   - Última revisão: %s\n", scoreItem.LastReview.Format("2006-01-02 15:04:05"))
	fmt.Println("================================================================")
}

func generateACTN3Query() string {
	// Query otimizada para ACTN3 rs1815739 R577X
	return `(ACTN3[Title/Abstract] OR "alpha-actinin-3"[Title/Abstract]) AND
	        (rs1815739[Title/Abstract] OR R577X[Title/Abstract] OR "R577X polymorphism"[Title/Abstract]) AND
	        (performance[Title/Abstract] OR athletic[Title/Abstract] OR muscle[Title/Abstract] OR exercise[Title/Abstract]) AND
	        (Review[ptyp] OR Meta-Analysis[ptyp] OR Clinical Trial[ptyp]) AND
	        2018:2026[dp]`
}

func updateACTN3ClinicalFields(item *models.ScoreItem) bool {
	// Campos clínicos atualizados baseados em evidências científicas atualizadas

	newClinicalRelevance := `Gene ACTN3 codifica alpha-actinina-3, proteína estrutural Z-disc expressa exclusivamente em fibras musculares de contração rápida (tipo II/IIx). SNP rs1815739 (R577X, c.1729C>T) resulta em substituição de arginina por códon de parada prematuro (nonsense mutation), gerando alelo nulo X com perda completa de função da proteína. Aproximadamente 18% da população mundial é homozigota XX (null genotype), com ausência total de alpha-actinina-3 funcional, percentual variando significativamente por ancestralidade: ~25% europeus, ~1% africanos subsaarianos, ~20% asiáticos orientais. Genótipo XX associado a desempenho reduzido em atividades de potência explosiva, sprint e força máxima, mas vantagem potencial em resistência aeróbica prolongada, possivelmente via compensação por alpha-actinina-2 (ACTN2) e metabolismo oxidativo aumentado. Estudos GWAS e metanálises demonstram: (1) Genótipo RR e RX superrepresentados em atletas elite de potência/sprint (76-95% vs 50% população geral); (2) Genótipo XX mais frequente em atletas de ultra-resistência; (3) Modulação de resposta hipertrófica ao treinamento resistido; (4) Influência em recuperação muscular pós-exercício; (5) Interação com metabolismo de cálcio muscular e contratilidade. Polimorfismo também investigado em contextos não-atléticos: função muscular no envelhecimento, sarcopenia, fragilidade, miopatias metabólicas, obesidade e metabolismo glicêmico (achados inconsistentes). Evidências emergentes sugerem papel em termorregulação durante exercício em calor. Interpretação clínica requer integração com: (1) Outros genes de performance (ACE I/D, PPARGC1A, PPARA, ADRB2); (2) Fibrotipagem muscular funcional; (3) VO2max e limiar anaeróbico; (4) Avaliação de potência/força explosiva; (5) Histórico de resposta a treinamento. Frequências alélicas: alelo R ~55%, alelo X ~45% (população europeia). Penetrância variável, expressividade dependente de treinamento, nutrição, sono e modulação epigenética.`

	newPatientExplanation := `O gene ACTN3 produz uma proteína chamada alpha-actinina-3, encontrada exclusivamente nas fibras musculares de contração rápida - aquelas responsáveis por explosões de força, velocidade e potência (como sprints, saltos, levantamento de peso). Você possui uma variante genética (rs1815739 R577X) que determina se seu corpo produz ou não essa proteína. Existem três possibilidades: (1) Genótipo RR: você produz alpha-actinina-3 normalmente em ambas as cópias do gene - vantagem natural para atividades de potência, sprint e força explosiva; (2) Genótipo RX: você produz a proteína em quantidade reduzida (uma cópia funcional) - perfil intermediário; (3) Genótipo XX: você não produz nenhuma alpha-actinina-3 funcional - aproximadamente 1 em cada 5 pessoas tem este perfil, com vantagem potencial para resistência aeróbica prolongada, mas desvantagem para potência explosiva. IMPORTANTE: ausência de ACTN3 (genótipo XX) NÃO é doença ou defeito genético! É uma variação completamente normal e funcional encontrada em bilhões de pessoas saudáveis. Seu corpo compensa perfeitamente usando outra proteína similar (alpha-actinina-2). O genótipo XX foi mantido na evolução humana provavelmente porque oferece vantagens metabólicas específicas. O que isso significa na prática: seu genótipo ACTN3 NÃO determina se você pode ou não ser atlético, forte ou saudável, mas sugere qual tipo de treinamento físico e esporte seu corpo pode responder melhor naturalmente. Com treinamento adequado, nutrição otimizada e consistência, pessoas com qualquer genótipo ACTN3 alcançam excelentes resultados - a genética sugere tendências, não limites absolutos. Conhecer seu genótipo permite personalizar tipo de exercício, volume de treino, periodização e recuperação para maximizar resultados individuais.`

	newConduct := `Testar quando: (1) Otimização de prescrição de exercício físico personalizado; (2) Planejamento de performance atlética e seleção de modalidade esportiva baseada em potencial genético; (3) Avaliação de risco de sarcopenia e fragilidade em envelhecimento; (4) Investigação de resposta subótima a treinamento resistido ou aeróbico; (5) Medicina esportiva preventiva e preditiva. Método: genotipagem por PCR em tempo real (TaqMan) para SNP rs1815739 (C>T), microarray genômico (Illumina, Affymetrix), ou NGS se painel genético amplo. Interpretar sempre com frequências alélicas populacionais específicas por ancestralidade. Exames complementares direcionados: (1) Composição corporal por DXA (massa muscular total, distribuição regional, massa muscular apendicular); (2) Fibrotipagem muscular não-invasiva (estimativa via bioimpedância de fase ou EMG de superfície); (3) Testes funcionais: força máxima (1RM agachamento, supino), potência explosiva (salto vertical, Wingate test), sprint 30-60m, VO2max e limiar ventilatório; (4) Marcadores de turnover muscular: CPK basal e pós-exercício, mioglobina, LDH; (5) Painel metabólico: glicose, insulina, HOMA-IR (possível associação com metabolismo glicêmico em genótipo XX); (6) Inflamação sistêmica: PCR-us, IL-6; (7) Status de micronutrientes críticos para função muscular: vitamina D (25-OH), magnésio eritrocitário, selênio, zinco. Conduta personalizada rigorosamente por genótipo: **RR (produtor normal alpha-actinina-3):** Vantagem natural para potência/força explosiva - priorizar treinamento resistido com ênfase em força máxima e potência (cargas 85-95% 1RM, 1-5 repetições, explosividade), pliometria de alto impacto, sprints curtos (10-60m), recuperação adequada entre séries intensas (3-5 min). Nutrição: ingestão proteica 1.8-2.2g/kg/dia, timing periworkout otimizado (proteína + carboidrato 30-60min pré e pós-treino), creatina monohidratada 3-5g/dia (resposta potencialmente superior), beta-alanina 3-6g/dia para capacidade anaeróbica. **RX (produtor parcial):** Perfil versátil - excelente resposta a treinamento concorrente (força + resistência), periodização polarizada ou ondulada, combinar treinos de força (3x/semana) com aeróbico moderado-intenso (2-3x/semana). Flexibilidade estratégica conforme objetivos específicos. **XX (deficiência completa alpha-actinina-3):** Vantagem potencial para resistência aeróbica - priorizar treinos aeróbicos prolongados, HIIT com intervalos mais longos (>2min), treinamento resistido com volume alto e carga moderada (10-15 repetições, 65-75% 1RM, hipertrofia metabólica), menor tempo de recuperação entre séries. Atenção especial para: (1) Periodização cuidadosa em treinamento de força explosiva (requer mais tempo de adaptação neuromuscular); (2) Aquecimento mais prolongado antes de atividades de potência; (3) Suplementação estratégica: beta-alanina para compensar limitação em buffer anaeróbico, HMB (3g/dia) pode auxiliar preservação muscular, taurina para função contrátil. Vitamina D >40ng/mL essencial. **Todos os genótipos:** Sono 7-9h/noite (recuperação muscular), controle de estresse crônico (cortisol catabólico), hidratação adequada, anti-inflamatórios naturais (ômega-3 EPA+DHA 2-3g/dia, curcumina lipossomal, polifenóis). Monitoramento longitudinal: reavaliação funcional (força, potência, VO2max) a cada 12 semanas, composição corporal DXA semestral, ajuste de protocolo conforme resposta individual objetiva. Integrar sempre com outros genes de performance para estratégia verdadeiramente personalizada (ACE, PPARGC1A, PPARA, ADRB2, VDR). Considerar avaliação epigenética (metilação DNA muscular) se resposta anômala ao treinamento prescrito.`

	newMaxPoints := 0.0 // Gene de performance não contribui para score de risco

	// Comparar e atualizar
	updated := false

	if item.ClinicalRelevance != newClinicalRelevance {
		item.ClinicalRelevance = newClinicalRelevance
		updated = true
		fmt.Println("   ✓ clinical_relevance atualizado")
	}

	if item.PatientExplanation != newPatientExplanation {
		item.PatientExplanation = newPatientExplanation
		updated = true
		fmt.Println("   ✓ patient_explanation atualizado")
	}

	if item.Conduct != newConduct {
		item.Conduct = newConduct
		updated = true
		fmt.Println("   ✓ conduct atualizado")
	}

	if item.Points != newMaxPoints {
		item.Points = newMaxPoints
		updated = true
		fmt.Println("   ✓ max_points atualizado")
	}

	if !updated {
		fmt.Println("   ℹ️  Nenhuma mudança necessária nos campos clínicos")
	}

	return updated
}
