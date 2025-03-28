package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/repository"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/usecase"
	"gorm.io/gorm"
)

func extractMonsterID(c *gin.Context) (uuid.UUID, error) {
	id := c.Param("id")
	return uuid.Parse(id)
}

func CreateMonster(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var m entity.Monster
		if err := c.ShouldBindJSON(&m); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		m.ID = uuid.New()
		m.IsAlive = true
		m.LastUpdated = time.Now()
		db.Create(&m)
		c.JSON(http.StatusCreated, m)
	}
}

func GetMonster(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		uuidID, err := uuid.Parse(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}

		monsterFound, err := repository.FindMonsterById(db, uuidID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Monster not found"})
			return
		}

		c.JSON(http.StatusOK, monsterFound)
	}
}

func FeedMonster(db *gorm.DB) gin.HandlerFunc {

	useCase := usecase.NewFeedMonsterUseCase(db)

	return func(c *gin.Context) {
		id, err := extractMonsterID(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}

		monster, err := useCase.Execute(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, monster)

	}
}

func PlayMonster(db *gorm.DB) gin.HandlerFunc {
	useCase := usecase.NewPlayMonsterUseCase(db)

	return func(c *gin.Context) {
		id, err := extractMonsterID(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}

		monster, err := useCase.Execute(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, monster)
	}
}

func SleepMonster(db *gorm.DB) gin.HandlerFunc {
	useCase := usecase.NewSleepMonsterUseCase(db)

	return func(c *gin.Context) {
		id, err := extractMonsterID(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
			return
		}

		monster, err := useCase.Execute(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, monster)
	}
}
