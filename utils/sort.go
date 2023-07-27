package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

var _ sort.Interface = (*Sortable[any, string])(nil)

// Sortable wraps a slice of values with a type that implements sort.Interface.
type Sortable[T any, K constraints.Ordered] struct {
	ts        []T
	extractor Extractor[T, K]
}

// Extractor extracts an Ordered value from a type T
type Extractor[T any, K constraints.Ordered] func(T) K

// AsSortable builds a Sortable from a slice of values and an extractor function.
// The extractor function extracts an Ordered value from the slice elements which is used as the comparassion value.
//
// eg. suppose T represents a Person, the extractor function:
// func (p *Person) string { return p.Name }
// would be used to sort all Persons in the slice by their Name.
func AsSortable[T any, K constraints.Ordered](vals []T, extractor Extractor[T, K]) Sortable[T, K] {
	return Sortable[T, K]{
		ts:        vals,
		extractor: extractor,
	}
}

func (s *Sortable[T, K]) Len() int {
	return len(s.ts)
}

func (s *Sortable[T, K]) Swap(i, j int) {
	s.ts[i], s.ts[j] = s.ts[j], s.ts[i]
}

func (s *Sortable[T, K]) Less(i, j int) bool {
	ti, tj := s.ts[i], s.ts[j]
	return s.extractor(ti) < s.extractor(tj)
}

func (s Sortable[T, K]) Sort() {
	sort.Stable(&s)
}
