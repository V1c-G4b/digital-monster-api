package monster

import (
	"time"

	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
	"gorm.io/gorm"
)

func StartDecayRoutine(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		var monsters []entity.Monster

		db.Where("is_alive= ?", true).Find(&monsters)

		for _, m := range monsters {
			m.Hunger += 5
			m.Happiness -= 3
			m.Energy -= 2

			if m.Hunger > 100 {
				m.Health -= 5
			}

			if m.Health <= 0 {
				m.IsAlive = false
			}

			m.LastUpdated = time.Now()
			db.Save(&m)
		}
	}
}

func StartExploreRoutine(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		var monsters []entity.Monster

		db.Where("is_alive= ?", true).Find(&monsters)
	}
}
