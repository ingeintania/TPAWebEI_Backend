package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ingeintania/database"
	"github.com/ingeintania/graphql/mutations"
	"github.com/ingeintania/graphql/query"
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
	fmt.Print(":D")
	wrapped := middleware.CorsMiddleware(h)

	token := "akuasehat"

	http.Handle("/"+token+"/api",wrapped)
	log.Fatalln(http.ListenAndServe(":2003",nil))
}
