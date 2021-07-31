package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/goombeer/ask-yourself/src/application/useecases"
	"github.com/goombeer/ask-yourself/src/config"
)
func InitAppServerDependency(cfg config.Config) (*gin.Engine, func()) {
	wire.Build(
		//DBSet,
		ApplicationSet,
		//GraphqlSet,
		NewGinEngine,
	)
	return nil, nil
}

func InitConfig() config.Config {
	wire.Build(loadConfig)
	return config.Config{}
}

var ApplicationSet = wire.NewSet(
	useecases.NewHealthcheckUsecase,
)

//var GraphqlSet = wire.NewSet(
//	resolvers.NewResolver,
//	graphql.NewGraphQLRequestHandlerGenerator,
//)