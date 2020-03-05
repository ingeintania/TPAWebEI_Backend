package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/ingeintania/model"
	"time"
)

func GetUsers(p graphql.ResolveParams) (i interface{}, e error){
	users, error:= model.GetAllUser()

	return users, error
}

func GetUserByEmail(p graphql.ResolveParams) (i interface{}, e error){
	email:=p.Args["user_email"].(string)
	user,err:= model.GetUserByEmail(email)

	return user,err

}

func GetUserByPhone(p graphql.ResolveParams) (i interface{}, e error){
	phone:=p.Args["user_phone"].(string)
	user,err:= model.GetUserByPhone(phone)

	return user,err

}

func GetUserByPhoneOrEmail(p graphql.ResolveParams) (i interface{},e error){
	arg:= p.Args["arg"].(string)

	user, err:= model.GetUserByPhoneOrEmail(arg)

	return user, err
}

func GetLocations(p graphql.ResolveParams) (i interface{}, e error){
	locations, error:= model.GetAllLocation()

	return locations, error
}

func GetHoteltypes(p graphql.ResolveParams) (i interface{}, e error){
	hoteltypes, error:= model.GetAllHoteltype()

	return hoteltypes, error
}

func GetHotelfacilities(p graphql.ResolveParams) (i interface{}, e error){
	hotelfacility, error:= model.GetAllHotelfacility()

	return hotelfacility, error
}

func GetHotels(p graphql.ResolveParams) (i interface{}, e error){
	hotel, error:= model.GetAllHotel()

	return hotel, error
}

func GetHotelByName(p graphql.ResolveParams) (i interface{}, e error){
	name:=p.Args["hotel_name"].(string)
	hotel,err:= model.GetHotelByName(name)

	return hotel,err
}

func GetHotelByPrice(p graphql.ResolveParams) (i interface{}, e error){
	price:=p.Args["hotel_price"].(int)
	hotel,err:= model.GetHotelByPrice(price)

	return hotel,err
}

func GetHotelByFacility(p graphql.ResolveParams) (i interface{}, e error){
	facility:=p.Args["hotelfacility_id"].(int)
	hotel,err:= model.GetHotelByFacility(facility)

	return hotel,err
}

func GetHotelByArea(p graphql.ResolveParams) (i interface{}, e error){
	hotel_latitude:=p.Args["hotel_latitude"].(string)
	hotel_longitude:=p.Args["hotel_longitude"].(string)
	hotel,err:= model.GetHotelByArea(hotel_latitude, hotel_longitude)

	return hotel,err
}

func GetEntertainments(p graphql.ResolveParams) (i interface{}, e error){
	entertainments, error:= model.GetAllEntertainment()

	return entertainments, error
}

func GetEntertainmentByName(p graphql.ResolveParams) (i interface{}, e error){
	name:=p.Args["entertainment_name"].(string)
	entertainments, error:= model.GetEntertainmentByName(name)

	return entertainments, error
}

func GetEntertainmentByPrice(p graphql.ResolveParams) (i interface{}, e error){
	price:=p.Args["entertainment_price"].(int)
	entertainments, error:= model.GetEntertainmentByPrice(price)

	return entertainments, error
}

func GetEntertainmentByType(p graphql.ResolveParams) (i interface{}, e error){
	types:=p.Args["entertainment_type"].(string)
	entertainments, error:= model.GetEntertainmentByType(types)

	return entertainments, error
}

func GetEntertainmentByStartDate(p graphql.ResolveParams) (i interface{}, e error){
	startDate:=p.Args["entertainment_start_date"].(time.Time)
	entertainments, error:= model.GetEntertainmentByStartDate(startDate)

	return entertainments, error
}

func GetEntertainmentByEndDate(p graphql.ResolveParams) (i interface{}, e error){
	endDate:=p.Args["entertainment_end_date"].(time.Time)
	entertainments, error:= model.GetEntertainmentByStartDate(endDate)

	return entertainments, error
}

func GetBlogs(p graphql.ResolveParams) (i interface{}, e error){
	blogs, error:= model.GetAllBlog()

	return blogs, error
}

func GetTrains(p graphql.ResolveParams) (i interface{}, e error){
	trains, error:= model.GetAllTrain()

	return trains, error
}

func GetAeroplanes(p graphql.ResolveParams) (i interface{}, e error){
	aeroplanes, error:= model.GetAllAeroplane()

	return aeroplanes, error
}

func GetPromos(p graphql.ResolveParams) (i interface{}, e error){
	promos, error:= model.GetAllPromo()

	return promos, error
}

func GetCars(p graphql.ResolveParams) (i interface{}, e error){
	cars, error:= model.GetAllCar()

	return cars, error
}

func GetRooms(p graphql.ResolveParams) (i interface{}, e error){
	rooms, error:= model.GetAllRoom()

	return rooms, error
}

func GetRoomsByHotelId(p graphql.ResolveParams) (i interface{}, e error){
	hotel_id:=p.Args["hotel_id"].(int)
	rooms, error:= model.GetRoomByHotelId(hotel_id)

	return rooms, error
}

func SendSubscriptionToAll(p graphql.ResolveParams) (i interface{}, err error) {
	subscription := model.SendSubscriptionToAll()

	return subscription, nil
}

func GetLogin(p graphql.ResolveParams) (i interface{}, err error) {
	user_email:=p.Args["user_email"].(string)
	user_password:=p.Args["user_password"].(string)
	users, error:= model.GetLogin(user_email, user_password)

	return users, error
}

func GetAllOrder(p graphql.ResolveParams) (i interface{}, err error) {
	orders, error:= model.GetAllOrder()

	return orders, error
}

