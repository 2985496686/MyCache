package test

import (
	"MyCache"
	"fmt"
	"log"
	"testing"
)

var db = map[string]string{
	"jok": "545",
	"klo": "323",
	"los": "232",
}

func TestGroup(t *testing.T) {
	cache := MyCache.NewGroup("cach1", 2<<10, MyCache.GetterFunc(func(key string) ([]byte, error) {
		log.Printf("%s search in db\n", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s is not find in db ", key)
	}))

	cache.Get("jok")
	cache.Get("klo")
	cache.Get("klo")
	cache.Get("jok")
	cache.Get("los")
	cache.Get("los")
}
