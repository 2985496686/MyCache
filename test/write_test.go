package test

import (
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	http.Handle("/http", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	}))
	http.ListenAndServe("localhost:8080", nil)
}
