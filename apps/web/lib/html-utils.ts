/**
 * HTML utility functions for text processing
 */

/**
 * Strips all HTML tags from a string and returns plain text
 * @param html - HTML string to strip
 * @returns Plain text without HTML tags
 */
export function stripHtmlTags(html: string): string {
  if (!html) return ''

  // Remove HTML tags
  let text = html.replace(/<[^>]*>/g, '')

  // Decode HTML entities
  text = text
    .replace(/&nbsp;/g, ' ')
    .replace(/&lt;/g, '<')
    .replace(/&gt;/g, '>')
    .replace(/&amp;/g, '&')
    .replace(/&quot;/g, '"')
    .replace(/&#39;/g, "'")

  // Normalize whitespace
  text = text.replace(/\s+/g, ' ').trim()

  return text
}

/**
 * Extracts plain text from HTML preserving some structure
 * Converts headings, paragraphs, and lists to readable plain text
 * @param html - HTML string to convert
 * @returns Structured plain text
 */
export function htmlToPlainText(html: string): string {
  if (!html) return ''

  let text = html

  // Convert headings to text with newlines
  text = text.replace(/<h[1-6][^>]*>(.*?)<\/h[1-6]>/gi, '\n\n$1\n\n')

  // Convert paragraphs to text with newlines
  text = text.replace(/<p[^>]*>(.*?)<\/p>/gi, '$1\n\n')

  // Convert list items
  text = text.replace(/<li[^>]*>(.*?)<\/li>/gi, '• $1\n')

  // Convert line breaks
  text = text.replace(/<br\s*\/?>/gi, '\n')

  // Remove all remaining HTML tags
  text = text.replace(/<[^>]*>/g, '')

  // Decode HTML entities
  text = text
    .replace(/&nbsp;/g, ' ')
    .replace(/&lt;/g, '<')
    .replace(/&gt;/g, '>')
    .replace(/&amp;/g, '&')
    .replace(/&quot;/g, '"')
    .replace(/&#39;/g, "'")

  // Normalize excessive whitespace (but preserve intentional line breaks)
  text = text.replace(/ +/g, ' ') // Multiple spaces to single space
  text = text.replace(/\n\n\n+/g, '\n\n') // Max 2 consecutive newlines
  text = text.trim()

  return text
}

/**
 * Validates if HTML string has actual text content (not just tags)
 * @param html - HTML string to validate
 * @returns true if has meaningful text content
 */
export function hasTextContent(html: string): boolean {
  const plainText = stripHtmlTags(html)
  return plainText.length > 0
}
