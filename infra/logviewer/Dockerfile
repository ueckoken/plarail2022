FROM golang:1.20.3-bullseye AS builder

ENV CGO_ENABLED=0
WORKDIR /app
COPY . /app
RUN GOARCH=amd64 CGO_ENABLED=0 go build -a -tags "netgo" -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o app main.go

FROM gcr.io/distroless/static-debian11:nonroot AS runner
COPY --from=builder --chown=nonroot:nonroot /app/app /app
ENTRYPOINT [ "/app" ]
