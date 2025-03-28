package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/repository"
	"gorm.io/gorm"
)

type PlayMonsterUseCase struct {
	DB *gorm.DB
}

func NewPlayMonsterUseCase(db *gorm.DB) *PlayMonsterUseCase {
	return &PlayMonsterUseCase{DB: db}
}

func (uc *PlayMonsterUseCase) Execute(id uuid.UUID) (*entity.Monster, error) {
	m, err := repository.FindMonsterById(uc.DB, id)
	if err != nil {
		return nil, err
	}

	if !m.IsAlive {
		return nil, errors.New("your monster is dead")
	}

	m.Play()

	if err := uc.DB.Save(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}
