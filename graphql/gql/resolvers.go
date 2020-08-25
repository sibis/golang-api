package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/sibis/golang-api/graphql/postgres"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *postgres.Db
}

// UserResolver resolves our user query through a db call to GetUserByName
func (r *Resolver) ProductResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	id, ok := p.Args["id"].(int)
	if ok {
		product := r.db.GetProductByID(id)
		return product, nil
	}

	return nil, nil
}
