package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"repath.io/config"
	"repath.io/graph"
	"repath.io/graph/generated"

	// "repath.io/data"
)

const defaultPort = "8080"

func main() {
	configuration := config.LoadConfiguration()
	fmt.Printf("starting %s service\n", configuration.ServiceName)
	// datastore := data.NewPersonDataStore(configuration.Neo4j.URI, configuration.Neo4j.Username, configuration.Neo4j.Password)
	//defer datastore.Close()

	// initialize API and pass datastore
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
