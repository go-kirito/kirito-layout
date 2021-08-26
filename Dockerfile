FROM golang:1.16 AS builder

COPY . /build
WORKDIR /build

RUN make build

FROM alpine

COPY --from=builder /build/bin /app
COPY --from=builder /build/config /app/config

WORKDIR /app

EXPOSE 8000
EXPOSE 9000

CMD ["./app", "-f", "config/config.yaml"]