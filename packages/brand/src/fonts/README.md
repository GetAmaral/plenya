# Plenya Fonts

Fontes proprietárias da marca devem ser colocadas aqui em formatos `woff2` (preferencial) e `woff`.

## Estrutura esperada

```
fonts/
├── nalieta/
│   ├── Nalieta-Regular.woff2
│   ├── Nalieta-Medium.woff2
│   ├── Nalieta-SemiBold.woff2
│   └── Nalieta-Bold.woff2
└── polarband/
    ├── PolarbandWOO-Front-Regular.woff2
    └── PolarbandWOO-Front-Bold.woff2
```

## CSS @font-face

Cada app que consumir essas fontes deve declarar `@font-face` apontando para os arquivos via `import` ou `next/font/local`. Exemplo no `apps/site`:

```ts
import localFont from 'next/font/local';

export const nalieta = localFont({
  src: [
    { path: '../node_modules/@plenya/brand/src/fonts/nalieta/Nalieta-Regular.woff2', weight: '400', style: 'normal' },
    { path: '../node_modules/@plenya/brand/src/fonts/nalieta/Nalieta-Bold.woff2', weight: '700', style: 'normal' },
  ],
  variable: '--font-heading',
  display: 'swap',
});
```

## Fallback temporário

Enquanto os arquivos não estão disponíveis, o `tailwind-preset` declara fallbacks (Cormorant Garamond para Nalieta, IBM Plex Mono para Polarband). Em produção isso deve ser substituído pelas fontes oficiais.
