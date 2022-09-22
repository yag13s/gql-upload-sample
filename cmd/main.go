package main

import (
	"fmt"

	"github.com/ShinsakuYagi/gql-upload-sample/graphql"
	"github.com/ShinsakuYagi/gql-upload-sample/graphql/generated"
	"github.com/ShinsakuYagi/gql-upload-sample/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func main() {
	loadEnv()
	// Setting up Gin
	r := gin.Default()
	fmt.Println("server version :", model.Version)

	r = defineRoutes(r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func defineRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/query", graphqlHandler())
	return r
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graphql.Resolver{},
			},
		),
	)

	// Post only
	//
	// Since we are not using Playground, Get is not necessary.
	h.AddTransport(transport.POST{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
