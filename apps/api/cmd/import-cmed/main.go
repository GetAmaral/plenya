// Comando de importação do catálogo de medicamentos a partir da Lista de Preços de
// Medicamentos (PMC) publicada pela CMED/ANVISA.
//
// A lista é pública e mensal, baixada de:
//
//	https://www.gov.br/anvisa/pt-br/assuntos/medicamentos/cmed/precos
//
// O arquivo (xls_conformidade_site_YYYYMMDD_*.xlsx, ~12 MB) NÃO é versionado no repositório:
// muda todo mês e só serve de insumo. Rotina mensal:
//
//	docker cp xls_conformidade_site_*.xlsx <container>:/tmp/cmed.xlsx
//	./import-cmed --file /tmp/cmed.xlsx --dry-run     # confere o relatório
//	./import-cmed --file /tmp/cmed.xlsx
//
// Idempotente por GGREM. Preserva o que o médico curou à mão: em linhas com curated_at
// preenchido só os campos de FONTE são atualizados (nome, apresentação, laboratório, preço,
// tarja, classe); categoria, concentração, via e regras ficam como ele deixou. Quem sumiu da
// lista nova é desativado, nunca apagado — há prescrições antigas apontando para essas linhas.
package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/cmed"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var reVersionFromFilename = regexp.MustCompile(`(\d{4})(\d{2})\d{2}`)

func main() {
	file := flag.String("file", "", "XLSX da Lista de Preços (obrigatório)")
	version := flag.String("version", "", "Edição da lista no formato YYYYMM (default: deduzido do nome do arquivo)")
	dryRun := flag.Bool("dry-run", false, "Só imprime o relatório, não escreve no banco")
	deactivateMissing := flag.Bool("deactivate-missing", true, "Marca is_active=false em quem sumiu desta edição")
	batchSize := flag.Int("batch", 500, "Linhas por lote no upsert")
	flag.Parse()

	if *file == "" {
		log.Fatal("--file é obrigatório (caminho do XLSX da CMED)")
	}

	edition := *version
	if edition == "" {
		edition = versionFromFilename(*file)
	}
	if edition == "" {
		log.Fatal("não consegui deduzir a edição do nome do arquivo — passe --version YYYYMM")
	}

	rows := toRows(parse(*file))
	printReport(*file, edition, rows)

	if *dryRun {
		fmt.Println("\n--dry-run: nada foi escrito.")
		return
	}

	db := connect()

	before := countCMED(db)
	affected := upsertAll(db, rows, edition, *batchSize)
	after := countCMED(db)

	inserted := after - before
	updated := affected - inserted
	preserved := countCurated(db)

	deactivated := int64(0)
	if *deactivateMissing {
		deactivated = deactivate(db, edition)
	}

	fmt.Printf("\nupsert: %d inseridos · %d atualizados · %d desativados\n", inserted, updated, deactivated)
	fmt.Printf("linhas com curadoria preservada: %d\n", preserved)
}

func parse(path string) []cmed.Record {
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatalf("abrir %s: %v", path, err)
	}
	defer func() { _ = f.Close() }()

	records, err := cmed.ParseFile(f)
	if err != nil {
		log.Fatalf("ler planilha: %v", err)
	}
	if len(records) == 0 {
		log.Fatal("planilha lida, mas nenhuma linha com GGREM — formato inesperado")
	}
	return records
}

// row junta a linha da CMED com o que foi derivado e classificado a partir dela.
type row struct {
	rec        cmed.Record
	derived    cmed.Derived
	class      cmed.Classification
	rules      cmed.Rules
	commonName string
}

func toRows(records []cmed.Record) []row {
	out := make([]row, 0, len(records))
	for _, r := range records {
		c := cmed.Classify(r.Substance, r.TherapeuticClassCode, r.Stripe)

		// O nome que o médico procura é o do produto; a apresentação desempata na tela, não
		// no nome. Linha sem produto (acontece em registros defeituosos) cai na substância.
		name := r.Product
		if name == "" {
			name = r.Substance
		}

		out = append(out, row{
			rec:        r,
			derived:    cmed.DeriveFromPresentation(r.Presentation),
			class:      c,
			rules:      cmed.RulesFor(c.Category),
			commonName: name,
		})
	}
	return out
}

