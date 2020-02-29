package RequestBody

type UpdateUserBalanceRequestBody struct {
	PhoneNumber  string   `json:"phone_number"`
	Amount       int      `json:"amount"`
}