package monster

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Monster struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Hunger      int       `gorm:"type:int;not null;check:hunger >= 0 AND hunger <= 100"`
	Happiness   int       `gorm:"type:int;not null;check:happiness >= 0 AND happiness <= 100"`
	Energy      int       `gorm:"type:int;not null;check:energy >= 0 AND energy <= 100"`
	Health      int       `gorm:"type:int;not null;check:health >= 0 AND health <= 100"`
	Age         int       `gorm:"type:int;not null"`
	Stage       string    `gorm:"type:varchar(50);not null"`
	IsAlive     bool      `gorm:"type:boolean;not null"`
	LastUpdated time.Time `gorm:"type:timestamp;not null"`
}

func (m *Monster) Feed() error {

	if !m.IsAlive {
		return errors.New("monster is dead")
	}

	m.Hunger -= 20
	m.Energy += 10
	if m.Hunger < 0 {
		m.Hunger = 0
	}

	return nil
}

func (m *Monster) Play() error {

	if !m.IsAlive {
		return errors.New("monster is dead")
	}

	m.Happiness += 10
	m.Energy -= 15
	m.Hunger += 10

	return nil
}

func (m *Monster) Sleep() error {

	if !m.IsAlive {
		return errors.New("monster is dead")
	}

	m.Energy += 30
	m.Hunger += 5

	return nil
}

// Hunger > 90, Health decreases by 10 per hour

// Hunger > 70, Health decreases by 5 per hour

// Happiness < 30, Energy decreases by 2 per hour

// Health <= 0 or Energy <= 0, Monster dies

// Age >= 100, Monster dies

// Age == 5, Monster evolves to next stage, but need to meet certain conditions. Health, Hunger, Happiness, and Energy must be above 70.
