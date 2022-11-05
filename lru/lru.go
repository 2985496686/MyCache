package lru

import (
	"container/list"
	"errors"
)

type Cache struct {
	//最大占用内存
	maxBytes uint64
	//当前已使用内存内存
	nbytes uint64
	//用于查询节点的map
	cache map[string]*list.Element
	//存储数据的双向循环队列
	ll *list.List
	//删除数据的回调函数
	onEvicted func(key string, value Value)
}

type Value interface {
	Len() uint64
}

type Entity struct {
	key   string
	value Value
}

func NewCache(maxBytes uint64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		cache:     map[string]*list.Element{},
		ll:        list.New(),
		onEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (v Value, ok bool) {
	ele := c.cache[key]
	if ele != nil {
		c.ll.MoveToFront(ele)
		v = ele.Value.(*Entity).value
		ok = true
	}
	return
}

func (c *Cache) removeOldest() {
	oldest := c.ll.Remove(c.ll.Back()).(*Entity)
	delete(c.cache, oldest.key)
	if oldest != nil {
		c.nbytes = c.nbytes - uint64(len(oldest.key)) - oldest.value.Len()
		if c.onEvicted != nil {
			c.onEvicted(oldest.key, oldest.value)
		}
	}
}

func (c *Cache) Remove(key string) {
	ele, ok := c.cache[key]
	if ok {
		entity := c.ll.Remove(ele).(*Entity)
		delete(c.cache, key)
		c.nbytes -= uint64(len(key)) + entity.value.Len()
	}
}

func (c *Cache) Add(key string, v Value) (err error, ok bool) {
	c.Remove(key)
	levBytes := c.maxBytes - c.nbytes
	vBytes := uint64(len(key)) + v.Len()
	if vBytes > c.maxBytes {
		err = errors.New("entity is too big")
		return
	} else if vBytes > levBytes {
		c.removeOldest()
		return c.Add(key, v)
	}
	c.nbytes += vBytes
	ele := c.ll.PushFront(&Entity{key, v})
	c.cache[key] = ele
	ok = true
	return
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
