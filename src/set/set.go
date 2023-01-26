package set

type Set[T comparable] interface {
	Insert(el T) bool
	Erase(el T)
	Cardinality() int
	Clear()
	Contains(el ...T) bool
	Difference(Set[T]) Set[T]
	Equals(Set[T]) bool
	Intersection(Set[T]) Set[T]
	IsSubset(Set[T]) bool
	IsProperSubset(Set[T]) bool
	IsSuperset(Set[T]) bool
	IsProperSuperset(Set[T]) bool
	Union(Set[T]) Set[T]
	String() string
}

func New[T comparable](vals ...T) Set[T] {
	s := newSet[T]()
	for _, el := range vals {
		s.Insert(el)
	}
	return &s
}
