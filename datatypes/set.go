package datatypes

type Set[T comparable] struct {
	members map[T]Unit
}

func NewSet[T comparable](members ...T) *Set[T] {
	set := &Set[T]{map[T]Unit{}}
	for _, member := range members {
		set.Add(member)
	}
	return set
}

func (set *Set[T]) Add(value T) {
	set.members[value] = Unit{}
}

func (set *Set[T]) Remove(value T) {
	delete(set.members, value)
}

func (set *Set[T]) Has(value T) bool {
	_, has := set.members[value]
	return has
}
