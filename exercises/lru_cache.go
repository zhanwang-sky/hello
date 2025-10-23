package exercises

import "container/list"

type lruNode struct {
	key   int
	value int
}

type LRUCache struct {
	cap int
	ls  *list.List
	ha  map[int]*list.Element
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cap: capacity,
		ls:  list.New(),
		ha:  make(map[int]*list.Element),
	}
}

func (lru *LRUCache) access(key, value int) int {
	if ele, prs := lru.ha[key]; prs {
		node := ele.Value.(lruNode)
		oldVal := node.value
		if value >= 0 {
			node.value = value
		}
		lru.ls.Remove(ele)
		lru.ha[key] = lru.ls.PushBack(node)
		return oldVal
	} else {
		if value < 0 {
			return -1
		}
		node := lruNode{key, value}
		lru.ha[key] = lru.ls.PushBack(node)
		if lru.ls.Len() > lru.cap {
			frontEle := lru.ls.Front()
			frontNode := frontEle.Value.(lruNode)
			lru.ls.Remove(frontEle)
			delete(lru.ha, frontNode.key)
		}
		return value
	}
}

func (lru *LRUCache) Get(key int) int {
	return lru.access(key, -1)
}

func (lru *LRUCache) Put(key int, value int) {
	lru.access(key, value)
}
