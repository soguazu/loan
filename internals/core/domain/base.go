package domain

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

// Base entity that is reused in all entity
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate hooks run to before database insertion occurs to populate the ID field
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewV4()
	return
}
