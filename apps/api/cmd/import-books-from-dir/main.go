package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// BookDef define os metadados de um livro a ser importado
type BookDef struct {
	Title    string
	Authors  string
	Journal  string // editora / publisher
	Language string
	Year     int
	ChapDir  string // caminho relativo a apps/api/examples/
	OrigFile string // arquivo original (PDF/EPUB), relativo a apps/api/examples/
}

var books = []BookDef{
	{
		Title:    "Laboratory Evaluations for Integrative and Functional Medicine",
		Authors:  "J. Alexander Bralley, Richard S. Lord",
		Journal:  "Metametrix Institute",
		Language: "en",
		Year:     2008,
		ChapDir:  "chapters",
		OrigFile: "Funcional Medicine Laboratory.pdf",
	},
	{
		Title:    "Fisiologia do Exercício - McArdle, Katch & Katch (8ª ed.)",
		Authors:  "William D. McArdle, Frank I. Katch, Victor L. Katch",
		Journal:  "Guanabara Koogan",
		Language: "pt",
		Year:     2015,
		ChapDir:  "chapters_mcardle_final",
		OrigFile: "Fisiologia Do Exercício - Mcardle 8ª ed._compressed_compressed.pdf",
	},
	{
		Title:    "Bioquímica Básica",
		Authors:  "Anita Marzzoco, Bayardo Baptista Torres",
		Journal:  "Guanabara Koogan",
		Language: "pt",
		Year:     2015,
		ChapDir:  "chapters_bioquimica",
		OrigFile: "Bioquímica Básica 4Ed Marzzoco.pdf",
	},
	{
		Title:    "Biomechanics and Exercise Physiology",
		Authors:  "Arthur T. Johnson",
		Journal:  "Wiley-Interscience",
		Language: "en",
		Year:     1991,
		ChapDir:  "chapters_biomechanics",
		OrigFile: "Biomechanics and Exercise Physiology -- Arthur T_ Johnson -- 1, 1991 -- A Wiley-Interscience Publication John Wiley & Sons -- 9780471853985 -- b18885ed77ea49b80c0c05ce026af3a2 -- Anna's Archive (1).pdf",
	},
	{
		Title:    "NSCA's Guide to Tests and Assessments",
		Authors:  "Todd A. Miller",
		Journal:  "Human Kinetics",
		Language: "en",
		Year:     2012,
		ChapDir:  "chapters_nsca",
		OrigFile: "pdfcoffee.com_nsca-guide-to-tests-and-assessments-pdf-free.pdf",
	},
	{
		Title:    "The Biomechanics Method for Corrective Exercise",
		Authors:  "Justin Price",
		Journal:  "Human Kinetics",
		Language: "en",
		Year:     2019,
		ChapDir:  "chapters_price",
		OrigFile: "The biomechanics method for corrective exercise -- (Personal trainer) Justin Price -- 1, Champaign, IL, 2019 -- Human Kinetics Human Kinetics -- 9781492545668 -- ba4e4964ccf99a058fd02506a03465cb -- Anna's Archive.pdf",
	},
	{
		Title:    "Biomecânica Básica",
		Authors:  "Susan J. Hall",
		Journal:  "Guanabara Koogan",
		Language: "pt",
		Year:     2016,
		ChapDir:  "chapters_hall",
		OrigFile: "Biomecanica_Basica_-_Susan_J._Hall.pdf",
	},
	{
		Title:    "NSCA's Essentials of Strength Training and Conditioning (4th ed.)",
		Authors:  "G. Gregory Haff, N. Travis Triplett",
		Journal:  "Human Kinetics",
		Language: "en",
		Year:     2016,
		ChapDir:  "chapters_essentials",
		OrigFile: "Essencial_of_Strenght.pdf",
	},
	{
		Title:    "Exercise Physiology for Health, Fitness, and Performance (5th ed.)",
		Authors:  "Sharon A. Plowman, Denise L. Smith",
		Journal:  "Wolters Kluwer",
		Language: "en",
		Year:     2017,
		ChapDir:  "chapters_plowman",
		OrigFile: "vdoc.pub_exercise-physiology-for-health-fitness-and-performance.epub",
	},
}

