package model

import (
	"fmt"
	"github.com/ingeintania/database"
)

type Aeroplane struct {
	Aeroplane_id				int			`gorm:"primary_key"`
	Aeroplane_name 				string	`gorm:"type:varchar(1000)"`
	Aeroplane_company			string	`gorm:"type:varchar(1000)"`
	Aeroplane_type				string	`gorm:"type:varchar(1000)"`
	Aeroplane_image_path		string	`gorm:"type:varchar(1000)"`
	Aeroplane_depart_date		string	`gorm:"type:varchar(1000)"`
	Aeroplane_depart_time		string	`gorm:"type:varchar(1000)"`
	Aeroplane_arrive_date		string	`gorm:"type:varchar(1000)"`
	Aeroplane_arrive_time		string	`gorm:"type:varchar(1000)"`
	Aeroplane_duration			int		`gorm:"type:int"`
	Aeroplane_price				int		`gorm:"type:int"`
	Aeroplane_price_raw			int		`gorm:"type:int"`
	Aeroplane_price_tax			int		`gorm:"type:int"`
	Aeroplane_transit			string	`gorm:"type:varchar(1000)"`
	Aeroplane_transit_duration	int		`gorm:"type:int"`
	Aeroplane_facilities		string	`gorm:"type:varchar(1000)"`
	Aeroplane_depart_location	string	`gorm:"type:varchar(1000)"`
	Aeroplane_arrive_location	string	`gorm:"type:varchar(1000)"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&Aeroplane{})
}

func GetAllAeroplane()([]Aeroplane, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allAeroplane []Aeroplane

	db.Find(&allAeroplane)

	defer db.Close()
	return allAeroplane, nil

}

func CreateAeroplane(name string, company string, types string, images_path string, depart_date string,
	depart_time string, arrive_date string, arrive_time string, duration int,
	price int, price_raw int, price_tax int, transit string,
	transit_duration int, facilities string,
	depart_location string, arrive_location string)(*Aeroplane, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var aeroplane = Aeroplane{
		Aeroplane_name:name,
		Aeroplane_company:company,
		Aeroplane_type:types,
		Aeroplane_image_path:images_path,
		Aeroplane_depart_date:depart_date,
		Aeroplane_depart_time:depart_time,
		Aeroplane_arrive_date:arrive_date,
		Aeroplane_arrive_time:arrive_time,
		Aeroplane_duration:duration,
		Aeroplane_price:price,
		Aeroplane_price_raw:price_raw,
		Aeroplane_price_tax:price_tax,
		Aeroplane_transit:transit,
		Aeroplane_transit_duration:transit_duration,
		Aeroplane_facilities:facilities,
		Aeroplane_depart_location:depart_location,
		Aeroplane_arrive_location:arrive_location,
	}
	if name==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if company==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if types==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if images_path==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_date==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_time==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_location==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_date==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_time==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_location==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if duration==0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if price==0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if price_raw==0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if price_tax==0{
		fmt.Print(name)
		panic("DB SALAH")
	}



	if db.NewRecord(aeroplane){
		db.Create(&aeroplane)
		fmt.Print("New Flight has added!")
	}

	return &aeroplane,nil
}

func UpdateAeroplane(id int, name string, company string, types string, images_path string, depart_date string,
	depart_time string, arrive_date string, arrive_time string, duration int,
	price int, price_raw int, price_tax int, transit string,
	transit_duration int, facilities string,
	depart_location string, arrive_location string)(*Aeroplane, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var aeroplane = Aeroplane{
		Aeroplane_id:id,
	}
	if types==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if images_path==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_date==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_time==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_location==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_date==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_time==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_location==""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if duration==0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if price==0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if price_raw==0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if price_tax==0{
		fmt.Print(name)
		panic("DB SALAH")
	}

	db.Model(&aeroplane).Where("aeroplane_id=?", id).
		Update(map[string]interface{}{"aeroplane_name": name, "aeroplane_company":company,
			"aeroplane_type":types, "aeroplane_images_path":images_path,
			"aeroplane_depart_date":depart_date, "aeroplane_depart_time":depart_time,
			"aeroplane_arrive_date":arrive_date, "aeroplane_arrive_time":arrive_time,
			"aeroplane_duration":duration, "aeroplane_price":price,
			"aeroplane_price_raw":price_raw, "aeroplane_price_tax":price_tax,
			"aeroplane_transit":transit,
			"aeroplane_transit_duration":transit_duration, "aeroplane_facilities":facilities,
			"aeroplane_depart_location":depart_location, "aeroplane_arrive_location":arrive_location,})

	fmt.Print("Flight has been updated!")
	return &aeroplane,nil
}

func DeleteAeroplane(id int)(*Aeroplane, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var aeroplane = Aeroplane{
		Aeroplane_id:id,
	}

	db.Where("aeroplane_id=?", id).Find(&aeroplane)
	db.Delete(aeroplane)

	return &aeroplane,nil
}