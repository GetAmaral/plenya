package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	gomessage "github.com/emersion/go-message"
	"github.com/emersion/go-message/charset"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// EmailIngestService mantém workers IMAP IDLE pra cada (account, folder) configurado.
// Cada folder roda em sua própria goroutine; reconectam com backoff exponencial em falha.
//
// Fluxo:
//  1. Connect TLS na porta 993, login, SELECT folder
//  2. Resume: SEARCH UID > last_uid (do email_ingest_states), processa pendentes
//  3. IDLE loop: aguarda EXISTS, faz catch-up via UID search, repete
//
// Inbox → ProcessInboundEmail (cria/atualiza Lead).
// Sent  → RecordOutboundEmailMirror (registra histórico, não cria Lead).
type EmailIngestService struct {
	db      *gorm.DB
	cfg     *config.Config
	leadSvc *LeadService
}

// NewEmailIngestService constrói o serviço (não inicia o worker — chame Start).
func NewEmailIngestService(db *gorm.DB, cfg *config.Config, leadSvc *LeadService) *EmailIngestService {
	return &EmailIngestService{db: db, cfg: cfg, leadSvc: leadSvc}
}

// Start lança um worker por folder configurado. Retorna imediatamente.
// ctx.Cancel() encerra todos os workers (eles fazem logout best-effort).
func (s *EmailIngestService) Start(ctx context.Context) {
	mc := s.cfg.MailIngest
	if !mc.Enabled {
		log.Println("📧 [MAIL INGEST] disabled (set MAIL_INGEST_ENABLED=true para ativar)")
		return
	}
	if mc.Username == "" || mc.Password == "" {
		log.Println("⚠️  [MAIL INGEST] credenciais ausentes (MAIL_INGEST_USERNAME/PASSWORD) — worker desativado")
		return
	}
	if mc.IMAPHost == "" {
		log.Println("⚠️  [MAIL INGEST] MAIL_INGEST_IMAP_HOST vazio — worker desativado")
		return
	}
	if err := os.MkdirAll(mc.AttachmentDir, 0o755); err != nil {
		log.Printf("⚠️  [MAIL INGEST] não consigo criar AttachmentDir %s: %v", mc.AttachmentDir, err)
	}

	// Configura charset reader pro go-message conseguir decodificar latin-1, etc.
	gomessage.CharsetReader = charset.Reader

	folders := mc.Folders
	if len(folders) == 0 {
		folders = []string{"INBOX"}
	}
	log.Printf("📧 [MAIL INGEST] iniciando workers — host=%s user=%s folders=%v",
		mc.IMAPHost, mc.Username, folders)

	for _, folder := range folders {
		f := folder
		goSafe(fmt.Sprintf("mail_ingest_%s", f), func() {
			s.runFolderWorker(ctx, f)
		})
	}
}

