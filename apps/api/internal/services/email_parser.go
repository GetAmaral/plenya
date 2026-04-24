package services

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/mail"
	"regexp"
	"strings"
	"time"

	gomessage "github.com/emersion/go-message"
	"github.com/emersion/go-message/charset"
	gomail "github.com/emersion/go-message/mail"
)

// ParsedEmail é o resultado da parse de uma mensagem RFC822.
type ParsedEmail struct {
	MessageID  string
	InReplyTo  string
	References []string
	Date       time.Time

	From []ParsedAddress
	To   []ParsedAddress
	Cc   []ParsedAddress

	Subject  string
	BodyText string // text/plain (preferido). Se ausente, derivado de BodyHTML.
	BodyHTML string // text/html (fallback)

	Attachments []ParsedAttachment
}

// ParsedAddress é um endereço de email com nome opcional.
type ParsedAddress struct {
	Name    string
	Address string // já em lower-case
}

// ParsedAttachment é um anexo extraído (ainda em memória).
type ParsedAttachment struct {
	Filename    string
	ContentType string
	Data        []byte
}

// ParseRFC822 analisa um email completo (headers + body) usando go-message/mail.
// Tolera mensagens malformadas — devolve o que conseguir parsear, com erro se total.
//
// Sempre que possível use go-message charset reader para decoder Subject/From corretos.
func ParseRFC822(raw []byte) (*ParsedEmail, error) {
	if len(raw) == 0 {
		return nil, fmt.Errorf("email parser: payload vazio")
	}

	// Configura charset reader globalmente uma vez (idempotente).
	gomessage.CharsetReader = charset.Reader

	mr, err := gomail.CreateReader(bytes.NewReader(raw))
	if err != nil {
		// Fallback: tenta header-only via net/mail pra ainda extrair From/Subject/Message-ID.
		return parseHeaderOnly(raw, err)
	}
	defer mr.Close()

	pe := &ParsedEmail{}
	header := mr.Header

	if v, err := header.MessageID(); err == nil {
		pe.MessageID = v
	}
	if v, err := header.Date(); err == nil {
		pe.Date = v
	}
	if v, err := header.Text("Subject"); err == nil {
		pe.Subject = strings.TrimSpace(v)
	}
	if v, err := header.AddressList("From"); err == nil {
		pe.From = toParsedAddrs(v)
	}
	if v, err := header.AddressList("To"); err == nil {
		pe.To = toParsedAddrs(v)
	}
	if v, err := header.AddressList("Cc"); err == nil {
		pe.Cc = toParsedAddrs(v)
	}
	if v := header.Get("In-Reply-To"); v != "" {
		pe.InReplyTo = strings.TrimSpace(v)
	}
	if v := header.Get("References"); v != "" {
		pe.References = splitMessageIDs(v)
	}

	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			// Parte malformada — ignora e segue.
			continue
		}
		switch h := p.Header.(type) {
		case *gomail.InlineHeader:
			contentType, _, _ := h.ContentType()
			body, readErr := io.ReadAll(p.Body)
			if readErr != nil {
				continue
			}
			text := string(body)
			switch strings.ToLower(contentType) {
			case "text/plain":
				if pe.BodyText == "" {
					pe.BodyText = text
				}
			case "text/html":
				if pe.BodyHTML == "" {
					pe.BodyHTML = text
				}
			default:
				// Outras inline parts: se ainda não temos text, salva como texto bruto.
				if pe.BodyText == "" && strings.HasPrefix(strings.ToLower(contentType), "text/") {
					pe.BodyText = text
				}
			}
		case *gomail.AttachmentHeader:
			filename, _ := h.Filename()
			contentType, _, _ := h.ContentType()
			data, err := io.ReadAll(p.Body)
			if err != nil {
				continue
			}
			pe.Attachments = append(pe.Attachments, ParsedAttachment{
				Filename:    sanitizeFilename(filename),
				ContentType: contentType,
				Data:        data,
			})
		}
	}

	if pe.BodyText == "" && pe.BodyHTML != "" {
		pe.BodyText = ExtractPlainBody(pe.BodyHTML)
	}

	return pe, nil
}

