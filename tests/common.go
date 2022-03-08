package tests

import (
    "reflect"
    "testing"
)

func AssertEqual(t *testing.T, actual, expected interface{}) {
    if !reflect.DeepEqual(actual, expected) {
        t.Log("Assertion failed:\n[LOG] Expected: ", expected, "\n[LOG] Got: ", actual)
        t.FailNow()
    }
}
