# Variáveis
DOCKER_COMPOSE = docker-compose
ENV_FILE = .env

# Comandos
.PHONY: build up down logs restart clean prune db-migrate db-migrate-fresh db-migrate-reset

## Roda o microserviço
run:
	go run cmd/main.go run

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

## Roda as migrações e aplica novas
migrate:
	$(DOCKER_COMPOSE) exec microservice_catalog go run cmd/main.go migrate

## Roda as migrações e reinicia o banco de dados
migrate-fresh:
	$(DOCKER_COMPOSE) exec microservice_catalog go run cmd/main.go migrate-fresh

## Reverte todas as migrações
migrate-reset:
	$(DOCKER_COMPOSE) exec microservice_catalog go run cmd/main.go migrate-reset