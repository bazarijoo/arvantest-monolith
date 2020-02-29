package Entity

type VoucherEntity struct {
	Code      		string       				`gorm:"primary_key;unique;not null"`
	IsActive  		bool         				`gorm:"default:false"`
	CountUsed 		int         				`gorm:"default:0"`
	Amount			int						    `gorm:"default:0"`
	VoucherUsed 	[]VoucherUserEntity         `gorm:"foreignkey:phone_number;association_foreignkey:voucher_code"`

}
func (VoucherEntity) TableName() string {
	return "voucher_entity"
}