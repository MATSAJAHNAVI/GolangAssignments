package fifo

import (
	cacheinterface "cachee/cache_interface"
	"container/list"
	"errors"
)

type PairF struct {
	Value int
	Ptr   *list.Element
}

type FIFO struct {
	fifoList     *list.List
	fifoMap      map[int]*PairF
	fifoCapacity int
}

func CreateFIFOCache(capacity int) (cacheinterface.Cache, error) {
	if capacity < 0 {
		return nil, errors.New("constructor creating negative capacity cache")
	} else {
		return &FIFO{
			fifoList:     list.New(),
			fifoMap:      make(map[int]*PairF),
			fifoCapacity: capacity,
		}, nil
	}
}

func (f *FIFO) Put(key int, value int) {
	item, present := f.fifoMap[key]

	if !present {
		if f.fifoCapacity == len(f.fifoMap) {
			first := f.fifoList.Front()
			f.fifoList.Remove(first)
			delete(f.fifoMap, first.Value.(int))
		}
		f.fifoMap[key] = &PairF{
			Value: value,
			Ptr:   f.fifoList.PushBack(key),
		}
	} else {
		item.Value = value
		f.fifoMap[key] = item
	}
}

func (f *FIFO) Get(key int) int {
	item, present := f.fifoMap[key]

	if present {
		return item.Value
	}
	return -1
}

// func (f *FIFO) Check() int {
// 	r := len(f.MapF)
// 	return r
// }

// func (f *FIFO) AllCacheElements() {
// 	for _, ele := range f.MapF {
// 		fmt.Println(ele)
// 	}
// }
