FROM golang:1.19.2-bullseye AS builder

ENV CGO_ENABLED=0
WORKDIR /control
COPY ./go.mod ./go.sum /control/
RUN go mod download
COPY . /control/
RUN make build

FROM gcr.io/distroless/static-debian11:nonroot AS runner
COPY --from=builder --chown=nonroot:nonroot /control/app /internal
ENTRYPOINT [ "/internal" ]
