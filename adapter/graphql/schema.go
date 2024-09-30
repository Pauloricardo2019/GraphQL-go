package graphql

import (
	"fmt"
	serviceIF "github.com/Pauloricardo2019/graphql-teste/ports/service"
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// Define as queries
// NewSchema cria um novo schema GraphQL com a injeção de dependência do userService
func NewSchema(userService serviceIF.UserServiceIF) (graphql.Schema, error) {
	// Defina o QueryType com o userService injetado
	var queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if !ok {
						return nil, nil
					}
					// Usa o userService injetado
					return userService.GetUser(id)
				},
			},
		},
	})

	// Defina o MutationType com o userService injetado
	var mutationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					name := p.Args["name"].(string)
					age := p.Args["age"].(int)

					// Usa o userService injetado
					return userService.CreateUser(id, name, age)
				},
			},
			"updateUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					name := p.Args["name"].(string)
					age := p.Args["age"].(int)

					// Usa o userService injetado
					return userService.UpdateUser(id, name, age)
				},
			},
			"deleteUser": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "DeleteResponse",
					Fields: graphql.Fields{
						"message": &graphql.Field{
							Type: graphql.String,
						},
					},
				}),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)

					// Usa o userService injetado
					err := userService.DeleteUser(id)
					if err != nil {
						return nil, err
					}

					return map[string]interface{}{
						"message": fmt.Sprintf("User %s deleted successfully", id),
					}, nil
				},
			},
		},
	})

	// Cria o schema GraphQL com as queries e mutations definidas
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
