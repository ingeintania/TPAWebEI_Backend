package model

import "github.com/ingeintania/database"

type Location struct {
	Location_id		int	`gorm:"primary_key"`
	Location_name	string	`gorm:"type:varchar(100);not null"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.DropTableIfExists()
	db.AutoMigrate(&Location{})
}

func GetAllLocation()([]Location, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allLocation []Location

	db.Find(&allLocation)

	defer db.Close()
	return allLocation, nil
}

func CreateLocation(name string)(*Location, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var location = Location{
		Location_name:name,
	}
	if db.NewRecord(location){
		db.Create(&location)
	}

	return &location,nil
}
