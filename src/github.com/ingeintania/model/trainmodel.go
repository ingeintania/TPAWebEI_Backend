package model

import (
	"fmt"
	"github.com/ingeintania/database"
)

type Train struct {
	Train_id				int			`gorm:"primary_key"`
	Train_name	 			string	`gorm:"type:varchar(100)"`
	Train_class	 			string	`gorm:"type:varchar(100)"`
	Train_type				string	`gorm:"type:varchar(100)"`
	Train_depart_time		string	`gorm:"type:varchar(100)"`
	Train_depart_location	string	`gorm:"type:varchar(100)"`
	Train_arrive_time		string	`gorm:"type:varchar(100)"`
	Train_arrive_location	string	`gorm:"type:varchar(100)"`
	Train_duration			int		`gorm:"type:int"`
	Train_detail_price		int		`gorm:"type:varchar(100)"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.AutoMigrate(&Train{})
}

func GetAllTrain()([]Train, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	var allTrain []Train

	db.Find(&allTrain)

	defer db.Close()
	return allTrain, nil

}

func CreateTrain(name string, classes string, types string, depart_time string,
	depart_location string, arrive_time string, arrive_location string,
	durations int, detail_price int)(*Train, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var train = Train{
		Train_name:name,
		Train_class:classes,
		Train_type:types,
		Train_depart_time:depart_time,
		Train_depart_location:depart_location,
		Train_arrive_time:arrive_time,
		Train_arrive_location:arrive_location,
		Train_duration:durations,
		Train_detail_price:detail_price,
	}

	if name == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if classes == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if types == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_time == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_location == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_time == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_location == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if durations == 0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if detail_price == 0{
		fmt.Print(name)
		panic("DB SALAH")
	}


	if db.NewRecord(train){
		db.Create(&train)
		fmt.Print("Train has been added!")
	}

	return &train,nil
}

func UpdateTrain(id int, name string, classes string, types string, depart_time string,
	depart_location string, arrive_time string, arrive_location string,
	durations int, detail_price int)(*Train, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var train = Train{
		Train_id:id,
	}
	if name == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if classes == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if types == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_time == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if depart_location == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_time == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if arrive_location == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if durations == 0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if detail_price == 0{
		fmt.Print(name)
		panic("DB SALAH")
	}
	db.Model(&train).Where("train_id=?", id).
		Update(map[string]interface{}{"train_name": name, "train_class":classes,
			"train_type":types, "train_depart_time":depart_time,
			"train_depart_location":depart_location, "train_arrive_time":arrive_time,
			"train_arrive_location":arrive_location, "train_duration":durations,
			"train_detail_price":detail_price,})
	fmt.Print("Train has been updated!")

	return &train,nil
}

func DeleteTrain(id int)(*Train, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var train = Train{
		Train_id:id,
	}
	db.Where("train_id=?", id).Find(&train)
	db.Delete(train)

	return &train,nil
}

