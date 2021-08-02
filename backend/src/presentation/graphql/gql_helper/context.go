// See https://gqlgen.com/recipes/gin/#accessing-gincontext

package gql_helper

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type ginContextKeyType struct{}

var GinContextKey = ginContextKeyType{}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey)
	if ginContext == nil {
		err := errors.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := errors.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}