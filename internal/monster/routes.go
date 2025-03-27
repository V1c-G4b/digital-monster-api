package monster

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	grp := router.Group("/monsters")
	{
		grp.POST("/", CreateMonster(db))

		grp.GET("/:id", GetMonster(db))

		grp.PATCH("feed/:id", FeedMonster(db))

		grp.PATCH("play/:id", PlayMonster(db))

		grp.PATCH("sleep/:id", SleepMonster(db))
	}
}
