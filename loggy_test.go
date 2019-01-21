package loggy

import (
	"testing"
)

func Test_Logger_f(t *testing.T) {
	l := New("test")
	l.Tracef("this is trace %d", 1)
	l.Infof("this is info %d", 2)
	l.Warnf("this is warn %d", 3)
	l.Errorf("this is error %d", 4)
}

func Test_Logger_Discard(t *testing.T) {
	l := Discard()
	l.Tracef("do not see this trace")
	l.Infof("or this info")
	l.Warnf("or this warn")
	l.Errorf("or this error")
}
