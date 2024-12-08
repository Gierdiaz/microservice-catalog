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