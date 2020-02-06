package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ingeintania/database"
	"github.com/ingeintania/middleware"
	"log"
	"net/http"
)

func main(){
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query.GetRoot(),
		Mutation: mutations.GetRoot(),
	})

	if err != nil{
		panic(err.Error())
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty:	true,
		GraphiQL: true,
		Playground: true,
	})

	db,error := database.Connect()
	if error != nil{
		fmt.Print("Ini ga connect")
	}else {
		fmt.Print("Ini Connect")
	}
	db.Close()
	fmt.Print(":((((")
	wrapped := middleware.CorsMiddleware(h)
	log.Fatalln(http.ListenAndServe(":2000",wrapped))
}
