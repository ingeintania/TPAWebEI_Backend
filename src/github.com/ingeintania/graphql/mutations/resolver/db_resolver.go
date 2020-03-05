package resolver

import (
	"github.com/graphql-go/graphql"
	"github.com/ingeintania/model"
)

func InsertSubscription(p graphql.ResolveParams) (i interface{}, err error) {
	email := p.Args["email"].(string)

	newSubscription := model.InsertNewSubscription(email)

	return newSubscription, nil
}

func SendSubscriptionToAll(p graphql.ResolveParams) (i interface{}, err error) {
	subscription := model.SendSubscriptionToAll()

	return subscription, nil
}

func CreateUser(p graphql.ResolveParams) (i interface{}, e error){
	user_email:=p.Args["user_email"].(string)
	user_phone:=p.Args["user_phone"].(string)
	user_gender:=p.Args["user_gender"].(string)
	user_firstname:=p.Args["user_firstname"].(string)
	user_lastname:=p.Args["user_lastname"].(string)
	user_city:=p.Args["user_city"].(string)
	user_address:=p.Args["user_address"].(string)
	user_post_code:=p.Args["user_post_code"].(string)
	user_password:=p.Args["user_password"].(string)
	user_language:=p.Args["user_language"].(string)

	users, error := model.CreateUser(user_email, user_phone, user_gender, user_firstname,
		user_lastname, user_city, user_address, user_post_code, user_password, user_language)

	return users, error;
}

func UpdateUser(p graphql.ResolveParams) (i interface{}, e error){
	user_id:=p.Args["user_id"].(int)
	user_email:=p.Args["user_email"].(string)
	user_phone:=p.Args["user_phone"].(string)
	user_gender:=p.Args["user_gender"].(string)
	user_firstname:=p.Args["user_firstname"].(string)
	user_lastname:=p.Args["user_lastname"].(string)
	user_city:=p.Args["user_city"].(string)
	user_address:=p.Args["user_address"].(string)
	user_post_code:=p.Args["user_post_code"].(string)
	user_password:=p.Args["user_password"].(string)
	user_language:=p.Args["user_language"].(string)

	users, error := model.UpdateUser(user_id, user_email, user_phone, user_gender, user_firstname,
		user_lastname, user_city, user_address, user_post_code, user_password, user_language)

	return users, error;
}

func CreateLocation(p graphql.ResolveParams) (i interface{}, e error){
	location_name:=p.Args["location_name"].(string)

	locations, error := model.CreateLocation(location_name)
	return locations, error;
}

func CreateHoteltype(p graphql.ResolveParams) (i interface{}, e error){
	hoteltype_name:=p.Args["hoteltype_name"].(string)

	hoteltypes, error := model.CreateHoteltype(hoteltype_name)
	return hoteltypes, error;
}

func CreateHotelfacility(p graphql.ResolveParams) (i interface{}, e error){
	hotel_id:=p.Args["hotel_id"].(int)
	hotelfacility_name:=p.Args["hotelfacility_name"].(string)

	hotelfacilities, error := model.CreateHotelfacility(hotelfacility_name, hotel_id)
	return hotelfacilities, error;
}

func CreateHotel(p graphql.ResolveParams) (i interface{}, e error){
	hotel_name:=p.Args["hotel_name"].(string)
	hotel_price:=p.Args["hotel_price"].(int)
	hotel_image_path:=p.Args["hotel_image_path"].(string)
	location_id:=p.Args["location_id"].(int)
	hoteltype_id:=p.Args["hoteltype_id"].(int)
	hotel_address:=p.Args["hotel_address"].(string)
	hotel_latitude:=p.Args["hotel_latitude"].(string)
	hotel_longitude:=p.Args["hotel_longitude"].(string)
	hotel_star:=p.Args["hotel_star"].(int)
	hotel_rating_count:=p.Args["hotel_rating_count"].(int)
	hotel_availability:=p.Args["hotel_availability"].(string)
	hotel_facilities:=p.Args["hotel_facilities"].(string)

	hotel, error := model.CreateHotel(hotel_name, hotel_price, hotel_image_path, location_id,
		hoteltype_id, hotel_address, hotel_latitude, hotel_longitude,
		hotel_star, hotel_rating_count, hotel_availability, hotel_facilities)
	return hotel, error;
}

