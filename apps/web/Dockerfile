### Production Dockerfile for @plenya/web (EMR)
### Multi-stage. Builds with workspace context (root: ../..).

FROM node:24-alpine AS base
RUN corepack enable && corepack prepare pnpm@10.28.1 --activate
WORKDIR /repo

###
### deps — install full workspace deps (frozen)
###
FROM base AS deps
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml turbo.json ./
COPY apps/web/package.json ./apps/web/
COPY packages/brand/package.json ./packages/brand/
COPY packages/types/package.json ./packages/types/
COPY packages/ui/package.json ./packages/ui/
RUN pnpm install --frozen-lockfile --filter "@plenya/web..."

###
### builder — copy source + build
###
FROM base AS builder
COPY --from=deps /repo/node_modules ./node_modules
COPY --from=deps /repo/apps/web/node_modules ./apps/web/node_modules
# @plenya/ui é transpilado pelo web (transpilePackages) → precisa do node_modules
# dele (react peer) no isolated-install do Docker. (web tem ignoreBuildErrors,
# então o tsc não bloqueia, mas mantém o type-check de packages/ui funcional e
# espelha o Dockerfile do site.)
COPY --from=deps /repo/packages/ui/node_modules ./packages/ui/node_modules
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=${NEXT_PUBLIC_API_URL}
RUN pnpm --filter @plenya/web build

###
### runner — minimal production image (Next standalone, sem pnpm/node_modules de runtime)
###
FROM node:24-alpine AS runner
WORKDIR /repo

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3000

# Standalone traz server.js + node_modules tree-shaken (inclui react-reader e @plenya/ui
# transpilados, via outputFileTracingRoot na raiz do monorepo). Middleware embutido.
COPY --from=builder /repo/apps/web/.next/standalone ./
# static e public NÃO entram no standalone — copiar pro lado do server.js do app.
COPY --from=builder /repo/apps/web/.next/static ./apps/web/.next/static
COPY --from=builder /repo/apps/web/public ./apps/web/public

EXPOSE 3000

WORKDIR /repo/apps/web
CMD ["node", "server.js"]
