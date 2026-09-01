package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
	"github.com/plenya/api/internal/utils"
)

var (
	ErrPatientPlanNotFound = errors.New("plano do paciente não encontrado")
	ErrPatientPlanEmpty    = errors.New("plano sem slides — não há o que publicar")
)

// ErrPatientPlanOverflow — algum slide transborda a moldura de 1920×1080.
//
// Não é aviso, é bloqueio de publicação. O slide tem altura fixa e `overflow:hidden`: conteúdo
// demais não empurra nada nem gera erro, simplesmente SOME do PDF. Publicar assim entrega ao
// paciente um documento com pedaço faltando, e ninguém percebe.
type ErrPatientPlanOverflow struct {
	Slides []pdfdoc.DeckOverflow
}

func (e *ErrPatientPlanOverflow) Error() string {
	parts := make([]string, 0, len(e.Slides))
	for _, s := range e.Slides {
		parts = append(parts, fmt.Sprintf("slide %02d (%s) passa %.0fpx", s.Slide, s.Title, maxF(s.Right, s.Bottom)))
	}
	return "conteúdo não cabe no slide: " + strings.Join(parts, "; ")
}

func maxF(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// PatientPlanService — o plano de devolutiva: rascunho, edição e publicação.
type PatientPlanService struct {
	db        *gorm.DB
	documents *PatientDocumentsService
	signature *SignatureService
	revisions *PatientPlanRevisionService
	dossiers  *PatientPlanDossierService
}

// SetDossierService liga o serviço de dossiê depois da construção. É injeção tardia porque o
// dossiê depende de repositórios que nascem depois do serviço de plano no main.
func (s *PatientPlanService) SetDossierService(d *PatientPlanDossierService) { s.dossiers = d }

func NewPatientPlanService(db *gorm.DB, documents *PatientDocumentsService, signature *SignatureService) *PatientPlanService {
	return &PatientPlanService{
		db: db, documents: documents, signature: signature,
		revisions: NewPatientPlanRevisionService(db),
	}
}

func (s *PatientPlanService) ListByPatient(patientID uuid.UUID) ([]dto.PatientPlanResponse, error) {
	var rows []models.PatientPlan
	if err := s.db.Where("patient_id = ?", patientID).
		Order("created_at DESC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PatientPlanResponse, len(rows))
	for i := range rows {
		out[i] = *toPatientPlanDTO(&rows[i])
	}
	return out, nil
}

// ListPublished — o que o PACIENTE pode ver.
//
// Dois filtros, e os dois importam: `published_at IS NOT NULL` diz que ESTE plano já chegou ao
// paciente alguma vez, e `published_content` é o que ele recebeu. Filtrar por `status` seria
// errado — o médico voltar o plano para rascunho para ajustar uma frase não pode apagar a
// devolutiva que o paciente já tem na mão.
func (s *PatientPlanService) ListPublished(patientID uuid.UUID) ([]dto.PatientPlanResponse, error) {
	var rows []models.PatientPlan
	if err := s.db.
		Where("patient_id = ? AND published_at IS NOT NULL AND published_content IS NOT NULL", patientID).
		Order("published_at DESC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PatientPlanResponse, len(rows))
	for i := range rows {
		out[i] = *toPublishedPlanDTO(&rows[i])
	}
	return out, nil
}

// GetPublished — um plano publicado do próprio paciente. O filtro entra na CONSULTA, não depois:
// um plano nunca publicado tem que responder "não existe", não "existe mas você não pode ver".
//
// Ordem dos argumentos igual à das irmãs (planID, patientID): os dois são uuid.UUID, e trocar num
// call-site futuro compilaria em silêncio — neste caminho, o erro seria buscar o plano do paciente
// errado.
func (s *PatientPlanService) GetPublished(planID, patientID uuid.UUID) (*dto.PatientPlanResponse, error) {
	var plan models.PatientPlan
	err := s.db.Where("id = ? AND patient_id = ? AND published_at IS NOT NULL AND published_content IS NOT NULL",
		planID, patientID).First(&plan).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrPatientPlanNotFound
	}
	if err != nil {
		return nil, err
	}
	return toPublishedPlanDTO(&plan), nil
}

func (s *PatientPlanService) Get(planID, patientID uuid.UUID) (*dto.PatientPlanResponse, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return nil, err
	}
	return toPatientPlanDTO(plan), nil
}

func (s *PatientPlanService) Create(patientID, authorID uuid.UUID, req *dto.SavePatientPlanRequest) (*dto.PatientPlanResponse, error) {
	title := strings.TrimSpace(req.Title)
	if title == "" {
		title = "Seus exames"
	}
	plan := models.PatientPlan{
		PatientID:    patientID,
		Title:        title,
		Status:       models.PatientPlanDraft,
		Version:      1,
		Content:      req.Content,
		AuthorUserID: authorID,
	}
	if plan.Content == nil {
		plan.Content = []pdfdoc.DeckSlide{}
	}
	if req.SourceSnapshotID != nil && *req.SourceSnapshotID != "" {
		id, err := uuid.Parse(*req.SourceSnapshotID)
		if err != nil {
			return nil, errors.New("sourceSnapshotId inválido")
		}
		plan.SourceSnapshotID = &id
	}
	plan.Content, _ = EnsureSlideIDs(plan.Content)

	// O plano nasce já com a revisão 1: histórico que começa vazio faz o `revision_seq` mentir e
	// deixa a primeira edição sem base de comparação.
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&plan).Error; err != nil {
			return err
		}
		seq, err := s.revisions.Record(tx, RecordRevisionInput{
			Plan: &plan, Title: plan.Title, Content: plan.Content,
			AuthorKind: models.PlanAuthorHuman, CreatedByID: authorID,
			Reason: models.PlanRevisionEdit,
		})
		if err != nil {
			return err
		}
		plan.RevisionSeq = seq
		if err := tx.Model(&plan).Update("revision_seq", seq).Error; err != nil {
			return err
		}
		// Congela o prontuário no nascimento do plano. É o chão contra o qual todo número do deck
		// vai ser conferido, e ele precisa ser o mesmo do começo ao fim da autoria.
		if s.dossiers != nil {
			if _, err := s.dossiers.Freeze(tx, plan.ID, patientID, &authorID); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return toPatientPlanDTO(&plan), nil
}

// Update reescreve o conteúdo do plano. Publicar de novo é o que fecha uma versão nova.
//
// Toda gravação passa a deixar revisão. `req.ExpectedRevision`, quando vem, é o token de
// concorrência: cliente que carregou o plano antes de outra escrita (a do assistente, tipicamente)
// leva conflito em vez de apagar o que não viu.
func (s *PatientPlanService) Update(planID, patientID, userID uuid.UUID, req *dto.SavePatientPlanRequest) (*dto.PatientPlanResponse, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return nil, err
	}
	if err := CheckExpected(plan, req.ExpectedRevision); err != nil {
		return nil, err
	}
	if t := strings.TrimSpace(req.Title); t != "" {
		plan.Title = t
	}
	if req.Content != nil {
		plan.Content, _ = EnsureSlideIDs(req.Content)
	}
	// Editar volta o plano para rascunho e solta os ponteiros de documento: o conteúdo mudou, então
	// aqueles PDFs não representam mais este rascunho. O que o PACIENTE vê não se mexe —
	// `PublishedContent` fica como está, e é dele que o portal lê. `PublishedAt` fica como registro,
	// e a tela do médico mostra "rascunho, vN no portal desde <data>".
	plan.Status = models.PatientPlanDraft
	plan.Document16x9ID = nil
	plan.DocumentA4ID = nil

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		seq, err := s.revisions.Record(tx, RecordRevisionInput{
			Plan: plan, Title: plan.Title, Content: plan.Content,
			AuthorKind: models.PlanAuthorHuman, CreatedByID: userID,
			Reason: models.PlanRevisionEdit,
		})
		if err != nil {
			return err
		}
		plan.RevisionSeq = seq
		return tx.Save(plan).Error
	}); err != nil {
		return nil, err
	}
	return toPatientPlanDTO(plan), nil
}

