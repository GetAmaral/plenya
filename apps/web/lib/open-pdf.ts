import { apiClient } from '@/lib/api-client'

/**
 * Entrega ao usuário um PDF que veio de rota autenticada, COM O NOME que o servidor mandou.
 *
 * Todo documento do EMR sai do sistema: o médico reenvia pelo WhatsApp, o paciente guarda na pasta
 * de downloads. O servidor já monta o nome no padrão da casa
 * ("Ana-Cláudia_Receita_2026-08-31_01a0592b.pdf", ver utils.DocumentFileName no back), mas o nome
 * viaja no Content-Disposition — e um `URL.createObjectURL(blob)` aberto com `window.open` joga
 * esse cabeçalho fora. O arquivo então chega ao WhatsApp e à pasta de downloads como
 * **unknown.pdf**, todos iguais, um sobrescrevendo o outro.
 *
 * Por isso o blob vira um `File` nomeado e a entrega é por `<a download>`. No celular, quando o
 * aparelho sabe compartilhar arquivo, abre a folha de compartilhamento com o File — é o caminho
 * que leva o nome certo direto para o WhatsApp, sem passar pela pasta de downloads.
 *
 * O nome só chega ao JS porque a API expõe o cabeçalho no CORS (ExposeHeaders:
 * "Content-Disposition", em cmd/server/main.go). Se algum dia isso sair de lá, todo download volta
 * silenciosamente para o nome genérico — daí o fallback obrigatório.
 */
export async function openServerPdf(endpoint: string, fallbackName: string): Promise<void> {
  const { blob, filename } = await apiClient.getBlobWithName(endpoint)
  const name = filename || fallbackName
  const file = new File([blob], name, { type: blob.type || 'application/pdf' })

  // Folha de compartilhamento só no aparelho de toque. `canShare` sozinho não serve de teste:
  // o Chrome no Windows também compartilha arquivo, e aí um botão escrito "Baixar" abria a folha
  // do sistema. Pior, abria de forma imprevisível — `navigator.share` exige gesto recente do
  // usuário, e o download do blob vem antes: num laudo de 13 MB o gesto expira e o mesmo clique
  // ora compartilha, ora baixa, conforme a rede. No celular a folha é o ponto: é ela que entrega
  // o arquivo nomeado direto ao WhatsApp, sem passar pela pasta de downloads.
  const aparelhoDeToque =
    typeof matchMedia === 'function' && matchMedia('(pointer: coarse)').matches
  const canShareFile =
    aparelhoDeToque &&
    typeof navigator !== 'undefined' &&
    typeof navigator.canShare === 'function' &&
    navigator.canShare({ files: [file] })

  if (canShareFile) {
    try {
      await navigator.share({ files: [file], title: name })
      return
    } catch (e: any) {
      // Usuário fechou a folha de compartilhamento: não é erro, e não abre aba por cima.
      if (e?.name === 'AbortError') return
    }
  }

  const url = URL.createObjectURL(file)
  const a = document.createElement('a')
  a.href = url
  a.download = name
  a.target = '_blank'
  a.rel = 'noopener'
  // O Firefox só dispara o download de uma âncora que esteja no documento; solta, o clique é
  // ignorado em silêncio e o botão não faz nada.
  document.body.appendChild(a)
  a.click()
  a.remove()
  // Revogar na hora corta o download no meio em alguns navegadores.
  setTimeout(() => URL.revokeObjectURL(url), 60_000)
}
