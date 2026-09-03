"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { Loader2 } from "lucide-react";

import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient";

/**
 * Porta de entrada do plano de devolutiva pelo menu.
 *
 * A tela de verdade mora em `/patients/[id]/plano`, que precisa do id na URL. O menu lateral só
 * tem href estático, então esta rota existe para fazer a ponte: resolve o paciente selecionado e
 * redireciona.
 *
 * Sem ela, a feature inteira (dossiê, cartões, conversa, histórico) ficava acessível só por um
 * botão "Plano" dentro da página do paciente — quem procurasse no menu não encontrava.
 *
 * `useRequireSelectedPatient` cobre o caso de não haver paciente: manda para `/patients/select`
 * guardando o destino, e o seletor devolve para cá depois da escolha. É o mesmo caminho de
 * Anamneses, Escores e Exames.
 */
export default function PlanoRedirectPage() {
  const router = useRouter();
  const { selectedPatient, isLoading } = useRequireSelectedPatient();

  useEffect(() => {
    if (isLoading || !selectedPatient?.id) return;
    // `replace` e não `push`: esta rota é um trampolim, e não deve virar um passo no histórico
    // do navegador — voltar teria que pular por cima dela.
    router.replace(`/patients/${selectedPatient.id}/plano`);
  }, [isLoading, selectedPatient?.id, router]);

  return (
    <div className="flex h-[60vh] items-center justify-center">
      <Loader2 className="h-6 w-6 animate-spin text-muted-foreground" />
    </div>
  );
}
