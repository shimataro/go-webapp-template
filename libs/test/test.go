package test

import (
	"reflect"
	"testing"
)

func Expect(t *testing.T, expected interface{}, actual interface{}, err error) {
	if err != nil {
		t.Errorf("Got error: %v", err.Error())
	}
	if expected != actual {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)",
			expected, reflect.TypeOf(expected),
			actual, reflect.TypeOf(actual))
	}
}
