package services

import (
	"errors"
	"sort"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

// Quando o paciente fez cada exame — e a resposta certa depende de olhar os analitos FILHOS.
//
// Um painel do catálogo quase nunca tem resultado próprio: o laboratório não reporta "hemograma
// completo", reporta hemoglobina, leucócitos, plaquetas. Medido neste banco, "Hemograma completo"
// tem **0 resultados próprios e 336 nos filhos**; perfil lipídico, 0 e 91. Quem cruza o protocolo
// com o prontuário olhando só o painel conclui "nunca feito" e manda repetir o exame de quem
// acabou de fazer — já aconteceu, e era o motivo de o cruzamento ser feito à mão com uma regra
// decorada.
//
// O inverso também existe: "Gasometria venosa" tem 16 resultados próprios e nenhum nos filhos.
// Então a regra é a união — o painel conta como feito pela data mais recente entre o que ele mesmo
// tem e o que qualquer filho tem.

// LabCoverageService responde "quando isto foi feito" para cada exame do catálogo. Só lê.
type LabCoverageService struct {
	db *gorm.DB
}

func NewLabCoverageService(db *gorm.DB) *LabCoverageService {
	return &LabCoverageService{db: db}
}

// Build monta a cobertura do paciente. `onlyRequestable` limita ao que dá para pedir; `doneOnly`
// devolve só o que já foi feito (o catálogo inteiro são 528 entradas e ~57KB, e quem consome usa
// as ~90 feitas).
func (s *LabCoverageService) Build(patientID uuid.UUID, onlyRequestable, doneOnly bool) (*dto.LabCoverageResponse, error) {
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	defs, err := s.loadDefinitions()
	if err != nil {
		return nil, err
	}
	lastByDef, err := s.loadLastDates(patientID)
	if err != nil {
		return nil, err
	}

	// O mapa pai→filhos é montado com TODAS as definições, inclusive as inativas e as apagadas em
	// soft delete: um analito desativado no catálogo depois da coleta não desfaz a coleta, e
	// ignorá-lo devolveria o falso "nunca feito" que este serviço existe para matar.
	filhos := map[uuid.UUID][]uuid.UUID{}
	for _, d := range defs {
		if d.ParentID != nil {
			filhos[*d.ParentID] = append(filhos[*d.ParentID], d.ID)
		}
	}

	out := &dto.LabCoverageResponse{
		PatientID:   patientID.String(),
		Entries:     []dto.LabCoverageEntry{},
		GeneratedAt: time.Now().In(saoPaulo()).Format(time.RFC3339),
	}
	hoje := time.Now().In(saoPaulo())

	for _, d := range defs {
		// Recorte de SAÍDA (a árvore acima já usou todas as definições).
		if d.Apagado || !d.IsActive {
			continue
		}
		if onlyRequestable && !d.IsRequestable {
			continue
		}
		e := dto.LabCoverageEntry{Code: d.Code, Name: d.Name, Via: dto.LabCoverageNever}

		proprio, temProprio := lastByDef[d.ID]

		// Descendentes, não só filhos diretos: um catálogo de três níveis (painel → subgrupo →
		// analito) resolveria como "nunca feito" se olhássemos um nível só.
		desc := descendentes(filhos, d.ID)
		var maisRecenteFilho time.Time
		feitosEver := 0
		for _, fid := range desc {
			if dt, ok := lastByDef[fid]; ok {
				feitosEver++
				if dt.After(maisRecenteFilho) {
					maisRecenteFilho = dt
				}
			}
		}
		e.ChildrenTotal = len(desc)
		e.ChildrenDoneEver = feitosEver

		quando, via := resolveCoverage(proprio, temProprio, maisRecenteFilho)
		e.Via = via

		if !quando.IsZero() {
			e.LastDoneAt = collectionDay(quando)
			dias := int(hoje.Sub(quando).Hours() / 24)
			if dias < 0 {
				dias = 0
			}
			e.DaysAgo = &dias
			// Quantos analitos vieram NESTA coleta. Comparado por dia-calendário: a data do lote e
			// a do resultado são gravadas de formas diferentes.
			for _, fid := range desc {
				if dt, ok := lastByDef[fid]; ok && collectionDay(dt) == e.LastDoneAt {
					e.ChildrenDone++
				}
			}
		}

		if doneOnly && e.Via == dto.LabCoverageNever {
			continue
		}
		out.Entries = append(out.Entries, e)
	}

	sort.Slice(out.Entries, func(i, j int) bool { return out.Entries[i].Name < out.Entries[j].Name })
	return out, nil
}

// resolveCoverage decide a data e a origem de um exame.
//
// A união é o ponto: painel quase sempre só tem dado nos filhos (hemograma, 0 próprios e 336 nos
// filhos), mas há o inverso (gasometria venosa, 16 próprios e nenhum filho). Olhar só um dos lados
// erra metade do catálogo, e o erro que importa é o falso "nunca feito", que manda repetir exame de
// quem acabou de fazer.
func resolveCoverage(proprio time.Time, temProprio bool, maisRecenteFilho time.Time) (time.Time, dto.LabCoverageVia) {
	temFilho := !maisRecenteFilho.IsZero()
	switch {
	case temProprio && temFilho:
		// Os dois existem: vale a coleta mais recente, e a origem é quem deu a data.
		if proprio.After(maisRecenteFilho) {
			return proprio, dto.LabCoverageOwn
		}
		return maisRecenteFilho, dto.LabCoverageChildren
	case temProprio:
		return proprio, dto.LabCoverageOwn
	case temFilho:
		return maisRecenteFilho, dto.LabCoverageChildren
	}
	return time.Time{}, dto.LabCoverageNever
}

type coverageDef struct {
	ID            uuid.UUID
	ParentID      *uuid.UUID
	Code          string
	Name          string
	IsRequestable bool
	IsActive      bool
	Apagado       bool
}

// descendentes devolve toda a descendência de um exame, em qualquer profundidade. A guarda de
// visitados existe porque `parent_test_id` é um campo livre: um ciclo no catálogo travaria o
// serviço em laço infinito em vez de dar erro.
func descendentes(filhos map[uuid.UUID][]uuid.UUID, raiz uuid.UUID) []uuid.UUID {
	var out []uuid.UUID
	visto := map[uuid.UUID]bool{raiz: true}
	fila := append([]uuid.UUID(nil), filhos[raiz]...)
	for len(fila) > 0 {
		id := fila[0]
		fila = fila[1:]
		if visto[id] {
			continue
		}
		visto[id] = true
		out = append(out, id)
		fila = append(fila, filhos[id]...)
	}
	return out
}

// loadDefinitions traz TODAS as definições, inclusive inativas e apagadas: elas não entram na
// resposta, mas são necessárias para montar a árvore pai→filho. O recorte de saída é feito depois.
func (s *LabCoverageService) loadDefinitions() ([]coverageDef, error) {
	var rows []coverageDef
	err := s.db.
		Table("lab_test_definitions").
		Select("id, parent_test_id AS parent_id, code, name, is_requestable, is_active, deleted_at IS NOT NULL AS apagado").
		Scan(&rows).Error
	return rows, err
}

// loadLastDates devolve, por definição de exame, a data mais recente em que ESTE paciente tem
// resultado. Não filtra status do lote: um laudo recém-importado, ainda sem revisão, já é prova de
// que o exame foi feito — e é justamente ele que evita o pedido repetido.
func (s *LabCoverageService) loadLastDates(patientID uuid.UUID) (map[uuid.UUID]time.Time, error) {
	var rows []struct {
		DefID uuid.UUID
		Last  time.Time
	}
	err := s.db.
		Table("lab_results AS lr").
		Select(`lr.lab_test_definition_id AS def_id,
		        MAX(COALESCE(lr.collection_date, b.collection_date)) AS last`).
		Joins("JOIN lab_result_batches b ON b.id = lr.lab_result_batch_id AND b.deleted_at IS NULL").
		Where("b.patient_id = ? AND lr.deleted_at IS NULL AND lr.lab_test_definition_id IS NOT NULL", patientID).
		Group("lr.lab_test_definition_id").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make(map[uuid.UUID]time.Time, len(rows))
	for _, r := range rows {
		out[r.DefID] = r.Last
	}
	return out, nil
}
