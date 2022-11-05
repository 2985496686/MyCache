package test

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	ll := list.New()
	ll.PushFront(1)
	ll.PushFront(2)
	ll.PushFront(3)
	fmt.Println(ll.Back().Value)
	fmt.Println(ll.Back().Value)
	fmt.Println(ll.Back().Value)

	m := map[string]string{"1": "www"}
	delete(m, "1")
	s := m["1"]
	fmt.Println(s)

}
