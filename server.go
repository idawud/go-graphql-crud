package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)


func initDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "movies.db")
	if err != nil {
		panic("Failed To Open Connection: " + err.Error())
	}
	log.Println("DB connection Successful")

	DBConn.AutoMigrate(&Movie{})
	log.Println("Migration Completed")
}

func main(){
	initDatabase()
	defer DBConn.Close()
	// Setup schema
	schemaConfig := graphql.SchemaConfig{Query: rootQuery, Mutation: mutationType}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create graphql schema, error %v", err)
	}

	// create Movie

	mutate := `
	 mutation {
        updateMovie(id: 2, movie: { title: "updated movie again"}) {
            title,
			id
        }
    }
	`
	executeQuery(schema, mutate)
	query := `
	 query {
        movie(id: 2){
            title,
			minutes,
			id
        }
    }
	`
	executeQuery(schema, query)
}


func executeQuery(schema graphql.Schema, query string) {
	params := graphql.Params{Schema: schema, RequestString: query}
	response := graphql.Do(params)

	if len(response.Errors) > 0 {
		log.Fatalf("Unable to process graphql query, err %+v", response.Errors)
	}

	data, _ := json.Marshal(response)
	fmt.Printf("%s \n", data)
}

