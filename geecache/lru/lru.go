package lru

import (
	"container/list"
)

type Value interface {
	Len() int
}

type Cache struct {
	maxBytes	int64
	nBytes		int64
	ll 			*list.List
	cache 		map[string]*list.Element
	OnEvicted 	func(key string, value Value)
}

type entry struct{
	key string
	value Value
}

func New(maxBytes int64, onEvicted func(key string, value Value)) *Cache{
	return &Cache {
		ll: list.New(),
		cache: make(map[string]*list.Element),
		maxBytes: maxBytes,
		OnEvicted: onEvicted,
	}
}

func (c *Cache)Get(key string)(value Value, ok bool){
	if  element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		return kv.value, true
	}

	return 
}


func (c *Cache)moveOldest(){
	oldest := c.ll.Back()
	if oldest != nil {
		c.ll.Remove(oldest)
		kv := oldest.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache)Add(key string, value Value){
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	}else {
		newEle := c.ll.PushFront(&entry{key:key, value: value})
		c.cache[key] = newEle
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.nBytes > c.maxBytes {
		c.moveOldest()
	}
}

func (c *Cache)Len() int {
	return c.ll.Len()
}