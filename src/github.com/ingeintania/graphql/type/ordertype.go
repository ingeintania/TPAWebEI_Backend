package _type

import "github.com/graphql-go/graphql"

var order_type *graphql.Object

func GetOrderType() *graphql.Object{
	if order_type == nil{
		order_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"order_type",
			Fields:graphql.Fields{
				"order_id": &graphql.Field{
					Type:graphql.Int,
				},
				"user_name": &graphql.Field{
					Type:graphql.String,
				},
				"user_email": &graphql.Field{
					Type:graphql.String,
				},
				"user_phone": &graphql.Field{
					Type:graphql.String,
				},
				"ticket_quantity": &graphql.Field{
					Type:graphql.Int,
				},
				"ticket_title": &graphql.Field{
					Type:graphql.String,
				},
				"ticket_name": &graphql.Field{
					Type:graphql.String,
				},
				"ticket_nationality": &graphql.Field{
					Type:graphql.String,
				},
				"order_payment_method": &graphql.Field{
					Type:graphql.String,
				},
				"order_promo_code": &graphql.Field{
					Type:graphql.String,
				},
				"order_price": &graphql.Field{
					Type:graphql.Int,
				},
			},
		})
	}
	return order_type

}
