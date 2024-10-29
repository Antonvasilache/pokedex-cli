package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T){
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
	{
		key: "https://example.com",
		val: []byte("testdata"),
	},
	{
		key: "https://example.com/path",
		val: []byte("moretestdata"),
	},
}
for i, c := range cases {
	t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T){
		cache := NewCache(interval)
		cache.Add(c.key, c.val)
		val, ok := cache.Get(c.key)
		if !ok{
			t.Errorf("expected to find key")
			return
		}
		if string(val) != string(c.val){
			t.Errorf("expected to find value")
			return
		}
	})
}
}

func TestReapLoop(t *testing.T){
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestEmptyCache(t *testing.T){
	cache := NewCache(5 * time.Second)
	_, ok := cache.Get("nonexistent-key")
	if ok {
		t.Errorf("expected cache miss for empty cache")
	}
}

func TestMultipleAddsSameKey(t *testing.T){
	cache := NewCache(5 * time.Second)
	cache.Add("test-key", []byte("first-value"))
	cache.Add("test-key", []byte("second-value"))

	val, ok := cache.Get("test-key")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	if string(val) != "second-value" {
		t.Errorf("expected second value to override first")
	}
}

func TestConcurrentAccess (t *testing.T){
	cache := NewCache(5 * time.Second)
	const goroutines = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i:=0; i < goroutines; i++ {
		go func(i int){
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			cache.Add(key, []byte("value"))
			_, _ = cache.Get(key)
		}(i)
	}
	wg.Wait()
}