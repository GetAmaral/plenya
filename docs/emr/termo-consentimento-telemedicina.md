# Termo de Consentimento — Telemedicina (CFM 2.314/2022)

> **Status:** VALIDADO e IMPLANTADO (2026-06-02). Decisões do Getúlio: **gravação incluída**
> (cláusula 7) + captura **verbal** (médico registra no início da sala). Texto canônico vive em
> `services/appointment_service.go` const `TelemedConsentTermText` (alterar lá ao revisar).
> Base: Resolução CFM nº 2.314/2022 (regula a telemedicina no Brasil). O consentimento de
> telemedicina é exigência **médica/CFM** (consentir com o *formato remoto*), distinto da base
> LGPD do ato clínico ("tutela da saúde") e do consentimento de CID em atestado.

## O que o CFM 2.314/2022 exige (e este termo cobre)
1. Consentimento **livre e esclarecido** do paciente (ou representante legal) para o atendimento remoto.
2. **Registro no prontuário** (data/hora + texto apresentado + quem registrou).
3. Informar o **direito de optar pelo atendimento presencial**, sem prejuízo.
4. Informar **limitações** da telemedicina (ex.: exame físico restrito) e quando o presencial é necessário.
5. **Sigilo médico** e **proteção de dados** (LGPD) mantidos.
6. **Gravação** (se houver) só com consentimento explícito; integra o prontuário.

---

## Termo (versão completa — vira `telemed_consent_text`)

**TERMO DE CONSENTIMENTO LIVRE E ESCLARECIDO PARA ATENDIMENTO POR TELEMEDICINA**

Atendimento prestado por Plenya Serviços de Saúde Ltda. (CNPJ 66.991.259/0001-50), nos termos da
Resolução CFM nº 2.314/2022.

Declaro estar ciente e de acordo com o seguinte:

1. Este atendimento será realizado por **telemedicina**, de forma remota, por meio de vídeo em
   plataforma segura, sem a presença física do médico no mesmo ambiente.
2. Fui informado(a) de que tenho o **direito de optar pelo atendimento presencial** a qualquer
   momento, sem qualquer prejuízo ao meu cuidado.
3. Compreendo que a telemedicina possui **limitações** — em especial quanto ao exame físico — e que
   o médico poderá recomendar avaliação presencial, exames complementares ou encaminhamento sempre
   que julgar necessário para a minha segurança.
4. Mantêm-se integralmente o **sigilo médico** e a **proteção dos meus dados pessoais e de saúde**,
   tratados conforme a legislação aplicável (LGPD).
5. Posso **interromper** o atendimento a qualquer momento e esclarecer dúvidas com o médico.
6. Em caso de **emergência ou urgência**, devo procurar imediatamente um serviço de pronto
   atendimento presencial.

Declaro que fui devidamente esclarecido(a), tive a oportunidade de tirar dúvidas e **consinto
livremente** com a realização do atendimento por telemedicina.

_(Registrado no prontuário do paciente em <data/hora>, no início da teleconsulta.)_

### Cláusula opcional de gravação (incluir só se a clínica for gravar)
Autorizo a **gravação** desta teleconsulta, que integrará o meu prontuário e ficará sujeita ao
mesmo sigilo e proteção de dados.

---

## Versão curta (rótulo do checkbox / consentimento verbal registrado pelo médico)
> "O paciente foi esclarecido sobre o atendimento por telemedicina, seus direitos e limitações
> (Resolução CFM nº 2.314/2022), e **consente livremente** com a teleconsulta."

---

## Pontos para o Getúlio validar
- [ ] Redação geral aprovada?
- [ ] Incluir a **cláusula de gravação** agora ou deixar de fora (sem gravação por ora)?
- [ ] Modalidade de captura: **(A) verbal** — o médico marca "paciente consentiu" no início da sala
      (mais simples, o CFM aceita); ou **(B) escrita** — o paciente confirma/assina no portal/sala
      antes de entrar (mais robusto, exige ação do paciente). Recomendo começar por (A).
- [ ] Identificação jurídica correta (Plenya Serviços de Saúde Ltda. / CNPJ 66.991.259/0001-50)?

---

## Plano de implantação (após validação) — peça pequena, sem bloqueio
**Backend** (migration `00011`, aditiva em `appointments`):
- `telemed_consent_at *time.Time`, `telemed_consent_text text`, `telemed_consent_by_user_id *uuid`,
  `telemed_consent_mode varchar` (`verbal|written`).
- Endpoint `POST /api/v1/appointments/:id/telemed-consent` (RequireAnyStaff/Clinician + AuditLog):
  carimba `consent_at=now`, grava o texto vigente e quem registrou. Idempotente.
- DTO + expor os campos no `AppointmentResponse`.

**Frontend** (workspace, só quando `type=telemedicine`):
- Bloco do termo **lado a lado com o vídeo** (o plano original pedia "doc lado-a-lado"); enquanto
  `telemed_consent_at` for nulo, exibir o termo + botão **"Registrar consentimento"** (modo A) ou
  o fluxo de confirmação do paciente (modo B).
- Após registrado: selo "Consentimento de telemedicina registrado em <data>" e o termo entra na
  timeline/prontuário da consulta.

**Auditoria/registro:** o carimbo + texto satisfazem o "registro no prontuário" do CFM.
**Opcional futuro:** plugar `transcription_service` à gravação da Daily.co (casa com a cláusula de gravação).
