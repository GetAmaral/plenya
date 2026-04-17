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
RUN pnpm install --frozen-lockfile --filter "@plenya/site..."

###
### builder — copy source + build
###
FROM base AS builder
COPY --from=deps /repo/node_modules ./node_modules
COPY --from=deps /repo/apps/site/node_modules ./apps/site/node_modules
COPY --from=deps /repo/packages/brand/node_modules ./packages/brand/node_modules
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1
RUN pnpm --filter @plenya/site build

###
### runner — minimal production image
###
FROM node:24-alpine AS runner
RUN corepack enable && corepack prepare pnpm@10.28.1 --activate
WORKDIR /repo

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3002

COPY --from=builder /repo/package.json /repo/pnpm-lock.yaml /repo/pnpm-workspace.yaml /repo/turbo.json ./
COPY --from=builder /repo/apps/site/package.json ./apps/site/
COPY --from=builder /repo/packages/brand ./packages/brand
COPY --from=builder /repo/apps/site/.next ./apps/site/.next
COPY --from=builder /repo/apps/site/public ./apps/site/public
COPY --from=builder /repo/apps/site/next.config.ts ./apps/site/
COPY --from=builder /repo/apps/site/content ./apps/site/content

RUN pnpm install --prod --frozen-lockfile --filter "@plenya/site..."

EXPOSE 3002

WORKDIR /repo/apps/site
CMD ["pnpm", "start"]
