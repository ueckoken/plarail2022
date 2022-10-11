FROM golang:1.19.2-bullseye as builder

ENV CGO_ENABLED=0
WORKDIR /positioning
COPY ./go.mod ./go.sum /positioning/
RUN go mod download
COPY . /positioning
RUN make build

FROM gcr.io/distroless/static-debian11:nonroot AS runner
COPY --from=builder --chown=nonroot:nonroot /positioning/main /positioning
ENTRYPOINT [ "/positioning" ]
