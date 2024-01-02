package routes

import (
	"financify/pkg/controllers"
	middlewares "financify/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var UserRoutes = func(db *gorm.DB, router *mux.Router) {
	userController := controllers.NewUserController(db)
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/register", userController.Register).Methods("POST")
	userRouter.HandleFunc("/login", userController.Login).Methods("POST")
	userRouter.HandleFunc("/avatar", userController.UpdateAvatar).Methods("POST").Name("updateAvatar").Handler(middlewares.Protected(http.HandlerFunc(userController.UpdateAvatar)))
}
