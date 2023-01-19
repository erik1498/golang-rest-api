package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kelas struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid"`
	Name, WaliKelas string
}
