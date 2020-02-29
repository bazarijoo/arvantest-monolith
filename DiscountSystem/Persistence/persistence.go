package Persistence

import (
	. "arvantest-monolith/DiscountSystem/Entity"
	"arvantest-monolith/DiscountSystem/Model"
	. "arvantest-monolith/WalletSystem/Entity"
	"arvantest-monolith/config"
	"errors"
	"fmt"
	"github.com/captaincodeman/couponcode"
	_ "github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/t-tiger/gorm-bulk-insert"
	"log"
)

func EnableVoucherCode(code string,amount int) error {
	db, _ := config.GetDB()

	validatedCode, err := couponcode.Validate(code)
	if err!=nil{
		return err
	}

	var voucher VoucherEntity
	if err := db.Where("code = ?", validatedCode).Find(&voucher).Error; err != nil {
		log.Println(err)
		return err
	}

	if err:=db.Model(&voucher).Updates(VoucherEntity{
		Code:voucher.Code,
		IsActive:true,
		CountUsed:0,
		Amount:amount,
	}).Error;err!=nil{
		log.Println(err)
		return err
	}

	var users []UserEntity
	if err:=db.Find(&users).Error; err!=nil{
		log.Println(err)
		return err
	}

	var voucherUsers []interface{}
	voucherUsers = Model.GetVoucherUserEntities(users,validatedCode)

	gormbulk.BulkInsert(db, voucherUsers, 200)

	return nil
}

func SubmitVoucherCode(voucherCode, phoneNumber string) error {

	db, _ := config.GetDB()

	//if err := db.Exec("PRAGMA serializable = true").Error; err != nil { /// or serializable
	//	return err
	//}
	//Transactions in SQLite are SERIALIZABLE. and changes are not visible  to ther db connections prior commit
	tx:=db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	validatedCode,err:=couponcode.Validate(voucherCode)
	if err!=nil{
		return err
	}

	var voucher VoucherEntity
	if err := db.Where("code = ?", validatedCode).Find(&voucher).Error; err != nil {
		return err
	}
	if voucher.IsActive==false{
		return errors.New("Voucher Code not activated")
	}

	if voucher.CountUsed >= 1000{
		return errors.New("gift card limit reached, not valid any more")
	}

	var voucherUser VoucherUserEntity
	if err := db.Where("phone_number = ? AND voucher_code = ? ", phoneNumber,voucherCode).Find(&voucherUser).Error; err != nil {
		fmt.Println("voucher_user not found")
		log.Fatal(err)
		return err
	}

	if voucherUser.IsUsed==true{
		return errors.New("gift code is used before")
	}

	newIsActive :=true
	if voucher.CountUsed+1>=1000{
		newIsActive=false
	}

	if err:=db.Model(&voucher).Updates(VoucherEntity{
		Code:voucher.Code,
		IsActive:newIsActive,
		CountUsed:voucher.CountUsed+1,
		Amount:voucher.Amount,
	}).Error;err!=nil{
		log.Fatal(err)
		return err
	}
	if err:=db.Model(&voucherUser).Updates(VoucherUserEntity{
		VoucherCode:voucher.Code,
		PhoneNumber:voucherUser.PhoneNumber,
		IsUsed:true,
	}).Error;err!=nil{
		log.Fatal(err)
		return err
	}

	err =Model.UpdateBalanceAPICall(voucherUser.PhoneNumber,voucher.Amount)
	if err!=nil{
		log.Println(err)
		return err
	}

	return tx.Commit().Error
}