var chapterNumRe = regexp.MustCompile(`^chapter_(\d+)_`)
var h2TitleRe = regexp.MustCompile(`^##\s+(.+)$`)
var h1TitleRe = regexp.MustCompile(`^#\s+(.+)$`)

// parseChapterMD extrai número e título do arquivo MD de capítulo
func parseChapterMD(path string) (int, string, string, error) {
	filename := filepath.Base(path)

	// Número do capítulo pelo nome do arquivo
	chNum := 0
	if m := chapterNumRe.FindStringSubmatch(filename); len(m) > 1 {
		n, _ := strconv.Atoi(m[1])
		chNum = n
	}

	// Ler conteúdo completo
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, "", "", err
	}
	content := string(data)

	// Extrair título da primeira linha H2 ou H1
	title := ""
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if m := h2TitleRe.FindStringSubmatch(line); len(m) > 1 {
			title = strings.TrimSpace(m[1])
			// Remover prefixo "Capítulo N:" ou "Chapter N:" se presente
			title = regexp.MustCompile(`^(Capítulo|Chapter)\s+\d+[:.]\s*`).ReplaceAllString(title, "")
			break
		}
		if m := h1TitleRe.FindStringSubmatch(line); len(m) > 1 {
			// H1 é o título do livro; continuar procurando H2
			continue
		}
	}

	if title == "" {
		// Fallback: derivar do filename
		base := strings.TrimSuffix(filename, ".md")
		base = chapterNumRe.ReplaceAllString(base, "")
		title = strings.ReplaceAll(base, "_", " ")
	}

	return chNum, title, content, nil
}

// importBook importa um livro e seus capítulos no banco de dados
func importBook(db *gorm.DB, qSvc *services.EmbeddingQueueService, basePath string, book BookDef, userID uuid.UUID, queueEmbeddings bool) error {
	fmt.Printf("\n📚 Importando: %s\n", book.Title)

	chapDir := filepath.Join(basePath, book.ChapDir)
	origFile := filepath.Join(basePath, book.OrigFile)

	// Verificar se já existe
	var existing models.Article
	if err := db.Where("title = ? AND source_type = ?", book.Title, "book").First(&existing).Error; err == nil {
		fmt.Printf("   ⏭️  Já existe (ID: %s) — pulando\n", existing.ID.String()[:8])
		return nil
	}

	// Ler MDs dos capítulos
	entries, err := os.ReadDir(chapDir)
	if err != nil {
		return fmt.Errorf("erro ao ler diretório %s: %w", chapDir, err)
	}

	type chapEntry struct {
		num     int
		title   string
		content string
		path    string
	}
	var chapters []chapEntry

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		path := filepath.Join(chapDir, e.Name())
		num, title, content, err := parseChapterMD(path)
		if err != nil {
			fmt.Printf("   ⚠️  Erro ao ler %s: %v\n", e.Name(), err)
			continue
		}
		chapters = append(chapters, chapEntry{num, title, content, e.Name()})
	}

	// Ordenar por número de capítulo
	sort.Slice(chapters, func(i, j int) bool {
		return chapters[i].num < chapters[j].num
	})

	totalChapters := len(chapters)
	if totalChapters == 0 {
		return fmt.Errorf("nenhum capítulo encontrado em %s", chapDir)
	}

	// Montar InternalLink para o arquivo original
	var internalLink *string
	if _, statErr := os.Stat(origFile); statErr == nil {
		link := origFile
		internalLink = &link
	}

	// Publish date
	publishDate := time.Date(book.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	// ── Criar Article do livro (pai) ──
	bookArticle := &models.Article{
		Title:         book.Title,
		Authors:       book.Authors,
		Journal:       book.Journal,
		Language:      book.Language,
		PublishDate:   publishDate,
		ArticleType:   "review",
		SourceType:    "book",
		TotalChapters: &totalChapters,
		InternalLink:  internalLink,
		CreatedBy:     &userID,
		UpdatedBy:     &userID,
		EmbeddingStatus: "pending",
	}

	if err := db.Create(bookArticle).Error; err != nil {
		return fmt.Errorf("erro ao criar livro pai: %w", err)
	}

	fmt.Printf("   ✅ Livro criado: %s (%d capítulos)\n", bookArticle.ID.String()[:8], totalChapters)

	// ── Criar Article para cada capítulo ──
	for i, ch := range chapters {
		chNum := ch.num
		if chNum == 0 {
			chNum = i + 1
		}
		chTitle := ch.title

		chapterTitle := fmt.Sprintf("Capítulo %d: %s", chNum, chTitle)
		if book.Language == "en" {
			chapterTitle = fmt.Sprintf("Chapter %d: %s", chNum, chTitle)
		}

		chapterArticle := &models.Article{
			Title:           fmt.Sprintf("%s — %s", book.Title, chTitle),
			Authors:         book.Authors,
			Journal:         book.Journal,
			Language:        book.Language,
			PublishDate:     publishDate,
			ArticleType:     "review",
			SourceType:      "book_chapter",
			ParentArticleID: &bookArticle.ID,
			ChapterNumber:   &chNum,
			ChapterTitle:    &chapterTitle,
			FullContent:     &ch.content,
			CreatedBy:       &userID,
			UpdatedBy:       &userID,
			EmbeddingStatus: "pending",
		}

		if err := db.Create(chapterArticle).Error; err != nil {
			fmt.Printf("   ⚠️  Erro capítulo %d (%s): %v\n", chNum, ch.path, err)
			continue
		}

		if queueEmbeddings {
			if err := qSvc.QueueArticle(chapterArticle.ID); err != nil {
				fmt.Printf("   ⚠️  Erro ao enfileirar embedding cap %d: %v\n", chNum, err)
			}
		}

		fmt.Printf("   📖 Cap %2d: %s\n", chNum, chTitle)
	}

	return nil
}