// runFolderWorker roda o loop de connect/idle/reconnect pra uma única pasta.
// Fica vivo até ctx.Done().
func (s *EmailIngestService) runFolderWorker(ctx context.Context, folder string) {
	backoff := time.Second
	const maxBackoff = 60 * time.Second

	for {
		if ctx.Err() != nil {
			return
		}
		err := s.runFolderSession(ctx, folder)
		if ctx.Err() != nil {
			return
		}
		if err != nil {
			log.Printf("⚠️  [MAIL INGEST %s] sessão terminou com erro: %v — reconnect em %s", folder, err, backoff)
		} else {
			log.Printf("📧 [MAIL INGEST %s] sessão terminou limpa — reconnect em %s", folder, backoff)
		}
		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			return
		}
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

// runFolderSession faz uma única sessão IMAP completa (connect → idle loop → close).
// Retorna quando a conexão cai ou quando ctx é cancelado.
func (s *EmailIngestService) runFolderSession(ctx context.Context, folder string) error {
	mc := s.cfg.MailIngest
	addr := fmt.Sprintf("%s:%d", mc.IMAPHost, mc.IMAPPort)

	// Canal pra sinalizar EXISTS notification do IDLE (mailbox.NumMessages atualizado).
	mailboxCh := make(chan struct{}, 8)
	opts := &imapclient.Options{
		WordDecoder: &mime.WordDecoder{CharsetReader: charset.Reader},
		UnilateralDataHandler: &imapclient.UnilateralDataHandler{
			Mailbox: func(data *imapclient.UnilateralDataMailbox) {
				if data.NumMessages != nil {
					select {
					case mailboxCh <- struct{}{}:
					default:
					}
				}
			},
		},
	}

	c, err := imapclient.DialTLS(addr, opts)
	if err != nil {
		return fmt.Errorf("dial tls %s: %w", addr, err)
	}
	defer c.Close()

	if err := c.Login(mc.Username, mc.Password).Wait(); err != nil {
		return fmt.Errorf("login: %w", err)
	}
	log.Printf("📧 [MAIL INGEST %s] conectado e autenticado em %s", folder, addr)

	if _, err := c.Select(folder, nil).Wait(); err != nil {
		return fmt.Errorf("select %q: %w", folder, err)
	}

	// Resume: processa tudo > last_uid.
	if err := s.catchUp(c, folder); err != nil {
		log.Printf("⚠️  [MAIL INGEST %s] catch-up: %v", folder, err)
		// Não retorna erro — segue pra IDLE mesmo assim.
	}

	// IDLE loop. Restart a cada notificação ou timeout, processa, volta.
	for {
		if ctx.Err() != nil {
			_ = c.Logout().Wait()
			return nil
		}

		idleCmd, err := c.Idle()
		if err != nil {
			return fmt.Errorf("idle start: %w", err)
		}

		// Espera: EXISTS notification, ctx cancel, ou timeout periódico (sanity).
		select {
		case <-mailboxCh:
			// Drena buffer (múltiplas notifs em sequência viram 1 catch-up).
			drainSignal(mailboxCh)
		case <-ctx.Done():
			_ = idleCmd.Close()
			_ = c.Logout().Wait()
			return nil
		case <-time.After(20 * time.Minute):
			// Sanity ping — força reconnect periódico pra detectar conexão zumbi.
			_ = idleCmd.Close()
			return nil
		}

		if err := idleCmd.Close(); err != nil {
			return fmt.Errorf("idle close: %w", err)
		}

		if err := s.catchUp(c, folder); err != nil {
			log.Printf("⚠️  [MAIL INGEST %s] catch-up pós-idle: %v", folder, err)
		}
	}
}

// catchUp lê last_uid do banco e processa todos os UIDs > last_uid na pasta selecionada.
func (s *EmailIngestService) catchUp(c *imapclient.Client, folder string) error {
	mc := s.cfg.MailIngest
	state, err := s.loadOrCreateState(mc.Username, folder)
	if err != nil {
		return fmt.Errorf("load state: %w", err)
	}

	// SEARCH UID > last_uid. Se last_uid==0, pega tudo (primeiro boot).
	criteria := &imap.SearchCriteria{}
	if state.LastUID > 0 {
		criteria.UID = []imap.UIDSet{{{Start: imap.UID(state.LastUID + 1), Stop: 0}}}
	}
	searchData, err := c.UIDSearch(criteria, nil).Wait()
	if err != nil {
		return fmt.Errorf("uid search: %w", err)
	}
	uids := searchData.AllUIDs()
	if len(uids) == 0 {
		return nil
	}
	log.Printf("📧 [MAIL INGEST %s] %d emails pendentes (last_uid=%d)", folder, len(uids), state.LastUID)

	// Fetch BODY[] + UID. Processa um a um pra atualizar UID incremental
	// (assim falha no meio não reprocessa tudo).
	bodySection := &imap.FetchItemBodySection{}
	for _, uid := range uids {
		fetchOpts := &imap.FetchOptions{
			UID:         true,
			BodySection: []*imap.FetchItemBodySection{bodySection},
		}
		msgs, err := c.Fetch(imap.UIDSetNum(uid), fetchOpts).Collect()
		if err != nil {
			log.Printf("⚠️  [MAIL INGEST %s] fetch uid=%d falhou: %v", folder, uid, err)
			continue
		}
		if len(msgs) == 0 || len(msgs[0].BodySection) == 0 {
			log.Printf("⚠️  [MAIL INGEST %s] uid=%d sem body — skip", folder, uid)
			s.advanceUID(state, uint32(uid))
			continue
		}
		raw := msgs[0].BodySection[0].Bytes
		if err := s.processEmail(folder, uint32(uid), raw); err != nil {
			log.Printf("⚠️  [MAIL INGEST %s] process uid=%d: %v", folder, uid, err)
			// Avança UID mesmo em erro pra não ficar preso em mensagem corrompida.
		}
		s.advanceUID(state, uint32(uid))
	}
	return nil
}

// processEmail parseia a mensagem e despacha pro handler certo (inbox vs sent).
// Defensivo: erros de parse / handler nunca crasham (caller usa goSafe na cadeia).
func (s *EmailIngestService) processEmail(folder string, uid uint32, raw []byte) error {
	pe, err := ParseRFC822(raw)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	if len(pe.From) == 0 {
		return errors.New("sem header From — skip")
	}
	fromAddr := pe.From[0].Address
	var fromName *string
	if pe.From[0].Name != "" {
		n := pe.From[0].Name
		fromName = &n
	}

	if IsBot(fromAddr) {
		log.Printf("📧 [MAIL INGEST %s] uid=%d skip (bot/newsletter): %s", folder, uid, fromAddr)
		return nil
	}

	receivedAt := pe.Date
	if receivedAt.IsZero() {
		receivedAt = time.Now().UTC()
	}

	isSent := isSentFolder(folder)

	// Salva attachments em disco pré-handler (o handler só recebe metadata).
	// LeadID ainda é desconhecido pra inbound novo — usamos pasta temp/uid;
	// pós-create poderíamos mover pro path do lead, mas pra simplificar usamos
	// hash do MessageID (estável, idempotente).
	attachmentRefs, err := s.saveAttachments(pe)
	if err != nil {
		log.Printf("⚠️  [MAIL INGEST %s] uid=%d save attachments: %v", folder, uid, err)
		// continua — corpo é mais importante que anexos
	}

	in := InboundEmailInput{
		FromEmail:   fromAddr,
		FromName:    fromName,
		Subject:     pe.Subject,
		BodyText:    pe.BodyText,
		BodyHTML:    pe.BodyHTML,
		MessageID:   pe.MessageID,
		InReplyTo:   pe.InReplyTo,
		References:  pe.References,
		ReceivedAt:  receivedAt,
		Folder:      folder,
		Attachments: attachmentRefs,
	}

	if isSent {
		// Skip emails enviados pelo próprio EMR — eles já foram registrados como
		// LeadActivity por SendLeadReply. O header X-Plenya-Source é setado lá.
		if isEMROriginated(raw) {
			log.Printf("📧 [MAIL INGEST %s] uid=%d skip (origem EMR — já registrado)", folder, uid)
			return nil
		}
		// Espelha pra cada To/Cc — cada destinatário vira uma activity (se tiver lead).
		recipients := append([]ParsedAddress{}, pe.To...)
		recipients = append(recipients, pe.Cc...)
		seen := map[string]bool{}
		for _, r := range recipients {
			if r.Address == "" || seen[r.Address] {
				continue
			}
			seen[r.Address] = true
			if err := s.leadSvc.RecordOutboundEmailMirror(r.Address, in); err != nil {
				log.Printf("⚠️  [MAIL INGEST %s] mirror outbound to=%s: %v", folder, r.Address, err)
			}
		}
		return nil
	}

	// Inbound: roteia pra Patient OU Lead (Bloco A — Central de Conversas).
	// Worker não precisa do owner concreto; só importa que persistiu sem erro.
	if _, err := s.leadSvc.ProcessInboundEmail(in); err != nil {
		return fmt.Errorf("process inbound: %w", err)
	}
	return nil
}

// saveAttachments persiste anexos em disco respeitando MaxAttachmentMB.
// Path: <AttachmentDir>/<message-id-hash>/<filename>
// Retorna referencias (path absoluto) pra incluir em metadata.
func (s *EmailIngestService) saveAttachments(pe *ParsedEmail) ([]EmailAttachmentRef, error) {
	if len(pe.Attachments) == 0 {
		return nil, nil
	}
	mc := s.cfg.MailIngest
	maxBytes := mc.MaxAttachmentMB * 1024 * 1024
	if maxBytes <= 0 {
		maxBytes = 50 * 1024 * 1024
	}

	bucket := messageIDBucket(pe.MessageID)
	dir := filepath.Join(mc.AttachmentDir, bucket)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("mkdir %s: %w", dir, err)
	}

	out := make([]EmailAttachmentRef, 0, len(pe.Attachments))
	for _, a := range pe.Attachments {
		if len(a.Data) == 0 {
			continue
		}
		if len(a.Data) > maxBytes {
			log.Printf("📧 [MAIL INGEST] skip attachment %q (%d bytes excede %d MB)",
				a.Filename, len(a.Data), mc.MaxAttachmentMB)
			continue
		}
		fname := a.Filename
		if fname == "" {
			fname = "attachment.bin"
		}
		full := filepath.Join(dir, fname)
		// Idempotente: se arquivo já existe e mesmo tamanho, skip write.
		if st, err := os.Stat(full); err == nil && st.Size() == int64(len(a.Data)) {
			out = append(out, EmailAttachmentRef{
				Filename:    fname,
				Path:        full,
				ContentType: a.ContentType,
				SizeBytes:   len(a.Data),
			})
			continue
		}
		if err := os.WriteFile(full, a.Data, 0o644); err != nil {
			log.Printf("⚠️  [MAIL INGEST] write %s: %v", full, err)
			continue
		}
		out = append(out, EmailAttachmentRef{
			Filename:    fname,
			Path:        full,
			ContentType: a.ContentType,
			SizeBytes:   len(a.Data),
		})
	}
	return out, nil
}

