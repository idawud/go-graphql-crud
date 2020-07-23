package main

import (
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"log"
	"net/http"
)

type Route struct {
	logger *log.Logger
	schema graphql.Schema
}

func (route *Route) IndexRoute( rw http.ResponseWriter, r *http.Request)  {
	route.logger.Println("Index Get")
	_, _ = rw.Write([]byte("graphQl running on http://localhost:8080/graphql"))
}

func (route *Route) GraphqlRoute( rw http.ResponseWriter, r *http.Request)  {
	route.logger.Println("Graphql post")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		route.logger.Printf("could not read query %v", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// execute query
	json, err := executeQuery(route.schema, string(body))
	if err != nil {
		route.logger.Printf("could not process query %v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = rw.Write(json)
}

/// Middleware
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
