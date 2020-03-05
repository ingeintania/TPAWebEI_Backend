package _type

import "github.com/graphql-go/graphql"

var promo_type *graphql.Object

func GetPromoType() *graphql.Object{
	if promo_type == nil{
		promo_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"promo_type",
			Fields:graphql.Fields{
				"promo_id": &graphql.Field{
					Type:graphql.Int,
				},
				"promo_image_path": &graphql.Field{
					Type:graphql.String,
				},
				"promo_title": &graphql.Field{
					Type:graphql.String,
				},
				"promo_content": &graphql.Field{
					Type:graphql.String,
				},
				"promo_category": &graphql.Field{
					Type:graphql.String,
				},
				"promo_periode": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return promo_type

}
