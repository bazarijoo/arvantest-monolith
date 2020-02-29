package RequestBody

type EnableVoucherCodeRequestBody struct {
	VoucherCode string  `json:"voucher_code"`
	Amount       int   `json:"amount"`
}
