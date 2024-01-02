package routes

import (
	"financify/pkg/controllers"
	middlewares "financify/pkg/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var TransactionsRoutes = func(db *gorm.DB, router *mux.Router) {
	transactionsRouter := router.PathPrefix("/transactions").Subrouter()
	transactionsController := controllers.NewTransactionController(db)
	transactionsRouter.Use(middlewares.Protected)
	transactionsRouter.HandleFunc("/", transactionsController.GetTransactionsByUser).Methods("GET")
}
