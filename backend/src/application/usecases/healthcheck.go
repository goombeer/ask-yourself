package usecases

import (
	"github.com/goombeer/ask-yourself/backend/src/config"
	"github.com/goombeer/ask-yourself/backend/src/domain/healthcheck"
)

type HealthcheckUsecase struct {
	serverConfig config.ServerConfig
}

func NewHealthcheckUsecase(conf config.Config) HealthcheckUsecase {
	return HealthcheckUsecase{
		serverConfig: conf.Server,
	}
}

func (usecase *HealthcheckUsecase) Execute() *healthcheck.HealthcheckResult {
	return &healthcheck.HealthcheckResult{
		IsOK:   true,
		Server: usecase.serverConfig.Environment.String(),
	}
}