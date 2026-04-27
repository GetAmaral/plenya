type Props = {
  /**
   * Optional override for the date shown. Defaults to current month/year.
   */
  reviewedAt?: string;
};

function formatPtMonthYear(d: Date) {
  return d.toLocaleDateString('pt-BR', { month: 'long', year: 'numeric' });
}

export function ClinicalReviewBadge({ reviewedAt }: Props) {
  const date = reviewedAt ? new Date(reviewedAt) : new Date();
  return (
    <aside className="bg-paper border-l-2 border-gold py-4 px-5 my-12 max-w-2xl">
      <p className="label-upper text-gold mb-1">Revisão clínica</p>
      <p className="text-petrol/75 text-sm leading-relaxed">
        Conteúdo médico revisado por <strong>Dr. Getúlio José Mattos do Amaral Filho</strong> ·
        Direção Clínica da Plenya · CRM-PR 21.876 · RQE 16.038 (Nefrologia) ·{' '}
        {formatPtMonthYear(date)}.
      </p>
    </aside>
  );
}
