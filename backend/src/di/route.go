package di

import (
	"github.com/gin-gonic/gin"
	"github.com/goombeer/ask-yourself/src/config"
)

func NewGinEngine(
	conf config.Config,
	gqlHandlerGenerator *graphql.GraphQLRequestHandlerGenerator,
) *gin.Engine {
	r := gin.New()

	// Routing
	api := r.Group("/api")
	{

		graphqlGroup := api.Group("/graphql/v1")
		{
			graphqlGroup.Use(gql_helper.GinContextToContextMiddleware())
			graphqlGroup.POST("/query", gqlHandlerGenerator.GraphQLHandler())
			if conf.Server.Environment.AllowGraqhQLPlayGround() {
				graphqlGroup.GET("/playground", gqlHandlerGenerator.GraphQLPlaygroundHandler())
			}
		}
	}
	return r
}