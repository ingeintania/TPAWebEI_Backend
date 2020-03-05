package _type

import "github.com/graphql-go/graphql"

var hoteltype_type *graphql.Object

func GetHoteltypeType() *graphql.Object{
	if hoteltype_type == nil{
		hoteltype_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"hoteltype_type",
			Fields:graphql.Fields{
				"hoteltype_id": &graphql.Field{
					Type:graphql.Int,
				},
				"hoteltype_name": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return hoteltype_type
}
