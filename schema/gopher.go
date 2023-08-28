package schema

import "github.com/graphql-go/graphql"

var GopherType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Gopher",
	// Fields is the field values to declare the structure of the object

	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.ID,
			Description: "The ID is used to identify unique gopher",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the gopher",
		},
		"hired": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "True if the Gopher is employeed",
		},
		"profession": &graphql.Field{
			Type:        graphql.String,
			Description: "The gophers last/current profession",
		},
	},
})
