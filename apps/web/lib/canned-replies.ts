// Respostas rápidas (canned text) da recepção para a inbox omnichannel.
// {nome} é substituído pelo primeiro nome do contato na inserção.
// Voz de marca: prosa clínica PT-BR, sem travessões, sem slogans, sem promessas.

export interface CannedReply {
  id: string;
  label: string;
  text: string;
}

export const CANNED_REPLIES: CannedReply[] = [
  {
    id: "saudacao",
    label: "Saudação",
    text: "Olá, {nome}. Aqui é da Plenya. Como posso ajudar?",
  },
  {
    id: "horario",
    label: "Horário de atendimento",
    text: "Nosso atendimento é de segunda a sexta, das 8h às 18h. Posso te ajudar a encontrar um horário?",
  },
  {
    id: "endereco",
    label: "Endereço e acesso",
    text: "{nome}, assim que sua consulta for confirmada eu envio o endereço completo e as orientações de acesso.",
  },
  {
    id: "preparo",
    label: "Orientações de preparo",
    text: "{nome}, para a sua consulta vale trazer exames recentes e a lista de medicações em uso. Qualquer dúvida, fico à disposição.",
  },
  {
    id: "agendamento",
    label: "Agendar consulta",
    text: "{nome}, me diga os melhores dias e horários para você que eu verifico a disponibilidade na agenda.",
  },
  {
    id: "confirmacao",
    label: "Confirmar consulta",
    text: "{nome}, estou confirmando a sua consulta. Se precisar remarcar, é só me avisar por aqui.",
  },
];

/** Substitui {nome} pelo primeiro nome do contato. Sem nome, remove o placeholder. */
export function applyCannedReply(text: string, contactName?: string): string {
  const first = (contactName ?? "").trim().split(/\s+/)[0] ?? "";
  return text
    .replace(/\{nome\}/g, first)
    .replace(/, \. /g, ". ") // limpa "Olá, . " quando não há nome
    .replace(/Olá, \./g, "Olá.")
    .trim();
}
