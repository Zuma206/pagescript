package datatypes

import (
	"math/rand"
	"runtime"
	"testing"
)

const nValues = 500

func TestGettingAndSetting(t *testing.T) {
	weakMap := NewWeakMap[bool, int]()
	values := map[*bool]int{}
	for range nValues {
		key := true
		value := rand.Int()
		values[&key] = value
		weakMap.Set(&key, value)
	}
	for key, expected := range values {
		actual, ok := weakMap.Get(key)
		if !ok {
			t.Errorf("no value set for %p", key)
		} else if actual != expected {
			t.Errorf("key %p was set to %q rather when %q was expected", key, actual, expected)
		}
	}
}

func TestGarbageCollection(t *testing.T) {
	weakMap := NewWeakMap[bool, int]()
	{
		key := true
		weakMap.Set(&key, rand.Int())
		if length := weakMap.Len(); length != 1 {
			t.Errorf("weakmap has length %d, when a length of 1 was expected", length)
		}
	}
	runtime.GC()
	if length := weakMap.Len(); length != 0 {
		t.Errorf("weakmap has length %d, when a length of 0 was expected", length)
	}
}
