package controllers

import (
	"encoding/json"
	"financify/pkg/models"
	"financify/pkg/utils"
	"net/http"

	"gorm.io/gorm"
)

type TransactionController struct {
	DB *gorm.DB
}

func NewTransactionController(db *gorm.DB) *TransactionController {
	return &TransactionController{DB: db}
}

func (tc *TransactionController) GetTransactionsByUser(w http.ResponseWriter, r *http.Request) {
	uid, err := utils.ExtractUserID(r)

	if err != nil {
		http.Error(w, "Error extracting user id", http.StatusInternalServerError)
	}
	userTransactions := models.Transaction{}

	result, err := userTransactions.GetTransactionsByUser(tc.DB, uid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
