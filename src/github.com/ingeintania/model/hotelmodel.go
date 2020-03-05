package model

import (
	"fmt"
	"github.com/ingeintania/database"
	"strconv"
)

type Hotel struct {
	Hotel_id           int      `gorm:"primary_key"`
	Hotel_name         string   `gorm:"type:varchar(1000)"`
	Hotel_price        int      `gorm:"type:int"`
	Hotel_image_path   string   `gorm:"type:varchar(1000)"`
	HotelFacility   []Hotelfacility `gorm:"foreign_key:hotelid"`
	Location           Location `gorm:"foreign_key:Location_id"`
	Location_id        int
	Hoteltype          Hoteltype `gorm:"foreign_key:Hoteltype_id"`
	Hoteltype_id       int
	Hotel_address      string `gorm:"type:varchar(1000)"`
	Hotel_latitude      string `gorm:"type:varchar(1000)"`
	Hotel_longitude      string `gorm:"type:varchar(1000)"`
	Hotel_star         int    `gorm:"type:int"`
	Hotel_rating_count int    `gorm:"type:int"`
	Hotel_availability string `gorm:"type:varchar(1000)"`
	Hotel_facilities string `gorm:"type:varchar(1000)"`
}

func init() {
	db, error := database.Connect()

	if error != nil {
		panic("No Conection")
	}
	db.DropTableIfExists()
	db.AutoMigrate(&Hotel{})
	//.AddForeignKey("location_id", "locations(location_id)", "CASCADE", "CASCADE").
	//AddForeignKey("hoteltype_id", "hoteltypes(hoteltype_id)", "CASCADE", "CASCADE")
}

func GetAllHotel() ([]Hotel, error) {
	db, error := database.Connect()

	if error != nil {
		return nil, error
	}
	var allHotel []Hotel

	db.Find(&allHotel)

	for i, _ := range allHotel  {
		db.Model(allHotel[i]).Related(&allHotel[i].HotelFacility, "hotelid")
		fmt.Println(allHotel[i])
	}

	defer db.Close()
	return allHotel, nil
}

func GetHotelByName(name string) ([]Hotel, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotel []Hotel

	if db.Where("hotel_name=?", name).Find(&hotel).RecordNotFound() {
		return nil, nil
	}

	fmt.Println(hotel)

	return hotel, err
}

func GetHotelByPrice(price int) ([]Hotel, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotel []Hotel

	if db.Where("hotel_price>=?", price).Find(&hotel).RecordNotFound() {
		return nil, nil
	}

	fmt.Println(hotel)

	return hotel, err
}

func GetHotelByFacility(facility int) ([]Hotel, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotel []Hotel

	if db.Joins("JOIN hotel_detail ON hotel_detail.hotel_id=hotel.hotel_id").Where("hotelfacility_id=?", facility).Find(&hotel).RecordNotFound() {
		return nil, nil
	}

	fmt.Println(hotel)

	return hotel, err
}

func GetHotelByArea(latitude string, longitude string) ([]Hotel, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotel []Hotel

	float_latitude, err := strconv.ParseFloat(latitude, 64)
	float_longitude, err := strconv.ParseFloat(longitude, 64)

	db.Where("(3959 * acos( cos(radians(?)) * cos(radians("+
		"CAST(hotel_latitude AS DECIMAL(10,6)))) * cos(radians(CAST(hotel_longitude AS DECIMAL(10,6))) - radians(?)) + sin(radians("+
		"?))* sin(radians(CAST(hotel_latitude AS DECIMAL(10,6))))))"+
		"< 25", float_latitude, float_longitude, float_latitude).
		Limit(10).
		Order("(3959 * acos( cos(radians(" + latitude + ")) * cos(" +
			"radians(" +
			"CAST(hotel_latitude AS DECIMAL(10,6)))) * cos(radians(CAST(hotel_longitude AS DECIMAL(10,6))) - radians(" +
			"" + longitude + ")) + sin(radians(" +
			"" + latitude + "))* sin(radians(CAST(hotel_latitude AS DECIMAL(10,6))))))").
		Find(&hotel)

	fmt.Println(hotel)

	return hotel, err
}

