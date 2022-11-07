package test

import (
	"MyCache"
	"fmt"
	"log"
	"net/http"
	"testing"
)

var db2 = map[string]string{
	"jok": "545",
	"klo": "323",
	"los": "232",
}

func TestHttpServer(t *testing.T) {
	h := MyCache.NewHTTPPool("localhost:8080")

	MyCache.NewGroup("cach1", 2<<10, MyCache.GetterFunc(func(key string) ([]byte, error) {
		log.Printf("%s search in db\n", key)
		if v, ok := db2[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s is not find in db ", key)
	}))
	pool := MyCache.NewHTTPPool("localhost:8080")
	http.ListenAndServe("localhost:8080", pool)

	log.Fatal(http.ListenAndServe("localhost:8080", h))

}
