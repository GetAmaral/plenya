FROM golang:1.25-alpine AS builder

WORKDIR /app

# Dependências
COPY go.mod go.sum* ./
RUN go mod download

# Build (templates HTML são embedados via go:embed)
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server ./cmd/server \
 && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /migrate ./cmd/migrate \
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
COPY --from=builder /migrate .
COPY --from=builder /backfill-patient-phones .
COPY prod-entrypoint.sh .
RUN chmod +x prod-entrypoint.sh

# Assets de geração de PDF (gofpdf lê estes caminhos absolutos em runtime).
# Sem eles, TODO PDF falha: pedidos de exame, prescrições, documentos, pagamentos.
#   - Fontes OpenSans (Regular/Bold/SemiBold) em /usr/share/fonts/opensans/
#   - Timbrado, selo ICP-Brasil, logo e watermark em /app/ e /app/assets/
COPY assets/fonts/opensans/ /usr/share/fonts/opensans/
COPY PlenyaA4-150dpi.png icp-brasil-seal.png /app/
COPY assets/logo-only.png assets/watermark-only.png /app/assets/

# Volumes: uploads (PDFs), enrichment-batch (RAG)
RUN mkdir -p /app/uploads /app/enrichment-batch
VOLUME ["/app/uploads", "/app/enrichment-batch"]

EXPOSE 3001

# prod-entrypoint aplica migrations (goose) só se RUN_MIGRATIONS=true; senão só sobe o server.
CMD ["./prod-entrypoint.sh"]
