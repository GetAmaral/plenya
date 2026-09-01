package handlers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Id inválido no path tem que dar 400 "id inválido", nunca 404 "não encontrado".
//
// A armadilha é do Fiber: `c.Status(400).JSON(...)` devolve **nil** quando consegue escrever a
// resposta. Testar esse retorno com `if err != nil` — que é o jeito que parece óbvio — deixa o
// handler seguir em frente com uuid.Nil depois de já ter escrito o 400, e o cliente recebe um 404
// enganoso dizendo que o plano não existe. Aconteceu, e este teste existe para não voltar.
func TestPatientPlanIDsRecusaIDInvalidoCom400(t *testing.T) {
	h := &PatientPlanHandler{}
	app := fiber.New()
	app.Get("/p/:id/plans/:planId", func(c *fiber.Ctx) error {
		_, _, resp, ok := h.ids(c)
		if !ok {
			return resp
		}
		return c.JSON(fiber.Map{"ok": true})
	})

	valido := uuid.Must(uuid.NewV7()).String()
	cases := []struct {
		nome     string
		path     string
		status   int
		contains string
	}{
		{"planId não é uuid", "/p/" + valido + "/plans/nao-e-uuid", fiber.StatusBadRequest, "invalid plan id"},
		{"patientId não é uuid", "/p/xxx/plans/" + valido, fiber.StatusBadRequest, "invalid patient id"},
		{"os dois válidos", "/p/" + valido + "/plans/" + valido, fiber.StatusOK, "ok"},
	}
	for _, tc := range cases {
		t.Run(tc.nome, func(t *testing.T) {
			res, err := app.Test(httptest.NewRequest(fiber.MethodGet, tc.path, nil))
			if err != nil {
				t.Fatalf("request: %v", err)
			}
			defer res.Body.Close()
			if res.StatusCode != tc.status {
				t.Errorf("status = %d, quer %d", res.StatusCode, tc.status)
			}
			body, _ := io.ReadAll(res.Body)
			var payload map[string]any
			if err := json.Unmarshal(body, &payload); err != nil {
				t.Fatalf("resposta não é JSON: %s", body)
			}
			// Uma resposta só: escrever 400 e depois seguir emendaria dois JSONs no mesmo corpo.
			if len(body) == 0 {
				t.Fatal("corpo vazio")
			}
			found := false
			for _, v := range payload {
				if s, isStr := v.(string); isStr && s == tc.contains {
					found = true
				}
				if b, isBool := v.(bool); isBool && b && tc.contains == "ok" {
					found = true
				}
			}
			if !found {
				t.Errorf("resposta %s não traz %q", body, tc.contains)
			}
		})
	}
}
