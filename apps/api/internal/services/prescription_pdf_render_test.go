package services

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
)

// O padrão da receita é o do Dr. Getúlio, e é uma frase, não uma lista de campos:
//
//  1. Losartana 50mg ................. uso contínuo
//     Tomar 1 comprimido uma vez ao dia
//
// A duração e a quantidade saem da posologia e vão para o campo da direita; a linha de baixo é a
// instrução que o PACIENTE lê, com verbo.
func TestMedPosologyEhFraseComVerbo(t *testing.T) {
	casos := []struct {
		nome string
		med  models.PrescriptionMedication
		quer string
	}{
		{"oral omite a via, que o leitor já assume",
			models.PrescriptionMedication{Dosage: "1 comprimido", Frequency: "uma vez ao dia", Route: "oral"},
			"Tomar 1 comprimido uma vez ao dia"},
		{"via em branco NÃO é oral: o PDF não pode inventar 'tomar' numa seringa",
			models.PrescriptionMedication{Dosage: "1 seringa", Frequency: "de 12 em 12 horas"},
			"Usar 1 seringa de 12 em 12 horas"},
		{"registro antigo sem dosagem nem frequência ainda diz a via",
			models.PrescriptionMedication{Route: "subcutânea"}, "Via subcutânea"},
		{"sem nada, nada — melhor branco que frase inventada",
			models.PrescriptionMedication{}, ""},
		{"sublingual é tomar, mas a via aparece",
			models.PrescriptionMedication{Dosage: "1 comprimido", Frequency: "ao deitar", Route: "sublingual"},
			"Tomar 1 comprimido ao deitar, via sublingual"},
		{"injetável não se toma",
			models.PrescriptionMedication{Dosage: "1 seringa", Frequency: "uma vez ao dia", Route: "subcutânea"},
			"Aplicar 1 seringa uma vez ao dia, via subcutânea"},
		{"colírio não se toma",
			models.PrescriptionMedication{Dosage: "1 gota", Frequency: "de 8 em 8 horas", Route: "oftálmica"},
			"Aplicar 1 gota de 8 em 8 horas, via oftálmica"},
		{"duração sozinha NÃO entra na frase: ela ocupa o campo da direita",
			models.PrescriptionMedication{Dosage: "1 comprimido", Frequency: "uma vez ao dia", Route: "oral", Duration: 30},
			"Tomar 1 comprimido uma vez ao dia"},
		{"com quantidade, o campo da direita é dela e a duração volta para a frase — senão o prazo some do papel",
			models.PrescriptionMedication{Dosage: "1 comprimido", Frequency: "uma vez ao dia", Route: "oral", Duration: 30, Quantity: 60},
			"Tomar 1 comprimido uma vez ao dia, por 30 dias"},
		{"quantidade e duração numa injetável: via e prazo, nesta ordem",
			models.PrescriptionMedication{Dosage: "1 seringa", Frequency: "uma vez ao dia", Route: "subcutânea", Duration: 7, Quantity: 7},
			"Aplicar 1 seringa uma vez ao dia, via subcutânea, por 7 dias"},
		{"quantidade SEM duração: o 'uso contínuo' migra para a frase, senão some do papel",
			models.PrescriptionMedication{Dosage: "1 comprimido", Frequency: "uma vez ao dia", Route: "oral", Quantity: 30},
			"Tomar 1 comprimido uma vez ao dia, uso contínuo"},
	}
	for _, c := range casos {
		if got := medPosology(c.med); got != c.quer {
			t.Errorf("%s:\n  quer %q\n  veio %q", c.nome, c.quer, got)
		}
	}
}

// O campo da direita é o que a farmácia lê. Em branco ele MENTE por omissão: quem lê não sabe se o
// médico esqueceu o prazo ou se não há prazo.
func TestMedDispenseNuncaFicaVazio(t *testing.T) {
	casos := []struct {
		nome string
		med  models.PrescriptionMedication
		quer string
	}{
		{"sem duração e sem quantidade é uso contínuo",
			models.PrescriptionMedication{}, "uso contínuo"},
		{"quantidade manda, com o extenso entre parênteses",
			models.PrescriptionMedication{Quantity: 60, QuantityInWords: "sessenta comprimidos", Duration: 30},
			"60 (sessenta comprimidos)"},
		{"quantidade sem extenso leva rótulo: um '30' pelado lê como prazo",
			models.PrescriptionMedication{Quantity: 30}, "Quantidade: 30"},
		{"só duração vira o prazo",
			models.PrescriptionMedication{Duration: 7}, "por 7 dias"},
		{"um dia é singular",
			models.PrescriptionMedication{Duration: 1}, "por 1 dia"},
	}
	for _, c := range casos {
		if got := medDispense(c.med); got != c.quer {
			t.Errorf("%s:\n  quer %q\n  veio %q", c.nome, c.quer, got)
		}
	}
}

// Manipulado segue a mesma regra do prazo: sem duração, a receita DIZ que é contínuo.
func TestFormulaPosologyDizUsoContinuo(t *testing.T) {
	f := models.PrescriptionFormula{Posology: "Tomar 1 cápsula ao dia", Route: "oral"}
	if got := formulaPosologyLine(f); got != "Tomar 1 cápsula ao dia · via oral · uso contínuo" {
		t.Errorf("sem duração: veio %q", got)
	}
	f.Duration = 1
	if got := formulaPosologyLine(f); got != "Tomar 1 cápsula ao dia · via oral · por 1 dia" {
		t.Errorf("um dia: veio %q", got)
	}
}

// O rascunho vai para a mesma pasta de downloads da receita assinada, e no mesmo dia. Se os dois
// nomes forem iguais, o segundo sobrescreve o primeiro — e o que sobra pode ser o rascunho.
func TestNomeDoRascunhoSeparaDaReceitaAssinada(t *testing.T) {
	p := models.Prescription{
		ID:               uuid.MustParse("01a07670-3e4b-7e31-9015-c132b5ba60e0"),
		Type:             models.PrescriptionCommercial,
		PrescriptionDate: time.Date(2026, 9, 6, 0, 0, 0, 0, time.UTC),
		Patient:          models.Patient{Name: "Elisane Albuquerque"},
	}
	rascunho := prescriptionDraftFileName(&p)
	assinada := prescriptionFileName(&p)
	if rascunho == assinada {
		t.Fatalf("rascunho e assinada com o mesmo nome: %q", rascunho)
	}
	if want := "Elisane-Albuquerque_RascunhoReceita_2026-09-06_01a07670.pdf"; rascunho != want {
		t.Errorf("rascunho:\n  quer %q\n  veio %q", want, rascunho)
	}
	p.Type = models.PrescriptionCompounded
	if want := "Elisane-Albuquerque_RascunhoReceitaManipulado_2026-09-06_01a07670.pdf"; prescriptionDraftFileName(&p) != want {
		t.Errorf("manipulado:\n  quer %q\n  veio %q", want, prescriptionDraftFileName(&p))
	}
}
