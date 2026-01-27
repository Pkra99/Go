package main

import (
	"sync"
)

/*
1. create a class/struct Node, which will consist the key, value, prev, next properties
2. create a class/struct LRUCache, which will store the capacity, cache, head, tail (if using go then mutex)
3. Initialize the LRUCache, first check for edge cases like the capicity must be positive -> initilize the head and tail with emtpy Node
assign head.next to tail head.next = tail, and tail.prev to head tail.prev = head, and return the LRUCache with capcity, cache (which is just map of string), head and tail
4. Create Get func which takes the key as a parameter and will return either node.value or boolean value 
	4.1 Acquire the lcok for the data consistency and unlock at the end of the OP
	4.2 check if nodes exists or not, if not exits return nil, and flase, and if exits -> get the cache based on the key lru.cache[key] 
	4.2 move the current cache to the head -> return the node.value or nil
5.  Create Put func which takes the key, and value as a parameter, and returns nil
	5.1 Acquire lock and unlock at the end
	5.2 4.2 check if node exits or not, if exits -> get the cache based on the key lru.cache[key] 
	5.2 move the current cache to the head -> return the node.value or nil
	5.3 if not exists create a newNode with the key and value
	5.4 compare the lru cache size with the capacity len(lru.cache) >= lru.capacity, if tre evict the trail (remove the least used value to add the new one)
	5.5 set the lru.cache[key] = newNode -> update the head (movetohead(newNode))
6. Creaet addtohead func which takes the node as a parameter
	6.1 node.prev = lru.head
	6.2 node.next = lru.head.next
	6.3 lru.head.next.prev = node
	6.4	lru.head.prev = node 
7 Create removeNde func with a node parameter
	7.1 node.prev.next = node.prev
	7.2 node.next.prev = node.next
8. Creaet evitTail func 
	8.1 store the last node into a variable lastNode = lru.tail.prev
	8.2 call removeNode func with lastNode variable
	8.3 delete the cache with the lastNode key delete(lru.cahce, lastNode.key)
9. Create moveToHead func with node parameter
	9.1 call lru.removeNode with ndoe 
	9.2 call lru.addToHead with node
10. Create size func which returns the size of the lruCache
	10.1 acquire lock and unlock at the end
	10.2 return the size of cache with length func len(lru.cache)
11. Create keys func which returns all the keys in the form of string array []string
	11.1 acquire Rlock and unlock at the end
	11.2 create a empty keys array of string with the same size of the lru cache
	11.3 iterate over lru.cache and append the key in keys array
	11.4 return the keys array
	*/

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