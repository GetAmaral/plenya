FROM golang:1.25-alpine AS builder

WORKDIR /app

# Dependências
COPY go.mod go.sum* ./
RUN go mod download

# Build (templates HTML são embedados via go:embed)
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server ./cmd/server \
 && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /migrate ./cmd/migrate \
 && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /backfill-patient-phones ./cmd/backfill-patient-phones \
 && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /import-cmed ./cmd/import-cmed

# Runtime (imagem final com OCR dependencies)
FROM alpine:3.21

# Dependências de sistema: OCR/processamento de PDF + Chromium (geração de PDF vetorial via go-rod)
RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    poppler-utils \
    tesseract-ocr \
    tesseract-ocr-data-por \
    imagemagick \
    # ffmpeg: transcodifica notas de voz do WhatsApp (OGG/Opus) p/ AAC/MP4 ao servir —
    # Safari/iOS não toca o container OGG, então sem isto o áudio não roda no iPhone.
    ffmpeg \
    # Chromium headless p/ o pacote pdfdoc (papelaria world-class) — assets vão embutidos no binário
    chromium \
    nss \
    freetype \
    harfbuzz \
    ttf-freefont

WORKDIR /app

COPY --from=builder /server .
COPY --from=builder /migrate .
COPY --from=builder /backfill-patient-phones .
# Import mensal do catálogo de medicamentos (Lista de Preços CMED/ANVISA). Roda por
# docker exec com o XLSX copiado para /tmp — ver o comentário em cmd/import-cmed.
COPY --from=builder /import-cmed .
COPY prod-entrypoint.sh .
RUN chmod +x prod-entrypoint.sh

# Assets de geração de PDF (gofpdf lê estes caminhos absolutos em runtime).
# Sem eles, TODO PDF falha: pedidos de exame, prescrições, documentos, pagamentos.
#   - Fontes OpenSans (Regular/Bold/SemiBold) em /usr/share/fonts/opensans/
#   - Timbrado, selo ICP-Brasil, logo e watermark em /app/ e /app/assets/
COPY assets/fonts/opensans/ /usr/share/fonts/opensans/
COPY PlenyaA4-150dpi.png icp-brasil-seal.png /app/
COPY assets/logo-only.png assets/watermark-only.png /app/assets/

# Diretórios de dados persistentes: uploads (mídia WhatsApp + patient-docs) e
# enrichment-batch (RAG). NÃO usar a diretiva VOLUME aqui: ela faz o Docker criar
# um volume ANÔNIMO novo a cada deploy quando não há mount explícito, órfãnando os
# dados (mídia/exames somem após redeploy). A persistência é garantida por bind
# mount configurado no Coolify (local_persistent_volumes → /data/coolify/.../uploads).
RUN mkdir -p /app/uploads /app/enrichment-batch

EXPOSE 3001

# prod-entrypoint aplica migrations (goose) só se RUN_MIGRATIONS=true; senão só sobe o server.
CMD ["./prod-entrypoint.sh"]
