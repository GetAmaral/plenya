FROM golang:1.23-alpine AS builder

WORKDIR /app

# Dependências
COPY go.mod go.sum* ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server ./cmd/server

# Runtime (imagem final ~15MB)
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /server .

EXPOSE 3001

CMD ["./server"]
