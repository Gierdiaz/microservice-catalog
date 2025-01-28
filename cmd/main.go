package main

import (
	"log"

	"github.com/Gierdiaz/Book/config"
	"github.com/Gierdiaz/Book/internal/database"
	"github.com/Gierdiaz/Book/internal/endpoints"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	db, err := database.InitDatabase(config)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	err = database.MigrateFresh(db)
	if err != nil {
		log.Fatalf("Erro ao rodar as migrações fresh: %v", err)
	}

	router := endpoints.InitRouter(config, db)
	router.Run(":" + config.Server.APP_PORT)

}