func CreateHotel(name string, price int, image_path string, locationId int, hotelTypeId int, address string,
	latitude string, longitude string,
	star int, rating_count int, availability string, facilities string) (*Hotel, error) {



	db, error := database.Connect()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	var hotel = Hotel{

		Hotel_name:         name,
		Hotel_price:        price,
		Hotel_image_path:   image_path,
		Location_id:        locationId,
		Hoteltype_id:       hotelTypeId,
		Hotel_address:      address,
		Hotel_latitude:latitude,
		Hotel_longitude:longitude,
		Hotel_star:         star,
		Hotel_rating_count: rating_count,
		Hotel_availability: availability,
		Hotel_facilities:facilities,
	}
	if name == "" || len(name) < 10{
		fmt.Print(name)
		panic("DB NAME SALAH")
	}

	if price == 0{
		fmt.Print(price)
		panic("DB PRICE SALAH")
	}

	if image_path == ""{
		fmt.Print(image_path)
		panic("DB IMAGE SALAH")
	}

	if locationId == 0{
		fmt.Print(locationId)
		panic("DB LOC SALAH")
	}

	if hotelTypeId == 0 {
		fmt.Print(hotelTypeId)
		panic("DB TYPE SALAH")
	}

	if address == ""{
		fmt.Print(address)
		panic("DB ADDRESS SALAH")
	}

	if latitude == ""{
		fmt.Print(latitude)
		panic("DB LAT SALAH")
	}

	if longitude == ""{
		fmt.Print(longitude)
		panic("DB LONG SALAH")
	}

	if star == 0{
		fmt.Print(star)
		panic("DB STAR SALAH")
	}

	if rating_count == 0{
		fmt.Print(rating_count)
		panic("DB RATING SALAH")
	}

	if availability == ""{
		fmt.Print(availability)
		panic("DB AVAIL SALAH")
	}

	if facilities == ""{
		fmt.Print(facilities)
		panic("DB PRINT SALAH")
	}



	if db.NewRecord(hotel) {
		db.Create(&hotel)
		fmt.Print("New Hotel has been added!")
	}



	return &hotel, nil
}

func UpdateHotel(id int, name string, price int, image_path string, locationId int, hotelTypeId int, address string,
	latitude string, longitude string,
	star int,
	rating_count int, availability string, facilities string) (*Hotel, error) {
	db, error := database.Connect()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	var hotel = Hotel{
		Hotel_id:id,

	}
	if name == "" || len(name) < 10{
		fmt.Print(name)
		panic("DB NAME SALAH")
	}

	if price == 0{
		fmt.Print(price)
		panic("DB PRICE SALAH")
	}

	if image_path == ""{
		fmt.Print(image_path)
		panic("DB IMAGE SALAH")
	}

	if locationId == 0{
		fmt.Print(locationId)
		panic("DB LOCATION SALAH")
	}

	if hotelTypeId == 0 {
		fmt.Print(hotelTypeId)
		panic("DB TYPE SALAH")
	}

	if address == ""{
		fmt.Print(address)
		panic("DB ADDRESS SALAH")
	}

	if latitude == ""{
		fmt.Print(latitude)
		panic("DB LAT SALAH")
	}

	if longitude == ""{
		fmt.Print(longitude)
		panic("DB LONG SALAH")
	}

	if star == 0{
		fmt.Print(star)
		panic("DB STAR SALAH")
	}

	if rating_count == 0{
		fmt.Print(rating_count)
		panic("DB RATING SALAH")
	}

	if availability == ""{
		fmt.Print(availability)
		panic("DB AVAILABILITY SALAH")
	}

	if facilities == ""{
		fmt.Print(facilities)
		panic("DB FACILITY SALAH")
	}

	db.Model(&hotel).Where("hotel_id=?", id).
		Update(map[string]interface{}{"hotel_name": name, "hotel_price":price,
			"hotel_image_path":image_path, "location_id":locationId,
			"hoteltype_id":hotelTypeId, "hotel_address":address,
			"hotel_latitude":latitude, "hotel_longitude":longitude,
			"hotel_star":address, "hotel_rating_count":star,
			"hotel_availability":availability, "hotel_facilities":facilities,})

	fmt.Print("Hotel has been updated!")

	return &hotel, nil
}

func DeleteHotel(id int) (*Hotel, error) {
	db, error := database.Connect()

	if error != nil {
		return nil, error
	}
	defer db.Close()
	var hotel = Hotel{
		Hotel_id:id,

	}
	db.Where("hotel_id=?", id).Find(&hotel)
	db.Delete(hotel)

	return &hotel, nil
}