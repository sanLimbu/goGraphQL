package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/sanLimbu/gopheragency/gopher"
	"github.com/sanLimbu/gopheragency/job"
	"github.com/sanLimbu/gopheragency/schema"
)

func main() {

	//create a gopher repository
	gopherService := gopher.NewService(gopher.NewMemoryRepository(),
		job.NewMemoryRepository(),
	)

	schema, err := schema.GenerateSchema(&gopherService)
	if err != nil {
		panic(err)
	}
	StartServer(schema)

	// We create yet another Fields map, one which holds all the different queries

	// fields := graphql.Fields{
	// 	"gophers": &graphql.Field{
	// 		//It will return a list of GopherTypes, a list is an slice
	// 		//We defined our type in the schemas package earlier
	// 		Type: graphql.NewList(schema.GopherType),

	// 		Resolve:     gopherService.ResolveGophers,
	// 		Description: "Query all gophers",
	// 	},
	// }

	// Create the Root Query that is used to start each query
	// rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	// Now combine all Objects into a Schema Configuration
	// schemaConfig := graphql.SchemaConfig{
	// 	Query: graphql.NewObject(rootQuery)}
	//Create a new GraphQl schema
	// schema, err := graphql.NewSchema(schemaConfig)
	// if err != nil {
	// 	log.Fatalf("failed to create a new schema, errr : %v", err)
	// }
	//StartServer(&schema)
}

// StartServer will trigger the server with a Playground

func StartServer(schema *graphql.Schema) {

	//Create a new HTTP handler
	h := handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
