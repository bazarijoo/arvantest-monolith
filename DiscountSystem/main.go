package main

import (
	. "arvantest-monolith/DiscountSystem/Controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/enable-voucher-code", EnableVoucherCode).Methods("POST")
	router.HandleFunc("/submit-voucher-code", SubmitVoucherCode).Methods("POST")
	router.HandleFunc("/get-voucher-code-reports", GetVoucherCodeStatus).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {

	handleRequests()
}
