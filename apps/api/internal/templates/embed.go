// Package templates embeda os assets do pôster de escore no binário, para que a geração
// de PDF não dependa de arquivos no disco em runtime: a imagem de prod copia só o binário,
// sem a árvore de fontes. Em dev o bind-mount de ./apps/api em /app mascarava isso.
package templates

import "embed"

// FS contém o template do pôster e o logo referenciado por ele (o .png não é usado).
//
//go:embed score_poster.html logo_infinity.svg
var FS embed.FS
