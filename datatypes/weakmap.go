package datatypes

import (
	"runtime"
	"weak"
)

type WeakMapEntry[T any] struct {
	value   T
	cleanup runtime.Cleanup
}

type WeakMap[K any, T any] struct {
	values map[weak.Pointer[K]]WeakMapEntry[T]
}

func NewWeakMap[K any, T any]() *WeakMap[K, T] {
	return &WeakMap[K, T]{
		values: map[weak.Pointer[K]]WeakMapEntry[T]{},
	}
}

func (weakMap *WeakMap[K, T]) Set(keyPtr *K, value T) {
	weakKeyPtr := weak.Make(keyPtr)
	weakMap.values[weakKeyPtr] = WeakMapEntry[T]{
		cleanup: runtime.AddCleanup(keyPtr, func(_ Unit) {
			delete(weakMap.values, weakKeyPtr)
		}, Unit{}),
		value: value,
	}
}

func (weakMap *WeakMap[K, T]) Get(keyPtr *K) (T, bool) {
	var nilvalue T
	entry, ok := weakMap.values[weak.Make(keyPtr)]
	if !ok {
		return nilvalue, false
	}
	return entry.value, true
}

func (weakMap *WeakMap[K, T]) Delete(keyPtr *K) bool {
	weakKeyPtr := weak.Make(keyPtr)
	entry, ok := weakMap.values[weakKeyPtr]
	if !ok {
		return false
	}
	entry.cleanup.Stop()
	delete(weakMap.values, weakKeyPtr)
	return true
}

func (weakMap *WeakMap[K, T]) Len() int {
	return len(weakMap.values)
}
