package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kelas struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primary_key"`
	Name, WaliKelas string
}
