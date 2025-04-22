package models

import (
	"time"
)


type Transaction struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	SourceID      int64     `json:"source_id" gorm:"not null"`
	DestinationID int64     `json:"destination_id" gorm:"not null"`
	Amount        float64   `json:"amount" gorm:"not null;default:0"`
	Type          string    `json:"type" gorm:"not null;check:type IN ('credit','debit')"`
	Status        string    `json:"status" gorm:"not null;check:status IN ('pending','completed','failed')"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Source      Account `gorm:"foreignKey:SourceID"`
	Destination Account `gorm:"foreignKey:DestinationID"`
}


