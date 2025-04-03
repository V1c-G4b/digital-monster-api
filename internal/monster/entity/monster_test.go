package entity

import (
	"testing"

	"github.com/google/uuid"
)

func makeMonsterWithID(hunger, energy, happiness, health int, isAlive bool) Monster {
	return Monster{
		ID:        uuid.New(),
		Hunger:    hunger,
		Energy:    energy,
		Happiness: happiness,
		Health:    health,
		IsAlive:   isAlive,
	}
}

func TestFeed_TableDriven(t *testing.T) {
	tests := []struct {
		name              string
		monster           Monster
		expectedHunger    int
		expectedEnergy    int
		expectedHappiness int
	}{
		{"Feed básico", makeMonsterWithID(50, 50, 50, 50, true), 30, 60, 50},
		{"Feed com fome baixa", makeMonsterWithID(10, 50, 50, 50, true), 0, 60, 50},
		{"Monstro morto", makeMonsterWithID(50, 50, 50, 50, false), 50, 50, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.monster.Feed()

			if tt.monster.Hunger != tt.expectedHunger {
				t.Errorf("Hunger esperado %d, obteve %d", tt.expectedHunger, tt.monster.Hunger)
			}

			if tt.monster.Energy != tt.expectedEnergy {
				t.Errorf("Energy esperado %d, obteve %d", tt.expectedEnergy, tt.monster.Energy)
			}
		})
	}
}

func TestSleep_TableDriven(t *testing.T) {
	tests := []struct {
		name              string
		monster           Monster
		expectedHunger    int
		expectedEnergy    int
		expectedHappiness int
	}{
		{"Sleep Básico", makeMonsterWithID(50, 50, 50, 50, true), 55, 80, 50},
		{"Sleep Com fome ao extremo", makeMonsterWithID(100, 50, 50, 50, true), 100, 80, 50},
		{"Sleep com energia ao extremo", makeMonsterWithID(50, 100, 50, 50, true), 55, 100, 50},
		{"Monstro Morto", makeMonsterWithID(50, 50, 50, 50, false), 50, 50, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.monster.Sleep()

			if tt.monster.Hunger != tt.expectedHunger {
				t.Errorf("Hunger esperado %d, obteve %d", tt.expectedHunger, tt.monster.Hunger)
			}

			if tt.monster.Energy != tt.expectedEnergy {
				t.Errorf("Energy esperado %d, obteve %d", tt.expectedEnergy, tt.monster.Energy)
			}
		})
	}
}

//10, 15, 10

func TestPlay_TableDriven(t *testing.T) {
	tests := []struct {
		name              string
		monster           Monster
		expectedHunger    int
		expectedEnergy    int
		expectedHappiness int
	}{
		{"Play Básico", makeMonsterWithID(50, 50, 50, 50, true), 60, 35, 60},
		{"Play Com fome ao extremo", makeMonsterWithID(100, 50, 50, 50, true), 100, 35, 60},
		{"Play com energia ao extremo", makeMonsterWithID(50, 100, 50, 50, true), 60, 85, 60},
		{"Play com happiness ao extremo", makeMonsterWithID(50, 50, 100, 50, true), 60, 35, 100},
		{"Monstro Morto", makeMonsterWithID(50, 50, 50, 50, false), 50, 50, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.monster.Play()

			if tt.monster.Hunger != tt.expectedHunger {
				t.Errorf("Hunger esperado %d, obteve %d", tt.expectedHunger, tt.monster.Hunger)
			}

			if tt.monster.Energy != tt.expectedEnergy {
				t.Errorf("Energy esperado %d, obteve %d", tt.expectedEnergy, tt.monster.Energy)
			}

			if tt.monster.Happiness != tt.expectedHappiness {
				t.Errorf("Happiness esperado %d, obteve %d", tt.expectedHappiness, tt.monster.Happiness)
			}

		})
	}
}

func TestMonsterDie_DispatchEvent(t *testing.T) {
	m := makeMonsterWithID(50, 50, 50, 50, true)

	m.Die()

	if len(m.domainEvents) != 1 {
		t.Fatalf("esperava 1 evento, obteve %d", len(m.domainEvents))
	}

	event, ok := m.domainEvents[0].(MonsterDiedEvent)
	if !ok {
		t.Fatal("evento não é do tipo MonsterDiedEvent")
	}

	if event.MonsterID != m.ID.String() {
		t.Errorf("esperava ID %s no evento, obteve %s", m.ID.String(), event.MonsterID)
	}
}

func TestDie_DispatchEvent(t *testing.T) {
	m := makeMonsterWithID(50, 50, 50, 50, true)
	m.ID = uuid.New()

	m.Health = 0
	m.Die()

	if !m.IsAlive {
		t.Log("✅ monstro morreu corretamente")
	} else {
		t.Error("❌ monstro deveria estar morto")
	}

	if len(m.domainEvents) != 1 {
		t.Fatalf("esperava 1 evento, obteve %d", len(m.domainEvents))
	}

	event, ok := m.domainEvents[0].(MonsterDiedEvent)
	if !ok {
		t.Fatal("evento não é do tipo MonsterDiedEvent")
	}

	if event.MonsterID != m.ID.String() {
		t.Errorf("esperava ID %s no evento, obteve %s", m.ID.String(), event.MonsterID)
	}
}
