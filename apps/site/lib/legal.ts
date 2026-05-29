/**
 * Versão e metadata dos documentos legais.
 * Quando atualizar a Política ou Termos, incrementar a versão correspondente
 * — sessões/usuários registram qual versão aceitaram.
 */

export const PRIVACY_POLICY_VERSION = '2026-04-25.1';
export const TERMS_VERSION = '2026-04-22.1';

export const LEGAL_CONTACT = {
  dpoEmail: 'dpo@plenyasaude.com.br',
  dpoName: 'Dr. Getúlio Amaral',
  dpoTitle: 'Encarregado de Proteção de Dados (DPO)',
  controllerName: 'Plenya Serviços de Saúde Ltda.',
  controllerCnpj: '66.991.259/0001-50',
  controllerEmail: 'contato@plenyasaude.com.br',
  // Endereço da SEDE (CNPJ/fiscal) — uso legal apenas. O endereço de
  // atendimento (NAP público) vive em `@plenya/brand` → brand.address.
  controllerAddress: 'Av. Gil de Abreu e Souza, 2335, Casa 634 — Bairro Esperança, Londrina/PR, 86058-100',
} as const;
