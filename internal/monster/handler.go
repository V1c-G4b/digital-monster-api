package monster

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateMonster(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var m Monster
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

		monsterFound, err := FindMonsterById(db, uuidID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Monster not found"})
			return
		}

		c.JSON(http.StatusOK, monsterFound)
	}
}

func FeedMonster(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		uuidID, err := uuid.Parse(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		}

		monsterFound, err := FindMonsterById(db, uuidID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Monster not found"})
		}

		if err := monsterFound.Feed(); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		db.Save(&monsterFound)
		c.JSON(http.StatusOK, monsterFound)

	}
}

func PlayMonster(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		uuidID, err := uuid.Parse(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		}

		monsterFound, err := FindMonsterById(db, uuidID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Monster not found"})
		}

		monsterFound.Play()

		db.Save(&monsterFound)
		c.JSON(http.StatusOK, monsterFound)

	}
}

func SleepMonster(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		uuidID, err := uuid.Parse(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		}

		monsterFound, err := FindMonsterById(db, uuidID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Monster not found"})
		}

		monsterFound.Sleep()

		db.Save(&monsterFound)
		c.JSON(http.StatusOK, monsterFound)

	}
}
