package testutils

import (
	"runtime"
	"testing"
)

type Assert struct {
	*testing.T
}

func (a *Assert) Equal(x interface{}, y interface{}) {
	if x != y {
		_, file, line, _ := runtime.Caller(1)
		a.Errorf("%v:%v %+v != %+v", file, line, x, y)
	}
}

func (a *Assert) Nil(o interface{}, t *testing.T) {
	if o != nil {
		_, file, line, _ := runtime.Caller(1)
		a.Errorf("%v:%v Not nil: %v", file, line, o)
	}
}
