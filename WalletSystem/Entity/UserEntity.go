package Entity

type UserEntity struct {
	PhoneNumber     string `gorm:"primary_key;unique;not null"`
	Balance         int  `gorm:"default:0"`
}

func (UserEntity) TableName() string {
	return "user_entity"
}

