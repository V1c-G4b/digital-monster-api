package monster

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindMonsterById(db *gorm.DB, id uuid.UUID) (*Monster, error) {
	var monster Monster

	if err := db.Where("ID = ?", id).First(&monster).Error; err != nil {
		return nil, err
	}

	return &monster, nil
}

func SaveMonster(db *gorm.DB, m *Monster) error {
	return db.Save(m).Error
}
