package main

import (
	"CommentsGQL"
	"CommentsGQL/resolver"
	"CommentsGQL/storage"
	"CommentsGQL/storage/memory"
	"flag"
	"log"
	"net/http"
	"CommentsGQL/storage/postgre"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"
const defaultConn = "postgres://admin:password@localhost:5432/posts?sslmode=disable"

func main() {
	port := flag.String("port", defaultPort, "gql server port")
	storageType := flag.String("storage", "memo", "type of storage (memo/psql)")
	connectionString := flag.String("conn", defaultConn, "Postgres connection string")
	flag.Parse()

	var st storage.Storage
    switch *storageType {
    case "psql":
        st = postgre.New(*connectionString)
    case "memo":
        st = memory.New()
    default:
        log.Fatal("no such storage supported")
    }

	resolver := &resolver.Resolver{
		Storage: st,
	}

	srv := handler.NewDefaultServer(CommentsGQL.NewExecutableSchema(CommentsGQL.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}
