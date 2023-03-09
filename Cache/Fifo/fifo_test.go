package fifo

import (
	"testing"
)

func TestFIFOPut(t *testing.T) {
	//f := CreateFIFOCache(3).(*FIFO)
	c, _ := CreateFIFOCache(3)
	f := c.(*FIFO)
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

	for _, tc := range testCases {
		f.Put(tc.key, tc.value)
		//check whether element moved to back of list
		li := f.fifoList.Back().Value
		if li != tc.key {
			t.Errorf("error occured with %s expected a key of %v but got %v", tc.id, tc.key, li)
		}
		//check whether element moved into map
		m := f.fifoMap[tc.key].Value
		if m != tc.value {
			t.Errorf("error occured while moving element to map")
		}
	}

}

func TestFIFOGet(t *testing.T) {
	//f := CreateFIFOCache(3).(*FIFO)
	c, _ := CreateFIFOCache(3)
	f := c.(*FIFO)

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

	for _, tc := range testCases {
		f.Put(tc.key, tc.value)
	}

	//declaring a key that is not present in testcases and checking get returning -1
	check_key := 5
	check_val := f.Get(check_key)
	if check_val != -1 {
		t.Errorf("Returning a value for key which is absent in fifo cache")
	}

	//if key is present then check is get function returning correct value
	check_id1 := testCases[1].id
	check_key1 := testCases[1].key
	check_val1 := testCases[1].value
	get_val1 := f.Get(check_key1)
	if get_val1 != check_val1 {
		t.Errorf("error occured with %s : returning false value for a key", check_id1)
	}
}

func TestCreateFIFOCache(t *testing.T) {
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
			_, err := CreateFIFOCache(tt.args.capacity)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFIFOCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("CreateFIFOCache() = %v, want %v", got, tt.want)
			// }
		})
	}
}
