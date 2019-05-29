package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/underthebus/lazy-invoice/backend/graphql"
	"github.com/underthebus/lazy-invoice/backend/store"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	store := store.NewInMemoryStore()
	router.Handle(
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
	//log.Fatal(http.ListenAndServe(":"+port, nil))

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
