-- Veículo oleoso do catálogo magistral: triglicerídeos de cadeia média (TCM).
--
-- POR QUE FALTAVA: o catálogo tinha 290 componentes e nenhum veículo. Toda cápsula com ativo
-- lipossolúvel (vitaminas A, D, E, K, coenzima Q10, carotenoides) precisa de um, e sem ele a
-- fórmula não podia sequer ser escrita no módulo. Achado ao montar a fórmula da paciente
-- Ana Cláudia, que junta D3, E, K2 MK-7, A, CavaQ10 e palmitato de ascorbila numa cápsula só.
--
-- POR QUE A UNIDADE É ml E NÃO mg: o veículo entra como qsp, preenchendo o volume que sobra na
-- cápsula depois dos ativos. Ele não tem dose própria, então usual_dose/min_dose/max_dose ficam
-- nulos de propósito. Consequência desejada no motor: MassToMg (magistral_capsule.go) devolve
-- ok=false para ml, e magistral_component_service.go pula o componente antes de montar o
-- CapsuleInput. O TCM fica de fora do cálculo de volume, que é exatamente o certo: quem dimensiona
-- a cápsula são os ativos, e o veículo ocupa o resto.
--
-- DENSIDADE: 0,945 g/mL a 20 °C é constante física de ficha técnica do triglicerídeo
-- caprílico/cáprico, não default de classe. Por isso density_source = 'medida'. O campo fica
-- preenchido para o dia em que alguém prescrever o TCM em massa, e para não disparar o
-- "sem densidade cadastrada" do CalculateCapsule caso a unidade seja trocada.
--
-- ESTABILIDADE: TCM é saturado, o que o torna resistente à oxidação e à rancificação. É essa a
-- razão de ele ser o veículo preferido sobre óleos insaturados quando a cápsula carrega ativos
-- oxidação-sensíveis. Daí oxidation_sensitive e photosensitive ficarem em false.

INSERT INTO public.magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, bulk_density, density_source,
   hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok,
   indications, indication_bullets, dose_reference, notes, source, evidence_status, last_review)
SELECT * FROM (VALUES
  (uuidv7(), 'Triglicerídeos de cadeia média (TCM)',
   'TCM, MCT, medium chain triglycerides, óleo de coco fracionado, triglicerídeo caprílico/cáprico, ácidos caprílico e cáprico',
   'ml',
   NULL::numeric, NULL::numeric, NULL::numeric, 0.945::numeric, 'medida',
   false, false, false, false, 0::smallint, false,
   'Veículo oleoso para cápsulas com ativos lipossolúveis. Dissolve e dispersa vitaminas A, D, E e K, coenzima Q10 e carotenoides, e fornece a gordura de que a absorção dessas moléculas depende. Sendo saturado, resiste à oxidação melhor que os óleos vegetais insaturados usados como alternativa.',
   E'veículo de cápsula com ativos lipossolúveis\nabsorção de vitaminas A, D, E e K\ndispersão de coenzima Q10 e carotenoides',
   'Entra como qsp: preenche o volume restante da cápsula depois dos ativos. Não tem dose própria, e por isso não traz dose usual, mínima nem máxima.',
   'Veículo, não princípio ativo. Fica fora do cálculo de volume da cápsula de propósito, porque quem dimensiona a cápsula são os ativos e o TCM ocupa o que sobra. Densidade de 0,945 g/mL a 20 °C vem da ficha técnica do caprílico/cáprico.',
   'pesquisa', 'suggested', now())
) AS v(id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, bulk_density, density_source,
       hygroscopic, oxidizing, oxidation_sensitive, photosensitive, bitterness, sachet_ok,
       indications, indication_bullets, dose_reference, notes, source, evidence_status, last_review)
WHERE NOT EXISTS (
  SELECT 1 FROM public.magistral_components c
  WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.name))
);

SELECT name, default_unit, bulk_density, density_source, source, evidence_status
FROM public.magistral_components
WHERE name = 'Triglicerídeos de cadeia média (TCM)' AND deleted_at IS NULL;
