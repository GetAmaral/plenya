package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// WhatsAppTemplateService — catálogo de templates WhatsApp: cache + status sincronizado da Meta
// (fonte da verdade) + metadados nossos (purpose/variables/wiringNotes/enabled, preservados no sync).
type WhatsAppTemplateService struct {
	db     *gorm.DB
	cfg    *config.Config
	client *http.Client
}

func NewWhatsAppTemplateService(db *gorm.DB, cfg *config.Config) *WhatsAppTemplateService {
	return &WhatsAppTemplateService{db: db, cfg: cfg, client: &http.Client{Timeout: 15 * time.Second}}
}

var varRe = regexp.MustCompile(`\{\{(\d+)\}\}`)

func (s *WhatsAppTemplateService) graphVersion() string {
	if v := s.cfg.WhatsApp.GraphAPIVersion; v != "" {
		return v
	}
	return "v21.0"
}

// --- shape da resposta da Meta ---
type metaTemplate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Language   string `json:"language"`
	Category   string `json:"category"`
	Status     string `json:"status"`
	Components []struct {
		Type    string `json:"type"`
		Text    string `json:"text"`
		Example struct {
			BodyText [][]string `json:"body_text"`
		} `json:"example"`
	} `json:"components"`
}

type metaTemplatesResp struct {
	Data   []metaTemplate `json:"data"`
	Paging struct {
		Next string `json:"next"`
	} `json:"paging"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

// SyncFromMeta busca os templates da WABA e faz upsert por (name, language), preservando os campos
// nossos. Retorna quantos foram sincronizados.
func (s *WhatsAppTemplateService) SyncFromMeta(ctx context.Context) (int, error) {
	if s.cfg.WhatsApp.AccessToken == "" || s.cfg.WhatsApp.WABAID == "" {
		return 0, fmt.Errorf("whatsapp templates sync: sem WHATSAPP_ACCESS_TOKEN/WABA_ID")
	}
	next := fmt.Sprintf("https://graph.facebook.com/%s/%s/message_templates?fields=name,language,category,status,components&limit=200&access_token=%s",
		s.graphVersion(), s.cfg.WhatsApp.WABAID, url.QueryEscape(s.cfg.WhatsApp.AccessToken))

	count := 0
	for next != "" {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, next, nil)
		if err != nil {
			return count, err
		}
		resp, err := s.client.Do(req)
		if err != nil {
			return count, fmt.Errorf("whatsapp templates sync: %w", err)
		}
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
		resp.Body.Close()
		var parsed metaTemplatesResp
		if err := json.Unmarshal(body, &parsed); err != nil {
			return count, fmt.Errorf("whatsapp templates sync: parse: %w", err)
		}
		if parsed.Error != nil {
			return count, fmt.Errorf("whatsapp templates sync: meta: %s", parsed.Error.Message)
		}
		for _, mt := range parsed.Data {
			if err := s.upsert(mt); err != nil {
				return count, err
			}
			count++
		}
		next = parsed.Paging.Next
	}
	return count, nil
}

func (s *WhatsAppTemplateService) upsert(mt metaTemplate) error {
	body, example := "", []string(nil)
	for _, c := range mt.Components {
		if c.Type == "BODY" {
			body = c.Text
			if len(c.Example.BodyText) > 0 {
				example = c.Example.BodyText[0]
			}
		}
	}
	nvars := 0
	for _, m := range varRe.FindAllStringSubmatch(body, -1) {
		var n int
		fmt.Sscanf(m[1], "%d", &n)
		if n > nvars {
			nvars = n
		}
	}
	now := time.Now().UTC()
	lang := mt.Language
	if lang == "" {
		lang = "pt_BR"
	}

	var existing models.WhatsAppTemplate
	err := s.db.Where("name = ? AND language = ?", mt.Name, lang).First(&existing).Error
	if err == nil {
		// Atualiza campos da Meta; preserva purpose/wiringNotes/enabled. Variáveis: mescla (rótulo
		// editado à mão tem prioridade; senão usa o canônico; exemplo vem da Meta).
		vars := mergeVars(nvars, example, canonicalVarLabels(mt.Name), existing.Variables)
		return s.db.Model(&existing).Updates(map[string]interface{}{
			"meta_id":        mt.ID,
			"category":       mt.Category,
			"status":         mt.Status,
			"body_text":      body,
			"variable_count": nvars,
			"variables":      vars,
			"last_synced_at": now,
		}).Error
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	// Novo: variáveis já nascem com rótulo canônico (descrição do campo) + exemplo da Meta.
	vars := mergeVars(nvars, example, canonicalVarLabels(mt.Name), nil)
	return s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.WhatsAppTemplate{
		MetaID:        mt.ID,
		Name:          mt.Name,
		Language:      lang,
		Category:      mt.Category,
		Status:        mt.Status,
		BodyText:      body,
		VariableCount:  nvars,
		Variables:      vars,
		Enabled:        true,
		ChatSelectable: chatSelectableDefault(mt.Name),
		LastSyncedAt:   &now,
	}).Error
}

// mergeVars monta as variáveis: rótulo = o editado à mão (não-vazio) se houver, senão o canônico;
// exemplo = o da Meta se houver, senão o que já tinha.
func mergeVars(nvars int, example, canonical []string, prev []models.WhatsAppTemplateVar) []models.WhatsAppTemplateVar {
	out := make([]models.WhatsAppTemplateVar, 0, nvars)
	for i := 0; i < nvars; i++ {
		v := models.WhatsAppTemplateVar{Index: i + 1}
		if i < len(prev) {
			v.Label, v.Example = prev[i].Label, prev[i].Example
		}
		if v.Label == "" && i < len(canonical) {
			v.Label = canonical[i]
		}
		if i < len(example) && example[i] != "" {
			v.Example = example[i]
		}
		out = append(out, v)
	}
	return out
}

// canonicalVarLabels descreve o que é cada {{n}} de cada template (conhecimento nosso — a Meta não
// fornece). Garante que todo template venha documentado. Atualizar ao criar/alterar templates.
func canonicalVarLabels(name string) []string {
	switch name {
	case "aniversario_plenya":
		return []string{"Primeiro nome do paciente"}
	case "cancelamento_consulta":
		return []string{"Nome do paciente", "Data da consulta", "Hora"}
	case "confirmacao_consulta_dia":
		return []string{"Nome do paciente", "Hora", "Frase do dia (modalidade/local)"}
	case "confirmacao_consulta_semana":
		return []string{"Nome do paciente", "Data da consulta", "Hora", "Modalidade (presencial/online)"}
	case "confirmacao_consulta_vespera":
		return []string{"Nome do paciente", "Hora", "Modalidade (presencial/online)"}
	case "consultation_prep_invite":
		return []string{"Link seguro de preparação"}
	case "resposta_atendimento":
		return []string{"Primeiro nome do paciente", "A resposta enviada (texto livre)"}
	case "retorno_disponivel":
		return []string{"Primeiro nome do paciente"}
	case "documento_disponivel":
		return []string{"Primeiro nome do paciente", "Tipo do documento (ex.: Receita Médica)", "Link seguro do documento"}
	case "followup_pos_consulta":
		return []string{"Nome do paciente"}
	case "lead_alert":
		return []string{"Nome do lead", "Contato (e-mail/telefone)", "Origem do lead", "Link do painel administrativo"}
	case "magic_link":
		return []string{"Link do resultado do Escore Light"}
	case "reengajamento_lead":
		return []string{"Nome do lead"}
	case "remarcacao_consulta":
		return []string{"Nome do paciente", "Nova data", "Nova hora", "Modalidade (presencial/online)"}
	default:
		return nil
	}
}

// List retorna o catálogo (ordenado por status então nome).
func (s *WhatsAppTemplateService) List() ([]models.WhatsAppTemplate, error) {
	var out []models.WhatsAppTemplate
	err := s.db.Order("status ASC, name ASC").Find(&out).Error
	return out, err
}

// ListSendable retorna os templates que o atendente pode escolher no compositor: APPROVED + enabled
// + chat_selectable (exclui internos/automáticos). Substitui a lista hardcoded.
func (s *WhatsAppTemplateService) ListSendable() ([]models.WhatsAppTemplate, error) {
	var out []models.WhatsAppTemplate
	err := s.db.Where("enabled = ? AND status = ? AND chat_selectable = ?", true, "APPROVED", true).
		Order("name ASC").Find(&out).Error
	return out, err
}

// chatSelectableDefault: templates internos (notificação ao time) ou automáticos (precisam de URL
// gerada pelo sistema / teste) NÃO são escolhíveis à mão pelo atendente.
func chatSelectableDefault(name string) bool {
	switch name {
	case "lead_alert", "magic_link", "consultation_prep_invite", "documento_disponivel", "hello_world":
		return false
	default:
		return true
	}
}

// IsSendable é o GUARD: só permite enviar template APROVADO e habilitado. Se não houver registro
// no catálogo (ainda não sincronizou), permite (não bloqueia o que já funcionava) mas o caller loga.
func (s *WhatsAppTemplateService) IsSendable(name, lang string) (bool, string) {
	if lang == "" {
		lang = "pt_BR"
	}
	var t models.WhatsAppTemplate
	if err := s.db.Where("name = ? AND language = ?", name, lang).First(&t).Error; err != nil {
		return true, "" // sem catálogo => não bloqueia (defesa só quando há registro)
	}
	if !t.Enabled {
		return false, "template desabilitado no catálogo"
	}
	if t.Status != "APPROVED" {
		return false, "template não aprovado na Meta (status: " + t.Status + ")"
	}
	return true, ""
}

// UpdateMetaFields atualiza os campos NOSSOS (purpose/variables/wiringNotes/enabled).
type UpdateTemplateInput struct {
	Purpose        *string
	WiringNotes    *string
	Enabled        *bool
	ChatSelectable *bool
	Variables      *[]models.WhatsAppTemplateVar
}

func (s *WhatsAppTemplateService) UpdateMetaFields(id uuid.UUID, in UpdateTemplateInput) (*models.WhatsAppTemplate, error) {
	var t models.WhatsAppTemplate
	if err := s.db.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	upd := map[string]interface{}{}
	if in.Purpose != nil {
		upd["purpose"] = *in.Purpose
	}
	if in.WiringNotes != nil {
		upd["wiring_notes"] = *in.WiringNotes
	}
	if in.Enabled != nil {
		upd["enabled"] = *in.Enabled
	}
	if in.ChatSelectable != nil {
		upd["chat_selectable"] = *in.ChatSelectable
	}
	if in.Variables != nil {
		upd["variables"] = *in.Variables
	}
	if len(upd) > 0 {
		if err := s.db.Model(&t).Updates(upd).Error; err != nil {
			return nil, err
		}
	}
	return &t, nil
}
