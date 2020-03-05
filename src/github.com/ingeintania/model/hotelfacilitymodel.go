package model

import (
	"github.com/ingeintania/database"
)

type Hotelfacility struct {
	Hotelfacility_id   int    `gorm:"primary_key"`
	Hotelid int `gorm:"type:integer;not null"`
	Hotelfacility_name string `gorm:"type:varchar(100);not null"`
}

func init() {
	db, error := database.Connect()

	if error != nil {
		panic("No Conection")
	}
	db.DropTableIfExists()
	db.AutoMigrate(&Hotelfacility{})
}

func GetAllHotelfacility() ([]Hotelfacility, error) {
	db, error := database.Connect()

	if error != nil {
		return nil, error
	}
	var allHotelfacility []Hotelfacility

	db.Find(&allHotelfacility)

	defer db.Close()
	return allHotelfacility, nil
}

func CreateHotelfacility(name string, hotelId int) (*Hotelfacility, error) {
	db, error := database.Connect()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	var hotelfacility = Hotelfacility{
		Hotelfacility_name: name,
		Hotelid:hotelId,
	}
	if db.NewRecord(hotelfacility) {
		db.Create(&hotelfacility)
	}

	return &hotelfacility, nil
}
