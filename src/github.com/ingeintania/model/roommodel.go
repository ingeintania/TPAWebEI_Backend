package model

import (
	"fmt"
	"github.com/ingeintania/database"
)

type Room struct {
	Room_id				int			`gorm:"primary_key"`
	Hotel_id			int		`gorm:"type:int"`
	Room_image_path 	string	`gorm:"type:varchar(1000)"`
	Room_title			string	`gorm:"type:varchar(1000)"`
	Room_cancel_status	string	`gorm:"type:varchar(1000)"`
	Room_information	string	`gorm:"type:varchar(1000)"`
	Room_facilities		string	`gorm:"type:varchar(1000)"`
	Room_availability	int	`gorm:"type:int"`
	Room_price			string	`gorm:"type:varchar(1000)"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&Room{})
}

func GetAllRoom()([]Room, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allRoom []Room

	db.Find(&allRoom)

	defer db.Close()
	return allRoom, nil
}

func GetRoomByHotelId(id int)([]Room, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var room []Room

	if db.Where("hotel_id = ?",id).Find(&room).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(room)

	return room,err
}