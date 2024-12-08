# Variáveis
DOCKER_COMPOSE = docker-compose
ENV_FILE = .env

# Comandos
.PHONY: build up down logs restart clean prune db-migrate

## Roda o microserviço
run:
	go run cmd/server/main.go

## Roda os testes
test:
	go test ./...

## Formata arquivos .go
fmt:
	go fmt ./...

## Builda as imagens do Docker
build:
	$(DOCKER_COMPOSE) up --build -d

## Sobe os containers em modo detached
up:
	$(DOCKER_COMPOSE) up -d

## Para e remove os containers
down:
	$(DOCKER_COMPOSE) down

## Mostra os logs de todos os containers
logs:
	$(DOCKER_COMPOSE) logs -f

## Reinicia os containers
restart: down up

## Remove volumes, containers e imagens não utilizados
clean:
	$(DOCKER_COMPOSE) down -v --remove-orphans

## Remove tudo não utilizado pelo Docker (atenção, isto remove volumes também)
prune:
	docker system prune -a --volumes --force

## Roda as migrações no banco de dados
db-migrate:
	docker exec postgres_db bash -c "psql -U $$POSTGRES_USER -d $$POSTGRES_DB -f /app/infrastructure/database/migrations/*.sql"