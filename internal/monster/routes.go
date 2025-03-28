package monster

import (
	"github.com/gin-gonic/gin"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/handler"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	grp := router.Group("/monsters")
	{
		grp.POST("/", handler.CreateMonster(db))

		grp.GET("/:id", handler.GetMonster(db))

		grp.PATCH("feed/:id", handler.FeedMonster(db))

		grp.PATCH("play/:id", handler.PlayMonster(db))

		grp.PATCH("sleep/:id", handler.SleepMonster(db))
	}
}