func UpdateHotel(p graphql.ResolveParams) (i interface{}, e error){
	hotel_id:=p.Args["hotel_id"].(int)
	hotel_name:=p.Args["hotel_name"].(string)
	hotel_price:=p.Args["hotel_price"].(int)
	hotel_image_path:=p.Args["hotel_image_path"].(string)
	location_id:=p.Args["location_id"].(int)
	hoteltype_id:=p.Args["hoteltype_id"].(int)
	hotel_address:=p.Args["hotel_address"].(string)
	hotel_latitude:=p.Args["hotel_latitude"].(string)
	hotel_longitude:=p.Args["hotel_longitude"].(string)
	hotel_star:=p.Args["hotel_star"].(int)
	hotel_rating_count:=p.Args["hotel_rating_count"].(int)
	hotel_availability:=p.Args["hotel_availability"].(string)
	hotel_facilities:=p.Args["hotel_facilities"].(string)

	hotel, error := model.UpdateHotel(hotel_id, hotel_name, hotel_price, hotel_image_path, location_id,
		hoteltype_id, hotel_address, hotel_latitude, hotel_longitude,
		hotel_star, hotel_rating_count, hotel_availability, hotel_facilities)
	return hotel, error;
}

func DeleteHotel(p graphql.ResolveParams) (i interface{}, e error){
	hotel_id:=p.Args["hotel_id"].(int)

	hotel, error := model.DeleteHotel(hotel_id)
	return hotel, error;
}


func CreateEntertainment(p graphql.ResolveParams) (i interface{}, e error){
	entertainment_name:=p.Args["entertainment_name"].(string)
	entertainment_price:=p.Args["entertainment_price"].(int)
	entertainment_image_path1:=p.Args["entertainment_image_path1"].(string)
	entertainment_image_path2:=p.Args["entertainment_image_path2"].(string)
	entertainment_image_path3:=p.Args["entertainment_image_path3"].(string)
	entertainment_image_path4:=p.Args["entertainment_image_path4"].(string)
	entertainment_image_path5:=p.Args["entertainment_image_path5"].(string)
	entertainment_address:=p.Args["entertainment_address"].(string)
	location_id:=p.Args["location_id"].(int)
	entertainment_latitude:=p.Args["entertainment_latitude"].(string)
	entertainment_longitude:=p.Args["entertainment_longitude"].(string)
	entertainment_type:=p.Args["entertainment_type"].(string)
	entertainment_start_date:=p.Args["entertainment_start_date"].(string)
	entertainment_end_date:=p.Args["entertainment_end_date"].(string)
	entertainment_description:=p.Args["entertainment_description"].(string)
	entertainment_t_c:=p.Args["entertainment_t_c"].(string)

	entertainment, error := model.CreateEntertainment(entertainment_name, entertainment_price,
		entertainment_image_path1,entertainment_image_path2,entertainment_image_path3,entertainment_image_path4,
		entertainment_image_path5, location_id, entertainment_address,
		entertainment_latitude, entertainment_longitude, entertainment_type, entertainment_start_date, entertainment_end_date,
		entertainment_description, entertainment_t_c)
	return entertainment, error;
}

