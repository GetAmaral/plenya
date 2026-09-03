package services

import (
	"strings"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/pdfdoc"
)

// Hidratação das réguas do deck gerado.
//
// O modelo escolhe QUAL exame vira régua e escreve o que é autoral — `display` (o nome que o
// paciente reconhece), `sub` (o que o exame mede) e `note` (a leitura clínica). Tudo o mais vem do
// dossiê congelado, copiado pelo servidor: `code`, `unit`, `axis`, `segments` e `history`.
//
// Isto não é conveniência, é a mesma regra que a triagem da conversa já aplica: esses campos são
// DO DOSSIÊ, e operação que tenta escrevê-los é recusada, nem vira sugestão. Na primeira geração
// o problema apareceu na forma mais cara possível: o modelo devolveu onze slides bonitos com as
// réguas vazias — `code` em branco, `axis: [0,0]`, `segments: null` — que renderizariam como barra
// sem escala e sem o ponto do paciente. Um deck que parece pronto e não mostra dado nenhum é pior
// que um deck vazio, porque ninguém desconfia dele.
//
// Régua cujo código não existe no dossiê é REMOVIDA e vira aviso. Deixá-la passar significaria
// desenhar uma escala inventada ao lado do nome de um exame real.
func hidrataReguas(slides []pdfdoc.DeckSlide, d *dto.PlanDossierResponse) ([]pdfdoc.DeckSlide, []dto.PlanGenWarning) {
	if d == nil {
		return slides, nil
	}
	porCodigo := map[string]dto.PlanRuler{}
	porNome := map[string]dto.PlanRuler{}
	for _, r := range d.Rulers {
		if r.Code != "" {
			porCodigo[strings.ToUpper(strings.TrimSpace(r.Code))] = r
		}
		porNome[normalizeTestName(r.Name)] = r
	}

	var avisos []dto.PlanGenWarning
	for i := range slides {
		if len(slides[i].Rulers) == 0 {
			continue
		}
		mantidas := make([]pdfdoc.DeckRulerBlock, 0, len(slides[i].Rulers))
		for _, bloco := range slides[i].Rulers {
			fonte, ok := porCodigo[strings.ToUpper(strings.TrimSpace(bloco.Code))]
			if !ok {
				// Fallback pelo nome que o modelo escreveu: ele acerta o exame e erra o código com
				// frequência, e descartar uma régua certa por causa do código seria desperdício.
				fonte, ok = porNome[normalizeTestName(bloco.Display)]
			}
			if !ok {
				avisos = append(avisos, dto.PlanGenWarning{
					SlideIndex: i + 1, SlideID: slides[i].ID, Title: slides[i].Title,
					Numeral: bloco.Display,
					Reason:  "régua removida: este exame não está no prontuário compilado",
				})
				continue
			}
			mantidas = append(mantidas, pdfdoc.DeckRulerBlock{Ruler: reguaDoDossie(fonte, bloco)})
		}
		slides[i].Rulers = mantidas
	}
	return slides, avisos
}

// reguaDoDossie monta a régua do deck: escala e histórico do dossiê, texto do autor.
func reguaDoDossie(fonte dto.PlanRuler, autoral pdfdoc.DeckRulerBlock) pdfdoc.Ruler {
	r := pdfdoc.Ruler{
		Code:    fonte.Code,
		Unit:    fonte.Unit,
		Display: strings.TrimSpace(autoral.Display),
		Sub:     strings.TrimSpace(autoral.Sub),
		Note:    strings.TrimSpace(autoral.Note),
	}
	// `display` vazio cai no nome do catálogo, que é o pior nome para o paciente ("Ferritina -
	// Homens Pós-Menopausa") mas é melhor que régua sem rótulo nenhum.
	if r.Display == "" {
		r.Display = fonte.Name
	}
	if len(fonte.Axis) >= 2 {
		r.Axis = [2]float64{fonte.Axis[0], fonte.Axis[1]}
	}
	for _, s := range fonte.Segments {
		r.Segments = append(r.Segments, pdfdoc.RulerSegment{Level: s.Level, A: s.A, B: s.B})
	}
	for _, p := range fonte.History {
		r.History = append(r.History, pdfdoc.RulerPoint{Value: p.Plot, Text: p.Text, Date: p.Date})
	}
	return r
}
