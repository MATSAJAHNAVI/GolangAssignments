package lru

import (
	"testing"
)

func TestLRUPut(t *testing.T) {
	c, _ := CreateLRUCache(3) //.(*LRU)
	l := c.(*LRU)

	type pair struct {
		id    string
		key   int
		value int
	}
	testCases := []pair{
		{id: "tc1", key: 1, value: 21},
		{id: "tc2", key: 2, value: 22},
		{id: "tc3", key: 3, value: 23},
		{id: "tc4", key: 4, value: 24},
	}

	//ranging over testcases
	for _, tc := range testCases {
		l.Put(tc.key, tc.value)
		//to check whether the element moved to front in the list
		li := l.lrulist.Front().Value
		if li != tc.key {
			t.Errorf("error occured at %s expected key of %v got %v ", tc.id, tc.key, li)
		}
		//to check whether the element has been added to map
		m := l.lrumap[tc.key].value
		if m != tc.value {
			t.Errorf("error occured while moving element to map")
		}
	}

}
func TestLRUGet(t *testing.T) {
	c, _ := CreateLRUCache(3)
	l := c.(*LRU)

	type pair struct {
		id    string
		key   int
		value int
	}
	testCases := []pair{
		{id: "tc1", key: 1, value: 21},
		{id: "tc2", key: 2, value: 22},
		{id: "tc3", key: 3, value: 23},
		{id: "tc4", key: 4, value: 24},
	}
	//Ranging over testcases
	for _, tc := range testCases {
		l.Put(tc.key, tc.value)
	}
	//declaring a key that is not present in testcases and checking get returning -1
	check_key := 5
	check_val := l.Get(check_key)
	if check_val != -1 {
		t.Errorf("Returning a value for key which is absent in lru cache")
	}
	//if key is present then check is get function returning correct value
	check_id1 := testCases[1].id
	check_key1 := testCases[1].key
	check_val1 := testCases[1].value
	get_val1 := l.Get(check_key1)
	if get_val1 != check_val1 {
		t.Errorf("error occured with %s : returning false value for a key", check_id1)
	}
	//if key is present check  whether it has moved to front of the list as it is recently used
	check_id2 := testCases[2].id
	check_key2 := testCases[2].key
	//t_val2 := testCases[2].value
	l.Get(check_key2)
	li := l.lrulist.Front().Value
	if li != check_key2 {
		t.Errorf("error occured with %s : element not moving to front of list although it is recently used", check_id2)
	}

}

func TestCreateLRUCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		//want    cacheinterface.Cache
		wantErr bool
	}{
		{name: "test1", args: args{capacity: 3}, wantErr: false},
		{name: "test2", args: args{capacity: -3}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateLRUCache(tt.args.capacity)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateLRUCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("CreateLRUCache() = %v, want %v", got, tt.want)
			// }
		})
	}
}
