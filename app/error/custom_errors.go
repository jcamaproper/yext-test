package error

import "errors"

var (
	ErrInvalidJSON      = errors.New("invalid JSON format")
	ErrInvalidSortKeys  = errors.New("sortKeys must be a non-empty array")
	ErrInvalidPayload   = errors.New("payload must be a non-empty map")
	ErrInvalidArrayType = errors.New("array values must be of type string or number")
)
