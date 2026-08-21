import type { User } from "./auth-store";

/**
 * Para onde mandar quem JÁ está logado.
 *
 * Mesma regra usada logo após o login: secretaria sem papel clínico cai na Recepção; o resto
 * vai pro dashboard. Fica aqui — e não dentro da tela de login — porque três lugares precisam
 * dela: o login, a raiz `/` e a volta do PWA.
 */
export function homeFor(user: User | null): string {
  const roles = user?.roles ?? [];
  const secretaryOnly =
    roles.includes("secretary") &&
    !roles.some((r) =>
      ["doctor", "nurse", "nutritionist", "psychologist", "physicalEducator"].includes(r),
    );
  return secretaryOnly ? "/recepcao" : "/dashboard";
}
