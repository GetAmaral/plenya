// Package migrations embeda os arquivos .sql de migration goose no binário,
// para que `cmd/migrate` os aplique sem depender de arquivos no disco em runtime.
package migrations

import "embed"

// FS contém as migrations goose deste diretório (não inclui _legacy/).
//
//go:embed *.sql
var FS embed.FS
