import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

/**
 * Normaliza string para busca: remove acentos, converte para minúsculas e remove espaços extras
 * Exemplo: "Sódio" -> "sodio", "GLICOSE" -> "glicose"
 */
export function normalizeForSearch(str: string): string {
  if (!str) return ''

  return str
    .normalize('NFD') // Decompõe caracteres acentuados (á -> a + ´)
    .replace(/[\u0300-\u036f]/g, '') // Remove marcas diacríticas (acentos)
    .toLowerCase()
    .trim()
}
