package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PhoneNumber struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	PhoneNumber string    `gorm:"not null;uniqueIndex:unique_number_idx"`
	Provider    string
}

func (phoneNumber *PhoneNumber) BeforeCreate(tx *gorm.DB) (err error) {
	phoneNumber.ID = uuid.New()
	return
}
