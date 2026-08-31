-- Densidade real da vitamina E no catálogo magistral.
--
-- PROBLEMA: a vitamina E estava com bulk_density 0,45 e density_source 'classe', que é o default
-- de PÓ. Em cápsula oleosa ela entra como acetato de tocoferol LÍQUIDO, densidade ~0,95 g/mL.
-- O erro dobrava o volume estimado do componente, e o CalculateCapsule passava a recomendar duas
-- cápsulas onde uma bastava. Achado na fórmula lipossolúvel da paciente Ana Cláudia: com 0,45 o
-- motor devolvia 1,546 mL com margem, acima da 000; com a densidade certa devolve 1,151 mL.
--
-- LIMIAR: para essa fórmula caber numa 000, bastava densidade >= 0,59. A real, 0,95, resolve com
-- folga e também acerta qualquer outra fórmula oleosa.
--
-- RESSALVA REGISTRADA NA NOTA: a vitamina E também é dispensada como pó adsorvido em sílica, e
-- nessa forma a densidade volta para a casa de 0,45. Como o campo é único por componente, fica
-- valendo a forma líquida, que é a usada nas cápsulas oleosas, e a nota avisa do outro caso.
-- Se um dia valer a pena separar, o caminho é criar dois componentes distintos.

UPDATE public.magistral_components SET
  bulk_density   = 0.95,
  density_source = 'medida',
  notes = 'Fórmula BR: 400 mg. Teto IN 28/2018 para suplemento: 1.000 mg/dia (consultivo, não se aplica a magistral). Densidade de 0,95 g/mL é a do acetato de tocoferol líquido, que é a forma usada em cápsula oleosa. Na forma de pó adsorvido em sílica a densidade cai para perto de 0,45 e o volume calculado fica subestimado.',
  dose_reference = 'Dose em mg de d-alfa-tocoferol. Conversão: 1 UI equivale a cerca de 0,67 mg de d-alfa-tocoferol, ou 0,735 mg do acetato.',
  last_review = now(),
  updated_at = now()
WHERE name = 'Vitamina E' AND deleted_at IS NULL;

SELECT name, bulk_density, density_source, dose_reference
FROM public.magistral_components WHERE name = 'Vitamina E' AND deleted_at IS NULL;
