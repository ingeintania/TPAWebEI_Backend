package model

import "github.com/ingeintania/database"

type Promo struct {
	Promo_id			int		`gorm:"primary_key"`
	Promo_image_path 	string	`gorm:"type:varchar(1000)"`
	Promo_title			string	`gorm:"type:varchar(1000)"`
	Promo_content		string	`gorm:"type:varchar(1000)"`
	Promo_category		string	`gorm:"type:varchar(1000)"`
	Promo_periode		string	`gorm:"type:varchar(1000)"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&Promo{})
}

func GetAllPromo()([]Promo, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allPromo []Promo

	db.Find(&allPromo)

	defer db.Close()
	return allPromo, nil

}

func CreatePromo(image_path string, title string, content string, category string, periode string)(*Promo, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var promo = Promo{
		Promo_image_path: image_path,
		Promo_title:      title,
		Promo_content:    content,
		Promo_category:   category,
		Promo_periode:periode,
	}
	if db.NewRecord(promo){
		db.Create(&promo)
	}

	return &promo,nil
}