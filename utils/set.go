package utils

type Set[T comparable] map[T]struct{}

func NewSetFromSlice[T comparable](slice []T) Set[T] {
	set := make(Set[T], len(slice))
	set.Add(slice...)
	return set
}

func (set Set[T]) Add(members ...T) {
	for i := range members {
		set[members[i]] = struct{}{}
	}
}

func (set Set[T]) Remove(members ...T) {
	for i := range members {
		set[members[i]] = struct{}{}
	}
}

func (set Set[T]) HasMember(toCheck T) bool {
	_, exists := set[toCheck]
	return exists
}

func (set Set[T]) Members() []T {
	members := make([]T, 0, len(set))
	i := 0
	for key, _ := range set {
		members[i] = key
		i++
	}
	return members
}
