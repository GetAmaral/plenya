/**
 * Layout do portal do paciente — minimalista, sem sidebar profissional.
 * Atende usuários que entraram via magic link do Escore Plenya Light.
 */
export default function PatientPortalLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="min-h-screen bg-background">
      <header className="border-b">
        <div className="container mx-auto px-4 py-5 flex items-center justify-between">
          <a href="/" className="font-medium tracking-wide text-lg">
            Plenya
          </a>
          <span className="text-sm text-muted-foreground">Área do paciente</span>
        </div>
      </header>
      <main className="container mx-auto px-4 py-10">{children}</main>
    </div>
  );
}
