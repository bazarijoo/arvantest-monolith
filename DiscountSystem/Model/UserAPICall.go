package Model

import (
	"arvantest-monolith/DiscountSystem/Model/RequestBody"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func UpdateBalanceAPICall(phoneNumber string,amount int)error{

	url := "http://localhost:8080/update-balance"
	request:= RequestBody.UpdateUserBalanceRequestBody{PhoneNumber:phoneNumber,Amount:amount}

	jsonReq,_:=json.Marshal(request)

	_, err := http.Post(url, "application/json",  bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		panic(err)
		return err
	}

	return nil
}