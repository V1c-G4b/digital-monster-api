package monster

import (
	"time"

	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
	"gorm.io/gorm"
)

func StartDecayRoutine(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Minute)
	dispatcher := entity.NewEventDispatcher()
	for range ticker.C {
		var monster []entity.Monster

		db.Where("is_alive= ?", true).Find(&monster)

		for _, m := range monster {
			m.Hunger += 5
			m.Happiness -= 3
			m.Energy -= 2

			if m.Happiness < 0 {
				m.Happiness = 0
			}

			if m.Energy < 0 {
				m.Energy = 0
			}

			if m.Hunger > 100 {
				m.Hunger = 100
				m.Health -= 5
			}

			if m.Health <= 0 && m.IsAlive {
				m.IsAlive = false
				dispatcher.Dispatch(entity.MonsterDiedEvent{MonsterID: m.ID.String()})
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
