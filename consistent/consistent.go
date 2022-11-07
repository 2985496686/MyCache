package consistent

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func([]byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[uint32]string
}

func NewMap(replicas int, hash Hash) *Map {
	m := Map{
		hash:     hash,
		replicas: replicas,
		hashMap:  make(map[uint32]string),
	}
	if hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return &m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := m.hash([]byte(strconv.Itoa(i) + key))
			m.keys = append(m.keys, int(hash))
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	search := sort.Search(len(m.keys), func(i int) bool {
		if m.keys[i] >= hash {
			return true
		}
		return false
	})
	return m.hashMap[uint32(m.keys[search%len(m.keys)])]
}
