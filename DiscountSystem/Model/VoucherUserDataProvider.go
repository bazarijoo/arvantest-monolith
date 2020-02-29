package Model

import (
	. "arvantest-monolith/DiscountSystem/Entity"
	. "arvantest-monolith/WalletSystem/Entity"
)

func GetVoucherUserEntities(users []UserEntity,voucherCode string)[]interface{} {

	var voucherUsers []interface{}
	for _, user := range users {
		voucherUsers = append(voucherUsers, VoucherUserEntity{
			VoucherCode: voucherCode,
			PhoneNumber: user.PhoneNumber,
			IsUsed:      false,
		})
	}
	return voucherUsers
}