func (s *PatientPlanService) Delete(planID, patientID uuid.UUID) error {
	res := s.db.Where("id = ? AND patient_id = ?", planID, patientID).Delete(&models.PatientPlan{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrPatientPlanNotFound
	}
	return nil
}

// Preview devolve o HTML do deck sem publicar nada — é o que a tela de montagem mostra.
//
// `fontBase` faz as fontes virem por link em vez de embutidas. Com elas embutidas o HTML tem
// 1,97 MB, dos quais 96% são as fontes; por link são 75 KB, e o navegador cacheia. Vazio mantém o
// comportamento antigo, para nenhum chamador ficar sem fonte por engano.
func (s *PatientPlanService) Preview(planID, patientID uuid.UUID, fontBase string) (string, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return "", err
	}
	// Rascunho recém-criado ainda não tem slide: é caso normal, e a tela precisa de um 400 com
	// explicação, não de um 500.
	if len(plan.Content) == 0 {
		return "", ErrPatientPlanEmpty
	}
	return pdfdoc.DeckHTMLForBrowser(s.deck(plan), fontBase)
}

// CheckOverflow mede se algum slide transborda, sem publicar.
func (s *PatientPlanService) CheckOverflow(planID, patientID uuid.UUID) ([]pdfdoc.DeckOverflow, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return nil, err
	}
	if len(plan.Content) == 0 {
		return nil, ErrPatientPlanEmpty
	}
	return pdfdoc.CheckDeckOverflow(s.deck(plan))
}