func UpdateEntertainment(p graphql.ResolveParams) (i interface{}, e error){
	entertainment_id:=p.Args["entertainment_id"].(int)
	entertainment_name:=p.Args["entertainment_name"].(string)
	entertainment_price:=p.Args["entertainment_price"].(int)
	entertainment_image_path1:=p.Args["entertainment_image_path1"].(string)
	entertainment_image_path2:=p.Args["entertainment_image_path2"].(string)
	entertainment_image_path3:=p.Args["entertainment_image_path3"].(string)
	entertainment_image_path4:=p.Args["entertainment_image_path4"].(string)
	entertainment_image_path5:=p.Args["entertainment_image_path5"].(string)
	entertainment_address:=p.Args["entertainment_address"].(string)
	location_id:=p.Args["location_id"].(int)
	entertainment_latitude:=p.Args["entertainment_latitude"].(string)
	entertainment_longitude:=p.Args["entertainment_longitude"].(string)
	entertainment_type:=p.Args["entertainment_type"].(string)
	entertainment_start_date:=p.Args["entertainment_start_date"].(string)
	entertainment_end_date:=p.Args["entertainment_end_date"].(string)
	entertainment_description:=p.Args["entertainment_description"].(string)
	entertainment_t_c:=p.Args["entertainment_t_c"].(string)

	entertainment, error := model.UpdateEntertainment(entertainment_id, entertainment_name, entertainment_price,
		entertainment_image_path1,entertainment_image_path2,entertainment_image_path3,entertainment_image_path4,
		entertainment_image_path5, location_id, entertainment_address,
		entertainment_latitude, entertainment_longitude, entertainment_type, entertainment_start_date, entertainment_end_date,
		entertainment_description, entertainment_t_c)
	return entertainment, error;
}

func DeleteEntertainment(p graphql.ResolveParams) (i interface{}, e error){
	entertainment_id:=p.Args["entertainment_id"].(int)


	entertainment, error := model.DeleteEntertainment(entertainment_id)
	return entertainment, error;
}

func CreateBlog(p graphql.ResolveParams) (i interface{}, e error){
	blog_image_path:=p.Args["blog_image_path"].(string)
	blog_title:=p.Args["blog_title"].(string)
	blog_content:=p.Args["blog_content"].(string)
	blog_viewers:=p.Args["blog_viewers"].(int)
	blog_category:=p.Args["blog_category"].(string)
	blog_writer:=p.Args["blog_writer"].(string)

	blogs, error := model.CreateBlog(blog_image_path, blog_title, blog_content, blog_viewers, blog_category, blog_writer)
	return blogs, error;
}

func UpdateBlog(p graphql.ResolveParams) (i interface{}, e error){
	blog_id:=p.Args["blog_id"].(int)
	blog_image_path:=p.Args["blog_image_path"].(string)
	blog_title:=p.Args["blog_title"].(string)
	blog_content:=p.Args["blog_content"].(string)
	blog_viewers:=p.Args["blog_viewers"].(int)
	blog_category:=p.Args["blog_category"].(string)
	blog_writer:=p.Args["blog_writer"].(string)

	blogs, error := model.UpdateBlog(blog_id,blog_image_path, blog_title, blog_content, blog_viewers, blog_category, blog_writer)
	return blogs, error;
}

func DeleteBlog(p graphql.ResolveParams) (i interface{}, e error){
	blog_id:=p.Args["blog_id"].(int)

	blogs, error := model.DeleteBlog(blog_id)
	return blogs, error;
}

func CreateTrain(p graphql.ResolveParams) (i interface{}, e error){
	train_name	 :=p.Args["train_name"].(string)
	train_class:=p.Args["train_class"].(string)
	train_type:=p.Args["train_type"].(string)
	train_depart_time:=p.Args["train_depart_time"].(string)
	train_depart_location:=p.Args["train_depart_location"].(string)
	train_arrive_time:=p.Args["train_arrive_time"].(string)
	train_arrive_location:=p.Args["train_arrive_location"].(string)
	train_duration:=p.Args["train_duration"].(int)
	train_detail_price:=p.Args["train_detail_price"].(int)

	trains, error := model.CreateTrain(train_name, train_class,
		train_type, train_depart_time, train_depart_location,
		train_arrive_time, train_arrive_location, train_duration, train_detail_price)
	return trains, error;
}

