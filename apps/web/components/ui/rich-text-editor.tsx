'use client'

import { useEditor, EditorContent } from '@tiptap/react'
import StarterKit from '@tiptap/starter-kit'
import { Underline } from '@tiptap/extension-underline'
import { TextAlign } from '@tiptap/extension-text-align'
import { TextStyle } from '@tiptap/extension-text-style'
import { Color } from '@tiptap/extension-color'
import { Highlight } from '@tiptap/extension-highlight'
import {
  Bold,
  Italic,
  Underline as UnderlineIcon,
  Strikethrough,
  AlignLeft,
  AlignCenter,
  AlignRight,
  AlignJustify,
  List,
  ListOrdered,
  Undo,
  Redo,
  Heading1,
  Heading2,
  Heading3,
  Palette,
  Paintbrush,
} from 'lucide-react'
import { Button } from './button'
import { cn } from '@/lib/utils'
import { Toggle } from './toggle'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from './popover'
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from './tooltip'

interface RichTextEditorProps {
  value: string
  onChange: (html: string) => void
  placeholder?: string
  className?: string
  minHeight?: string
}

export function RichTextEditor({
  value,
  onChange,
  placeholder = 'Digite aqui...',
  className,
  minHeight = '200px',
}: RichTextEditorProps) {
  const editor = useEditor({
    immediatelyRender: false, // Fix SSR hydration issues in Next.js
    extensions: [
      StarterKit, // Already includes: bold, italic, strike, heading, paragraph, lists, etc.
      Underline, // Add underline (not in StarterKit by default)
      TextAlign.configure({
        types: ['heading', 'paragraph'],
      }),
      TextStyle,
      Color,
      Highlight.configure({
        multicolor: true,
      }),
    ],
    content: value,
    onUpdate: ({ editor }) => {
      onChange(editor.getHTML())
    },
    editorProps: {
      attributes: {
        class: cn(
          'prose prose-sm max-w-none focus:outline-none',
          'px-3 py-2',
          className
        ),
        style: `min-height: ${minHeight}`,
      },
    },
  })

  if (!editor) {
    return null
  }

  return (
    <TooltipProvider delayDuration={300}>
      <div className="border rounded-lg overflow-hidden">
        {/* Toolbar */}
        <div className="border-b bg-muted/50 p-2 flex flex-wrap gap-1">
        {/* Text formatting */}
        <div className="flex gap-1 border-r pr-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('bold')}
                onPressedChange={() => editor.chain().focus().toggleBold().run()}
                aria-label="Negrito"
              >
                <Bold className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Negrito (Ctrl+B)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('italic')}
                onPressedChange={() => editor.chain().focus().toggleItalic().run()}
                aria-label="Itálico"
              >
                <Italic className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Itálico (Ctrl+I)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('underline')}
                onPressedChange={() => editor.chain().focus().toggleUnderline().run()}
                aria-label="Sublinhado"
              >
                <UnderlineIcon className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Sublinhado (Ctrl+U)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('strike')}
                onPressedChange={() => editor.chain().focus().toggleStrike().run()}
                aria-label="Tachado"
              >
                <Strikethrough className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Tachado (Ctrl+Shift+X)</p>
            </TooltipContent>
          </Tooltip>
        </div>

        {/* Headings */}
        <div className="flex gap-1 border-r pr-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('heading', { level: 1 })}
                onPressedChange={() =>
                  editor.chain().focus().toggleHeading({ level: 1 }).run()
                }
                aria-label="Título 1"
              >
                <Heading1 className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Título 1 (Ctrl+Alt+1)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('heading', { level: 2 })}
                onPressedChange={() =>
                  editor.chain().focus().toggleHeading({ level: 2 }).run()
                }
                aria-label="Título 2"
              >
                <Heading2 className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Título 2 (Ctrl+Alt+2)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('heading', { level: 3 })}
                onPressedChange={() =>
                  editor.chain().focus().toggleHeading({ level: 3 }).run()
                }
                aria-label="Título 3"
              >
                <Heading3 className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Título 3 (Ctrl+Alt+3)</p>
            </TooltipContent>
          </Tooltip>
        </div>

        {/* Alignment */}
        <div className="flex gap-1 border-r pr-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive({ textAlign: 'left' })}
                onPressedChange={() => editor.chain().focus().setTextAlign('left').run()}
                aria-label="Alinhar à esquerda"
              >
                <AlignLeft className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Alinhar à esquerda (Ctrl+Shift+L)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive({ textAlign: 'center' })}
                onPressedChange={() => editor.chain().focus().setTextAlign('center').run()}
                aria-label="Centralizar"
              >
                <AlignCenter className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Centralizar (Ctrl+Shift+E)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive({ textAlign: 'right' })}
                onPressedChange={() => editor.chain().focus().setTextAlign('right').run()}
                aria-label="Alinhar à direita"
              >
                <AlignRight className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Alinhar à direita (Ctrl+Shift+R)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive({ textAlign: 'justify' })}
                onPressedChange={() =>
                  editor.chain().focus().setTextAlign('justify').run()
                }
                aria-label="Justificar"
              >
                <AlignJustify className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Justificar (Ctrl+Shift+J)</p>
            </TooltipContent>
          </Tooltip>
        </div>

        {/* Lists */}
        <div className="flex gap-1 border-r pr-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('bulletList')}
                onPressedChange={() => editor.chain().focus().toggleBulletList().run()}
                aria-label="Lista com marcadores"
              >
                <List className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Lista com marcadores (Ctrl+Shift+8)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Toggle
                size="sm"
                pressed={editor.isActive('orderedList')}
                onPressedChange={() => editor.chain().focus().toggleOrderedList().run()}
                aria-label="Lista numerada"
              >
                <ListOrdered className="h-4 w-4" />
              </Toggle>
            </TooltipTrigger>
            <TooltipContent>
              <p>Lista numerada (Ctrl+Shift+7)</p>
            </TooltipContent>
          </Tooltip>
        </div>

        {/* Colors */}
        <div className="flex gap-1 border-r pr-2">
          {/* Text Color */}
          <Popover>
            <Tooltip>
              <TooltipTrigger asChild>
                <PopoverTrigger asChild>
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    aria-label="Cor do texto"
                  >
                    <Palette className="h-4 w-4" />
                  </Button>
                </PopoverTrigger>
              </TooltipTrigger>
              <TooltipContent>
                <p>Cor do texto</p>
              </TooltipContent>
            </Tooltip>
            <PopoverContent className="w-auto p-2">
              <div className="space-y-2">
                <p className="text-xs font-medium mb-2">Cor do Texto</p>
                <div className="grid grid-cols-6 gap-1">
                  {[
                    { color: '#000000', label: 'Preto' },
                    { color: '#DC2626', label: 'Vermelho' },
                    { color: '#EA580C', label: 'Laranja' },
                    { color: '#CA8A04', label: 'Amarelo' },
                    { color: '#16A34A', label: 'Verde' },
                    { color: '#2563EB', label: 'Azul' },
                    { color: '#9333EA', label: 'Roxo' },
                    { color: '#DB2777', label: 'Rosa' },
                    { color: '#64748B', label: 'Cinza' },
                    { color: '#0891B2', label: 'Ciano' },
                    { color: '#84CC16', label: 'Lima' },
                    { color: '#F59E0B', label: 'Âmbar' },
                  ].map(({ color, label }) => (
                    <button
                      key={color}
                      type="button"
                      className="w-6 h-6 rounded border-2 border-border hover:border-primary transition-colors"
                      style={{ backgroundColor: color }}
                      onClick={() => editor.chain().focus().setColor(color).run()}
                      title={label}
                    />
                  ))}
                </div>
                <Button
                  type="button"
                  variant="outline"
                  size="sm"
                  className="w-full"
                  onClick={() => editor.chain().focus().unsetColor().run()}
                >
                  Remover Cor
                </Button>
              </div>
            </PopoverContent>
          </Popover>

          {/* Background Color (Highlight) */}
          <Popover>
            <Tooltip>
              <TooltipTrigger asChild>
                <PopoverTrigger asChild>
                  <Button
                    type="button"
                    variant="ghost"
                    size="sm"
                    aria-label="Cor do fundo"
                  >
                    <Paintbrush className="h-4 w-4" />
                  </Button>
                </PopoverTrigger>
              </TooltipTrigger>
              <TooltipContent>
                <p>Cor do fundo (destaque)</p>
              </TooltipContent>
            </Tooltip>
            <PopoverContent className="w-auto p-2">
              <div className="space-y-2">
                <p className="text-xs font-medium mb-2">Cor do Fundo</p>
                <div className="grid grid-cols-6 gap-1">
                  {[
                    { color: '#FEF08A', label: 'Amarelo' },
                    { color: '#FED7AA', label: 'Laranja' },
                    { color: '#FECACA', label: 'Vermelho' },
                    { color: '#BBF7D0', label: 'Verde' },
                    { color: '#BFDBFE', label: 'Azul' },
                    { color: '#DDD6FE', label: 'Roxo' },
                    { color: '#FBCFE8', label: 'Rosa' },
                    { color: '#A5F3FC', label: 'Ciano' },
                    { color: '#E0E7FF', label: 'Índigo' },
                    { color: '#FEF3C7', label: 'Âmbar' },
                    { color: '#D9F99D', label: 'Lima' },
                    { color: '#E5E7EB', label: 'Cinza' },
                  ].map(({ color, label }) => (
                    <button
                      key={color}
                      type="button"
                      className="w-6 h-6 rounded border-2 border-border hover:border-primary transition-colors"
                      style={{ backgroundColor: color }}
                      onClick={() =>
                        editor.chain().focus().toggleHighlight({ color }).run()
                      }
                      title={label}
                    />
                  ))}
                </div>
                <Button
                  type="button"
                  variant="outline"
                  size="sm"
                  className="w-full"
                  onClick={() => editor.chain().focus().unsetHighlight().run()}
                >
                  Remover Destaque
                </Button>
              </div>
            </PopoverContent>
          </Popover>
        </div>

        {/* Undo/Redo */}
        <div className="flex gap-1">
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                type="button"
                variant="ghost"
                size="sm"
                onClick={() => editor.chain().focus().undo().run()}
                disabled={!editor.can().undo()}
                aria-label="Desfazer"
              >
                <Undo className="h-4 w-4" />
              </Button>
            </TooltipTrigger>
            <TooltipContent>
              <p>Desfazer (Ctrl+Z)</p>
            </TooltipContent>
          </Tooltip>

          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                type="button"
                variant="ghost"
                size="sm"
                onClick={() => editor.chain().focus().redo().run()}
                disabled={!editor.can().redo()}
                aria-label="Refazer"
              >
                <Redo className="h-4 w-4" />
              </Button>
            </TooltipTrigger>
            <TooltipContent>
              <p>Refazer (Ctrl+Y)</p>
            </TooltipContent>
          </Tooltip>
        </div>
      </div>

      {/* Editor */}
      <EditorContent editor={editor} className="bg-background" />
      </div>
    </TooltipProvider>
  )
}
