package models

import "gorm.io/gorm"

const (
	Income  string = "income"
	Outcome string = "outcome"
)

type Transaction struct {
	gorm.Model
	Type   string `gorm:"not null" json:"type"`
	UserId uint32  `gorm:"not null" json:"user_id"`
}

func (*Transaction) GetTransactionsByUser(db *gorm.DB, uid uint32) ([]Transaction, error) {
	var transactions []Transaction
	result := db.Find(&transactions).Where("user_id = ?", uid)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}
