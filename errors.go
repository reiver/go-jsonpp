package jsonpp

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errInternalError = erorr.Error("jsonpp: internal error")
	errNilWriter     = erorr.Error("jsonpp: nil writer")
)
