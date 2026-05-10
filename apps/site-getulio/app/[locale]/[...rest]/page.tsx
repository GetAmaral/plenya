import { notFound } from 'next/navigation';

// Catch-all que força qualquer rota não casada a cair em
// app/[locale]/not-found.tsx — sem isso, o Next renderiza o /_not-found
// global default (sem layout, em inglês).
export default function CatchAllPage() {
  notFound();
}
