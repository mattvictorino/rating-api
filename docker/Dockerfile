# ---------- BUILD STAGE ----------
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Copia go.mod primeiro (cache)
COPY ../go.mod ../go.sum ./
RUN go mod download

# Copia o restante do código
COPY .. .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .


# ---------- RUNTIME STAGE ----------
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app .

EXPOSE 8080

ENV SERVER_PORT=8080

CMD ["./app", "api"]
