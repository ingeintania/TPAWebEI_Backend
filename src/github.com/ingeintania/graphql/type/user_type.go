package _type

import "github.com/graphql-go/graphql"

var user_type *graphql.Object

func GetUserType() *graphql.Object{
	if user_type == nil{
		user_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"user_type",
			Fields:graphql.Fields{
				"user_id": &graphql.Field{
					Type:graphql.Int,
				},
				"user_email": &graphql.Field{
					Type:graphql.String,
				},
				"user_phone": &graphql.Field{
					Type:graphql.String,
				},
				"user_gender": &graphql.Field{
					Type:graphql.String,
				},
				"user_firstname": &graphql.Field{
					Type:graphql.String,
				},
				"user_lastname": &graphql.Field{
					Type:graphql.String,
				},
				"user_city": &graphql.Field{
					Type:graphql.String,
				},
				"user_address": &graphql.Field{
					Type:graphql.String,
				},
				"user_post_code": &graphql.Field{
					Type:graphql.String,
				},
				"user_password": &graphql.Field{
					Type:graphql.String,
				},
				"user_language": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return user_type

}