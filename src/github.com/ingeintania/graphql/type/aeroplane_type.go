package _type

import "github.com/graphql-go/graphql"

var aeroplane_type *graphql.Object

func GetAeroplaneType() *graphql.Object{
	if aeroplane_type == nil{
		aeroplane_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"aeroplane_type",
			Fields:graphql.Fields{
				"aeroplane_id": &graphql.Field{
					Type:graphql.Int,
				},
				"aeroplane_name": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_company": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_type": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_image_path": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_depart_date": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_depart_time": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_arrive_date": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_arrive_time": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_duration": &graphql.Field{
					Type:graphql.Int,
				},
				"aeroplane_price": &graphql.Field{
					Type:graphql.Int,
				},
				"aeroplane_price_raw": &graphql.Field{
					Type:graphql.Int,
				},
				"aeroplane_price_tax": &graphql.Field{
					Type:graphql.Int,
				},
				"aeroplane_transit": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_transit_duration": &graphql.Field{
					Type:graphql.Int,
				},
				"aeroplane_facilities": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_depart_location": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplane_arrive_location": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return aeroplane_type

}
