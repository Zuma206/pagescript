package datatypes

import (
	"math/rand"
	"testing"
)

const nValues = 500

func TestGettingAndSetting(t *testing.T) {
	weakMap := NewWeakMap[Unit, int]()
	values := map[*Unit]int{}
	for range nValues {
		key := &Unit{}
		value := rand.Int()
		values[key] = value
		weakMap.Set(key, value)
	}
	for key, expected := range values {
		actual, ok := weakMap.Get(key)
		if !ok {
			t.Errorf("no value set for %q", key)
		} else if actual != expected {
			t.Errorf("key %q was set to %q rather when %q was expected", key, actual, expected)
		}
	}
}
