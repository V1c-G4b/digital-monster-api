package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/v1c-g4b/digital-monster-api/config"
	db "github.com/v1c-g4b/digital-monster-api/db/migrations"
	"github.com/v1c-g4b/digital-monster-api/internal/monster"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("erro ao carregar confi: %v", err)
	}

	dbConn := db.Connect(cfg.Database.Path)

	dbConn.AutoMigrate(&entity.Monster{})

	go monster.StartDecayRoutine(dbConn)

	router := gin.Default()

	monster.RegisterRoutes(router, dbConn)

	router.Run(":" + cfg.Server.Port)
}
