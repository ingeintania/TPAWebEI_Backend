package model

import (
	"fmt"
	"github.com/ingeintania/database"
)

type User struct {
	User_id			int	`gorm:"primary_key"`
	User_email		string	`gorm:"type:varchar(1000);not null"`
	User_phone		string	`gorm:"type:varchar(1000);not null"`
	User_gender		string	`gorm:"type:varchar(1000);not null"`
	User_firstname 	string	`gorm:"type:varchar(1000);not null"`
	User_lastname 	string	`gorm:"type:varchar(1000);not null"`
	User_city		string	`gorm:"type:varchar(1000);not null"`
	User_address	string	`gorm:"type:varchar(1000);not null"`
	User_post_code	string	`gorm:"type:varchar(1000);not null"`
	User_password	string	`gorm:"type:varchar(1000);not null"`
	User_language	string	`gorm:"type:varchar(1000)"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&User{})
}

func GetLogin(email string, password string)([]User, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var user []User

	if db.Where("user_email = ?",email).Where("user_password = ?", password).Find(&user).RecordNotFound(){
		return nil, nil
	}

	fmt.Println(user)

	return user,err
}

func GetUserByEmail(email string)([]User, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var user []User

	if db.Where("user_email = ?",email).Find(&user).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(user)

	return user,err
}

func GetUserByPhone(phone string)([]User, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var user []User

	if db.Where("user_phone = ?",phone).Find(&user).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(user)

	return user,err
}

func GetUserByPhoneOrEmail(arg string) ([]User, error){
	db, err:= database.Connect()

	if err != nil{
		return nil,err
	}
	defer db.Close()

	var user []User

	if db.Where("user_email = ?",arg).Or("user_phone = ?", arg).Find(&user).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(user)

	return user,nil
}

func GetAllUser()([]User, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allUser []User

	db.Find(&allUser)

	defer db.Close()
	return allUser, nil

}

func CreateUser(email string, phone string, gender string,
	firstname string, lastname string, city string,
	address string, post_code string,
	password string, language string)(*User, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var user = User{
		User_firstname: firstname,
		User_lastname: lastname,
		User_password: password,
		User_email:    email,
		User_phone:    phone,
		User_address: address,
		User_city:city,
		User_post_code:post_code,
		User_gender:gender,
		User_language:language,
	}
	if db.NewRecord(user){
		db.Create(&user)
	}

	return &user,nil
}

func UpdateUser(id int, email string, phone string, gender string,
	firstname string, lastname string, city string,
	address string, post_code string,
	password string, language string)(*User, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var user = User{
		User_id:id,
	}

	db.Model(&user).Where("user_id=?", id).
		Update(map[string]interface{}{"user_email": email, "user_phone":phone,
			"user_gender":gender, "user_firstname":firstname,
			"user_lastname":lastname, "user_city":city,
			"user_address":address, "user_post_code":post_code,
			"user_password":password, "user_language":language,})
	fmt.Print("User has been updated!")

	return &user,nil
}