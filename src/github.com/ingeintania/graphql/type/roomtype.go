package _type

import "github.com/graphql-go/graphql"

var room_type *graphql.Object

func GetRoomType() *graphql.Object{
	if room_type == nil{
		room_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"room_type",
			Fields:graphql.Fields{
				"room_id": &graphql.Field{
					Type:graphql.Int,
				},
				"hotel_id": &graphql.Field{
					Type:graphql.Int,
				},
				"room_image_path": &graphql.Field{
					Type:graphql.String,
				},
				"room_title": &graphql.Field{
					Type:graphql.String,
				},
				"room_depart_time": &graphql.Field{
					Type:graphql.String,
				},
				"room_cancel_status": &graphql.Field{
					Type:graphql.String,
				},
				"room_information": &graphql.Field{
					Type:graphql.String,
				},
				"room_facilities": &graphql.Field{
					Type:graphql.String,
				},
				"room_availability": &graphql.Field{
					Type:graphql.Int,
				},
				"room_price": &graphql.Field{
					Type:graphql.Int,
				},
			},
		})
	}
	return room_type

}