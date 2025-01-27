package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

// aplica apenas as migrações pendentes
func RunMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("erro ao configurar o driver PostgreSQL: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("erro ao configurar migração: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("erro ao aplicar migrações: %w", err)
	}

	log.Println("Migrações pendentes aplicadas com sucesso")
	return nil
}

// remove todas as tabelas e reaplica todas as migrações
func MigrateFresh(db *sqlx.DB) error {
	log.Println("Resetando banco de dados (Drop e reaplicação)...")
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("erro ao configurar o driver PostgreSQL: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("erro ao configurar migração: %w", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Erro ao rodar Down: %v", err)
		return fmt.Errorf("erro ao limpar banco de dados: %w", err)
	}
	
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Erro ao rodar Up: %v", err)
		return fmt.Errorf("erro ao aplicar migrações fresh: %w", err)
	}
	

	log.Println("Banco de dados resetado e migrações reaplicadas com sucesso")
	return nil
}

// reverte todas as migrações na ordem inversa
func MigrateReset(db *sqlx.DB) error {
	log.Println("Revertendo todas as migrações...")
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("erro ao configurar o driver PostgreSQL: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("erro ao configurar migração: %w", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("erro ao reverter migrações: %w", err)
	}

	log.Println("Todas as migrações revertidas com sucesso")
	return nil
}
