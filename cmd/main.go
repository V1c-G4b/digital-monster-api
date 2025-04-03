package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/v1c-g4b/digital-monster-api/config"
	db "github.com/v1c-g4b/digital-monster-api/db/migrations"
	"github.com/v1c-g4b/digital-monster-api/internal/monster"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
)

func main() {

	dispatcher := entity.NewEventDispatcher()

	dispatcher.Register("MonsterDied", func(event entity.DomainEvent) {
		e := event.(entity.MonsterDiedEvent)
		fmt.Printf("🟣 Evento: O monstro %s morreu!\n", e.MonsterID)
	})

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("erro ao carregar confi: %v", err)
	}

	println(cfg.Database.Path)

	dbConn := db.Connect(cfg.Database.Path)

	dbConn.AutoMigrate(&entity.Monster{})

	go monster.StartDecayRoutine(dbConn)

	router := gin.Default()
	router.Use(cors.Default())

	monster.RegisterRoutes(router, dbConn)

	router.Run(":" + cfg.Server.Port)
}
