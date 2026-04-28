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
### runner — minimal production image
###
FROM node:24-alpine AS runner
RUN corepack enable && corepack prepare pnpm@10.28.1 --activate
WORKDIR /repo

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3003

COPY --from=builder /repo/package.json /repo/pnpm-lock.yaml /repo/pnpm-workspace.yaml /repo/turbo.json ./
COPY --from=builder /repo/apps/site-getulio/package.json ./apps/site-getulio/
COPY --from=builder /repo/apps/site-getulio/.next ./apps/site-getulio/.next
COPY --from=builder /repo/apps/site-getulio/public ./apps/site-getulio/public
COPY --from=builder /repo/apps/site-getulio/next.config.ts ./apps/site-getulio/
COPY --from=builder /repo/apps/site-getulio/content ./apps/site-getulio/content
# Posts mirror — site-getulio lê /apps/site/content/blog/pt em runtime SSR.
# Sem este COPY o /escritos retorna [] e exibe "Nenhum artigo nesse pilar ainda".
COPY --from=builder /repo/apps/site/content ./apps/site/content

RUN pnpm install --prod --frozen-lockfile --filter "@plenya/site-getulio..."

EXPOSE 3003

WORKDIR /repo/apps/site-getulio
CMD ["pnpm", "start"]
