FROM golang:1.25-alpine AS builder

WORKDIR /app

# Dependências
COPY go.mod go.sum* ./
RUN go mod download

# Build (templates HTML são embedados via go:embed)
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server ./cmd/server \
 && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /backfill-patient-phones ./cmd/backfill-patient-phones

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
COPY --from=builder /backfill-patient-phones .

# Volumes: uploads (PDFs), enrichment-batch (RAG)
RUN mkdir -p /app/uploads /app/enrichment-batch
VOLUME ["/app/uploads", "/app/enrichment-batch"]

EXPOSE 3001

CMD ["./server"]
