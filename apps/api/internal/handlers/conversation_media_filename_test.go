package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/services"
)

// O nome do PDF que o paciente manda pelo WhatsApp vem com acento ("Laboratório.pdf").
// O header precisa levar o ASCII como fallback E o UTF-8 real, senão o exame chega ao
// médico como "Laboratrio.pdf".
func TestSetContentDisposition(t *testing.T) {
	cases := []struct {
		name    string
		res     services.ActivityMediaResult
		want    string
		wantSet bool
	}{
		{
			name:    "acento preservado no filename*",
			res:     services.ActivityMediaResult{Filename: "Laboratrio.pdf", DisplayFilename: "Laboratório.pdf"},
			want:    `inline; filename="Laboratrio.pdf"; filename*=UTF-8''Laborat%C3%B3rio.pdf`,
			wantSet: true,
		},
		{
			name:    "sem acento não duplica o parâmetro",
			res:     services.ActivityMediaResult{Filename: "exame.pdf", DisplayFilename: "exame.pdf"},
			want:    `inline; filename="exame.pdf"`,
			wantSet: true,
		},
		{
			name: "ponto-e-vírgula e quebra de linha não escapam do valor",
			res: services.ActivityMediaResult{
				Filename:        "a.pdf",
				DisplayFilename: "a;b\r\nX-Injected: 1.pdf",
			},
			want:    `inline; filename="a.pdf"; filename*=UTF-8''a%3Bb%0D%0AX-Injected%3A%201.pdf`,
			wantSet: true,
		},
		{
			name: "caminho no nome vira só o arquivo",
			res: services.ActivityMediaResult{
				Filename:        "passwd",
				DisplayFilename: "../../etc/passwd",
			},
			want:    `inline; filename="passwd"`,
			wantSet: true,
		},
		{name: "sem nome, sem header", res: services.ActivityMediaResult{}, wantSet: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/x", func(c *fiber.Ctx) error {
				res := tc.res
				setContentDisposition(c, &res)
				return c.SendString("ok")
			})
			resp, err := app.Test(httptest.NewRequest("GET", "/x", nil))
			if err != nil {
				t.Fatalf("request: %v", err)
			}
			got := resp.Header.Get("Content-Disposition")
			if !tc.wantSet {
				if got != "" {
					t.Fatalf("esperava header ausente, veio %q", got)
				}
				return
			}
			if got != tc.want {
				t.Fatalf("Content-Disposition\n  veio:     %q\n  esperado: %q", got, tc.want)
			}
		})
	}
}
