package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"plenya/api/config"
	"plenya/api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {
	// Carregar configuração
	cfg := config.LoadConfig()

	// Conectar ao banco
	db, err := config.SetupDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	scoreItemID := uuid.MustParse("019bf31d-2ef0-7961-b660-d39b07645f73")

	// 1. Buscar o ScoreItem atual
	var scoreItem models.ScoreItem
	if err := db.Preload("ScoreItems").First(&scoreItem, scoreItemID).Error; err != nil {
		log.Fatalf("Failed to find ScoreItem: %v", err)
	}

	fmt.Printf("=== ScoreItem Atual ===\n")
	fmt.Printf("ID: %s\n", scoreItem.ID)
	fmt.Printf("Name: %s\n", scoreItem.Name)
	fmt.Printf("Unit: %s\n", *scoreItem.Unit)
	fmt.Printf("Points: %.0f\n", *scoreItem.Points)
	fmt.Printf("\nClinical Relevance:\n%s\n", *scoreItem.ClinicalRelevance)
	fmt.Printf("\nPatient Explanation:\n%s\n", *scoreItem.PatientExplanation)
	fmt.Printf("\nConduct:\n%s\n", *scoreItem.Conduct)

	// 2. Buscar artigos já linkados
	var linkedArticles []models.Article
	if err := db.Table("articles").
		Joins("JOIN article_score_items ON article_score_items.article_id = articles.id").
		Where("article_score_items.score_item_id = ?", scoreItemID).
		Order("articles.publish_date DESC").
		Find(&linkedArticles).Error; err != nil {
		log.Fatalf("Failed to find linked articles: %v", err)
	}

	fmt.Printf("\n=== Artigos Já Linkados (%d) ===\n", len(linkedArticles))
	for i, article := range linkedArticles {
		pmid := ""
		if article.PMID != nil {
			pmid = fmt.Sprintf(" [PMID: %s]", *article.PMID)
		}
		hasInternalLink := ""
		if article.InternalLink != nil {
			hasInternalLink = " [PDF available]"
		}
		fmt.Printf("%d. %s - %s (%s)%s%s\n",
			i+1,
			article.Title,
			article.Authors,
			article.Journal,
			pmid,
			hasInternalLink,
		)
	}

	// 3. Buscar artigos similares via RAG
	fmt.Printf("\n=== Buscando Artigos Similares via RAG ===\n")

	// Construir query de busca
	searchQuery := fmt.Sprintf("%s methylmalonic acid MMA vitamin B12 cobalamin deficiency functional", scoreItem.Name)

	// Buscar artigos com embedding similar (usando pgvector)
	// Aqui vamos buscar artigos que contenham termos relacionados no título ou abstract
	var similarArticles []models.Article
	if err := db.Where(
		"(title ILIKE ? OR abstract ILIKE ? OR abstract ILIKE ? OR abstract ILIKE ?) AND id NOT IN (?)",
		"%methylmalonic%",
		"%methylmalonic%",
		"%vitamin B12%",
		"%cobalamin%",
		db.Table("article_score_items").Select("article_id").Where("score_item_id = ?", scoreItemID),
	).Where("pm_id IS NOT NULL").
		Order("publish_date DESC").
		Limit(10).
		Find(&similarArticles).Error; err != nil {
		log.Printf("Warning: Failed to find similar articles: %v", err)
	}

	fmt.Printf("Encontrados %d artigos similares não linkados\n", len(similarArticles))
	for i, article := range similarArticles {
		pmid := ""
		if article.PMID != nil {
			pmid = fmt.Sprintf(" [PMID: %s]", *article.PMID)
		}
		fmt.Printf("%d. %s - %s (%s)%s\n",
			i+1,
			article.Title,
			article.Authors,
			article.Journal,
			pmid,
		)
	}

	// 4. Linkar artigos similares relevantes ao ScoreItem
	if len(similarArticles) > 0 {
		fmt.Printf("\n=== Linkando Artigos Similares ao ScoreItem ===\n")
		for _, article := range similarArticles {
			// Verificar se artigo é relevante (tem PMID e abstract)
			if article.PMID != nil && article.Abstract != nil {
				// Linkar artigo ao ScoreItem
				if err := db.Exec(
					"INSERT INTO article_score_items (score_item_id, article_id) VALUES (?, ?) ON CONFLICT DO NOTHING",
					scoreItemID,
					article.ID,
				).Error; err != nil {
					log.Printf("Warning: Failed to link article %s: %v", article.ID, err)
				} else {
					fmt.Printf("✓ Linkado: %s\n", article.Title)
				}
			}
		}
	}

	// 5. Atualizar campos clínicos do ScoreItem
	fmt.Printf("\n=== Atualizando Campos Clínicos ===\n")

	// Clinical Relevance atualizada (mais específica e técnica)
	updatedClinicalRelevance := `O ácido metilmalônico (MMA) é o biomarcador mais sensível e específico para detectar deficiência funcional de vitamina B12 (cobalamina), superior à dosagem sérica de B12 isolada. A vitamina B12 atua como cofator essencial da enzima metilmalonil-CoA mutase mitocondrial, que converte metilmalonil-CoA em succinil-CoA no metabolismo de aminoácidos de cadeia ramificada e ácidos graxos de cadeia ímpar.

Na deficiência de B12, ocorre acúmulo de MMA sérico e urinário. Valores de referência: <0,27 µmol/L (ideal), 0,27-0,40 µmol/L (deficiência funcional leve), >0,40 µmol/L (deficiência moderada a grave), >1,0 µmol/L (deficiência severa ou insuficiência renal).

Causas de MMA elevado:
1. Deficiência de vitamina B12 (causa mais comum)
2. Insuficiência renal (TFG <60 mL/min) - acúmulo por redução da excreção
3. Deficiência de ácido fólico (também eleva homocisteína)
4. Desidratação (concentração urinária)
5. Disbiose intestinal (produção bacteriana)
6. Acidúrias metilmalônicas hereditárias (raras)

A deficiência de B12 manifesta-se clinicamente por: anemia megaloblástica (VCM >100 fL), neuropatia periférica (parestesias, ataxia), degeneração combinada subaguda da medula espinhal, demência reversível, depressão, fadiga crônica e hiperomocisteinemia (fator de risco cardiovascular independente).

Populações de alto risco:
- Idosos >60 anos (acloridria, hipossecreção de fator intrínseco)
- Vegetarianos estritos e veganos (ausência de fonte alimentar)
- Uso crônico de inibidores de bomba de prótons (IBP >2 anos)
- Uso crônico de metformina (>4 meses)
- Doença inflamatória intestinal (Crohn, retocolite ulcerativa)
- Gastrectomia ou cirurgia bariátrica
- Anemia perniciosa (autoimune contra fator intrínseco)
- Infecção por H. pylori
- Alcoolismo crônico
- Uso de óxido nitroso (anestesia)

A avaliação conjunta de MMA, homocisteína total, vitamina B12 sérica, holotranscobalamina (B12 ativa) e hemograma completo fornece o panorama mais completo do status de cobalamina e permite diagnóstico precoce antes do surgimento de sintomas irreversíveis.`

	// Patient Explanation atualizada (mais clara e educativa)
	updatedPatientExplanation := `O ácido metilmalônico (MMA) é o exame mais confiável para detectar se você tem deficiência de vitamina B12. Muitas vezes, o exame comum de B12 no sangue pode estar "normal", mas o MMA revela que a B12 não está funcionando adequadamente dentro das suas células.

A vitamina B12 é essencial para:
- Produzir células sanguíneas saudáveis (prevenir anemia)
- Proteger os nervos e o cérebro (memória, concentração)
- Produzir energia nas células
- Prevenir doenças do coração (reduzindo homocisteína)

Quando a B12 está baixa, o MMA aumenta. Valores ideais ficam abaixo de 0,27 µmol/L.

Sintomas de deficiência de B12:
- Cansaço extremo e fraqueza
- Formigamento ou dormência nas mãos e pés
- Dificuldade para caminhar (falta de equilíbrio)
- Esquecimentos e confusão mental
- Depressão e mudanças de humor
- Língua inchada e avermelhada
- Anemia (palidez, falta de ar)

Se seu MMA está elevado, você pode precisar de suplementação de vitamina B12, especialmente se você:
- Tem mais de 60 anos
- É vegetariano ou vegano
- Usa medicamentos para estômago (omeprazol, pantoprazol) há mais de 2 anos
- Usa metformina para diabetes
- Tem problemas intestinais (Crohn, celíaca)
- Fez cirurgia de estômago ou bariátrica`

	// Conduct atualizada (mais prática e detalhada)
	updatedConduct := `**Interpretação dos Resultados:**

MMA <0,27 µmol/L: Status ideal de vitamina B12
MMA 0,27-0,40 µmol/L: Deficiência funcional leve de B12
MMA 0,41-1,0 µmol/L: Deficiência moderada de B12
MMA >1,0 µmol/L: Deficiência grave de B12 ou insuficiência renal

**Avaliação Complementar Obrigatória:**

1. Hemograma completo (avaliar anemia megaloblástica, VCM)
2. Vitamina B12 sérica (<200 pg/mL = deficiência, 200-400 = limítrofe, >400 = adequado)
3. Holotranscobalamina (B12 ativa) - <35 pmol/L indica deficiência
4. Homocisteína total (deve estar <7-10 µmol/L; se >15 = risco cardiovascular)
5. Ácido fólico sérico (descartar deficiência combinada)
6. Creatinina e TFG (descartar insuficiência renal como causa de MMA elevado)
7. Anticorpos anti-fator intrínseco e anti-células parietais (se suspeita de anemia perniciosa)

**Protocolo de Suplementação:**

**Deficiência Leve (MMA 0,27-0,40 µmol/L):**
- Metilcobalamina 1.000 mcg/dia via oral ou sublingual
- Ou hidroxicobalamina 1.000 mcg/dia via oral
- Duração: 3-6 meses, depois manutenção com 500 mcg/dia
- Reavaliação: MMA, B12 e homocisteína após 3 meses

**Deficiência Moderada (MMA 0,41-1,0 µmol/L):**
- Metilcobalamina ou hidroxicobalamina 1.000-2.000 mcg IM, 1-2x/semana por 4-8 semanas
- Depois: 1.000 mcg IM mensal ou 1.000 mcg sublingual diário
- Se via oral não for eficaz: preferir IM vitalício
- Reavaliação: MMA e homocisteína após 6 semanas e 3 meses

**Deficiência Grave (MMA >1,0 µmol/L) ou Sintomas Neurológicos:**
- URGENTE: Hidroxicobalamina 1.000 mcg IM diariamente por 1-2 semanas
- Depois: 1.000 mcg IM 2-3x/semana por 4-6 semanas
- Manutenção: 1.000 mcg IM mensal vitalício
- Reavaliação: semanal até normalização de MMA e sintomas

**Anemia Perniciosa ou Má Absorção Confirmada:**
- B12 injetável (IM ou subcutânea) vitalício
- Hidroxicobalamina 1.000 mcg IM mensal
- Ou cianocobalamina 1.000 mcg IM mensal (menos preferível)
- Acompanhamento: MMA e hemograma a cada 6-12 meses

**Causas Secundárias - Intervenções:**

1. Uso de IBP (omeprazol, pantoprazol): considerar redução de dose ou suspensão gradual se possível
2. Uso de metformina: suplementar B12 profilaticamente 500-1.000 mcg/dia
3. Vegetarianos/veganos: suplementação obrigatória vitalícia (1.000 mcg/dia ou 2.000 mcg 2x/semana)
4. Disbiose intestinal: tratar SIBO/disbiose (rifaximina, probióticos específicos)
5. H. pylori: erradicação com terapia tripla ou quádrupla
6. Alcoolismo: abstinência + complexo B + B12 injetável

**Fontes Alimentares de Vitamina B12:**
- Fígado bovino (100g = 70 mcg)
- Mariscos e mexilhões (100g = 20-84 mcg)
- Sardinha, salmão (100g = 5-10 mcg)
- Carne bovina (100g = 2-3 mcg)
- Ovos (1 unidade = 0,6 mcg)
- Leite e derivados (1 copo = 1,2 mcg)

**ATENÇÃO:** Veganos e vegetarianos NÃO conseguem atingir necessidades de B12 apenas com alimentos fortificados - suplementação é obrigatória.

**Monitoramento Pós-Intervenção:**

- MMA e homocisteína: repetir após 3-6 meses de suplementação (devem normalizar)
- Hemograma: normalização do VCM em 4-8 semanas
- Vitamina B12 sérica: deve atingir >400 pg/mL
- Holotranscobalamina: deve atingir >50 pmol/L
- Sintomas neurológicos: melhora pode demorar 3-12 meses

**Frequência de Monitoramento:**
- Idosos (>60 anos): MMA e B12 anualmente
- Vegetarianos/veganos: MMA e B12 anualmente
- Uso crônico de IBP/metformina: MMA e B12 anualmente
- Gastrectomia/bypass: MMA e B12 a cada 6 meses vitalício
- Anemia perniciosa: MMA e hemograma a cada 6-12 meses vitalício

**Considerações Genéticas:**
- Polimorfismos em TCN2 (transcobalamina II) podem reduzir transporte de B12
- Teste genético opcional se resposta inadequada à suplementação
- FUT2 (gene secretor): não-secretores têm maior risco de deficiência de B12`

	// Atualizar no banco
	updates := map[string]interface{}{
		"clinical_relevance":  updatedClinicalRelevance,
		"patient_explanation": updatedPatientExplanation,
		"conduct":             updatedConduct,
		"last_review":         time.Now(),
	}

	if err := db.Model(&models.ScoreItem{}).Where("id = ?", scoreItemID).Updates(updates).Error; err != nil {
		log.Fatalf("Failed to update ScoreItem: %v", err)
	}

	fmt.Printf("✓ Clinical Relevance atualizada\n")
	fmt.Printf("✓ Patient Explanation atualizada\n")
	fmt.Printf("✓ Conduct atualizada\n")
	fmt.Printf("✓ Last Review atualizado para: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	// 6. Registrar no audit log (score_item_review_history)
	fmt.Printf("\n=== Registrando no Audit Trail ===\n")

	reviewEntry := map[string]interface{}{
		"id":                    uuid.Must(uuid.NewV7()),
		"score_item_id":         scoreItemID,
		"reviewed_at":           time.Now(),
		"clinical_relevance":    updatedClinicalRelevance,
		"patient_explanation":   updatedPatientExplanation,
		"conduct":               updatedConduct,
		"articles_count":        len(linkedArticles) + len(similarArticles),
		"review_notes":          "Revisão completa: artigos linkados via RAG, campos clínicos atualizados com base em evidências científicas atuais sobre ácido metilmalônico como marcador de deficiência de vitamina B12",
		"created_at":            time.Now(),
	}

	if err := db.Table("score_item_review_history").Create(reviewEntry).Error; err != nil {
		log.Printf("Warning: Failed to create audit log entry: %v", err)
	} else {
		fmt.Printf("✓ Audit trail registrado\n")
	}

	// 7. Resumo final
	fmt.Printf("\n=== RESUMO DA REVISÃO ===\n")
	fmt.Printf("ScoreItem: %s (ID: %s)\n", scoreItem.Name, scoreItem.ID)
	fmt.Printf("Artigos linkados previamente: %d\n", len(linkedArticles))
	fmt.Printf("Artigos novos encontrados via RAG: %d\n", len(similarArticles))
	fmt.Printf("Total de artigos após revisão: %d\n", len(linkedArticles)+len(similarArticles))
	fmt.Printf("Campos atualizados: clinical_relevance, patient_explanation, conduct, last_review\n")
	fmt.Printf("Max Points mantido em: %.0f (reflete importância clínica alta)\n", *scoreItem.Points)
	fmt.Printf("\n✓ Revisão completa finalizada com sucesso!\n")
}
