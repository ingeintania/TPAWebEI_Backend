package _type

import "github.com/graphql-go/graphql"

var hotel_type *graphql.Object

func GetHotelType() *graphql.Object{
	if hotel_type == nil{
		hotel_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"hotel_type",
			Fields:graphql.Fields{
				"hotel_id": &graphql.Field{
					Type:graphql.Int,
				},
				"hotel_name": &graphql.Field{
					Type:graphql.String,
				},
				"hotel_price": &graphql.Field{
					Type:graphql.Int,
				},
				"hotel_image_path": &graphql.Field{
					Type:graphql.String,
				},
				"location_id": &graphql.Field{
					Type:graphql.Int,
				},
				"hoteltype_id": &graphql.Field{
					Type:graphql.Int,
				},
				"hotel_address": &graphql.Field{
					Type:graphql.String,
				},
				"hotel_latitude": &graphql.Field{
					Type:graphql.String,
				},
				"hotel_longitude": &graphql.Field{
					Type:graphql.String,
				},
				"hotel_star": &graphql.Field{
					Type:graphql.Int,
				},
				"hotel_rating_count": &graphql.Field{
					Type:graphql.Int,
				},
				"hotel_availability": &graphql.Field{
					Type:graphql.String,
				},
				"hotelFacility": &graphql.Field{
					Type:graphql.NewList(GetHotelfacilityType()),
				},
				"hotel_facilities": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return hotel_type
}