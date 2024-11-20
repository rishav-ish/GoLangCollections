package collections

type Set struct {
	data map[any]bool
}

func (set *Set) New() {
	set.data = make(map[any]bool)
}

func (set *Set) ContainsElement(element any) bool {
	_, exists := set.data[element]
	return exists
}

func (set *Set) AddElement(element any) {
	if !set.ContainsElement(element) {
		set.data[element] = true
	}
}

func (set *Set) DeleteElement(element any) {
	delete(set.data, element)
}

func (set *Set) Intersect(anotherSet *Set) *Set {
	newSet := &Set{}
	newSet.New()

	for key := range set.data {
		if anotherSet.ContainsElement(key) {
			newSet.AddElement(key)
		}
	}
	return newSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	newSet := &Set{}
	newSet.New()

	for key := range set.data {
		newSet.AddElement(key)
	}

	for key := range anotherSet.data {
		newSet.AddElement(key)
	}
	return newSet
}

func (set *Set) Difference(anotherSet *Set) *Set {
	newSet := &Set{}
	newSet.New()

	for key := range set.data {
		if !anotherSet.ContainsElement(key) {
			newSet.AddElement(key)
		}
	}
	return newSet
}
