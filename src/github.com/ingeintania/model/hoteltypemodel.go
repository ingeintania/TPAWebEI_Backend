package model

import "github.com/ingeintania/database"

type Hoteltype struct {
	Hoteltype_id		int	`gorm:"primary_key"`
	Hoteltype_name	string	`gorm:"type:varchar(100);not null"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.DropTableIfExists()
	db.AutoMigrate(&Hoteltype{})
}

func GetAllHoteltype()([]Hoteltype, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allHoteltype []Hoteltype

	db.Find(&allHoteltype)

	defer db.Close()
	return allHoteltype, nil
}

func CreateHoteltype(name string)(*Hoteltype, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var hoteltype = Hoteltype{
		Hoteltype_name:name,
	}
	if db.NewRecord(hoteltype){
		db.Create(&hoteltype)
	}

	return &hoteltype,nil
}
