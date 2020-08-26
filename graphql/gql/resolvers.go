package gql

import (
	"fmt"

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
		fmt.Println(product)
		return product, nil
	}

	return nil, nil
}

func (r *Resolver) ProductsResolver(p graphql.ResolveParams) (interface{}, error) {
	product := r.db.GetProducts()
	if len(product) > 0 {
		return product, nil
	}
	return nil, nil
}

func (r *Resolver) CreateProductResolver(p graphql.ResolveParams) (interface{}, error) {
	name, _ := p.Args["name"].(string)
	price, _ := p.Args["price"].(float64)

	product := r.db.CreateProduct(name, price)
	return product, nil
}
