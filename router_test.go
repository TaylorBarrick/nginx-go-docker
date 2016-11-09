package main

import (
	"net/http/httptest"
	"testing"
)

func TestGetURIParams(t *testing.T) {
	var expected = []string{
		"blah",
		"thing",
		"stuff",
	}
	r := httptest.NewRequest("GET", "/blah/thing/stuff", nil)
	values := getURIParams(r)
	for i, e := range expected {
		if e != values[i] {
			t.Errorf("Expected %v index param to be %s and was %s", i, e, values[i])
		}
	}
}