// loadOrCreateState retorna o registro de estado pra (account, folder) — cria se ausente.
func (s *EmailIngestService) loadOrCreateState(account, folder string) (*models.EmailIngestState, error) {
	var state models.EmailIngestState
	err := s.db.Where("account = ? AND folder = ?", account, folder).First(&state).Error
	if err == nil {
		return &state, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	state = models.EmailIngestState{
		ID:      uuid.Must(uuid.NewV7()),
		Account: account,
		Folder:  folder,
		LastUID: 0,
	}
	if err := s.db.Create(&state).Error; err != nil {
		// Race com outra réplica do worker — refetch.
		if err2 := s.db.Where("account = ? AND folder = ?", account, folder).First(&state).Error; err2 == nil {
			return &state, nil
		}
		return nil, err
	}
	return &state, nil
}

// advanceUID atualiza last_uid em memória + banco. Best-effort em erro.
func (s *EmailIngestService) advanceUID(state *models.EmailIngestState, uid uint32) {
	if uid <= state.LastUID {
		return
	}
	state.LastUID = uid
	if err := s.db.Model(state).Updates(map[string]any{
		"last_uid":     uid,
		"last_seen_at": time.Now().UTC(),
	}).Error; err != nil {
		log.Printf("⚠️  [MAIL INGEST] persist last_uid=%d folder=%s: %v", uid, state.Folder, err)
	}
}

// isEMROriginated detecta o header X-Plenya-Source: emr no raw RFC822.
// Match case-insensitive, valor após ':' trimado. Header parsing manual mínimo
// (mais barato que rerodar gomail.CreateReader só pra isso).
func isEMROriginated(raw []byte) bool {
	// Headers terminam na primeira linha vazia (CRLF CRLF ou LF LF).
	headerEnd := -1
	for i := 0; i+1 < len(raw); i++ {
		if raw[i] == '\n' && raw[i+1] == '\n' {
			headerEnd = i
			break
		}
		if i+3 < len(raw) && raw[i] == '\r' && raw[i+1] == '\n' && raw[i+2] == '\r' && raw[i+3] == '\n' {
			headerEnd = i
			break
		}
	}
	if headerEnd < 0 {
		headerEnd = len(raw)
	}
	headers := strings.ToLower(string(raw[:headerEnd]))
	prefix := strings.ToLower(XPlenyaSourceHeader) + ":"
	for _, line := range strings.Split(headers, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, prefix) {
			continue
		}
		val := strings.TrimSpace(line[len(prefix):])
		if val == strings.ToLower(XPlenyaSourceEMR) {
			return true
		}
	}
	return false
}

