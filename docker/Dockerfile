# Etapa de desenvolvimento
FROM golang:1.23.4-alpine AS dev

WORKDIR /go/src/app

RUN go install github.com/air-verse/air@latest

RUN apk add --no-cache bash curl git

COPY go.mod go.sum ./
RUN go mod tidy && go mod vendor

COPY . .

COPY .air.toml .air.toml

CMD ["air", "-c", ".air.toml"]

# Etapa de produção
FROM golang:1.23.4-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod tidy && go mod vendor

COPY . .

# Build para produção
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/app

COPY --from=builder /go/src/app/server .
COPY --from=builder /go/src/app/.env .

EXPOSE 8080

ENTRYPOINT ["./server"]
