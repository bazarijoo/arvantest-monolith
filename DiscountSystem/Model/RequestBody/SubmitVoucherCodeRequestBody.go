package RequestBody

type SubmitVoucherCodeRequestBody struct {
	PhoneNumber string `json:"phone_number"`
	VoucherCode string `json:"voucher_code"`
}