func UpdateTrain(p graphql.ResolveParams) (i interface{}, e error){
	train_id	 :=p.Args["train_id"].(int)
	train_name	 :=p.Args["train_name"].(string)
	train_class:=p.Args["train_class"].(string)
	train_type:=p.Args["train_type"].(string)
	train_depart_time:=p.Args["train_depart_time"].(string)
	train_depart_location:=p.Args["train_depart_location"].(string)
	train_arrive_time:=p.Args["train_arrive_time"].(string)
	train_arrive_location:=p.Args["train_arrive_location"].(string)
	train_duration:=p.Args["train_duration"].(int)
	train_detail_price:=p.Args["train_detail_price"].(int)

	trains, error := model.UpdateTrain(train_id, train_name, train_class,
		train_type, train_depart_time, train_depart_location,
		train_arrive_time, train_arrive_location, train_duration, train_detail_price)
	return trains, error;
}

func DeleteTrain(p graphql.ResolveParams) (i interface{}, e error){
	train_id	 :=p.Args["train_id"].(int)


	trains, error := model.DeleteTrain(train_id)
	return trains, error;
}

func CreateAeroplane(p graphql.ResolveParams) (i interface{}, e error){
	aeroplane_name	 :=p.Args["aeroplane_name"].(string)
	aeroplane_company:=p.Args["aeroplane_company"].(string)
	aeroplane_type:=p.Args["aeroplane_type"].(string)
	aeroplane_image_path:=p.Args["aeroplane_image_path"].(string)
	aeroplane_depart_date:=p.Args["aeroplane_depart_date"].(string)
	aeroplane_depart_time:=p.Args["aeroplane_depart_time"].(string)
	aeroplane_arrive_date:=p.Args["aeroplane_arrive_date"].(string)
	aeroplane_arrive_time:=p.Args["aeroplane_arrive_time"].(string)
	aeroplane_duration:=p.Args["aeroplane_duration"].(int)
	aeroplane_price:=p.Args["aeroplane_price"].(int)
	aeroplane_price_raw:=p.Args["aeroplane_price_raw"].(int)
	aeroplane_price_tax:=p.Args["aeroplane_price_tax"].(int)
	aeroplane_transit:=p.Args["aeroplane_transit"].(string)
	aeroplane_transit_duration:=p.Args["aeroplane_transit_duration"].(int)
	aeroplane_facilities:=p.Args["aeroplane_facilities"].(string)
	aeroplane_depart_location:=p.Args["aeroplane_depart_location"].(string)
	aeroplane_arrive_location:=p.Args["aeroplane_arrive_location"].(string)

	aeroplanes, error := model.CreateAeroplane(
		aeroplane_name, aeroplane_company, aeroplane_type,
		aeroplane_image_path, aeroplane_depart_date,
		aeroplane_depart_time, aeroplane_arrive_date,
		aeroplane_arrive_time, aeroplane_duration,
		aeroplane_price, aeroplane_price_raw, aeroplane_price_tax,
		aeroplane_transit, aeroplane_transit_duration,
		aeroplane_facilities, aeroplane_depart_location,
		aeroplane_arrive_location)
	return aeroplanes, error;
}

