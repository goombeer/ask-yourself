package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/goombeer/ask-yourself/backend/src/domain/healthcheck"
	"github.com/goombeer/ask-yourself/backend/src/presentation/graphql/generated"
	"github.com/goombeer/ask-yourself/backend/src/presentation/graphql/model"
)

func (r *healthcheckResultResolver) Server(ctx context.Context, obj *healthcheck.HealthcheckResult) (model.ServerEnvType, error) {
	return model.ServerEnvType(obj.Server), nil
}

func (r *queryResolver) Healthcheck(ctx context.Context) (*healthcheck.HealthcheckResult, error) {
	return r.getHealthcheckResult.Execute(), nil
}

// HealthcheckResult returns generated.HealthcheckResultResolver implementation.
func (r *Resolver) HealthcheckResult() generated.HealthcheckResultResolver {
	return &healthcheckResultResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type healthcheckResultResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
