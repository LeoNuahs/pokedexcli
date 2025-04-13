package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct{
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2", 
			val: []byte("Sample Data"),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area", 
			val: []byte("Example Data"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T){
			cache := NewCache(interval)
			cache.Add(c.key, c.val)

			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key: %v", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value: %v\ninstead got value: %v", c.val, val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const intervalTime = 5 * time.Millisecond
	const waitTime = intervalTime + 5 * time.Millisecond

	cache := NewCache(intervalTime)
	cache.Add("https://pokeapi.co/api/v2/location-area?offset=20&limit=20", []byte("Testing Data"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area?offset=20&limit=20")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	
	time.Sleep(waitTime)
	
	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area?offset=20&limit=20")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
