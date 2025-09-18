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

type List[T any] struct {
	base *list.List
}

func NewList[T any]() *List[T] {
	return &List[T]{list.New()}
}

func (l *List[T]) Back() *Element[T] {
	if e := l.base.Back(); e != nil {
		return &Element[T]{e}
	}
	return nil
}

func (l *List[T]) Front() *Element[T] {
	if e := l.base.Front(); e != nil {
		return &Element[T]{e}
	}
	return nil
}
