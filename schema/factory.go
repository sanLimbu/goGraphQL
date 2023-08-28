package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/sanLimbu/gopheragency/gopher"
)

var jobType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Job",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"employeeID": &graphql.Field{
			Type: graphql.ID,
		},
		"company": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"start": &graphql.Field{
			Type: graphql.String,
		},
		"end": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// generateJobsField will build the GraphQL Field for jobs

func generateJobsField(gs *gopher.GopherService) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(jobType),
		Description: "A list of all jobs gopher have",
		Resolve:     gs.ResolveJobs,

		//Args are the possible arguments
		Args: graphql.FieldConfigArgument{
			"company": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
	}
}

// genereateGopherType will assemble the Gophertype and all related fields
func generateGopherType(gs *gopher.GopherService) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Gopher",
		// Fields is the field values to declare the structure of the object
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.ID,
				Description: "The ID that is used to identify unique gophers",
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
			// Here we create a graphql.Field which is depending on the jobs repository, notice how the Gopher struct does not contain any information about jobs
			// But this still works
			"jobs": generateJobsField(gs),
		}})
}

// GenerateSchema will create a GraphQL Schema and set the Resolvers found in the GopherService

func GenerateSchema(gs *gopher.GopherService) (*graphql.Schema, error) {
	gopherType := generateGopherType(gs)

	fields := graphql.Fields{
		"gophers": &graphql.Field{
			Type:        graphql.NewList(gopherType),
			Resolve:     gs.ResolveGophers,
			Description: "Query all Gophers",
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	//build RootMutation
	rootMutation := generateRootMutation(gs)

	//Now combine all objects into a schema config
	schemaCofig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: rootMutation,
	}
	//Create a new GraphQL schema
	schema, err := graphql.NewSchema(schemaCofig)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}

// generateGraphQLField is a generic builder factory to create graphql fields
func generateGraphQLField(output graphql.Output, resolver graphql.FieldResolveFn, description string, args graphql.FieldConfigArgument) *graphql.Field {

	return &graphql.Field{
		Type:        output,
		Resolve:     resolver,
		Description: description,
		Args:        args,
	}
}
