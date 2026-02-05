/**
 * Converte números para extenso em português
 * Usado para quantidade de medicamentos em prescrições
 * Suporta números de 0 a 9999
 */

const unidades = [
  '', 'um', 'dois', 'três', 'quatro', 'cinco', 'seis', 'sete', 'oito', 'nove'
]

const dez_a_dezenove = [
  'dez', 'onze', 'doze', 'treze', 'quatorze', 'quinze',
  'dezesseis', 'dezessete', 'dezoito', 'dezenove'
]

const dezenas = [
  '', '', 'vinte', 'trinta', 'quarenta', 'cinquenta',
  'sessenta', 'setenta', 'oitenta', 'noventa'
]

const centenas = [
  '', 'cento', 'duzentos', 'trezentos', 'quatrocentos', 'quinhentos',
  'seiscentos', 'setecentos', 'oitocentos', 'novecentos'
]

/**
 * Converte número para extenso (0-99)
 */
function converteDezena(num: number): string {
  if (num < 10) {
    return unidades[num]
  }
  if (num >= 10 && num < 20) {
    return dez_a_dezenove[num - 10]
  }
  const dezena = Math.floor(num / 10)
  const unidade = num % 10
  if (unidade === 0) {
    return dezenas[dezena]
  }
  return `${dezenas[dezena]} e ${unidades[unidade]}`
}

/**
 * Converte número para extenso (0-999)
 */
function converteCentena(num: number): string {
  if (num === 100) {
    return 'cem'
  }
  if (num < 100) {
    return converteDezena(num)
  }
  const centena = Math.floor(num / 100)
  const resto = num % 100
  if (resto === 0) {
    return centenas[centena]
  }
  return `${centenas[centena]} e ${converteDezena(resto)}`
}

/**
 * Converte número para extenso (0-9999)
 *
 * @param num - Número inteiro entre 0 e 9999
 * @returns String com número por extenso
 *
 * @example
 * numberToWords(0)    // 'zero'
 * numberToWords(1)    // 'um'
 * numberToWords(15)   // 'quinze'
 * numberToWords(30)   // 'trinta'
 * numberToWords(100)  // 'cem'
 * numberToWords(250)  // 'duzentos e cinquenta'
 * numberToWords(1000) // 'mil'
 * numberToWords(1500) // 'mil e quinhentos'
 */
export function numberToWords(num: number): string {
  // Validação
  if (!Number.isInteger(num) || num < 0 || num > 9999) {
    throw new Error('Número deve ser inteiro entre 0 e 9999')
  }

  // Casos especiais
  if (num === 0) return 'zero'
  if (num === 1000) return 'mil'

  // Milhar
  if (num >= 1000) {
    const milhar = Math.floor(num / 1000)
    const resto = num % 1000

    let resultado = ''
    if (milhar === 1) {
      resultado = 'mil'
    } else {
      resultado = `${unidades[milhar]} mil`
    }

    if (resto > 0) {
      if (resto < 100) {
        resultado += ` e ${converteCentena(resto)}`
      } else {
        resultado += ` e ${converteCentena(resto)}`
      }
    }
    return resultado
  }

  // Centenas
  return converteCentena(num)
}

/**
 * Converte número para extenso com unidade de medicamento
 *
 * @param quantity - Quantidade numérica
 * @param unit - Unidade (comprimidos, cápsulas, ampolas, etc.)
 * @returns String formatada para prescrição
 *
 * @example
 * numberToWordsWithUnit(30, 'comprimidos') // 'trinta comprimidos'
 * numberToWordsWithUnit(1, 'ampola')       // 'uma ampola'
 * numberToWordsWithUnit(2, 'frascos')      // 'dois frascos'
 */
export function numberToWordsWithUnit(quantity: number, unit: string): string {
  const words = numberToWords(quantity)

  // Ajustar "um" para "uma" se a unidade for feminina
  const feminineUnits = ['ampola', 'cápsula', 'dose', 'unidade', 'bisnaga']
  const isFeminine = feminineUnits.some(u => unit.toLowerCase().includes(u))

  if (quantity === 1 && isFeminine) {
    return `uma ${unit}`
  }

  return `${words} ${unit}`
}

/**
 * Valida se a quantidade é adequada para prescrição
 * CFM/ANVISA: máximo de 3 medicamentos controlados por prescrição
 *
 * @param quantity - Quantidade a validar
 * @param category - Categoria do medicamento (c1, c5, etc.)
 * @returns true se válido
 */
export function isValidPrescriptionQuantity(
  quantity: number,
  category: string
): boolean {
  if (quantity <= 0 || quantity > 999) {
    return false
  }

  // Controlados (C1, C5): máximo 180 unidades (6 meses de tratamento)
  if (category === 'c1' || category === 'c5') {
    return quantity <= 180
  }

  // Antibióticos: máximo 30 unidades (10 dias)
  if (category === 'antibiotic') {
    return quantity <= 30
  }

  // Simples: máximo 999
  return true
}
