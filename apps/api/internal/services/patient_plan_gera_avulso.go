package services

// Geração de deck SEM banco.
//
// O caminho normal (`GenerateDraft`) cria o plano, congela o dossiê e grava revisões. Para comparar
// o gerador contra um deck aprovado isso é ruído: o que se compara é o DECK, não a persistência.
// Esta função roda exatamente as mesmas etapas — arco em código, seções em paralelo com o cache
// aquecido, hidratação das réguas, regras do servidor, conferência e reparo — e devolve os slides.
//
// Existe para medir e comparar. Não é chamada pelo produto.

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/pdfdoc"
)

// GeraDeckParaMedida devolve o deck e o consumo de tokens.
func GeraDeckParaMedida(ai *AIService, dossie *dto.PlanDossierResponse, modelo string) (
	[]pdfdoc.DeckSlide, AICallMeta, error,
) {
	var meta AICallMeta
	if dossie == nil {
		return nil, meta, errors.New("dossiê vazio")
	}
	secoes := montaArco(dossie)
	if semMaterial(secoes, dossie) {
		return nil, meta, errors.New("sem material para uma devolutiva")
	}
	arcoTexto := descreveArco(secoes)
	dossieJSON, err := podaDossieParaPrompt(dossie, nil)
	if err != nil {
		return nil, meta, err
	}

	porSecao := make([][]pdfdoc.DeckSlide, len(secoes))
	erros := make([]error, len(secoes))
	var mu sync.Mutex

	escreve := func(i int, sec PlanArcSection) {
		r, m, e := ai.GeneratePlanSection(PlanSectionRequest{
			DossierJSON: dossieJSON, ArcoJSON: arcoTexto, Secao: sec, Indice: i, Model: modelo,
		})
		mu.Lock()
		defer mu.Unlock()
		if e != nil {
			erros[i] = e
			return
		}
		porSecao[i] = r.Slides
		if t := strings.TrimSpace(r.Label); t != "" {
			secoes[i].Label = t
		}
		meta.Model = m.Model
		meta.InputTokens += m.InputTokens
		meta.OutputTokens += m.OutputTokens
		meta.CacheReadTokens += m.CacheReadTokens
		meta.CacheWriteTokens += m.CacheWriteTokens
	}

	// A primeira sozinha, para aquecer o cache; o resto em paralelo com teto de 3.
	escreve(0, secoes[0])
	if erros[0] != nil {
		return nil, meta, fmt.Errorf("seção %q: %w", secoes[0].Label, erros[0])
	}
	var wg sync.WaitGroup
	vagas := make(chan struct{}, 3)
	for i, sec := range secoes[1:] {
		i++
		wg.Add(1)
		go func(i int, sec PlanArcSection) {
			defer wg.Done()
			vagas <- struct{}{}
			defer func() { <-vagas }()
			escreve(i, sec)
		}(i, sec)
	}
	wg.Wait()
	for i, e := range erros {
		if e != nil {
			return nil, meta, fmt.Errorf("seção %q: %w", secoes[i].Label, e)
		}
	}

	brutos := aplicaRegrasDoDeck(secoes, porSecao,
		primeiroNomeDe(dossie.Patient.Name), formataDataPT(dossie.GeneratedAt))
	for i := range brutos {
		brutos[i].ID = ""
	}
	slides, _ := EnsureSlideIDs(brutos)
	slides, _ = hidrataReguas(slides, dossie)
	poeLegenda(slides)

	// Uma rodada de reparo, como no produto.
	estouro, _ := pdfdoc.CheckDeckOverflow(pdfdoc.Deck{Title: "Seus exames", Slides: slides})
	estilo := confereDeck(slides)
	if len(estouro) > 0 || len(estilo) > 0 {
		if novos, m, ok := reparaAvulso(ai, slides, estouro, estilo, modelo); ok {
			slides = novos
			meta.InputTokens += m.InputTokens
			meta.OutputTokens += m.OutputTokens
			meta.CacheReadTokens += m.CacheReadTokens
			meta.CacheWriteTokens += m.CacheWriteTokens
		}
	}
	return slides, meta, nil
}

// reparaAvulso é o reparo sem banco, com a mesma regra de preservar o que o servidor decide.
func reparaAvulso(ai *AIService, slides []pdfdoc.DeckSlide, estouro []pdfdoc.DeckOverflow,
	estilo []dto.PlanGenWarning, modelo string) ([]pdfdoc.DeckSlide, AICallMeta, bool) {
	var meta AICallMeta
	problemas := map[string][]string{}
	porIndice := map[int]pdfdoc.DeckOverflow{}
	for _, o := range estouro {
		porIndice[o.Slide] = o
	}
	for i, sl := range slides {
		if o, ok := porIndice[i+1]; ok {
			if o.Bottom > 1 {
				problemas[sl.ID] = append(problemas[sl.ID], fmt.Sprintf("sobraram %.0fpx de ALTURA", o.Bottom))
			}
			if o.Right > 1 {
				problemas[sl.ID] = append(problemas[sl.ID], fmt.Sprintf("sobraram %.0fpx de LARGURA", o.Right))
			}
		}
	}
	for _, w := range estilo {
		if w.SlideID != "" {
			problemas[w.SlideID] = append(problemas[w.SlideID], w.Reason)
		}
	}
	var quebrados []pdfdoc.DeckSlide
	var linhas []string
	for _, sl := range slides {
		if ps, ok := problemas[sl.ID]; ok {
			quebrados = append(quebrados, sl)
			linhas = append(linhas, "slide id="+sl.ID+": "+strings.Join(ps, "; "))
		}
	}
	if len(quebrados) == 0 {
		return nil, meta, false
	}
	b, _ := json.Marshal(quebrados)
	corrigidos, m, err := ai.RepairOverflow(PlanRepairRequest{
		SlidesJSON: string(b), Excessos: strings.Join(linhas, "\n"), Model: modelo,
	})
	meta = m
	if err != nil || len(corrigidos) == 0 {
		return nil, meta, false
	}
	porID := map[string]pdfdoc.DeckSlide{}
	for _, c := range corrigidos {
		if c.ID != "" {
			porID[c.ID] = c
		}
	}
	out := make([]pdfdoc.DeckSlide, len(slides))
	copy(out, slides)
	for i := range out {
		novo, ok := porID[out[i].ID]
		if !ok {
			continue
		}
		novo.Kind, novo.ID = out[i].Kind, out[i].ID
		novo.Eyebrow, novo.Variant, novo.Legend = out[i].Eyebrow, out[i].Variant, out[i].Legend
		novo.Rulers = out[i].Rulers
		if out[i].Summary != nil {
			novo.Summary = out[i].Summary
		}
		out[i] = novo
	}
	return out, meta, true
}
