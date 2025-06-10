package containers

// Credits: Make p set from  github.com/badgerodon/collections generic

type (
	Set[T comparable] struct {
		hash map[T]nothing
	}

	nothing struct{}
)

// Create a new set
func NewSet[T comparable](initial ...T) *Set[T] {
	s := &Set[T]{make(map[T]nothing)}
	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

// Transforms Set to Slice
func (p *Set[T]) ToSlice() *Slice[T] {
	n := []T{}
	for k := range p.hash {
		n = append(n, k)
	}

	return NewSlice(n...)
}

// Get the sets values
func (p *Set[T]) Values() []T {
	n := []T{}
	for k := range p.hash {
		n = append(n, k)
	}

	return n
}

// Find the difference between two sets
func (p *Set[T]) Difference(set *Set[T]) *Set[T] {
	n := make(map[T]nothing)
	for k := range p.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = nothing{}
		}
	}

	return &Set[T]{n}
}

// Do calls f for each item in the set
func (p *Set[T]) Do(f func(T)) {
	for k := range p.hash {
		f(k)
	}
}

// ForEach calls the given function for each element in the set.
//
// The function takes a parameter 'f' of type 'func(T) error'. It should
// receive an element of type 'T' and return an error if any.
//
// It returns an error if the function 'f' returns an error during iteration.
// Otherwise, it returns nil.
func (p *Set[T]) ForEach(f func(T) error) error {
	for k := range p.hash {
		if err := f(k); err != nil {
			return err
		}
	}

	return nil
}

// Has checks if the given element exists in the Set.
//
// Parameters:
// - element: the element to check for existence in the Set.
//
// Returns:
// - bool: true if the element exists in the Set, false otherwise.
func (p *Set[T]) Has(element T) bool {
	_, exists := p.hash[element]
	return exists
}

// Test if Set has any Items
func (p *Set[T]) HasItems() bool {
	return len(p.hash) > 0
}

// Add one or more elements to the set
func (p *Set[T]) Insert(elements ...T) {
	for _, k := range elements {
		p.hash[k] = nothing{}
	}
}

// Find the intersection of two sets
func (p *Set[T]) Intersection(set *Set[T]) *Set[T] {
	n := make(map[T]nothing)

	for k := range p.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = nothing{}
		}
	}

	return &Set[T]{n}
}

// Return the number of items in the set
func (p *Set[T]) Len() int {
	return len(p.hash)
}

// Test whether or not p set is a proper subset of "set"
func (p *Set[T]) ProperSubsetOf(set *Set[T]) bool {
	return p.SubsetOf(set) && p.Len() < set.Len()
}

// Remove given elements from the set
func (p *Set[T]) Remove(elements ...T) {
	for _, k := range elements {
		delete(p.hash, k)
	}
}

// Clear removes all elements from the set.
func (p *Set[T]) Clear() {
	p.hash = make(map[T]nothing)
}

// Remove a set from the set
func (p *Set[T]) RemoveSet(set *Set[T]) {
	for _, k := range set.Values() {
		delete(p.hash, k)
	}
}

// Test whether or not this set is a subset of "set"
func (p *Set[T]) SubsetOf(set *Set[T]) bool {
	if p.Len() > set.Len() {
		return false
	}

	for k := range p.hash {
		if _, exists := set.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Find the union of two sets
func (p *Set[T]) Union(set *Set[T]) *Set[T] {
	n := make(map[T]nothing)
	for k := range p.hash {
		n[k] = nothing{}
	}
	for k := range set.hash {
		n[k] = nothing{}
	}

	return &Set[T]{n}
}
