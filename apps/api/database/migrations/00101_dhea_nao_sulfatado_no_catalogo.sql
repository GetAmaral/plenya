-- +goose Up
-- DHEA NÃO SULFATADO: ENTRA NO CATÁLOGO, NÃO ENTRA NO ESCORE.
--
-- Um laudo trouxe "DHEA - DEHIDROEPIANDROSTERONA 1,4 ng/mL" onde tinha sido pedido DHEA-S, e a
-- carga ficou sem onde pousar: o catálogo só conhecia o sulfato. As duas decisões estão em
-- docs/emr/dhea-livre-vale-no-escore.md e são estas.
--
-- 1) Sem item de escore. O DHEA tem meia-vida de cerca de 25 minutos, é secretado em pulsos e segue
--    ritmo circadiano acoplado ao ACTH. Duas coletas do mesmo paciente em horários diferentes dão
--    números diferentes sem que nada tenha mudado nele, e um item de escore transformaria isso em
--    nível oscilante. O DHEA-S (meia-vida de 10 a 16 horas, sem ritmo circadiano, cerca de 99% do
--    androgênio adrenal circulante) é o que já tem os doze itens calibrados por sexo e década.
--
-- 2) Com `alt_names` QUALIFICADOS, e nenhum genérico. A tentação era reivindicar "dhea" e
--    "dehidroepiandrosterona" soltos, e ela produz exatamente o dano que esta migration quer
--    evitar, só que ao contrário. `MatchIndex.Resolve` cai para "o nome indexado mais longo que
--    contém ou é contido" quando não há match exato, e o laudo do sulfato costuma vir como
--    "DEHIDROEPIANDROSTERONA (DHEA) SULFATO", que normaliza para
--    "dehidroepiandrosterona dhea sulfato". O alternativo do DHEA-S "dehidroepiandrosterona
--    sulfato" NÃO é substring dessa linha, porque o token "dhea" se intromete no meio; mas
--    "dehidroepiandrosterona" solto seria, com 22 caracteres, e ganharia. Um valor de DHEA-S em
--    µg/dL passaria a ser gravado contra uma definição em ng/mL, sem que a rede de plausibilidade
--    tivesse como perceber. Por isso só entram formas que carregam o qualificador.
INSERT INTO public.lab_test_definitions
  (id, code, name, short_name, alt_names, category, is_requestable, unit, result_type,
   specimen_type, sex_applicability, is_active, display_order,
   clinical_significance, created_at, updated_at)
SELECT uuid_generate_v7(), 'PLNDHEALIVRE', 'DHEA (dehidroepiandrosterona)', 'DHEA',
       '["dhea - dehidroepiandrosterona", "dhea dehidroepiandrosterona", "dehidroepiandrosterona nao sulfatada", "dehidroepiandrosterona nao conjugada", "dhea nao sulfatado", "dhea livre", "dhea total"]'::jsonb,
       'hormones', true, 'ng/mL', 'numeric', 'Soro', 'all', true,
       -- COALESCE porque o catálogo de exames NÃO é semeado por migration: numa base nova (CI, um
       -- ambiente recriado) o DHEA-S não existe, a subquery escalar devolve NULL, e `display_order`
       -- é NOT NULL. Sem isto a migration aborta e leva o deploy do api junto.
       coalesce((SELECT display_order + 1 FROM public.lab_test_definitions
                  WHERE code = 'PLN82DBE091' AND deleted_at IS NULL), 0),
       'Forma não sulfatada do androgênio adrenal. Arquivada no prontuário sem entrar no escore: '
       'meia-vida de cerca de 25 minutos, secreção em pulsos e ritmo circadiano acoplado ao ACTH '
       'fazem o valor variar com o horário da coleta. Para pontuar androgênio adrenal, o exame é o '
       'DHEA-S (PLN82DBE091), que é estável ao longo do dia.',
       now(), now()
WHERE NOT EXISTS (SELECT 1 FROM public.lab_test_definitions WHERE code = 'PLNDHEALIVRE');

-- +goose Down
DELETE FROM public.lab_test_definitions WHERE code = 'PLNDHEALIVRE';
