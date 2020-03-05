package mutations

import (
	"github.com/graphql-go/graphql"
	res "github.com/ingeintania/graphql/mutations/resolver"
	typ "github.com/ingeintania/graphql/type"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name:"RootMutation",
		Fields: graphql.Fields{
			"insertsubscription": &graphql.Field{
				Type: typ.GetSubscriptionType(),
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: res.InsertSubscription,
			},
			"createuser":{
				Type: typ.GetUserType() ,
				Args: graphql.FieldConfigArgument{
					"user_email": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_phone": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_gender": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_firstname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_lastname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_city": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_address": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_post_code": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_password": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_language": &graphql.ArgumentConfig{
						Type: graphql.String,

					},

				},
				Resolve: res.CreateUser,
				Description: "Get User",
			},
			"updateuser":{
				Type: typ.GetUserType() ,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"user_email": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_phone": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_gender": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_firstname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_lastname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_city": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_address": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_post_code": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_password": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"user_language": &graphql.ArgumentConfig{
						Type: graphql.String,

					},

				},
				Resolve: res.UpdateUser,
				Description: "Update User",
			},
			"createlocation":{
				Type: typ.GetLocationType() ,
				Args: graphql.FieldConfigArgument{
					"location_name": &graphql.ArgumentConfig{
						Type: graphql.String,

					},

				},
				Resolve: res.CreateLocation,
				Description: "Get Location",
			},
			"createhoteltype":{
				Type: typ.GetHoteltypeType() ,
				Args: graphql.FieldConfigArgument{
					"hoteltype_name": &graphql.ArgumentConfig{
						Type: graphql.String,

					},

				},
				Resolve: res.CreateHoteltype,
				Description: "Get Hotel Type",
			},
			"createhotelfacility":{
				Type: typ.GetHotelfacilityType() ,
				Args: graphql.FieldConfigArgument{
					"hotel_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotelfacility_name": &graphql.ArgumentConfig{
						Type: graphql.String,

					},

				},
				Resolve: res.CreateHotelfacility,
				Description: "Get Hotel Facility",
			},
			"createhotel":{
				Type: typ.GetHotelType() ,
				Args: graphql.FieldConfigArgument{
					"hotel_name": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"hotel_price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_image_path": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"location_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hoteltype_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_address": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"hotel_latitude": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"hotel_longitude": &graphql.ArgumentConfig{
						Type: graphql.String,

					},

					"hotel_star": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_rating_count": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_availability": &graphql.ArgumentConfig{
						Type: graphql.String,

					},

					"hotel_facilities": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.CreateHotel,
				Description: "Get Hotel",
			},
			"updatehotel":{
				Type: typ.GetHotelType() ,
				Args: graphql.FieldConfigArgument{
					"hotel_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_name": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"hotel_price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_image_path": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"location_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hoteltype_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_address": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"hotel_latitude": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"hotel_longitude": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"hotel_star": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_rating_count": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"hotel_availability": &graphql.ArgumentConfig{
						Type: graphql.String,

					},

					"hotel_facilities": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.UpdateHotel,
				Description: "Update Hotel",
			},
			"deletehotel":{
				Type: typ.GetHotelType() ,
				Args: graphql.FieldConfigArgument{
					"hotel_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},

				},
				Resolve: res.DeleteHotel,
				Description: "Delete Hotel",
			},
			"createentertainment":{
				Type: typ.GetEntertainmentType() ,
				Args: graphql.FieldConfigArgument{
					"entertainment_name": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"entertainment_image_path1": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_image_path2": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_image_path3": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_image_path4": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_image_path5": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"location_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"entertainment_address": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_latitude": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_longitude": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_type": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_start_date": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_end_date": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_description": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_t_c": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.CreateEntertainment,
				Description: "Get Entertainment",
			},
			"updateentertainment":{
				Type: typ.GetEntertainmentType() ,
				Args: graphql.FieldConfigArgument{
					"entertainment_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"entertainment_name": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"entertainment_image_path1": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_image_path2": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_image_path3": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_image_path4": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_image_path5": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"location_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"entertainment_address": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_latitude": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_longitude": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_type": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_start_date": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_end_date": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_description": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"entertainment_t_c": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.UpdateEntertainment,
				Description: "Update Entertainment",
			},
			"deleteentertainment":{
				Type: typ.GetEntertainmentType() ,
				Args: graphql.FieldConfigArgument{
					"entertainment_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},

				},
				Resolve: res.DeleteEntertainment,
				Description: "Delete Entertainment",
			},
			"createblog":{
				Type: typ.GetBlogType() ,
				Args: graphql.FieldConfigArgument{
					"blog_image_path": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"blog_title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"blog_content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"blog_viewers": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"blog_category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"blog_writer": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateBlog,
				Description: "Get Blog",
			},
			"updateblog":{
				Type: typ.GetBlogType() ,
				Args: graphql.FieldConfigArgument{
					"blog_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"blog_image_path": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"blog_title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"blog_content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"blog_viewers": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"blog_category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"blog_writer": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateBlog,
				Description: "Update Blog",
			},
			"deleteblog":{
				Type: typ.GetBlogType() ,
				Args: graphql.FieldConfigArgument{
					"blog_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},

				},
				Resolve: res.DeleteBlog,
				Description: "Delete Blog",
			},
			"createtrain":{
				Type: typ.GetTrainType() ,
				Args: graphql.FieldConfigArgument{
					"train_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_class": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_depart_time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_depart_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_arrive_time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_arrive_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"train_detail_price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateTrain,
				Description: "Get Train",
			},
			"updatetrain":{
				Type: typ.GetTrainType() ,
				Args: graphql.FieldConfigArgument{
					"train_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"train_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_class": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_depart_time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_depart_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_arrive_time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_arrive_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"train_duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"train_detail_price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.UpdateTrain,
				Description: "Update Train",
			},
			"deletetrain":{
				Type: typ.GetTrainType() ,
				Args: graphql.FieldConfigArgument{
					"train_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},

				},
				Resolve: res.DeleteTrain,
				Description: "Delete Train",
			},
			"createaeroplane":{
				Type: typ.GetAeroplaneType() ,
				Args: graphql.FieldConfigArgument{
					"aeroplane_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_company": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_image_path": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_depart_date": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_depart_time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_arrive_date": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_arrive_time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_price_raw": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_price_tax": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_transit": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_transit_duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_facilities": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_depart_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_arrive_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateAeroplane,
				Description: "Get Aeroplane",
			},
			"updateaeroplane":{
				Type: typ.GetAeroplaneType() ,
				Args: graphql.FieldConfigArgument{
					"aeroplane_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_company": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_image_path": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_depart_date": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_depart_time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_arrive_date": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_arrive_time": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_price_raw": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_price_tax": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_transit": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_transit_duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"aeroplane_facilities": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_depart_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"aeroplane_arrive_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.UpdateAeroplane,
				Description: "Update Aeroplane",
			},
			"deleteaeroplane":{
				Type: typ.GetAeroplaneType() ,
				Args: graphql.FieldConfigArgument{
					"aeroplane_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},

				},
				Resolve: res.DeleteAeroplane,
				Description: "Delete Aeroplane",
			},
			"createpromo":{
				Type: typ.GetPromoType() ,
				Args: graphql.FieldConfigArgument{
					"promo_image_path": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"promo_title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"promo_content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"promo_category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"promo_periode": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreatePromo,
				Description: "Get Promo",
			},
			"createcar":{
				Type: typ.GetCarType() ,
				Args: graphql.FieldConfigArgument{
					"car_image_path": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"car_model": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"car_price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"car_passenger": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"car_luggage": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"car_vendor": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"car_location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.CreateCar,
				Description: "Get Car",
			},
			"createorder":{
				Type: typ.GetOrderType() ,
				Args: graphql.FieldConfigArgument{
					"user_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"user_email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"user_phone": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"ticket_quantity": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"ticket_title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"ticket_name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"ticket_nationality": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"order_payment_method": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"order_promo_code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"order_price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.CreateOrder,
				Description: "Get Car",
			},
		},
	})
}
