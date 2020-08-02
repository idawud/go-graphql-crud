package routes

import (
	"github.com/graphql-go/graphql"
	"github.com/idawud/go-graphql-crud/gql"
	"io/ioutil"
	"log"
	"net/http"
)

type Route struct {
	Logger *log.Logger
	Schema graphql.Schema
}

func (route *Route) IndexRoute( rw http.ResponseWriter, r *http.Request)  {
	route.Logger.Println("Index Get")
	_, _ = rw.Write([]byte("graphQl running on http://localhost:8080/graphql"))
}

func (route *Route) GraphqlRoute( rw http.ResponseWriter, r *http.Request)  {
	route.Logger.Println("Graphql post")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		route.Logger.Printf("could not read query %v", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// execute query
	json, err := gql.ExecuteQuery(route.Schema, string(body))
	if err != nil {
		route.Logger.Printf("could not process query %v", err)
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
