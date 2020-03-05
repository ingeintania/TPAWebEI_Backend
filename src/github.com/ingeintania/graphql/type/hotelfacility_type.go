package _type

import "github.com/graphql-go/graphql"

var hotelfacility_type *graphql.Object

func GetHotelfacilityType() *graphql.Object{
	if hotelfacility_type == nil{
		hotelfacility_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"hotelfacility_type",
			Fields:graphql.Fields{
				"hotelfacility_id": &graphql.Field{
					Type:graphql.Int,
				},
				"hotelid": &graphql.Field{
					Type:graphql.Int,
				},
				"hotelfacility_name": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return hotelfacility_type
}
