package support

type ErrorCode = string

const (
	// general errors
	BadRequest          = ErrorCode("bad_request")
	OutdatedVersion     = ErrorCode("bad_request.version_outdated")
	Unauthorized        = ErrorCode("unauthorized")
	Forbidden           = ErrorCode("forbidden")
	NotFound            = ErrorCode("not_found")
	InternalServerError = ErrorCode("internal_server_error")
)