// Publish gera os dois PDFs e publica no portal do paciente.
//
// SÍNCRONO de propósito, apesar de o Chromium ser serializado por um mutex global: medido, um deck
// de 8 slides sai em ~1,0 s no 16:9 e ~0,2 s no A4, e são DUAS passagens (a medição de transbordo
// vem junto do render do 16:9). Uma publicação segura a fila da receita por poucos segundos. Vale
// reavaliar se aparecer deck muito maior ou uso concorrente de verdade; hoje uma fila de jobs só
// para isso seria peso sem retorno.
func (s *PatientPlanService) Publish(planID, patientID, publisherID uuid.UUID) (*dto.PatientPlanResponse, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return nil, err
	}
	if len(plan.Content) == 0 {
		return nil, ErrPatientPlanEmpty
	}

	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	deck := s.deck(plan)

	// Barreira: o que não cabe no slide some do PDF em silêncio. Melhor recusar e dizer onde.
	//
	// A medição sai de graça no render do 16:9 (o hook roda antes de imprimir), então são DUAS
	// passagens pelo Chromium e não três. O Chromium é serializado por mutex global e atende
	// também receita e pedido de exames: uma passagem a menos é fila a menos para todo mundo.
	pdf169, over, err := pdfdoc.RenderDeckMeasured(deck)
	if err != nil {
		return nil, fmt.Errorf("erro ao medir os slides: %w", err)
	}
	if len(over) > 0 {
		return nil, &ErrPatientPlanOverflow{Slides: over}
	}

	// Publicar depois de EDITAR é uma versão nova: o documento antigo continua no portal e o
	// `source_ref` novo é diferente, então nada é sobrescrito. Clicar "publicar" duas vezes num
	// plano intocado não é isso — ali a versão fica onde está, o source_ref bate e o
	// `CreateFromBytes` devolve os mesmos documentos, sem encher a lista do paciente de PDFs
	// idênticos. `Update` é quem devolve o plano para rascunho.
	version := plan.Version
	if plan.PublishedAt != nil && plan.Status == models.PatientPlanDraft {
		version++
	}
	now := time.Now()

	type output struct {
		paper    pdfdoc.DeckPaper
		kind     string
		suffix   string
		title    string
		assignTo **uuid.UUID
	}
	var doc169, docA4 *uuid.UUID
	outputs := []output{
		{pdfdoc.DeckPaper169, "Plano", "16x9", plan.Title, &doc169},
		{pdfdoc.DeckPaperA4, "PlanoImpressao", "a4", plan.Title + " (impressão)", &docA4},
	}

	// Renderiza os DOIS antes de publicar qualquer um. Publicando dentro do mesmo laço, uma falha
	// no A4 deixava o 16:9 já visível no portal enquanto o médico via "falha ao publicar" e o plano
	// continuava rascunho, sem registro do documento órfão.
	rendered := make([][]byte, len(outputs))
	for i, o := range outputs {
		if o.paper == pdfdoc.DeckPaper169 {
			rendered[i] = pdf169 // já veio da medição
			continue
		}
		bytesPDF, rErr := pdfdoc.RenderDeck(deck, o.paper)
		if rErr != nil {
			return nil, fmt.Errorf("erro ao gerar o PDF %s do plano: %w", o.paper, rErr)
		}
		rendered[i] = bytesPDF
	}

	for i, o := range outputs {
		bytesPDF := rendered[i]
		// Nome legível: o PDF sai do EMR e vai parar na pasta de downloads do paciente.
		filename := utils.DocumentFileName(patient.Name, o.kind, now, plan.ID)
		sourceRef := fmt.Sprintf("patient_plan:%s:v%d:%s", plan.ID, version, o.suffix)
		doc, dErr := s.documents.CreateFromBytes(CreateFromBytesInput{
			PatientID:  patientID,
			Bytes:      bytesPDF,
			Filename:   filename,
			Title:      o.title,
			Type:       models.DocumentTypeReport,
			Source:     models.DocumentSourceStaffUpload,
			UploadedBy: &publisherID,
			SourceRef:  &sourceRef,
		})
		if dErr != nil {
			return nil, fmt.Errorf("erro ao publicar o plano no portal: %w", dErr)
		}
		id := doc.ID
		*o.assignTo = &id
	}

	plan.Status = models.PatientPlanPublished
	plan.Version = version
	plan.PublishedAt = &now
	// Congela o que foi publicado: daqui para frente o médico edita `Content` sem tirar a
	// devolutiva da tela do paciente.
	plan.PublishedContent = plan.Content
	plan.Document16x9ID = doc169
	plan.DocumentA4ID = docA4

	// A revisão de publicação é o que preserva os bytes exatos de CADA versão. `PublishedContent`
	// guarda só a última: sem esta linha, republicar apagava para sempre o que o paciente leu na
	// versão anterior — o PDF continuava no portal, mas o banco deixava de saber o que havia nele.
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		seq, err := s.revisions.Record(tx, RecordRevisionInput{
			Plan: plan, Title: plan.Title, Content: plan.Content,
			AuthorKind: models.PlanAuthorHuman, CreatedByID: publisherID,
			Reason: models.PlanRevisionPublish, IsPublication: true,
		})
		if err != nil {
			return err
		}
		plan.RevisionSeq = seq
		return tx.Save(plan).Error
	}); err != nil {
		return nil, err
	}
	return toPatientPlanDTO(plan), nil
}

