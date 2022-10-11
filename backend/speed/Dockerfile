FROM golang:1.19.2-bullseye AS builder

ENV CGO_ENABLED=0
WORKDIR /speed
COPY go.mod go.sum /speed/
RUN go mod download
COPY . /speed
RUN go build -o /speed/app ./cmd/main.go

FROM gcr.io/distroless/static-debian11:nonroot as runner
COPY --from=builder --chown=nonroot:nonroot /speed/app /speed-app
CMD ["/speed-app"]
