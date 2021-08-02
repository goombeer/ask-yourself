package graphql

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/goombeer/ask-yourself/backend/src/presentation/support"
)

func CustomErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	status, code, msg := support.ResolveErrorCode(err)
	path := graphql.GetFieldContext(ctx).Path()
	if status >= 500 {
		logrus.Error(err.Error())
	}
	return &gqlerror.Error{
		Message: msg,
		Path:    path,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}