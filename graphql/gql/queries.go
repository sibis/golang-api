package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/sibis/golang-api/graphql/postgres"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *postgres.Db) *Root {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"product": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: Product,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: resolver.ProductResolver,
					},
					"products": &graphql.Field{
						Type:    graphql.NewList(Product),
						Resolve: resolver.ProductsResolver,
					},
				},
			},
		),
	}
	return &root
}

func NewMutation(db *postgres.Db) *Root {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"createProduct": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: Product,
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"price": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Float),
							},
						},
						Resolve: resolver.CreateProductResolver,
					},
				},
			},
		),
	}
	return &root
}
