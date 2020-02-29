package main

import (
	. "arvantest-monolith/WalletSystem/Controller"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func handleRequests(){
	router := mux.NewRouter()
	router.HandleFunc("/get-balance", GetBalance).Methods("POST")
	router.HandleFunc("/update-balance", UpdateBalance).Methods("POST")

	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {

	handleRequests()

}
