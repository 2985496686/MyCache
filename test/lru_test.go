package test

import (
	"MyCache/lru"
	"fmt"
	"testing"
)

type value struct {
	name string
}

func (v value) Len() uint64 {
	return uint64(len(v.name))
}

var m = lru.NewCache(uint64(10), nil)

func TestAddAndGet(t *testing.T) {
	m.Add("1", value{"ab"})
	m.Add("2", value{"cd"})
	m.Add("3", value{"ab"})
	m.Add("1", value{"cd"})
	m.Add("5", value{"ab"})
	m.Add("6", value{"cd"})
	v, ok := m.Get("1")
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("fail find ‘1’ belong to value")
	}
	v, ok = m.Get("5")
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("fail find ‘5’ belong to value")
	}
	v, ok = m.Get("6")
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("fail find ‘6’ belong to value")
	}
	fmt.Printf("len:%v", m.Len())
}

func TestRemoveOldest(t *testing.T) {

}

func TestGet(t *testing.T) {

}