func (s *PatientPlanService) deck(plan *models.PatientPlan) pdfdoc.Deck {
	return pdfdoc.Deck{Title: plan.Title, Slides: plan.Content}
}

func (s *PatientPlanService) load(planID, patientID uuid.UUID) (*models.PatientPlan, error) {
	var plan models.PatientPlan
	// Escopado pelo paciente SEMPRE: nenhuma tela de prontuário lê por id solto.
	err := s.db.Where("id = ? AND patient_id = ?", planID, patientID).First(&plan).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrPatientPlanNotFound
	}
	if err != nil {
		return nil, err
	}
	// Deck escrito fora da tela (a skill, um PUT direto) chega sem id de slide. Preencher aqui, e
	// não só na migration, é o que garante alvo estável para toda operação e sugestão.
	if slides, mudou := EnsureSlideIDs(plan.Content); mudou {
		plan.Content = slides
		if err := s.db.Model(&plan).Update("content", plan.Content).Error; err != nil {
			return nil, err
		}
	}
	return &plan, nil
}

// toPublishedPlanDTO é a visão do PACIENTE: o `content` que sai é o congelado na publicação, não o
// rascunho vivo do médico.
func toPublishedPlanDTO(p *models.PatientPlan) *dto.PatientPlanResponse {
	r := toPatientPlanDTO(p)
	r.Content = p.PublishedContent
	if r.Content == nil {
		r.Content = []pdfdoc.DeckSlide{}
	}
	r.Status = string(models.PatientPlanPublished)
	return r
}

func toPatientPlanDTO(p *models.PatientPlan) *dto.PatientPlanResponse {
	r := &dto.PatientPlanResponse{
		ID:           p.ID.String(),
		PatientID:    p.PatientID.String(),
		Title:        p.Title,
		Status:       string(p.Status),
		Version:      p.Version,
		RevisionSeq:  p.RevisionSeq,
		Content:      p.Content,
		AuthorUserID: p.AuthorUserID.String(),
		CreatedAt:    p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    p.UpdatedAt.Format(time.RFC3339),
	}
	if r.Content == nil {
		r.Content = []pdfdoc.DeckSlide{}
	}
	if p.SourceSnapshotID != nil {
		v := p.SourceSnapshotID.String()
		r.SourceSnapshotID = &v
	}
	if p.PublishedAt != nil {
		v := p.PublishedAt.In(saoPaulo()).Format(time.RFC3339)
		r.PublishedAt = &v
	}
	if p.Document16x9ID != nil {
		v := p.Document16x9ID.String()
		r.Document16x9ID = &v
	}
	if p.DocumentA4ID != nil {
		v := p.DocumentA4ID.String()
		r.DocumentA4ID = &v
	}
	return r
}

