package Controller

import (
	. "arvantest-monolith/WalletSystem/Model"
	"arvantest-monolith/WalletSystem/Persistence"
	"encoding/json"
	"net/http"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {

	var data GetBalanceRequestBody
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	balance, err := Persistence.GetBalance(data.PhoneNumber)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = make(map[string]interface{})
	if err != nil {
			response["error"]=err.Error()
			_=json.NewEncoder(w).Encode(response)
		return
	}

	response["balance"] = balance
	_=json.NewEncoder(w).Encode(response)

}

func UpdateBalance(w http.ResponseWriter, r *http.Request) {
	var requestBody UpdateBalanceRequestBody

	var response = make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = Persistence.UpdateBalance(requestBody.PhoneNumber,requestBody.Amount)
	if err != nil {
		panic(err)
	}

	response["result"] = "Ok"
	_=json.NewEncoder(w).Encode(response)

}

