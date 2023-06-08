FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


FROM builder AS DEV
WORKDIR /app

RUN go install github.com/cespare/reflex@latest

COPY . .

CMD reflex -s -r '\.go$$' go run main.go

FROM builder as prod_builder

COPY . .

RUN go build -ldflags "-w -s" -o dauth_v2_server


FROM alpine:3.17 AS PROD
WORKDIR /app

COPY --from=prod_builder /app/dauth_v2_server .
COPY --from=prod_builder /app/.env .
COPY --from=prod_builder /app/entry.sh .

ENTRYPOINT ["/app/entry.sh"]
CMD ["/app/dauth_v2_server"]