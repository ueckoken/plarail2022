FROM node:14.20.1-buster-slim as builder

WORKDIR /app
COPY package.json package-lock.json /app/
RUN npm ci

COPY *.js /app/
RUN npm run build

FROM nginxinc/nginx-unprivileged:1.23.2-alpine
COPY --from=builder /app/bundle.js /usr/share/nginx/html/
COPY ./index.html /usr/share/nginx/html/
