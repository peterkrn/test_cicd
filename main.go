package main

import (
	"fmt"
	"latihanbackend/controllers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected on port :8888")
	log.Println("Connected on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
