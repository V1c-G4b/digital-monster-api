package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/repository"
	"gorm.io/gorm"
)

type FeedMonsterUseCase struct {
	DB *gorm.DB
}

func NewFeedMonsterUseCase(db *gorm.DB) *FeedMonsterUseCase {
	return &FeedMonsterUseCase{DB: db}
}

func (uc *FeedMonsterUseCase) Execute(id uuid.UUID) (*entity.Monster, error) {
	m, err := repository.FindMonsterById(uc.DB, id)
	if err != nil {
		return nil, err
	}

	if !m.IsAlive {
		return nil, errors.New("your monster is dead")
	}

	m.Feed()

	if err := uc.DB.Save(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}
