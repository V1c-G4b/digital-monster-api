package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/repository"
	"gorm.io/gorm"
)

type SleepMonsterUseCase struct {
	DB *gorm.DB
}

func NewSleepMonsterUseCase(db *gorm.DB) *SleepMonsterUseCase {
	return &SleepMonsterUseCase{DB: db}
}

func (uc *SleepMonsterUseCase) Execute(id uuid.UUID) (*entity.Monster, error) {
	m, err := repository.FindMonsterById(uc.DB, id)
	if err != nil {
		return nil, err
	}

	if !m.IsAlive {
		return nil, errors.New("your monster is dead")
	}

	m.Sleep()

	if err := uc.DB.Save(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}
