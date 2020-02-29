package Controller

import (
	. "arvantest-monolith/DiscountSystem/Model/RequestBody"
	"arvantest-monolith/DiscountSystem/Persistence"
	"encoding/json"
	"github.com/captaincodeman/couponcode"
	"net/http"
)

func EnableVoucherCode(w http.ResponseWriter, r *http.Request) {

	var requestBody EnableVoucherCodeRequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = make(map[string]interface{})

	validatedCode,err := couponcode.Validate(requestBody.VoucherCode)

	if err != nil {
		response["error"]=err.Error()
		_=json.NewEncoder(w).Encode(response)
		return
	}
	err=Persistence.EnableVoucherCode(validatedCode,requestBody.Amount)
	if err != nil {
		response["error"]=err.Error()
		_=json.NewEncoder(w).Encode(response)
		return
	}

	response["result"]="OK"
	_=json.NewEncoder(w).Encode(response)
}

func SubmitVoucherCode(w http.ResponseWriter, r *http.Request) {

	var body SubmitVoucherCodeRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = make(map[string]interface{})

	err=Persistence.SubmitVoucherCode(body.VoucherCode,body.PhoneNumber)
	if err != nil {
		response["error"]=err.Error()
		_=json.NewEncoder(w).Encode(response)
		return
	}

	response["result"]="OK"
	_=json.NewEncoder(w).Encode(response)

}

func GetVoucherCodeStatus(w http.ResponseWriter, r *http.Request) {



	//var body SubmitVoucherCodeRequestBody
	//
	//err := json.NewDecoder(r.Body).Decode(&body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//
	//var response = make(map[string]interface{})
	//
	//err=Persistence.SubmitVoucherCode(body.VoucherCode,body.PhoneNumber)
	//if err != nil {
	//	response["error"]=err.Error()
	//	_=json.NewEncoder(w).Encode(response)
	//	return
	//}
	//
	//response["result"]="OK"
	//_=json.NewEncoder(w).Encode(response)

}
