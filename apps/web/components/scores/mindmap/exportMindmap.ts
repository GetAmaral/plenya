import { toPng } from 'html-to-image'
import { toast } from 'sonner'

export async function exportMindmapToPNG() {
  const reactFlowElement = document.querySelector('.react-flow') as HTMLElement

  if (!reactFlowElement) {
    toast.error('Elemento do mindmap não encontrado')
    return
  }

  try {
    toast.loading('Gerando imagem...', { id: 'export-mindmap' })

    const dataUrl = await toPng(reactFlowElement, {
      backgroundColor: '#ffffff',
      filter: (node) => {
        // Exclude controls, minimap, and legend from export
        if (
          node.classList?.contains('react-flow__controls') ||
          node.classList?.contains('react-flow__minimap') ||
          node.classList?.contains('react-flow__panel')
        ) {
          return false
        }
        return true
      },
      pixelRatio: 2, // Higher quality
    })

    // Download the image
    const link = document.createElement('a')
    link.download = `mindmap-escores-${new Date().toISOString().split('T')[0]}.png`
    link.href = dataUrl
    link.click()

    toast.success('Imagem exportada com sucesso!', { id: 'export-mindmap' })
  } catch (error) {
    console.error('Error exporting mindmap:', error)
    toast.error('Erro ao exportar imagem', { id: 'export-mindmap' })
  }
}