func main() {
	godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Erro ao carregar config:", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	userID, err := uuid.Parse("019b8139-b683-75b2-8d27-82c3db7deedd")
	if err != nil {
		log.Fatal("UUID inválido:", err)
	}

	qSvc := services.NewEmbeddingQueueService(db)

	// Caminho base dos exemplos — dentro do container é /app/examples
	// Fora do container (dev) é apps/api/examples
	basePath := "/app/examples"
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		// Tentar caminho local
		basePath = "examples"
		if _, err2 := os.Stat(basePath); os.IsNotExist(err2) {
			log.Fatal("Diretório examples não encontrado em /app/examples nem ./examples")
		}
	}

	fmt.Printf("📂 Base path: %s\n", basePath)
	fmt.Printf("👤 User ID: %s\n", userID.String()[:8])

	// Importar cada livro
	// O primeiro livro recebe embedding completo; os demais apenas são enfileirados
	for i, book := range books {
		queueEmb := true // queue embeddings para todos
		if err := importBook(db, qSvc, basePath, book, userID, queueEmb); err != nil {
			fmt.Printf("   ❌ Erro ao importar '%s': %v\n", book.Title, err)
		}
		_ = i
	}

	// Resumo final
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	var bookCount, chapterCount int64
	db.Model(&models.Article{}).Where("source_type = ?", "book").Count(&bookCount)
	db.Model(&models.Article{}).Where("source_type = ?", "book_chapter").Count(&chapterCount)
	fmt.Printf("📊 Total no banco: %d livros, %d capítulos\n", bookCount, chapterCount)

	var queueCount int64
	db.Table("embedding_queue").Where("entity_type = ? AND status IN (?)", "article", []string{"pending", "processing"}).Count(&queueCount)
	fmt.Printf("📥 Jobs de embedding na fila: %d\n", queueCount)
	fmt.Println("✅ Importação concluída!")
}
