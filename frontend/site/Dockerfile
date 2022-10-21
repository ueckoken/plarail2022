# Build the source code
FROM node:16.17.1-buster-slim AS builder
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile --immutable
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1
ENV NODE_ENV production
RUN yarn next build
RUN yarn next export

# Production image, copy all the files and run nginx
FROM nginxinc/nginx-unprivileged:1.23.1-alpine AS runner
COPY --from=builder --chown=nginx:nginx /app/out/ /usr/share/nginx/html/
