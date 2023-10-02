package testutil

import (
	"reflect"
	"strings"
	"testing"
)

// TestSuite represents a type which groups several tests.
// Implementing types should have internal Member functions whose name start with "Test"
// in order for the Runner to pick it up.
// See RunSuite for usage example
type TestSuite interface{}

// Run a TestSuite
// The suite type must have methods prefixed with "Test" and these methods must receive a
// *testing.T argument
func RunSuite(t *testing.T, suite TestSuite) {
	suiteVal := reflect.ValueOf(suite)
	suiteT := reflect.TypeOf(suite)

	var tests []reflect.Method
	for i := 0; i < suiteT.NumMethod(); i++ {
		method := suiteT.Method(i)
		if strings.HasPrefix(method.Name, "Test") {
			tests = append(tests, method)
		}
	}

	for _, test := range tests {
		f := test.Func
		t.Run(test.Name, func(t *testing.T) {
			tVal := reflect.ValueOf(t)
			args := []reflect.Value{suiteVal, tVal}
			f.Call(args)
		})
	}
}
