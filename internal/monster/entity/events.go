package entity

type MonsterDiedEvent struct {
	MonsterID string
}

func (e MonsterDiedEvent) EventName() string {
	return "MonsterDied"
}
