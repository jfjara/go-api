# Etapa 1: build
FROM golang:1.25.4-alpine AS builder

# Instalar git y ca-certificates si hace falta para dependencias
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copiamos go.mod y go.sum primero para cachear dependencias
COPY go.mod go.sum ./
COPY config/ ./config/
RUN go mod download

# Copiamos el resto del proyecto
COPY . .

# Build del binario
RUN go build -o mi-api ./cmd/api

# Etapa 2: runtime
FROM alpine:3.18

# Instalar certificados raíz
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copiamos el binario desde la etapa builder
COPY --from=builder /app/config ./config
COPY --from=builder /app/mi-api .

# Puerto en el que corre tu API
EXPOSE 3000

# Comando por defecto
CMD ["./mi-api"]
