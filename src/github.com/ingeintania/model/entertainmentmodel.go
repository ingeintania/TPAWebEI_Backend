package model

import (
	"fmt"
	"github.com/ingeintania/database"
	"time"
)

type Entertainment struct {
	Entertainment_id			int		`gorm:"primary_key"`
	Entertainment_name 			string 	`gorm:"type:varchar(1000)"`
	Entertainment_price 		int		`gorm:"type:int"`
	Entertainment_image_path1	string 	`gorm:"type:varchar(1000)"`
	Entertainment_image_path2	string 	`gorm:"type:varchar(1000)"`
	Entertainment_image_path3	string 	`gorm:"type:varchar(1000)"`
	Entertainment_image_path4	string 	`gorm:"type:varchar(1000)"`
	Entertainment_image_path5	string 	`gorm:"type:varchar(1000)"`
	Location Location			`gorm:"foreign_key:Location_id"`
	Location_id			int
	Entertainment_address 		string 	`gorm:"type:varchar(1000)"`
	Entertainment_latitude 		string 	`gorm:"type:varchar(1000)"`
	Entertainment_longitude 	string 	`gorm:"type:varchar(1000)"`
	Entertainment_type 			string 	`gorm:"type:varchar(1000)"`
	Entertainment_start_date	string 	`gorm:"type:varchar(1000)"`
	Entertainment_end_date		string 	`gorm:"type:varchar(1000)"`
	Entertainment_description	string 	`gorm:"type:varchar(1000)"`
	Entertainment_t_c			string 	`gorm:"type:varchar(1000)"`
}

func init(){
	db, error := database.Connect()

	if error != nil{
		panic("No Conection")
	}
	db.DropTableIfExists()
	db.AutoMigrate(&Entertainment{})
	//.AddForeignKey("location_id", "locations(location_id)", "CASCADE", "CASCADE").
	//AddForeignKey("entertainmenttype_id", "entertainmenttypes(entertainmenttype_id)", "CASCADE", "CASCADE")
}

func GetAllEntertainment()([]Entertainment, error){
	db, error := database.Connect()

	if error != nil{
		return nil, error
	}
	//var allEntertainment []Entertainment
	//db.Find(&allEntertainment)

	defer db.Close()

	var entertainment []Entertainment
	db.Find(&entertainment)

	for i := 0; i < len(entertainment); i++ {
		db.Model(entertainment[i]).Related(&entertainment[i].Location)
	}

	fmt.Println(entertainment)

	return entertainment, nil
}

func GetEntertainmentByName(name string)([]Entertainment, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var entertainment []Entertainment

	if db.Where("entertainment_name=?",name).Find(&entertainment).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(entertainment)

	return entertainment,err
}

func GetEntertainmentByPrice(price int)([]Entertainment, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var entertainment []Entertainment

	if db.Where("entertainment_price>=?",price).Find(&entertainment).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(entertainment)

	return entertainment,err
}

func GetEntertainmentByType(types string)([]Entertainment, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var entertainment []Entertainment

	if db.Where("entertainment_type=?",types).Find(&entertainment).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(entertainment)

	return entertainment,err
}

func GetEntertainmentByStartDate(startDate time.Time)([]Entertainment, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var entertainment []Entertainment

	if db.Where("entertainment_start_date=?",startDate).Find(&entertainment).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(entertainment)

	return entertainment,err
}

func GetEntertainmentByEndDate(endDate time.Time)([]Entertainment, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var entertainment []Entertainment

	if db.Where("entertainment_end_date=?",endDate).Find(&entertainment).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(entertainment)

	return entertainment,err
}

func CreateEntertainment(name string, price int, image_path1 string, image_path2 string, image_path3 string,image_path4 string, image_path5 string, locationId int, address string, latitude string, longitude string, types string, startDate string, endDate string, description string, t_c string)(*Entertainment, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var entertainment = Entertainment{
		Entertainment_name:name,
		Entertainment_price:price,
		Entertainment_image_path1:image_path1,
		Entertainment_image_path2:image_path2,
		Entertainment_image_path3:image_path3,
		Entertainment_image_path4:image_path4,
		Entertainment_image_path5:image_path5,
		Entertainment_address:address,
		Location_id:locationId,
		Entertainment_latitude:latitude,
		Entertainment_longitude:longitude,
		Entertainment_type:types,
		Entertainment_start_date:startDate,
		Entertainment_end_date:endDate,
		Entertainment_description:description,
		Entertainment_t_c:t_c,
	}

	if name == "" || len(name) > 10{
		fmt.Print(name)
		panic("DB SALAH")
	}
	if price == 0{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if image_path1 == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if image_path2 == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if image_path3 == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if image_path4 == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if image_path5 == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if locationId < 0 || locationId > 6{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if latitude == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if longitude == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if types == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if startDate == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if endDate == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if description == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if t_c == ""{
		fmt.Print(name)
		panic("DB SALAH")
	}

	if db.NewRecord(entertainment){
		db.Create(&entertainment)
		panic("New Entertainment has been added!")
	}



	return &entertainment,nil
}

func UpdateEntertainment(id int, name string, price int, image_path1 string, image_path2 string, image_path3 string,image_path4 string, image_path5 string, locationId int, address string, latitude string, longitude string, types string, startDate string, endDate string, description string, t_c string)(*Entertainment, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var entertainment = Entertainment{
		Entertainment_id:id,
	}

	db.Model(&entertainment).Where("Entertainment_id=?", id).
		Update(map[string]interface{}{"entertainment_name": name, "entertainment_price":price,
			"entertainment_image_path1":image_path1, "entertainment_image_path2":image_path2,
			"entertainment_image_path3":image_path3, "entertainment_image_path4":image_path4,
			"entertainment_image_path5":image_path5, "location_id":locationId, "entertainment_address":address,
			"entertainment_latitude":latitude, "entertainment_longitude":longitude,
			"entertainment_type":types, "entertainment_start_date":startDate,
			"entertainment_end_date":endDate, "entertainment_description":description,
			"entertainment_t_c":t_c,})
	panic("Entertainment has been updated!")

	return &entertainment,nil
}

func DeleteEntertainment(id int)(*Entertainment, error){
	db, error := database.Connect()

	if error !=nil{
		return nil,error
	}
	defer db.Close()
	var entertainment = Entertainment{
		Entertainment_id:id,
	}

	db.Where("entertainment_id=?", id).Find(&entertainment)
	db.Delete(entertainment)

	return &entertainment,nil
}