func UpdateAeroplane(p graphql.ResolveParams) (i interface{}, e error){
	aeroplane_id:=p.Args["aeroplane_id"].(int)
	aeroplane_name	 :=p.Args["aeroplane_name"].(string)
	aeroplane_company:=p.Args["aeroplane_company"].(string)
	aeroplane_type:=p.Args["aeroplane_type"].(string)
	aeroplane_image_path:=p.Args["aeroplane_image_path"].(string)
	aeroplane_depart_date:=p.Args["aeroplane_depart_date"].(string)
	aeroplane_depart_time:=p.Args["aeroplane_depart_time"].(string)
	aeroplane_arrive_date:=p.Args["aeroplane_arrive_date"].(string)
	aeroplane_arrive_time:=p.Args["aeroplane_arrive_time"].(string)
	aeroplane_duration:=p.Args["aeroplane_duration"].(int)
	aeroplane_price:=p.Args["aeroplane_price"].(int)
	aeroplane_price_raw:=p.Args["aeroplane_price_raw"].(int)
	aeroplane_price_tax:=p.Args["aeroplane_price_tax"].(int)
	aeroplane_transit:=p.Args["aeroplane_transit"].(string)
	aeroplane_transit_duration:=p.Args["aeroplane_transit_duration"].(int)
	aeroplane_facilities:=p.Args["aeroplane_facilities"].(string)
	aeroplane_depart_location:=p.Args["aeroplane_depart_location"].(string)
	aeroplane_arrive_location:=p.Args["aeroplane_arrive_location"].(string)

	aeroplanes, error := model.UpdateAeroplane(aeroplane_id,
		aeroplane_name, aeroplane_company, aeroplane_type,
		aeroplane_image_path, aeroplane_depart_date,
		aeroplane_depart_time, aeroplane_arrive_date,
		aeroplane_arrive_time, aeroplane_duration,
		aeroplane_price, aeroplane_price_raw, aeroplane_price_tax,
		aeroplane_transit, aeroplane_transit_duration,
		aeroplane_facilities, aeroplane_depart_location,
		aeroplane_arrive_location)
	return aeroplanes, error;
}

func DeleteAeroplane(p graphql.ResolveParams) (i interface{}, e error){
	aeroplane_id:=p.Args["aeroplane_id"].(int)

	aeroplanes, error := model.DeleteAeroplane(aeroplane_id)
	return aeroplanes, error;
}

func CreatePromo(p graphql.ResolveParams) (i interface{}, e error){
	promo_image_path:=p.Args["promo_image_path"].(string)
	promo_title:=p.Args["promo_title"].(string)
	promo_content:=p.Args["promo_content"].(string)
	promo_category:=p.Args["promo_category"].(string)
	promo_periode:=p.Args["promo_periode"].(string)

	promos, error := model.CreatePromo(promo_image_path, promo_title, promo_content, promo_category, promo_periode)
	return promos, error;
}

func CreateCar(p graphql.ResolveParams) (i interface{}, e error){
	car_image_path:=p.Args["car_image_path"].(string)
	car_model:=p.Args["car_model"].(string)
	car_price:=p.Args["car_price"].(int)
	car_passenger:=p.Args["car_passenger"].(int)
	car_luggage:=p.Args["car_luggage"].(int)
	car_vendor:=p.Args["car_vendor"].(string)
	car_location:=p.Args["car_location"].(string)

	cars, error := model.CreateCar(car_image_path, car_model, car_price, car_passenger, car_luggage, car_vendor, car_location)
	return cars, error;
}

func CreateOrder(p graphql.ResolveParams) (i interface{}, e error){
	user_name:=p.Args["user_name"].(string)
	user_email:=p.Args["user_email"].(string)
	user_phone:=p.Args["user_phone"].(string)
	ticket_quantity:=p.Args["ticket_quantity"].(int)
	ticket_title:=p.Args["ticket_title"].(string)
	ticket_name:=p.Args["ticket_name"].(string)
	ticket_nationality:=p.Args["ticket_nationality"].(string)
	order_payment_method:=p.Args["order_payment_method"].(string)
	order_promo_code:=p.Args["order_promo_code"].(string)
	order_price:=p.Args["order_price"].(int)

	promos, error := model.CreateOrder(user_name, user_email, user_phone,
		ticket_quantity, ticket_title, ticket_name, ticket_nationality,
		order_payment_method,order_promo_code,order_price)
	return promos, error;
}