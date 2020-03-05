package model

import "github.com/ingeintania/database"

type Car struct {
	Car_id					int			`gorm:"primary_key"`
	Car_image_path			string	`gorm:"type:varchar(1000)"`
	Car_model	 			string	`gorm:"type:varchar(1000)"`
	Car_price	 			int		`gorm:"type:int"`
	Car_passenger			int		`gorm:"type:int"`
	Car_luggage				int		`gorm:"type:int"`
	Car_vendor				string	`gorm:"type:varchar(1000)"`
	Car_location			string	`gorm:"type:varchar(1000)"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&Car{})
}

func GetAllCar()([]Car, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allCar []Car

	db.Find(&allCar)

	defer db.Close()
	return allCar, nil

}

func CreateCar(image_path string, model string,
	price int, passenger int, luggage int,
	vendor string, location string)(*Car, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var train = Car{
		Car_image_path:image_path,
		Car_model:model,
		Car_price:price,
		Car_passenger:passenger,
		Car_luggage:luggage,
		Car_vendor:vendor,
		Car_location:location,
	}
	if db.NewRecord(train){
		db.Create(&train)
	}

	return &train,nil
}