// ── Relatório ────────────────────────────────────────────────────────────────────────────
// É o artefato de verificação do import: roda igual no --dry-run e no import real, e é o que
// se compara entre edições para perceber que a planilha mudou de formato.

func printReport(path, edition string, rows []row) {
	byCategory := map[models.MedicationCategory]int{}
	bySource := map[string]int{}
	byConfidence := map[string]int{}
	needsReview, notPrescribable := 0, 0

	for _, r := range rows {
		byCategory[r.class.Category]++
		bySource[r.class.CategorySource]++
		byConfidence[r.derived.Confidence]++
		if r.class.NeedsReview {
			needsReview++
		}
		if !r.class.IsPrescribable {
			notPrescribable++
		}
	}

	fmt.Printf("CMED %s — %s\n", edition, filepath.Base(path))
	fmt.Printf("%d apresentações lidas (GGREM único)\n\n", len(rows))

	fmt.Println("categoria regulatória:")
	for _, cat := range []models.MedicationCategory{
		models.MedCategorySimple, models.MedCategoryC1, models.MedCategoryC5,
		models.MedCategoryAntibiotic, models.MedCategoryGLP1, models.MedCategoryAB,
	} {
		fmt.Printf("  %-12s %6d\n", cat, byCategory[cat])
	}

	fmt.Printf("\nclassificação derivada %d · palpite conservador %d\n",
		bySource[models.MedCategorySourceDerived], bySource[models.MedCategorySourceFallback])
	fmt.Printf("marcadas para revisão: %d (%.1f%%)\n", needsReview, pct(needsReview, len(rows)))
	fmt.Printf("fora do receituário (tarja preta): %d\n", notPrescribable)
	fmt.Printf("\nconcentração/forma derivadas: alta %d · parcial %d · nenhuma %d\n",
		byConfidence[cmed.ConfidenceHigh], byConfidence[cmed.ConfidenceMedium], byConfidence[cmed.ConfidenceNone])
}

func pct(part, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(part) * 100 / float64(total)
}

// ── Banco ────────────────────────────────────────────────────────────────────────────────

func connect() *gorm.DB {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("carregar config: ", err)
	}
	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("conectar no banco: ", err)
	}
	return db
}

// insertColumns — ordem única usada tanto na lista de colunas quanto na montagem dos args.
var insertColumns = []string{
	"id", "ggrem", "source", "source_version", "last_imported_at", "is_active",
	"common_name", "active_ingredient", "presentation", "laboratory", "product_type", "ean13",
	"therapeutic_class", "therapeutic_class_code", "stripe", "pmc_price", "anvisa_code",
	"concentration", "pharmaceutical_form", "route", "package_quantity", "derivation_confidence",
	"category", "category_source", "needs_review", "is_prescribable",
	"validity_days", "max_per_prescription", "max_treatment_days",
	"requires_digital_signature", "requires_sncr", "created_at", "updated_at",
}

// sourceColumns — a CMED é a autoridade sobre estes: sobrescreve sempre.
var sourceColumns = []string{
	"common_name", "active_ingredient", "presentation", "laboratory", "product_type", "ean13",
	"therapeutic_class", "therapeutic_class_code", "stripe", "pmc_price", "anvisa_code",
	"source_version", "last_imported_at",
}

// curatedColumns — estes são clínicos. Só o import escreve enquanto ninguém curou a linha;
// depois de curated_at, a palavra final é do médico e o reimport mensal não encosta.
var curatedColumns = []string{
	"concentration", "pharmaceutical_form", "route", "package_quantity", "derivation_confidence",
	"category", "category_source", "needs_review", "is_prescribable",
	"validity_days", "max_per_prescription", "max_treatment_days",
	"requires_digital_signature", "requires_sncr",
}

