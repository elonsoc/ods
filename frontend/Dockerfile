# syntax=docker/dockerfile:1.2
# copied shamelessly from https://github.com/vercel/next.js/blob/canary/examples/with-docker/Dockerfile
# with edits to use pnpm fetch, cribbed from https://pnpm.io/cli/fetch
FROM node:18-alpine AS base

# Install dependencies only when needed
FROM base AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
WORKDIR /app

RUN yarn global add pnpm@8.1.1
# Install dependencies based on the preferred package manager
COPY package.json pnpm-lock.yaml* ./
# RUN pnpm fetch --prod
RUN pnpm fetch --dev

ADD . ./
# RUN pnpm install -r --offline --prod
RUN pnpm install -r --offline --dev

# Rebuild the source code only when needed
FROM base AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

ENV NEXT_TELEMETRY_DISABLED 1

RUN yarn build

# Production image, copy all the files and run next
FROM base AS runner
WORKDIR /app

# ENV NODE_ENV production
ENV NODE_ENV development
# Uncomment the following line in case you want to disable telemetry during runtime.
ENV NEXT_TELEMETRY_DISABLED 1

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public

# Automatically leverage output traces to reduce image size
# https://nextjs.org/docs/advanced-features/output-file-tracing
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000

ENV PORT 3000

CMD ["node", "server.js"]
