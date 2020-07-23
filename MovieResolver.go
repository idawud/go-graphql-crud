package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

var MovieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "movie",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"minutes": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)


var MovieInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "MovieInputType",
	Fields: graphql.InputObjectConfigFieldMap{
	"title": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	"minutes": &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
})

// Mutations
var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createMovie": &graphql.Field{
			Type: MovieType,
			Description: "Create a new Movie",
			Args: graphql.FieldConfigArgument{
				"movie": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(MovieInputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				movieInput := params.Args["movie"]

				movie := Movie{}
				err := mapstructure.Decode(movieInput, &movie)
				if err != nil {
					return nil, err
				}

				// add to db
				DBConn.Create(&movie)
				if movie.ID == 0 {
					return nil, nil
				}
				return movie, nil
			},
		},
		"deleteMovie": &graphql.Field{
			Type: graphql.Boolean,
			Description: "delete a Movie",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"]
				if ok {
					var movie Movie
					DBConn.First(&movie, id)
					if movie.ID == 0 {
						return false, nil
					}
					DBConn.Delete(&movie)
					return true, nil
				}
				return false, fmt.Errorf("invalid ID %v", id)
			},
		},
		"updateMovie": &graphql.Field{
			Type: MovieType,
			Description: "update a Movie",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"movie": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(MovieInputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				movieInput := params.Args["movie"]
				id  := params.Args["id"]

				newMovie := Movie{}
				err := mapstructure.Decode(movieInput, &newMovie)
				if err != nil {
					return nil, err
				}

				// update db records
				var oldMovie Movie
				DBConn.First(&oldMovie, id)
				if oldMovie.ID == 0 {
					return nil, nil
				}

				oldMovie.Title = newMovie.Title
				if newMovie.Minutes != 0  {
					oldMovie.Minutes = newMovie.Minutes
				}

				DBConn.Save(&oldMovie)
				return oldMovie, nil
			},
		},
	},
})

// Queries
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "rootQuery",
	Fields: graphql.Fields{
		"movies": &graphql.Field{
			Type:   graphql.NewList(MovieType),
			Description: "Get all Movies",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var movies []Movie
				DBConn.Find(&movies)
				return movies, nil
			},
		},
		"movie": &graphql.Field{
			Type:   MovieType,
			Description: "Get all Movie by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"]
				if ok {
					var movie Movie
					DBConn.First(&movie, id)
					if movie.ID == 0 {
						return nil, nil
					}
					return movie, nil
				}
				return nil, fmt.Errorf("invalid ID %v", id)
			},
		},
	},
})
