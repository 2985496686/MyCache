package test

import (
	"MyCache/lru"
	"sync"
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

}

func TestRemoveOldest(t *testing.T) {

}

func TestGet(t *testing.T) {
	var rw sync.RWMutex
	rw.Lock()
	rw.Unlock()
}
