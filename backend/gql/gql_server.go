package gql

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)


func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "movies.db")
	if err != nil {
		panic("Failed To Open Connection: " + err.Error())
	}
	log.Println("DB connection Successful")

	DBConn.AutoMigrate(&Movie{})
	log.Println("Migration Completed")
}

func GetSchema() (graphql.Schema, error){
	// Setup schema
	schemaConfig := graphql.SchemaConfig{Query: rootQuery, Mutation: mutationType}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return schema, fmt.Errorf("failed to create graphql schema, error %v", err)
	}
	return schema, nil
}


func ExecuteQuery(schema graphql.Schema, query string) ([]byte, error){
	params := graphql.Params{Schema: schema, RequestString: query}
	response := graphql.Do(params)

	if len(response.Errors) > 0 {
		return nil, fmt.Errorf("unable to process graphql query, err %+v", response.Errors)
	}

	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return data, nil
}

