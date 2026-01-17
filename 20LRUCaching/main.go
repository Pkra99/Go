package main

import (
	"sync"
)

type Node struct {
	key   string
	value interface{}
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache map[string]*Node
	head *Node
	tail *Node
	mu *sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		panic("Capcity must be positive")
	}

	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache: make(map[string]*Node),
		head: head,
		tail: tail,
	}
}

func (lru *LRUCache) Get(key string) (interface{}, bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()
	if node, exists := lru.cache[key]; exists {
		lru.moveToHead(node)
		return node.value, true
	}
	return nil, false
}

func(lru *LRUCache) Put(key string, value interface{}) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if node, exists := lru.cache[key]; exists {
		node.value = value
		lru.moveToHead(node)
		return
	}

	newNode := &Node{
		key: key,
		value: value,
	}

	if len(lru.cache) >= lru.capacity {
		lru.evictTail()
	}
	lru.cache[key] = newNode
	lru.addToHead(newNode)
} 

func (lru *LRUCache) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.moveToHead(node)
}
func (lru *LRUCache) addToHead(node *Node) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache) evictTail() {
	lastNode := lru.tail.prev
	lru.removeNode(lastNode)
	delete(lru.cache, lastNode.key)
}

func (lru *LRUCache) size() int {
	lru.mu.RLock()
	defer lru.mu.Unlock()
	return len(lru.cache)
}

func (lru *LRUCache) Keys() []string {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	keys := make([]string, 0, len(lru.cache))

	for key := range lru.cache {
		keys = append(keys, key)
	}
	return keys
}



// generic implemenataion for type safety

// type LRUCache[K comparable, V any] struct {
//     capacity int
//     cache    map[K]*Node[K, V]
//     head     *Node[K, V]
//     tail     *Node[K, V]
//     mu       sync.RWMutex
// }
// type Node[K comparable, V any] struct {
//     key   K
//     value V
//     prev  *Node[K, V]
//     next  *Node[K, V]
// }
// // Usage with type safety
// cache := NewLRUCache[string, User](1000)
// user, exists := cache.Get("user:123")  // Returns User, not interface{}

// with the expiration node

// type Node[K comparable, V any] struct {
// 	key K
// 	value V
// 	prev *Node[K, V]
// 	next *Node[K, V]
// 	expiresAt time.Time
// }

// func (lru *LRUCache[K, V]) Get(key K) (V,bool) {
// 	lru.mu.Lock()
// 	defer lru.mu.Unlock()

// 	if node, exists := lru.cache[key]; exists {
// 		if time.Now().After(node.expiresAt) {
// 			lru.removeNode(node)
// 			delete(lru.cache, key)
// 			var zero V
// 			return  zero, false
// 		}
// 		lru.moveToHead(node)
// 		reutrn node.value, true	
// 	}
// 	var zero V
// 	return zero, false
// }