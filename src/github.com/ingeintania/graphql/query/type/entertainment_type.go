package _type

import "github.com/graphql-go/graphql"

var entertainment_type *graphql.Object

func GetEntertainmentType() *graphql.Object{
	if entertainment_type == nil{
		entertainment_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"entertainment_type",
			Fields:graphql.Fields{
				"entertainment_id": &graphql.Field{
					Type:graphql.Int,
				},
				"entertainment_name": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_price": &graphql.Field{
					Type:graphql.Int,
				},
				"entertainment_image_path1": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_image_path2": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_image_path3": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_image_path4": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_image_path5": &graphql.Field{
					Type:graphql.String,
				},
				"location": &graphql.Field{
					Type:GetLocationType(),
				},
				"entertainment_address": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_latitude": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_longitude": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_type": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_start_date": &graphql.Field{
					Type:graphql.DateTime,
				},
				"entertainment_end_date": &graphql.Field{
					Type:graphql.DateTime,
				},
				"entertainment_description": &graphql.Field{
					Type:graphql.String,
				},
				"entertainment_t_c": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return entertainment_type
}