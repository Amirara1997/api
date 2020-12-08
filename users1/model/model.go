package model

import (

	"users1/database"
	"users1/security"
)


type User struct {
	ID          uint         `gorm:"primary_key" json:"id"`
	Name        string       `json:"name"`
	Passwd         string       `json:"passwd"`
}

func (User) TableName() string {
	return "userha"
}

func (u *User) BeforeSave() error  {
	hashedPasswd , err := security.Hash(u.Passwd)
	if err != nil{
		return err
	}
	u.Passwd = string(hashedPasswd)
	return nil
}

func GetUserByUsername (name string) (User,error) {
	var (
		user User
		err error
	)
	db,_ := database.NewDB()
	if err = db.Find(&user,"name = ?",name ).Error;err != nil {
		return user,err
	}
	db.Commit()
	return user,err
}