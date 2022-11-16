FROM node:14.21.1-buster-slim as builder
ENV ESP_EYE_IP 192.168.12.199
WORKDIR /app
COPY package.json package-lock.json /app/
RUN npm ci

COPY *.js /app/
RUN npm run build

FROM nginxinc/nginx-unprivileged:1.23.2-alpine
COPY --from=builder /app/sender.bundle.js /usr/share/nginx/html/
COPY ./index.html /usr/share/nginx/html/
COPY ./index.js /usr/share/nginx/html/
COPY ./sender.html /usr/share/nginx/html/