package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/underthebus/lazy-invoice/backend/graphql"
	"github.com/underthebus/lazy-invoice/backend/store"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	store := store.NewInMemoryStore()
	http.Handle(
		"/query",
		handler.GraphQL(
			graphql.NewExecutableSchema(
				graphql.Config{
					Resolvers: graphql.NewResolver(store),
				},
			),
		),
	)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