// isSentFolder reconhece variações comuns de "Sent" usadas por servidores IMAP.
func isSentFolder(folder string) bool {
	f := strings.ToLower(strings.TrimSpace(folder))
	switch f {
	case "sent", "sent items", "sent messages", "sent mail",
		"enviados", "itens enviados", "[gmail]/sent mail":
		return true
	}
	return false
}

// messageIDBucket gera um diretório curto e estável a partir do Message-ID,
// evitando paths absurdamente longos e caracteres inválidos. Sem MessageID,
// usa um bucket fallback baseado em data UTC (granularidade hora).
func messageIDBucket(msgID string) string {
	id := strings.Trim(strings.TrimSpace(msgID), "<>")
	if id == "" {
		return time.Now().UTC().Format("2006-01-02-15") + "-unknown"
	}
	// Mantém apenas chars seguros pra path.
	var b strings.Builder
	for _, r := range id {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'),
			r == '-' || r == '_' || r == '.' || r == '@':
			b.WriteRune(r)
		default:
			b.WriteRune('_')
		}
	}
	out := b.String()
	if len(out) > 120 {
		out = out[:120]
	}
	return out
}

// drainSignal esvazia um canal de sinais sem bloquear.
func drainSignal(ch chan struct{}) {
	for {
		select {
		case <-ch:
		default:
			return
		}
	}
}
