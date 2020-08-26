package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sibis/golang-api/graphql/gql"
	"github.com/sibis/golang-api/graphql/postgres"
	"github.com/sibis/golang-api/graphql/server"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("🚀 server started ... ")
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString("localhost", "5432", "sibi", "postgres"),
	)
	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)
	rootMutation := gql.NewMutation(db)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery.Query,
			Mutation: rootMutation.Query,
		},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{GqlSchema: &schema}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQL())
	return router, db
}
