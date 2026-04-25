/**
 * Layout standalone do lobby de telemedicina.
 *
 * Sem sidebar, sem auth, sem branding pesado — paciente entra pelo link
 * direto do email/WhatsApp e o foco é a sala.
 */
export default function LobbyLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen bg-background">
      <header className="border-b">
        <div className="container mx-auto flex items-center px-4 py-4">
          <span className="text-lg font-medium tracking-wide">Plenya</span>
        </div>
      </header>
      <main className="container mx-auto px-4 py-6">{children}</main>
    </div>
  );
}
