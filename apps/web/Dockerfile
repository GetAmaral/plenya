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
RUN pnpm install --frozen-lockfile --filter "@plenya/web..."

###
### builder — copy source + build
###
FROM base AS builder
COPY --from=deps /repo/node_modules ./node_modules
COPY --from=deps /repo/apps/web/node_modules ./apps/web/node_modules
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=${NEXT_PUBLIC_API_URL}
RUN pnpm --filter @plenya/web build

###
### runner — minimal production image
###
FROM node:24-alpine AS runner
RUN corepack enable && corepack prepare pnpm@10.28.1 --activate
WORKDIR /repo

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3000

COPY --from=builder /repo/package.json /repo/pnpm-lock.yaml /repo/pnpm-workspace.yaml /repo/turbo.json ./
COPY --from=builder /repo/apps/web/package.json ./apps/web/
COPY --from=builder /repo/packages/brand ./packages/brand
COPY --from=builder /repo/packages/types ./packages/types
COPY --from=builder /repo/apps/web/.next ./apps/web/.next
COPY --from=builder /repo/apps/web/public ./apps/web/public
COPY --from=builder /repo/apps/web/next.config.ts ./apps/web/

RUN pnpm install --prod --frozen-lockfile --filter "@plenya/web..."

EXPOSE 3000

WORKDIR /repo/apps/web
CMD ["pnpm", "start"]
