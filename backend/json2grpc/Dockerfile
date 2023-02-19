FROM golang:1.20.1-bullseye AS builder

ENV CGO_ENABLED=0
WORKDIR /json2grpc
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . ./
RUN make build

FROM gcr.io/distroless/static-debian11:nonroot AS runner
COPY --from=builder --chown=nonroot:nonroot /json2grpc/json2grpc /app
ENTRYPOINT [ "/app" ]
