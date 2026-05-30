.PHONY: help dev build test clean migrate-up migrate-down migrate-status migrate-create swagger

help: ## Mostrar ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Rodar servidor em modo desenvolvimento
	go run cmd/server/main.go

build: ## Build do binário de produção
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o bin/server cmd/server/main.go

test: ## Rodar testes
	go test -v ./...

clean: ## Limpar binários e cache
	rm -rf bin/
	go clean

# Migrations (goose) — ver docs/emr/migrations-decisao.md
# Rodam dentro do container: docker compose exec -w /app api go run ./cmd/migrate <cmd>
migrate-up: ## Aplicar migrations pendentes
	docker compose exec -w /app api go run ./cmd/migrate up

migrate-down: ## Reverter última migration (apenas dev)
	docker compose exec -w /app api go run ./cmd/migrate down

migrate-status: ## Ver status das migrations
	docker compose exec -w /app api go run ./cmd/migrate status

migrate-create: ## Criar nova migration SQL (use NAME=nome) no formato goose
	@ts=$$(date +%Y%m%d%H%M%S); f="database/migrations/$${ts}_$(NAME).sql"; \
	printf -- '-- +goose Up\n-- +goose StatementBegin\n\n-- +goose StatementEnd\n\n-- +goose Down\n-- +goose StatementBegin\n\n-- +goose StatementEnd\n' > $$f; \
	echo "criada $$f"

# Swagger
swagger: ## Gerar documentação Swagger
	swag init -g cmd/server/main.go -o docs

# Dependências
deps: ## Instalar dependências
	go mod tidy
	go mod download

# Docker
docker-build: ## Build da imagem Docker
	docker build -t plenya-api:latest .

docker-run: ## Rodar container Docker
	docker run -p 3001:3001 --env-file .env plenya-api:latest

# Scripts utilitários
update-gender: ## Atualizar campo gender em score_items baseado no nome
	go run cmd/update-score-item-gender/main.go

test-gender: ## Testar lógica de detecção de gênero
	go test -v ./cmd/update-score-item-gender/

update-age-range: ## Atualizar campos ageRangeMin/Max em score_items baseado no nome
	go run cmd/update-score-item-age-range/main.go

test-age-range: ## Testar lógica de detecção de faixa etária
	go test -v ./cmd/update-score-item-age-range/
