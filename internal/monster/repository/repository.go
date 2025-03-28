package repository

import (
	"github.com/google/uuid"
	"github.com/v1c-g4b/digital-monster-api/internal/monster/entity"
	"gorm.io/gorm"
)

func FindMonsterById(db *gorm.DB, id uuid.UUID) (*entity.Monster, error) {
	var monster entity.Monster

	if err := db.Where("ID = ?", id).First(&monster).Error; err != nil {
		return nil, err
	}

	return &monster, nil
}

func SaveMonster(db *gorm.DB, m *entity.Monster) error {
	return db.Save(m).Error
}
