import type { Metadata } from "next";
import { Inter, Cormorant_Garamond, IBM_Plex_Mono } from "next/font/google";
import "./globals.css";
import { QueryProvider } from "@/lib/query-provider";
import { Toaster } from "@/components/ui/toaster";
import { CommandPalette } from "@/components/command-palette";
import { NavigationProgress } from "@/components/navigation-progress";
import { Suspense } from "react";
import "@/lib/suppress-dev-logs"; // Suppress verbose development logs

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-inter",
});

const cormorant = Cormorant_Garamond({
  subsets: ["latin"],
  weight: ["400", "500", "600"],
  variable: "--font-cormorant",
});

const plexMono = IBM_Plex_Mono({
  subsets: ["latin"],
  weight: ["400", "500"],
  variable: "--font-mono",
});

export const metadata: Metadata = {
  title: "Plenya EMR - Sistema de Prontuário Médico Eletrônico",
  description: "Sistema completo de prontuário eletrônico com web app, mobile apps e backend Go",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt-BR" data-scroll-behavior="smooth">
      <body className={`${inter.variable} ${cormorant.variable} ${plexMono.variable} font-sans antialiased`}>
        <QueryProvider>
          <div className="print:hidden">
            <Suspense fallback={null}>
              <NavigationProgress />
            </Suspense>
          </div>
          {children}
          <div className="print:hidden">
            <CommandPalette />
            <Toaster />
          </div>
        </QueryProvider>
      </body>
    </html>
  );
}
