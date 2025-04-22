package models

import ("time")

type Account struct {
	ID           int64        `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountNo    string       `json:"account_no" gorm:"not null;unique"`
	AccountType  string       `json:"account_type" gorm:"not null"`
	Balance      float64      `json:"balance" gorm:"not null;default:0"`
	Branch       string       `json:"branch" gorm:"not null"`
	UserID       int64        `json:"user_id" gorm:"not null"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:SourceID"` // or DestinationID if needed
}
