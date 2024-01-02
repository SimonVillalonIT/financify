package main

import (
	"financify/pkg/config"
	"financify/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	config.Connect()
	db := config.GetDB()
    config.MigrateDB()
	router := mux.NewRouter()
	routes.UserRoutes(db, router)
	routes.TransactionsRoutes(db, router)
	http.Handle("/", router)
	PORT := os.Getenv("PORT")
	log.Println("Listening on port: " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
