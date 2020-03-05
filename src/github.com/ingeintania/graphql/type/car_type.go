package _type

import "github.com/graphql-go/graphql"

var car_type *graphql.Object

func GetCarType() *graphql.Object{
	if car_type == nil{
		car_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"car_type",
			Fields:graphql.Fields{
				"car_id": &graphql.Field{
					Type:graphql.Int,
				},
				"car_image_path": &graphql.Field{
					Type:graphql.String,
				},
				"car_model": &graphql.Field{
					Type:graphql.String,
				},
				"car_price": &graphql.Field{
					Type:graphql.Int,
				},
				"car_passenger": &graphql.Field{
					Type:graphql.Int,
				},
				"car_luggage": &graphql.Field{
					Type:graphql.Int,
				},
				"car_vendor": &graphql.Field{
					Type:graphql.String,
				},
				"car_location": &graphql.Field{
					Type:graphql.String,
				},

			},
		})
	}
	return car_type

}
