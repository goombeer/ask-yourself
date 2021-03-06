package di

import (
	"github.com/gin-gonic/gin"
	"github.com/goombeer/ask-yourself/backend/src/config"
	"github.com/goombeer/ask-yourself/backend/src/presentation/graphql"
	"github.com/goombeer/ask-yourself/backend/src/presentation/graphql/gql_helper"
	"github.com/goombeer/ask-yourself/backend/src/presentation/support"
)

func NewGinEngine(
	conf config.Config,
	gqlHandlerGenerator *graphql.GraphQLRequestHandlerGenerator,
) *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(support.GinLogFormatter))
	r.Use(support.LogInfoMiddleware)
	r.Use(support.Logger(), gin.Recovery())

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