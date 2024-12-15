# Etapa de desenvolvimento
FROM golang:1.23.4-alpine AS dev

WORKDIR /go/src/app

RUN go install github.com/air-verse/air@latest


RUN apk add --no-cache bash curl git && \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s && \
    mv ./bin/air /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

COPY . .

CMD ["air", "-c", ".air.toml"]

# Etapa de produção
FROM golang:1.23.4-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify && go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/app

COPY --from=builder /go/src/app/server .
COPY --from=builder /go/src/app/.env .

EXPOSE 8080

ENTRYPOINT ["./server"]
