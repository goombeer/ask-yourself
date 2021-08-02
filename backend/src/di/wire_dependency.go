//+build wireinject

//execute `wire` in this package
package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/goombeer/ask-yourself/backend/src/application/usecases"
	"github.com/goombeer/ask-yourself/backend/src/config"
	"github.com/goombeer/ask-yourself/backend/src/presentation/graphql"
	"github.com/goombeer/ask-yourself/backend/src/presentation/graphql/resolvers"
)

func InitAppServerDependency(cfg config.Config) (*gin.Engine, func()) {
	wire.Build(
		ApplicationSet,
		GraphqlSet,
		NewGinEngine,
	)
	return nil, nil
}

func InitConfig() config.Config {
	wire.Build(loadConfig)
	return config.Config{}
}

var ApplicationSet = wire.NewSet(
	usecases.NewHealthcheckUsecase,
)

var GraphqlSet = wire.NewSet(
	resolvers.NewResolver,
	graphql.NewGraphQLRequestHandlerGenerator,
)