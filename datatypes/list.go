package datatypes

import "container/list"

type Element[T any] struct {
	base *list.Element
}

func (e *Element[T]) Next() *Element[T] {
	if next := e.base.Next(); next != nil {
		return &Element[T]{next}
	}
	return nil
}

func (e *Element[T]) Prev() *Element[T] {
	if prev := e.base.Prev(); prev != nil {
		return &Element[T]{prev}
	}
	return nil
}
