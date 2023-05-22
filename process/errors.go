package process

import "errors"

var errInvalidOperationType = errors.New("invalid/unknown operation type")

var errNilMarshaller = errors.New("nil marshaller provided")
