package models

import (
	"gorm.io/gorm"
)

const (
	Income  string = "income"
	Outcome string = "outcome"
)

type Transaction struct {
	gorm.Model
	Type   string `gorm:"not null" json:"type"`
	UserId uint32  `gorm:"not null" json:"user_id"`
    CategoryID uint32
    User
    Category
}

func (*Transaction) GetTransactionsByUser(db *gorm.DB, uid uint32) ([]Transaction, error) {
    var user User

    db.Preload("Transactions").First(&user, 1)

    return user.Transactions, nil
}
