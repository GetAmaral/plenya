export type Author = {
  slug: string;
  name: string;
  role: string;
  credentials: string;
  bio: string;
  photo?: string;
  social?: { instagram?: string; linkedin?: string; website?: string };
};

export const authors: Record<string, Author> = {
  'getulio-amaral': {
    slug: 'getulio-amaral',
    name: 'Dr. Getúlio Amaral Filho',
    role: 'Direção Clínica · Nefrologia · Medicina Funcional Integrativa',
    credentials: 'CRM-PR 21.876 · RQE 16.038',
    bio: 'Médico desde 2004. Especialista em nefrologia e clínica médica pela Santa Casa de Londrina, onde coordena a residência médica em nefrologia. Pós-graduado em medicina funcional integrativa pela ABMFI. Direção clínica da Plenya.',
    photo: '/images/dr-getulio.jpg',
    social: {
      instagram: 'https://instagram.com/drGetulioAmaralFilho',
      website: 'https://drGetulioAmaralFilho.com.br',
    },
  },
};

export function getAuthor(slug: string): Author | null {
  return authors[slug] ?? null;
}
