package set

func New(items ...interface{}) Set {
	st := set{
		elements: make(map[interface{}]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

type Set interface {
	Add(item interface{})
	Delete(item interface{})
	Len() int
	GetItems() []interface{}
	In(item interface{}) bool
	IsSubsetOf(set2 Set) bool
	IsSupersetOf(set2 Set) bool
	Union(set2 Set) Set
	Intersection(set2 Set) Set
	Difference(set2 Set) Set
	SymmetricDifference(set2 Set) Set
}

type set struct {
	elements map[interface{}]bool
}

func (st *set) Add(value interface{}) {
	st.elements[value] = true
}

func (st *set) Delete(value interface{}) {
	delete(st.elements, value)
}

func (st *set) GetItems() []interface{} {
	keys := make([]interface{}, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
	return keys
}

func (st *set) Len() int {
	return len(st.elements)
}

func (st *set) In(value interface{}) bool {
	if _, in := st.elements[value]; in {
		return true
	}
	return false
}

func (st *set) IsSubsetOf(superSet Set) bool {
	if st.Len() > superSet.Len() {
		return false
	}

	for _, item := range st.GetItems() {
		if !superSet.In(item) {
			return false
		}
	}
	return true
}

func (st *set) IsSupersetOf(subSet Set) bool {
	return subSet.IsSubsetOf(st)
}

func (st *set) Union(st2 Set) Set {
	unionSet := New()
	for _, item := range st.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range st2.GetItems() {
		unionSet.Add(item)
	}
	return unionSet
}

func (st *set) Intersection(st2 Set) Set {
	intersectionSet := New()
	var minSet, maxSet Set
	if st.Len() > st2.Len() {
		minSet = st2
		maxSet = st
	} else {
		minSet = st
		maxSet = st2
	}
	for _, item := range minSet.GetItems() {
		if maxSet.In(item) {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

func (st *set) Difference(st2 Set) Set {
	differenceSet := New()
	for _, item := range st.GetItems() {
		if !st2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

func (st *set) SymmetricDifference(st2 Set) Set {
	symmetricDifferenceSet := New()
	dropSet := New()
	for _, item := range st.GetItems() {
		if st2.In(item) {
			dropSet.Add(item)
		} else {
			symmetricDifferenceSet.Add(item)
		}
	}
	for _, item := range st2.GetItems() {
		if !dropSet.In(item) {
			symmetricDifferenceSet.Add(item)
		}
	}
	return symmetricDifferenceSet
}
