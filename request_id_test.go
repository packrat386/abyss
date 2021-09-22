package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithRequestID(t *testing.T) {
	called := false

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true

		rid := GetRequestID(r)

		if rid != "badfood" {
			t.Fatalf("got '%s' expected 'badfood'", rid)
		}
	})

	r := httptest.NewRequest("GET", "/who/cares", nil)

	wrapped := WithRequestID(h)
	wrapped.ServeHTTP(httptest.NewRecorder(), r)

	if !called {
		t.Fatal("never called the inside")
	}
}
