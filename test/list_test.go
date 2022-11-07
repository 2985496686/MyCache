package test

import (
	"container/list"
	"fmt"
	"sort"
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
	var i bool
	fmt.Println(i)

	//arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	x := 1
	search := sort.Search(10, func(i int) bool {
		if i >= x {
			return true
		}
		return false
	})
	fmt.Printf("index: %d\n ", search)
}
