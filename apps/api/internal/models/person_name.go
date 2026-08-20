package models

import (
	"regexp"
	"strings"
	"unicode"
)

// NormalizePersonName põe um nome de pessoa no padrão "João da Silva", venha ele em CAIXA ALTA,
// em minúsculas ou misturado.
//
// A regra "palavra de uma ou duas letras fica minúscula" resolve a maioria dos casos, mas quebra
// nomes reais: Sá, Pó e Ré são sobrenomes portugueses de duas letras, e Li, Xu e Ye são
// sobrenomes chineses comuns no Brasil. Por isso o que fica minúsculo é uma LISTA de partículas,
// não um tamanho — assim "maria de sá" vira "Maria de Sá", e não "Maria de sá".
//
// O que a função respeita:
//
//   - acento e cedilha, na caixa certa: "JOÃO GONÇALVES" → "João Gonçalves";
//   - partícula minúscula, exceto quando abre o nome: "DA SILVA, JOÃO" → "Da Silva, João";
//   - hífen e apóstrofo como separadores internos: "ana-maria d'ávila" → "Ana-Maria D'Ávila";
//   - inicial com ou sem ponto em maiúscula: "joão p. da silva" → "João P. da Silva";
//   - sufixo em algarismo romano no fim: "joão paulo ii" → "João Paulo II";
//   - espaço sobrando, que some.
//
// Nome de uma palavra só é caso corriqueiro aqui: a recepção cadastra paciente só pelo primeiro
// nome, e a função precisa aceitar isso sem inventar nada.
func NormalizePersonName(name string) string {
	limpo := strings.Join(strings.Fields(name), " ")
	if limpo == "" {
		return ""
	}

	palavras := strings.Split(limpo, " ")
	for i, palavra := range palavras {
		switch {
		case i > 0 && ehParticula(palavra):
			palavras[i] = strings.ToLower(palavra)
		case i == len(palavras)-1 && ehNumeralRomano(palavra):
			palavras[i] = strings.ToUpper(palavra)
		default:
			palavras[i] = capitalizaPalavra(palavra)
		}
	}
	return strings.Join(palavras, " ")
}

// Partículas de nome em português e nas origens mais comuns por aqui. "e" entra porque liga
// sobrenomes ("Maria e Silva"); "y" e "i" aparecem em nomes espanhóis e catalães.
var particulasDeNome = map[string]bool{
	"de": true, "da": true, "do": true, "das": true, "dos": true, "e": true,
	"di": true, "du": true, "dal": true, "del": true, "della": true, "dello": true,
	"van": true, "von": true, "der": true, "den": true, "ter": true,
	"la": true, "le": true, "las": true, "los": true, "y": true, "i": true,
}

func ehParticula(p string) bool {
	return particulasDeNome[strings.ToLower(p)]
}

var numeralRomano = regexp.MustCompile(`^(?i)(II|III|IV|V|VI|VII|VIII|IX|X|XI|XII)$`)

// ehNumeralRomano só reconhece sufixo dinástico, e só a partir do II: um "I" solto é quase sempre
// inicial, e "V" ou "X" sozinhos no fim de um nome brasileiro também.
func ehNumeralRomano(p string) bool {
	return numeralRomano.MatchString(p)
}

// capitalizaPalavra sobe a primeira letra de cada pedaço e desce o resto, tratando hífen e
// apóstrofo como fronteira interna: "d'ávila" → "D'Ávila", "ana-maria" → "Ana-Maria".
//
// O PONTO é caso à parte. Ele só abre pedaço novo quando a palavra inteira é feita de iniciais
// ("j.p." → "J.P."). Sem essa ressalva, um e-mail digitado no campo do nome — que acontece —
// virava "Getfilho@yahoo.Com.Br".
func capitalizaPalavra(p string) string {
	if ehIniciais(p) {
		return strings.ToUpper(p)
	}
	var b strings.Builder
	inicioDePedaco := true
	for _, r := range p {
		switch {
		case r == '-' || r == '\'' || r == '’':
			b.WriteRune(r)
			inicioDePedaco = true
		case inicioDePedaco:
			b.WriteRune(unicode.ToUpper(r))
			inicioDePedaco = false
		default:
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

// ehIniciais reconhece "p.", "j.p.", "a.b.c." — palavra formada só por letras soltas com ponto.
func ehIniciais(p string) bool {
	if !strings.Contains(p, ".") {
		return false
	}
	for _, pedaco := range strings.Split(p, ".") {
		if len([]rune(pedaco)) > 1 {
			return false
		}
	}
	return true
}
