FROM golang:1.23-alpine AS builder

WORKDIR /app

# Dependências
COPY go.mod go.sum* ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server ./cmd/server

# Runtime (imagem final com OCR dependencies)
FROM alpine:3.21

# Instalar dependências de sistema para OCR e processamento de PDF
RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    poppler-utils \
    tesseract-ocr \
    tesseract-ocr-data-por \
    imagemagick

WORKDIR /app

COPY --from=builder /server .

EXPOSE 3001

CMD ["./server"]
