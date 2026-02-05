/**
 * Formata CPF para exibição: 12345678900 -> 123.456.789-00
 */
export function formatCPF(cpf: string | null | undefined): string {
  if (!cpf) return ''

  // Remove tudo que não é número
  const numbers = cpf.replace(/\D/g, '')

  // Se não tiver 11 dígitos, retorna como está
  if (numbers.length !== 11) return cpf

  // Formata: 123.456.789-00
  return numbers.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, '$1.$2.$3-$4')
}

/**
 * Remove formatação do CPF: 123.456.789-00 -> 12345678900
 */
export function sanitizeCPF(cpf: string | null | undefined): string {
  if (!cpf) return ''
  return cpf.replace(/\D/g, '')
}

/**
 * Valida se CPF tem formato válido (11 dígitos)
 */
export function isValidCPFFormat(cpf: string | null | undefined): boolean {
  if (!cpf) return false
  const numbers = cpf.replace(/\D/g, '')
  return numbers.length === 11
}
