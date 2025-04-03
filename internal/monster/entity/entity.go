package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Monster struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;not null"`
	Hunger      int       `gorm:"not null;check:hunger >= 0 AND hunger <= 100"`
	Happiness   int       `gorm:"not null;check:happiness >= 0 AND happiness <= 100"`
	Energy      int       `gorm:"not null;check:energy >= 0 AND energy <= 100"`
	Health      int       `gorm:"not null;check:health >= 0 AND health <= 100"`
	Age         int       `gorm:"not null"`
	Stage       string    `gorm:"size:50;not null"`
	IsAlive     bool      `gorm:"not null"`
	LastUpdated time.Time `gorm:"not null"`

	domainEvents []DomainEvent
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

	if m.Energy >= 100 {
		m.Energy = 100
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

	if m.Energy <= 0 {
		m.Energy = 0
	}

	if m.Happiness >= 100 {
		m.Happiness = 100
	}

	if m.Hunger >= 100 {
		m.Hunger = 100
	}

	return nil
}

func (m *Monster) Sleep() error {

	if !m.IsAlive {
		return errors.New("monster is dead")
	}

	m.Energy += 30

	if m.Energy >= 100 {
		m.Energy = 100
	}

	m.Hunger += 5

	if m.Hunger >= 100 {
		m.Hunger = 100
	}

	return nil
}

func (m *Monster) AddEvent(event DomainEvent) {
	m.domainEvents = append(m.domainEvents, event)
}

func (m *Monster) PullEvents() []DomainEvent {
	events := m.domainEvents
	m.domainEvents = []DomainEvent{}
	return events
}

func (m *Monster) Die() {
	if m.IsAlive {
		m.IsAlive = false
		m.AddEvent(MonsterDiedEvent{MonsterID: m.ID.String()})
	}
}
