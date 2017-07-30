package testutils

import (
	"runtime"
	"testing"
)

type Assert struct {
	*testing.T
}

func (a *Assert) True(truth bool) {
	if !truth {
		_, file, line, _ := runtime.Caller(1)
		a.Errorf("%v:%v failed", file, line)
	}
}

func (a *Assert) Nil(o interface{}, t *testing.T) {
	if o != nil {
		_, file, line, _ := runtime.Caller(1)
		a.Errorf("%v:%v Not nil: %v", file, line, o)
	}
}
