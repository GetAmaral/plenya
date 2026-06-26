package pdfdoc

import "strings"

// ---- Relatório longitudinal AGIR (documento que flui em múltiplas páginas) ----

// ReportBiomarker — um item do escore com seu nível.
type ReportBiomarker struct {
	Name  string
	Label string // "Crítico" / "Subótimo" / "Ótimo" ...
	Kind  string // "critical" | "attention" | "good" | "optimal"
}

// ReportRec — recomendação do plano (com meta opcional).
type ReportRec struct {
	Text   string
	Target string // "" se não houver meta
}

// ReportPillar — pilar AGIR com suas recomendações.
type ReportPillar struct {
	Letter string // A G I R
	Name   string // Atividade / Gestão / Integração / Ritmo
	Recs   []ReportRec
}

// CarePlanReport — dados do relatório.
type CarePlanReport struct {
	Title       string // "Relatório de Saúde, Performance e Longevidade"
	Patient     Patient
	EmittedAt   string
	ScoreCalcAt string
	ScorePct    string // "78%"
	Attention   []ReportBiomarker
	Optimal     []ReportBiomarker
	Pillars     []ReportPillar
	Doctor      Doctor
	Signature   Signature
	Clinic      Clinic
}

// cor semântica do nível clínico (necessária para o significado; tons sóbrios harmonizados à marca).
var levelKindColor = map[string]string{
	"critical":  "#a8392b", // terracota
	"attention": "#b38645", // gold
	"good":      "#417e8e", // ocean
	"optimal":   "#3f7d5e", // sage-verde
}

// reportCSS — SÓ o que é específico do relatório (escore, biomarcadores, plano AGIR). Cabeçalho,
// rodapé, marca d'água, título e paginação vêm do motor único (renderDocument).
const reportCSS = `
.rpt-dates{ font-size:12px; color:var(--ink2); margin-top:6px; }
/* escore */
.score{ display:flex; align-items:baseline; gap:16px; margin-top:15px; padding:12px 20px;
        background:rgba(6,59,79,.04); border:1px solid rgba(179,134,69,.4); border-radius:4px; }
.score .pct{ font-family:'Cormorant',serif; font-size:40px; font-weight:600; color:var(--petrol); line-height:1; }
.score .lbl{ font-size:12.5px; color:var(--ink2); }
.score .lbl b{ color:var(--petrol); font-weight:600; }
/* seções */
.rsection{ margin-top:15px; break-inside:avoid; }
.rsection .eyebrow{ display:block; margin-bottom:9px; }
.bmrow{ display:flex; align-items:center; justify-content:space-between; padding:3.5px 0;
        border-bottom:1px solid rgba(6,59,79,.07); font-size:12.5px; color:var(--petrol); }
.bmrow:first-child{ border-top:1px solid rgba(6,59,79,.07); }
.bmbadge{ font-size:10px; font-weight:700; letter-spacing:.4px; text-transform:uppercase;
          color:#fff; padding:2px 9px; border-radius:10px; white-space:nowrap; }
/* plano AGIR */
.pillar{ margin-top:11px; break-inside:avoid; }
.pillar-h{ display:flex; align-items:center; gap:9px; }
.pillar-dot{ width:22px; height:22px; border-radius:5px; background:var(--petrol); color:var(--cream);
             font-size:12px; font-weight:700; display:flex; align-items:center; justify-content:center; }
.pillar-name{ font-size:14px; font-weight:600; color:var(--petrol); }
.prec{ display:flex; align-items:flex-start; gap:11px; padding:3px 0 3px 31px; font-size:12.5px; color:var(--petrol); }
.prec .bullet{ flex:0 0 auto; width:4.5px; height:4.5px; background:var(--gold); border-radius:1px; transform:rotate(45deg); margin-top:6px; }
.prec .target{ color:var(--ink2); }
/* rodapé do relatório */
.rfoot{ margin-top:18px; }
`

func reportBmRows(rows []ReportBiomarker) string {
	var b strings.Builder
	for _, r := range rows {
		col := levelKindColor[r.Kind]
		if col == "" {
			col = "#6b7280"
		}
		b.WriteString(`<div class="bmrow"><span>` + esc(r.Name) + `</span>` +
			`<span class="bmbadge" style="background:` + col + `">` + esc(r.Label) + `</span></div>`)
	}
	return b.String()
}

// RenderCarePlanReport gera o relatório longitudinal AGIR pelo motor único. Só monta o miolo
// (datas, escore, biomarcadores, plano AGIR); cabeçalho/rodapé/marca/paginação vêm do renderDocument.
func RenderCarePlanReport(in CarePlanReport) ([]byte, error) {
	if in.Title == "" {
		in.Title = "Relatório de Saúde, Performance e Longevidade"
	}

	var b strings.Builder
	// datas (bloco)
	meta := []string{}
	if in.EmittedAt != "" {
		meta = append(meta, "Relatório emitido em "+esc(in.EmittedAt))
	}
	if in.ScoreCalcAt != "" {
		meta = append(meta, "Escore calculado em "+esc(in.ScoreCalcAt))
	}
	if len(meta) > 0 {
		b.WriteString(`<div class="rpt-dates">` + strings.Join(meta, " · ") + `</div>`)
	}
	// escore (bloco)
	if in.ScorePct != "" {
		b.WriteString(`<div class="score"><span class="pct">` + esc(in.ScorePct) + `</span>` +
			`<span class="lbl"><b>Escore Plenya</b><br>saúde global</span></div>`)
	}
	// biomarcadores (cada seção é um bloco)
	if len(in.Attention) > 0 {
		b.WriteString(`<div class="rsection"><span class="eyebrow">Pontos de atenção</span>` + reportBmRows(in.Attention) + `</div>`)
	}
	if len(in.Optimal) > 0 {
		b.WriteString(`<div class="rsection"><span class="eyebrow">No ótimo</span>` + reportBmRows(in.Optimal) + `</div>`)
	}
	// plano AGIR: cabeçalho de seção + cada pilar como bloco (paginável)
	hasPillars := false
	for _, p := range in.Pillars {
		if len(p.Recs) > 0 {
			hasPillars = true
			break
		}
	}
	if hasPillars {
		b.WriteString(`<div class="sec"><span class="eyebrow">Plano de cuidado AGIR</span></div>`)
		for _, p := range in.Pillars {
			if len(p.Recs) == 0 {
				continue
			}
			b.WriteString(`<div class="pillar"><div class="pillar-h"><span class="pillar-dot">` + esc(p.Letter) +
				`</span><span class="pillar-name">` + esc(p.Name) + `</span></div>`)
			for _, rec := range p.Recs {
				b.WriteString(`<div class="prec"><span class="bullet"></span><span>` + esc(rec.Text))
				if rec.Target != "" {
					b.WriteString(` <span class="target">(meta: ` + esc(rec.Target) + `)</span>`)
				}
				b.WriteString(`</span></div>`)
			}
			b.WriteString(`</div>`)
		}
	}

	return renderDocument(Doc{
		Title:     in.Title,
		Patient:   &in.Patient,
		Body:      b.String(),
		Signature: signatureBlock(in.Doctor, in.Signature),
		ExtraCSS:  reportCSS,
		Clinic:    in.Clinic,
	})
}
