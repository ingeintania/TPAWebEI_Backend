package _type

import "github.com/graphql-go/graphql"

var train_type *graphql.Object

func GetTrainType() *graphql.Object{
	if train_type == nil{
		train_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"train_type",
			Fields:graphql.Fields{
				"train_id": &graphql.Field{
					Type:graphql.Int,
				},
				"train_name": &graphql.Field{
					Type:graphql.String,
				},
				"train_class": &graphql.Field{
					Type:graphql.String,
				},
				"train_type": &graphql.Field{
					Type:graphql.String,
				},
				"train_depart_time": &graphql.Field{
					Type:graphql.String,
				},
				"train_depart_location": &graphql.Field{
					Type:graphql.String,
				},
				"train_arrive_time": &graphql.Field{
					Type:graphql.String,
				},
				"train_arrive_location": &graphql.Field{
					Type:graphql.String,
				},
				"train_duration": &graphql.Field{
					Type:graphql.Int,
				},
				"train_detail_price": &graphql.Field{
					Type:graphql.Int,
				},
			},
		})
	}
	return train_type

}
