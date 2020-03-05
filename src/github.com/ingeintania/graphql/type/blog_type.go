package _type

import "github.com/graphql-go/graphql"

var blog_type *graphql.Object

func GetBlogType() *graphql.Object{
	if blog_type == nil{
		blog_type = graphql.NewObject(graphql.ObjectConfig{
			Name:"blog_type",
			Fields:graphql.Fields{
				"blog_id": &graphql.Field{
					Type:graphql.Int,
				},
				"blog_image_path": &graphql.Field{
					Type:graphql.String,
				},
				"blog_title": &graphql.Field{
					Type:graphql.String,
				},
				"blog_content": &graphql.Field{
					Type:graphql.String,
				},
				"blog_viewers": &graphql.Field{
					Type:graphql.Int,
				},
				"blog_category": &graphql.Field{
					Type:graphql.String,
				},
				"blog_writer": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}
	return blog_type

}
