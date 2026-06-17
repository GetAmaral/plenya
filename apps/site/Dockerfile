### Production Dockerfile for @plenya/site
### Multi-stage. Builds with workspace context.

FROM node:24-alpine AS base
RUN corepack enable && corepack prepare pnpm@10.28.1 --activate
WORKDIR /repo

###
### deps — install full workspace deps (frozen)
###
FROM base AS deps
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml turbo.json ./
COPY apps/site/package.json ./apps/site/
COPY packages/brand/package.json ./packages/brand/
COPY packages/types/package.json ./packages/types/
COPY packages/ui/package.json ./packages/ui/
RUN pnpm install --frozen-lockfile --filter "@plenya/site..."

###
### builder — copy source + build
###
FROM base AS builder
COPY --from=deps /repo/node_modules ./node_modules
COPY --from=deps /repo/apps/site/node_modules ./apps/site/node_modules
COPY --from=deps /repo/packages/brand/node_modules ./packages/brand/node_modules
# @plenya/ui é transpilado pelo site (transpilePackages) → o tsc type-checa os
# fontes dele; precisa do node_modules dele (react peer) no isolated-install do
# Docker, senão "Cannot find module 'react'" em packages/ui/src/score-radar.
COPY --from=deps /repo/packages/ui/node_modules ./packages/ui/node_modules
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1

# NEXT_PUBLIC_* precisam estar disponíveis no build (Next inline no bundle).
# Coolify passa estas envs como ARG durante o docker build.
ARG NEXT_PUBLIC_TURNSTILE_SITE_KEY
ENV NEXT_PUBLIC_TURNSTILE_SITE_KEY=$NEXT_PUBLIC_TURNSTILE_SITE_KEY

RUN pnpm --filter @plenya/site build

###
### runner — minimal production image (Next standalone, sem pnpm/node_modules de runtime)
###
FROM node:24-alpine AS runner
WORKDIR /repo

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3002

# Standalone traz server.js + node_modules tree-shaken (inclui @plenya/brand e @plenya/ui
# transpilados, via outputFileTracingRoot na raiz do monorepo).
COPY --from=builder /repo/apps/site/.next/standalone ./
# static e public NÃO entram no standalone — copiar pro lado do server.js do app.
COPY --from=builder /repo/apps/site/.next/static ./apps/site/.next/static
COPY --from=builder /repo/apps/site/public ./apps/site/public
# Conteúdo lido em runtime via process.cwd()/content (blog, testimonials, doctors, data) —
# cwd em runtime é /repo/apps/site (WORKDIR abaixo).
COPY --from=builder /repo/apps/site/content ./apps/site/content

EXPOSE 3002

WORKDIR /repo/apps/site
CMD ["node", "server.js"]
