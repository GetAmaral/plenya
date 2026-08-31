// Nome de arquivo dos documentos que saem do EMR.
//
// O PDF não fica só dentro do sistema: o médico reenvia pelo WhatsApp dele e o paciente guarda na
// pasta de downloads. Um arquivo chamado "receita_01a0592b-366c-751b-aa8b-663da085a25b.pdf" não
// identifica nada fora do EMR. O padrão nasceu no pedido de exames e vale para todo documento:
//
//	Nome-Do-Paciente_Receita_2026-08-31_01a0592b.pdf
//
// O sufixo curto do UUID evita colisão entre documentos do mesmo paciente no mesmo dia e permite
// achar o registro a partir do arquivo.
package utils

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// DocumentFileName monta o nome no padrão. kind é o rótulo do documento sem espaço
// ("PedidoExame", "Receita", "ReceitaManipulado"). Nome de paciente vazio cai em "Paciente".
func DocumentFileName(patientName, kind string, date time.Time, id uuid.UUID) string {
	patient := "Paciente"
	if n := CompactName(patientName); n != "" {
		patient = n
	}
	return fmt.Sprintf("%s_%s_%s_%s.pdf", patient, kind, date.Format("2006-01-02"), id.String()[:8])
}

// CompactName transforma "Luiz Gustavo José Carvalho" em "Luiz-Gustavo-José-Carvalho": une as
// partes do nome com hífen e descarta o que não for letra/dígito, para o nome do arquivo não
// depender de espaço nem de pontuação.
//
// O hífen entrou no lugar do CamelCase porque "LuizGustavoJoséCarvalho" só se lê com esforço, e
// esse nome vai parar no WhatsApp e na pasta de downloads de quem recebe.
func CompactName(name string) string {
	var partes []string
	var atual strings.Builder
	fecha := func() {
		if atual.Len() > 0 {
			partes = append(partes, atual.String())
			atual.Reset()
		}
	}
	primeiraDaParte := true
	for _, r := range name {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			if primeiraDaParte {
				atual.WriteRune(unicode.ToUpper(r))
				primeiraDaParte = false
			} else {
				atual.WriteRune(r)
			}
		default:
			fecha()
			primeiraDaParte = true
		}
	}
	fecha()
	return strings.Join(partes, "-")
}

// ASCIIFallback tira acentos e qualquer caractere fora do ASCII imprimível — é o nome que vai no
// filename= "clássico" do Content-Disposition, para cliente antigo que ignora o filename*=UTF-8.
func ASCIIFallback(s, generic string) string {
	normalized, _, err := transform.String(
		transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC), s)
	if err != nil {
		normalized = s
	}
	var b strings.Builder
	for _, r := range normalized {
		if r > 32 && r < 127 && r != '"' && r != '\\' {
			b.WriteRune(r)
		}
	}
	if b.Len() == 0 {
		return generic
	}
	return b.String()
}

// ContentDisposition monta o cabeçalho com as duas formas: filename= em ASCII puro para
// compatibilidade e filename*=UTF-8 com o nome real acentuado.
func ContentDisposition(fileName, generic string) string {
	return fmt.Sprintf(`inline; filename="%s"; filename*=UTF-8''%s`,
		ASCIIFallback(fileName, generic), url.PathEscape(fileName))
}
