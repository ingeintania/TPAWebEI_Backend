package query

import (
	res "github.com/ingeintania/graphql/query/resolver"
	typ "github.com/ingeintania/graphql/type"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name:"RootQuery",
		Fields: graphql.Fields{

			"users":{
				Type:	graphql.NewList(typ.GetUserType()),
				Resolve: res.GetUsers,
				Description: "Get All Users",
			},
			"getlogin":{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"user_email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"user_password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetLogin,
				Description: "Get Login",
			},
			"userbyemail":{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"user_email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetUserByEmail,
				Description: "Get User Based on Email",
			},
			"userbyphone":{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"user_phone": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.GetUserByPhone,
				Description: "Get User Based on Phone",
			},
			"userbyphoneoremail":{
				Type: graphql.NewList(typ.GetUserType()),
				Args:graphql.FieldConfigArgument{
					"arg": &graphql.ArgumentConfig{
						Type: graphql.String,
					},

				},
				Resolve: res.GetUserByPhoneOrEmail,
				Description: "Find User to Validate Login",
			},
			"locations":{
				Type:	graphql.NewList(typ.GetLocationType()),
				Resolve: res.GetLocations,
				Description: "Get All Locations",
			},
			"hoteltypes":{
				Type:	graphql.NewList(typ.GetHoteltypeType()),
				Resolve: res.GetHoteltypes,
				Description: "Get All Locations",
			},
			"hotelfacilities":{
				Type:	graphql.NewList(typ.GetHotelfacilityType()),
				Resolve: res.GetHotelfacilities,
				Description: "Get All Hotel Facilities",
			},
			"hotels":{
				Type:	graphql.NewList(typ.GetHotelType()),
				Resolve: res.GetHotels,
				Description: "Get All Hotels",
			},
			"hotelbyname":{
				Type: graphql.NewList(typ.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"hotel_name": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.GetHotelByName,
				Description: "Get Hotel Based on Name",
			},
			"hotelbyprice":{
				Type: graphql.NewList(typ.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"hotel_price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.GetHotelByPrice,
				Description: "Get Hotel Based on Price",
			},
			"hotelbyfacility":{
				Type: graphql.NewList(typ.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"hotelfacility_id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.GetHotelByFacility,
				Description: "Get Hotel Based on Facility",
			},
			"hotelbyarea":{
				Type: graphql.NewList(typ.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"hotel_latitude": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"hotel_longitude": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetHotelByArea,
				Description: "Get Hotel Based on Area",
			},
			"entertainments":{
				Type:	graphql.NewList(typ.GetEntertainmentType()),
				Resolve: res.GetEntertainments,
				Description: "Get All Entertainments Details",
			},
			"blogs":{
				Type:	graphql.NewList(typ.GetBlogType()),
				Resolve: res.GetBlogs,
				Description: "Get All Blogs",
			},
			"trains":{
				Type:	graphql.NewList(typ.GetTrainType()),
				Resolve: res.GetTrains,
				Description: "Get All Trains",
			},
			"aeroplanes":{
				Type:	graphql.NewList(typ.GetAeroplaneType()),
				Resolve: res.GetAeroplanes,
				Description: "Get All Aeroplanes",
			},
			"promos":{
				Type:	graphql.NewList(typ.GetPromoType()),
				Resolve: res.GetPromos,
				Description: "Get All Promos",
			},
			"cars":{
				Type:	graphql.NewList(typ.GetCarType()),
				Resolve: res.GetCars,
				Description: "Get All Cars",
			},
			"subscriptions": {
				Type:        typ.GetSubscriptionType(),
				Resolve:     res.SendSubscriptionToAll,
				Description: "Send Subscription",
			},
			"rooms": {
				Type:        graphql.NewList(typ.GetRoomType()),
				Resolve:     res.GetRooms,
				Description: "Get Room",
			},
			"roomByHotelId": {
				Type: graphql.NewList(typ.GetRoomType()),
				Args: graphql.FieldConfigArgument{
					"hotel_id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetRoomsByHotelId,
				Description: "Get Hotel Based on Area",
			},
			"orders": {
				Type:        graphql.NewList(typ.GetOrderType()),
				Resolve:     res.GetAllOrder,
				Description: "Get Order",
			},
		},
	})
}
