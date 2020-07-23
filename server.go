package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/idawud/go-graphql-crud/gql"
	handlers "github.com/idawud/go-graphql-crud/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "graphql-api ", log.LstdFlags)
	gql.InitDatabase()
	defer gql.DBConn.Close()
	// get schema
	schema, err := gql.GetSchema()
	if err != nil {
		panic("No Schema")
	}
	// get routes
	routes := handlers.Route{Logger: logger, Schema: schema}

	// setup mux
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", routes.IndexRoute)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.Use(handlers.JSONContentTypeMiddleware)
	postRouter.HandleFunc(`/graphql`, routes.GraphqlRoute)

		server := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout:120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout:1*time.Second,
	}

	fmt.Println("Server running on http://localhost:8080/" )
	log.Println(" Server started at ", time.Now().String())
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = server.Shutdown(ctx)
	
}
