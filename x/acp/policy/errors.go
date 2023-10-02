package policy

import (
	"errors"
)

var ErrInvalidPolicy = errors.New("invalid policy")
var ErrUnknownMarshalingType = errors.New("unknown marshaling type")
var ErrMalformedGraph = errors.New("malformed management graph")
var ErrInvalidCreator = errors.New("invalid creator")
var ErrUnmarshaling = errors.New("unmarshaling error")
var ErrResourceMissingOwnerRelation = errors.New("resource missing owner relation")