// parseHeaderOnly é o fallback quando go-message/mail falha (ex: MIME quebrado).
// Extrai apenas headers via net/mail pra ainda termos From/Subject/Message-ID.
func parseHeaderOnly(raw []byte, originalErr error) (*ParsedEmail, error) {
	msg, err := mail.ReadMessage(bytes.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("email parser: parse total falhou: %w (original: %v)", err, originalErr)
	}
	pe := &ParsedEmail{
		MessageID: strings.TrimSpace(msg.Header.Get("Message-ID")),
		Subject:   decodeWord(msg.Header.Get("Subject")),
		InReplyTo: strings.TrimSpace(msg.Header.Get("In-Reply-To")),
	}
	if v := msg.Header.Get("References"); v != "" {
		pe.References = splitMessageIDs(v)
	}
	if v := msg.Header.Get("Date"); v != "" {
		if t, err := mail.ParseDate(v); err == nil {
			pe.Date = t
		}
	}
	if addrs, err := msg.Header.AddressList("From"); err == nil {
		pe.From = toStdParsedAddrs(addrs)
	}
	if addrs, err := msg.Header.AddressList("To"); err == nil {
		pe.To = toStdParsedAddrs(addrs)
	}
	if addrs, err := msg.Header.AddressList("Cc"); err == nil {
		pe.Cc = toStdParsedAddrs(addrs)
	}
	if msg.Body != nil {
		if body, err := io.ReadAll(msg.Body); err == nil {
			pe.BodyText = string(body)
		}
	}
	return pe, nil
}

func toParsedAddrs(addrs []*mail.Address) []ParsedAddress {
	out := make([]ParsedAddress, 0, len(addrs))
	for _, a := range addrs {
		if a == nil {
			continue
		}
		out = append(out, ParsedAddress{
			Name:    strings.TrimSpace(a.Name),
			Address: strings.ToLower(strings.TrimSpace(a.Address)),
		})
	}
	return out
}

func toStdParsedAddrs(addrs []*mail.Address) []ParsedAddress {
	return toParsedAddrs(addrs)
}

func decodeWord(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	dec := &mime.WordDecoder{CharsetReader: charset.Reader}
	if v, err := dec.DecodeHeader(s); err == nil {
		return v
	}
	return s
}

// splitMessageIDs separa um header References ("<id1@x> <id2@x>") em slice.
func splitMessageIDs(s string) []string {
	fields := strings.Fields(s)
	out := make([]string, 0, len(fields))
	for _, f := range fields {
		f = strings.TrimSpace(f)
		if f != "" {
			out = append(out, f)
		}
	}
	return out
}

// sanitizeFilename remove path traversal e caracteres problemáticos do filesystem.
func sanitizeFilename(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "attachment.bin"
	}
	// Apenas basename — nunca path
	if i := strings.LastIndexAny(name, "/\\"); i >= 0 {
		name = name[i+1:]
	}
	// Remove caracteres de controle e separadores perigosos
	var b strings.Builder
	for _, r := range name {
		switch {
		case r < 0x20, r == 0x7f:
			// skip control chars
		case r == '/' || r == '\\' || r == ':' || r == '*' || r == '?' || r == '"' || r == '<' || r == '>' || r == '|':
			b.WriteRune('_')
		default:
			b.WriteRune(r)
		}
	}
	clean := strings.TrimSpace(b.String())
	if clean == "" || clean == "." || clean == ".." {
		return "attachment.bin"
	}
	// Limita comprimento
	if len(clean) > 200 {
		clean = clean[:200]
	}
	return clean
}

// stripHTMLRe é uma regex simples pra remover tags HTML.
// Não é XSS-safe — uso é só pra gerar uma versão "plain" de email pra timeline.
var (
	stripHTMLRe   = regexp.MustCompile(`(?is)<(script|style)[^>]*>.*?</(script|style)>`)
	stripTagsRe   = regexp.MustCompile(`<[^>]+>`)
	stripSpacesRe = regexp.MustCompile(`[ \t]+`)
	stripBlanksRe = regexp.MustCompile(`\n{3,}`)
)

// ExtractPlainBody converte HTML em texto plano (best-effort).
// Suficiente pra mostrar conteúdo na timeline do CRM. Não é replacement
// pra um sanitizer real (bluemonday, goquery, etc.).
func ExtractPlainBody(htmlBody string) string {
	if htmlBody == "" {
		return ""
	}
	s := htmlBody
	// Remove script/style blocos inteiros
	s = stripHTMLRe.ReplaceAllString(s, "")
	// Quebra de linha em <br>, <p>, <div>
	s = strings.ReplaceAll(s, "<br>", "\n")
	s = strings.ReplaceAll(s, "<br/>", "\n")
	s = strings.ReplaceAll(s, "<br />", "\n")
	s = strings.ReplaceAll(s, "</p>", "\n\n")
	s = strings.ReplaceAll(s, "</div>", "\n")
	// Remove tags restantes
	s = stripTagsRe.ReplaceAllString(s, "")
	// Decode entities mais comuns
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	s = strings.ReplaceAll(s, "&#39;", "'")
	// Normaliza espaços
	s = stripSpacesRe.ReplaceAllString(s, " ")
	s = stripBlanksRe.ReplaceAllString(s, "\n\n")
	return strings.TrimSpace(s)
}
