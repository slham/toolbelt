package cache

import (
	"errors"
)

// LRUCache consists of
// a map and a doubly linked list
type LRUCache struct {
	Cache

	Capacity int
	Map      map[string]*LRUCacheNode
	Head     *LRUCacheNode
	Tail     *LRUCacheNode
}

// LRUCacheNode represents cached data
type LRUCacheNode struct {
	Key   string
	Value interface{}
	Prev  *LRUCacheNode
	Next  *LRUCacheNode
}

// Creates new LRUCache with a given capacity above 0.
func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity <= 0 {
		return nil, errors.New("invalid capacity. capacity should be greater than 0.")
	}

	return &LRUCache{
		Capacity: capacity,
		Map:      make(map[string]*LRUCacheNode, capacity),
		Head:     nil,
		Tail:     nil,
	}, nil
}

// Place a value in the cache
func (c *LRUCache) Set(key string, value interface{}) {
	c.remove(key)
	c.prepend(key, value)

	if len(c.Map) > c.Capacity {
		c.remove(c.Tail.Key)
	}
}

// Returns the cached value associated with the key if present
func (c *LRUCache) Get(key string) interface{} {
	node := c.Map[key]
	if node == nil {
		return nil
	}

	c.remove(key)
	c.prepend(key, node.Value)
	return node.Value
}

// Adds value to cache
func (c *LRUCache) prepend(key string, value interface{}) {
	node := &LRUCacheNode{
		Key:   key,
		Value: value,
		Next:  c.Head,
		Prev:  nil,
	}

	if c.Head != nil {
		c.Head.Prev = node
	}
	if c.Tail == nil {
		c.Tail = node
	}

	c.Head = node
	c.Map[key] = node
}

// Removes value from cache
func (c *LRUCache) remove(key string) {
	node := c.Map[key]
	if node == nil {
		return
	}

	if key == c.Head.Key {
		c.Head = c.Head.Next
	}

	if key == c.Tail.Key {
		c.Tail = c.Tail.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	delete(c.Map, key)
}
