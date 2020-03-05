package _type

import "github.com/graphql-go/graphql"

var location_type *graphql.Object

func GetLocationType() *graphql.Object{
	if location_type == nil{
		location_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"location_type",
			Fields:graphql.Fields{
				"location_id": &graphql.Field{
					Type:graphql.Int,
				},
				"location_name": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return location_type
}