func upsertStatement(batch int) string {
	placeholders := make([]string, 0, batch)
	group := "(" + strings.TrimSuffix(strings.Repeat("?,", len(insertColumns)), ",") + ")"
	for i := 0; i < batch; i++ {
		placeholders = append(placeholders, group)
	}

	sets := make([]string, 0, len(sourceColumns)+len(curatedColumns)+2)
	for _, c := range sourceColumns {
		sets = append(sets, fmt.Sprintf("%s = EXCLUDED.%s", c, c))
	}
	for _, c := range curatedColumns {
		sets = append(sets, fmt.Sprintf(
			"%s = CASE WHEN medication_definitions.curated_at IS NULL THEN EXCLUDED.%s ELSE medication_definitions.%s END",
			c, c, c))
	}
	sets = append(sets, "is_active = true", "updated_at = now()")

	return fmt.Sprintf(
		"INSERT INTO medication_definitions (%s) VALUES %s "+
			"ON CONFLICT (ggrem) WHERE ggrem IS NOT NULL DO UPDATE SET %s",
		strings.Join(insertColumns, ", "),
		strings.Join(placeholders, ", "),
		strings.Join(sets, ", "),
	)
}

func upsertAll(db *gorm.DB, rows []row, edition string, batchSize int) int64 {
	now := time.Now().UTC()
	var affected int64

	for start := 0; start < len(rows); start += batchSize {
		end := start + batchSize
		if end > len(rows) {
			end = len(rows)
		}
		chunk := rows[start:end]

		args := make([]any, 0, len(chunk)*len(insertColumns))
		for _, r := range chunk {
			args = append(args, rowArgs(r, edition, now)...)
		}

		res := db.Exec(upsertStatement(len(chunk)), args...)
		if res.Error != nil {
			log.Fatalf("upsert do lote %d-%d: %v", start, end, res.Error)
		}
		affected += res.RowsAffected
	}
	return affected
}

func rowArgs(r row, edition string, now time.Time) []any {
	return []any{
		uuid.Must(uuid.NewV7()), r.rec.GGREM, models.MedSourceCMED, edition, now, true,
		clip(r.commonName, 500), r.rec.Substance,
		nullable(r.rec.Presentation), nullable(clip(r.rec.Laboratory, 200)),
		nullable(clip(r.rec.ProductType, 60)), nullable(r.rec.EAN13),
		nullable(clip(r.rec.TherapeuticClass, 200)), nullable(clip(r.rec.TherapeuticClassCode, 10)),
		nullable(r.rec.Stripe), r.rec.PMCPrice, nullable(clip(r.rec.Registro, 50)),
		nullable(clip(r.derived.Concentration, 120)), nullable(r.derived.Form), nullable(r.derived.Route),
		r.derived.PackageQty, r.derived.Confidence,
		string(r.class.Category), r.class.CategorySource, r.class.NeedsReview, r.class.IsPrescribable,
		r.rules.ValidityDays, r.rules.MaxPerPrescription, r.rules.MaxTreatmentDays,
		r.rules.RequiresDigitalSignature, r.rules.RequiresSNCR, now, now,
	}
}

func deactivate(db *gorm.DB, edition string) int64 {
	res := db.Exec(`
		UPDATE medication_definitions
		   SET is_active = false, updated_at = now()
		 WHERE source = ? AND is_active AND source_version IS DISTINCT FROM ?`,
		models.MedSourceCMED, edition)
	if res.Error != nil {
		log.Fatalf("desativar quem saiu da lista: %v", res.Error)
	}
	return res.RowsAffected
}

func countCMED(db *gorm.DB) int64 {
	var n int64
	db.Raw(`SELECT count(*) FROM medication_definitions WHERE source = ? AND deleted_at IS NULL`,
		models.MedSourceCMED).Scan(&n)
	return n
}

func countCurated(db *gorm.DB) int64 {
	var n int64
	db.Raw(`SELECT count(*) FROM medication_definitions WHERE curated_at IS NOT NULL AND deleted_at IS NULL`).Scan(&n)
	return n
}

func nullable(s string) any {
	if s == "" {
		return nil
	}
	return s
}

// clip corta o texto no limite da coluna. A planilha é de terceiros e muda todo mês: uma
// descrição maior que o esperado não pode derrubar o import inteiro no meio de um lote.
func clip(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return strings.TrimSpace(string(r[:max]))
}

func versionFromFilename(path string) string {
	m := reVersionFromFilename.FindStringSubmatch(filepath.Base(path))
	if m == nil {
		return ""
	}
	return m[1] + m[2]
}
