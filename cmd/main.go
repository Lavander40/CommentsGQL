package main

import (
	"CommentsGQL"
	"CommentsGQL/resolver"
	"log"
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
  port := defaultPort

  srv := handler.NewDefaultServer(CommentsGQL.NewExecutableSchema(CommentsGQL.Config{Resolvers: &resolver.Resolver{}}))

  http.Handle("/", playground.Handler("GraphQL playground", "/query"))
  http.Handle("/query", srv)

  log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
