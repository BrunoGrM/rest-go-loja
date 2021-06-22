package main

import (
	"log"
	"net/http"
	"rest-go-loja/controllers"
	"rest-go-loja/database"
	"rest-go-loja/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/products/get", controllers.FindAll).Methods("GET")
	router.HandleFunc("/products/get/{id}", controllers.FindByID).Methods("GET")
	router.HandleFunc("/products/create", controllers.Create).Methods("POST")
	router.HandleFunc("/products/update/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/products/delete/{id}", controllers.Delete).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "loja",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Product{})
}
