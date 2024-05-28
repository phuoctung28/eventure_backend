package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type EventType string

type Event struct {
	gorm.Model
	ID             string `gorm:"type:uuid;primary_key;"`
	Name           string `gorm:"type:longtext"`
	Description    string `gorm:"type:longtext"`
	Summary        string `gorm:"type:longtext"`
	Type           EventType
	StartDateTime  time.Time
	EndDateTime    time.Time
	AddressID      uint
	DeletedAt      gorm.DeletedAt
	OrganizationID uint
	IsPublished    bool
}

func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}
