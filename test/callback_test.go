package test

import (
	"MyCache"
	"fmt"
	"testing"
)

func TestCallBack(t *testing.T) {

	var f MyCache.Getter = MyCache.GetterFunc(func(key string) ([]byte, error) {
		bytes := []byte(key)
		return bytes, nil
	})

	bytes, _ := f.Get("qwer")
	fmt.Println(bytes)
}
