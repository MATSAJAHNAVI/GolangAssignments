package cacheinterface

import (
// "cachee/Lru"
// "container/list"
)

type Cache interface {
	Get(int) int
	Put(int, int)
	//Check() int
	//AllCacheElements()
}

// func Constructor(capacity int, s string) Cache {
// 	// if capacity < 0 {
// 	// 	return nil, "can't create negative capacity cache"
// 	// } else {
// 	switch {
// 	case s == "LRU":
// 		return &lru.LRU{
// 			lrlist:     list.New(),
// 			lrmap:      make(map[int]*Pair),
// 			lrcapacity: capacity,
// 		} //, "created lru cache"
// 	default:
// 		return &FIFO{
// 			ListF:     list.New(),
// 			MapF:      make(map[int]*PairF),
// 			CapacityF: capacity,
// 		} //, "created fifo cache"
// 	}
// 	//}
// }

// func Check(c Cache) string {
// 	fmt.Println(c.Capacity)
// }

// func (c *Cache) AddData(key int, value interface{}) string {
// 	return c.Put(key, value)
// }

// func RetrieveData(c Cache, key int) interface{} {
// 	return c.Get(key)
// }
