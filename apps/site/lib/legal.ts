/**
 * Versão e metadata dos documentos legais.
 * Quando atualizar a Política ou Termos, incrementar a versão correspondente
 * — sessões/usuários registram qual versão aceitaram.
 */

export const PRIVACY_POLICY_VERSION = '2026-04-22.1';
export const TERMS_VERSION = '2026-04-22.1';

export const LEGAL_CONTACT = {
  dpoEmail: 'dpo@plenyasaude.com.br',
  dpoName: 'Encarregado de Proteção de Dados — Plenya',
  controllerName: 'Plenya Saúde Ltda.',
  controllerEmail: 'contato@plenyasaude.com.br',
  // Endereço deve ser preenchido com o endereço da clínica
  controllerAddress: 'Londrina/PR, Brasil',
} as const;
