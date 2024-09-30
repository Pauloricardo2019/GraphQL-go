package main

import (
	"github.com/Pauloricardo2019/graphql-teste/adapter/graphql"
	"log"
	"net/http"
)

func main() {
	// Configura o handler GraphQL
	handler := graphql.NewGraphQLHandler()

	// Inicia o servidor HTTP
	http.Handle("/graphql", handler)

	log.Println("Servidor rodando em http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
