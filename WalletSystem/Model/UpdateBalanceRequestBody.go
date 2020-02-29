package Model

type UpdateBalanceRequestBody struct {
	PhoneNumber  string  `json:"phone_number"`
	Amount       int     `json:"amount"`
}