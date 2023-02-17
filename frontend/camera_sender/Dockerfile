FROM node:18.12-buster as builder
ENV ESP_EYE_IP 192.168.12.199
WORKDIR /app
COPY package.json package-lock.json /app/
RUN npm ci

COPY *.js ./
RUN npm run build

FROM nginxinc/nginx-unprivileged:1.23.3-alpine
WORKDIR /usr/share/nginx/html/
COPY --from=builder /app/sender.bundle.js ./
COPY ./index.html ./
COPY ./index.js ./
COPY ./sender.html ./
