package lru

import (
	cacheinterface "cachee/cache_interface"
	"container/list"
	"errors"
	//"fmt"
)

type Pair struct {
	value int
	ptr   *list.Element
}

type LRU struct {
	lrulist     *list.List
	lrumap      map[int]*Pair
	lrucapacity int
}

// func CreateLRUCache(capacity int) cacheinterface.Cache {
// 	return &LRU{
// 		lrulist:     list.New(),
// 		lrumap:      make(map[int]*Pair),
// 		lrucapacity: capacity,
// 	}
// }

func CreateLRUCache(capacity int) (cacheinterface.Cache, error) {
	if capacity < 0 {
		return nil, errors.New("could not create negative capacity cache")
	} else {
		return &LRU{
			lrulist:     list.New(),
			lrumap:      make(map[int]*Pair),
			lrucapacity: capacity,
		}, nil
	}
}

func (l *LRU) Put(key int, value int) {
	item, present := l.lrumap[key]

	if !present {
		if l.lrucapacity == len(l.lrumap) {
			last := l.lrulist.Back()
			l.lrulist.Remove(last)
			delete(l.lrumap, last.Value.(int))
		}
		l.lrumap[key] = &Pair{
			value: value,
			ptr:   l.lrulist.PushFront(key),
		}
	} else {
		item.value = value
		l.lrumap[key] = item
		l.lrulist.MoveToFront(item.ptr)
	}
	//return fmt.Sprintf("%d is added to cache", key)
}

func (l *LRU) Get(key int) int {
	item, present := l.lrumap[key]

	if present {
		l.lrulist.MoveToFront(item.ptr)
		return item.value
	}
	return -1
}

// func (l *LRU) Check() int {
// 	r := len(l.Map)
// 	return r
// }

// func (l *LRU) AllCacheElements() {
// 	fmt.Println("{")
// 	for e := l.List.Front(); e != nil; e = e.Next() {
// 		fmt.Print(e.Value)
// 		fmt.Println(" ")
// 	}
// 	fmt.Println("}")
// }
