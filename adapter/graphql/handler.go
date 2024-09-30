package graphql

import (
	"github.com/Pauloricardo2019/graphql-teste/internal/service"
	"github.com/graphql-go/handler"
	"log"
)

func NewGraphQLHandler() *handler.Handler {
	schema, err := NewSchema(service.NewService())
	if err != nil {
		log.Fatalf("Falha ao criar o schema GraphQL: %v", err)
	}
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // Habilita o GraphiQL para facilitar os testes
	})
}
