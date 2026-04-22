'use client';

/**
 * Formulário público de exercício dos direitos do titular (LGPD art. 18).
 *
 * Strategy: usa mailto: como entrega — abre o cliente de email do usuário
 * com payload pré-preenchido para o DPO. Sem backend dedicado.
 *
 * Vantagens: zero infra, zero risco de coleta indevida de dados, registro
 * fica no email do próprio titular (transparência).
 */
import { useState } from 'react';

const RIGHTS = [
  { id: 'access', label: 'Confirmação e acesso aos meus dados', desc: 'Saber que dados a Plenya tem sobre mim e obter cópia.' },
  { id: 'correct', label: 'Correção de dados', desc: 'Corrigir informações incompletas, inexatas ou desatualizadas.' },
  { id: 'delete', label: 'Exclusão dos meus dados', desc: 'Eliminar dados pessoais tratados com meu consentimento.' },
  { id: 'anonymize', label: 'Anonimização ou bloqueio', desc: 'Anonimizar ou bloquear dados desnecessários ou tratados em desconformidade.' },
  { id: 'portability', label: 'Portabilidade dos meus dados', desc: 'Receber meus dados em formato estruturado para transferir a outro serviço.' },
  { id: 'sharing', label: 'Informação sobre compartilhamento', desc: 'Saber com quais entidades a Plenya compartilhou meus dados.' },
  { id: 'revoke', label: 'Revogação do consentimento', desc: 'Retirar consentimento previamente concedido.' },
  { id: 'other', label: 'Outro', desc: 'Descreva no campo de detalhes abaixo.' },
];

export function LGPDRightsForm({ dpoEmail }: { dpoEmail: string }) {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [right, setRight] = useState<string>('');
  const [details, setDetails] = useState('');
  const [sessionCode, setSessionCode] = useState('');

  const isValid = name.trim() !== '' && email.trim() !== '' && right !== '';

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!isValid) return;

    const rightLabel = RIGHTS.find((r) => r.id === right)?.label ?? right;

    const subject = `[LGPD] ${rightLabel} — ${name}`;
    const body = `Solicitação de exercício de direito do titular (LGPD art. 18)

Direito solicitado: ${rightLabel}

Solicitante:
- Nome: ${name}
- Email/identificador: ${email}
${sessionCode ? `- Código da sessão Light (se aplicável): ${sessionCode}\n` : ''}
Detalhes adicionais:
${details || '(nenhum)'}

---
Solicitação enviada via formulário em /lgpd/direitos.
Resposta esperada em até 15 dias úteis (LGPD art. 19).
`;

    window.location.href = `mailto:${dpoEmail}?subject=${encodeURIComponent(subject)}&body=${encodeURIComponent(body)}`;
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-6 border border-petrol/15 bg-paper p-6 rounded-md">
      <div>
        <label className="block text-petrol/80 mb-2 text-sm font-medium">
          Qual direito você quer exercer? <span className="text-red-600">*</span>
        </label>
        <div className="space-y-2">
          {RIGHTS.map((r) => (
            <label key={r.id} className="flex gap-3 items-start cursor-pointer hover:bg-cream/60 p-2 rounded transition">
              <input
                type="radio"
                name="right"
                value={r.id}
                checked={right === r.id}
                onChange={() => setRight(r.id)}
                className="mt-1 accent-gold"
              />
              <span>
                <span className="text-petrol font-medium block">{r.label}</span>
                <span className="text-petrol/60 text-sm">{r.desc}</span>
              </span>
            </label>
          ))}
        </div>
      </div>

      <div className="grid sm:grid-cols-2 gap-4">
        <div>
          <label className="block text-petrol/80 mb-2 text-sm font-medium">
            Seu nome <span className="text-red-600">*</span>
          </label>
          <input
            type="text"
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
          />
        </div>
        <div>
          <label className="block text-petrol/80 mb-2 text-sm font-medium">
            Seu email <span className="text-red-600">*</span>
          </label>
          <input
            type="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
            placeholder="usado para identificá-lo e responder"
          />
        </div>
      </div>

      <div>
        <label className="block text-petrol/80 mb-2 text-sm font-medium">
          Código da sua sessão Light (opcional)
        </label>
        <input
          type="text"
          value={sessionCode}
          onChange={(e) => setSessionCode(e.target.value)}
          placeholder="ex: jG5RAMX2089a — encontre na URL do seu radar"
          className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none font-mono text-sm"
        />
        <p className="text-petrol/55 text-xs mt-1">
          Se você fez o Escore Plenya Light e quer agir sobre essa sessão, informe o
          código que aparece na URL do seu radar (ex.: <code>/resultado/<strong>jG5RAMX2089a</strong></code>).
        </p>
      </div>

      <div>
        <label className="block text-petrol/80 mb-2 text-sm font-medium">
          Detalhes adicionais
        </label>
        <textarea
          value={details}
          onChange={(e) => setDetails(e.target.value)}
          rows={4}
          className="w-full border border-petrol/20 bg-cream px-4 py-3 text-petrol focus:border-gold focus:outline-none"
          placeholder="opcional — descreva a solicitação com mais detalhes"
        />
      </div>

      <div className="bg-cream border-l-2 border-gold pl-3 py-2 text-petrol/70 text-xs">
        Ao enviar, abrirá seu cliente de email com a mensagem pré-preenchida. Você
        confere e envia da sua caixa — assim mantemos um registro claro da solicitação
        em seu próprio email.
      </div>

      <div>
        <button
          type="submit"
          disabled={!isValid}
          className={`btn-gold ${!isValid ? 'opacity-40 cursor-not-allowed' : ''}`}
        >
          Preparar email para o Encarregado
        </button>
      </div>
    </form>
  );
}
