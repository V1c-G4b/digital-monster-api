package entity

type DomainEvent interface {
	EventName() string
}
