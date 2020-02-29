package Persistence

import (
	. "arvantest-monolith/WalletSystem/Entity"
	"arvantest-monolith/config"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetBalance(phoneNumber string) (int,error){
	db, _ := config.GetDB()

	var user UserEntity
	if err := db.Where("phone_number = ?", phoneNumber).Find(&user).Error; err != nil {
		log.Println(err)
		return -1,err
	}

	return user.Balance,nil
}

func UpdateBalance(phoneNumber string, amount int) error {
	db, _ := config.GetDB()

	var user UserEntity
	if err := db.Where("phone_number = ?", phoneNumber).Find(&user).Error; err != nil {
		log.Println(err)
		return err
	}

	if err := db.Exec("PRAGMA read_committed = true").Error; err != nil {
		return err
	}

	tx:=db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err:=db.Model(&user).Update("balance",user.Balance + amount).Error;err!=nil{
		log.Println(err)
			return err
	}

	return tx.Commit().Error
}