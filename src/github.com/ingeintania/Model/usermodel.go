package Model

import "github.com/ingeintania/database"

type User struct {
	UserName string
	Password string
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&User{})
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