// PublishReport gera o TERCEIRO modo do plano: o relatório A4 assinado.
//
// Mesmo conteúdo dos slides, achatado no documento fluido da papelaria e assinado com ICP-Brasil.
// A separação é deliberada: o deck 16:9/A4 é peça de comunicação e não leva assinatura; o
// relatório é o documento clínico, e é ele que vale como registro assinado.
func (s *PatientPlanService) PublishReport(planID, patientID, doctorID uuid.UUID) (string, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return "", err
	}
	if len(plan.Content) == 0 {
		return "", ErrPatientPlanEmpty
	}

	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrPatientNotFound
		}
		return "", err
	}
	var doctor models.User
	if err := s.db.First(&doctor, doctorID).Error; err != nil {
		return "", err
	}

	// IssuedDocument em rascunho primeiro: o id dele é o que entra no QR de validação impresso.
	doc := &models.IssuedDocument{
		PatientID:      patientID,
		DoctorID:       doctorID,
		Type:           models.IssuedDocReport,
		Title:          plan.Title,
		Body:           "Relatório de devolutiva.",
		Status:         models.IssuedDocDraft,
		IssuedByUserID: doctorID,
	}
	if err := s.db.Create(doc).Error; err != nil {
		return "", err
	}

	validationURL := fmt.Sprintf("https://app.plenyasaude.com.br/documentos/validar/%s", doc.ID)
	now := time.Now()

	render := func(digital bool) ([]byte, error) {
		return pdfdoc.RenderPlanReport(pdfdoc.PlanReport{
			Title:     plan.Title,
			Patient:   pdfdoc.Patient{Name: patient.Name},
			Slides:    plan.Content,
			EmittedAt: now.In(saoPaulo()).Format("02/01/2006"),
			Doctor:    pdfdoc.Doctor{Name: doctor.Name, Credentials: doctorCredentials(&doctor)},
			Signature: pdfdoc.Signature{
				Digital:     digital,
				SignedAt:    signedAtPT(&now),
				ValidateURL: validationURL,
				PlaceDate:   placeDatePT(now),
			},
		})
	}

	// O IssuedDocument nasce em rascunho ANTES do render porque o id dele é o que vai impresso no
	// QR. Se o render ou a assinatura falhar, ele tem que sair: senão cada tentativa deixa uma
	// linha sem PDF e sem documento no prontuário, e a lista de documentos emitidos vai enchendo
	// de fantasma.
	descartaRascunho := func() {
		s.db.Delete(&models.IssuedDocument{}, doc.ID)
	}

	out, err := signOrDegrade(s.signature, &doctor, doctorID,
		"Assinatura Digital de Relatório Médico", "Plenya EMR - Relatório de devolutiva", render)
	if err != nil {
		descartaRascunho()
		return "", fmt.Errorf("erro ao gerar o relatório do plano: %w", err)
	}

	// A chave de idempotência é o IssuedDocument, NÃO a versão do plano.
	//
	// Cada geração de relatório é um documento assinado novo, com um QR de validação próprio
	// impresso dentro dele. Chaveando por versão, editar os slides sem republicar o deck (o que não
	// muda a versão) e gerar o relatório de novo casava com o source_ref antigo: o
	// `CreateFromBytes` devolvia o PDF ANTIGO, o render novo ia fora, e o paciente ficava com um
	// arquivo cujo QR aponta para um IssuedDocument diferente do que o banco registrou.
	sourceRef := "patient_plan_report:" + doc.ID.String()
	uploadedBy := doctorID
	patientDoc, err := s.documents.CreateFromBytes(CreateFromBytesInput{
		PatientID:  patientID,
		Bytes:      out.Bytes,
		Filename:   utils.DocumentFileName(patient.Name, "Relatorio", now, doc.ID),
		Title:      plan.Title,
		Type:       models.DocumentTypeReport,
		Source:     models.DocumentSourceStaffUpload,
		UploadedBy: &uploadedBy,
		SourceRef:  &sourceRef,
	})
	if err != nil {
		descartaRascunho()
		return "", fmt.Errorf("erro ao publicar o relatório: %w", err)
	}

	updates := map[string]interface{}{
		"status":                models.IssuedDocSigned,
		"has_digital_signature": out.Digital,
		"signed_at":             now,
		"signed_pdf_hash":       out.Hash,
		"signed_pdf_path":       patientDoc.FilePath,
		"qr_code_data":          validationURL,
		"patient_document_id":   patientDoc.ID,
	}
	if out.CertSerial != nil {
		updates["certificate_serial"] = *out.CertSerial
	}
	if err := s.db.Model(doc).Updates(updates).Error; err != nil {
		return "", err
	}
	return doc.ID.String(), nil
}
