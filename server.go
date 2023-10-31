package main

import (
	"context"
	"fmt"
	"graph/database"
	"graph/graph"
	"graph/repository"
	"graph/service"
	"net/http"
	"time"

	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/zerolog/log"
)

const defaultPort = "8081"

func main() {

	svc, err := StartApp()
	if err != nil {
		log.Info().Err(err).Msg("error in startapp")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Svc: svc,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal().Err(http.ListenAndServe(":"+port, nil))
}
func StartApp() (service.UserService, error) {

	// starting database connection
	db, err := database.Open()
	if err != nil {
		return &service.Service{}, fmt.Errorf("error in opening the database connection : %w", err)
	}

	pg, err := db.DB()
	if err != nil {
		return &service.Service{}, fmt.Errorf("error in getting the database instance")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return &service.Service{}, fmt.Errorf("database is not connected: %w", err)
	}
	repo, err := repository.NewRepo(db)
	if err != nil {
		return &service.Service{}, fmt.Errorf("repository not initialized: %w", err)
	}
	svc, err := service.NewService(repo)
	if err != nil {
		return &service.Service{}, fmt.Errorf("service layer not initialized: %w", err)
	}

	return svc, nil
}
