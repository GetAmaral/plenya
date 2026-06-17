### Production Dockerfile for @plenya/site-getulio
### Multi-stage. Builds with workspace context.

FROM node:24-alpine AS base
RUN corepack enable && corepack prepare pnpm@10.28.1 --activate
WORKDIR /repo

###
### deps — install workspace deps (frozen)
###
FROM base AS deps
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml turbo.json ./
COPY apps/site-getulio/package.json ./apps/site-getulio/
RUN pnpm install --frozen-lockfile --filter "@plenya/site-getulio..."

###
### builder — copy source + build
###
FROM base AS builder
COPY --from=deps /repo/node_modules ./node_modules
COPY --from=deps /repo/apps/site-getulio/node_modules ./apps/site-getulio/node_modules
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1
RUN pnpm --filter @plenya/site-getulio build

###
### runner — minimal production image (Next standalone, sem pnpm/node_modules de runtime)
###
FROM node:24-alpine AS runner
WORKDIR /repo

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3003

# Standalone traz server.js + node_modules tree-shaken (raiz = monorepo via outputFileTracingRoot).
COPY --from=builder /repo/apps/site-getulio/.next/standalone ./
# static e public NÃO entram no standalone — copiar pro lado do server.js do app.
COPY --from=builder /repo/apps/site-getulio/.next/static ./apps/site-getulio/.next/static
COPY --from=builder /repo/apps/site-getulio/public ./apps/site-getulio/public
# Conteúdo MDX: o app é SSG (baked no build), mas copiamos por paridade/segurança — blog.ts lê de
# process.cwd()/content/blog e o cwd em runtime é /repo/apps/site-getulio (WORKDIR abaixo).
COPY --from=builder /repo/apps/site-getulio/content ./apps/site-getulio/content
COPY --from=builder /repo/apps/site/content ./apps/site/content

EXPOSE 3003

WORKDIR /repo/apps/site-getulio
CMD ["node", "server.js"]
