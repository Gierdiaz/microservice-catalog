#!/bin/bash
# Aguardar o PostgreSQL ficar disponível para conexões

# Tentando conectar ao banco de dados até ele estar disponível
until nc -z -v -w30 $DB_HOST $DB_PORT; do
  echo "Aguardando banco de dados em $DB_HOST:$DB_PORT..."
  sleep 1
done

echo "Banco de dados disponível. Executando migrações."

# Agora, executa as migrações
go run cmd/main.go migrate-fresh
