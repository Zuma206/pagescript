package datatypes

import (
	"container/list"
	"iter"
)

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

func (e *Element[T]) Value() T {
	return e.base.Value.(T)
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

func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	if e := l.base.InsertAfter(v, mark.base); e != nil {
		return &Element[T]{e}
	}
	return nil
}

func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	if e := l.base.InsertBefore(v, mark.base); e != nil {
		return &Element[T]{e}
	}
	return nil
}

func (l *List[T]) Len() int {
	return l.base.Len()
}

func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.base.MoveAfter(e.base, mark.base)
}

func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.base.MoveBefore(e.base, mark.base)
}

func (l *List[T]) MoveToBack(e *Element[T]) {
	l.base.MoveToBack(e.base)
}

func (l *List[T]) MoveToFront(e *Element[T]) {
	l.base.MoveToFront(e.base)
}

func (l *List[T]) PushBack(v any) *Element[T] {
	return &Element[T]{l.base.PushBack(v)}
}

func (l *List[T]) PushBackList(other *List[T]) {
	l.base.PushBackList(other.base)
}

func (l *List[T]) PushFront(v any) *Element[T] {
	return &Element[T]{l.base.PushFront(v)}
}

func (l *List[T]) PushFrontList(other *List[T]) {
	l.base.PushFrontList(other.base)
}

func (l *List[T]) Remove(e *Element[T]) T {
	return l.base.Remove(e.base).(T)
}

func (l *List[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := l.base.Front(); e != nil; e = e.Next() {
			if !yield(e.Value.(T)) {
				return
			}
		}
	}
}
