package services

import (
	"strings"
)

// botLocalParts são local-parts que indicam mailer/bot/notification.
// Match exato OU prefixo seguido de "+" / "." / "-" / "_" + qualquer coisa.
var botLocalParts = []string{
	"mailer-daemon",
	"postmaster",
	"no-reply",
	"noreply",
	"do-not-reply",
	"donotreply",
	"bounce",
	"bounces",
	"notifications",
	"notification",
	"alerts",
	"alert",
	"info",
	"news",
	"newsletter",
	"marketing",
	"comms",
	"updates",
	"support-bot",
}

// botDomains são domínios inteiros de plataformas de envio em massa / bots.
// Match: From domain == bot OR From domain ends with "." + bot.
var botDomains = []string{
	"mailchimp.com",
	"sendgrid.net",
	"amazonses.com",
	"hubspot.com",
	"intercom-mail.com",
	"mailgun.org",
	"postmarkapp.com",
	"sparkpostmail.com",
	"mandrillapp.com",
	"em.sendgrid.net",
}

// IsBot retorna true se From é um remetente conhecido como mailer/bot/newsletter.
// Match é case-insensitive. Local-parts permitem variantes (noreply, noreply+x, noreply.x).
//
// Função pura — sem deps externas, fácil de testar.
func IsBot(from string) bool {
	from = strings.TrimSpace(strings.ToLower(from))
	if from == "" {
		return false
	}
	at := strings.LastIndex(from, "@")
	if at < 0 || at == 0 || at == len(from)-1 {
		return false
	}
	local := from[:at]
	domain := from[at+1:]

	// Remove eventual <…> envoltório
	local = strings.TrimPrefix(local, "<")
	domain = strings.TrimSuffix(domain, ">")

	// Domain match
	for _, bd := range botDomains {
		if domain == bd || strings.HasSuffix(domain, "."+bd) {
			return true
		}
	}

	// Local-part match (exato ou prefixo separado por +, ., -, _)
	for _, lp := range botLocalParts {
		if local == lp {
			return true
		}
		// noreply+anything, noreply.anything, etc.
		for _, sep := range []string{"+", ".", "-", "_"} {
			if strings.HasPrefix(local, lp+sep) {
				return true
			}
		}
	}

	return false
}
