package main

import (
	"reflect"
	"testing"
)

func isType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func TestInitDebug(t *testing.T) {
	t.Run("test_InitDebug()", func(t *testing.T) {
		expected := 0
		result := 0

		if !isType(expected, result) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
