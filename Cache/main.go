package main

import (
	fifo "cachee/Fifo"
	lru "cachee/Lru"
	"fmt"
)

func main() {

	l, _ := lru.CreateLRUCache(4)
	//fmt.Printf("%s \n", msgl)

	l.Put(1, 5)
	l.Put(2, 6)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
	l.Put(3, 7)
	f, _ := fifo.CreateFIFOCache(5)
	//fmt.Printf("%s \n", msgf)

	f.Put(1, 11)
	f.Put(2, 12)
	fmt.Println(f.Get(2))
	f.Put(3, 13)
	f.Put(4, 14)
	f.Put(5, 15)
	f.Put(6, 16)
	fmt.Println(f.Get(1))
	//a := l.Check()
	//fmt.Printf("current length is: %d", a)
	//l.AllCacheElements()

}
