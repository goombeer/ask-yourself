package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/goombeer/ask-yourself/backend/src/presentation/graphql/generated"
	"github.com/goombeer/ask-yourself/backend/src/presentation/graphql/resolvers"
)

// Ginを用いたアプリケーションでGraphQLのリクエストを捌けるよう、Handlerを生成する
type GraphQLRequestHandlerGenerator struct {
	resolver *resolvers.Resolver
}

func NewGraphQLRequestHandlerGenerator(resolver *resolvers.Resolver) *GraphQLRequestHandlerGenerator {
	return &GraphQLRequestHandlerGenerator{
		resolver: resolver,
	}
}

func (hg *GraphQLRequestHandlerGenerator) GraphQLHandler() gin.HandlerFunc {
	config := generated.Config{Resolvers: hg.resolver}

	server := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	server.SetErrorPresenter(CustomErrorPresenter)

	return func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	}
}

func (hg *GraphQLRequestHandlerGenerator) GraphQLPlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api/graphql/v1/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}