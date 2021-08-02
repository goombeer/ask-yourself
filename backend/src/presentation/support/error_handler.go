package support

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ResolveErrorCode(err error) (int, ErrorCode, string) {
	var causeErr error
	// gqlerror uses Unwrap instead of Cause
	if gqlErr, ok := err.(*gqlerror.Error); ok {
		causeErr = errors.Cause(gqlErr.Unwrap())
		// FIXME hotfix; handle gqlError appropriately
		if causeErr == nil {
			causeErr = gqlErr.Unwrap()
		}
	} else {
		causeErr = errors.Cause(err)
	}

	switch causeErr.(type) {
	default:
		return 500, InternalServerError, causeErr.Error()
	}
}
